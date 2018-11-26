package clock

import (
	"fmt"
)

// Clock is a type representing a clock
type Clock int

// New creates a new Clock with the time given in minutes
func New(hour, minute int) Clock {
	t := (hour*60 + minute) % (60 * 24)
	if t < 0 {
		t += 60 * 24
	}
	return Clock(t)
}

// String implements Stringer interface
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}

// Add adds minutes and return a new Clock instance with the new time
func (c Clock) Add(minutes int) Clock {
	t := minutes + int(c)
	return New(t/60, t%60)
}

// Add subtract minutes and return a new Clock instance with the new time
func (c Clock) Subtract(minutes int) Clock {
	t := int(c) - minutes
	return New(t/60, t%60)
}
