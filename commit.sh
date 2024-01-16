#!/bin/sh
git add -f extra src/*.go src/*.mod .gitignore commit.sh LICENSE README.md
git commit -m "$1"
git push
