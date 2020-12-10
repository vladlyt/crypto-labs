package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Game struct {
	Id           string    `json:"id"`
	Money        int64     `json:"money"`
	DeletionTime time.Time `json:"deletionTime"`
	mode string
}

type PlayResponse struct {
	Message    string `json:"message"`
	Game       Game   `json:"account"`
	RealNumber int    `json:"realNumber"`
}

func NewGame(mode string) *Game {
	ui, _ := uuid.NewUUID()
	fmt.Println(ui.String())
	game := Game{
		mode: mode,
	}

	client := http.Client{}
	request, err := http.NewRequest(
		"GET", fmt.Sprintf("%s/createacc?id=%s", CASINO_LINK, ui.String()),
		nil,
	)
	if err != nil {
		log.Fatalln(err)
	}

	r, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Body.Close()

	err = json.NewDecoder(r.Body).Decode(&game)
	if err != nil {
		log.Fatalln(err)
	}

	return &game
}

func (g *Game) MakeABet(bet int, number int) int {
	client := http.Client{}
	req, err := http.NewRequest(
		"GET", fmt.Sprintf("%s/play%s", CASINO_LINK, g.mode),
		nil,
	)

	q := req.URL.Query()
	q.Add("bet", strconv.Itoa(bet))
	q.Add("number", strconv.Itoa(number))
	q.Add("id", g.Id)

	req.URL.RawQuery = q.Encode()
	fmt.Println(req.URL.String())
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	playResp := PlayResponse{}
	if res.StatusCode != 200 {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Not 200: ", string(bodyBytes))
	} else {
		err = json.NewDecoder(res.Body).Decode(&playResp)
		fmt.Printf("%#v\n", playResp)
		fmt.Println(playResp.Message, playResp.Game.Money, playResp.RealNumber)
	}
	g.Money = playResp.Game.Money

	return playResp.RealNumber
}

func (g *Game) CreationTime() time.Time {
	return g.DeletionTime.Add(-time.Hour)
}