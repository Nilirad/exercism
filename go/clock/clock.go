package clock

import "fmt"

type Clock struct {
	h int
	m int
}

const minutesPerDay = 60 * 24

func New(h, m int) Clock {
	minutesTotal := h*60 + m
	minutesInDay := (minutesTotal%minutesPerDay + minutesPerDay) % minutesPerDay

	hours := minutesInDay / 60
	minutes := minutesInDay - 60*hours
	return Clock{h: hours, m: minutes}
}

func (c Clock) Add(m int) Clock {
	minutesTotal := c.h*60 + c.m + m
	return New(0, minutesTotal)
}

func (c Clock) Subtract(m int) Clock {
	minutesTotal := c.h*60 + c.m - m
	return New(0, minutesTotal)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}
