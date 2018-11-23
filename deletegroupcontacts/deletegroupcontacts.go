package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var contactId string = ""
	var deleteGroupContactsURL string = "https://amisend.com/api/contacts/delete/" + contactId + "/groups"

	// authentication
	var x_username string = ""
	var x_apikey string = ""

	// data
	groups := map[string]string{
		"groups": "",
	}

	params, _ := json.Marshal(groups)

	// request
	request, err := http.NewRequest("POST", deleteGroupContactsURL, bytes.NewBuffer(params))

	if err != nil {
		panic(err.Error())
	}

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
