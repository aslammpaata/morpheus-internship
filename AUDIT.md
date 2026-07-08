# Windows Onboarding Friction Audit — Morpheus-Lumerin-Node

**Status:** Living document (audit in progress)
**Started:** 2026-07-08
**Auditor:** MorpheusAI intern (firsthand run) + repo review
**Target:** [MorpheusAIs/Morpheus-Lumerin-Node](https://github.com/MorpheusAIs/Morpheus-Lumerin-Node), release v7.3.0 (2026-06-24)

---

## Purpose

Catalog every friction point on the **zero-to-first-inference** path on Windows, so we can fix
discovery/docs, and automate the painful developer flow. Each entry is a candidate for one of:
a documentation PR, a GitHub issue, or a feature of `morpheus-doctor`.

---

## The two paths

| Path | Audience | Verdict |
|------|----------|---------|
| **Desktop app** (`win-x64-morpheus-app-7.3.0.exe`) | Beginners / most users | Bundles wallet + chat + built-in router. Staking & sessions are smooth. **Already good.** |
| **Router + CLI** | Developers / power users | Greater session-management control, but a painful multi-step manual chain. **This is where the automation value is.** |

---

## Friction log

| # | Step | Friction | Sev | Owner |
|---|------|----------|-----|-------|
| F1 | Discovering the path | README/getting-started reads "dev / build from source"; the recommended desktop `.exe` isn't surfaced up top | High | Doc PR |
| F2 | Router+CLI: start router | Router must be started manually as a separate process before anything works | Med | Tool |
| F3 | Router+CLI: open session | Session/stake opened via raw `curl` calls | High | Tool |
| F4 | Router+CLI: session → inference | **Session ID must be hand-copied** from the open-session response into a separate inference command | High | Tool ⭐ |
| F5 | Two wallets | CLI auto-creates one wallet; desktop uses another — confusing until both are imported into MetaMask | Low | Doc note |
| F6 | Design axis | Desktop = easy but less control; CLI = powerful but painful (not a bug — defines the two tracks) | — | Design |
| F7 | Router+CLI on Windows | Documented commands are **Linux/systemd** (`systemctl --user`, `*.sh`, `node *.mjs`) — no native-Windows equivalent documented; likely requires WSL | High | Doc + Tool |
| F8 | Wallet setup command | Uses a **bespoke `everclaw` script**, not a canonical Morpheus command — cannot be recommended to the community as-is | Med | Investigate |
| F9 | Credential handling | The router API Basic Auth password is naturally copy-pasted verbatim; no guidance that it's a secret to protect/rotate | Med | Doc + Tool |

### What works well (do NOT "fix")
- Wallet **funding** was easy (Base ETH + MOR).
- **Desktop-app staking** is straightforward.

---

## The manual developer chain (the thing we automate)

Confirmed environment: **mainnet (Base)**, model **Kimi K2.5**
(`model_id 0xbb9e920d94ad3fa2861e1e209d0a969dbe9e1af1cf1ad95c49f76d7b63d32d93`),
router API on **`localhost:8082`**, session duration **600s (10 min)**.

> **SECURITY:** the `admin:...` value below is the router API password. It is a **secret**.
> It is redacted here as `<ROUTER_API_PASSWORD>` and must never be committed.

```bash
# 1. Start the router (Linux/systemd — see F7)
systemctl --user start morpheus-router
#    or: ~/morpheus/start-router.sh

# 2. Set up the wallet (bespoke everclaw script — see F8)
node skills/everclaw/scripts/everclaw-wallet.mjs setup

# 3. Open a session (stake MOR) — returns a sessionID
curl -s -u "admin:<ROUTER_API_PASSWORD>" \
  -X POST "http://localhost:8082/blockchain/models/0xbb9e920d94ad3fa2861e1e209d0a969dbe9e1af1cf1ad95c49f76d7b63d32d93/session" \
  -H "Content-Type: application/json" \
  -d '{"sessionDuration": 600}'
# -> { "sessionID": "..." }   <-- F4: this must be hand-copied into step 4

# 4. Run inference using the session ID
curl -s -u "admin:<ROUTER_API_PASSWORD>" \
  -X POST "http://localhost:8082/v1/chat/completions" \
  -H "Content-Type: application/json" \
  -H "session_id: <SESSION_ID>" \
  -H "model_id: 0xbb9e920d94ad3fa2861e1e209d0a969dbe9e1af1cf1ad95c49f76d7b63d32d93" \
  -d '{"model":"kimi-k2.5","messages":[{"role":"user","content":"Hello!"}],"stream":false}'
```

**`morpheus-doctor --dev` collapses steps 1, 3, and 4 into one command**, capturing the
session ID in memory (kills F4) and health-checking the router (kills F2).

---

## Open questions (need answers to finish the audit)

1. **WSL or native Windows?** Are the router+CLI commands run under WSL, or on a Linux box?
   Determines whether the dev path on real Windows needs a WSL guide or native `.exe` instructions. (F7)
2. **Canonical wallet command?** Is `everclaw` the intended wallet tool, or is there a stock
   Morpheus CLI wallet command we should document instead? (F8)
3. **Timing:** rough end-to-end minutes for the router+CLI path (for the "why the tool matters" story).
4. Does `win-x64-morpheus-cli-7.3.0.exe` expose the same session/inference flow natively on Windows?
