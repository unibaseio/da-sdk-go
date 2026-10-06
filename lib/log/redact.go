package log

import (
	"io"
	"net/url"
	"sort"
	"strings"
	"sync"
	"sync/atomic"

	"go.uber.org/zap/zapcore"
)

// Hosted RPC providers (Alchemy, Infura, QuickNode, ...) put the API key in the
// endpoint URL, and every transport error from net/http quotes the full URL
// (`Post "https://host/v2/<key>": ...`). Such errors travel up through many call
// sites before something logs them, so redaction is done where logs are
// written: secrets registered with RegisterURL are replaced in every zap entry
// and every stdlib log line.

const redacted = "<redacted>"

// a secret shorter than this is not a credential (e.g. "/rpc") and replacing it
// would mangle unrelated log text
const minSecretLen = 8

var (
	secretMu sync.Mutex
	secrets  = map[string]struct{}{}
	replacer atomic.Pointer[strings.Replacer] // nil: nothing registered
)

// RedactURL keeps only the scheme and host of an endpoint, for logging which
// endpoint is in use without its credentials.
func RedactURL(raw string) string {
	u, err := url.Parse(raw)
	if err != nil || u.Host == "" {
		return redacted
	}
	s := u.Scheme + "://" + u.Host
	if u.User != nil || strings.Trim(u.Path, "/") != "" || u.RawQuery != "" || u.Fragment != "" {
		s += "/" + redacted
	}
	return s
}

// RegisterURL marks the credential-bearing parts of an endpoint URL as secrets:
// from now on they are replaced in all log output. Registered are
//   - the userinfo: username and password (providers put the key in either;
//     net/http masks a password in its errors but not a username),
//   - token-like host labels (`https://<key>.rpc.example.com`),
//   - token-like path segments (`/v2/<key>`) — not the whole path, which
//     would also scrub ordinary routes such as /v1/ethereum/stats,
//   - the query string.
//
// "Token-like" is tokenLike: long, URL-safe, mixing letters and digits. A
// shorter key, or one of digits only, in a path or host is not caught.
func RegisterURL(raw string) {
	u, err := url.Parse(raw)
	if err != nil {
		// unparsable: the whole string is the only safe unit
		register(raw)
		return
	}
	if u.User != nil {
		name := u.User.Username()
		register(name)
		register(url.QueryEscape(name))
		register(url.PathEscape(name))
		if p, ok := u.User.Password(); ok {
			register(p)
			register(url.QueryEscape(p))
			register(url.PathEscape(p))
		}
	}
	for _, label := range strings.Split(u.Hostname(), ".") {
		if tokenLike(label, minHostTokenLen) {
			register(label)
		}
	}
	for _, path := range []string{u.Path, u.EscapedPath()} {
		for _, seg := range strings.Split(path, "/") {
			if tokenLike(seg, minPathTokenLen) {
				register(seg)
			}
		}
	}
	register(u.RawQuery)
}

// The shortest host label / path segment taken for a key. Provider keys are
// 16+ characters (Alchemy/Infura 32, QuickNode 40, Ankr 64, UUIDs 36). Path
// segments of an RPC URL are otherwise short version tags, so 16 is safe
// there; host labels are often long descriptive names, so they need 20.
const (
	minPathTokenLen = 16
	minHostTokenLen = 20
)

// tokenLike reports whether a host label or path segment looks like an API
// key: at least min characters of [A-Za-z0-9_-], with both a letter and a
// digit (so plain words such as "ethereum" or "base-sepolia-rpc" are not).
func tokenLike(s string, min int) bool {
	if len(s) < min {
		return false
	}
	letter, digit := false, false
	for _, r := range s {
		switch {
		case r >= 'a' && r <= 'z', r >= 'A' && r <= 'Z':
			letter = true
		case r >= '0' && r <= '9':
			digit = true
		case r == '-' || r == '_':
		default:
			return false
		}
	}
	return letter && digit
}

func register(s string) {
	if len(s) < minSecretLen {
		return
	}
	secretMu.Lock()
	defer secretMu.Unlock()
	if _, ok := secrets[s]; ok {
		return
	}
	secrets[s] = struct{}{}
	// longest first, so a secret containing another is replaced whole
	keys := make([]string, 0, len(secrets))
	for k := range secrets {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return len(keys[i]) > len(keys[j]) })
	pairs := make([]string, 0, 2*len(keys))
	for _, k := range keys {
		pairs = append(pairs, k, redacted)
	}
	replacer.Store(strings.NewReplacer(pairs...))
}

// Scrub replaces registered secrets in s.
func Scrub(s string) string {
	r := replacer.Load()
	if r == nil {
		return s
	}
	return r.Replace(s)
}

type scrubWriter struct{ w io.Writer }

// Write scrubs one log entry. zap and the stdlib logger each emit an entry in a
// single Write, so a secret is never split across calls. It reports len(p) on
// success: the caller wrote p, whatever length it became.
func (s scrubWriter) Write(p []byte) (int, error) {
	r := replacer.Load()
	if r == nil {
		return s.w.Write(p)
	}
	if _, err := io.WriteString(s.w, r.Replace(string(p))); err != nil {
		return 0, err
	}
	return len(p), nil
}

// ScrubWriter wraps w so registered secrets never reach it.
func ScrubWriter(w io.Writer) io.Writer { return scrubWriter{w} }

// scrubSyncer is scrubWriter for zap's WriteSyncer.
type scrubSyncer struct {
	scrubWriter
	ws zapcore.WriteSyncer
}

func (s scrubSyncer) Sync() error { return s.ws.Sync() }

func scrubWriteSyncer(ws zapcore.WriteSyncer) zapcore.WriteSyncer {
	return scrubSyncer{scrubWriter{ws}, ws}
}
