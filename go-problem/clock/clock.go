package clock

import (
	"fmt"
)

// The value of testVersion here must match `targetTestVersion` in the file
// clock_test.go.
const testVersion = 4

const secondInDay = 24 * 60 * 60
const secondInHours = 60 * 60
const secondInMinute = 60

// Clock struct
type Clock struct {
	sec int
}

// New return new Clock
func New(hour, minute int) (o Clock) {
	sec := hour*secondInHours + minute*secondInMinute
	o = Clock{sec}
	return o.Round()
}

// String return Clock in string
func (c Clock) String() string {
	hours := c.sec / secondInHours
	minute := (c.sec % secondInHours) / secondInMinute
	return fmt.Sprintf("%02d:%02d", hours, minute)
}

//Add function
func (c Clock) Add(minutes int) Clock {
	c.sec = c.sec + minutes*secondInMinute
	return c.Round()
}

//Round clock in a date
func (c Clock) Round() Clock {
	for c.sec < 0 {
		c.sec = c.sec + secondInDay
	}
	for c.sec >= secondInDay {
		c.sec = c.sec - secondInDay
	}
	return c
}
