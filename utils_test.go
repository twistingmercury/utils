package utils_test

import (
	"bytes"
	"context"
	"os"
	"syscall"
	"testing"
	"time"

	"github.com/twistingmercury/utils"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestShowVersion(t *testing.T) {
	// Set up test data
	pflag.Parse()
	viper.Set(utils.ViperShowVersionKey, true)
	viper.Set(utils.ViperServiceNameKey, "test-service")
	viper.Set(utils.ViperBuildVersionKey, "1.0.0")
	viper.Set(utils.ViperBuildDateKey, "2023-06-09")
	viper.Set(utils.ViperCommitHashKey, "abcdef")

	// Redirect stdout to capture the output
	stdout := os.Stdout
	defer func() { os.Stdout = stdout }()
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Override the exit function to avoid exiting during tests
	defer func() { utils.SetExitFunc(os.Exit) }()
	var exitCode int
	exitFunc := func(code int) { exitCode = code }
	utils.SetExitFunc(exitFunc)

	// Call the function
	utils.ShowVersion()

	// Capture the output
	w.Close()
	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	// Assert the output
	expectedOutput := "test-service\nversion: 1.0.0; build date: 2023-06-09; commit: abcdef\n"
	assert.Equal(t, expectedOutput, output)
	assert.Equal(t, 0, exitCode)
}

func TestShowHelp(t *testing.T) {
	// Set up test data
	pflag.Parse()
	viper.Set(utils.ViperShowHelpKey, true)
	viper.Set(utils.ViperServiceNameKey, "test-service")
	viper.Set(utils.ViperBuildVersionKey, "1.0.0")
	viper.Set(utils.ViperBuildDateKey, "2023-06-09")
	viper.Set(utils.ViperCommitHashKey, "abcdef")

	// Redirect stdout to capture the output
	stdout := os.Stdout
	defer func() { os.Stdout = stdout }()
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Override the exit function to avoid exiting during tests
	defer func() { utils.SetExitFunc(os.Exit) }()
	var exitCode int
	exitFunc := func(code int) { exitCode = code }
	utils.SetExitFunc(exitFunc)

	// Call the function
	utils.ShowHelp()

	// Capture the output
	w.Close()
	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	// Assert the output
	assert.Contains(t, output, "test-service\nversion: 1.0.0; build date: 2023-06-09; commit: abcdef")
	assert.Contains(t, output, "Usage of")
	assert.Equal(t, 0, exitCode)
}

func TestListenForInterrupt(t *testing.T) {
	// Create a context and cancel function
	ctx, cancel := context.WithCancel(context.Background())

	// Call the ListenForInterrupt function
	go utils.ListenForInterrupt(cancel)

	time.Sleep(250 * time.Millisecond)

	// Send an interrupt signal
	syscall.Kill(syscall.Getpid(), syscall.SIGINT)

	// Wait for the context to be canceled
	<-ctx.Done()

	// Assert that the context is canceled
	assert.Equal(t, context.Canceled, ctx.Err())
}

func TestFailFast(t *testing.T) {
	// Assert that FailFast panics when an error is provided
	assert.Panics(t, func() {
		utils.FailFast(assert.AnError, "test error")
	}, "FailFast should panic when an error is provided")

	// Assert that FailFast doesn't panic when no error is provided
	assert.NotPanics(t, func() {
		utils.FailFast(nil, "test success")
	}, "FailFast should not panic when no error is provided")
}
