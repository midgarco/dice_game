package game

import "github.com/midgarco/dice_game/internal/player"

// Game ...
type Game struct {
	Players []*player.Player `json:"players"`
}
