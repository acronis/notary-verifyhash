#!/usr/bin/env bash
export GOFLAGS=-mod=vendor

cd src/modules
case $1 in
    Windows_NT)
    export GOOS=windows
    go build -o verifyhash.exe
    mv verifyhash.exe /mnt
    ;;
    Darwin)
    export GOOS=darwin
    go build -o verifyhash
    mv verifyhash /mnt
    ;;
    *)
    go build -o verifyhash
    mv verifyhash /mnt
    ;;
esac