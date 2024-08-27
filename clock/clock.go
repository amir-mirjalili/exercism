package clock

import "fmt"

type Clock struct {
	Hours   int
	Minutes int
}

func New(h, m int) Clock {
	totalMinutes := (h*60 + m) % 1440
	if totalMinutes < 0 {
		totalMinutes += 1440
	}
	return Clock{totalMinutes / 60, totalMinutes % 60}
}

func (c Clock) Add(m int) Clock {
	return New(c.Hours, c.Minutes+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.Hours, c.Minutes-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hours, c.Minutes)
}
