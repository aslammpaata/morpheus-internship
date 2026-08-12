---
title: "Native Path: morctl — the Morpheus Developer CLI"
description: "A hands-on walkthrough of morctl, the fastest way to test sessions and inference from the command line."
audience: ["developer"]
last_verified: "morctl v1 (renamed from morpheus-doctor)"
---

# Native Path: `morctl` — the Morpheus Developer CLI

This guide walks through `morctl` the same hands-on way as the
[Desktop App walkthrough](./desktop-app-beginner-guide.md).

**What you'll do:** Build/download → point it at your router → preflight check → open a
session → chat → understand what happens when you're done.

**Time:** about 5 minutes, plus however long you spend chatting.

**If you just want the raw HTTP calls `morctl` automates**, see the
[headless/developer curl guide](./developer-headless-guide.md) instead.

---

## Before You Start

**What `morctl` is, in one sentence:** a single command-line tool that replaces the
manual "open a session, copy the session ID, paste it into another command, run
inference" chain with one call.

**What you'll need:**
- A running Morpheus proxy-router (Desktop App or headless).
- A **funded wallet** — at least 5 MOR and a small amount of Base ETH.
- **Go 1.21+**, only if building from source.

> 🔐 **About official `mor-cli`:** has several confirmed, unresolved bugs on Windows
> ([issue #792](https://github.com/MorpheusAIs/Morpheus-Lumerin-Node/issues/792)).
> `morctl` is the recommended alternative until those are fixed.

---

## Step 1: Get the Binary

```bash
git clone https://github.com/aslammpaata/morctl.git
cd morctl
go build -o morctl .          # Linux/macOS
# or: go build -o morctl.exe .   # Windows
```
Cross-compile: `GOOS=windows GOARCH=amd64 go build -o morctl.exe .` ·
`GOOS=darwin GOARCH=arm64 go build -o morctl .` (Apple Silicon) ·
`GOOS=darwin GOARCH=amd64 go build -o morctl .` (Intel)

**📸 Capture 1 — silent successful build**, annotate: *"No output = it worked."*

✅ You have a binary.

---

## Step 2: Point It at Your Router's Credentials

```bash
export COOKIE_FILE_PATH="/path/to/.cookie"    # macOS/Linux
$env:COOKIE_FILE_PATH = "C:\path\to\.cookie"   # Windows
```

**📸 Capture 2 — `echo $COOKIE_FILE_PATH` confirming it's set.**

✅ Credentials in place.

---

## Step 3: Preflight Check

```bash
./morctl
```

**📸 Capture 3 — successful output**, box the three key lines (router reachable,
wallet balance, "Preflight OK").

✅ Ready for Step 4.

---

## Step 4: Open a Session and Run Inference

```bash
./morctl --dev --prompt "Hello!"
```

**📸 Capture 4 — the full run**, numbered arrows at: session opened → inference ok →
reply → session live countdown. This is the single most important capture in the guide.

✅ Core loop works.

---

## Step 5: Have a Real Conversation

```bash
./morctl --dev --interactive
```

**📸 Capture 5 — a short exchange showing memory working within the session.**

Type `exit`/`quit` to stop chatting locally — the session keeps running in the
background.

✅ Multi-turn conversation confirmed.

---

## Step 6: Session Reuse

Run `--dev` again before expiry:

**📸 Capture 6 — the "reusing cached session" line.**

Force fresh: `./morctl --dev --fresh`

> ⚠️ **Known limitation:** session *reuse* correctly avoids a new stake, but
> conversation *history* is not currently persisted across separate `morctl` runs —
> exiting and restarting resets chat memory even though the underlying session is the
> same one. This is being worked on; see the project roadmap.

✅ You understand when a session opens vs. reuses.

---

## Step 7: What Happens When You're Done

⚠️ **Updated — recently changed, more involved than before.**

- Letting a session expire naturally only auto-settles the small session **cost** —
  not your full stake.
- Your stake enters a **~24-hour hold** before it's claimable — a network-wide rule,
  not specific to `morctl`.
- **Recovering it currently requires a manual call** to `withdrawUserStakes` on the
  Diamond contract — not yet wired into `morctl` itself. Use
  [Basescan's Write tab](https://basescan.org/address/0x6abe1d282f72b474e54527d93b979a4f64d3030a#writeContract)
  with your address and a small iteration count.
  ([Tracking issue #827](https://github.com/MorpheusAIs/Morpheus-Lumerin-Node/issues/827)
  — an API for this is proposed but not yet shipped.)
- Exiting `--interactive` only stops the local loop — never closes the session early,
  so no early-close penalty either way.

✅ Full loop understood, including the currently-manual recovery step.

---

## Understanding morctl: A Tour of the Flags

| Flag | For | When |
|---|---|---|
| `--dev` | Full session→inference flow | Every real chat |
| `--interactive` | Multi-turn mode | More than one exchange |
| `--fresh` | Skip session cache | Force a new session |
| `--model` / `--model-id` | Choose model | Testing a specific model |
| `--duration` | Session length (seconds) | Shorter = smaller stake needed |
| `--json` | Raw request/response | Debugging |
| `--cookie` | Explicit `.cookie` path | One-off override |
| `--host` | Non-default router | Remote/custom router |

---

## Troubleshooting

| Symptom | Fix |
|---|---|
| `no router credentials` | Point `--cookie`/`COOKIE_FILE_PATH` correctly (Step 2). |
| `router unreachable` | Confirm the Desktop App or headless router is running. |
| `insufficient MOR balance` | Fund further, or use `--duration` with a smaller value. |
| `no bids available` | Wrong/inactive model ID. |
| Stake missing after session ends | Expected right now — see Step 7. |
| Chat history gone after restart | Known limitation — see Step 6. |

---

## Glossary

- **Session** — time-boxed access to a model.
- **Stake** — refundable deposit; currently subject to a ~24h hold + manual claim.
- **Preflight** — the no-flags reachability + balance check.
- **`.cookie` file** — the router's auto-generated auth credential.