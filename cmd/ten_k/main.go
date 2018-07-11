package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/midgarco/dice_game/internal/game"
	"github.com/midgarco/dice_game/internal/player"
)

func main() {
	rand.Seed(time.Now().Unix())

	game := &game.Game{MaxScore: 10000, OpenScore: 650}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("How many players? \n")
	var playerCount string
	if scanner.Scan() {
		playerCount = scanner.Text()
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	// _ = scanner
	// playerCount = "2"

	numPlayers, err := strconv.Atoi(playerCount)
	if err != nil {
		panic(err)
	}
	if numPlayers < 2 {
		fmt.Println("Game requires at least 2 players")
		os.Exit(1)
	}
	game.Players = make([]*player.Player, numPlayers)

	for i := range game.Players {
		game.Players[i] = player.NewPlayer(i + 1)
	}

	// start the game
	game.Start()
}
