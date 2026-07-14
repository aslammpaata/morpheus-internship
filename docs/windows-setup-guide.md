---
title: "Getting Started with Morpheus on Windows"
description: "Chat with decentralized AI on Windows in minutes — from a one-click app to a free local model to the developer API."
audience: ["consumer", "developer"]
last_verified: "v7.3.0"
---

# Getting Started with Morpheus on Windows

Morpheus lets you use AI models over a **decentralized network** — no single company owns the
infrastructure. This guide gets you chatting on **Windows** as fast as possible. Pick the path that
fits you:

- **[Path A — Just use it](#path-a--just-use-it-easiest)** (installer app, ~5 min) — the smoothest way in.
- **[Path B — Try it free & offline](#path-b--try-it-free--offline)** (same install, no wallet, no tokens).
- **[Path C — Build on it](#path-c--build-on-it-developers)** (local API + hosted API, for developers).
- **On a phone / Mac / not Windows?** Open **[app.mor.org](https://app.mor.org)** in any browser and chat there.

---

## ⚠️ One download does everything — ignore any mention of a `.zip` archive

The [Releases page](https://github.com/MorpheusAIs/Morpheus-Lumerin-Node/releases) lists a few
Windows assets. For almost everyone, there is only **one you want**:

| Build | What it is | How you launch it | Best for |
|-------|-----------|-------------------|----------|
| **Installer** — `win-x64-morpheus-app-<ver>.exe` | A normal Windows installer. Installs to `%APPDATA%\morpheus-app\`. Bundles the desktop UI, the proxy-router, **and a free local model**. | From the **Start-menu / desktop shortcut** it creates. The proxy-router **auto-starts** with the app. | Paths A **and** B. Nearly everyone. |
| `win-x64-morpheus-router-<ver>.exe` | Headless proxy-router only, no GUI. For running a provider on a server. | Manual `.env` setup, command line. | Providers only — not covered here. |
| `win-x64-morpheus-cli-<ver>.exe` | Standalone CLI client. | — | **Currently broken on Windows** — see the note in Path C. Skip it for now. |

> 🔑 **If you've seen instructions elsewhere mentioning a `.zip` archive and `mor-launch.exe`:**
> that build does not currently exist on the Releases page (checked across every release from
> v7.1.12 through v7.3.3 — [tracked here](https://github.com/MorpheusAIs/Morpheus-Lumerin-Node/issues/796)).
> Ignore those instructions — the **installer** above covers everything they promised, including the
> free offline model.

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

> 💡 No funds yet? Do **Path B** first — it's free and needs no wallet, and it's the exact same install
> you just did.

---

## Path B — Try it free & offline

Great for a first taste, a demo, or working without internet. **No wallet, no tokens, no account —
and no separate download.** This uses the same installer from Path A.

1. If you haven't already, complete **step 1–3 of Path A** (download, install, open the app).
2. In the **Chat** tab, open the model picker and select the **bundled local model**
   (llama.cpp-based, runs entirely on your machine).
3. **Chat** — no wallet, no session, no MOR required. Responses come from your own hardware.

✅ Real AI, running locally, for free, from the same app you already installed.

---

## Path C — Build on it (developers)

The Desktop App runs a local **proxy-router API** on `http://localhost:8082` that is
**OpenAI-compatible**.

**Auth:** the router auto-generates an admin password in a **`.cookie`** file
(`admin:<random>`, default location `%APPDATA%\morpheus-app\services\.cookie` for the installer
build; path via `COOKIE_FILE_PATH`).
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

> ⚠️ **Multi-turn chat:** the router does **not** remember earlier messages in a session on its own.
> Resend the full conversation history in `messages` on every request (standard OpenAI-style client
> behavior) — see [issue #794](https://github.com/MorpheusAIs/Morpheus-Lumerin-Node/issues/794).

**Headless wallet:** set `WALLET_PRIVATE_KEY` (falls back to the system keychain when unset).

**Prefer the cloud?** The hosted API at **`https://api.mor.org/api/v1`** is also OpenAI-compatible —
create an account and API key at [app.mor.org](https://app.mor.org) and point any OpenAI client at it.

> 🛠️ **Tired of copy-pasting the session ID between steps 1 and 2, or hand-managing multi-turn
> history?** [`morpheus-doctor`](../morpheus-doctor/) does the whole `session → inference` handshake
> in one command (`--dev`), and handles multi-turn conversations correctly out of the box
> (`--dev --interactive`). Tested end-to-end against multiple models.
>
> **Skip `mor-cli` for now** — the official CLI binary currently fails to launch on Windows
> ([issue #792](https://github.com/MorpheusAIs/Morpheus-Lumerin-Node/issues/792)); `morpheus-doctor`
> or raw `curl` are the working alternatives until it's fixed.

---

## Verify it's working

- **Router up?** Visit `http://localhost:8082` (Swagger/API) — it should respond.
- **Inference?** A "Hello!" in the chat (Path A/B) or the curl call / `morpheus-doctor` (Path C) should
  return a reply.

---

## Troubleshooting

| Symptom | Cause & fix |
|---------|-------------|
| Instructions online mention `mor-launch.exe` or a `.zip` archive and it doesn't exist on the Releases page | That build isn't currently shipped ([issue #796](https://github.com/MorpheusAIs/Morpheus-Lumerin-Node/issues/796)). Use the installer — it includes everything, including the free local model (Path B). |
| SmartScreen / Defender blocks the app | Unsigned binary. **More info → Run anyway.** |
| Chat says no funds / can't open session | Network models need **Base ETH** (gas) + **≥5 MOR** (stake). Or use **Path B** (free local model, no funds needed). |
| I have two wallets and I'm confused | The app wallet and any CLI/`WALLET_PRIVATE_KEY` wallet are separate — import both into MetaMask to track them in one place. |
| Inference returns 401 / empty | Check the Basic Auth (`.cookie`) credentials and that the `session_id` header matches the one returned when you opened the session. |
| `mor-cli.exe` fails immediately, even on `--help` | Known bug on Windows ([issue #792](https://github.com/MorpheusAIs/Morpheus-Lumerin-Node/issues/792)). Use `morpheus-doctor` or raw `curl` (Path C) instead. |
| A "copy" button in the app (e.g. model ID) doesn't actually copy anything | Known UI bug ([issue #793](https://github.com/MorpheusAIs/Morpheus-Lumerin-Node/issues/793)). Retrieve the value via the API instead: `GET /blockchain/models`. |

---

## Where to go next

- **[app.mor.org](https://app.mor.org)** — use Morpheus from any device/browser.
- **[Official docs](https://nodedocs.mor.org)** — full consumer/provider guides.
- **[GitHub repo](https://github.com/MorpheusAIs/Morpheus-Lumerin-Node)** — run a provider, or contribute.