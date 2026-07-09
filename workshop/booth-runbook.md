# Morpheus Booth Runbook

**Event:** Web3 blockchain event — MorpheusAI booth (Aug 2026 — _confirm exact name/date/location_)
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
| **1 — Run it yourself** | Curious / beginners | Booth **Windows** laptop | ~5 min | Double-click desktop app **or** `mor-launch local` (free bundled model, **works offline**) | "It's running locally, for free, no account" |
| **2 — Build on it** | Developers | Booth Windows laptop | ~10 min | `morpheus-doctor --dev` → live `session → inference`; show the OpenAI-compatible `api.mor.org`; open a real staked session on Base | "One command, and I own the inference pipeline" |

**The story that ties it together (for Tier 2 / anyone technical):**
> "Setting this up the hard way took me **a week**. The easy path is **one double-click**. That gap is
> the problem we've been closing — here's the tooling and docs that get you from zero to inference fast."

---

## Booth physical setup

- **1 tablet or laptop** at the front edge, browser open + **logged into `app.mor.org`** → hand to
  walk-ups for instant Tier-0 chat (no signup needed to *try*).
- **2–3 Windows laptops** pre-provisioned for Tier 1/2 (see checklist).
- **QR code cards** (printed, on the table):
  - `app.mor.org` — "Try it on your phone / make your own account"
  - GitHub repo — "Run a node / contribute"
  - **Feedback form** — "60-second feedback → [swag]"
- **Monitor/TV** looping the 3–5 min demo video (see task #4 sibling deliverable).
- Stickers + a **one-page setup handout** (= the beginner Windows setup guide, task #2).

---

## Pre-provisioning checklist (DO BEFORE THE EVENT — not at the booth)

Conference wifi will betray you. Everything that can be done offline, do in advance.

- [ ] **Desktop app** installed on each Windows laptop; launched once so Defender/SmartScreen is
      already cleared (the "Allow Defender if prompted" step is done).
- [ ] **`mor-launch local` tested** and the **local model pre-downloaded** on each laptop (so there's
      **no model download over conference wifi**). Confirm chat works with wifi **off**.
- [ ] **`morpheus-doctor.exe`** built (`GOOS=windows GOARCH=amd64`) and on the desktop of each laptop.
- [ ] **A funded wallet** for the Tier-2 live-staking demo — small amount of **Base ETH + MOR**
      (enough for several 5–10 min sessions). Use a **dedicated demo wallet**, not a personal one.
- [ ] **Scoped API user** created on the router (`open_session` + `chat` only) so no admin `.cookie`
      is ever shown on-screen (AUDIT F9). **Never display the admin password during a demo.**
- [ ] **`app.mor.org` booth demo account** created and logged in on the front tablet; an **API key**
      generated for the Tier-2 `api.mor.org` demo.
- [ ] **QR codes** printed and laminated; **feedback form** live; **handout** printed.
- [ ] **Full offline dry-run**: unplug wifi, confirm Tier 1 still works end-to-end.

---

## Talk tracks (keep them short)

**Tier 0 (hand them the tablet):** "Type anything — ask it a question." … "That answer came through a
decentralized marketplace of GPU providers, not one company's servers. Want your own account? Scan
here." *(→ signup = lead)*

**Tier 1 (at a laptop):** "No account, no wallet — watch." Double-click `mor-launch local`. "That's a
full AI model running on *this* laptop, for free, even with the wifi off. The same app can also
connect you to bigger models on the network when you want them."

**Tier 2 (developer):** "You like a terminal? This took me a week the hard way." Run
`morpheus-doctor --dev`. "One command: it started a session, staked, captured the session ID, and ran
inference — the thing that used to be four manual curl calls. And the API is OpenAI-compatible, so
your existing code just works by pointing at `api.mor.org`."

---

## Contingencies

| If… | Then… |
|-----|-------|
| **Wifi dies** | Make **Tier 1 `mor-launch local`** the star — it's offline. Tell attendees "this is the point: it doesn't need the cloud." |
| **SmartScreen/Defender warning** | Pre-cleared on booth machines; for attendees' own machines, show them the "More info → Run anyway" step and note it's an unsigned-binary warning, not malware. |
| **Long line** | Front tablet + QR self-serve keeps Tier 0 flowing without staff. |
| **`app.mor.org` down / signup friction** | Fall back to Tier 1 local demo; collect leads via the feedback QR instead of signups. |
| **Someone wants to run a provider/earn** | Point to GitHub repo + provider quickstart; take their email for follow-up (out of scope for a booth). |

---

## Lead capture & adoption metrics (feeds the reports deliverable)

Track, with a simple tally sheet + digital form:
- Tier-0 chats handed out · **`app.mor.org` signups** · Tier-1 local runs · Tier-2 dev demos
- Feedback-form submissions · GitHub stars **before/after** the event · sessions opened on-chain
- Qualitative: top 3 questions asked, top confusion points (→ new AUDIT entries / doc PRs)

Post-event, compile these into an **adoption report** (ties into `reports/`), which doubles as
evidence of impact for the internship write-up.

---

## Open items / decisions needed

1. **Event name, exact date, location** — and are booth laptops **provided or BYO**?
2. **Budget** for the Tier-2 demo wallet (Base ETH + MOR).
3. **Does `mor-launch local` actually work?** — pending your ~10-min test. Gates Tier 1.
4. **Is `app.mor.org` chat free / does it include trial credits?** — confirm at `apidocs.mor.org`;
   affects whether Tier 0 costs anything per message.
5. **Who owns the booth `app.mor.org` demo account** and the demo wallet keys.
6. **Is `morpheus-doctor --dev` built in time?** If not, Tier 2 falls back to the manual curl demo
   (still fine — the "week vs double-click" story carries it).

---

## Deliverables this runbook depends on (tracked)

- Beginner Windows setup guide → the printed **handout** (task #2)
- `morpheus-doctor` → the **Tier-2 demo** (task #3)
- 3–5 min demo video → the **looping monitor** content (task #4)
