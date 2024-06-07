#!/usr/bin/env bash

SCRIPT_ROOT=$(dirname "${BASH_SOURCE}")
source "$SCRIPT_ROOT/common.sh"

common::checkenv "BUILD_DATE"
common::checkenv "BUILD_VER"
common::checkenv "DOCKERFILE_DIR"

printf "\n** Changing directory to '%s'\n" "$DOCKERFILE_DIR"
cd "$DOCKERFILE_DIR"

printf "\n** Building Docker image for 'MODULE_NAME/BIN_NAME', version '%s'\n" "$BUILD_VER"
docker build --force-rm \
	--build-arg BUILD_DATE="$BUILD_DATE" \
	--build-arg BUILD_VER="$BUILD_VER" \
	-t "BIN_NAME":"$BUILD_VER" -f Dockerfile .

docker system prune -f