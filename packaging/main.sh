#!/bin/sh

# get version number from cli package
cd ..
printf 'package main\nimport (\n"fmt"\n"github.com/rwinkhart/MUTN/src/cli"\n)\nfunc main() {\nfmt.Println(cli.MUTNVersion)\n}' > ./version.go
version=$(go run -ldflags="-s -w" -tags noMarkdown ./version.go)
rm ./version.go
cd ./packaging

# get revision number from user input, fallback to 1 if not provided
if [ -z "$2" ]; then
    revision=1
else
    revision="$2"
fi

# ensure output directory exists
mkdir -p ./1output

case "$1" in
    release-binaries)
        . ./resources/release-binaries.sh
        create_release_binaries
        ;;
    pkgbuild-git-stable-optimized)
        . ./resources/pkgbuild-git-stable-optimized.sh
        create_pkgbuild_git_stable_optimized
        ;;
    apkbuild-source-stable-optimized)
        . ./resources/apkbuild-source-stable-optimized.sh
        create_apkbuild_source_stable_optimized
        ;;
    *)
    printf '\nusage: package.sh [target] <revision>\n\ntargets: release-binaries pkgbuild-git-stable-optimized apkbuild-source-stable-optimized\n\n'
    ;;
esac
