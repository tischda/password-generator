@echo off

set CGO_ENABLED=1
set GOOS=windows
set GOARCH=amd64 

make build
