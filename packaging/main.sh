#!/bin/sh

# get version number from LibmuttonVersion
cd ..
printf 'package main\nimport (\n"fmt"\n"github.com/rwinkhart/MUTN/src/backend"\n)\nfunc main() {\nfmt.Println(backend.LibmuttonVersion)\n}' > ./version.go
version=$(go run ./version.go)
rm ./version.go
cd ./packaging

# get revision number from user input, fallback to 1 if not provided
if [ -z "$2" ]; then
    revision=1
else
    revision="$2"
fi

# ensure output directory exists
mkdir -p ./output

case "$1" in
    pkgbuild-git-stable)
        . ./pkgbuild-git-stable.sh
        create_pkgbuild_git_stable
        ;;
    *)
    printf '\nusage: package.sh [target] <revision>\n\ntargets: pkgbuild-git-stable\n\n'
    ;;
esac
