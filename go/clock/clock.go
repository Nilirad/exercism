package clock

import "fmt"

type Clock struct {
	minutes int
}

const minutesPerDay = 60 * 24

func New(h, m int) Clock {
	minutesModulo := ((h*60+m)%minutesPerDay + minutesPerDay) % minutesPerDay
	return Clock{minutes: minutesModulo}
}

func (c Clock) Add(m int) Clock {
	minutes := c.minutes + m
	return New(0, minutes)
}

func (c Clock) Subtract(m int) Clock {
	minutes := c.minutes - m
	return New(0, minutes)
}

func (c Clock) String() string {
	hours := c.minutes / 60
	minutes := c.minutes - 60*hours

	return fmt.Sprintf("%02d:%02d", hours, minutes)
}
