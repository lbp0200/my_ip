#!/usr/bin/env bash

env GOOS=linux GOARCH=amd64 go build -o out/linux_64_my_ip

env GOOS=windows GOARCH=amd64 go build -o out/windows_64_my_ip.exe