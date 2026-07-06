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
var ChainType = build.BNBTestnetDAO
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

const InHashID = hash.MIMC_BW6_761

func CheckENV() {
	if env.Str(env.ChainType, "") == "" {
		os.Setenv(env.ChainType, ChainType)
	}
	chaintype = build.CheckChain()
	logger.Warn("connect to chain: ", chaintype)
}

func CheckFileFull(ff types.FileFull, stream common.Address, fp string) ([]types.PieceCore, error) {
	logger.Debug("check stream handle of file: ", fp)
	res := make([]types.PieceCore, len(ff.Pieces))
	p, err := homedir.Expand(fp)
	if err != nil {
		return res, err
	}

	fi, err := os.Open(p)
	if err != nil {
		return res, err
	}
	defer fi.Close()

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
			n, err := fi.Read(buf)
			if err != nil {
				return nil, err
			}
			if n != int(size) {
				return nil, fmt.Errorf("short read length")
			}
			rest -= size

			cval := bls.Eval(bls.Split(31, buf), rnd)
			if cval.Cmp(&ew.ClaimedValues[j]) != 0 {
				return nil, fmt.Errorf("unequal val at %d %d", i, j)
			}
		}

		res[i] = types.PieceCore{
			Policy:   ff.Policy,
			Name:     ff.Pieces[i],
			Size:     ff.PieceSizes[i],
			Streamer: stream,
		}
	}
	return res, nil
}

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
)

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
	return issuedAt, nil
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

	pr := progressbar.NewReader(resp.Body, bar)
	res, err := io.ReadAll(&pr)
	if err != nil {
		return nil, err
	}
	bar.Finish()

	if resp.StatusCode != 200 {
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

	if resp.StatusCode != 200 {
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
