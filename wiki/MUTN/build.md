## Building
This repository is currently home to both the MUTN client and the general libmutton server software.

Official binaries are stripped of debug info for size and built without CGO (except for distribution packages) for portability, as follows:
```
CGO_ENABLED=0 go build -ldflags="-s -w" -trimpath ./mutn.go
CGO_ENABLED=0 go build -ldflags="-s -w" -trimpath ./libmuttonserver.go
```

Additionally, some custom build tags can be used to create different binaries. The following tags are not used in official builds:
- `wsl`: Allows creating a Linux binary that can interact with the Windows clipboard (for WSL)
- `termux`: Allows creating a Linux binary that can interact with the Termux clipboard (for Android)