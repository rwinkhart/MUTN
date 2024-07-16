#!/bin/sh
gofmt -l -w -s ./src/cli/*.go ./src/backend/*.go ./mutn.go ./libmuttonserver.go
git commit -am "$1"
git push
