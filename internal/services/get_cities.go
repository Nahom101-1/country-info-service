package services

import (
	"encoding/json"
	"github.com/Nahom101-1/country-info-service/internal/constants"
	"github.com/Nahom101-1/country-info-service/internal/models"
	"github.com/Nahom101-1/country-info-service/internal/utils"
	"log"
)

// GetCities fetches city data for a given country code
func GetCities(name string) (models.CitiesData, error) {
	apiURL := constants.CitiesURL // URL from constants

	// Prepare JSON payload
	payload := map[string]string{"country": name}

	// Use SendPostRequest to send a POST request
	response, err := utils.SendPostRequest(apiURL, payload)
	if err != nil {
		log.Println("Error sending POST request:", err)
		return models.CitiesData{}, err
	}

	// Read response body
	body, err := utils.ReadResponseBody(response)
	if err != nil {
		log.Println("Error reading response body:", err)
		return models.CitiesData{}, err
	}

	// Decode JSON response
	var cities models.CitiesData
	if err := json.Unmarshal(body, &cities); err != nil {
		log.Println("Error decoding city response JSON:", err)
		return models.CitiesData{}, err
	}

	return cities, nil
}
