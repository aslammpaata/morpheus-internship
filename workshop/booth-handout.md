# Booth Handout — one-page take-home

**Format:** single page (A4/Letter), print-friendly. This is the markdown source; drop the QR images
in where marked and export to PDF. Keep it to **one side** if possible; the "How it works" block is
optional back-side content.

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
1. Download the **archive build** from the releases page.
2. Extract it, then run `mor-launch.exe local`.
3. Chat with a model running **on your own machine** — no wallet, no tokens, works offline.
`[ QR → https://github.com/MorpheusAIs/Morpheus-Lumerin-Node/releases ]`

> Just want the simplest install? Grab the **desktop app** (`win-x64-morpheus-app`) instead and launch
> it from the Start menu — it bundles the wallet, chat, and router.

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
- Verify the exact **archive `.zip` asset name** on the Releases page before printing (see setup guide).
- Consider a short link / vanity URL under each QR for people who prefer typing.
- Match Morpheus brand colors/logo when you lay it out for print.
