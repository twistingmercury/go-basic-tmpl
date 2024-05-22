#!/usr/bin/env bash

SCRIPT_ROOT=$(dirname "${BASH_SOURCE}")
source "$SCRIPT_ROOT/common.sh"

common::checkenv "BUILD_DATE"
common::checkenv "BUILD_VER"
common::checkenv "GIT_COMMIT"
common::checkenv "DOCKERFILE_DIR"
common::checkenv "ALPINE_VERSION"
common::checkenv "GO_VERSION"

printf "\n** Changing directory to '%s'\n" "$DOCKERFILE_DIR"
cd "$DOCKERFILE_DIR"

printf "\n** Building Docker image for 'tunnelvision', version '%s'\n" "$BUILD_VER"
docker build --force-rm \
	--build-arg BUILD_DATE="$BUILD_DATE" \
	--build-arg BUILD_VER="$BUILD_VER" \
	--build-arg GIT_COMMIT="$GIT_COMMIT" \
	--build-arg ALPINE_VERSION="$ALPINE_VERSION" \
	--build-arg GO_VERSION="$GO_VERSION" \
	-t "tunnelvision":"$BUILD_VER" -f Dockerfile .

docker system prune -f