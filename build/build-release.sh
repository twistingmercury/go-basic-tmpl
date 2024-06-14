#!/usr/bin/env bash

set -e

SCRIPT_ROOT=$(dirname "${BASH_SOURCE}")
source "$SCRIPT_ROOT/common.sh"
PROJ_ROOT=$(cd -- "${SCRIPT_ROOT}/.." && pwd)

common::checkenv "BUILD_DATE"
common::checkenv "BUILD_VER"

printf "\n** Building Docker image for 'token_go_module/token_bin', version '%s'\n" "$BUILD_VER"
DOCKER_BUILDKIT=1 docker build --force-rm \
	--build-arg BUILD_DATE="$BUILD_DATE" \
	--build-arg BUILD_VER="$BUILD_VER" \
  -f build/Dockerfile "$PROJ_ROOT" \
	-t token_bin:"$BUILD_VER"

docker system prune -f