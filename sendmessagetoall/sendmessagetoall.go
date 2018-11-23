package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// endpoint
	var sendMessageURL string = "https://amisend.com/api/sms/send/all"

	// authentication
	var x_username string = ""
	var x_apikey string = ""

	// data
	createMessage := map[string]string{
		"message":  "",
		"senderId": "",
		"refId":    "",
	}

	params, err := json.Marshal(createMessage)

	if err != nil {
		panic(err.Error())
	}

	// request
	request, err := http.NewRequest("POST", sendMessageURL, bytes.NewBuffer(params))

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
