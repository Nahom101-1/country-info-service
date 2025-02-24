package main

import (
	"fmt"
	"github.com/Nahom101-1/country-info-service/internal/constants"
	"github.com/Nahom101-1/country-info-service/internal/handlers"
	"github.com/Nahom101-1/country-info-service/internal/handlers/country"
	"github.com/Nahom101-1/country-info-service/internal/handlers/population"
	"github.com/Nahom101-1/country-info-service/internal/utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	port := utils.GetPort()
	//end point handlers
	r := mux.NewRouter()
	//Empty path
	r.HandleFunc(constants.EmptyPath, handlers.EmptyHandler)
	//Country info onnly allow get as method
	r.HandleFunc(constants.CountryInfoEndpoint, country.Handler).Methods(http.MethodGet)
	// TODO: ... these functions PopulationEndpoint and ServiceStatusEndpoint
	r.HandleFunc(constants.PopulationEndpoint, population.Handler).Methods(http.MethodGet)
	/*		r.HandleFunc(constants.ServiceStatusEndpoint, handlers.EmptyHandler)
	 */
	fmt.Println(" Starting server:, :", port, r)
	//start serverd
	err := http.ListenAndServe(":"+port, r)

	//handle error when starting server
	if err != nil {
		log.Fatal(err)
	}
}
