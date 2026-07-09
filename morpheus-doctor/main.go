// Command morpheus-doctor diagnoses a local Morpheus proxy-router and, with --dev,
// runs the full session -> inference handshake in one step so you never have to
// copy-paste a session ID between commands.
//
// See SPEC.md for the design. Router API credentials are a secret: they are read
// from MORPHEUS_ROUTER_AUTH ("user:pass") or the router .cookie file, never from a
// command-line flag.
package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	defaultHost     = "localhost:8082"
	defaultModel    = "kimi-k2.5"
	defaultModelID  = "0xbb9e920d94ad3fa2861e1e209d0a969dbe9e1af1cf1ad95c49f76d7b63d32d93"
	defaultPrompt   = "Hello!"
	defaultDuration = 600 // seconds (10 minutes)
	defaultCookie   = "./.cookie"
)

type config struct {
	host     string
	model    string
	modelID  string
	prompt   string
	duration int
	dev      bool
	jsonOut  bool
	cookie   string
	timeout  time.Duration
}

func main() {
	cfg := parseFlags()
	if err := run(cfg); err != nil {
		fmt.Fprintln(os.Stderr, "✖ "+err.Error())
		os.Exit(1)
	}
}

func parseFlags() config {
	var cfg config
	flag.StringVar(&cfg.host, "host", defaultHost, "proxy-router API host:port")
	flag.StringVar(&cfg.model, "model", defaultModel, "model name for the completion request")
	flag.StringVar(&cfg.modelID, "model-id", defaultModelID, "on-chain model id (0x...)")
	flag.StringVar(&cfg.prompt, "prompt", defaultPrompt, "prompt to send in --dev mode")
	flag.IntVar(&cfg.duration, "duration", defaultDuration, "session duration in seconds")
	flag.BoolVar(&cfg.dev, "dev", false, "developer flow: open a session and run inference")
	flag.BoolVar(&cfg.jsonOut, "json", false, "also print raw JSON responses")
	flag.StringVar(&cfg.cookie, "cookie", cookieDefault(), "path to the router .cookie file")
	flag.Parse()
	cfg.timeout = 60 * time.Second
	return cfg
}

func cookieDefault() string {
	if p := os.Getenv("COOKIE_FILE_PATH"); p != "" {
		return p
	}
	return defaultCookie
}

func (c config) baseURL() string {
	h := c.host
	if !strings.HasPrefix(h, "http://") && !strings.HasPrefix(h, "https://") {
		h = "http://" + h
	}
	return strings.TrimRight(h, "/")
}

// loadAuth returns the router API user and password from MORPHEUS_ROUTER_AUTH or
// the .cookie file (format "user:pass"). It never accepts credentials via a flag.
func loadAuth(cfg config) (user, pass string, err error) {
	raw := os.Getenv("MORPHEUS_ROUTER_AUTH")
	src := "MORPHEUS_ROUTER_AUTH"
	if raw == "" {
		b, e := os.ReadFile(cfg.cookie)
		if e != nil {
			return "", "", fmt.Errorf("no router credentials: set MORPHEUS_ROUTER_AUTH=user:pass, or point --cookie at the router .cookie file (%v)", e)
		}
		raw = string(b)
		src = cfg.cookie
	}
	u, p, ok := strings.Cut(strings.TrimSpace(raw), ":")
	if !ok || u == "" {
		return "", "", fmt.Errorf("credentials from %s are not in user:pass form", src)
	}
	return u, p, nil
}

type client struct {
	http *http.Client
	base string
	user string
	pass string
	json bool
}

func newClient(cfg config) (*client, error) {
	u, p, err := loadAuth(cfg)
	if err != nil {
		return nil, err
	}
	return &client{
		http: &http.Client{Timeout: cfg.timeout},
		base: cfg.baseURL(),
		user: u,
		pass: p,
		json: cfg.jsonOut,
	}, nil
}

// do performs an authenticated JSON request. out (if non-nil) is decoded from the
// response body. Any transport error propagates; any non-2xx status is an error.
func (c *client) do(method, path string, headers map[string]string, body, out any) error {
	var rdr io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return err
		}
		rdr = bytes.NewReader(b)
	}
	req, err := http.NewRequest(method, c.base+path, rdr)
	if err != nil {
		return err
	}
	req.SetBasicAuth(c.user, c.pass)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if c.json {
		fmt.Printf("  %s %s -> %s\n  %s\n", method, path, resp.Status, snippet(data))
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("%s %s -> %s: %s", method, path, resp.Status, snippet(data))
	}
	if out != nil {
		if err := json.Unmarshal(data, out); err != nil {
			return fmt.Errorf("decode %s response: %w", path, err)
		}
	}
	return nil
}

// reachable reports whether the router answers HTTP at all. Any HTTP response
// (even 401/404) means the router is up; only a transport error means it is down.
func (c *client) reachable() error {
	req, err := http.NewRequest(http.MethodGet, c.base+"/healthcheck", nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(c.user, c.pass)
	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}

func run(cfg config) error {
	c, err := newClient(cfg)
	if err != nil {
		return err
	}
	fmt.Printf("morpheus-doctor  ->  %s\n", c.base)

	if err := c.reachable(); err != nil {
		fmt.Println("✖ router not reachable")
		fmt.Println("  Start Morpheus first, then re-run:")
		fmt.Println("   - Installer build: open the app from the Start-menu/desktop shortcut (router auto-starts)")
		fmt.Println("   - Archive build:   run  mor-launch.exe  in the extracted folder")
		return fmt.Errorf("router unreachable at %s (%v)", c.base, err)
	}
	fmt.Println("✔ router reachable")

	if !cfg.dev {
		fmt.Println("\nPreflight OK. Re-run with --dev to open a session and run inference.")
		return nil
	}
	return devFlow(c, cfg)
}

type sessionResp struct {
	SessionID string `json:"sessionID"`
}

type chatResp struct {
	Choices []struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

// devFlow opens a session, captures the session ID in memory, and runs one
// inference request with it -- replacing the manual open-session / copy-id /
// run-inference chain (AUDIT F3/F4).
func devFlow(c *client, cfg config) error {
	fmt.Printf("\n> opening session (model %s, %ds)...\n", cfg.model, cfg.duration)
	var sess sessionResp
	err := c.do(http.MethodPost, "/blockchain/models/"+cfg.modelID+"/session", nil,
		map[string]any{"sessionDuration": cfg.duration}, &sess)
	if err != nil {
		return fmt.Errorf("open session: %w", err)
	}
	if sess.SessionID == "" {
		return errors.New("open session: response contained no sessionID")
	}
	fmt.Printf("✔ session opened: %s\n", sess.SessionID)

	fmt.Println("> running inference...")
	var chat chatResp
	err = c.do(http.MethodPost, "/v1/chat/completions",
		map[string]string{"session_id": sess.SessionID, "model_id": cfg.modelID},
		map[string]any{
			"model":    cfg.model,
			"messages": []map[string]string{{"role": "user", "content": cfg.prompt}},
			"stream":   false,
		}, &chat)
	if err != nil {
		return fmt.Errorf("inference: %w", err)
	}
	fmt.Println("✔ inference ok")

	if len(chat.Choices) > 0 {
		fmt.Printf("\n  %s\n", strings.TrimSpace(chat.Choices[0].Message.Content))
	}
	fmt.Printf("\n✅ session %s live (~%dm)\n", short(sess.SessionID), cfg.duration/60)
	return nil
}

func snippet(b []byte) string {
	s := strings.TrimSpace(string(b))
	if len(s) > 300 {
		return s[:300] + "..."
	}
	return s
}

func short(s string) string {
	if len(s) > 10 {
		return s[:10] + "..."
	}
	return s
}
