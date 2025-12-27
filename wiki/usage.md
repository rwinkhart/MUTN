## Setup/Usage Tips - MUTN CLI
Most information relevant to the usage of MUTN can be viewed with `mutn help`, in [the man page](https://raw.githubusercontent.com/rwinkhart/MUTN/main/docs/man), and in the demo animation from the README.

Because of this, this wiki page will focus on initial setup as well as some tips that may not be immediately obvious.

### Initial Setup
libmutton operates on a client-server model, and thus sync support requires you to have access to your own server (whether it be a physical home server or a cloud rental) with remote access via SSH.

For security reasons, libmutton deliberately does not support password-based SSH authentication; you must use key-based authentication to connect to your server (password-protected keys are also supported). You are responsible for setting up SSH connectivity on your own.

The libmutton server software is shipped as its own dedicated binary (in [libmutton releases](https://github.com/rwinkhart/libmutton/releases)) and must be installed on the server and configured with `libmuttonserver init`. The server binary _**MUST**_ be in your $PATH and _**MUST**_ be named "libmuttonserver" for clients to successfully reach it.

Once this has been done, you may proceed with the client setup:
1. Install MUTN
2. Run `mutn init` and follow the prompts
3. If you already have entries on the server, sync them using `mutn sync`
4. Optionally, shell completions (Bash, ZSH, and PowerShell 7+) can be [enabled with your shell's respective method](https://github.com/rwinkhart/MUTN/blob/main/wiki/completions.md)
5. Optionally, [import passwords](https://github.com/rwinkhart/libmutton/blob/main/wiki/migration.md) from another password manager. After doing this, you may desire to use the "Age all entries" option from the `mutn tweak` menu to add initial age tracking data to the imported passwords

### Tips
#### Operating on Passwords
Any action that operates on the password attached to an entry (copying, editing, adding a new password-focused entry) does not require the user to specify the `password` or `-pw` options. In all cases, `password` is the assumed target when no other option is specified.

#### TOTP Support
See [here](https://github.com/rwinkhart/libmutton/blob/main/wiki/tips.md#totp-support).

#### Granular Help
Argument-specific help can be obtained by running `mutn <argument>` without an entry name or any options. For example, `mutn edit` will display help for the `edit` argument.
