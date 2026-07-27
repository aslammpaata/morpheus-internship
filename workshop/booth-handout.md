# Booth Handout — one-page take-home

**Format:** single page (A4/Letter), print-friendly. This is the markdown source; drop the QR images
in where marked and export to PDF. Keep it to **one side** if possible; the "How it works" block is
optional back-side content.

> ⚠️ **Do not send to print until the onboarding risk in `booth-runbook.md` is resolved** — the
> "On your PC" instructions below assume the Desktop App's wallet/local-model flow works on a clean
> install. Confirm this on a real booth laptop first.

---

<!-- ================== FRONT ================== -->

# 🧠 Try Morpheus — Decentralized AI You Can Run Yourself

**AI inference with no single owner.** Models run across a peer-to-peer marketplace of GPU providers,
coordinated by smart contracts on Base. Use it in your browser, or run it on your own machine.

## Three ways to try it — pick one

### 1️⃣ On your phone — 30 seconds
Open **app.mor.org**, sign up, and start chatting.
`[ QR → https://app.mor.org ]`

### 2️⃣ On your PC — free & offline
1. Download the **desktop app** (`win-x64-morpheus-app`) from the releases page and install it.
2. Open it, go to **Chat**, and select the **bundled local model**.
3. Chat with a model running **on your own machine** — no wallet, no tokens, works offline.
`[ QR → https://github.com/MorpheusAIs/Morpheus-Lumerin-Node/releases ]`

> Same install also gets you onto the full network — fund a wallet with a little MOR + ETH to open a
> session with any model, right from the same app.

### 3️⃣ For developers — OpenAI-compatible API
Point any OpenAI client at **`https://api.mor.org/api/v1`** (get a key at app.mor.org).
Running a local node? The proxy-router serves the same API at `localhost:8082`.
`[ QR → https://nodedocs.mor.org ]`

---

**⭐ Star the repo:** github.com/MorpheusAIs/Morpheus-Lumerin-Node
**💬 60-second feedback (+ swag):** `[ QR → feedback form ]`
**👋 Questions? Talk to us at the booth.**

<!-- ================== BACK (optional) ================== -->

---

## How it works (optional back side)

```
You  →  Morpheus app / proxy-router  →  Base blockchain (session + payment)  →  GPU provider  →  model reply
```

- **Consumers** stake **MOR** to open a session and pay per use.
- **Providers** register models and earn for serving inference.
- The **proxy-router** matches you to a provider and streams the response back — OpenAI-compatible,
  so existing tools work unchanged.

**Want to provide compute and earn?** See the Provider quickstart at **nodedocs.mor.org**.

---

### Production notes (not printed)

- Generate QR codes for: `app.mor.org`, the Releases page, `nodedocs.mor.org`, and your feedback form.
- There is **one Windows download** (the installer) — no separate archive build exists. Don't print
  or QR-link anything implying otherwise.
- Match Morpheus brand colors/logo when you lay it out for print.