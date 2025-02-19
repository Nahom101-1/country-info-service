package country

import (
	"encoding/json"
	"fmt"
	"github.com/Nahom101-1/country-info-service/internal/models"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func HandleGetRequest(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	countryCode := vars["code"] // get code from {code} URL
	fmt.Println("testtest", countryCode)

	limit := r.URL.Query().Get("limit")
	if limit == "" {
		limit = "10" // default limit
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// URL to invoke //for testing
	url := "https://restcountries.com/v3.1/alpha/NO"

	// create new request
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Errorf("Error in creating request:", err.Error())
	}

	// Setting content type -> effect depends on the service provider
	r.Header.Add("Accept", "application/json")

	client := &http.Client{}
	defer client.CloseIdleConnections()
	//issue request
	res, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()

	//Devode JSON
	decoder := json.NewDecoder(res.Body)
	var test models.CountryInfo
	if err := decoder.Decode(&test); err != nil {
		log.Fatal(err)
	}

	/*	fmt.Println("Raw Response: ", string(body))

		// HTTP Header content
		fmt.Println("Status:", res.Status)
		fmt.Println("Status code:", res.StatusCode)

		fmt.Println("Content type:", res.Header.Get("content-type"))
		fmt.Println("Protocol:", res.Proto)*/
}
