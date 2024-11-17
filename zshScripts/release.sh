#!/bin/sh

cd /Users/bigsexy/Desktop/Go/projects/ulld/cli || exit

. /Users/bigsexy/Desktop/Go/projects/ulld/cli/.env.zsh

git add .

git commit

TAG=$(gum input --placeholder "tag")

git tag -a "$TAG"

goreleaser
