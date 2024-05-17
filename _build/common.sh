checkEnv() {
    if [ -z "$1" ]; then
        printf "** env var %s is not assigned\n" "$2"
        help
    fi
}

help() {
    echo "\nHow to use build.sh:"
    echo '  BUILD_DATE="$(BUILD_DATE)" \'
	echo '  BUILD_VER="$(BUILD_VER)" \'
	echo '  GIT_COMMIT="$(GIT_COMMIT)" \'
	echo '  DOCKERFILE_DIR="$(PWD)" \'
	echo '  ALPINE_VERSION="$(ALPINE_VERSION)" \'
	echo '  GO_VERSION="$(GO_VERSION)" \'
	echo '  DESCRIPTION="$(DESCRIPTION)" \'
	echo '  VENDOR="$(VENDOR)" \'
	echo '  ./_build/build.sh'
    echo "\nEnvironment variables:"
    echo "  BUILD_DATE:     The build date of the binary"
    echo "  BUILD_VER:      The build semantic version (if a release candidate) of the binary."
    echo "  GIT_COMMIT:     The short commit hash of the commit being used for the build."
    echo "  DOCKERFILE_DIR: The directory containing the target dockerfile."
    echo "  ALPINE_VERSION: The version of the alpine image to use."
    echo "  GO_VERSION:     The version of the go image to use, as available in the Alpine go images."
    echo "  DESCRIPTION:    The description of the binary to be used in the Dockerfile."
    echo "  VENDOR:         The vendor of the binary."
    echo "  TARGET:         The file that will be targeted for the build, e.g., main.go."
    exit 1
}