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
checkEnv "$BIN_NAME" "BIN_NAME"
checkEnv "$MODULE_NAME" "MODULE_NAME"
checkEnv "$DOCKERFILE_DIR" "DOCKERFILE_DIR"
checkEnv "$ALPINE_VERSION" "ALPINE_VERSION"
checkEnv "$GO_VERSION" "GO_VERSION"
checkEnv "$DESCRIPTION" "DESCRIPTION"
checkEnv "$VENDOR" "VENDOR"

printf "\n** Changing directory to '%s'\n" "$DOCKERFILE_DIR"
cd "$DOCKERFILE_DIR"

## Generate the ENTRYPOINT command for the Dockerfile
ENTRY_POINT="\nENTRYPOINT [\"/app/$BIN_NAME\"]"
printf "\nENTRYPOINT [\"/app/$BIN_NAME\"]" >> dockerfile

printf "\n** Building Docker image for '%s' with version '%s'\n" "$BIN_NAME" "$BUILD_VER"
docker build --force-rm \
	--build-arg BUILD_DATE="$BUILD_DATE" \
	--build-arg BUILD_VER="$BUILD_VER" \
	--build-arg GIT_COMMIT="$GIT_COMMIT" \
	--build-arg BIN_NAME="$BIN_NAME" \
	--build-arg MODULE_NAME="$MODULE_NAME" \
	--build-arg ALPINE_VERSION="$ALPINE_VERSION" \
	--build-arg GO_VERSION="$GO_VERSION" \
    --build-arg DESCRIPTION="$ENTRY_POINT" \
	-t "$BIN_NAME":"$BUILD_VER" -f dockerfile .

docker system prune -f

## Undo the addition of the ENTRYPOINT command from the Dockerfile
perl -i -ne 'print unless /^ENTRYPOINT/' dockerfile