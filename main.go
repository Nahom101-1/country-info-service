package main

import (
	"fmt"
	"github.com/Nahom101-1/country-info-service/internal/constants"
	"github.com/Nahom101-1/country-info-service/internal/handlers"
	"github.com/Nahom101-1/country-info-service/internal/handlers/country"
	"github.com/Nahom101-1/country-info-service/internal/utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	port := utils.GetPort()
	//end point handlers
	// TODO: handle dynamic URL
	r := mux.NewRouter()
	//Empty path
	r.HandleFunc(constants.EmptyPath, handlers.EmptyHandler)
	fmt.Println(" Received request:", constants.CountryInfoEndpoint)
	//Country info onnly allow get as method
	r.HandleFunc(constants.CountryInfoEndpoint, country.Handler).Methods(http.MethodGet)

	r.HandleFunc(constants.PopulationEndpoint, handlers.EmptyHandler)
	r.HandleFunc(constants.ServiceStatusEndpoint, handlers.EmptyHandler)

	fmt.Println(" Starting server:, :", port, r)
	//start server
	err := http.ListenAndServe(":"+port, r)

	//handle error when starting server
	if err != nil {
		log.Fatal(err)
	}
}
