package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
)

func httpRequest(method string, url string, headers map[string]string, payload []byte) (*http.Response, error) {

	var req = &http.Request{}
	var err error

	log.Println("Request URL: " + url)

	if payload != nil {
		req, err = http.NewRequest(method, url, bytes.NewBuffer(payload))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	log.Println("Response Code: " + strconv.Itoa(resp.StatusCode))

	// Throw an error if the response is over 400
	if resp.StatusCode >= 400 {
		err = errors.New("ERROR Status Code is " + strconv.Itoa(resp.StatusCode))
	}

	return resp, err
}

func callSendApi(data MessageRequestData) {

	jsonPayload, _ := json.Marshal(data)

	log.Println("https://graph.facebook.com/v2.6/me/messages?access_token=" + os.Getenv("PAGE_ACCESS_TOKEN"))
	log.Println(string(jsonPayload))

	var headers = map[string]string{
		"Authorization": "Bearer " + os.Getenv("LINE_CHANNEL_ACCESS_TOKEN"),
		"Content-Type":  "application/json",
	}

	resp, err := httpRequest("POST", "https://graph.facebook.com/v2.6/me/messages?access_token="+os.Getenv("PAGE_ACCESS_TOKEN"), headers, jsonPayload)
	if err != nil {
		log.Println("Error sending reply" + err.Error())

		return
	}
	// Close the Body after using. (Find a better way to do this later. It's kind of weird doing it in a different method)
	defer resp.Body.Close()

}
