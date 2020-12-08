package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/google/uuid"
)

const (
	LINK = "http://95.217.177.249/casino"
)

type Game struct {
	id     string
	amount int
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func createGame() *Game {
	ui, _ := uuid.NewUUID()
	game := Game{
		id:     ui.String(),
		amount: 1000,
	}

	client := http.Client{}
	request, err := http.NewRequest("GET", "http://95.217.177.249/casino/createacc?id="+game.id, nil)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}
	log.Println(result, resp)

	return &game
}

func (g *Game) makeABet(amount int, number int, mode string) int {
	// client := http.Client{}
	// request, err := http.NewRequest("GET", "http://95.217.177.249/casino/play"+mode+"?id="+g.id+"&bet="+amount+"&number="+number, nil)
	return 0
}

func main() {
	mt := initMT19937(0)
	mt.Seed(1303091290)

	fmt.Println(mt.mtToFloat())

	//mt.Seed()
	//game := createGame()

	//for m in range(min_m, max_m)

	// playLcg?id=312312313&bet=2&number=2414241241
	///play{Mode}?id={playerID}&bet={amountOfMoney}&number={theNumberYouBetOn}
}
