package country

import (
	"encoding/json"
	"github.com/Nahom101-1/country-info-service/internal/models"
	"github.com/Nahom101-1/country-info-service/internal/services"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleGetRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	countryCode := vars["code"] // get code from {code} URL
	log.Println("Fetching country", countryCode)

	limit := r.URL.Query().Get("limit")
	if limit == "" {
		limit = "10" // default limit
	}

	countryData, err := services.GetCountry(countryCode)
	if err != nil {
		log.Println("Error fetching country", err)
		http.Error(w, "Faild to fetch country data", http.StatusInternalServerError)
		return
	}

	citiesData, err := services.GetCities(countryData.Name.Common)
	if err != nil {
		log.Println("Error fetching country", err)
		http.Error(w, "Faild to fetch country data", http.StatusInternalServerError)
		return
	}

	response := models.CountryWithCities{
		Name:       countryData.Name.Common,
		Continents: countryData.Continents,
		Population: countryData.Population,
		Languages:  countryData.Languages,
		Borders:    countryData.Borders,
		Flag:       countryData.Flags.PNG,
		Capital:    countryData.Capital,
		Cities:     citiesData.Cities,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("Error encoding response", err)
		http.Error(w, "Faild to fetch country data", http.StatusInternalServerError)
		return
	}
}
