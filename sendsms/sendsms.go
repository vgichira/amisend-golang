package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
	// endpoint
	endPoint := "https://amisend.com/api/sms/send"

	// authentication

	userName := "vgichira"
	apiKey := "ami_f8e36721dc387e81ebc6111668da96b9fffee4a4"

	// data

	messageData := map[string]string{
		"phoneNumbers": "+254725089232",
		"message":      "Test message",
		"senderId":     "", // leave blank if you do not have a custom sender Id
	}

	params, _ := json.Marshal(messageData)

	request, err := http.NewRequest("POST", endPoint, bytes.NewBuffer(params))

	if err != nil {
		panic(err.Error())
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Set("x-api-user", userName)
	request.Header.Set("x-api-key", apiKey)
	request.Header.Set("Content-Length", strconv.Itoa(len(params)))

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
