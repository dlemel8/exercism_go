package clock

import (
	"fmt"
	"time"
)

const testVersion = 4

type Clock struct {
	hour, minute int
}

func New(hour, minute int) Clock {
	var c Clock
	return c.Add(hour*60 + minute)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

func (c Clock) Add(minutes int) Clock {
	maxMinutes := int(24 * time.Hour / time.Minute)

	totalMinutes := (c.hour*60 + c.minute + minutes) % maxMinutes
	if totalMinutes < 0 {
		totalMinutes += maxMinutes
	}

	return Clock{
		(totalMinutes / 60) % 24,
		totalMinutes % 60,
	}
}
