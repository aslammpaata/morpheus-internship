---
title: "Native Path: morctl — the Morpheus Developer CLI"
description: "A hands-on walkthrough of morctl, the fastest way to test sessions and inference from the command line."
audience: ["developer"]
last_verified: "morctl v1 (renamed from morpheus-doctor)"
---

# Native Path: `morctl` — the Morpheus Developer CLI

This guide walks through `morctl` the same hands-on way as the
[Desktop App walkthrough](./desktop-app-beginner-guide.md) — one step at a time, with a
terminal capture at each point so you know exactly what you should be seeing.

**What you'll do:** Build/download → point it at your router → preflight check → open a
session → chat → understand what happens when you're done.

**Time:** about 5 minutes, plus however long you spend chatting.

**If you just want the raw HTTP calls `morctl` automates**, see the
[headless/developer curl guide](./developer-headless-guide.md) instead — that's the right
doc if you're integrating Morpheus into your own tooling rather than using ours.

---

## Before You Start

**What `morctl` is, in one sentence:** a single command-line tool that replaces the
manual "open a session, copy the session ID, paste it into another command, run
inference" chain with one call — no copy-pasting, no juggling terminal output by hand.

**What you'll need:**
- A running Morpheus proxy-router — either the **Desktop App** (it starts its own router
  automatically) or a **headless router** you've started yourself. If you don't have
  either yet, follow the [Desktop App guide](./desktop-app-beginner-guide.md) or the
  [headless setup guide](./developer-headless-guide.md) first, then come back here.
- A **funded wallet** on that router — at least 5 MOR and a small amount of Base ETH.
- **Go 1.21+**, only if you're building from source. Skip this if you're using a
  prebuilt binary.

> 🔐 **About the official `mor-cli`:** you may run into Morpheus's own official CLI,
> shipped alongside the router. As of this writing it has several confirmed,
> unresolved bugs on Windows (fails to launch, slow model loading, unreliable chat —
> [issue #792](https://github.com/MorpheusAIs/Morpheus-Lumerin-Node/issues/792)).
> `morctl` is the recommended alternative until those are fixed — you're not missing
> anything by skipping straight to this guide.

---

## Step 1: Get the Binary

**Option A — build from source:**
```bash
git clone https://github.com/aslammpaata/morctl.git
cd morctl
go build -o morctl .          # Linux/macOS
# or: go build -o morctl.exe .   # Windows
```

**Option B — cross-compile for another platform** (e.g. building the Windows binary from
a Mac to hand to a colleague):
```bash
GOOS=windows GOARCH=amd64 go build -o morctl.exe .   # Windows
GOOS=darwin GOARCH=arm64 go build -o morctl .         # macOS Apple Silicon
GOOS=darwin GOARCH=amd64 go build -o morctl .         # macOS Intel
```

**Option C — prebuilt binary**, if one's attached to a release — download and, on
macOS/Linux, mark it executable: `chmod +x morctl`.

**📸 Capture 1 — the build command and its (silent) success.** A successful `go build`
prints nothing at all — worth showing that explicitly, since a beginner might expect
some confirmation text and wonder if it failed. Annotate: *"No output = it worked. Check
for the file with `ls` / `dir` next."*

> macOS: if you see a Gatekeeper warning on first run, right-click → **Open** → **Open
> anyway**, or run `xattr -d com.apple.quarantine ./morctl` in a terminal.

✅ You have a `morctl` (or `morctl.exe`) binary.

---

## Step 2: Point It at Your Router's Credentials

`morctl` needs the router's `.cookie` file, or an environment variable, to authenticate.

| Setup | Default `.cookie` location |
|---|---|
| Desktop App (Windows) | `%APPDATA%\morpheus-app\services\.cookie` |
| Desktop App (macOS) | `~/Library/Application Support/morpheus-app/services/.cookie` |
| Headless router | wherever your `.env`'s `COOKIE_FILE_PATH` points, often `./.cookie` |

Either pass it explicitly each time you run a command:
```bash
./morctl --cookie "/path/to/.cookie"
```
or set it once per terminal session so you don't have to repeat it:
```bash
export COOKIE_FILE_PATH="/path/to/.cookie"    # macOS/Linux
$env:COOKIE_FILE_PATH = "C:\path\to\.cookie"   # Windows PowerShell
```

**📸 Capture 2 — `echo $COOKIE_FILE_PATH` (or `echo $env:COOKIE_FILE_PATH` on Windows)**
showing the path is actually set. Annotate: *"Confirm this before moving on — a wrong or
missing path is the #1 cause of the errors in the Troubleshooting section below."*

✅ Your credentials are in place.

---

## Step 3: Preflight Check

Run with no flags to confirm everything's reachable before doing anything else:
```bash
./morctl
```

**📸 Capture 3 — successful preflight output.** Box the three key lines:
```
morctl  ->  http://localhost:8082
✔ router reachable
  wallet: 18.9442 MOR, 0.0021 ETH

Preflight OK. Re-run with --dev to open a session and run inference.
```
Callout: *"These three lines are your green light. If you see a low-balance warning
instead, fund your wallet before continuing — sessions need at least 5 MOR."*

✅ Router reachable, wallet funded, ready for Step 4.

---

## Step 4: Open a Session and Run Inference

```bash
./morctl --dev --prompt "Hello!"
```

**📸 Capture 4 — the full `--dev` run, top to bottom.** This is the single most
important capture in the guide — annotate each stage with a numbered arrow:
```
1→ > opening session (model kimi-k2.5, 600s)...
2→ ✔ session opened: 0x4c8e4022...
3→ > running inference...
4→ ✔ inference ok

     Hello! How can I help you today?

5→ ✅ session 0x4c8e4022... live (9m57s remaining)
```
Callout under arrow 2: *"This is the step that used to require a separate copy-paste —
now it's automatic."*

**Try a different model**, once you've seen the basic flow work once:
```bash
./morctl --dev --model "gpt-oss-120b:web" --model-id "0x2e7228..." --prompt "Say hi"
```
Find any model's full ID in the app's Models tab, or via `GET /blockchain/models` on the
router API.

✅ You've opened a session and received a reply — the core loop works.

---

## Step 5: Have a Real Conversation

A single `--dev` call is one question, one answer. For genuine back-and-forth:
```bash
./morctl --dev --interactive
```

**📸 Capture 5 — a short interactive exchange showing memory working**, e.g.:
```
you> My name is Alex and I'm testing morctl.
model> Nice to meet you, Alex! ...
you> What's my name?
model> Your name is Alex.
you> exit
```
Callout: *"morctl resends your whole conversation every turn behind the scenes — the
router itself doesn't remember previous messages on its own. This is what makes that
memory possible."*

Type `exit` or `quit` to stop chatting — the session keeps running in the background
until it naturally expires (Step 7 explains what happens then).

✅ You've had a multi-turn conversation with context retained correctly.

---

## Step 6: Session Reuse

Run `--dev` again against the same model before your first session expires:

**📸 Capture 6 — a second `--dev` run showing the reuse line:**
```
✔ reusing cached session: 0x4c8e4022... (7m12s remaining)
```
Callout: *"No new stake, no new session — just reused what was already open."*

Force a brand-new session regardless of the cache:
```bash
./morctl --dev --fresh
```

✅ You understand when a new session opens vs. when one's reused.

---

## Step 7: What Happens When You're Done

You don't need to close anything manually.

- **Let it expire naturally.** When the countdown from Capture 4 reaches zero, your node
  automatically submits the close transaction, and your staked MOR returns to your
  wallet within about a minute.
- **Exiting `--interactive`** (via `exit`/`quit`) only stops the local chat loop — it
  does **not** close the session early, so there's no early-close penalty to worry about.

✅ Full loop complete: built the tool, ran a session, chatted, and understand how your
stake comes back.

---

## Understanding morctl: A Tour of the Flags

Now that you've used the core flow once, here's what the rest of the flag set is for.

| Flag | What it's for | When you'd use it |
|---|---|---|
| `--dev` | Turns on the whole session→inference flow | Every time you actually want to chat, not just preflight |
| `--interactive` | Multi-turn conversation mode | When one question/answer isn't enough |
| `--fresh` | Skips the session cache, forces a new one | Testing a fresh session deliberately, or the cached one is behaving oddly |
| `--model` / `--model-id` | Choose which model to talk to | Testing a specific model instead of the default |
| `--duration` | How long a session should last (seconds) | Shorter sessions need a smaller stake — useful with a limited balance |
| `--json` | Print raw request/response JSON | Debugging, or understanding exactly what's sent over the wire |
| `--cookie` | Explicit path to your `.cookie` file | Overriding `COOKIE_FILE_PATH` for a one-off run |
| `--host` | Point at a router that isn't `localhost:8082` | Testing against a remote or non-default router |

---

## Troubleshooting & Common Questions

| Symptom | Fix |
|---|---|
| `no router credentials` | Point `--cookie` or `COOKIE_FILE_PATH` at the right `.cookie` file (Step 2). |
| `router unreachable` | Make sure the Desktop App or headless router is actually running. |
| `open session: ... insufficient MOR balance` | Fund your wallet with more MOR, or use `--duration` with a smaller value to reduce the required stake. |
| `open session: ... no bids available` | The model ID is wrong, or has no active providers right now — double-check the ID or try a different model. |
| Second `--dev` run opened a new session instead of reusing one | The cached session may have expired (check the remaining-time shown in Capture 4), or you used `--fresh`. |

---

## Glossary

- **Session** — a time-boxed window during which you can chat with a chosen model.
- **Stake** — a refundable deposit locked to open a session; not a payment.
- **Preflight** — the no-flags check that confirms the router is reachable and shows
  your wallet balance, before you commit to opening a session.
- **`.cookie` file** — the router's auto-generated password file, used to authenticate
  every `morctl` command against it.