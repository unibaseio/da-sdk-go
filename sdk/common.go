package sdk

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/unibaseio/da-sdk-go/build"
	"github.com/unibaseio/da-sdk-go/lib/bls"
	"github.com/unibaseio/da-sdk-go/lib/bls/erasure"
	"github.com/unibaseio/da-sdk-go/lib/env"
	"github.com/unibaseio/da-sdk-go/lib/log"
	"github.com/unibaseio/da-sdk-go/lib/types"
	"github.com/unibaseio/da-sdk-go/lib/utils"

	"github.com/consensys/gnark-crypto/hash"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mitchellh/go-homedir"
	"github.com/schollz/progressbar/v3"
	"golang.org/x/crypto/sha3"
)

var logger = log.Logger("sdk")

// ChainType is the chain used when CHAIN_TYPE is unset. CheckENV exports it, and
// CLI flags reading CHAIN_TYPE see it before their own default, so it must be
// the live chain: DA anchors on Base (BSC testnets are retired).
var ChainType = build.BaseSepolia
var chaintype = ""

func init() {
	// local test
	//os.Setenv("CHAIN_TYPE", ChainType)
	CheckENV()
	// was os.Getenv("LogLevel") — wrong name (the log pkg reads LOG_LEVEL),
	// so this never took effect; use the same key now.
	log.SetLogLevel(env.Str(env.LogLevel, "DEBUG"))
}

var ServerURL = build.ServerURL

// MetaTimeout bounds the small metadata calls a node makes in its polling
// loops (edge list, dispatch list, slot request), so one unresponsive stream
// or gateway cannot stall the loop.
var MetaTimeout = 30 * time.Second

const InHashID = hash.MIMC_BW6_761

func CheckENV() {
	if env.Str(env.ChainType, "") == "" {
		os.Setenv(env.ChainType, ChainType)
	}
	chaintype = build.CheckChain()
	logger.Warn("connect to chain: ", chaintype)
}

// MaxPieceSize is the most raw bytes one piece holds under policy p (what the
// stream reads per piece; every piece but the last of a file is this size).
func MaxPieceSize(p types.Policy) int64 {
	if p.K == 0 {
		return 0
	}
	k := int64(p.K)
	return (int64(bls.MaxSize) / (bls.UnPadSize * k)) * k * bls.UnPadSize
}

// CheckFileFullShape checks the structure of a stream's upload answer before
// anything is decoded or evaluated: a supported policy (equal to want unless
// want is the zero Policy), one proof and one size per piece, every witness
// holding exactly N commits and K move/limit commits and claimed values, piece
// sizes within a piece and summing to Size, and a well-formed sha256 Hash.
// It does not read the file; CheckFileFull does that.
func CheckFileFullShape(ff types.FileFull, want types.Policy) error {
	pol := ff.Policy
	if err := pol.Check(); err != nil {
		return err
	}
	if pol.K == 0 || pol.N <= pol.K {
		return fmt.Errorf("bad rs policy %d/%d", pol.N, pol.K)
	}
	if want != (types.Policy{}) && pol != want {
		return fmt.Errorf("stream encoded with policy %d/%d, requested %d/%d", pol.N, pol.K, want.N, want.K)
	}
	if len(ff.Proofs) != len(ff.Pieces) || len(ff.PieceSizes) != len(ff.Pieces) {
		return fmt.Errorf("receipt has %d pieces, %d proofs, %d sizes", len(ff.Pieces), len(ff.Proofs), len(ff.PieceSizes))
	}
	if hb, err := hex.DecodeString(ff.Hash); err != nil || len(hb) != sha256.Size {
		return fmt.Errorf("receipt hash %q is not a sha256", ff.Hash)
	}
	if ff.Size < 0 {
		return fmt.Errorf("negative file size %d", ff.Size)
	}
	maxp := MaxPieceSize(pol)
	total := int64(0)
	for i, ps := range ff.PieceSizes {
		if ps <= 0 || ps > maxp {
			return fmt.Errorf("piece %d has size %d, want 1..%d", i, ps, maxp)
		}
		total += ps
		if total > ff.Size {
			return fmt.Errorf("piece sizes exceed the file size %d", ff.Size)
		}
		if err := bls.CheckEncodeWitnessShape(ff.Proofs[i], int(pol.N), int(pol.K)); err != nil {
			return fmt.Errorf("piece %d witness: %w", i, err)
		}
	}
	if total != ff.Size {
		return fmt.Errorf("pieces cover %d bytes, file is %d", total, ff.Size)
	}
	return nil
}

// CheckFileFull verifies a stream's upload answer against the local file at
// fp: see CheckFileFullPolicy. Any supported policy is accepted; callers that
// know the policy they asked for should use CheckFileFullPolicy.
func CheckFileFull(ff types.FileFull, stream common.Address, fp string) ([]types.PieceCore, error) {
	return CheckFileFullPolicy(ff, stream, fp, types.Policy{})
}

// CheckFileFullPolicy verifies that the stream encoded exactly the file at fp
// under policy want (zero Policy: any supported one) and returns the pieces to
// register. Besides the per-piece encoding check (Eval(data)==ClaimedValues at
// the Fiat-Shamir point, Σ MoveCommits==Root, RS-valid parity commits) it
// requires each piece name to be its witness root, the pieces to cover the
// whole file with nothing left over, and Size and Hash to be the file's own.
func CheckFileFullPolicy(ff types.FileFull, stream common.Address, fp string, want types.Policy) ([]types.PieceCore, error) {
	logger.Debug("check stream handle of file: ", fp)
	if err := CheckFileFullShape(ff, want); err != nil {
		return nil, err
	}
	p, err := homedir.Expand(fp)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	st, err := f.Stat()
	if err != nil {
		return nil, err
	}
	if !st.Mode().IsRegular() {
		return nil, fmt.Errorf("%s is not a regular file", fp)
	}
	if st.Size() != ff.Size {
		return nil, fmt.Errorf("receipt size %d, file is %d bytes", ff.Size, st.Size())
	}

	h := sha256.New()
	fi := io.TeeReader(f, h)

	res := make([]types.PieceCore, len(ff.Pieces))
	var rnd bls.Fr
	for i := 0; i < len(ff.Pieces); i++ {
		ew := new(bls.EncodeWitness)
		err := ew.Deserialize(ff.Proofs[i])
		if err != nil {
			return nil, err
		}

		err = CheckWitness(int(ff.Policy.N), int(ff.Policy.K), ew)
		if err != nil {
			return nil, err
		}

		// the encoder's accumulated opening binds the move/limit commitments
		// (hence the piece root) to the claimed values checked below
		vk, err := encodingVerifyKey()
		if err != nil {
			return nil, err
		}
		shardElems := 1 + (ff.PieceSizes[i]-1)/(31*int64(ff.Policy.K))
		if err := ew.VerifyOpening(vk, stream.Bytes(), int(shardElems)); err != nil {
			return nil, fmt.Errorf("piece %d: %w", i, err)
		}

		root := ew.Root.Bytes()
		name := hex.EncodeToString(root[:])
		if !strings.EqualFold(name, ff.Pieces[i]) {
			return nil, fmt.Errorf("piece %d is named %s, its witness root is %s", i, ff.Pieces[i], name)
		}

		// Fiat-Shamir point — shared with the encoder via bls.Challenge so the
		// transcript byte-layout can't drift between encode and verify.
		point := ew.Challenge(stream.Bytes())
		rnd.SetBytes(point)

		slen := (1 + (ff.PieceSizes[i]-1)/(31*int64(ff.Policy.K))) * 31
		rest := ff.PieceSizes[i]
		for j := 0; j < int(ff.Policy.K); j++ {
			size := slen
			if rest < slen {
				size = rest
			}
			buf := make([]byte, size)
			if _, err := io.ReadFull(fi, buf); err != nil {
				return nil, fmt.Errorf("read piece %d shard %d: %w", i, j, err)
			}
			rest -= size

			cval := bls.Eval(bls.Split(31, buf), rnd)
			if cval.Cmp(&ew.ClaimedValues[j]) != 0 {
				return nil, fmt.Errorf("unequal val at %d %d", i, j)
			}
		}

		res[i] = types.PieceCore{
			Policy:   ff.Policy,
			Name:     name,
			Size:     ff.PieceSizes[i],
			Streamer: stream,
		}
	}

	// the pieces must cover the whole file, and the receipt hash must be its own
	if n, _ := io.Copy(io.Discard, fi); n != 0 {
		return nil, fmt.Errorf("%d bytes of %s are not in any piece", n, fp)
	}
	if got := hex.EncodeToString(h.Sum(nil)); !strings.EqualFold(got, ff.Hash) {
		return nil, fmt.Errorf("receipt hash %s, file hashes to %s", ff.Hash, got)
	}
	return res, nil
}

// encodingVerifyKey is the KZG verifying key encoding openings are checked
// with: the shared SRS's (tests swap in a key of their own).
var encodingVerifyKey = bls.SRSVerifyKey

func CheckWitness(rsn, rsk int, ew *bls.EncodeWitness) error {
	if len(ew.Commits) != rsn {
		return fmt.Errorf("invalid commit count")
	}

	if len(ew.MoveCommits) != rsk {
		return fmt.Errorf("invalid move commit count")
	}

	if len(ew.LimitCommits) != rsk {
		return fmt.Errorf("invalid limit commit count")
	}

	if len(ew.ClaimedValues) != rsk {
		return fmt.Errorf("invalid proof value count")
	}

	dv := make([][]byte, rsk)
	var sum bls.G1
	for i := 0; i < rsk; i++ {
		dv[i] = ew.Commits[i].Marshal()
		sum.Add(&sum, &ew.MoveCommits[i])
	}

	if !sum.Equal(&ew.Root) {
		return fmt.Errorf("unequal root")
	}

	need := make([]int, rsn-rsk)
	pv := make([][]byte, rsn-rsk)
	for i := rsk; i < rsn; i++ {
		need[i-rsk] = i
		pv[i-rsk] = ew.Commits[i].Marshal()
	}

	rs, err := erasure.NewRS(rsn, rsk)
	if err != nil {
		return err
	}
	return rs.Check(dv, pv, need)
}

func DecodeAuth(authstr string) (types.Auth, error) {
	au := types.Auth{}
	if authstr == "" {
		return au, fmt.Errorf("nil authorization")
	}

	ab, err := hex.DecodeString(authstr)
	if err == nil {
		err = json.Unmarshal(ab, &au)
		if err != nil {
			return au, err
		}
	} else {
		type JSAuth struct {
			Type string
			Addr common.Address
			Time int64
			Hash string
			Sign string
			Msg  string
		}
		jau := JSAuth{}
		err := json.Unmarshal([]byte(authstr), &jau)
		if err != nil {
			return au, err
		}

		au.Type = jau.Type
		au.Addr = jau.Addr
		au.Time = jau.Time
		au.Msg = jau.Msg
		if strings.HasPrefix(jau.Hash, "0x") {
			au.Hash, err = hex.DecodeString(jau.Hash[2:])
			if err != nil {
				return au, err
			}
		}

		if strings.HasPrefix(jau.Sign, "0x") {
			au.Sign, err = hex.DecodeString(jau.Sign[2:])
			if err != nil {
				return au, err
			}
		}
	}
	return au, nil
}

func VerifyAuth(au types.Auth) error {
	if len(au.Sign) == 65 {
		recoveryID := int(au.Sign[64])
		if recoveryID >= 27 && recoveryID <= 34 {
			recoveryID -= 27
		} else if recoveryID >= 35 && recoveryID <= 38 {
			recoveryID -= 35
		}
		au.Sign[64] = byte(recoveryID)
	}

	b := make([]byte, len(au.Hash)+8)
	copy(b, au.Hash)
	binary.BigEndian.PutUint64(b[len(au.Hash):], uint64(au.Time))
	var sum []byte
	switch au.Type {
	case "personal_sign":
		hash := sha3.NewLegacyKeccak256()
		hash.Write([]byte{0x19})
		hash.Write([]byte("Ethereum Signed Message:"))
		hash.Write([]byte{0x0A})
		hash.Write([]byte(strconv.Itoa(len(b))))
		hash.Write(b)
		sum = hash.Sum(nil)
	default:
		sums := sha256.Sum256(b)
		sum = sums[:]
	}

	rePub, err := crypto.Ecrecover(sum, au.Sign)
	if err != nil {
		return err
	}

	if !bytes.Equal(au.Addr.Bytes(), utils.ToEthAddress(rePub)) {
		return fmt.Errorf("invalid auth %s", au.Addr)
	}

	return nil
}

// personalSignDigest computes the EIP-191 personal_sign digest of msg:
// keccak256("\x19Ethereum Signed Message:\n" + len(msg) + msg). Same construction
// VerifyAuth uses for the "personal_sign" type, factored for reuse by VerifySIWE.
func personalSignDigest(msg []byte) []byte {
	h := sha3.NewLegacyKeccak256()
	h.Write([]byte{0x19})
	h.Write([]byte("Ethereum Signed Message:"))
	h.Write([]byte{0x0A})
	h.Write([]byte(strconv.Itoa(len(msg))))
	h.Write(msg)
	return h.Sum(nil)
}

var (
	siweAddrRe   = regexp.MustCompile(`(?m)^0x[0-9a-fA-F]{40}$`)
	siweIssuedRe = regexp.MustCompile(`(?mi)^Issued At:[ \t]*(.+?)[ \t]*$`)
	siweExpRe    = regexp.MustCompile(`(?mi)^Expiration Time:[ \t]*(.+?)[ \t]*$`)
	siweNbfRe    = regexp.MustCompile(`(?mi)^Not Before:[ \t]*(.+?)[ \t]*$`)
	siweDomainRe = regexp.MustCompile(`^(\S+) wants you to sign in with your Ethereum account:`)
	siweChainRe  = regexp.MustCompile(`(?mi)^Chain ID:[ \t]*(.*?)[ \t]*$`)
	siweURIRe    = regexp.MustCompile(`(?mi)^URI:[ \t]*(.*?)[ \t]*$`)
)

// siweTime parses an optional RFC3339 field; ok=false when it is absent.
func siweTime(re *regexp.Regexp, msg string) (t time.Time, ok bool, err error) {
	m := re.FindStringSubmatch(msg)
	if m == nil {
		return time.Time{}, false, nil
	}
	t, err = time.Parse(time.RFC3339, strings.TrimSpace(m[1]))
	return t, true, err
}

// siweDomain returns the domain an EIP-4361 message was issued for (the
// first line's "<domain> wants you to sign in ..."), "" if it has none.
func siweDomain(msg string) string {
	if m := siweDomainRe.FindStringSubmatch(msg); m != nil {
		return strings.ToLower(m[1])
	}
	return ""
}

// parseSIWE pulls the two security-relevant fields out of a SIWE / EIP-4361
// message: the account address line (^0x…40 hex…$) and the "Issued At:" RFC3339
// timestamp. The verifier does NOT reconstruct the message (it checks the
// received bytes), so the rest of the text is free-form/human-readable.
func parseSIWE(msg string) (addr string, issuedAt int64, err error) {
	a := siweAddrRe.FindString(msg)
	if a == "" {
		return "", 0, fmt.Errorf("siwe: missing account address line")
	}
	m := siweIssuedRe.FindStringSubmatch(msg)
	if m == nil {
		return "", 0, fmt.Errorf("siwe: missing 'Issued At'")
	}
	t, perr := time.Parse(time.RFC3339, strings.TrimSpace(m[1]))
	if perr != nil {
		return "", 0, fmt.Errorf("siwe: bad 'Issued At' %q: %w", m[1], perr)
	}
	return strings.ToLower(a), t.Unix(), nil
}

// DefaultAuthDrift is how far (seconds) a signed timestamp may be from now
// before VerifyAuthFresh rejects it.
const DefaultAuthDrift int64 = 600

// VerifyAuthFresh verifies au's signature and that the timestamp bound into
// it (au.Time for the legacy Hash||be64(Time) bytes, "Issued At" for SIWE) is
// within drift seconds of now, so a captured header cannot be replayed later.
// The freshness check uses the signed timestamp, so editing the envelope
// cannot widen the window.
func VerifyAuthFresh(au types.Auth, drift int64) error {
	var signedAt int64
	if len(au.Msg) > 0 {
		t, err := VerifySIWE(au)
		if err != nil {
			return fmt.Errorf("verify auth: %w", err)
		}
		signedAt = t
	} else {
		signedAt = au.Time
		if err := VerifyAuth(au); err != nil {
			return fmt.Errorf("verify auth: %w", err)
		}
	}

	delta := time.Now().Unix() - signedAt
	if delta < 0 {
		delta = -delta
	}
	if delta > drift {
		return fmt.Errorf("auth timestamp out of window (delta=%ds, max=%ds)", delta, drift)
	}
	return nil
}

// VerifySIWE verifies a human-readable EIP-4361 / SIWE auth. au.Sign must be a
// personal_sign over the EXACT au.Msg text and recover au.Addr, and the account
// address embedded in au.Msg must equal au.Addr. It returns the message's
// "Issued At" as a unix timestamp so the caller can enforce its own freshness
// window. Security model is identical to VerifyAuth (prove control of Addr +
// a fresh in-signature timestamp); only the signed bytes are readable.
func VerifySIWE(au types.Auth) (int64, error) {
	if len(au.Msg) == 0 {
		return 0, fmt.Errorf("siwe: empty message")
	}
	if len(au.Sign) != 65 {
		return 0, fmt.Errorf("siwe: bad signature length %d", len(au.Sign))
	}
	sig := make([]byte, 65)
	copy(sig, au.Sign)
	if sig[64] >= 27 && sig[64] <= 34 {
		sig[64] -= 27
	} else if sig[64] >= 35 && sig[64] <= 38 {
		sig[64] -= 35
	}

	rePub, err := crypto.Ecrecover(personalSignDigest([]byte(au.Msg)), sig)
	if err != nil {
		return 0, err
	}
	if !bytes.Equal(au.Addr.Bytes(), utils.ToEthAddress(rePub)) {
		return 0, fmt.Errorf("siwe: signature does not match %s", au.Addr)
	}

	addrInMsg, issuedAt, err := parseSIWE(au.Msg)
	if err != nil {
		return 0, err
	}
	if addrInMsg != strings.ToLower(au.Addr.Hex()) {
		return 0, fmt.Errorf("siwe: message address %s != envelope %s", addrInMsg, au.Addr)
	}
	now := time.Now()
	if exp, ok, err := siweTime(siweExpRe, au.Msg); err != nil {
		return 0, fmt.Errorf("siwe: bad 'Expiration Time': %w", err)
	} else if ok && !now.Before(exp) {
		return 0, fmt.Errorf("siwe: message expired at %s", exp.Format(time.RFC3339))
	}
	if nbf, ok, err := siweTime(siweNbfRe, au.Msg); err != nil {
		return 0, fmt.Errorf("siwe: bad 'Not Before': %w", err)
	} else if ok && now.Before(nbf) {
		return 0, fmt.Errorf("siwe: message not valid before %s", nbf.Format(time.RFC3339))
	}
	return issuedAt, nil
}

// VerifyAuthFreshDomains is VerifyAuthFresh that, for a SIWE message, also
// requires the domain it was issued for to be one of domains (and its URI to
// point at one of them). Without this a sign-in message any other site
// obtained from the user would be accepted here. An empty domains list skips
// the domain check. See VerifyAuthFreshPolicy for the full set of checks.
func VerifyAuthFreshDomains(au types.Auth, drift int64, domains []string) error {
	return VerifyAuthFreshPolicy(au, drift, SIWEPolicy{Domains: domains})
}

// SIWEPolicy is what a service requires of a sign-in message (au.Msg) on top
// of a valid signature and a fresh "Issued At".
type SIWEPolicy struct {
	// Domains the message must be issued for ("<domain> wants you to sign
	// in ..."); its "URI:" must then also be an http(s) URI on one of them.
	// Empty: no domain or URI check (unless RequireDomains).
	Domains []string
	// RequireDomains refuses every sign-in message while Domains is empty:
	// for services that are not signed in to from a website (the DA gateway
	// and stream) a message bound to no site must not be accepted.
	RequireDomains bool
	// ChainIDs, when set, are the chains a message may name: a SIWE message
	// must carry exactly one "Chain ID:" among them; any other message that
	// carries one must also match.
	ChainIDs []int64
}

// The SIWE "Nonce:" is not checked. A nonce only stops a replay if the
// service issued it and accepts it once, which needs a challenge endpoint and
// clients that fetch a nonce before signing: a protocol change for every
// client, deferred. Until then a captured sign-in can be replayed for as long
// as its "Issued At" is within the drift window (and its Expiration Time, if
// any, has not passed).

// VerifyAuthFreshPolicy is VerifyAuthFresh plus policy p for a sign-in
// message. Envelopes without a message (the binary hash||time formats) are
// not affected by p.
func VerifyAuthFreshPolicy(au types.Auth, drift int64, p SIWEPolicy) error {
	if err := VerifyAuthFresh(au, drift); err != nil {
		return err
	}
	if len(au.Msg) == 0 {
		return nil
	}
	if err := checkSIWEPolicy(au.Msg, p); err != nil {
		return fmt.Errorf("verify auth: %w", err)
	}
	return nil
}

func checkSIWEPolicy(msg string, p SIWEPolicy) error {
	domains := make([]string, 0, len(p.Domains))
	for _, d := range p.Domains {
		if d = strings.ToLower(strings.TrimSpace(d)); d != "" {
			domains = append(domains, d)
		}
	}
	if len(domains) == 0 && p.RequireDomains {
		return fmt.Errorf("sign-in messages are not accepted by this service (no domain allowlist configured)")
	}
	domain := siweDomain(msg)

	if len(p.ChainIDs) > 0 {
		ids := siweChainRe.FindAllStringSubmatch(msg, -1)
		switch {
		case len(ids) > 1:
			return fmt.Errorf("siwe message has %d 'Chain ID' lines", len(ids))
		case len(ids) == 0 && domain != "":
			return fmt.Errorf("siwe message has no 'Chain ID'")
		case len(ids) == 1:
			id, err := strconv.ParseInt(ids[0][1], 10, 64)
			if err != nil {
				return fmt.Errorf("siwe: bad 'Chain ID' %q", ids[0][1])
			}
			if !slices.Contains(p.ChainIDs, id) {
				return fmt.Errorf("siwe message is for chain %d, this service is on %v", id, p.ChainIDs)
			}
		}
	}

	if len(domains) == 0 {
		return nil
	}
	if domain == "" || !slices.Contains(domains, domain) {
		return fmt.Errorf("siwe message issued for %q, not for this service", domain)
	}
	uris := siweURIRe.FindAllStringSubmatch(msg, -1)
	if len(uris) != 1 {
		return fmt.Errorf("siwe message has %d 'URI' lines, want 1", len(uris))
	}
	u, err := url.Parse(uris[0][1])
	if err != nil || (u.Scheme != "https" && u.Scheme != "http") || u.Host == "" {
		return fmt.Errorf("siwe: bad 'URI' %q", uris[0][1])
	}
	if !slices.Contains(domains, strings.ToLower(u.Host)) {
		return fmt.Errorf("siwe message URI %q is not on an allowed domain", uris[0][1])
	}
	return nil
}

// hash is random byte now
func BuildAuth(addr, privk string, hash []byte) types.Auth {
	h := sha256.New()
	ts := time.Now().Unix()
	h.Write(hash)
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(ts))
	h.Write(b)

	sum := h.Sum(nil)

	pb, _ := crypto.HexToECDSA(privk)

	sign, _ := crypto.Sign(sum[:], pb)

	au := types.Auth{
		Addr: utils.HexToAddress(addr),
		Time: ts,
		Hash: hash,
		Sign: sign,
	}
	return au
}

func doRequest(ctx context.Context, baseUrl, method, ctype string, au types.Auth, r io.Reader) ([]byte, error) {
	return doRequestLimit(ctx, baseUrl, method, ctype, au, r, -1)
}

// doRequestLimit is doRequest reading at most max bytes of the response body
// (max < 0: unbounded). A longer body is an error.
func doRequestLimit(ctx context.Context, baseUrl, method, ctype string, au types.Auth, r io.Reader, max int64) ([]byte, error) {
	haddr := baseUrl + method
	hreq, err := http.NewRequestWithContext(ctx, "POST", haddr, r)
	if err != nil {
		return nil, err
	}

	if len(au.Sign) > 0 {
		aub, err := json.Marshal(au)
		if err != nil {
			return nil, err
		}
		hreq.Header.Add("Authorization", hex.EncodeToString(aub))
	}

	if ctype == "" {
		ctype = "application/x-www-form-urlencoded"
	}
	hreq.Header.Add("Content-Type", ctype)

	defaultHTTPClient := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			ForceAttemptHTTP2:     true,
			WriteBufferSize:       16 << 10, // 16KiB moving up from 4KiB default
			ReadBufferSize:        16 << 10, // 16KiB moving up from 4KiB default
			MaxIdleConns:          100,
			MaxIdleConnsPerHost:   100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			DisableCompression:    true,
		},
	}

	bar := progressbar.DefaultBytes(-1, baseUrl+" download:")

	resp, err := defaultHTTPClient.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var body io.Reader = resp.Body
	if max >= 0 {
		body = io.LimitReader(resp.Body, max+1)
	}
	pr := progressbar.NewReader(body, bar)
	res, err := io.ReadAll(&pr)
	if err != nil {
		return nil, err
	}
	bar.Finish()
	if max >= 0 && int64(len(res)) > max {
		return nil, fmt.Errorf("response from %s exceeds %d bytes", haddr, max)
	}

	// Accept any 2xx — the clean /v1 writes use proper REST codes (POST /v1/files
	// and /v1/edges return 201 Created), not just 200.
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("response: %s, msg: %s", resp.Status, res)
	}

	return res, nil
}

func Get(ctx context.Context, baseUrl string) ([]byte, error) {
	haddr := baseUrl
	hreq, err := http.NewRequestWithContext(ctx, "GET", haddr, nil)
	if err != nil {
		return nil, err
	}

	hreq.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	defaultHTTPClient := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			ForceAttemptHTTP2:     true,
			WriteBufferSize:       16 << 10, // 16KiB moving up from 4KiB default
			ReadBufferSize:        16 << 10, // 16KiB moving up from 4KiB default
			MaxIdleConns:          100,
			MaxIdleConnsPerHost:   100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			DisableCompression:    true,
		},
	}

	bar := progressbar.DefaultBytes(-1, baseUrl+" download:")

	resp, err := defaultHTTPClient.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	pr := progressbar.NewReader(resp.Body, bar)
	res, err := io.ReadAll(&pr)
	if err != nil {
		return nil, err
	}
	bar.Finish()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("response: %s, msg: %s", resp.Status, res)
	}

	return res, nil
}

// unwrapItems decodes a clean /v1 list envelope {"items":[...],...} into out.
// The /v1 list endpoints (gateway + nodes) all return this uniform envelope
// instead of the legacy per-type keys (Pieces/Edges/Files/…); SDK list helpers
// keep their existing return structs by unwrapping items into the inner slice,
// so callers (el.Edges, lr.Pieces, …) are unchanged.
func unwrapItems(resByte []byte, out any) error {
	var env struct {
		Items json.RawMessage `json:"items"`
	}
	if err := json.Unmarshal(resByte, &env); err != nil {
		return err
	}
	if len(env.Items) == 0 {
		return nil
	}
	return json.Unmarshal(env.Items, out)
}

// v1URL joins baseUrl + path and appends ?chain= when a chain type is set, so
// the public GET reads select the chain the same way the legacy ?chaintype did.
func v1URL(baseUrl, path string) string {
	u := baseUrl + path
	if chaintype != "" {
		u += "?chain=" + url.QueryEscape(chaintype)
	}
	return u
}

// andSep returns the query separator to append another param to u: "&" when u
// already carries a query string, "?" otherwise.
func andSep(u string) string {
	if strings.Contains(u, "?") {
		return "&"
	}
	return "?"
}

func Disorder(array []types.EdgeReceipt) {
	var temp types.EdgeReceipt
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := len(array) - 1; i >= 0; i-- {
		num := r.Intn(i + 1)
		temp = array[i]
		array[i] = array[num]
		array[num] = temp
	}
}
