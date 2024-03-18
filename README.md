# MUTN Password Manager
Pronounced as: "mutton", "muh·tn"

MUTN is a simple, self-hosted, SSH-synchronized password manager based on libmutton.

MUTN is an expanded re-implementation of [sshyp](https://github.com/rwinkhart/sshyp) written in Go.

Though MUTN will feel very familiar to users of sshyp, it is intended to differ and will break entry compatibility before reaching a stable release.
# WARNING
It is your responsibility to assess the security and stability of MUTN and to ensure it meets your needs before using it.
I am not responsible for any data loss or breaches of your information resulting from the use of MUTN.
MUTN is a new project that is constantly being updated, and though safety and security are priorities, they cannot be guaranteed.

# Mission Statement
MUTN aims to make it as simple as possible to manage passwords and notes via CLI across multiple devices in a secure, self-hosted fashion.

# Building
MUTN is currently in early development and does not yet have online synchronization functionality (offline only).

To build the current version of MUTN, simply clone this repository and run `go build` in its root directory.

# Roadmap
- Refactor for better modularity
  - Release v0.1.0 - No binaries
- Crete server software
- Add native sync
- Add netpin (quick-unlock replacement) functionality
  - Release v0.2.0 - Make repo public - x86_64_v1 binary
- Println text-wrapping
- Markdown support for notes
  - Release v0.2.1 - x86_64_v1 binary
- Swap to native encryption, consider making GPG optional or removing entirely
- Swap to native SSH if not already done
  - Release v0.3.0 - x86_64_v1 binary
- Replace init with interactive ttyPod-based menu
  - Release v0.3.1 - x86_64_v1 binary
- Password aging support
  - Release v0.4.0 - x86_64_v1 binary
- Perform extensive testing, fixing, and optimizing
- Create artwork
- Create manpage
- Create packaging script
  - Release v1.0.0
    - Official packages for:
      - Arch Linux (AUR PKGBUILD, source-based with user's architecture feature level)
      - Debian 12+/Ubuntu 24.04+ Linux
      - Alpine Linux
      - Fedora Linux
      - FreeBSD
      - Windows
      - MacOS (aarch64 only)