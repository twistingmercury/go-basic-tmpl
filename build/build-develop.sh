#!/usr/bin/env bash

set -e

SCRIPT_ROOT=$(dirname "${BASH_SOURCE}")
source "$SCRIPT_ROOT/common.sh"

PROJ_ROOT=$(cd -- "${SCRIPT_ROOT}/.." && pwd)
START_OF_YEAR=$(date -j -f "%Y-%m-%d %H:%M:%S" "$(date +%Y)-01-01 00:00:00" +%s)
CURRENT_TIME=$(date +%s)
DATE=$((CURRENT_TIME - START_OF_YEAR))

echo $DATE

DOCKER_BUILDKIT=1 docker build --force-rm \
	--build-arg BUILD_DATE="$DATE" \
	--build-arg BUILD_VER="$DATE-develop" \
  -f build/Dockerfile "$PROJ_ROOT" \
	-t token_bin:"$DATE-develop"

docker system prune -f