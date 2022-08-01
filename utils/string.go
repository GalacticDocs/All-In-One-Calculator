package conversions

import (
	"strings"

	"github.com/GalacticDocs/Conversions"
)

func FormatString(str string, values []string) string {
	var (
		value string = str
		format (func(int) string) = func(i int) string { return "{" + conversions.Stringify().IntToString(i) + "}" }
	)

	for x := 0; x < len(values); x++ {
		strings.Replace(value, format(x), values[x], -1)
	}

	return value
}