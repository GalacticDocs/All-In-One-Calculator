package main

import (
	"github.com/iVitaliya/logger-go"
)

func main() {
	var Logger = logger.ConfigureLogger(logger.LoggerConfig{
		Severity: logger.DebugState,
	})

	Logger.Log("Testy test test")
}
