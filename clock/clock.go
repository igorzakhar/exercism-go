package clock

import (
	"fmt"
)

type Clock int

func New(hour, minute int) Clock {
	t := (hour*60 + minute) % (60 * 24)
	if t < 0 {
		t += 60 * 24
	}
	return Clock(t)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}

func (c Clock) Add(minutes int) Clock {
	t := minutes + int(c)
	return New(t/60, t%60)
}

func (c Clock) Subtract(minutes int) Clock {
	t := int(c) - minutes
	return New(t/60, t%60)
}
