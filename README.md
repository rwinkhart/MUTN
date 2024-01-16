# MUTN Password Manager
Pronounced as: "mutton", "muh·tn"

MUTN is a very simple self-hosted, SSH-synchronized password manager based on libmutton.

MUTN is an expanded re-implementation of [sshyp](https://github.com/rwinkhart/sshyp) written in Go.

Though MUTN will feel very familiar to users of sshyp, it is intended to differ and eventually break compatibility.

# WARNING
It is your responsibility to assess the security and stability of MUTN before using it and ensure it meets your needs.
I am not responsible for any data loss or breaches of your information resulting from the use of MUTN.
MUTN is a new project that is constantly being updated, and though safety and security are priorities, they cannot be guaranteed.

# Mission Statement
MUTN aims to make it as simple as possible to manage passwords and notes via CLI across multiple devices in a secure, self-hosted fashion.

# Building
MUTN is currently in early development and does not yet serve its intended purposes.

To build a test version of MUTN, simply clone this repository and run `go build main` in the "src" directory.

# Roadmap
Short-term Goals:

- Re-implement all of sshyp's functionality, including the sshyp-mfa extension from [sshyp-labs](https://github.com/rwinkhart/sshyp-labs)

Long-term Goals:

- Migrate away from GPG
- Consider using a Go-based SSH package, rather than relying on OpenSSH directly
- Break the entry format to allow for storing password aging data (how old a password is)
- Add markdown support for notes
