![banner](https://raw.githubusercontent.com/rwinkhart/sshyp-labs/main/extra/artwork/MUTN-banner.webp)
---
Pronounced as: "mutton", "muh·tn"

MUTN is a simple, self-hosted, SSH-synchronized password and note manager based on libmutton. It is the successor to [sshyp](https://github.com/rwinkhart/sshyp).

> [!WARNING]
>It is your responsibility to assess the security and stability of MUTN and to ensure it meets your needs before using it.
>I am not responsible for any data loss or breaches of your information resulting from the use of MUTN.
>MUTN is a new project that is constantly being updated, and though safety and security are priorities, they cannot be guaranteed.

# Demo Tape
![mutn-demo.webp](https://raw.githubusercontent.com/rwinkhart/sshyp-labs/main/extra/mutn-vhs/mutn-demo.webp)

# Mission Statement
MUTN aims to make it as easy as possible to manage passwords and notes via CLI across multiple devices in a secure, self-hosted fashion.

# Installation/Building
See the [installation guide](https://github.com/rwinkhart/MUTN/blob/main/wiki/MUTN/install.md).

Additionally, MUTN is available as a source package ("mutn") on the AUR.

After installing, please review the [usage guide](https://github.com/rwinkhart/MUTN/blob/main/wiki/MUTN/usage.md).

# Roadmap
#### Release v0.2.1
- [ ] Only run getSSHClient once to prevent being asked for keyfile password multiple times
    - [ ] After this, handle all errors in client.go
- [ ] Ensure all config files and entry files are created with 0600 permissions
- [ ] Add fail-specific error codes
- [ ] Split into separate repos
    1. libmutton
    2. libmutton-server
    3. MUTN
#### Release v0.3.0
- [ ] Replace Glamour with a more minimal Markdown renderer (likely custom)
#### Release v0.4.0
- [ ] Swap to native (cascade) encryption (custom)
- [ ] Implement "netpin" (quick-unlock) with new encryption
#### Release v0.5.0
- [ ] Password aging support
    - [ ] Append UNIX timestamp to entry names
        - [ ] Add yellow/red dot indicators to entry list readout for when passwords should be changed
#### Release v0.6.0
- [ ] Re-implement init menu
- [ ] Implement tweak menu
- [ ] Add refresh/re-encrypt functionality
#### Release v1.0.0 - Distribution packages
- [ ] Create packaging scripts
- Perform extensive testing
