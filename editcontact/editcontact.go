package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var contactId string = "1"
	var endPoint string = "https://amisend.com/api/contacts/edit/" + contactId

	// authentication
	var userName string = ""
	var apiKey string = ""

	// data
	newContactDetails := map[string]string{
		"name": "",
		"tags": "",
	}

	params, _ := json.Marshal(newContactDetails)

	// request
	request, err := http.NewRequest("POST", endPoint, bytes.NewBuffer(params))

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

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err.Error())
	}

	defer response.Body.Close()

	fmt.Println(string(body))
}
