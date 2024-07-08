## Migrating From Other Password Managers
**Important Notice**: The libmutton entry format is not final and has two [breaking changes planned prior to release v1.0.0](https://github.com/rwinkhart/MUTN/blob/main/wiki/libmutton/breaking.md). This guide will be updated accordingly.
### pass
libmutton-based password managers *currently* use GnuPG encryption and an entry format similar to that of [pass](https://www.passwordstore.org/). Because of this, any entries in `pass` format can simply be dropped into `~/.local/share/libmutton`.
### sshyp
sshyp, though also `pass`-compatible, makes some changes to the entry format that take effect once the entry has been imported. The changes made by sshyp are not compatible with libmutton, and as such sshyp entries must be converted before they can be used. A script for doing that has been created and is available in the sshyp extension store. Simply run `sshyp tweak`, go to the "extension management" menu, and download the "export-to-libmutton" extension. After doing this, the `sshyp export` command can be used to export entries in libmutton format.
### Other
The formats for many other password managers can be converted to the `pass` format with community scripts. Some of these scripts are listed [here](https://www.passwordstore.org/#migration). Once converted, entries can be dropped into `~/.local/share/libmutton`.