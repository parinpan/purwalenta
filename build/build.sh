#!/bin/bash

dep ensure -v
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./cmd/purwalenta/purwalenta-bin ./cmd/purwalenta/
