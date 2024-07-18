## Developer Guide (for third-party libmutton-based clients)
**Important Notice**: libmutton is in early development and is currently a moving target to develop off of. Feel free to jump in early, but greater change stability will be met with release v1.0.0. Check [here](https://github.com/rwinkhart/MUTN/blob/main/wiki/libmutton/breaking.md) for planned breaking changes.

libmutton was designed to be usable as a library for building other compatible password managers off of. [MUTN](https://github.com/rwinkhart/MUTN) is the official reference CLI password manager, however libmutton can be implemented in many other unique ways.

All functionality in the `backend` and `sync` packages are designed to be used in other implementations.

If any functionality in these two packages proves to be difficult to implement in a third-party client, please open an issue so that it can be addressed.

## Build Tags
Custom build tags can (and sometimes must) be used to achieve desired results.

These are as follows:
- `returnOnExit`: If making an interactive interface (GUI/TUI/interactive CLI), you probably need to use this build tag. Without it, your entire program will exit after any given operation is completed. This behavior is only desired for non-interactive CLI implementations, such as MUTN. Currently, errors will result in the program exiting **even with this build tag**. This may be changed in the future (under evaluation).
- `wsl`: Allows creating a Linux binary that can interact with the Windows clipboard (for WSL)
- `termux`: Allows creating a Linux binary that can interact with the Termux clipboard (for Android)

## Required Arguments
libmutton-based password manager clients should accept at least one specific required argument, as well as another recommended one:
- `clipclear`: This argument is required for correct functionality. In order to clear the clipboard on a timer, libmutton-based password managers call another instance of their executable with the `clipclear` argument (e.g. `mutn clipclear`) with the intended clipboard contents provided via STDIN. If after 30 seconds the clipboard contents have not changed, they are cleared. Please accept a `clipclear` argument that is processed before the launch of any interactive interface. All this argument needs to do is call `backend.ClipClearArgument()`.
- `init`: This argument is optional, but recommended for CLI interfaces. Some error messages request the user to use the `init` argument to fix configuration issues. If the argument does not exist, this may confuse the user.

## Configuration
libmutton-based password manager clients should all share the same INI configuration file.

On UNIX-like systems, this is located at `~/.config/libmutton/libmutton.ini`. On Windows, it is located at `~\AppData\Local\libmutton\config\libmutton.ini`.

### Base `libmutton.ini` Layout
The current base layout of `libmutton.ini` will change leading up to release v1.0.0. As of right now, the specification is as follows:
```
[LIBMUTTON]
gpgID = <gpg key id>
sshUser = <remote user>
sshIP = <remote ip>
sshPort = <remote ssh port>
sshKey = <ssh private key identity file path>
sshKeyProtected = <true/false>
sshEntryRoot = <remote entry root>
sshIsWindows = <true/false>
```
If creating a third-party client that requires extra configuration to be stored, please use the same file and create a new INI section for your application-specific configuration, e.g.:
```
[THIRD-PARTY-CLIENT-NAME]
configKey = <value>
```
This ensures that a user can use multiple client applications with the same configuration while avoiding conflicts.

# Relevant Bugs Affecting Third-Party Implementations
- Password-protected SSH identity files currently only prompt for password entry in the CLI, and thus they are not yet supported in GUI/TUI implementations
