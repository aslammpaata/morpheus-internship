# Morpheus Onboarding System

Tooling, documentation, and workshop material that shrink the distance between a newcomer and
their first working inference on the [Morpheus](https://mor.org) / Lumerin decentralized AI network.

Built during an internship at MorpheusAI, ahead of a Web3 booth event (Aug 2026). Designed to keep
providing value — and to be maintained by the team — long after the internship ends.

## Why this exists

The [Morpheus-Lumerin-Node](https://github.com/MorpheusAIs/Morpheus-Lumerin-Node) already ships a
solid, cross-platform Proxy Router and a recommended desktop app. The remaining barrier to adoption
is **onboarding friction** — discovering the easy path, funding a wallet, and (for developers) the
painful manual `router → session → inference` chain. This repo owns that gap.

## Components

| Path | What it is | Status |
|------|-----------|--------|
| [`AUDIT.md`](./AUDIT.md) | Living Windows onboarding friction log — source of truth for everything else | 🟡 In progress |
| [`docs/windows-setup-guide.md`](./docs/windows-setup-guide.md) | Beginner-friendly, Windows-first getting-started guide (PR target: Mintlify docs) | ⬜ Draft |
| [`morpheus-doctor/`](./morpheus-doctor/) | Single-binary Go tool: preflight + one-command `router → session → inference` for the developer path | ⬜ Spec'd |
| Workshop + booth kit | Tiered live-demo script + short video (built on the tool) | ⬜ Planned |

## Design principles

- **Don't rebuild what exists.** Automate and document around the official node.
- **Two audiences, two tracks.** Beginners → desktop app (docs + workshop). Developers → CLI (the tool).
- **Docs are a deliverable, not an afterthought.** Every finding becomes a doc, an issue, or a feature.
- **Upstreamable.** Prefer small, high-quality PRs to the official repo over a fork nobody merges.

## Flagship deliverable

`morpheus-doctor` — a single static Windows `.exe` that replaces the manual developer chain
(start router, open session, copy-paste session ID, run inference) with **one command**.
See [`morpheus-doctor/SPEC.md`](./morpheus-doctor/SPEC.md).
