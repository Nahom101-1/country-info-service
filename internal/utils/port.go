package utils

import (
	"log"
	"os"
)

// GetPort retrieves the port number from environment variables or defaults to "8080".
func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT not set, using default port 8080")
		port = "8080"
	}
	return port
}
