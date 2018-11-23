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
	// endpoint
	var addToGroupURL string = "https://amisend.com/api/contacts/add/" + contactId + "/groups"

	// authentication

	var x_username string = ""
	var x_apikey string = ""

	// data

	groupData := map[string]string{
		"groups": "Group Test",
	}

	params, _ := json.Marshal(groupData)

	// request

	request, err := http.NewRequest("POST", addToGroupURL, bytes.NewBuffer(params))

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
