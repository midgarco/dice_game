package models

import (
	"math/rand"
)

// Die model
type Die struct {
	Sides int
}

// NewSixSidedDie quickly creates a six sided die
func NewSixSidedDie() *Die {
	return &Die{Sides: 6}
}

// Roll the die
func (d Die) Roll() int {
	r := rand.New(rand.NewSource(int64(d.Sides)))
	return r.Intn(d.Sides) + 1
}
