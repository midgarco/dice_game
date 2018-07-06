package die

import (
	"math/rand"
)

// Die model
type Die struct {
	Sides int `json:"side"`
}

// NewSixSidedDie quickly creates a six sided die
func NewSixSidedDie() *Die {
	return &Die{Sides: 6}
}

// Roll the die
func (d Die) Roll() int {
	return rand.Intn(d.Sides) + 1
}
