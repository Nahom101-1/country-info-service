package population

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleGetRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	countryCode := vars["code"]
	startYear := vars["startYear"]
	endYear := vars["endYear"]
	log.Println(" Get Request for:", countryCode, startYear, endYear)

	limit := r.URL.Query().Get("limit")
	if limit == "" {
		limit = "10" // default limit
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
