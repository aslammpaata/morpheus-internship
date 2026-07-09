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

## Config & secrets (canonical names, from /reference/env-proxy-router)

- Router API auth comes from the **`.cookie`** file (`admin:<random>`, default `./.cookie`; path via
  `COOKIE_FILE_PATH`). Treat as a **secret** (AUDIT F9). `morpheus-doctor` reads it from the cookie
  file or `MORPHEUS_ROUTER_AUTH` env — **never** a flag or committed default. Prefer creating a
  **scoped `chat`/`open_session` user** (`POST /auth/users`) over using admin.
- Router API address: `WEB_ADDRESS` (default `0.0.0.0:8082`).
- Headless wallet: **`WALLET_PRIVATE_KEY`** (falls back to system keychain when unset).
- Network: `ETH_NODE_ADDRESS` (RPC), `ETH_NODE_CHAIN_ID` (`8453` main / `84532` test),
  `DIAMOND_CONTRACT_ADDRESS`, `MOR_TOKEN_ADDRESS`.
- Flags: `--model`, `--prompt`, `--duration`, `--host` (default `localhost:8082`), `--json`.

## Design decisions (resolved from AUDIT)

- **Target:** build native `windows/amd64` `.exe` (cross-compiled from the author's WSL dev env).
- **Router lifecycle (F7):** canonical Windows run is `mor-launch.exe` (no systemd). `--dev`
  **assumes the router is up and health-checks it**; it does not manage systemd. Keeps the tool thin.
- **Wallet (F8):** assumes a funded wallet exists (desktop app or `WALLET_PRIVATE_KEY`); the tool
  does **not** wrap wallet creation.

## Non-goals

- Not an installer for the desktop app (that path is already good — AUDIT).
- Not a replacement for the node or its API — a thin, honest wrapper only.
