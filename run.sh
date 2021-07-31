#!/bin/sh

export ROOT_FOLDER=${PWD}/temp
export SERVER_ADDRESS="0.0.0.0"
export SERVER_PORT=3440
export SECRET=secret

go run ./cmd/server/main.go