# Utils Example

This is an example application that demonstrates the usage of the `utils` package. The `utils` package provides utility functions for handling common tasks in a Go application.

## Features

- Shows version information of the application
- Shows help information about the application
- Listens for interrupt signals and gracefully shuts down the server
- Provides a simple HTTP server that responds with "hello world"

## Prerequisites

- Go programming language (version 1.21 or later)
- `github.com/spf13/pflag` package
- `github.com/spf13/viper` package
- `github.com/twistingmercury/utils` package

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/twistingmercury/utils
   ```

2. Change to the project directory:
   ```
   cd utils-example
   ```

3. Build the application:
   ```
   go build -ldflags="-X main.buildVersion=1.0.0 -X main.buildDate=2021-09-01T00:00:00Z -X main.commitHash=22ae825" -o bin/utils-example main.go
   ```

   Note: The `ldflags` are used to set the build version, build date, and commit hash during the build process.

## Usage

1. Run the application:
   ```
   ./utils-example
   ```

   This will start the HTTP server and listen on port 8080.

2. Open a web browser or use a tool like `curl` to make a request to `http://localhost:8080`. You should see the response "hello world".

3. To show the version information, run the application with the `--version` flag:
   ```
   ./utils-example --version
   ```

4. To show the help information, run the application with the `--help` flag:
   ```
   ./utils-example --help
   ```

5. To stop the server, press `Ctrl+C` in the terminal where the application is running.

## Configuration

The `Initialize` function in the `main` package sets up the configuration using the `pflag` and `viper` packages. It binds the command-line flags to the corresponding viper keys and sets the default values for the build version, build date, and commit hash.

The `buildVersion`, `buildDate`, and `commitHash` variables are intended to be set during the build process using `ldflags`.

## Customization

You can customize the behavior of the application by modifying the code in the `main` package. For example, you can add more routes to the HTTP server or change the configuration settings.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).