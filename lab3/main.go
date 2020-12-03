package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	accID()
	//MakeRequest("http://95.217.177.249/casino/createacc?id=${playerid}")
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
