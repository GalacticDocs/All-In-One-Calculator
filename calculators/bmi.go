package calculators

import (
	"fmt"
	"regexp"
	"strconv"

	CLRS "github.com/iVitaliya/colors-go"
)

func readOnlyNumbers(x string) string {
	re := regexp.MustCompile(`^(?:(?:0|[1-9]\d*)(?:\.\d*)?|\.\d+)$`)
	return re.FindString(x)
}

func turnIntoString(x string) float64 {
	i, err := strconv.ParseFloat(readOnlyNumbers(x), 32)
	if err != nil {
		fmt.Println(err.Error())
	}

	return i
}

func readSingularDecimal(value float64) string {
	float := strconv.FormatFloat(value, 'f', 1, 64)

	return float
}

func BMI() {
	
}