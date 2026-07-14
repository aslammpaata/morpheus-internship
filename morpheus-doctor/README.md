# morpheus-doctor

Preflight + one-command `session → inference` for the Morpheus proxy-router.
Turns the manual "start router → open session → copy the session ID → run inference"
chain into a single command. See [`SPEC.md`](./SPEC.md).

## Build

Requires **Go 1.21+**. From this folder:

```bash
go build -o morpheus-doctor .
```

Cross-compile a Windows `.exe` (e.g. from WSL/Linux — ideal for the booth laptops):

```bash
GOOS=windows GOARCH=amd64 go build -o morpheus-doctor.exe .
```

> The module path (`github.com/aslamwaiswa/morpheus-doctor`) is cosmetic for a single-binary tool.
> Rename it any time with `go mod edit -module <your/path>`.

## Credentials (never pass the password on the command line)

The router API password is a secret (the auto-generated `.cookie`). Provide it via **either**:

- an environment variable: `MORPHEUS_ROUTER_AUTH=admin:<password>`, or
- the router `.cookie` file: `--cookie ./.cookie` (or set `COOKIE_FILE_PATH`).

For anything shared (a booth), prefer a **scoped user** (`open_session` + `chat` only) over admin.

## Use

Preflight — just check the router is up:

```bash
./morpheus-doctor
```

Developer flow — open a session and run inference in one step:

```bash
./morpheus-doctor --dev --prompt "Hello!"
```

Multi-turn conversation (same session, full history resent each turn):

```bash
./morpheus-doctor --dev --interactive
```

Flags: `--host` (default `localhost:8082`), `--model`, `--model-id`, `--duration`,
`--interactive` (keep the session open for multiple prompts — type `exit`/`quit`
to end; conversation history is resent each turn), `--json` (print raw
responses), `--cookie`.


### Status

**Tested against a live router (Base mainnet).** Verified: preflight (router
up/down), `--dev` happy path against two different models (proving it's not
hardcoded to one), `--interactive` multi-turn chat with correct context
retention, and graceful handling of router-down, malformed/unregistered model
IDs, and a real mid-session provider disconnect (no crash, failed turn is
retried by the user, not left in history).

**Known limitation:** exiting `--interactive` (via `exit`/`quit` or Ctrl+C)
stops the local loop only — it does **not** close the session early. Your
stake refunds automatically at natural session expiry, same as any Morpheus
session; there is no penalty for exiting early this way, you just wait out the
remaining duration for the refund, same as if you'd simply stopped chatting in
the desktop app without closing.

Next: scoped-user creation helper (AUDIT F9), `--local` beginner preflight
track (SPEC.md).