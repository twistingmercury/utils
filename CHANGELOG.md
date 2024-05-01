# Changelog
All notable changes to the utils package will be documented in this file.
The format is based on Keep a Changelog,
and this project adheres to Semantic Versioning.

## 1.0.0 - 2023-06-09

### Added

* Initial release of the utils package.
* Added ShowVersion function to display version information of the application.
  * Reads version details from viper configuration keys.
  * Prints the version, build date, and commit hash.
  * Exits the program after displaying the version.
  
* Added ShowHelp function to display help information of the application.
  * Reads help details from viper configuration keys.
  * Prints the version, build date, commit hash, and usage information.
  * Exits the program after displaying the help.

* Added ListenForInterrupt function to handle interrupt signals gracefully.
  * Listens for interrupt signals (SIGINT, SIGTERM) in a separate goroutine.
  * Cancels the provided context when an interrupt signal is received.
  * Prints a message indicating graceful stopping of the application.

* Added FailFast function intended to be used when an action by the app should ause the app to fail.
  * Panics if the provided error is not nil.
  * Logs an error message along with the error details.

* Defined constants for viper configuration keys.
  * `ViperShowHelpKey`: Key for the flag to show help information.
  * `ViperShowVersionKey`: Key for the flag to show version information.
  * `ViperBuildVersionKey`: Key for the build version.
  * `ViperBuildDateKey`: Key for the build date.
  * `ViperCommitHashKey`: Key for the commit hash.
  * `ViperServiceNameKey`: Key for the service name.