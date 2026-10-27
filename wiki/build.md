## Building
Official binaries are built without CGO (for portability), built position-independent (for security, allows [ASLR](https://en.wikipedia.org/wiki/Address_space_layout_randomization)), and stripped of debug info (for size reduction), as follows:
```
CGO_ENABLED=0 go build -buildmode=pie -ldflags="-s -w" -trimpath -tags=stxAll ./mutn.go
```

Additionally, some custom build tags can be used to create different binaries. The following tags are not used in official builds:
- `wsl`: Allows creating a Linux binary that can interact with the Windows clipboard (for WSL)
