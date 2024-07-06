#!/bin/sh
gofmt -l -w -s ./src/cli/*.go
gofmt -l -w -s ./src/backend/*.go
git add -f extra src wiki .gitignore commit.sh go.mod go.sum LICENSE mutn.go libmuttonserver.go README.md
git commit -m "$1"
git push
