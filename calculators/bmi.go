package calculators

import (
	conversions "github.com/GalacticDocs/Conversions"
	"github.com/GalacticDocs/Conversions/converter"
	"github.com/GalacticDocs/Conversions/reader"
	log "github.com/iVitaliya/logger-go"
)

var Logger = log.ConfigureLogger(log.LoggerConfig{
	Severity: log.InfoState,
})

type IBMI struct {
	Height float64
	Weight float64
	BMI    float64
}

func BMI() *IBMI {
	var (
		height string
		weight string
	)

	Logger.Log("Enter your height in cm: ")
	log.Scan(&height)
	log.LogNextLine(1)

	Logger.Log("Enter you weight in kg: ")
	log.Scan(&weight)

	var (
		fWeight float64                     = float64(converter.StringToInt(weight))
		fHeight float64                     = float64(converter.StringToInt(height))
		builder *conversions.IStringBuilder = conversions.StringBuilder("")

		BMI       float64 = (fWeight / fHeight / fHeight) * 10000
		FORMATTED string  = reader.ReadSingularDecimal(BMI)
	)

	builder.Append("Your Body Mass Index is: ")
	builder.Append(FORMATTED)

	log.LogNextLine(2)
	Logger.Log(builder.Build())
	log.LogNextLine(1)

	var (
		FORMATTED_FLOAT float64 = converter.StringToFloat64(FORMATTED)

		UNDERWEIGHT_VALUE float64 = converter.StringToFloat64("18.5")
		HEALTHY_VALUE     float64 = converter.StringToFloat64("24.9")
		OVERWEIGHT_VALUE  float64 = converter.StringToFloat64("29.9")
	)

	if FORMATTED_FLOAT <= UNDERWEIGHT_VALUE {
		Logger.Log("Result: You are underweight.")
	} else if FORMATTED_FLOAT <= HEALTHY_VALUE {
		Logger.Log("Result: Awesome! You are healthy.")
	} else if FORMATTED_FLOAT <= OVERWEIGHT_VALUE {
		Logger.Log("Result: You are overweight.")
	} else {
		Logger.Log("Result: You are obese.")
	}

	return &IBMI{
		Height: fHeight,
		Weight: fWeight,
		BMI:    FORMATTED_FLOAT,
	}
}
