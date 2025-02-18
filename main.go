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
	//end point handlers
	// TODO: handle dynamic URL
	http.HandleFunc(constants.EmptyPath, handlers.EmptyHandler)
	http.HandleFunc(constants.CountryInfoEndpoint, handlers.EmptyHandler)
	http.HandleFunc(constants.PopulationEndpoint, handlers.EmptyHandler)
	http.HandleFunc(constants.ServiceStatusEndpoint, handlers.EmptyHandler)

	//start server
	err := http.ListenAndServe(":"+port, nil)

	//handle error when starting server
	if err != nil {
		log.Fatal(err)
	}
}
