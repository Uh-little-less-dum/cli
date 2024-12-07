#!/bin/zsh

cd /Users/bigsexy/Desktop/Go/projects/github.com/Uh-little-less-dum/cli && go test ./... -coverprofile coverage/coverage.out
go-cover-treemap -coverprofile coverage/coverage.out > coverage/coverage.svg
go tool cover -func=coverage/coverage.out
