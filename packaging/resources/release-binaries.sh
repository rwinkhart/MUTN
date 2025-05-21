#!/bin/sh
# This script generates portable MUTN release binaries for the following platforms:
# - Linux (x86_64_v2)
# - Linux (arm64v8.0)
# - Linux (arm64v8.7+crypto)
# - Windows (x86_64_v2)
# - Windows (arm64v8.7+crypto)

create_release_binaries() {
    printf '\nGenerating release binaries...\n'
    cd ..
    GOOS=linux CGO_ENABLED=0 GOAMD64=v2 go build -o ./packaging/1output/mutn-linux-x86_64_v2 -ldflags="-s -w" -trimpath ./mutn.go
    GOOS=linux GOARCH=arm64 GOARM64=v8.0 CGO_ENABLED=0 go build -o ./packaging/1output/mutn-linux-arm64v8.0 -ldflags="-s -w" -trimpath ./mutn.go
    GOOS=linux GOARCH=arm64 GOARM64=v8.7,crypto CGO_ENABLED=0 go build -o ./packaging/1output/mutn-linux-arm64v8.7+crypto -ldflags="-s -w" -trimpath ./mutn.go
    GOOS=windows CGO_ENABLED=0 GOAMD64=v2 go build -o ./packaging/1output/mutn-windows-x86_64_v2.exe -ldflags="-s -w" -trimpath ./mutn.go
    GOOS=windows GOARCH=arm64 GOARM64=v8.7,crypto CGO_ENABLED=0 go build -o ./packaging/1output/mutn-windows-arm64v8.7+crypto.exe -ldflags="-s -w" -trimpath ./mutn.go
    cd ./packaging
}
