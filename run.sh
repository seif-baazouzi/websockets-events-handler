#!/bin/bash

if [[ $1 == "server" ]]; then
    go run ./test-server/main.go

fi


if [[ $1 == "client" ]]; then
    go run ./test-client/main.go

fi
