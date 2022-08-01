package calculators

import (
	"github.com/GalacticDocs/Conversions"
	"github.com/GalacticDocs/Conversions/converter"
	"github.com/GalacticDocs/Conversions/reader"
	"github.com/iVitaliya/logger-go"
)

func BMI() {
	var (
		height string
		weight string
	)

	logger.Debug("Enter your height in cm: ")
	logger.Scan(&height)
	logger.LogNextLine(1)

	logger.Debug("Enter you weight in kg: ")
	logger.Scan(&weight)

	var (
		fWeight float64                    = float64(converter.StringToInt(weight))
		fHeight float64                    = float64(converter.StringToInt(height))
		builder *conversions.IStringBuilder = conversions.StringBuilder("")

		BMI       float64 = (fWeight / fHeight / fHeight) * 10000
		FORMATTED string  = reader.ReadSingularDecimal(BMI)
	)

	builder.Append("Your Body Mass Index is: ")
	builder.Append(FORMATTED)

	logger.LogNextLine(2)
	logger.Info(builder.Build())
	logger.LogNextLine(1)

	var (
		FORMATTED_FLOAT float64 = converter.StringToFloat64(FORMATTED)

		UNDERWEIGHT_VALUE float64 = converter.StringToFloat64("18.5")
		HEALTHY_VALUE     float64 = converter.StringToFloat64("24.9")
		OVERWEIGHT_VALUE  float64 = converter.StringToFloat64("29.9")
	)

	if FORMATTED_FLOAT <= UNDERWEIGHT_VALUE {
		logger.Info("Result: You are underweight.")
	} else if FORMATTED_FLOAT <= HEALTHY_VALUE {
		logger.Info("Result: Awesome! You are healthy.")
	} else if FORMATTED_FLOAT <= OVERWEIGHT_VALUE {
		logger.Info("Result: You are overweight.")
	} else {
		logger.Info("Result: You are obese.")
	}
}
