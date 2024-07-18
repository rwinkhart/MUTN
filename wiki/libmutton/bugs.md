# Known Bugs - libmutton
- On Windows, GPG is sometimes (seems unpredictable) incredibly slow to start (often after a reboot), leading to many operations seemingly hanging
    - **This will be addressed** in the migration off of GPG that will take place before v1.0.0
- If a client is reconfigured to use a new device ID, the old one will remain present on the server. This will not cause any immediate issues, however it will lead to the server creating unnecessary deletions files for the old device ID. Currently, it is recommended to manually delete the old device ID from the server's "devices" directory.
    - **This will be addressed** in v0.2.1
- Using a password-protected SSH identity file will prompt for the key's password multiple times when syncing
    - **This will be addressed** in v0.2.1
