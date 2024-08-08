#!/bin/sh
gofmt -l -w -s ./src/cli/*.go ./mutn.go
git commit -am "$1"
git push
