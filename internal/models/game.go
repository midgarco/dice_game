package models

// Game ...
type Game struct {
	Players []*Player `json:"players"`
}
