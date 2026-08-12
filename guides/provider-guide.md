---
title: "Native Path: Becoming a Morpheus Provider"
description: "A hands-on walkthrough for registering your compute and earning MOR by serving inference."
audience: ["provider", "developer"]
last_verified: "v7.3.0 CLI/API path verified; myprovider.mor.org unverified"
---

# Native Path: Becoming a Morpheus Provider

This guide walks through registering as a compute provider on Morpheus — staking MOR,
listing a model, and serving real inference requests to earn fees — the same hands-on
way as the [Desktop App](./desktop-app-beginner-guide.md) and
[morctl](./morctl-native-guide.md) guides.

**What you'll do:** Fund a wallet → install the router → stand up a model server →
configure securely → register on-chain → open your firewall → start serving → verify.

**Time:** 30–60 minutes, depending on model download size.

---

## Before You Start

**What a provider is, in one sentence:** someone who runs a computer with a GPU,
connects it to the Morpheus network, and gets paid in MOR whenever a consumer's session
routes an inference request to it.

**What you'll need:**

| | Minimum | Recommended |
|---|---|---|
| GPU | Any CUDA GPU, 4GB+ VRAM | RTX 3060+, 12GB+ VRAM |
| RAM | 8GB | 16GB+ |
| Storage | 20GB free | 50GB+ (model weights) |
| Network | Public IP or port-forwarding | Static public IP |
| OS | Linux, Windows, macOS | Ubuntu 22.04+ |

Plus: a wallet with a small amount of **Base ETH** (gas) and **MOR** (1–10 for staking),
and Python 3.10+ if you're running your own model server.

> 🔐 **Never put your private key in a config file.** This guide leaves
> `WALLET_PRIVATE_KEY` blank in `.env` and injects it at runtime from your OS's secure
> credential store instead — see Step 4.

---

## Step 1: Fund a Wallet on Base

1. Install MetaMask (or your preferred EVM wallet), add Base:
   Network `Base` · RPC `https://mainnet.base.org` · Chain ID `8453` · Currency `ETH` ·
   Explorer `https://basescan.org`.
2. Send a small amount of **Base ETH** (~0.01 to start) and **MOR**
   (`0x7431aDa8a591C955a994a21710752EF9b882b8e3`) to this wallet — buy MOR on
   [Uniswap](https://app.uniswap.org/explore/tokens/base/0x7431ada8a591c955a994a21710752ef9b882b8e3)
   or [Aerodrome](https://aerodrome.finance/swap?from=eth&to=0x7431ada8a591c955a994a21710752ef9b882b8e3)
   if needed.

**📸 Capture 1 — MetaMask showing nonzero ETH and MOR, on the Base network specifically.**
Annotate: *"Confirm the network selector says Base — sending to the wrong chain is
unrecoverable."*

✅ Wallet funded.

---

## Step 2: Get the Router

**Option A — prebuilt binary (recommended):**
```bash
# From https://github.com/MorpheusAIs/Morpheus-Lumerin-Node/releases
unzip linux-x86_64-morpheus-router-*.zip -d ~/morpheus/
chmod +x ~/morpheus/proxy-router
```

**Option B — build from source:**
```bash
git clone https://github.com/MorpheusAIs/Morpheus-Lumerin-Node.git
cd Morpheus-Lumerin-Node
# follow the repo's README build instructions
```

**Option C — Desktop App (Windows/macOS)**, if you'd rather use a GUI: download
`win-x64-morpheus-app-<version>.exe` from the same releases page and launch it — bundles
the router with a UI.

> ⚠️ Release filenames change between versions — always check the
> [releases page](https://github.com/MorpheusAIs/Morpheus-Lumerin-Node/releases) for
> current asset names rather than trusting an old link.

**📸 Capture 2 — successful extract + `chmod +x`, no errors.**

✅ You have a router binary.

---

## Step 3: Set Up a Model Server

The router forwards inference requests to a backend you run yourself, which must expose
an **OpenAI-compatible API**.

**Easiest — Ollama:**
```bash
curl -fsSL https://ollama.com/install.sh | sh
ollama pull qwen2.5:7b
ollama serve   # port 11434
```

**For GPU throughput — vLLM:**
```bash
pip install vllm
python -m vllm.entrypoints.openai.api_server \
    --model Qwen/Qwen2.5-7B-Instruct-AWQ --port 8000 --host 0.0.0.0
```

**📸 Capture 3 — a direct test against your model server, before connecting the router:**
```bash
curl http://localhost:8000/v1/models
```
Annotate: *"Get this working alone first — if inference fails later, you'll immediately
know whether the fault is here or on the Morpheus side."*

✅ Your model server answers on its own.

---

## Step 4: Configure the Router Securely

Create `~/morpheus/.env`, **leaving the private key field blank**:
```bash
cat > ~/morpheus/.env << 'EOF'
ETH_NODE_ADDRESS=https://base-mainnet.public.blastapi.io
ETH_NODE_CHAIN_ID=8453
DIAMOND_CONTRACT_ADDRESS=0x6aBE1d282f72B474E54527D93b979A4f64d3030a
MOR_TOKEN_ADDRESS=0x7431aDa8a591C955a994a21710752EF9b882b8e3
WALLET_PRIVATE_KEY=
PROXY_ADDRESS=0.0.0.0:3333
PROXY_STORAGE_PATH=./data/badger/
PROXY_STORE_CHAT_CONTEXT=true
PROXY_FORWARD_CHAT_CONTEXT=true
MODELS_CONFIG_PATH=./models-config.json
WEB_ADDRESS=0.0.0.0:8082
WEB_PUBLIC_URL=http://localhost:8082
AUTH_CONFIG_FILE_PATH=./proxy.conf
COOKIE_FILE_PATH=./.cookie
LOG_LEVEL_APP=info
ENVIRONMENT=production
EOF
```

Then point it at your model server:
```json
{
  "models": [
    { "modelId": "0xYOUR_MODEL_ID", "modelName": "Qwen2.5-7B-Instruct-AWQ",
      "apiType": "openai", "apiUrl": "http://localhost:8000" }
  ]
}
```
Save as `~/morpheus/models-config.json`.

**Store your private key in the OS keychain, once:**

*Linux:*
```bash
secret-tool store --service=morpheus-wallet-key --account=provider <<< "YOUR_PRIVATE_KEY"
```
*macOS:* use Keychain Access, or `security add-generic-password -a "$(whoami)" -s "morpheus-wallet-key" -w "YOUR_PRIVATE_KEY"`

**📸 Capture 4 — the `.env` file's contents, with the blank `WALLET_PRIVATE_KEY=` line
visible.** Annotate: *"This stays blank in the file — the key gets injected at startup,
next step."*

✅ Router configured, no plaintext key anywhere on disk.

---

## Step 5: Register On-Chain

**This is the step people most often miss.** A running, reachable node with no on-chain
registration is invisible to the network — see the real example at the end of this guide.

**Path A — CLI/API (verified, recommended):**
```bash
# 1. Approve MOR for staking
curl -s -u "$(cat ~/morpheus/.cookie)" -X POST \
  "http://localhost:8082/blockchain/approve?spender=0x6aBE1d282f72B474E54527D93b979A4f64d3030a&amount=1000000000000000000000"

# 2. Register as a provider
curl -s -u "$(cat ~/morpheus/.cookie)" -X POST \
  "http://localhost:8082/blockchain/createBlockchainProvider" \
  -H "Content-Type: application/json" \
  -d '{"endpoint": "http://YOUR_PUBLIC_IP:3333", "stakeAmount": "10000000000000000000"}'

# 3. Register a model bid, linking your model to your endpoint
curl -s -u "$(cat ~/morpheus/.cookie)" -X POST \
  "http://localhost:8082/blockchain/createBlockchainProviderBid" \
  -H "Content-Type: application/json" \
  -d '{"modelId": "0xYOUR_MODEL_ID", "endpoint": "http://YOUR_PUBLIC_IP:3333", "pricePerSecond": "1000000000000"}'
```
> Amounts are in wei (18 decimals). Endpoint names may shift between router versions —
> check `http://localhost:8082/swagger/index.html` for your version's exact API surface.

**Path B — web portal** (`myprovider.mor.org`) — **unverified in this project**; the site
is a JS SPA we haven't confirmed end-to-end. If you try it: connect your wallet, enter
your endpoint, register your model, stake, confirm in MetaMask. Treat Path A as the
tested default until this is confirmed.

**Path C — direct contract call (advanced)**, via
[Basescan's Write tab](https://basescan.org/address/0x6aBE1d282f72B474E54527D93b979A4f64d3030a#writeContract):
use `ModelFacet`/`SessionFacet`/`AllowanceFacet` directly. Requires understanding
EIP-2535 Diamond ABI encoding — the same mechanism we used to manually recover a stuck
consumer stake elsewhere in this project.

**📸 Capture 5 — the on-chain confirmation**, whichever path you used: the Basescan
transaction, or the CLI's JSON success response with your new model ID. Annotate:
*"This transaction is what actually makes you visible — not your node being online."*

✅ You have a real, on-chain model ID.

---

## Step 6: Open Your Firewall

```bash
sudo ufw allow 3333/tcp
```
Behind NAT: forward TCP 3333 to your machine in your router's admin panel. On a cloud
VPS: open 3333 in your provider's security group/firewall rules.

**Test from *outside* your network** (e.g. your phone's hotspot, not the same wifi):
```bash
nc -zv YOUR_PUBLIC_IP 3333 -w 5
```

**📸 Capture 6 — the successful `nc` connection, run from a genuinely different network.**

✅ Reachable from the internet, not just locally.

---

## Step 7: Start the Router

```bash
export WALLET_PRIVATE_KEY=$(secret-tool lookup --service=morpheus-wallet-key --account=provider)  # Linux
# or: export WALLET_PRIVATE_KEY=$(security find-generic-password -a "$(whoami)" -s "morpheus-wallet-key" -w)  # macOS
cd ~/morpheus && ./proxy-router
```

**Run it as a persistent service (Linux/systemd)** so it survives reboots and logouts —
wrap the key-injection in a startup script:
```bash
cat > ~/morpheus/start-provider.sh << 'SCRIPT'
#!/bin/bash
export WALLET_PRIVATE_KEY=$(secret-tool lookup --service=morpheus-wallet-key --account=provider)
cd ~/morpheus && exec ./proxy-router
SCRIPT
chmod +x ~/morpheus/start-provider.sh

cat > ~/.config/systemd/user/morpheus-provider.service << 'EOF'
[Unit]
Description=Morpheus Provider
After=network.target
[Service]
Type=simple
ExecStart=%h/morpheus/start-provider.sh
Restart=on-failure
RestartSec=10
[Install]
WantedBy=default.target
EOF

systemctl --user daemon-reload
systemctl --user enable --now morpheus-provider
```

**📸 Capture 7 — the router's startup log**, boxing: wallet loaded, listening on 3333,
no config errors.

✅ Running, and will survive a restart.

---

## Step 8: Verify

```bash
curl -s -u "$(cat ~/morpheus/.cookie)" \
  "http://localhost:8082/blockchain/models/YOUR_MODEL_ID" | python3 -m json.tool
```
A real response (not empty/404) confirms you're discoverable. Cross-check on
[Basescan](https://basescan.org) — your wallet should show an Approve, a registration,
and a staking transaction against the Diamond contract.

✅ Confirmed live and discoverable.

---

## Step 9: Test Real Inference

```bash
curl -s -u "$(cat ~/morpheus/.cookie)" -X POST \
  "http://localhost:8082/blockchain/models/YOUR_MODEL_ID/session" \
  -H "Content-Type: application/json" -d '{"sessionDuration": 600}'

curl -s -u "$(cat ~/morpheus/.cookie)" -X POST \
  "http://localhost:8082/v1/chat/completions" \
  -H "Content-Type: application/json" \
  -H "session_id: SESSION_ID" -H "model_id: YOUR_MODEL_ID" \
  -d '{"model": "Qwen2.5-7B-Instruct-AWQ", "messages": [{"role": "user", "content": "Hello!"}], "stream": false}'
```
A real reply confirms the full loop — you're earning.

✅ Provider setup complete, end to end.

> 🔒 **Note on the test session's stake:** the `session` you opened in this step to test
> your own model behaves like any consumer session — its stake goes through a ~24h
> hold/maturation period before it's claimable, requiring a manual
> `withdrawUserStakes` call to recover (not yet automated by any client). This is
> separate from your **provider** stake in Step 5, which follows different rules. See
> the [headless guide's Step 7](./developer-headless-guide.md#step-7-session-lifecycle--what-happens-next)
> for the full explanation.

---

## Real Example: Why a Setup Wasn't Working

The single most common mistake, from real experience:

| Step | Status | Issue |
|---|---|---|
| Router running | ✅ Online, port 3333 reachable | — |
| Wallet funded | ✅ 0.4 MOR | Below the recommended 1–10 MOR |
| Model server | ✅ Running | — |
| **On-chain registration** | ❌ **Not done** | **This was the actual blocker** |
| MOR staked | ⚠️ 0.4 MOR | Below threshold |
| Firewall | ✅ Open | — |

**Result:** consumers could TCP-connect, but the network had no on-chain record — session
creation returned 404, because the router queries the **blockchain** for model IDs, not
the provider's local config. **Fix:** complete Step 5 and stake sufficient MOR.

---

## Troubleshooting

| Symptom | Cause & fix |
|---|---|
| Model ID returns 404 | Not registered on-chain (Step 5) — a running node isn't enough. |
| Session opens, no response | Test endpoint and model server independently (Steps 3, 6). |
| "No providers found for model" | Under-staked — increase your stake. |
| Connection refused | Router not running, or port 3333 not actually open (`sudo netstat -tlnp \| grep 3333`, `sudo ufw status`). |
| Insufficient ETH/MOR | Fund the wallet further (Step 1). |

---

## Provider Economics

| Price (MOR/sec) | 1-hour session cost | Positioning |
|---|---|---|
| 0.0001 | 0.36 MOR | Budget/competitive |
| 0.001 | 3.6 MOR | Standard |
| 0.01 | 36 MOR | Premium/large models |

**Costs:** one-time MOR stake (unstakeable later), small ETH gas per transaction,
ongoing electricity/bandwidth.

---

## Security Best Practices

1. Never store a private key in a config file — OS keychain or runtime env var only.
2. Use a **dedicated wallet** for providing, not your main one.
3. Keep the router updated.
4. Monitor logs for abuse or errors.
5. Rate-limit your model server if needed.

---

## Glossary

- **Provider** — someone serving inference and earning MOR.
- **Stake** — MOR locked to register a provider/model listing (distinct from a
  consumer's per-session stake).
- **Endpoint** — your node's public IP + port, registered on-chain so sessions route to you.
- **Model server** — the local OpenAI-compatible engine your router forwards requests to.

---

## Useful Links

| Resource | Status |
|---|---|
| [Releases](https://github.com/MorpheusAIs/Morpheus-Lumerin-Node/releases) | Verified |
| [Diamond contract](https://basescan.org/address/0x6abe1d282f72b474e54527d93b979a4f64d3030a) | Verified |
| [MOR token](https://basescan.org/token/0x7431ada8a591c955a994a21710752ef9b882b8e3) | Verified |
| [myprovider.mor.org](https://myprovider.mor.org) | **Unverified** |