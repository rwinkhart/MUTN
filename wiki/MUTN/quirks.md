## Windows Quirks
MUTN for Windows is fully functional, though it does exhibit a few quirks/bugs not found on other platforms. Watch out for these, and if you have the know-how, pull requests addressing them are welcomed!

- Reading entries (especially those containing Markdown notes) often results in failure to interpret ANSI escape codes, leading to the raw escape codes being dumped to stdout. This happens seemingly randomly and re-running the exact same command ALWAYS yields the correct output. It's like the terminal has to "warm up" to ASNI escape codes or something. **Help wanted**.
- GPG is sometimes (seems unpredictable) incredibly slow to start on Windows (often after a reboot), leading to many operations seemingly hanging
  - **This will be addressed** in the migration off of GPG that will take place before v1.0.0
- PowerShell tab completions will not complete after a directory name if the directory contains a space