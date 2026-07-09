# Windows Onboarding Friction Audit — Morpheus-Lumerin-Node

**Status:** Living document (audit in progress)
**Started:** 2026-07-08
**Auditor:** MorpheusAI intern (firsthand run) + repo & official-docs review
**Target:** [MorpheusAIs/Morpheus-Lumerin-Node](https://github.com/MorpheusAIs/Morpheus-Lumerin-Node), release v7.3.0 (2026-06-24)
**Docs reviewed:** [nodedocs.mor.org](https://nodedocs.mor.org) — quickstart-consumer, reference/api-auth

---

## Purpose

Catalog every friction point on the **zero-to-first-inference** path on Windows, so we can fix
discovery/docs and automate the painful developer flow. Each entry is a candidate for a
documentation PR, a GitHub issue, or a feature of `morpheus-doctor`.

---

## The paths (updated with canonical docs)

| Path | How you actually run it (canonical) | Audience | Verdict |
|------|-------------------------------------|----------|---------|
| **Installer app** (`win-x64-morpheus-app-*.exe`) | Single-file installer → installs to `%APPDATA%\morpheus-app\`; launch from **Start-menu/desktop shortcut** (router auto-starts). **No `mor-launch.exe`** in this build (F12) | Beginners | Smooth. **Already good.** |
| **Archive build** (`.zip`, contains `mor-launch.exe`) | Extract, then run `mor-launch.exe` (starts proxy-router + CLI + UI). `mor-launch local` starts a **free bundled llama.cpp model** — offline, no wallet | Beginners (free trial) & devs | The free/offline path; **most people never find it (F10)** |
| **Hand-rolled router + CLI** (`systemctl`, `everclaw`) | What the auditor did in WSL — Linux/systemd service + bespoke wallet script + raw curl | Advanced / non-standard | Painful; **non-canonical** — see F7/F8/F10 |

---

## Friction log

| # | Step | Friction | Sev | Status / Owner |
|---|------|----------|-----|----------------|
| F1 | Discovering the path | README/getting-started reads "dev / build from source"; the recommended desktop `.exe` and `mor-launch.exe` aren't surfaced up top | High | Open — Doc PR |
| F2 | Start the router | Auditor started it manually via `systemctl`; **canonically unnecessary** — `mor-launch.exe` starts it | Med | ✅ Resolved (use mor-launch); Doc + Tool |
| F3 | Open session | Session/stake via raw `curl` for API users | High | Open — Tool |
| F4 | Session → inference | **Session ID hand-copied** from open-session response into a separate inference call | High | Open — Tool ⭐ |
| F5 | Two wallets | CLI-created wallet vs desktop wallet — confusing until both imported into MetaMask | Low | Open — Doc note |
| F6 | Design axis | Desktop = easy but less control; API = powerful but manual (defines the two tracks) | — | Design |
| F7 | Windows run method | Auditor's `systemctl` / `*.sh` path is **not** the documented Windows path — canonical is `mor-launch.exe` (no systemd) | High | ✅ Resolved — Doc |
| F8 | Wallet setup command | Canonical wallet = **in-app** (set password → create/recover via mnemonic) OR headless via **`WALLET_PRIVATE_KEY`** env (falls back to system keychain when unset). The `everclaw` script is **bespoke**, not recommendable to community | Med | ✅ Resolved — Doc |
| F9 | Credential handling | Router password = auto-generated **`.cookie`** (`admin:<random>`). Docs: *"Rotate the admin password regularly. Keep `proxy.conf` out of version control."* API supports **scoped users** (`POST /auth/users`, 29 perms incl. `open_session`, `chat`) | Med | ✅ Resolved — Doc + Tool (create scoped user for booth) |
| **F10** | **Discoverability of golden path** | A reasonable intern ended up on the HARD path (Linux/systemd/everclaw) while `mor-launch.exe` exists and does it in one double-click. Strong evidence newcomers will mis-path. | **High** | Open — Doc PR (highest ROI) |
| **F11** | **(Opportunity)** | `mor-launch local` runs a **free bundled llama.cpp model** → inference with **no wallet, no MOR, no staking**. **Only in the archive (.zip) build** (see F12). Ideal booth beginner demo. | — | Workshop + Doc |
| **F12** | **Two Windows distributions** | **Single-file installer** (`win-x64-morpheus-app-*.exe`) → `%APPDATA%\morpheus-app\`, launched from Start-menu shortcut, router auto-starts, **no `mor-launch.exe`**. **Archive (.zip)** → extract-and-run, ships `mor-launch.exe` + the `local` free-model flow. Docs' `mor-launch local` commands assume the archive build, so **installer users hit "mor-launch.exe not found."** | High | Open — Doc PR |

### What works well (do NOT "fix")
- Wallet **funding** was easy (Base ETH + MOR).
- **Desktop-app staking** is straightforward.
- **`mor-launch local`** gives a zero-cost inference demo out of the box.

---

## The manual developer chain (what `morpheus-doctor --dev` automates)

Confirmed env: **mainnet (Base, chain 8453)**, model **Kimi K2.5**
(`model_id 0xbb9e920d94ad3fa2861e1e209d0a969dbe9e1af1cf1ad95c49f76d7b63d32d93`),
router API **`localhost:8082`**, session **600s (10 min)**.

> **SECURITY:** the `admin:...` value is the router's `.cookie` password — a **secret**.
> Redacted as `<ROUTER_API_PASSWORD>`; never commit. Prefer a **scoped user** (F9) over admin.

```bash
# 1. Start the router
#    Canonical Windows: double-click mor-launch.exe   (NOT systemctl — see F7)
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

**`morpheus-doctor --dev`** health-checks the router (or launches it), opens the session,
**captures `sessionID` in memory (kills F4)**, runs inference, and can create a **scoped `chat`
user (F9)** instead of using admin.

---

## Open questions

1. ✅ **WSL or native — answered:** auditor develops on a **Windows WSL** client via PowerShell;
   the `systemctl`/`.sh` path lives in WSL, while the easy path (double-click desktop `.exe`) is what
   they use to actually run it. Strongly confirms **F10**: even a capable intern split across two paths
   and did the hard one in WSL.
2. ✅ **Booth machines — answered:** booth's own laptops are **majority Windows**, but attendees will
   bring **mixed devices** (Mac/Linux/phones). Implication: attendees can't be assumed to install
   anything → lead with a **no-install path (Hosted Inference API / web)** for their devices, and use
   the booth's Windows laptops for hands-on desktop-app / `mor-launch local` demos.
3. ⏳ **`mor-launch local` — not yet tried.** Auditor only double-clicks the desktop `.exe`. Worth a
   ~10-min test; if it works it becomes the zero-fund beginner booth demo (F11).
4. ✅ **CLI wallet env var — answered:** `WALLET_PRIVATE_KEY` (falls back to system keychain). F8 closed.
5. ✅ **Timing — answered:** the hand-rolled router+CLI (WSL) path took **~1 week** to fully set up
   and run, vs **one double-click** for the desktop app. This is the headline stat for the "why
   discoverability matters" story (F10).

## Canonical config surface (for morpheus-doctor)

From `/reference/env-proxy-router`:
- `WEB_ADDRESS` (default `0.0.0.0:8082`) — local API bind
- `WALLET_PRIVATE_KEY` — headless wallet
- `ETH_NODE_ADDRESS` (RPC), `ETH_NODE_CHAIN_ID` (`8453` main / `84532` test)
- `DIAMOND_CONTRACT_ADDRESS` (main `0x6aBE1d282f72B474E54527D93b979A4f64d3030a`)
- `MOR_TOKEN_ADDRESS` (main `0x7431aDa8a591C955a994a21710752EF9b882b8e3`)
- Auth: `.cookie` (default `./.cookie`, `admin:<pw>`), seed via `COOKIE_CONTENT`; `proxy.conf` (`AUTH_CONFIG_FILE_PATH`) holds `rpcauth=`/`rpcwhitelist=`
