# About this template

This is a basic template for creating Go projects with a focus on observability, containerization, and best practices. The template provides a solid foundation for building robust and maintainable Go applications, incorporating essential components such as logging, tracing, metrics, and health monitoring.

## Features

- Structured logging using the [zerolog](https://pkg.go.dev/github.com/rs/zerolog) package
- Tracing with OpenTelemetry and OTLP exporter
- Metrics collection Prometheus
- Configuration management using the [viper](https://pkg.go.dev/github.com/spf13/viper) package
- Health monitoring with a heartbeat endpoint
- Dockerfile for containerization
- Shell scripts for building the application
- Makefile for common development tasks

## Getting Started

To create a new Go project using this template, you'll need to use the [scaffolder](https://github.com/twistingmercury/scaffolder) CLI tool. The [scaffolder](https://github.com/twistingmercury/scaffolder) helps you initialize a new project by replacing the template's tokens with your project-specific values.

Since this is a template, understand that cloning this project directly and attempting to compile it will fail. This is due to the replacement tokens scattered throughout the template:

- `MODULE_NAME`     This is the name of the module, as declared in the go.mod file, i.e., `github.com/your-name/your-project`
- `BIN_NAME`        This will be the name of the compiled binary, the name of the root folder, and the name of the docker image.
- `IMG_VENDOR_NAME    ` This is the name of the vendor to be added to the `org.opencontainers.image.vendor` label in the docker image.
- `IMG_DESCRIPTION`     This is a description of the app to be added to the `org.opencontainers.image.description` label in the docker image.

### Prerequisites

- Go 1.21 or higher
- Docker (for containerization)
- [scaffolder](https://github.com/twistingmercury/scaffolder) to use the template
- GNU Make

### Initializing a New Project

You don't have to clone this project to use it as a template. Scaffolder will *always* clone the project itself, always pulling from the `main` branch. Even if you do clone it, it will not use the locally cloned project.

To initialize a new Go project using this template, run the following command:

```bash
scaffolder init --module "my/module/name" --bin-name "myapp" --vendor-name "my name" --description "this is a short description"
```

Replace the following placeholders with your project-specific values:

- `my/module/name`:              The Go module name for your project
- `myapp`:                       The name of the binary that will be created
- `my name`:                     Your name or the name of the project's vendor
- `this is a short description`: A brief description of your project

The `scaffolder` will create a new directory with your project's name and replace the template's tokens (`MODULE_NAME`, `BIN_NAME`, `IMG_VENDOR_NAME`, and `IMG_DESCRIPTION`) with the provided values.

### Using the Makefile

The provided Makefile includes several targets to help with common development tasks:

- `clean`: Remove build artifacts and clean up the project
- `bin`: Build the binary and output to the `_bin/` directory
- `test`: Run all unit tests and generate a coverage report
- `image`: Build the Docker image using a multi-stage build
- `run`: Run the `main.go` file to start the server
- `test-heartbeat`: Test the heartbeat endpoint using cURL

To use the Makefile, simply run `make <target>` in your terminal, replacing `<target>` with the desired target name.

## Contributing

Instead of modifying this template directly, I recommend that you fork this project and modify it to your liking. The 
[scaffolder](https://github.com/twistingmercury/scaffolder) can be pointed to a different repository by using its `--template` flag. The default for that flag is this project.

## License

This Go template project template is released under the [MIT License](./LICENSE).