package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const ( // viper conf keys
	ViperShowHelpKey     = "help"
	ViperShowVersionKey  = "version"
	ViperBuildVersionKey = "build_version"
	ViperBuildDateKey    = "build_date"
	ViperCommitHashKey   = "commit_hash"
	ViperServiceNameKey  = "service_name"
)

var exitFunc = os.Exit

// ShowVersion prints the version information and exits the program.
func ShowVersion() {
	if !viper.GetBool(ViperShowVersionKey) {
		return
	}

	fmt.Printf("%s\nversion: %s; build date: %s; commit: %s\n",
		viper.GetString(ViperServiceNameKey),
		viper.GetString(ViperBuildVersionKey),
		viper.GetString(ViperBuildDateKey),
		viper.GetString(ViperCommitHashKey),
	)
	exitFunc(0)
}

// ShowHelp prints the help information and exits the program.
func ShowHelp() {
	if !viper.GetBool(ViperShowHelpKey) {
		return
	}

	fmt.Printf("%s\nversion: %s; build date: %s; commit: %s\n",
		viper.GetString(ViperServiceNameKey),
		viper.GetString(ViperBuildVersionKey),
		viper.GetString(ViperBuildDateKey),
		viper.GetString(ViperCommitHashKey),
	)
	fmt.Printf("Usage of %s:\n", viper.GetString(ViperServiceNameKey))
	pflag.PrintDefaults()
	println()
	exitFunc(0)
}

// ListenForInterrupt listens for an interrupt signal and gracefully stops the svr.
func ListenForInterrupt(cancelFunc context.CancelFunc) {
	go func() {
		sigChannel := make(chan os.Signal, 1)
		signal.Notify(sigChannel, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
		<-sigChannel
		cancelFunc()
		fmt.Print("\rgracefully stopping...\n")
	}()
}

// FailFast panics if err is not nil.
func FailFast(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err.Error())
	}
}
