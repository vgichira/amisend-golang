package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var groupId string = ""
	// endpoint
	var editGroupURL string = "https://amisend.com/api/contacts/groups/edit/" + groupId

	// authentication
	var x_username string = ""
	var x_apikey string = ""

	// data
	groupData := map[string]string{
		"name": "",
	}

	params, _ := json.Marshal(groupData)

	request, err := http.NewRequest("POST", editGroupURL, bytes.NewBuffer(params))

	if err != nil {
		panic(err.Error())
	}

	// headers
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("x-api-user", x_username)
	request.Header.Set("x-api-key", x_apikey)

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
