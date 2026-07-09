# Morpheus Booth Demo — Video Script & Shot List

**Purpose:** (a) loops silently on the booth monitor to draw people in, and (b) a shareable
"zero-to-inference" tutorial for social/GitHub. **Target length:** ~3–3.5 min.
**Format:** 1080p screen capture. **Audio optional** — design **on-screen captions** so it works
muted on the booth loop.

---

## Pre-record checklist (do NOT skip — security)

- [ ] Use a **scoped `chat` user**, not the admin `.cookie`. **Never show a real password or wallet
      seed phrase on screen** (AUDIT F9). Set auth via env var off-camera.
- [ ] `mor-launch local` model **pre-downloaded** so there's no wait on camera.
- [ ] `morpheus-doctor.exe` already **built**; a terminal open in its folder.
- [ ] Browser signed into a **demo `app.mor.org`** account; clean desktop; hide personal info.
- [ ] Screen resolution 1920×1080; font size bumped in terminal for legibility on a booth TV.

---

## Shot list

| Time | On screen | Voiceover / caption |
|------|-----------|---------------------|
| **0:00–0:20 — Hook** | Split text card: "Setting this up the hard way: **~1 week**" → "The easy way: **60 seconds**". | "This is AI that no company owns — and you can run it yourself. Here are three ways to try it, fastest first." |
| **0:20–1:00 — Tier 0: any device** | Phone (or mobile-sized browser) opens **app.mor.org**, type a prompt, answer streams in. | "On any phone or browser, open app.mor.org and chat. That answer routed through a decentralized marketplace of GPU providers — not one company's servers." |
| **1:00–1:50 — Tier 1: free & local** | Terminal in the extracted archive folder: type `mor-launch.exe local`. App/UI opens, chat with the local model — then **toggle wifi OFF**, chat again, still works. | "Want it fully in your control? One command runs a model on **your** machine — free, no wallet, and — watch — it keeps working with the **wifi off**." |
| **1:50–3:00 — Tier 2: developers** | Terminal: run `morpheus-doctor --dev --prompt "Hello!"`. Show the output lines: ✔ router reachable → ✔ session opened → ✔ inference ok → reply. | "For developers: this one command opens a paid session, grabs the session ID, and runs inference — the four manual API calls, collapsed into one. And the API is OpenAI-compatible, so your existing code just points at Morpheus." |
| **3:00–3:30 — What it is + CTA** | Simple diagram: You → proxy-router → Base blockchain → providers. End card with **three QR codes** (app.mor.org · GitHub · feedback) and booth location. | "Decentralized inference, coordinated on Base. Scan to try it, star the repo, or come talk to us at the booth." |

---

## Silent-loop version (booth monitor)

- Cut a **60–75s** version: Hook → Tier 0 → the **wifi-off** moment (Tier 1) → end card.
- Burn in **large captions** for every step; assume no sound.
- Loop it; the wifi-off beat is the scroll-stopper — lead the loop with it if needed.

## Production notes

- Keep each terminal command **on screen long enough to read** (2–3s min).
- Highlight the returned model reply (zoom or box) so the "it actually answered" moment lands.
- Export: MP4 (H.264), plus a muted GIF of the `morpheus-doctor --dev` run for the repo README.
