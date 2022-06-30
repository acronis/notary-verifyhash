#!/usr/bin/env bash
export GOFLAGS=-mod=vendor

cd src/modules
go build -o verifyhash
mv verifyhash /mnt
