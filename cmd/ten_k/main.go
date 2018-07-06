package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/midgarco/dice_game/internal/models"
)

func main() {
	rand.Seed(time.Now().Unix())

	d1 := models.NewSixSidedDie()
	d2 := models.NewSixSidedDie()

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println(d1.Roll())
	}()
	go func() {
		defer wg.Done()
		fmt.Println(d2.Roll())
	}()

	wg.Wait()
}
