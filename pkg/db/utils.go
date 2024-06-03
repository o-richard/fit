package db

import "strings"

// month is lowercased 3/4 characters of the month name
func GetMonthNumber(month string) int {
	month = strings.ToLower(month)
	switch month {
	case "jan":
		return 1
	case "feb":
		return 2
	case "mar":
		return 3
	case "apr":
		return 4
	case "may":
		return 5
	case "june":
		return 6
	case "july":
		return 7
	case "aug":
		return 8
	case "sep":
		return 9
	case "oct":
		return 10
	case "nov":
		return 11
	case "dec":
		return 12
	default:
		return 0
	}
}

// month is a 2-digit integer string. returns the titlecased 3/4 characters of the month name
func GetMonthName(month string) string {
	switch month {
	case "01":
		return "Jan"
	case "02":
		return "Feb"
	case "03":
		return "Mar"
	case "04":
		return "Apr"
	case "05":
		return "May"
	case "06":
		return "June"
	case "07":
		return "July"
	case "08":
		return "Aug"
	case "09":
		return "Sep"
	case "10":
		return "Oct"
	case "11":
		return "Nov"
	case "12":
		return "Dec"
	default:
		return ""
	}
}
