#!/bin/sh

# get version number from backend TODO embed Go code that reads the value of LibmuttonVersion
version=$(grep 'LibmuttonVersion = "' ../src/backend/1globals.go | cut -c22-26)

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
