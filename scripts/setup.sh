#!/bin/sh

cp ./scripts/pre-commit ./.git/hooks/
cp .env-example .env

go install github.com/gordonklaus/ineffassign@latest
go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
go install honnef.co/go/tools/cmd/staticcheck@latest