# MUTN Password Manager
Pronounced as: "mutton", "muh·tn"

MUTN is a simple, self-hosted, SSH-synchronized password manager based on libmutton.

MUTN is an expanded re-implementation of [sshyp](https://github.com/rwinkhart/sshyp) written in Go.

Though MUTN will feel very familiar to users of sshyp, it is intended to differ and breaks compatibility with its entry format.

> [!WARNING]
>It is your responsibility to assess the security and stability of MUTN and to ensure it meets your needs before using it.
>I am not responsible for any data loss or breaches of your information resulting from the use of MUTN.
>MUTN is a new project that is constantly being updated, and though safety and security are priorities, they cannot be guaranteed.

# Mission Statement
MUTN aims to make it as simple as possible to manage passwords and notes via CLI across multiple devices in a secure, self-hosted fashion.

# Building
This repository is currently home to both the MUTN client and the general libmutton server software.

Official binaries are stripped of debug info for size and built without CGO (except for distribution packages) for portability, as follows:
```
CGO_ENABLED=0 go build -ldflags="-s -w" ./mutn.go
CGO_ENABLED=0 go build -ldflags="-s -w" ./libmuttonserver.go
```

# Roadmap
#### Release v0.2.0 - Make repo public - No binaries
- Add custom build option for WSL support
- Add fail-specific error codes
- Address most TODOs
- Hunt and fix bugs in preparation for public debut
#### Release v0.2.1 - No binaries
- Replace init with interactive ttyPod-based menu
#### Release v0.3.0 - x86_64_v1 binary
- Swap to native encryption, remove GPG support
#### Release v0.4.0 - x86_64_v1 binary
- Password aging support
  - Add password aging info to entry names
    - Make constant character count to easily trim for user interaction
#### Release v1.0.0 - Distribution packages (from here on out)
- Perform extensive testing, fixing, and optimizing
- Create artwork
- Create manpage
- Create packaging script
#### Release v1.1.0
- Cascading encryption support
#### Release v1.2.0
- Custom color scheme support
#### Release v1.3.0
- Add build option for breaking markdown support and menus into separate binaries
  - Greatly reduces startup time