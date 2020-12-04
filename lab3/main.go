package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func main() {

	ui, _ := uuid.NewUUID()
	fmt.Println(ui.String())
	MakeRequest("http://95.217.177.249/casino/createacc?id=7bd9912c-367e-11eb-a5f7-acde48001122")
	//for m in range(min_m, max_m)

	// playLcg?id=312312313&bet=2&number=2414241241
	///play{Mode}?id={playerID}&bet={amountOfMoney}&number={theNumberYouBetOn}
}

func MakeRequest(url string) {

	client := http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	log.Println(result)
}
