package country

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func HandleGetRequest(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	countryCode := vars["code"] // get code from {code} URL
	fmt.Println("testtest")

	limit := r.URL.Query().Get("limit")
	if limit == "" {
		limit = "10" // default limit
	}
	// test
	fmt.Fprintf(w, "Country Code: %s\n", countryCode)
	fmt.Fprintf(w, "Limit: %s\n", limit)

}
