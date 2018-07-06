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

	game := &game.Game{}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("How many players? ")
	var playerCount string
	if scanner.Scan() {
		playerCount = scanner.Text()
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	numPlayers, err := strconv.Atoi(playerCount)
	if err != nil {
		panic(err)
	}
	game.Players = make([]*player.Player, numPlayers)

	for i, p := range game.Players {
		fmt.Printf("Player %d Name: ", i+1)
		p = &player.Player{}
		if scanner.Scan() {
			p.Name = scanner.Text()
		}
		game.Players[i] = p
	}

	fmt.Println(game)
}
