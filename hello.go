package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func rollDice() string {
	die1 := rand.Intn(6) + 1
	die2 := rand.Intn(6) + 1
	total := die1 + die2
	return fmt.Sprintf("%d (%d, %d)", total, die1, die2)
}

func main() {
	fmt.Println(rollDice())
}
