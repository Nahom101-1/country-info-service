package utils

import (
	"log"
	"net/http"
)

// SendGetRequest Sender request to given url and returns response/err
func SendGetRequest(url string) (*http.Response, error) {
	//Make get Request via http method get
	response, err := http.Get(url)
	if err != nil {
		//If faild to make request log error and return nil,err
		log.Println("Error sending GET request:", err)
		return nil, err
	}
	return response, nil //return response and error if any
}
