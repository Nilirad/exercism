package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	time, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		fmt.Printf("Parse error: %s\n", err)
	}

	return time
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	schedule, err := time.Parse("January 2, 2006 15:04:05", date)
	if err != nil {
		fmt.Printf("Parse error: %s\n", err)
	}

	return time.Now().After(schedule)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	schedule, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if err != nil {
		fmt.Printf("Parse error: %s\n", err)
	}

	return schedule.Hour() >= 12 && schedule.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	datetime := Schedule(date)
	return datetime.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	year := time.Now().Year()
	return time.Date(year, time.September, 15, 0, 0, 0, 0, time.UTC)
}
