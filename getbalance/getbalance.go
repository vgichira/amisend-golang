package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	// endpoint
	var endPoint string = "https://amisend.com/api/topups/balance"

	var userName string = ""
	var apiKey string = ""

	// request
	req, err := http.NewRequest("GET", endPoint, nil)

	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	// headers
	req.Header.Add("x-api-user", userName)
	req.Header.Add("x-api-key", apiKey)
	req.Header.Add("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err.Error())
	}

	defer response.Body.Close()

	fmt.Println(string(body))
}
