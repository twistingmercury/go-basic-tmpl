# How to use and update the build scripts

This directory contains two shell scripts, `build.sh` and `common.sh`, which are used to build the Docker image for the Go application.

## build.sh

The `build.sh` script is responsible for building the Docker image for the Go application. It takes several environment variables as input to customize the build process.

### Usage

To use the `build.sh` script, run the following command:

```
BUILD_DATE="$(BUILD_DATE)" \
BUILD_VER="$(BUILD_VER)" \
GIT_COMMIT="$(GIT_COMMIT)" \
DOCKERFILE_DIR="$(PWD)" \
ALPINE_VERSION="$(ALPINE_VERSION)" \
GO_VERSION="$(GO_VERSION)" \
DESCRIPTION="$(DESCRIPTION)" \
VENDOR="$(VENDOR)" \
TARGET="$(TARGET)" \
./_build/build.sh
```

### Environment Variables

The following environment variables are used by the `build.sh` script:

- `BUILD_DATE`: The build date of the binary
- `BUILD_VER`: The build semantic version (if a release candidate) of the binary
- `GIT_COMMIT`: The short commit hash of the commit being used for the build
- `DOCKERFILE_DIR`: The directory containing the target Dockerfile
- `ALPINE_VERSION`: The version of the Alpine image to use
- `GO_VERSION`: The version of the Go image to use, as available in the Alpine Go images
- `DESCRIPTION`: The description of the binary to be used in the Dockerfile
- `VENDOR`: The vendor of the binary
- `TARGET`: The file that will be targeted for the build, e.g., `main.go`

### Functionality

The `build.sh` script performs the following steps:

1. Sources the `common.sh` script to access shared functions
2. Checks if the required environment variables are set using the `checkEnv` function from `common.sh`
3. Changes the current directory to the `DOCKERFILE_DIR`
4. Builds the Docker image using the provided environment variables and the Dockerfile in the current directory
5. Prunes the Docker system to remove unused images and containers

## common.sh

The `common.sh` script contains helper functions used by the `build.sh` script.

### Functions

- `checkEnv`: Checks if an environment variable is set and prints an error message if it's not set
- `help`: Prints usage information for the `build.sh` script, including the required environment variables

### Usage

The `common.sh` script is sourced by the `build.sh` script and is not meant to be run directly. Its functions are used to validate environment variables and provide help information.

## Adding Environment Variables

If you need to add additional environment variables to customize the build process, follow these steps:

1. Open the `build.sh` script in a text editor.

2. Locate the section where the existing environment variables are checked using the `checkEnv` function, e.g.:

```bash
common::checkEnv "BUILD_DATE"
common::checkEnv "BUILD_VER"
common::checkEnv "GIT_COMMIT"
# ...
```

3. Add a new line for each environment variable you want to add, following the same format:

```bash
common::checkEnv "YOUR_NEW_ENV_VAR"
```

Replace `YOUR_NEW_ENV_VAR` with the name of your new environment variable.

4. Locate the `docker build` command in the `build.sh` script and add your new environment variable as a build argument:

```bash
docker build --force-rm \
  --build-arg BUILD_DATE="$BUILD_DATE" \
  --build-arg BUILD_VER="$BUILD_VER" \
  --build-arg GIT_COMMIT="$GIT_COMMIT" \
  # ...
  --build-arg YOUR_NEW_ENV_VAR="$YOUR_NEW_ENV_VAR" \
  -t "{{bin_name}}":"$BUILD_VER" -f dockerfile .
```

Replace `YOUR_NEW_ENV_VAR` with the name of your new environment variable.

5. Open the `dockerfile` in a text editor and add a corresponding `ARG` for your new environment variable:

```dockerfile
ARG YOUR_NEW_ENV_VAR
```

Replace `YOUR_NEW_ENV_VAR` with the name of your new environment variable. This allows the Dockerfile to accept the value passed from the `build.sh` script.

6. Open the `common.sh` script in a text editor.

7. Locate the `help` function and add a new line to the "Environment variables" section for each new environment variable you added:

```bash
echo "  YOUR_NEW_ENV_VAR: Description of your new environment variable"
```

Replace `YOUR_NEW_ENV_VAR` with the name of your new environment variable and provide a brief description of its purpose.

8. Save the changes to `build.sh`, `dockerfile`, and `common.sh`.

Now, when running the `build.sh` script, make sure to provide a value for your new environment variable:

```bash
YOUR_NEW_ENV_VAR="value" ./build.sh
```

The `build.sh` script will check for the presence of your new environment variable and pass it to the Docker build command as a build argument. The Dockerfile will accept the value using the corresponding `ARG`. The `help` function in `common.sh` will also display information about your new environment variable when invoked.

Remember to use the value of your new environment variable within the Dockerfile as needed, since you've added the corresponding `ARG`.

## Contributing

If you'd like to contribute to the build scripts, please follow the contribution guidelines outlined in the main README file of the project.

## License

The build scripts are released under the [MIT License](LICENSE).