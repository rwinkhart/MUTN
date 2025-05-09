## Installation
MUTN is still in early development and thus no installation packages are officially distributed.

For now, please manually install the binaries from the [latest release](https://github.com/rwinkhart/MUTN/releases) or [compile from source](https://github.com/rwinkhart/MUTN/blob/main/wiki/build.md). All pre-built binaries were compiled with CGO disabled and thus are libc-independent.

Additionally, MUTN is available as a source package ("[mutn](https://aur.archlinux.org/packages/mutn)") on the AUR.

Please note that the server binary (available in [libmutton releases](https://github.com/rwinkhart/libmutton/releases)) must be called "libmuttonserver" in order for clients to successfully reach it.

After placing the binaries in your $PATH, it is highly recommended to also download and correctly place/source the relevant [shell completions scripts](https://github.com/rwinkhart/MUTN/blob/main/wiki/completions.md).

Please see the [usage guide](https://github.com/rwinkhart/MUTN/blob/main/wiki/usage.md) for help getting started.

### Dependencies (required for all installations)
- A text editor (preferably CLI-based) for writing entry notes
- A private key for SSH key-based authentication is required to use MUTN/libmutton in online/synced mode
