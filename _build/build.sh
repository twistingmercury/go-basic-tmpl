#!/usr/bin/env bash

set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE}")
source $SCRIPT_ROOT/common.sh

if $1 == "help"; then
    help
fi

checkEnv "$BUILD_DATE" "BUILD_DATE"
checkEnv "$BUILD_VER" "BUILD_VER"
checkEnv "$GIT_COMMIT" "GIT_COMMIT"
checkEnv "$DOCKERFILE_DIR" "DOCKERFILE_DIR"
checkEnv "$ALPINE_VERSION" "ALPINE_VERSION"
checkEnv "$GO_VERSION" "GO_VERSION"
checkEnv "$TARGET" "TARGET"

printf "\n** Changing directory to '%s'\n" "$DOCKERFILE_DIR"
cd "$DOCKERFILE_DIR"

printf "\n** Building Docker image for '%s' with version '%s'\n" "{{bin_name}}" "$BUILD_VER"
docker build --force-rm \
	--build-arg BUILD_DATE="$BUILD_DATE" \
	--build-arg BUILD_VER="$BUILD_VER" \
	--build-arg GIT_COMMIT="$GIT_COMMIT" \
	--build-arg ALPINE_VERSION="$ALPINE_VERSION" \
	--build-arg GO_VERSION="$GO_VERSION" \
	--build-arg TARGET="$TARGET" \
	-t "{{bin_name}}":"$BUILD_VER" -f dockerfile .

docker system prune -f