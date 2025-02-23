package utils

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// SendPostRequest Sender post request given url and payload and returns reponse/err
func SendPostRequest(url string, payload interface{}) (*http.Response, error) {
	// Prepare JSON payload
	jsonPayLoad, err := json.Marshal(payload) // convert payload to JSON
	if err != nil {
		//If payload is invalid log error and return
		log.Println("Error marshalling JSON payload:", err)
		return nil, err
	}
	//Create a new HTTP post request
	//bytes.NewBuffer(jsonPayload) wraps the JSON data into a readable format
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonPayLoad))
	if err != nil {
		//If http request faild log error and return
		log.Println("Error creating request:", err)
		return nil, err
	}

	//Make sure server knows it is receiving JSON
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}        //create client to handle request
	response, err := client.Do(req) //Send request to server and wait for response
	if err != nil {
		// If server faild to reespond log error and return
		log.Println("Error posting request:", err)
		return nil, err
	}

	return response, nil
}
