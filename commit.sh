#!/bin/sh
gofmt -l -w -s ./src/cli/*.go
gofmt -l -w -s ./src/offline/*.go
git add -f extra glamour-styles src .gitignore commit.sh go.mod go.sum LICENSE main.go README.md
git commit -m "$1"
git push
