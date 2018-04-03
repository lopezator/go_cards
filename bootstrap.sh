#!/bin/sh

go get github.com/derekparker/delve/cmd/dlv;
go build -gcflags='-N -l' github.com/lopezator/go_cards &&
dlv --listen=:5432 --headless=true --api-version=2 exec go_cards