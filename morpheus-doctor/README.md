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

Flags: `--host` (default `localhost:8082`), `--model`, `--model-id`, `--duration`,
`--json` (print raw responses), `--cookie`.

## Status

**Scaffold.** The happy path — reachable check → open session → capture session ID → inference —
is implemented against the documented API, with endpoint/body shapes matching the verified curl
flow in [`../AUDIT.md`](../AUDIT.md) (mainnet Base, Kimi K2.5). Next steps: build & run against a
live router, add the scoped-user creation helper (AUDIT F9), and a `--local` beginner preflight.
