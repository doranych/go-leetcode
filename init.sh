#!/bin/bash

echo "Init new problem"

pn=$1
pkg=$(echo "$pn" | tr '[:upper:]' '[:lower:]' | tr ' ' '_')

mkdir -p $pkg
echo "package $pkg" > "$pkg"/main.go
cp "$pkg"/main.go "$pkg"/main_test.go
printf 'Paste a link to the problem\n'

read link
printf "[%s](%s)\n===" "$pn" "$link" > "$pkg"/problem.md

echo "Done"

git add ./*