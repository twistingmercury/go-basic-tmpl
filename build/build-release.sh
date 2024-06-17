#!/usr/bin/env bash

set -e

SCRIPT_ROOT=$(dirname "${BASH_SOURCE}")
source "$SCRIPT_ROOT/common.sh"
PROJ_ROOT=$(cd -- "${SCRIPT_ROOT}/.." && pwd)
DATE=$(date +"%Y-%m-%d")

common::checkenv "BUILD_VER"

DOCKER_BUILDKIT=1 docker build --force-rm \
--build-arg BUILD_DATE="$DATE" \
--build-arg BUILD_VER="$BUILD_VER" \
-f build/Dockerfile "$PROJ_ROOT" \
-t token_go_bin:"$BUILD_VER"

docker system prune -f
