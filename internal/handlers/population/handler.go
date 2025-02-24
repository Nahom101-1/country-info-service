package population

import (
	"fmt"
	"log"
	"net/http"
)

// Handler handles incoming HTTP requests
func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request:", r.Method, r.URL.Path)
	switch r.Method {
	case http.MethodGet:
		HandleGetRequest(w, r)
	default:
		http.Error(w,
			fmt.Sprintf(`{"error": "REST Method '%s' not supported. Supported methods: '%s'."}`,
				r.Method, http.MethodGet),
			http.StatusMethodNotAllowed)
	}
}
