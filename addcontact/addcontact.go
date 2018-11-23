package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func main() {
	// endpoint
	addContactURL := "https://amisend.com/api/contacts/add"

	// authentication
	var x_username string = ""
	var x_apikey string = ""

	// data
	createContact := map[string]string{
		"name":   "",
		"phone":  "",
		"tags":   "",
		"groups": "",
	}

	params, _ := json.Marshal(createContact)

	// request
	request, err := http.NewRequest("POST", addContactURL, bytes.NewBuffer(params))

	if err != nil {
		log.Fatal(err)
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Set("x-api-user", x_username)
	request.Header.Set("x-api-key", x_apikey)
	request.Header.Set("Content-Length", strconv.Itoa(len(params)))

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
