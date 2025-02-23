package services

import (
	"encoding/json"
	"fmt"
	"github.com/Nahom101-1/country-info-service/internal/constants"
	"github.com/Nahom101-1/country-info-service/internal/models"
	"github.com/Nahom101-1/country-info-service/internal/utils"
	"log"
	"net/url"
)

// GetCountry fetches country info based on country code
func GetCountry(code string) (models.CountryData, error) {
	// Construct API URL
	baseUrl := fmt.Sprintf(constants.CountryURL, code)
	apiURL, err := url.Parse(baseUrl)
	if err != nil {
		log.Println("Error parsing URL:", err)
		return models.CountryData{}, err
	}
	apiURL.RawQuery = "fields=name,capital,continents,population,languages,borders,flags"

	// Log the request URL
	log.Println("Requesting:", apiURL.String())

	// Use SendGetRequest from utils
	response, err := utils.SendGetRequest(apiURL.String())
	if err != nil {
		return models.CountryData{}, err
	}

	// Read response body
	body, err := utils.ReadResponseBody(response)
	if err != nil {
		return models.CountryData{}, err
	}

	// Decode JSON response
	var countryData models.CountryData
	if err := json.Unmarshal(body, &countryData); err != nil {
		log.Println("Error decoding country response JSON:", err)
		return models.CountryData{}, err
	}

	return countryData, nil
}
