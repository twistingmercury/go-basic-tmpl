
---

NOTE: This directory is suggested by [Standard Go Project Layout](https://github.com/golang-standards/project-layout/tree/master).

See: [project-layout
/build/](https://github.com/golang-standards/project-layout/tree/master/build) for a description.

---

# Directory: build



This directory contains three shell scripts, `build-develop.sh`, `build-release.sh`, and `common.sh`. These are used to build the Docker image for the application. The `build-develop.sh` script is responsible for building a development build Docker image for the Go application. The `build-release.sh` script builds the image intended for release.

### Usage

### Environment Variables

The following environment variables are used by the `build-release.sh` script:

- `BUILD_VER`: The build semantic version (if a release candidate) of the binary

The `build-develop.sh` doesn't require any environmental variables.

### Functionality

The `build.sh` script performs the following steps:

1. Sources the `common.sh` script to access shared functions
2. Checks if there are any required environment variables using the `checkEnv` function from `common.sh`
3. Builds the Docker image using any provided environment variables.
4. Prunes the Docker system to remove unused images and containers.

## common.sh

The `common.sh` script contains helper functions used by the build scripts.

### Functions

- `checkEnv`: Checks if an environment variable is set and prints an error message if it's not set
- `help`: Prints usage information for the `build.sh` script, including the required environment variables

### Usage

The `common.sh` script is sourced by the build scripts and is not meant to be run directly. Its functions are used to validate environment variables and provide help information.

## Adding Environment Variables

If you need to add additional environment variables to customize the build process, follow these steps:

1. Open the `build-[develop | release].sh` script in a text editor.

2. Locate the section where the existing environment variables are checked using the `checkEnv` function, e.g.:
   ```bash
    # In build-release.sh 
    common::checkEnv "BUILD_VER"
    # ...
   ```

3. Add a new line for each environment variable you want to add, following the same format:
    ```bash
    common::checkEnv "YOUR_NEW_ENV_VAR"
    ```
   
4. Locate the `docker build` command in the `build.sh` script and add your new environment variable as a build argument:
    ```bash
    docker build --force-rm \
      --build-arg BUILD_DATE="$BUILD_DATE" \
      --build-arg BUILD_VER="$BUILD_VER" \
      # ...
      --build-arg YOUR_NEW_ENV_VAR="$YOUR_NEW_ENV_VAR" \
      -t "devapp":"$BUILD_VER" -f Dockerfile .
    ```

5. Open the `dockerfile` in a text editor and add a corresponding `ARG` for your new environment variable:
    ```dockerfile
    ARG YOUR_NEW_ENV_VAR
    ```

6. Open the `common.sh` script in a text editor. Locate the `help` function and add a new line to the "Environment variables" section for each new environment variable you added:
    ```bash
    echo "  YOUR_NEW_ENV_VAR: Description of your new environment variable"
    ```
7. Save the changes to `build.sh`, `dockerfile`, and `common.sh`.

Now, when running the `build.sh` script, make sure to provide a value for your new environment variable:

```bash
YOUR_NEW_ENV_VAR="value" ./build-release.sh
```

The `build.sh` script will check for the presence of your new environment variable and pass it to the Docker build command as a build argument. The Dockerfile will accept the value using the corresponding `ARG`. The `help` function in `common.sh` will also display information about your new environment variable when invoked.

Remember to use the value of your new environment variable within the Dockerfile as needed, since you've added the corresponding `ARG`.
