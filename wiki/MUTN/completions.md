## Shell Completions Troubleshooting
[Shell completions are provided for ZSH, Bash, and PowerShell 7+](https://github.com/rwinkhart/MUTN/tree/main/completions).
### ZSH completions not working?
Make sure your ~/.zshrc contains the following (if sourcing from default completions location, usually `/usr/share/zsh/functions/Completion/Unix/_mutn`):
```shell
autoload -Uz compinit && compinit
```
...and then restart your shell.
***
### Bash completions not working? 
Install your distribution's 'bash-completion' package or source the completion script manually.

For most environments, manual sourcing means adding the following to your ~/.bashrc:
```shell
source /usr/share/bash-completion/completions/mutn
```
Note that this directory is different on FreeBSD.

FreeBSD:
```shell
source /usr/local/share/bash-completion/completions/mutn
```
...and then restart your shell.

*Please note that Bash completions are slightly more limited than ZSH completions, and as such, new entries will not be auto-completed until the completions script is re-sourced.*
***
### PowerShell completions not working?
Ensure you are using modern PowerShell (7+) as opposed to the built-in Windows PowerShell.

Currently, completions must be manually installed and sourced in your PowerShell profile as such:
```shell
Set-PSReadlineKeyHandler -Key Tab -Function Complete  # this line is optional; it makes tab completion function more similarly to Bash/ZSH
. <drive letter>:\path\to\mutn.ps1
```
...and then restart your shell.