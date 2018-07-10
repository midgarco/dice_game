package game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/midgarco/dice_game/internal/die"
	"github.com/midgarco/dice_game/internal/player"
)

// Game ...
type Game struct {
	Players       []*player.Player `json:"players"`
	Dice          []*die.Die       `json:"dice"`
	CurrentPlayer *player.Player   `json:"current_player"`
}

// Start the game
func (g *Game) Start() {
	g.Dice = make([]*die.Die, 5)
	for d := range g.Dice {
		g.Dice[d] = die.NewSixSidedDie()
	}
	g.NextTurn()
}

// NextTurn ...
func (g *Game) NextTurn() {
	g.CurrentPlayer, g.Players = g.Players[0], g.Players[1:]

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Player %s [%d], it's your turn: [R]oll\n", g.CurrentPlayer.Name, g.CurrentPlayer.Score)
	if scanner.Scan() {
		if strings.ToLower(scanner.Text()) != "r" {
			g.EndTurn()
			return
		}
	}

	g.CurrentPlayer.Turn.New()
	for {
		diceRoll := g.CurrentPlayer.Turn.Roll(g.Dice[:g.CurrentPlayer.Turn.RemainingDie])
		for i := range diceRoll {
			fmt.Printf("[%d]:%d ", i, diceRoll[i].Value)
		}
		runningTally := die.Tally(diceRoll)
		fmt.Printf(">> %d\n", runningTally)

		if runningTally == 0 {
			break // turn is over!
		}

		g.Bank(diceRoll)

		// prompt to save their score
		fmt.Println("remaining die", g.CurrentPlayer.Turn.RemainingDie)
		if (g.CurrentPlayer.Score >= 650 || runningTally >= 650 || g.CurrentPlayer.Turn.Banked >= 650) && g.CurrentPlayer.Turn.RemainingDie > 0 {
			fmt.Println("Do you want to [S]ave?")
			if scanner.Scan() {
				if strings.ToLower(scanner.Text()) == "s" {
					g.CurrentPlayer.Save()
					break
				}
			}
		}
	}

	fmt.Printf("Player %s, your score is: %d\n", g.CurrentPlayer.Name, g.CurrentPlayer.Score)

	g.EndTurn()
	return
}

// Bank ...
func (g *Game) Bank(dice []*die.Die) {
	fmt.Println("Select which die to bank: ")
	var bank string
	fmt.Scanln(&bank)

	var banked []*die.Die
	for r := range bank {
		if idx, err := strconv.Atoi(string(bank[r])); err == nil {
			banked = append(banked, dice[idx])
		}
	}
	err := g.CurrentPlayer.Turn.Bank(banked)
	if err != nil {
		return
	}
	fmt.Println("Banked score: ", g.CurrentPlayer.Turn.Banked)
}

// EndTurn ...
func (g *Game) EndTurn() {
	g.Players = append(g.Players, g.CurrentPlayer)
	g.CurrentPlayer = &player.Player{}
	g.NextTurn()
}
