#!/bin/sh

cd /Users/bigsexy/Desktop/Go/projects/ulld/cli || exit

source /Users/bigsexy/Desktop/Go/projects/ulld/cli/.env.zsh

git add .

git commit

TAG=$(gum input --placeholder "tag")

git tag -a "$TAG"

echo "Now create a tag with git tag -a x.x.x and then you can run gorelease without any issues... hopefully."
