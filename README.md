![banner](https://raw.githubusercontent.com/rwinkhart/sshyp-labs/main/extra/artwork/MUTN-banner.webp)
---
Pronounced as: "mutton", "muh·tn"

MUTN is a simple, self-hosted, SSH-synchronized password and note manager based on [libmutton](https://github.com/rwinkhart/libmutton). It is the successor to [sshyp](https://github.com/rwinkhart/sshyp).

> [!WARNING]
>It is your responsibility to assess the security and stability of MUTN and to ensure it meets your needs before using it.
>I am not responsible for any data loss or breaches of your information resulting from the use of MUTN.
>MUTN is a new project that is constantly being updated, and though safety and security are priorities, they cannot be guaranteed.

# Demo Tape
![mutn-demo.webp](https://raw.githubusercontent.com/rwinkhart/sshyp-labs/main/extra/mutn-vhs/mutn-demo.webp)

# Mission Statement
MUTN aims to make it as easy as possible to manage passwords and notes via CLI across multiple devices in a secure, self-hosted fashion.

# Installation/Building
See the [installation guide](https://github.com/rwinkhart/MUTN/blob/main/wiki/install.md).

Additionally, MUTN is available as a source package ("[mutn](https://aur.archlinux.org/packages/mutn)") on the AUR.

After installing, please review the [usage guide](https://github.com/rwinkhart/MUTN/blob/main/wiki/usage.md).

# Roadmap
### Release v0.4.0
- [ ] Switch to fully compliant Markdown (do not preserve new lines)
- [ ] Switch to alternative Markdown renderer (minimark or BEAN)
### Release v0.5.0
- [ ] libmutton v0.5.0
    - [ ] Implement "netpin" (quick-unlock)
    - [ ] Password aging support
        - [ ] Add yellow/red dot indicators to entry list readout for when passwords should be changed
### Release v1.0.0 - Distribution packages
- [ ] Create packaging scripts
    - [x] Stable source PKGBUILD
    - [x] Stable source APKBUILD
    - [ ] Debian/Ubuntu
    - [ ] Fedora
    - [ ] FreeBSD
    - [ ] Windows installer
    - [ ] Brew formula
- Perform extensive testing
