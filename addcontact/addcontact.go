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
	endPoint := "https://amisend.com/api/contacts/add"

	userName := ""
	apiKey := ""

	postData := map[string]string{
		"name":   "",
		"phone":  "",
		"tags":   "",
		"groups": "",
	}

	params, _ := json.Marshal(postData)

	request, err := http.NewRequest("POST", endPoint, bytes.NewBuffer(params))

	if err != nil {
		log.Fatal(err)
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