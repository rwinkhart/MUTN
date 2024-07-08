# MUTN Password Manager
Pronounced as: "mutton", "muh·tn"

MUTN is a simple, self-hosted, SSH-synchronized password manager based on libmutton. It is the successor to [sshyp](https://github.com/rwinkhart/sshyp).

> [!WARNING]
>It is your responsibility to assess the security and stability of MUTN and to ensure it meets your needs before using it.
>I am not responsible for any data loss or breaches of your information resulting from the use of MUTN.
>MUTN is a new project that is constantly being updated, and though safety and security are priorities, they cannot be guaranteed.

# Mission Statement
MUTN aims to make it as simple as possible to manage passwords and notes via CLI across multiple devices in a secure, self-hosted fashion.

# Building
See the [building guide](https://github.com/rwinkhart/MUTN/blob/main/wiki/MUTN/build.md).

# Roadmap
#### Release v0.2.0 - Make repo public - No binaries
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