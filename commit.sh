#!/bin/sh
gofmt -l -w -s ./src/cli/*.go
gofmt -l -w -s ./src/backend/*.go
git add -f extra src .gitignore commit.sh go.mod go.sum LICENSE main.go mutn.go libmuttonserver.go README.md
git commit -m "$1"
git push
