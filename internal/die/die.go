package die

import (
	"math/rand"
)

// Die model
type Die struct {
	Sides int `json:"sides"`
	Value int `json:"value"`
}

// NewSixSidedDie quickly creates a six sided die
func NewSixSidedDie() *Die {
	return &Die{Sides: 6}
}

// Roll the die
func (d *Die) Roll() {
	d.Value = rand.Intn(d.Sides) + 1
}
