#!/usr/bin/env bash

set -euo pipefail

TOOLS=(
  "github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.7.1"
)

for tool in "${TOOLS[@]}"; do
  go install "$tool"
done
