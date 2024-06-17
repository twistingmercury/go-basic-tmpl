#!/usr/bin/env bash

function common::checkenv() {
    if [ -z "${!1}" ]; then
        printf "** Error: $1 must be defined\n"
        common::help
    fi
}

function common::help() {
    echo "Usage:"
    echo "  BUILD_VER=<BUILD_VER> ./build/build-[develop | release].sh"
    echo "\nEnvironment variables:"
    echo "  BUILD_VER:      The build semantic version (if a release candidate) of the binary"
    exit 1
}