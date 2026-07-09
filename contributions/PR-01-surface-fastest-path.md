# PR-01 — Surface the fastest onboarding path in the docs

**Target repo:** `MorpheusAIs/Morpheus-Lumerin-Node`
**Target file:** `docs/get-started/introduction.mdx`
**Type:** Documentation (small, single-purpose)
**Related findings:** AUDIT F1, F10 (discoverability), F11 (free local model), F12 (installer vs archive build)

---

## Problem

The Introduction page opens with *what the node is* and an architecture diagram, then offers
role cards. A newcomer in a hurry has to read conceptual material before learning the one thing
they want: **how do I try this right now?** In practice this leads people onto slower, more
technical paths (building/running components manually) when a one-download desktop app and a
**free, no-wallet local model** already exist. Front-loading the fastest path costs nothing and
removes the most common early drop-off.

## Change

Insert a single Mintlify `<Note>` callout **immediately after the frontmatter, before the first
prose paragraph**, matching the page's existing component style. No other content is removed.

```mdx
<Note>
  **In a hurry? Pick the fastest path.**

  - **Just want to use it (~5 min):** download the desktop app from the
    [latest release](https://github.com/MorpheusAIs/Morpheus-Lumerin-Node/releases) and run it —
    wallet, chat, and proxy-router are bundled. See the [Consumer quickstart](/get-started/quickstart-consumer).
  - **Try it free — no wallet, no tokens:** download the **archive (.zip)** build (the one containing
    `mor-launch.exe`) and run `mor-launch local` to start the bundled local model and chat immediately.
  - **Building on the marketplace?** jump to the [Hosted Inference API](/inference-api/overview).

  New to Morpheus? Read on for how the pieces fit together.
</Note>
```

**Optional in same PR:** bump the frontmatter `last_verified: "v7.0.0"` → `"v7.3.0"` (current release).

## PR title

```
docs(get-started): surface the fastest path (desktop app + free local model) at the top of Introduction
```

## PR description (ready to paste)

> **What & why**
> The Introduction page currently leads with architecture before telling a new user how to start.
> This adds a short callout at the very top that points hurried users to the two fastest paths —
> the bundled desktop app, and `mor-launch local` for a free, no-wallet local model — while leaving
> all existing content intact.
>
> **Change**
> - Add a `<Note>` callout after the frontmatter with three "pick your path" links.
> - (Optional) bump `last_verified` to v7.3.0.
>
> **Impact**
> Reduces early drop-off for newcomers and makes the zero-cost local demo discoverable — useful for
> workshops and events where attendees have no wallet or tokens.

## How to submit (you don't need write access — fork + PR)

```bash
# 1. Fork on GitHub, then:
gh repo fork MorpheusAIs/Morpheus-Lumerin-Node --clone
cd Morpheus-Lumerin-Node
git checkout -b docs/surface-fastest-path

# 2. Edit docs/get-started/introduction.mdx — paste the <Note> block after the frontmatter.

# 3. Commit and open the PR
git commit -am "docs(get-started): surface fastest path at top of Introduction"
git push -u origin docs/surface-fastest-path
gh pr create --fill
```

> ⚠️ Confirm the exact surrounding lines against the real file when you clone — placement is
> "after frontmatter, before first paragraph." The block itself is self-contained and style-matched.

## Follow-up (separate PR, not this one)

The GitHub **repo README** gives a stronger "developer / build" first impression than the desktop
app deserves. A second small PR adding a "Just want to use it? Download the app" line near the top
of `README.md` would close the same gap for people who arrive via GitHub rather than the docs site.
Keep it separate so each PR stays single-purpose and easy to merge.
