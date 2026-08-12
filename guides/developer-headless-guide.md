---
title: "Native Path: Headless / Developer Setup (Raw API)"
description: "Run a Morpheus proxy-router headless and talk to it with raw curl — for integrating Morpheus into your own tooling."
audience: ["developer"]
last_verified: "v7.3.0, tested on Windows PowerShell and WSL2/Linux"
---

# Native Path: Headless / Developer Setup (Raw API)

This guide walks through running a Morpheus proxy-router with no GUI at all, and talking
to it with plain `curl` — the exact HTTP calls that [morctl](./morctl-native-guide.md)
automates for you. Use this guide if you're integrating Morpheus into your **own**
tooling and want to see precisely what's sent over the wire, rather than using a
pre-built client.

**If you just want to use Morpheus, not build on it**, use the
[Desktop App guide](./desktop-app-beginner-guide.md) or [morctl](./morctl-native-guide.md)
instead — both are faster paths to the same result.

**What you'll do:** Install the router → configure it → inject a wallet → open a
session → run inference → understand session lifecycle → (optional) multi-turn context.

Covers **Windows (PowerShell)** and **WSL2/Linux** — pick your tab at each step; the
concepts are identical, only the shell syntax differs.

---

## Before You Start

**What you'll need:**
- **Windows:** PowerShell (built in).
- **WSL2/Linux:** a working WSL2 install or native Linux, `bash`.
- A funded wallet — at least 5 MOR and a small amount of Base ETH on Base mainnet.
- `curl`, available by default on both platforms.

> 🔐 **You'll be handling a real private key in this guide.** Never write it to a
> committed file, never paste it into chat tools, and clear it from your shell history
> when you're done (commands at the end of Step 3).

---

## Step 1: Get the Router Binary

**WSL2/Linux:**
```bash
mkdir -p ~/morpheus && cd ~/morpheus
wget "https://github.com/MorpheusAIs/Morpheus-Lumerin-Node/releases/download/v7.3.0/linux-x86_64-morpheus-router-7.3.0" -O proxy-router
chmod +x proxy-router
```

**Windows (PowerShell):**
```powershell
mkdir C:\morpheus; cd C:\morpheus
Invoke-WebRequest "https://github.com/MorpheusAIs/Morpheus-Lumerin-Node/releases/download/v7.3.0/win-x64-morpheus-router-7.3.0.exe" -OutFile proxy-router.exe
```

> ⚠️ Check the [releases page](https://github.com/MorpheusAIs/Morpheus-Lumerin-Node/releases)
> for the current version — filenames change between releases. Avoid any `-test`-suffixed
> release unless you specifically want testnet.

**📸 Capture 1 — the download completing and the binary present in the folder** (`ls` /
`dir`). Annotate: *"Confirm no `-test` in the filename — that's a testnet build."*

✅ You have a router binary.

---

## Step 2: Configure `.env`

Both platforms use the same variables — only the shell syntax for creating the file differs.

**WSL2/Linux:**
```bash
cat > .env << 'EOF'
ETH_NODE_ADDRESS=https://mainnet.base.org
ETH_NODE_CHAIN_ID=8453
DIAMOND_CONTRACT_ADDRESS=0x6aBE1d282f72B474E54527D93b979A4f64d3030a
MOR_TOKEN_ADDRESS=0x7431aDa8a591C955a994a21710752EF9b882b8e3
WALLET_PRIVATE_KEY=
MODELS_CONFIG_PATH=./models-config.json
COOKIE_FILE_PATH=./.cookie
WEB_ADDRESS=0.0.0.0:8082
PROXY_ADDRESS=0.0.0.0:3333
PROXY_STORAGE_PATH=./data/
EOF
```

**Windows (PowerShell):**
```powershell
@"
ETH_NODE_ADDRESS=https://mainnet.base.org
ETH_NODE_CHAIN_ID=8453
DIAMOND_CONTRACT_ADDRESS=0x6aBE1d282f72B474E54527D93b979A4f64d3030a
MOR_TOKEN_ADDRESS=0x7431aDa8a591C955a994a21710752EF9b882b8e3
WALLET_PRIVATE_KEY=
MODELS_CONFIG_PATH=./models-config.json
COOKIE_FILE_PATH=./.cookie
WEB_ADDRESS=0.0.0.0:8082
PROXY_ADDRESS=0.0.0.0:3333
PROXY_STORAGE_PATH=./data/
"@ | Out-File -Encoding ascii .env
```

Note `WALLET_PRIVATE_KEY` is left **blank** — you'll inject it live in Step 3, not store
it in this file.

**📸 Capture 2 — the `.env` file contents printed** (`cat .env` / `type .env`), with the
blank key line visible.

✅ Router configured.

---

## Step 3: Start the Router and Provide a Wallet

**Terminal 1 — start the router, leave it running:**
```bash
./proxy-router          # Linux/WSL2
.\proxy-router.exe       # Windows
```
Wait for `proxy state: waiting for private key` in the log.

**📸 Capture 3 — the log showing "waiting for private key."**

**Terminal 2 — inject your wallet:**

*WSL2/Linux:*
```bash
cd ~/morpheus
read -rs PK
printf '{"privateKey":"%s"}' "$PK" \
  | curl -X POST 'http://localhost:8082/wallet/privateKey' \
    -H 'Content-Type: application/json' -u "$(cat .cookie)" --data-binary @-
unset PK
```

*Windows (PowerShell):*
```powershell
cd C:\morpheus
$PK = Read-Host -AsSecureString "Private key" | ConvertFrom-SecureString -AsPlainText
$cookie = Get-Content .cookie
$user, $pass = $cookie -split ":", 2
$b64 = [Convert]::ToBase64String([Text.Encoding]::UTF8.GetBytes("$($user):$($pass)"))
$body = @{ privateKey = $PK } | ConvertTo-Json
Invoke-RestMethod -Uri "http://localhost:8082/wallet/privateKey" -Method POST `
  -Headers @{ Authorization = "Basic $b64" } -ContentType "application/json" -Body $body
Remove-Variable PK
```

**Back in Terminal 1**, the log should flip from "waiting" to reporting your wallet
address and MOR/ETH balance.

**📸 Capture 4 — Terminal 1's log showing the wallet address and balance after
injection.** This is your confirmation the key worked — annotate: *"If this doesn't
appear, the key wasn't accepted — check Terminal 2's response for an error."*

✅ Router is running with a real, funded wallet, key never written to disk.

---

## Step 4: Open a Session

```bash
curl -s -X POST 'http://localhost:8082/blockchain/models/<MODEL_ID>/session' \
  -H 'Content-Type: application/json' \
  -u "$(cat .cookie)" \
  -d '{"sessionDuration": 600, "failover": true, "directPayment": false}'
```
(Same command on Windows PowerShell, or use `Invoke-RestMethod` with the auth header
pattern from Step 3.)

Get a real `<MODEL_ID>` from `GET /blockchain/models`, or from the Desktop App's Models
tab.

**📸 Capture 5 — the response**, boxing the `sessionID` field:
```json
{"sessionID":"0xa3e0ac1a374208b4420d84af96a2f7b27b34d5bbe742be9f8c656"}
```
Annotate: *"Copy this — you need it for every inference call in this session."*

✅ Session open, MOR staked.

---

## Step 5: Run Inference

```bash
curl -s -X POST 'http://localhost:8082/v1/chat/completions' \
  -H 'Content-Type: application/json' \
  -H "session_id: <SESSION_ID>" \
  -u "$(cat .cookie)" \
  -d '{"messages": [{"role": "user", "content": "Hello. Tell me about AI."}]}'
```

**📸 Capture 6 — the model's reply in the JSON response.** Tip: append
`| python3 -m json.tool` for readable output.

✅ You've run raw inference with no client library at all.

---

## Step 6: Multi-Turn Conversation (Optional)

⚠️ **The router does not remember previous messages on its own** — this is intended
behavior, not a bug (confirmed by maintainers). Two ways to get real conversation memory:

**Option A — resend full history yourself** (what `morctl --interactive` does
automatically): include every prior message in the `messages` array on each call.

**Option B — use a `chat_id`** to have the router persist and prepend history for you:
```bash
CHAT_ID=0x$(openssl rand -hex 32)

curl -s -X POST 'http://localhost:8082/v1/chat/completions' \
  -H 'Content-Type: application/json' \
  -H "session_id: <SESSION_ID>" -H "chat_id: $CHAT_ID" \
  -u "$(cat .cookie)" \
  -d '{"messages":[{"role":"user","content":"My name is Alex."}]}'

curl -s -X POST 'http://localhost:8082/v1/chat/completions' \
  -H 'Content-Type: application/json' \
  -H "session_id: <SESSION_ID>" -H "chat_id: $CHAT_ID" \
  -u "$(cat .cookie)" \
  -d '{"messages":[{"role":"user","content":"What is my name?"}]}'
```
Requires `PROXY_STORE_CHAT_CONTEXT=true` and `PROXY_FORWARD_CHAT_CONTEXT=true` in your
`.env`. History for a given `chat_id` is viewable locally at
`./data/chats/<CHAT_ID>.json`.

**📸 Capture 7 — the second response correctly recalling "Alex" from the first message.**

✅ Genuine multi-turn conversation via the raw API.

---

## Step 7: Session Lifecycle — What Happens Next

- **Let it expire naturally.** Your node auto-submits the close transaction ~1 minute
  after `endsAt`, and your stake is eligible for return.
- **Stake return is not automatic or instant.** As of the current router version, the
  full stake goes through a hold/maturation period (roughly 24 hours) before it's
  claimable, and claiming currently requires a **separate manual call** to
  `withdrawUserStakes` on the Diamond contract (via Basescan's Write tab, or a direct
  contract call) — this is not yet exposed through the router's HTTP API. Track
  [issue #827](https://github.com/MorpheusAIs/Morpheus-Lumerin-Node/issues/827) for
  when an API for this ships.
- **Don't close early** unless you have a reason to — early closes route through the
  same hold mechanism with no offsetting benefit.

✅ You understand the full lifecycle, including the part that trips people up.

---

## Troubleshooting

| Symptom | Cause & fix |
|---|---|
| `waiting for private key` never resolves | Terminal 2's POST failed — check its response for an error before assuming Terminal 1 is broken. |
| `401 Unauthorized` | Wrong `.cookie` value — re-read the file, it may have regenerated since you last checked it. |
| `insufficient MOR balance` on session open | Fund the wallet further, or request a shorter `sessionDuration`. |
| Second chat message has no memory of the first | You didn't resend history or use a `chat_id` — see Step 6. |
| Stake didn't come back after closing | Expected right now — see Step 7. Not a bug in your setup. |

---

## Cross-guide reference

The stake-lock behavior in Step 7 above applies identically regardless of which client
opened the session — Desktop App, `morctl`, or raw curl as in this guide. See
[morctl guide](./morctl-native-guide.md) and
[Desktop App guide](./desktop-app-beginner-guide.md) for the same warning in those
contexts.

## Glossary

- **`.cookie`** — the router's auto-generated Basic Auth credential (`user:pass`).
- **Session** — a time-boxed window during which you can run inference against a model.
- **`chat_id`** — an optional header that has the router persist and prepend
  conversation history for you, as an alternative to resending it yourself.
- **Stake hold** — the maturation period between a session closing and its stake
  becoming claimable; currently requires a manual on-chain call to release.