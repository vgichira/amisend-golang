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
	endPoint := "https://amisend.com/api/contacts/groups/add"

	// authentication
	userName := ""
	apiKey := ""

	// data
	groupData := map[string]string{
		"name": "",
		"tags": "",
	}

	params, _ := json.Marshal(groupData)

	request, err := http.NewRequest("POST", endPoint, bytes.NewBuffer(params))

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
