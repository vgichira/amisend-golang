package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var groupId string = ""
	// endpoint
	var deleteGroupURL string = "https://amisend.com/api/contacts/groups/delete/" + groupId

	// authentication

	var x_username string = ""
	var x_apikey string = ""

	// request
	request, err := http.NewRequest("POST", deleteGroupURL, nil)

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

	fmt.Println(string(body))

}
