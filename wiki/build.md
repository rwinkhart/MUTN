## Building
Official binaries are stripped of debug info for size and built without CGO (except for distribution packages) for portability, as follows:
```
CGO_ENABLED=0 go build -ldflags="-s -w" -trimpath ./mutn.go
```

Additionally, some custom build tags can be used to create different binaries. The following tags are not used in official builds:
- `BEAN`: Creates a binary with the experimental [BEAN](https://github.com/Trojan2021/BEAN) Markdown renderer in place of Glamour
- `noMarkdown`: Creates a binary without Markdown support (much smaller binary size and faster launch time VS Glamour)
- `wsl`: Allows creating a Linux binary that can interact with the Windows clipboard (for WSL)
- `termux`: Allows creating a Linux binary that can interact with the Termux clipboard (for Android)
