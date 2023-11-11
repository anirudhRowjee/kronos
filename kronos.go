package main

import (
	"fmt"
	"math"
)

type LogicalClock struct {
	Value      int
	Prev_value int // This field is public only because testing apparatus wants it
}

// Initalize a new logical clock
func New() LogicalClock {
	return LogicalClock{0, 0}
}

// Update the clock based on the new one
func (self *LogicalClock) Update(new *LogicalClock) (int, error) {

	// This is for archival purposes
	self.Prev_value = self.Value

	// TODO assertion to see if self.Value >= self.Prev_value, feels important
	// Saves from bad inits

	// Update the latest timestamp seen to be the latest timestamp
	// This is guaranteed to be safe as self.value and new.value are both ints
	self.Value = int(math.Max(
		float64(self.Value), float64(new.Value),
	))

	// Since we only update whenever we recieve or send a message, we can
	// add a +1 to this, as the send and recieve is also an event
	self.Value += 1

	return self.Value, nil
}

func main() {
	fmt.Println("Hello, world!")
}
