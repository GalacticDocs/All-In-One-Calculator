package calculators

import (
	"time"

	conversions "github.com/GalacticDocs/Conversions"
)

func stringAppend(original string, appendable string) string {
	newString := original + appendable

	return newString
}

// =================================
// ========== [[ TIME ]] ==========
// =================================

func dayFormatter(day int) string {
	var result string

	if day == 1 || day == 21 || day == 31 {
		result = conversions.Stringify().IntToString(day) + "st"
	} else if day == 2 || day == 22 {
		result = conversions.Stringify().IntToString(day) + "nd"
	} else if day == 3 || day == 23 {
		result = conversions.Stringify().IntToString(day) + "rd"
	} else {
		result = conversions.Stringify().IntToString(day) + "th"
	}

	return result
}

func monthFormatter(month time.Month) string {
	var result string

	switch month {
	case time.January:
		result = "January"

	case time.February:
		result = "February"

	case time.March:
		result = "March"

	case time.April:
		result = "April"

	case time.May:
		result = "May"

	case time.June:
		result = "June"

	case time.July:
		result = "July"

	case time.August:
		result = "August"

	case time.September:
		result = "September"

	case time.October:
		result = "October"

	case time.November:
		result = "November"

	case time.December:
		result = "December"
	}

	return result
}

func Date() string {
	var (
		str string = ""
		now        = time.Now()
	)
	var (
		day   = dayFormatter(now.Day())
		month = monthFormatter(now.Month())
		year  = conversions.Stringify().IntToString(now.Year())
	)

	str = stringAppend(str, day)
	str = stringAppend(str, " of ")
	str = stringAppend(str, month)
	str = stringAppend(str, year)

	return str
}
