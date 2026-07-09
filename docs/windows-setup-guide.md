---
title: "Getting Started with Morpheus on Windows"
description: "Chat with decentralized AI on Windows in minutes — from a one-click app to a free offline model to the developer API."
audience: ["consumer", "developer"]
last_verified: "v7.3.0"
---

# Getting Started with Morpheus on Windows

Morpheus lets you use AI models over a **decentralized network** — no single company owns the
infrastructure. This guide gets you chatting on **Windows** as fast as possible. Pick the path that
fits you:

- **[Path A — Just use it](#path-a--just-use-it-easiest)** (installer app, ~5 min) — the smoothest way in.
- **[Path B — Try it free & offline](#path-b--try-it-free--offline)** (archive build, no wallet, no tokens).
- **[Path C — Build on it](#path-c--build-on-it-developers)** (local API + hosted API, for developers).
- **On a phone / Mac / not Windows?** Open **[app.mor.org](https://app.mor.org)** in any browser and chat there.

---

## ⚠️ First, pick the right download — this trips people up

The [Releases page](https://github.com/MorpheusAIs/Morpheus-Lumerin-Node/releases) offers **two
different Windows builds**. They are **not** the same, and choosing wrong is the #1 early stumble:

| Build | What it is | How you launch it | Best for |
|-------|-----------|-------------------|----------|
| **Installer** — `win-x64-morpheus-app-<ver>.exe` | A normal Windows installer. Installs to `%APPDATA%\morpheus-app\`. | From the **Start-menu / desktop shortcut** it creates. The proxy-router **auto-starts** with the app. | Path A. Most people. |
| **Archive** — the `.zip` build | A folder you extract and run in place. Contains **`mor-launch.exe`**, the CLI, and the router. | Extract, then run **`mor-launch.exe`** from the extracted folder. | Paths B & C. Free local model + command line. |

> 🔑 **The gotcha:** `mor-launch.exe` and the free `mor-launch local` model **only exist in the
> archive build**. If you installed the single-file installer and try to run `mor-launch.exe`, Windows
> will say *"file not found"* — that's expected. Use the Start-menu shortcut instead, or download the
> archive if you specifically want the local-model / command-line flow.

---

## Path A — Just use it (easiest)

1. **Download** `win-x64-morpheus-app-<ver>.exe` from the
   [Releases page](https://github.com/MorpheusAIs/Morpheus-Lumerin-Node/releases).
2. **Run the installer.** Windows SmartScreen may warn about an unknown publisher — click
   **More info → Run anyway** (it's an unsigned binary, not malware).
3. **Open Morpheus** from the Start-menu or desktop shortcut. The proxy-router starts automatically.
4. **Set up your wallet** in the app: set a password, then **create a new wallet** or **recover** an
   existing one with your seed phrase.
   > 🔐 Your seed phrase / private key is a **secret**. Never share it or paste it into a chat.
5. **Add a little funding.** Network models are paid per session, so your wallet needs a small amount of:
   - **ETH on Base** — for transaction gas.
   - **MOR** — staked to open a session (minimum **5 MOR**).
6. **Chat.** Go to **Chat → Change Model**, pick a model, click **Open Session**, stake, and start typing.

✅ You're using decentralized AI.

> 💡 No funds yet? Do **Path B** first — it's free and needs no wallet.

---

## Path B — Try it free & offline

Great for a first taste, a demo, or working without internet. **No wallet, no tokens, no account.**

1. **Download the archive (.zip) build** from the
   [Releases page](https://github.com/MorpheusAIs/Morpheus-Lumerin-Node/releases) — the one that
   contains `mor-launch.exe` (not the single-file installer).
   > _Author note: confirm the exact `.zip` asset name against the live Releases page before publishing._
2. **Extract** it to a folder you can find (e.g. `C:\morpheus`).
3. **Open a terminal** in that folder and run:
   ```powershell
   .\mor-launch.exe local
   ```
   This starts a **bundled local model** (llama.cpp). The first run may download the model once.
4. **Chat** — the UI opens and you can talk to the model running **on your own machine**, offline.

✅ Real AI, running locally, for free.

---

## Path C — Build on it (developers)

Both builds run a local **proxy-router API** on `http://localhost:8082` that is **OpenAI-compatible**.

**Auth:** the router auto-generates an admin password in a **`.cookie`** file
(`admin:<random>`, next to the router; path via `COOKIE_FILE_PATH`).
> 🔐 Treat it as a **secret** — don't commit it or show it on screen. For anything shared, create a
> **scoped user** (`POST /auth/users`) limited to `open_session` + `chat` instead of using admin.

**Open a session, then run inference** (session ID is returned by the first call):

```bash
# 1. Open a session (stakes MOR) -> returns { "sessionID": "..." }
curl -s -u "admin:<ROUTER_API_PASSWORD>" \
  -X POST "http://localhost:8082/blockchain/models/<MODEL_ID>/session" \
  -H "Content-Type: application/json" \
  -d '{"sessionDuration": 600}'

# 2. Run inference with that session ID
curl -s -u "admin:<ROUTER_API_PASSWORD>" \
  -X POST "http://localhost:8082/v1/chat/completions" \
  -H "Content-Type: application/json" \
  -H "session_id: <SESSION_ID>" \
  -H "model_id: <MODEL_ID>" \
  -d '{"model":"<MODEL>","messages":[{"role":"user","content":"Hello!"}],"stream":false}'
```

**Headless wallet:** set `WALLET_PRIVATE_KEY` (falls back to the system keychain when unset).

**Prefer the cloud?** The hosted API at **`https://api.mor.org/api/v1`** is also OpenAI-compatible —
create an account and API key at [app.mor.org](https://app.mor.org) and point any OpenAI client at it.

> 🛠️ Tired of copy-pasting the session ID between steps 1 and 2? `morpheus-doctor --dev` does the whole
> `session → inference` handshake in one command.

---

## Verify it's working

- **Router up?** Visit `http://localhost:8082` (Swagger/API) — it should respond.
- **Inference?** A "Hello!" in the chat (Path A/B) or the curl call (Path C) should return a reply.

---

## Troubleshooting

| Symptom | Cause & fix |
|---------|-------------|
| `mor-launch.exe` not found | You installed the **single-file installer**, which has no `mor-launch.exe`. Launch from the Start-menu shortcut, or download the **archive** build for the launcher / local model. |
| SmartScreen / Defender blocks the app | Unsigned binary. **More info → Run anyway.** |
| Chat says no funds / can't open session | Network models need **Base ETH** (gas) + **≥5 MOR** (stake). Or use **Path B** (free local model). |
| I have two wallets and I'm confused | The app wallet and any CLI/`WALLET_PRIVATE_KEY` wallet are separate — import both into MetaMask to track them in one place. |
| Inference returns 401 / empty | Check the Basic Auth (`.cookie`) credentials and that the `session_id` header matches the one returned when you opened the session. |

---

## Where to go next

- **[app.mor.org](https://app.mor.org)** — use Morpheus from any device/browser.
- **[Official docs](https://nodedocs.mor.org)** — full consumer/provider guides.
- **[GitHub repo](https://github.com/MorpheusAIs/Morpheus-Lumerin-Node)** — run a provider, or contribute.
