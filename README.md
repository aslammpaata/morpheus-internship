# Morpheus Internship — Onboarding & Contributions

Project work for the Morpheus onboarding-friction audit, upstream
contributions, and Web3 event (Jul 30–31) prep.

The developer tool that came out of this work, `morctl`, now lives
in its own repo: **[morctl](https://github.com/aslammpaata/morctl)**.

## Structure

```
contributions/     Friction audit + tracked status of every upstream PR/issue
guides/            Step-by-step setup guides (Windows installer, headless/dev path)
workshop/          Web3 event booth materials (runbook, handout, demo script)
```

## Contributions

- **[AUDIT.md](contributions/AUDIT.md)** — full friction log from a firsthand
  zero-to-first-inference run, findings F1–F16+, cross-referenced against
  official docs.
- **[STATUS.md](contributions/STATUS.md)** — every PR/issue filed upstream,
  and whether "Closed" on GitHub actually means resolved.

## Guides

- **[windows-setup-guide.md](guides/windows-setup-guide.md)** — installer
  path, free local model, developer/API path, troubleshooting.
- **developer-headless-guide.md** — *(in progress)* WSL2/headless
  proxy-router setup, wallet injection, session + inference via `curl` and
  via `morctl`.

## Workshop

Booth materials for the Web3 event, kept current against what's actually
been tested rather than the original plan — see
[STATUS.md's booth risk summary](contributions/STATUS.md#booth-risk-summary)
for the current go/no-go state of each onboarding path.

## Key findings so far

- The Desktop App's Windows onboarding is currently broken on a clean
  install, confirmed on multiple machines/versions ([#811](https://github.com/MorpheusAIs/Morpheus-Lumerin-Node/issues/811), critical).
- `mor-cli` has multiple confirmed issues on Windows and in its interactive
  chat flow ([#792](https://github.com/MorpheusAIs/Morpheus-Lumerin-Node/issues/792)).
- The headless router + `morctl` path is fully tested and working —
  current fallback for both dev demos and, if needed, the booth itself.
