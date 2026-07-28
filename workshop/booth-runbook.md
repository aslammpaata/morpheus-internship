# Morpheus Booth Runbook

**Event:** Web3 blockchain event — MorpheusAI booth, **Jul 30–31, 2026**
**Goal:** Get as many attendees as possible to *personally experience* decentralized AI, capture
leads/feedback, and drive signups + GitHub interest.
**Design principle:** attendees arrive on **mixed devices** (Windows/Mac/Linux/phones) and range from
total beginners to developers. The flow is **tiered** so anyone gets a "wow" in under a minute, and
the curious can go deeper.

---

## The hook (say this in one line)

> "This is AI that **no single company owns** — you can run it on your own device, and right now,
> for free. Want to try it?"

Keep the wallet/blockchain/staking talk **out** of the opening. Lead with the experience; reveal the
decentralization underneath *after* they're impressed.

---

## The three tiers

| Tier | Who | Device | Time | What they do | The "aha" |
|------|-----|--------|------|--------------|-----------|
| **0 — Try it now** | Everyone | Their phone / booth tablet | ~30–60s | Chat at **app.mor.org** (booth tablet is pre-logged-in for instant use; QR for take-home signup) | "I'm talking to decentralized AI from my phone" |
| **1 — Run it yourself** | Curious / beginners | Booth **Windows** laptop | ~5 min | Open the Desktop App, complete wallet setup (now confirmed working on a clean install — #811), and pick the bundled local model in the Chat tab for a free, offline demo — or fund it to try the full network. *(There is no `.zip`/`mor-launch.exe` archive build — confirmed absent from every release, [issue #796](https://github.com/MorpheusAIs/Morpheus-Lumerin-Node/issues/796). This is a single installer, not two builds.)* | "It's running locally, for free, no account" |
| **2 — Build on it** | Developers | Booth Windows laptop | ~10 min | `morpheus-doctor --dev` → live `session → inference` in one command; show the OpenAI-compatible `api.mor.org`; open a real staked session on Base | "One command, and I own the inference pipeline" |

**The story that ties it together (for Tier 2 / anyone technical):**
> "Setting this up the hard way took me **a week**. The easy path should be **one double-click** —
> we're still closing that gap; here's the tooling and the bugs we've found and reported along the way."

---

## Booth physical setup

- **1 tablet or laptop** at the front edge, browser open + **logged into `app.mor.org`** → hand to
  walk-ups for instant Tier-0 chat (no signup needed to *try*).
- **2–3 Windows laptops** pre-provisioned for Tier 1/2 (see checklist).
- **QR code cards** (printed, on the table):
  - `app.mor.org` — "Try it on your phone / make your own account"
  - GitHub repo — "Run a node / contribute"
  - **Feedback form** — "60-second feedback → [swag]"
- **Monitor/TV** looping the 3–5 min demo video.
- Stickers + a **one-page setup handout** (the beginner Windows setup guide).

---

## Pre-provisioning checklist (DO BEFORE THE EVENT — not at the booth)

Conference wifi will betray you. Everything that can be done offline, do in advance.

- ✅ Onboarding confirmed working on a clean install (#811) — still do one clean-install test per booth laptop as a sanity check, not a risk-gate.
- [ ] **Desktop App** installed on each Windows laptop; launched once so Defender/SmartScreen is
      already cleared.
- [ ] **Confirm the bundled local model works and is fully downloaded** — open Chat, select the
      local model, send a message, then **turn off wifi and send another** to confirm it truly runs
      offline. No model download should ever happen over conference wifi.
- [ ] **`morpheus-doctor.exe`** built (`GOOS=windows GOARCH=amd64`), tested, and on the desktop of
      each laptop. *(Confirmed working end-to-end — see `morpheus-doctor` repo.)*
- [ ] **A funded wallet** for the Tier-2 live-staking demo — small amount of **Base ETH + MOR**
      (enough for several 5–10 min sessions). Use a **dedicated demo wallet**, not a personal one.
      If Desktop App onboarding is broken, fund this wallet via the **headless router +
      `WALLET_PRIVATE_KEY`** path instead — proven working, doesn't touch the broken onboarding flow.
- [ ] **Scoped API user** created on the router (`open_session` + `chat` only) so no admin `.cookie`
      is ever shown on-screen. **Never display the admin password during a demo.**
- [ ] **`app.mor.org` booth demo account** created and logged in on the front tablet; an **API key**
      generated for the Tier-2 `api.mor.org` demo.
- [ ] **QR codes** printed and laminated; **feedback form** live; **handout** printed.
- [ ] **Full offline dry-run**: unplug wifi, confirm Tier 1 still works end-to-end.
- [ ] **`mor-cli` is not part of any demo** — confirmed broken on Windows and in its chat flow
      ([issue #792](https://github.com/MorpheusAIs/Morpheus-Lumerin-Node/issues/792)). Don't
      improvise it in on the day.

---

## Talk tracks (keep them short)

**Tier 0 (hand them the tablet):** "Type anything — ask it a question." … "That answer came through a
decentralized marketplace of GPU providers, not one company's servers. Want your own account? Scan
here." *(→ signup = lead)*

**Tier 1 (at a laptop):** "No account, no wallet — watch." Open the Chat tab, select the local model.
"That's a full AI model running on *this* laptop, for free, even with the wifi off. The same app can
also connect you to bigger models on the network when you want them."

**Tier 2 (developer):** "You like a terminal? This took me a week the hard way." Run
`morpheus-doctor --dev`. "One command: it started a session, staked, captured the session ID, and ran
inference — the thing that used to be four manual curl calls. And the API is OpenAI-compatible, so
your existing code just works by pointing at `api.mor.org`."

---

## Contingencies

| If… | Then… |
|-----|-------|
| **Wifi dies** | Tier 1 (local model) is the star — it's offline. Tell attendees "this is the point: it doesn't need the cloud." |
| **SmartScreen/Defender warning** | Pre-cleared on booth machines; for attendees' own machines, show them "More info → Run anyway" and note it's an unsigned-binary warning, not malware. |
| **Long line** | Front tablet + QR self-serve keeps Tier 0 flowing without staff. |
| **`app.mor.org` down / signup friction** | Fall back to Tier 1 local demo; collect leads via the feedback QR instead of signups. |
| **Someone wants to run a provider/earn** | Point to GitHub repo + provider quickstart; take their email for follow-up (out of scope for a booth). |

---

## Lead capture & adoption metrics

Track, with a simple tally sheet + digital form:
- Tier-0 chats handed out · **`app.mor.org` signups** · Tier-1 local runs · Tier-2 dev demos
- Feedback-form submissions · GitHub stars **before/after** the event · sessions opened on-chain
- Qualitative: top 3 questions asked, top confusion points (→ new AUDIT entries / doc PRs)

Post-event, compile these into an adoption report — doubles as evidence of impact for the internship
write-up.

---

## Open items / decisions needed

1. ~~Does `mor-launch local` work?~~ **Resolved: no such build exists.** Tier 1 uses the Desktop
   App's bundled local model instead.
2. ~~Is `morpheus-doctor --dev` built in time?~~ **Resolved: yes, built and thoroughly tested.**
3. Is Desktop App onboarding fixed? Resolved — confirmed working, #811.
4. Budget for the Tier-2 demo wallet (Base ETH + MOR).
5. Is `app.mor.org` chat free / does it include trial credits? — confirm at `apidocs.mor.org`.
6. Who owns the booth `app.mor.org` demo account and the demo wallet keys.

---

## Deliverables this runbook depends on

- Beginner Windows setup guide → the printed handout
- `morpheus-doctor` → the Tier-2 demo (built, tested, in its own repo)
- 3–5 min demo video → the looping monitor content