package main

import (
	"github.com/iVitaliya/logger-go"
)

func main() {
	var calc_type string

	logger.Debug("Enter the type of calculator you'd like to use (math, bmi, angel): ")
	logger.Scan(&calc_type)
	logger.LogNextLine(2)

	logger.Scan()
}
