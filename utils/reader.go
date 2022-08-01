package conversions

import (
	"regexp"
	"strconv"
)

var Regex = regexp.MustCompile(`^(?:(?:0|[1-9]\d*)(?:\.\d*)?|\.\d+)$`)

func ReadNumbers(value string) string {
	return Regex.FindString(value)
}

func ReadSingularDecimal(value float64) string {
	float := strconv.FormatFloat(value, 'f', 1, 64)

	return float
}
