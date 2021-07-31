#!/bin/sh

export ROOT_URL=http://localhost:3440
export SECRET=secret

go run ./cmd/client/main.go 0.0.2 /home/apin/go-deploy/test_file