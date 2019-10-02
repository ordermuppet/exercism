package clock

import "fmt"

const minutesInDay = 1440

// Clock handles time but not dates.
type Clock struct {
	minute int
}

// New creates a new clock.
func New(hour, minute int) Clock {
	a := minute + hour*60
	n := minutesInDay
	return Clock{minute: ((a % n) + n) % n}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minute/60, c.minute%60)
}

// Add returns a new clock with more minutes.
func (c Clock) Add(minutes int) Clock {
	return New(0, c.minute+minutes)
}

// Subtract returns a new clock with fewer minutes.
func (c Clock) Subtract(minutes int) Clock {
	return New(0, c.minute-minutes)
}
