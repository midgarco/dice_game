package game

import (
	"github.com/midgarco/dice_game/internal/die"
	"github.com/midgarco/dice_game/internal/player"
)

// Game ...
type Game struct {
	Players []*player.Player `json:"players"`
	Dice    []*die.Die       `json:"dice"`
}

// Start the game
func (g *Game) Start() {
	g.Dice = make([]*die.Die, 5)
	for d := range g.Dice {
		die := &die.Die{}
		die.Sides = 6
		g.Dice[d] = die
	}
}
