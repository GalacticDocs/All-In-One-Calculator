package conversions

import (
	"strconv"
)

// func formatable(number int) string { return "{" + string(rune(number)) + "}" }
// func FormatString(str string, value []string) string {
// 	for i := 0; i < len(value); i++ {
// 		strings.Replace(str, formatable(i), value[i], 0)
// 	}

// 	return str
// }

func IntToString(value int) string {
	return string(rune(value))
}

func Float32ToString(value float32) string {
	return strconv.FormatFloat(float64(value), 'f', 1, 32)
}

func Float64ToString(value float64) string {
	return strconv.FormatFloat(value, 'f', 1, 64)
}

func StringToFloat32(value string) float32 {
	val, err := strconv.Atoi(value)
	if err != nil {

	}

	return float32(val)
}

func StringToInt(value string) int {
	val, err := strconv.Atoi(value)
	if err != nil {

	}

	return val
}
