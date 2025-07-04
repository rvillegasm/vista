#!/usr/bin/env bash
set -euo pipefail

# Formatting check
unformatted=$(gofmt -l .)
if [[ -n "$unformatted" ]]; then
  echo "The following files need gofmt formatting:\n$unformatted"
  exit 1
fi

# Run vet, tests and lint

go vet ./...

go test ./... -v

golangci-lint run --timeout 5m 