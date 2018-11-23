package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var endPoint string = "https://amisend.com/api/contacts/fetch"

	// authentication

	var userName string = ""
	var apiKey string = ""

	// request

	request, err := http.NewRequest("GET", endPoint, nil)

	if err != nil {
		panic(err.Error())
	}

	// headers

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("x-api-user", userName)
	request.Header.Set("x-api-key", apiKey)

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		panic(err.Error())
	}

	body, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(body))

}
