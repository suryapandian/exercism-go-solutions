package clock

import "fmt"

const (
	hrsPerDay  = 24
	minsPerHr  = 60
	minsPerDay = 1440
)

// Clock contains hour and minute
type Clock struct {
	hour, min int
}

// New Clock
func New(hour, minute int) Clock {
	for minute >= minsPerHr {
		hour++
		minute -= minsPerHr
	}
	for minute < 0 {
		hour--
		minute += minsPerHr
	}
	return Clock{fixHour(hour), minute}
}

func fixHour(h int) int {
	h = h % hrsPerDay
	if h < 0 {
		h += hrsPerDay
	}
	return h
}

// String of a Clock hh:mm
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.min)
}

// Add minutes to a Clock
func (c Clock) Add(minutes int) Clock {
	return New(c.hour, c.min+minutes)
}

// Subtract minutes from a Clock
func (c Clock) Subtract(minutes int) Clock {
	return New(c.hour, c.min-minutes)
}