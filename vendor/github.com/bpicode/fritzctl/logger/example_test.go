package logger_test

import (
	"flag"

	"github.com/bpicode/fritzctl/logger"
)

// Logging can be turned off manually.
func ExampleLevel_Set() {
	l := &logger.Level{}
	l.Set("none")
	logger.Warn("log statement")
	// Output:
}

// Level can be used with the flag package.
func ExampleLevel() {
	l := logger.Level{}
	flag.Var(&l, "log", "logging verbosity, e.g. 'info'")
	flag.Parse()
	// Output:
}

// Log on "info" level.
func ExampleInfo() {
	logger.Info("informational message")
	// Output:
}

// Log on "warn" level.
func ExampleWarn() {
	logger.Warn("a warning")
	// Output:
}

// Log on "error" level.
func ExampleError() {
	logger.Error("an error occurred")
	// Output:
}

// Log on "debug" level.
func ExampleDebug() {
	logger.Debug("debug message")
	// Output:
}

// Log on "success" level.
func ExampleSuccess() {
	logger.Success("successfully reached a milestone in my program flow")
	// Output:
}
