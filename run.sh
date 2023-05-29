#!/bin/sh

go build -o server ./cmd/server/main.go

exit_status=$?
if [ $exit_status -ne 0 ]; then
    echo "Error"
    exit $exit_status
fi

export ROOT_FOLDER=${PWD}/temp
export SERVER_ADDRESS="0.0.0.0"
export SERVER_PORT=3440
export SECRET=secret

./server