# Contribution Status

Tracks every PR/issue filed against `MorpheusAIs/Morpheus-Lumerin-Node`, and
whether closure actually means resolved. GitHub's "Closed" label doesn't
always mean fixed — see #792/#794 below.

**Last updated:** 2026-07-22

## Pull Requests

| # | Title | Status | Notes |
|---|---|---|---|
| [#791](https://github.com/MorpheusAIs/Morpheus-Lumerin-Node/pull/791) | Surface fastest path at top of Introduction | ✅ Merged (`dev` branch) | Live-site deploy status unconfirmed — deploy check did not complete on merge; flagged to team |

## Issues

| # | Title | Filed | Status | Actually resolved? |
|---|---|---|---|---|
| [#792](https://github.com/MorpheusAIs/Morpheus-Lumerin-Node/issues/792) | mor-cli Windows cookie-path bug | ✅ | Closed via PR #807 | **Partial.** Core path-duplication bug fixed. Reopened — 3 follow-up findings from later testing (unbatched bid loading, 3600s default duration causing balance failures, chat TUI not sending messages) were not addressed |
| [#793](https://github.com/MorpheusAIs/Morpheus-Lumerin-Node/issues/793) | Model card copy-to-clipboard silently fails | ✅ | Open | Not yet triaged |
| [#794](https://github.com/MorpheusAIs/Morpheus-Lumerin-Node/issues/794) | Docs: chat context not retained by default | ✅ | Closed, no linked fix | Clarified, not a bug. Maintainer confirmed this is intended: the router is stateless by design, and client-side history management is the expected pattern — which morctl --interactive already implements correctly. Doc clarification on /v1/chat/completions still arguably worth adding for future users hitting the same confusion (Asia team independently did), but the underlying behavior is correct as-is. |
| [#795](https://github.com/MorpheusAIs/Morpheus-Lumerin-Node/issues/795) | Docs: fastest path not surfaced | ✅ | — | Check linked PR status before relying on this row |
| [#796](https://github.com/MorpheusAIs/Morpheus-Lumerin-Node/issues/796) | Docs reference nonexistent `.zip`/`mor-launch.exe` | ✅ | Closed via PR #814/#815 | **Resolved.** Real docs fix, stale references removed |
| [#811](https://github.com/MorpheusAIs/Morpheus-Lumerin-Node/issues/811) | Onboarding completely broken on fresh install | ✅ | Closed via PR #831 | **Resolved** — confirmed on multiple machines, multiple versions. |

## Open follow-ups (not yet filed as separate issues)

- `mor-cli chat`: unbatched/serial bid-loading (~8-9 min load, no progress indicator) — currently a comment on #792
- `mor-cli chat`: unprompted 3600s default session duration, causing opaque insufficient-balance failures — currently a comment on #792
- `mor-cli chat`: interactive TUI fails to send messages even with a valid session (verified via parallel curl test) — currently a comment on #792

## Booth risk summary

| Path | Status | Booth-usable? |
|---|---|---|
| Desktop App onboarding | Broken (#811) | ✅ Working |
| `mor-cli` | 4 confirmed issues | ❌ No |
| Headless router + curl | Working, tested | ✅ Yes |
| `morctl` | Working, tested | ✅ Yes |

**Contingency:** if #811 is unresolved by ~Jul 24-25, default to the headless
router + `morctl` path for the booth's dev-track demo, and consider
a pre-provisioned/pre-onboarded wallet workaround for the beginner track
(see workshop/booth-runbook.md).