package main

import (
	"github.com/Nahom101-1/country-info-service/internal/constants"
	"github.com/Nahom101-1/country-info-service/internal/handlers"
	"github.com/Nahom101-1/country-info-service/utils"
	"log"
	"net/http"
)

func main() {
	port := utils.GetPort()
	http.HandleFunc(constants.EmptyPath, handlers.EmptyHandler)
	http.HandleFunc(constants.CountryInfoEndpoint, handlers.EmptyHandler)
	http.HandleFunc(constants.PopulationEndpoint, handlers.EmptyHandler)
	http.HandleFunc(constants.ServiceStatusEndpoint, handlers.EmptyHandler)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
