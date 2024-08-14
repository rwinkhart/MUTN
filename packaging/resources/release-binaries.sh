#!/bin/sh

create_release_binaries() {
    printf '\nGenerating release binaries...\n'
    cd ..
    GOOS=linux CGO_ENABLED=0 GOAMD64=v1 go build -o ./packaging/1output/mutn-linux-x86_64_v1 -ldflags="-s -w" -trimpath ./mutn.go
    GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o ./packaging/1output/mutn-linux-aarch64 -ldflags="-s -w" -trimpath ./mutn.go
    GOOS=windows CGO_ENABLED=0 GOAMD64=v1 go build -o ./packaging/1output/mutn-windows-x86_64_v1.exe -ldflags="-s -w" -trimpath ./mutn.go
    GOOS=windows GOARCH=arm64 CGO_ENABLED=0 go build -o ./packaging/1output/mutn-windows-aarch64.exe -ldflags="-s -w" -trimpath ./mutn.go
    cd ./packaging
}
