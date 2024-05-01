[![Run Tests](https://github.com/twistingmercury/utils/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/twistingmercury/utils/actions/workflows/tests.yml)
[![codecov](https://codecov.io/gh/twistingmercury/utils/branch/main/graph/badge.svg?token=4FKAKJS1H0)](https://codecov.io/gh/twistingmercury/utils)
[![CodeQL](https://github.com/twistingmercury/utils/actions/workflows/codeql.yml/badge.svg?branch=main)](https://github.com/twistingmercury/utils/actions/workflows/codeql.yml)
# Utils Package

The `utils` package is a small collection of utility functions that provide common functionality for Go applications. It aims to simplify tasks such as showing version information, displaying help text, handling interrupt signals, and performing error handling. I decided to put this in its own importable package since I basically use this in all non-trivial projects.

## Features

- `ShowVersion`: Prints the version information of the application and exits.
- `ShowHelp`: Prints the help information of the application and exits.
- `ListenForInterrupt`: Listens for interrupt signals (e.g., Ctrl+C) and gracefully stops the application.
- `FailFast`: Panics if an error is encountered, providing a convenient way to handle errors.

## Installation

To use the `utils` package in your Go application, you can install it using the following command:

```
go get github.com/twistingmercury/utils
```

## Usage

Here's an example of how to use the `utils` package in your Go application:

```go
package main

import (
	"context"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/twistingmercury/utils"
	"log"
)

func main() {
	Initialize()
	ctx, cancelFunc := context.WithCancel(context.Background())
	utils.ListenForInterrupt(cancelFunc)
	utils.ShowHelp()
	utils.ShowVersion()

	// Set up your application logic here

	log.Println("Application started")
	log.Println("Press Ctrl+C to stop the application")
	<-ctx.Done()
}

func Initialize() {
	pflag.Parse()
	viper.BindPFlag(utils.ViperShowHelpKey, pflag.Lookup(utils.ViperShowHelpKey))
	viper.BindPFlag(utils.ViperShowVersionKey, pflag.Lookup(utils.ViperShowVersionKey))
	viper.Set(utils.ViperServiceNameKey, "your-app-name")

	// Set the default values for the viper keys
	viper.Set(utils.ViperBuildVersionKey, "1.0.0")
	viper.Set(utils.ViperBuildDateKey, "2023-06-09")
	viper.Set(utils.ViperCommitHashKey, "abcdef")
}
```

In the example above, the `utils` package is imported, and its functions are used to handle common tasks such as listening for interrupt signals, showing version and help information, and handling errors.

The `Initialize` function sets up the configuration using the `pflag` and `viper` packages. It binds the command-line flags to the corresponding viper keys and sets the default values for the build version, build date, and commit hash.

## Configuration

The `utils` package uses the following configuration keys:

- `ViperShowHelpKey`: Determines whether to show the help information.
- `ViperShowVersionKey`: Determines whether to show the version information.
- `ViperBuildVersionKey`: Specifies the build version of the application.
- `ViperBuildDateKey`: Specifies the build date of the application.
- `ViperCommitHashKey`: Specifies the commit hash of the application.
- `ViperServiceNameKey`: Specifies the name of the service or application.

These configuration keys can be set using command-line flags or environment variables.

## Contributing

Contributions to the `utils` package are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request on the [GitHub repository](https://github.com/twistingmercury/utils).

## License

The `utils` package is open-source software licensed under the [MIT License](https://opensource.org/licenses/MIT).