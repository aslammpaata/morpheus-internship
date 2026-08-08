# Windows Onboarding Friction Audit — Morpheus-Lumerin-Node

**Status:** Living document (audit in progress)
**Started:** 2026-07-08
**Last updated:** 2026-07-14
**Auditor:** MorpheusAI intern (firsthand run) + repo & official-docs review
**Target:** [MorpheusAIs/Morpheus-Lumerin-Node](https://github.com/MorpheusAIs/Morpheus-Lumerin-Node), release v7.3.0 (2026-06-24)
**Docs reviewed:** [nodedocs.mor.org](https://nodedocs.mor.org) — quickstart-consumer, reference/api-auth, reference/api-auth

---

## Purpose

Catalog every friction point on the **zero-to-first-inference** path on Windows, so we can fix
discovery/docs and automate the painful developer flow. Each entry is a candidate for a
documentation PR, a GitHub issue, or a feature of `morctl`.

---

## The paths (corrected 2026-07-14 — see F12)

| Path | How you actually run it (verified) | Audience | Verdict |
|------|-------------------------------------|----------|---------|
| **Installer app** (`win-x64-morpheus-app-*.exe`) | Single-file installer → installs to `%APPDATA%\morpheus-app\`; launch from **Start-menu/desktop shortcut** (router auto-starts). Bundles a **free local llama.cpp model** (`services\ai-model.gguf`) usable with no wallet/MOR. | Beginners | Smooth. **Already good, and the only officially-shipped Windows path.** |
| **Headless proxy-router** (`win-x64-morpheus-router-*.exe`) | Standalone binary, no GUI. Requires manual `.env`/`models-config.json` setup. | Providers, servers | Correct for its audience; not a consumer path. |
| **`mor-cli`** (`win-x64-morpheus-cli-*.exe`) | Standalone CLI binary. **Currently broken on Windows — see F15.** | Scripting/automation (intended) | **Not usable today.** |
| ~~Archive build (`.zip`, `mor-launch.exe`)~~ | **Does not exist in any release checked (v7.1.12 → v7.3.3).** Referenced in `docs/04-consumer-setup.md` and the provider quickstart, but no such asset ships. | — | **Finding, not a real path — see F12.** |

---

## Friction log

| # | Step | Friction | Sev | Status / Owner |
|---|------|----------|-----|----------------|
| F1 | Discovering the path | README/getting-started reads "dev / build from source"; the recommended desktop `.exe` isn't surfaced up top. (Note: original wording also cited `mor-launch.exe` as under-surfaced — drop that; see F12.) | High | Open — Doc PR (**PR-01 submitted**) |
| F2 | Start the router | Auditor started it manually via `systemctl`; **canonically unnecessary** — the Desktop App starts it automatically | Med | ✅ Resolved; Doc + Tool |
| F3 | Open session | Session/stake via raw `curl` for API users | High | ✅ Automated — `morctl --dev` |
| F4 | Session → inference | **Session ID hand-copied** from open-session response into a separate inference call | High | ✅ Resolved — `morctl --dev` captures it in memory |
| F5 | Two wallets | CLI-created wallet vs desktop wallet — confusing until both imported into MetaMask | Low | Open — Doc note |
| F6 | Design axis | Desktop = easy but less control; API = powerful but manual (defines the two tracks) | — | Design |
| F7 | Windows run method | Auditor's `systemctl`/`*.sh` path is **not** the documented Windows path — canonical is the Desktop App's Start-menu shortcut (no systemd, no `mor-launch.exe`) | High | ✅ Resolved — Doc |
| F8 | Wallet setup command | Canonical wallet = **in-app** (set password → create/recover via mnemonic) OR headless via **`WALLET_PRIVATE_KEY`** env (falls back to system keychain when unset). The `everclaw` script is **bespoke**, not recommendable to community | Med | ✅ Resolved — Doc |
| F9 | Credential handling | Router password = auto-generated **`.cookie`** (`admin:<random>`), stored under `%APPDATA%\morpheus-app\services\.cookie` for the installer build. Docs: *"Rotate the admin password regularly."* API supports **scoped users** (`POST /auth/users`, 29 perms incl. `open_session`, `chat`) | Med | ✅ Resolved — Doc + Tool (create scoped user for booth) |
| **F10** | **Discoverability of golden path** | A reasonable intern ended up on the HARD path (Linux/systemd/everclaw) while the Desktop App does it in one double-click. Strong evidence newcomers will mis-path. (Corrected: the easy path is the **installer**, not `mor-launch.exe` — see F12.) | **High** | Open — Doc PR (**PR-01 submitted**, highest ROI) |
| **F11** | **(Opportunity)** | The Desktop App installer bundles a **free local llama.cpp model** (`services\ai-model.gguf`) → inference with **no wallet, no MOR, no staking**, selectable directly in the Chat tab. (Corrected: **not** archive-only — it's part of the standard installer.) | — | Workshop + Doc |
| **F12** | **Archive build does not exist (was: "two Windows distributions")** | Official docs (`docs/04-consumer-setup.md`, provider quickstart at nodedocs.mor.org) describe an "extract the `.zip`, run `mor-launch.exe`" flow. **Verified against every release from v7.1.12 through v7.3.3: no `.zip` archive asset exists.** Current releases ship only: Desktop App installer, headless proxy-router binary, CLI binary, Docker image, TEE compose. Following the documented archive instructions is a dead end — there is no file to download. This is stale documentation, not a fabrication or user error. | **High** | Open — Doc PR (docs need to drop the archive/`mor-launch` references or the team needs to reinstate the archive build) |
| **F13** | **Model-card "copy" button silently fails** | In the Models tab, clicking the copy icon next to a model's `ID` shows a success toast but does not actually populate the clipboard (confirmed: pasting after the click yields nothing). Forces users to retrieve the full model ID via the router API instead (`GET /blockchain/models`). | Low–Med | Open — bug report |
| **F14** | **Chat sessions do not share conversation context by default** | Router config flags `StoreChatContext`/`ForwardChatContext` do **not** mean the router remembers prior turns in a session — they govern chat logging, not model memory. Confirmed via direct test: repeated prompts in the same `session_id` with only the latest message sent produced zero recollection of earlier turns ("I don't have access to your name..."). Standard OpenAI-style behavior (client resends full history) is required and is **not** the default assumption a new API user would make. Worth a clear docs callout on `/v1/chat/completions`. `morctl --interactive` implements client-side history correctly as a reference. | Med | Open — Doc PR |
| **F15** | **`mor-cli` fails to launch on Windows (cookie path bug)** | `mor-cli.exe --help` (and any subcommand) fails unconditionally with: `can't read cookie file: open <dir>\<dir>\.cookie: The filename, directory name, or volume label syntax is incorrect.` The printed path is the router's base directory duplicated onto itself. **Confirmed independent of**: working directory, `COOKIE_FILE_PATH` env var, `AUTH_USER`/`AUTH_PASSWORD` env vars, and presence of a valid `.cookie` file at the expected location (5 separate test configurations, all producing byte-identical output). The failure occurs during startup/config resolution, before any documented override (`COOKIE_FILE_PATH`, `AUTH_CONFIG_FILE_PATH` per `/reference/api-auth`) is consulted. Tested binary: `win-x64-morpheus-cli-7.3.0.exe`. **`mor-cli` is currently unusable out of the box on Windows — cannot even print help text.** | **High** | Open — bug report (blocking) |
| **F16** | **Desktop App onboarding completely broken** | Fresh install, no prior state, wallet setup fails unconditionally: "Failed to finish onboarding." Confirmed on multiple machines and multiple versions (v7.3.0 and an earlier tested release). Also reproduces when switching an existing wallet on an already-onboarded install. Blocks every new user at step one. | **Critical** | ✅ Resolved — #811, confirmed working on a clean install as of [27/07/2026] |

### What works well (do NOT "fix")
- Wallet **funding** was easy (Base ETH + MOR).
- **Desktop-app staking** is straightforward *once a wallet is successfully onboarded* — but onboarding itself is currently broken (F16).- The Desktop App's **bundled local model** gives a zero-cost inference demo out of the box.
- `morctl --dev`/`--interactive` — tested end-to-end against two models, graceful failure handling, correct client-side context.

---

## The manual developer chain (what `morctl --dev` automates)

Confirmed env: **mainnet (Base, chain 8453)**, model **Kimi K2.5**
(`model_id 0xbb9e920d94ad3fa2861e1e209d0a969dbe9e1af1cf1ad95c49f76d7b63d32d93`),
router API **`localhost:8082`**, session **600s (10 min)**.

> **SECURITY:** the `admin:...` value is the router's `.cookie` password — a **secret**.
> Redacted as `<ROUTER_API_PASSWORD>`; never commit. Prefer a **scoped user** (F9) over admin.

```bash
# 1. Start the router
#    Canonical Windows: launch the Desktop App from the Start-menu shortcut (router auto-starts)
#    Auditor's non-standard path:  systemctl --user start morpheus-router

# 2. Wallet
#    Canonical: in-app (set password -> create/recover via mnemonic)   (see F8)
#    Auditor's bespoke path:  node skills/everclaw/scripts/everclaw-wallet.mjs setup

# 3. Open a session (stake MOR) — returns a sessionID
curl -s -u "admin:<ROUTER_API_PASSWORD>" \
  -X POST "http://localhost:8082/blockchain/models/0xbb9e...32d93/session" \
  -H "Content-Type: application/json" \
  -d '{"sessionDuration": 600}'
# -> { "sessionID": "..." }   <-- F4: must be hand-copied into step 4

# 4. Run inference using the session ID
curl -s -u "admin:<ROUTER_API_PASSWORD>" \
  -X POST "http://localhost:8082/v1/chat/completions" \
  -H "Content-Type: application/json" \
  -H "session_id: <SESSION_ID>" \
  -H "model_id: 0xbb9e...32d93" \
  -d '{"model":"kimi-k2.5","messages":[{"role":"user","content":"Hello!"}],"stream":false}'
```

**`morctl --dev`** (✅ built and tested) health-checks the router, opens the session,
**captures `sessionID` in memory (kills F4)**, runs inference. **`morctl --dev --interactive`**
(✅ built and tested) additionally maintains full conversation history client-side and resends it each
turn (fix for F14), since the router does not retain it. A scoped `chat`/`open_session` user (F9) as
an alternative to admin remains a future enhancement, not yet built.

---

## Open questions

1. ✅ **WSL or native — answered.**
2. ✅ **Booth machines — answered.**
3. ✅ **Local free model — answered, differently than expected.** There is no `mor-launch.exe` to
   test (see F12). The free local model is bundled directly in the Desktop App installer
   (`services\ai-model.gguf`) and is reachable from the Chat tab with no separate download. Verified
   working.
4. ✅ **CLI wallet env var — answered:** `WALLET_PRIVATE_KEY` (falls back to system keychain). F8 closed.
5. ✅ **Timing — answered:** the hand-rolled router+CLI (WSL) path took **~1 week** vs **one
   double-click** for the desktop app. Headline stat for F10.
6. ⏳ **New — is `mor-cli` broken on other platforms too, or Windows-only?** Not tested on
   macOS/Linux. Worth a maintainer confirming scope before assuming Windows-specific.

## Canonical config surface (for morctl)

From `/reference/env-proxy-router`:
- `WEB_ADDRESS` (default `0.0.0.0:8082`) — local API bind
- `WALLET_PRIVATE_KEY` — headless wallet
- `ETH_NODE_ADDRESS` (RPC), `ETH_NODE_CHAIN_ID` (`8453` main / `84532` test)
- `DIAMOND_CONTRACT_ADDRESS` (main `0x6aBE1d282f72B474E54527D93b979A4f64d3030a`)
- `MOR_TOKEN_ADDRESS` (main `0x7431aDa8a591C955a994a21710752EF9b882b8e3`)
- Auth: `.cookie` (default `./.cookie`, `admin:<pw>`), seed via `COOKIE_CONTENT`; `proxy.conf`
  (`AUTH_CONFIG_FILE_PATH`) holds `rpcauth=`/`rpcwhitelist=`. **`mor-cli` does not currently respect
  `COOKIE_FILE_PATH` — see F15.**