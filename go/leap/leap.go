// Utilities for calculating leap years.
package leap

// Checks whether `year` is a leap year.
func IsLeapYear(year int) bool {
	period4 := year%4 == 0
	period100 := year%100 == 0
	period400 := year%400 == 0

	return period400 || (period4 && !period100)
}
