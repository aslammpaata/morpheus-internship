# `morpheus-doctor` — Specification (draft)

A single static **Go** binary that takes a user from a cold machine to a verified first inference,
and eliminates the manual developer chain documented in [`../AUDIT.md`](../AUDIT.md).

Ships as `morpheus-doctor.exe` (Windows-first) with no runtime dependencies — download and run.

## Why Go

- Compiles to one static `.exe` — no Python/Node runtime for a booth beginner to install.
- Same language as the Lumerin node → path to upstreaming as a `contrib/` tool.

## Two tracks

### `morpheus-doctor` (default — beginner preflight)
Diagnoses the *desktop-app* path without touching it. Read-only, friendly ✅/❌ report:
- Is the desktop app / router process reachable? (health-check `localhost:8082`)
- Is a wallet configured and funded? (Base ETH for gas, MOR for sessions)
- Network sanity: correct chain (Base mainnet 8453), RPC reachable.
- Can it reach a provider/model? (e.g. Kimi K2.5)
- Prints exactly what's missing and how to fix it — no jargon.

### `morpheus-doctor --dev` (the flagship — automate the manual chain)
Replaces steps 1/3/4 of the manual `curl` dance (AUDIT F2/F3/F4) with **one command**:

1. **Start / verify** the router is up (health-check; start it if needed).
2. **Open a session** — POST to `/blockchain/models/<model_id>/session`, parse the `sessionID`
   from the response **in memory** (this is the fix for F4 — no hand-copying).
3. **Run inference** — POST to `/v1/chat/completions` with the captured `session_id` + `model_id`.
4. **Report** — show the model's reply and a "✅ session <id> live for N minutes" summary.

```
$ morpheus-doctor --dev --model kimi-k2.5 --prompt "Hello!"
✔ router healthy (localhost:8082)
✔ session opened  (id: 3f9a… , 10 min)
✔ inference ok
  > "Hello! How can I help you today?"
```

## Config & secrets

- Router API password (Basic Auth) is a **secret** (AUDIT F9). Read from env
  (`MORPHEUS_ROUTER_AUTH`) or a git-ignored config file — **never** a flag or committed default.
- Flags: `--model`, `--prompt`, `--duration`, `--host` (default `localhost:8082`), `--json`.

## Open design questions (blocked on AUDIT answers)

- **WSL vs native (F7):** if the router only runs under WSL/systemd, does `--dev` start it via WSL,
  or do we target the native `win-x64-morpheus-router-7.3.0.exe`? Prefer native for a true one-`.exe` story.
- **Wallet setup (F8):** should `--dev` also wrap wallet creation, or assume a funded wallet exists?
  Depends on whether there's a canonical (non-`everclaw`) wallet command.

## Non-goals

- Not an installer for the desktop app (that path is already good — AUDIT).
- Not a replacement for the node or its API — a thin, honest wrapper only.
