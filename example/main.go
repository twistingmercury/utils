package main

import (
	"log"
	"net/http"

	"context"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/twistingmercury/utils"
)

func main() {
	Initialize()
	ctx, cancelFunc := context.WithCancel(context.Background())
	utils.ListenForInterrupt(cancelFunc)
	utils.ShowHelp()
	utils.ShowVersion()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	go http.ListenAndServe(":8080", mux)

	log.Println("server started")
	log.Println("press ctrl+c to stop the server")
	<-ctx.Done()
}

// The following three are intended to be set by the build process using ldflags
// I typically defined these in a package `conf`, not in the main package
var (
	buildVersion = "1.0.0"
	buildDate    = "2021-09-01T00:00:00Z"
	commitHash   = "22ae825"
)

// I typically defined these in a package `conf`, not in the main package
var (
	_ = pflag.Bool(utils.ViperShowHelpKey, false, "show version information")
	_ = pflag.Bool(utils.ViperShowVersionKey, false, "show help information")
)

// I typically defined these in a package `conf`, not in the main package
func Initialize() {
	pflag.Parse()
	viper.BindPFlag(utils.ViperShowHelpKey, pflag.Lookup(utils.ViperShowHelpKey))
	viper.BindPFlag(utils.ViperShowVersionKey, pflag.Lookup(utils.ViperShowVersionKey))
	viper.Set(utils.ViperServiceNameKey, "utils-example")

	// Set the default values for the viper keys
	// The following three are intended to be set by the build process using ldflags
	viper.Set(utils.ViperBuildVersionKey, buildVersion)
	viper.Set(utils.ViperBuildDateKey, buildDate)
	viper.Set(utils.ViperCommitHashKey, commitHash)
}
