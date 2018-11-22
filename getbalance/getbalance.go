package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	req, err := http.NewRequest("GET", "https://amisend.com/api/topups/balance", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("x-api-user", "")
	q.Add("x-api-key", "")
	q.Add("Content-Type", "application/json")
	req.URL.RawQuery = q.Encode()

	response, err := http.DefaultClient.Do(req)

	fmt.Println(response)
	// Output:
	// http://api.themoviedb.org/3/tv/popular?another_thing=foo+%26+bar&api_key=key_from_environment_or_flag
}
