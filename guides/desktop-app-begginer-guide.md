---
title: "Morpheus Desktop App — A Beginner's Walkthrough"
description: "A hands-on, screenshot-driven guide from download to your first AI reply — no crypto experience required."
audience: ["consumer", "beginner"]
last_verified: "v7.3.0"
---

# Morpheus Desktop App — A Beginner's Walkthrough

This guide walks you through the Morpheus Desktop App from a blank computer to your
first AI reply, one screen at a time.

**What you'll do:** Download → Install → Launch → Set up a wallet → Add a little
funding → Pick a model → Open a session → Chat → Understand what happens next.

**Time:** about 10 minutes, plus however long you spend chatting.

---

## Before You Start

**What Morpheus is, in one sentence:** it's a way to chat with AI models that run on
other people's computers around the world, instead of one company's servers.

**What you'll need:** Windows 10/11, an internet connection (skip this entirely with the
free offline local model, see Step 4's tip), and about 10 minutes.

> 🔐 This app creates a real crypto wallet. You'll get a clear warning at the exact
> moment this matters (Step 3).

---

## Step 1: Download & Install

1. Go to the [Releases page](https://github.com/MorpheusAIs/Morpheus-Lumerin-Node/releases)
   and download `win-x64-morpheus-app-<version>.exe`.

   **📸 Screenshot 1 — Releases page**, correct installer boxed, other assets marked
   "not this one."

2. Run it — Windows may show a SmartScreen warning; **More info → Run anyway**.

   **📸 Screenshot 2 — SmartScreen warning**, arrows to both buttons in order.

✅ Installed.

---

## Step 2: First Launch

Open Morpheus from the Start menu.

**📸 Screenshot 3 — first screen on launch.**

---

## Step 3: Set Up Your Wallet

1. Set a password.

   **📸 Screenshot 4 — password screen.**

2. **Create** (new user) or **Recover** (existing seed phrase).

   **📸 Screenshot 5 — Create vs. Recover.**

3. If creating: you'll see a **seed phrase**.

   **📸 Screenshot 6 — seed phrase screen.** 🔴 Heaviest annotation in the guide:
   *"Never share this. Never screenshot and post it. Write it down on paper."*

✅ Wallet exists, currently empty — expected.

---

## Step 4: Fund Your Wallet

| | For | Think of it as |
|---|---|---|
| ETH (Base) | Gas | The toll for using the road |
| MOR | Using the network | What you're here for |

1. Wallet tab shows `0 MOR / 0 ETH`.

   **📸 Screenshot 7 — unfunded wallet**, arrows to address/copy/Receive.

2. Send MOR + ETH, or use Receive's QR code.

   **📸 Screenshot 8 — funded wallet**, placed beside #7 for before/after.

✅ Funded.

> 💡 **Free path, no funding needed:** in Chat, pick the bundled **local model** — runs
> entirely on your machine.

---

## Step 5: Choose a Model

**📸 Screenshot 9 — Models → Registry**, circle 2-3 cards. Callout: *"Fee and Stake here
are the model's own listing bond — not what you'll pay to use it."*

---

## Step 6: Open a Session (Staking, Explained)

A **stake** is a refundable deposit, not a payment — like a hotel deposit.

1. Chat → Change Model → pick one.

   **📸 Screenshot 10 — model picker.**

2. Open Session → confirm.

   **📸 Screenshot 11 — stake confirmation**, if shown. Annotate: *"Refundable deposit,
   not a bill."*

3. Wallet tab shows available MOR dropped, Staked Balance up.

   **📸 Screenshot 12 — wallet right after staking.** 🔴 *"This is expected — your MOR
   is locked, not spent."*

4. Session header shows ID, timer, small cost figure.

   **📸 Screenshot 13 — session header**, arrows distinguishing timer/cost from stake.

✅ Session open.

---

## Step 7: Chat & Run Inference

**📸 Screenshot 14 — a reply in the chat window.**

✅ You've used decentralized AI.

---

## Step 8: What Happens When You're Done

⚠️ **Updated — this behavior recently changed and is more involved than it used to be.**

- **Let the session run out.** When the timer hits zero, your node auto-submits the
  close transaction after ~1 minute — but this **only settles the session cost**, not
  your full stake.
- **Your stake goes into a hold period (~24 hours)** before it's claimable. This is a
  recently-enforced network rule, not a bug in your setup.
- **Recovering it currently requires a manual step**: after ~24 hours, someone needs to
  call `withdrawUserStakes` directly on the Diamond contract — via
  [Basescan's Write tab](https://basescan.org/address/0x6abe1d282f72b474e54527d93b979a4f64d3030a#writeContract),
  connecting the same wallet, entering your address and a small iteration count (e.g.
  `1`–`3`). The app itself does not do this for you yet.
  ([Tracking issue #827](https://github.com/MorpheusAIs/Morpheus-Lumerin-Node/issues/827).)

**📸 Screenshot 15 — wallet balance before and after a manual claim**, if you choose to
walk through this step live in the guide.

✅ You understand the full loop, including the part that currently needs a manual step.

---

## Understanding the App: A Tour of Each Section

**📸 Screenshots 16a–16g — one per tab.**

- **Wallet** — balances, address, staking activity.
- **Chat** — talk to models, manage sessions.
- **Models** — browse the network's registry.
- **Agents** — task-oriented AI beyond simple chat.
- **Provider Hub** — register your own machine to *earn* MOR (a different flow — see
  the [Provider guide](./provider-guide.md)).
- **Settings** — preferences, network config.
- **Help** — docs and support.

---

## Troubleshooting & Common Questions

| Question | Answer |
|---|---|
| My MOR dropped a lot more than the session cost — lost money? | No — see Step 6. Refundable stake, not a payment. |
| My session ended but I still don't see my full MOR back | Expected right now — see Step 8. There's a ~24h hold, then a manual claim is needed. |
| The Models tab "Stake" doesn't match my session's stake | Different things — model listing bond vs. your session stake. |
| Copy icon on a model ID does nothing | Known bug — select and copy the text manually instead. |

---

## Glossary

- **MOR** — token used to access AI models.
- **ETH (Base)** — small gas fee token.
- **Stake** — refundable deposit to open a session — currently subject to a ~24h hold
  and manual claim before it returns (Step 8).
- **Session** — a time-boxed window to chat with a model.
- **Provider** — someone else's computer serving the model.