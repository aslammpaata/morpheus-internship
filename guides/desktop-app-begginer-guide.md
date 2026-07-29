---
title: "Morpheus Desktop App — A Beginner's Walkthrough"
description: "A hands-on, screenshot-driven guide from download to your first AI reply — no crypto experience required."
audience: ["consumer", "beginner"]
last_verified: "v7.3.0"
---

# Morpheus Desktop App — A Beginner's Walkthrough

This guide walks you through the Morpheus Desktop App from a blank computer to your first
AI reply, one screen at a time. It assumes no prior blockchain or crypto experience. If you
want the conceptual deep-dive instead, see the [official docs](https://nodedocs.mor.org) —
this guide is the hands-on companion to those.

**What you'll do:** Download → Install → Launch → Set up a wallet → Add a little funding →
Pick a model → Open a session → Chat → See your funds return.

**Time:** about 10 minutes, plus however long you spend chatting.

---

## Before You Start

**What Morpheus is, in one sentence:** it's a way to chat with AI models that run on other
people's computers around the world, instead of one company's servers — coordinated by a
public ledger (a blockchain called Base) instead of a central authority.

**What you'll need:**
- Windows 10 or 11 *(confirm minimum supported version before publishing)*.
- An internet connection, if you want to use network models. *(You can skip straight to a
  free, offline local model with no internet and no funding at all — see the tip at the end
  of Step 4.)*
- About 10 minutes and, if you want to try a real network session, a small amount of two
  things: **MOR** and **ETH**. Both are explained in Step 4 — you don't need to know anything
  about them yet.

> 🔐 **One thing to know before you begin:** this app creates a real crypto wallet, which can
> hold real value. You'll get a clear warning at the exact moment this matters (Step 3), but
> it's worth knowing upfront that this isn't just a chat app — treat the wallet part with the
> same care you'd give any financial account.

---

## Step 1: Download & Install

1. Go to the [Releases page](https://github.com/MorpheusAIs/Morpheus-Lumerin-Node/releases).
2. Download the file named **`win-x64-morpheus-app-<version>.exe`**.

   **📸 Screenshot 1 — Releases page.** Highlight the correct installer filename with a box.
   This page lists a few different files (installer, router, CLI) — draw arrows to the other
   two marked "not this one," since picking the wrong file is the single most common
   first mistake.

3. Run the installer. Windows will likely show a blue **"Windows protected your PC"** screen —
   this is normal for a new, unsigned application, not a sign of malware.

   **📸 Screenshot 2 — SmartScreen warning.** Arrow pointing to **"More info"**, then a second
   arrow to the **"Run anyway"** button that appears after clicking it.

4. Follow the installer prompts (default options are fine).

✅ Morpheus is now installed.

---

## Step 2: First Launch

1. Open **Morpheus** from your Start menu or desktop shortcut.
2. The app starts its background service automatically — you don't need to do anything else
   to "turn it on."

   **📸 Screenshot 3 — First screen you see on launch**, whatever that is before wallet setup
   begins (a splash screen, or straight into Step 3 below). No annotation needed — this is
   just an orientation shot so the reader recognizes they're in the right place.

---

## Step 3: Set Up Your Wallet

A wallet is what holds your funds and identifies you to the network. You'll set one up once.

1. **Set a password** for the app. This protects your wallet on this device.

   **📸 Screenshot 4 — Password setup screen.**

2. Choose **Create a new wallet** (if this is your first time) or **Recover** (if you already
   have a wallet from somewhere else, using your seed phrase).

   **📸 Screenshot 5 — Create vs. Recover screen.** Callout box: *"New here? Choose Create.
   Already have a seed phrase? Choose Recover."*

3. If you chose Create, you'll be shown a **seed phrase** — a list of words that is the master
   key to your wallet.

   **📸 Screenshot 6 — Seed phrase screen.** This is the most important annotation in the
   whole guide. Use a bold red callout box directly over the phrase:
   > 🔴 **Never share this with anyone. Never type it into a website. Never screenshot and
   > post it. Write it down on paper and store it somewhere safe.** Anyone with these words
   > has full control of your wallet — including Morpheus staff, who will never ask for it.

4. Confirm your seed phrase if prompted, and finish setup.

✅ Your wallet exists. Right now it's empty — that's expected, and totally fine. Next, we'll
check that, and add a little funding.

---

## Step 4: Fund Your Wallet

Two things live in your wallet, and they do different jobs:

| | What it's for | Think of it as |
|---|---|---|
| **ETH** (on the Base network) | Small transaction fees ("gas") | The toll you pay to use the road |
| **MOR** | What you actually use to access AI models | The thing you're really here for |

You'll need a small amount of both — ETH just needs to be enough for gas (a few cents' worth
goes a long way), MOR needs to be at least **5** to open a session.

1. Go to the **Wallet** tab. You'll see `0 MOR` and `0 ETH` — expected for a brand-new wallet.

   **📸 Screenshot 7 — Wallet page, unfunded.** Arrow to the wallet address and its copy
   button, arrow to the **Receive** button.

2. Send MOR and ETH to this address from an exchange or another wallet you control, or use
   the **Receive** button to get a QR code / shareable address.
3. Once funds arrive, the Wallet page updates automatically.

   **📸 Screenshot 8 — Wallet page, funded.** Place directly next to Screenshot 7 so the
   before/after is visually obvious.

✅ Your wallet is funded and ready.

> 💡 **Want to try Morpheus for free first, with no funding at all?** Skip ahead — in the
> **Chat** tab, you can select a bundled **local model** that runs entirely on your own
> computer, no wallet balance required. Come back to Step 5 when you're ready for the full
> network.

---

## Step 5: Choose a Model

1. Go to the **Models** tab.

   **📸 Screenshot 9 — Models → Registry view.** Circle 2-3 different model cards. Callout:
   *"The Fee and Stake numbers on this card belong to the model's listing — they are not
   what you'll pay to use it. Your actual cost is set per-session in the next step."* This is
   a common point of confusion worth heading off here, not later.

2. Note a model's name — you'll pick it from the Chat tab next.

---

## Step 6: Open a Session (Staking, Explained Simply)

This is the step that trips people up the most, so read this part before you click anything.

**What a session is:** a time-boxed window during which you can chat with a specific model.
**What staking means:** to open a session, you temporarily lock some MOR as a deposit. This
deposit is **not spent** — it comes back to your wallet automatically once the session ends.
The *actual* cost of your session is a much smaller number, paid separately from network
rewards, not from your stake.

> 💡 Think of it like a hotel deposit: you put down more than the room actually costs, and
> get it back when you check out — as long as nothing goes wrong.

1. Go to **Chat → Change Model**, and pick the model you noted in Step 5.

   **📸 Screenshot 10 — Model picker.**

2. Click **Open Session**. You may see the amount about to be staked before confirming.

   **📸 Screenshot 11 — Stake confirmation.** Annotate: *"This number is your refundable
   deposit, not a bill."*

3. Confirm. Your session opens.

4. Check the **Wallet** tab — your available MOR will have dropped by the staked amount, and
   your **Staked Balance** will show the same amount.

   **📸 Screenshot 12 — Wallet page immediately after staking.** This needs the strongest
   callout in this section:
   > 🔴 **This is expected!** Your MOR isn't gone — it's temporarily locked, not spent. It
   > returns automatically when your session ends naturally (see Step 8).

5. At the top of the Chat screen, you'll see your session's ID, a countdown timer, and a small
   cost figure.

   **📸 Screenshot 13 — Session header.** Arrow to the timer, and a separate arrow to the
   small cost number, with a callout distinguishing it from the (much larger) stake shown in
   Screenshot 12.

✅ Your session is open and counting down.

---

## Step 7: Chat & Run Inference

1. Type a message in the chat box and send it.

   **📸 Screenshot 14 — A chat conversation with a model reply visible.**

2. A reply appears from the model — that's "inference," working exactly like any AI chat
   you've used before. Under the hood, that reply came from someone else's computer on the
   network, verified and paid for through the session you just opened.

✅ You've used decentralized AI.

---

## Step 8: What Happens When You're Done

You don't need to manually "close" your session — in fact, for a first session, it's easiest
not to.

- **Let it run out.** When the timer in Screenshot 13 reaches zero, your session ends on its
  own, and your staked MOR is automatically returned to your wallet within about a minute.
- **If you close it early**, part of your stake goes into a short hold period (about a day)
  before you can access it again — not lost, just delayed. There's rarely a good reason to do
  this as a beginner, so we won't cover it further here.

**📸 Screenshot 15 — Wallet page after the session has expired**, showing the balance
restored. Place next to Screenshot 12 to visually close the loop: locked → returned.

✅ Full loop complete: you funded a wallet, opened a session, chatted with a decentralized AI
model, and got your deposit back automatically.

---

## Understanding the App: A Tour of Each Section

Now that you've been through the core flow once, here's what the rest of the sidebar is for.

**📸 Screenshots 16a–16g — one small screenshot per tab below**, simple numbered callout,
no heavy annotation needed.

- **Wallet** — your balances, address, and staking activity. You'll return here often to
  check funds.
- **Chat** — where you talk to models and manage sessions, as you just did.
- **Models** — browse everything available on the network, including fees and tags for each.
- **Agents** — for interacting with more complex, task-oriented AI agents registered on the
  network (beyond simple chat). Worth exploring once you're comfortable with the basics.
- **Provider Hub** — for the *other* side of the marketplace: registering your own computer
  to serve models and earn MOR. Not something you need for using Morpheus, only for hosting.
- **Settings** — app preferences, network configuration, and advanced options.
- **Help** — documentation links and support if you get stuck.

---

## Troubleshooting & Common Questions

| Question / symptom | Answer |
|---|---|
| My MOR balance dropped a lot more than my session cost — did I lose money? | No — see Step 6. That's your refundable stake, not a payment. It returns automatically. |
| The "Stake" number on a model's card doesn't match what I paid to open a session | These are two different things. The model card's stake belongs to whoever listed the model; your session stake is calculated separately when you open a chat. |
| I clicked the copy icon next to a model ID and it didn't paste anywhere | Known display bug — the copy sometimes silently fails. Try selecting and copying the text manually instead. |
| Do I need to close my session manually? | No — just let it run out. See Step 8. |
| Windows blocked the installer | Normal for a new app — see Step 1, "More info → Run anyway." |

---

## Glossary

- **MOR** — the token used to access AI models on Morpheus.
- **ETH (on Base)** — a small transaction fee token, needed in tiny amounts alongside MOR.
- **Stake** — a refundable deposit locked to open a session; not a payment.
- **Session** — a time-boxed window during which you can chat with a chosen model.
- **Provider** — someone else's computer, registered on the network, that actually runs the
  AI model and answers your prompts.
- **Bid** — the price a provider has set to serve a particular model; determines how much
  session time your stake buys.