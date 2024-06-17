#!/usr/bin/env bash

function common::checkenv() {
    if [ -z "${!1}" ]; then
        printf "** Error: $1 must be defined\n"
        common::help
    fi
}

function common::help() {
    echo "Usage:"
    echo "  BUILD_DATE=<BUILD_DATE> BUILD_VER=<BUILD_VER> DOCKERFILE_DIR=<DOCKERFILE_DIR> ./build/build-image.sh"
    echo "\nEnvironment variables:"
    echo "  BUILD_DATE:     The build date of the binary"
    echo "  BUILD_VER:      The build semantic version (if a release candidate) of the binary"
    exit 1
}