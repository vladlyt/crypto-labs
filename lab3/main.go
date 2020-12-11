package main

import (
	"fmt"
	"time"
)

const (
	CASINO_LINK = "http://95.217.177.249/casino"
	LCG         = "Lcg"
	MT          = "Mt"
	IMPROVED_MT = "BetterMt"
	M           = 2 << 31
	a, c        = 1664525, 1013904223
)

func lcgNextValue(x int) int {
	return (x*a + c) % M
}

func lcgCrack() {
	game := NewGame(LCG)
	fmt.Printf("%#v\n", game)
	x := game.MakeABet(1, 33)
	fmt.Println(x)
	for game.Money <= 1000000 || game.Money == 0 {
		x = lcgNextValue(x)
		game.MakeABet(900, x)
	}
	if game.Money != 0 {
		fmt.Println("YEAH WIN")
	} else {
		fmt.Println("LOST")
	}
}

func mtCrack() {
	game := NewGame(MT)

	seed := game.CreationTime().UTC().Sub(
		time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
	).Seconds()
	fmt.Println("SEED:", seed)

	mt := initMT19937()
	mt.Seed(uint32(seed))

	for game.Money <= 1000000 || game.Money == 0 {
		expected := mt.Next()
		game.MakeABet(900, int(expected))
	}
}

func improvedMtCrack() {
	game := NewGame(IMPROVED_MT)

	inputs := make([]uint32, 624)
	for i := 0; i < 624; i++ {
		result := game.MakeABet(1, 33)
		inputs[i] = uint32(result)
	}

	mt := initMTImproved(inputs).MakeRange()

	for game.Money <= 1000000 || game.Money == 0 {
		expected := mt.Next()
		bet := game.Money - 1
		game.MakeABet(int(bet), int(expected))
	}
}

func main() {
	//lcgCrack()
	//mtCrack()
	improvedMtCrack()
}
