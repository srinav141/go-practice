/*
Package leap implements a simple function
to check whether a given year in leap or not

*/
package leap

// IsLeapYear evaluates a given year and return true if it is a leap year
func IsLeapYear(year int) bool {
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		return true
	}
	return false
}
