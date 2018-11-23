package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// endpoint
	var fetchGroupContactsURL string = "https://amisend.com/api/contacts/groups/fetch"

	// authentication
	var x_username string = "vgichira"
	var x_apikey string = "vgichira"

	// request
	request, err := http.NewRequest("GET", fetchGroupContactsURL, nil)

	if err != nil {
		panic(err.Error())
	}

	// headers
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("x-api-user", x_username)
	request.Header.Set("x-api-key", x_apikey)

	// response
	response, err := http.DefaultClient.Do(request)

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
