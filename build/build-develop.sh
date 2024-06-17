#!/usr/bin/env bash

set -e

SCRIPT_ROOT=$(dirname "${BASH_SOURCE}")
source "$SCRIPT_ROOT/common.sh"
PROJ_ROOT=$(cd -- "${SCRIPT_ROOT}/.." && pwd)
DATE=$(date +"%Y-%m-%d")

DOCKER_BUILDKIT=1 docker build --force-rm \
--build-arg BUILD_DATE="$DATE" \
--build-arg BUILD_VER="$DATE-develop" \
-f build/Dockerfile "$PROJ_ROOT" \
--target build \
-t token_go_bin:"develop-$DATE"

docker system prune -f
