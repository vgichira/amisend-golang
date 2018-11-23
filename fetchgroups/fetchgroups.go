package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// endpoint
	var fetchGroupsURL string = "https://amisend.com/api/contacts/groups/fetch"

	// authentication
	var x_username string = "vgichira"
	var x_apikey string = "ami_f8e36721dc387e81ebc6111668da96b9fffee4a4"

	// request
	request, err := http.NewRequest("GET", fetchGroupsURL, nil)

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

	body, _ := ioutil.ReadAll(response.Body)

	defer response.Body.Close()

	fmt.Println(string(body))
}
