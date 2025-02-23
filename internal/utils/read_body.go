package utils

import (
	"io"
	"log"
	"net/http"
)

// ReadResponseBody reads and returns the body of an HTTP response
func ReadResponseBody(response *http.Response) ([]byte, error) {
	defer response.Body.Close()            // close response body after reading
	body, err := io.ReadAll(response.Body) //Read the response into body
	if err != nil {
		//If failed to read log error and return nil
		log.Println("Error reading response body:", err)
		return nil, err
	}
	return body, nil // Return body or error if any
}
