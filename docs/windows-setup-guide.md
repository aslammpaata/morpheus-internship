# Getting Started on Windows (Draft)

> Draft — sections marked _TODO_ are pending audit answers. Target format: Mintlify (nodedocs.mor.org).

Morpheus lets you run and use AI models over a decentralized network. This guide gets you from
zero to your first chat on **Windows**. Pick the path that fits you:

- **Just want to use it?** → [Path A: Desktop App](#path-a-desktop-app-recommended) (5 minutes)
- **A developer who wants control?** → [Path B: Router + CLI](#path-b-router--cli-advanced)

---

## Path A: Desktop App (recommended)

The desktop app bundles everything — wallet, chat, and the proxy router — in one download.

1. **Download** `win-x64-morpheus-app-<version>.exe` from the
   [Releases page](https://github.com/MorpheusAIs/Morpheus-Lumerin-Node/releases).
2. **Run it.** _TODO: SmartScreen / unsigned-binary warning steps + screenshot._
3. **Set up your wallet.** _TODO: does the app create one, or import via MetaMask? screenshot._
4. **Fund your wallet** — you need a little **ETH on Base** (for gas) and **MOR** (for sessions).
   _TODO: where to get each; testnet vs mainnet note._
5. **Open a chat**, pick a model, and send your first message. _TODO: screenshot._

✅ That's it — you're using decentralized AI.

---

## Path B: Router + CLI (advanced)

More control over session management, at the cost of more steps.

> ⚠️ **Windows note:** the current router+CLI flow uses Linux-style commands
> (`systemctl`, shell scripts). _TODO (AUDIT F7): document the WSL path or the native
> `win-x64-morpheus-router.exe` equivalent._

> 🔐 **Security:** your router API password is a **secret**. Don't paste it into shared docs,
> screenshots, or issues. Store it in an environment variable.

1. **Start the router.** _TODO: native Windows / WSL command._
2. **Set up a wallet.** _TODO (AUDIT F8): canonical wallet command (not the `everclaw` script)._
3. **Open a session** (stakes MOR, returns a session ID).
4. **Run inference** with that session ID.

> 💡 **Tip:** [`morpheus-doctor --dev`](../morpheus-doctor/SPEC.md) does steps 1, 3, and 4 for you
> in a single command — no copy-pasting session IDs.

---

## Troubleshooting

_TODO: build from the audit's friction log (ports, funding, wrong network, auth errors)._
