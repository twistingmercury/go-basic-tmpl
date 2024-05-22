#!/usr/bin/env bash

function common::checkenv() {
    if [ -z "${!1}" ]; then
        printf "** Error: $1 must be defined\n"
        common::help
        exit 1
    fi
}

function common::help() {
    echo "Usage:"
    echo "  BUILD_DATE=<BUILD_DATE> \\n  BUILD_VER=<BUILD_VER> \\n  GIT_COMMIT=<GIT_COMMIT> \\n  BIN_NAME=<BIN_NAME> \\n  MODULE_NAME=<MODULE_NAME> \\n  DOCKERFILE_DIR=<DOCKERFILE_DIR> \\n  ./build/build-image.sh"
    echo "\nEnvironment variables:"
    echo "  BUILD_DATE:     The build date of the binary"
    echo "  BUILD_VER:      The build semantic version (if a release candidate) of the binary"
    echo "  GIT_COMMIT:     The short commit hash of the commit being used for the build"
    echo "  BIN_NAME:       The binary name as well as the name of the process to be executed in the ENTRYPOINT"
    echo "  MODULE_NAME:    The module name as defined in go.mod"
    echo "  DOCKERFILE_DIR: The directory containing the target dockerfile"
    exit 1
}