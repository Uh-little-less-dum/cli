#!/bin/bash

ADD="Add"
RESET="Reset"
UNSTAGE="Unstage"

ACTION=$(gum choose "$ADD" "$RESET" "$UNSTAGE")

if [ "$ACTION" == "$ADD" ]; then
    git status --short | cut -c 4- | gum choose --no-limit | xargs git add
elif [ "$ACTION" == "$UNSTAGE" ]; then
    git diff --name-only --cached | gum choose --no-limit | xargs git restore --staged
else
    git status --short | cut -c 4- | gum choose --no-limit | xargs git restore
fi
