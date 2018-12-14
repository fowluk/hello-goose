#!/bin/bash
env GOOS=linux GOARCH=amd64 go build -o app/binary ../go/main.go
