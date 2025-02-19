package country

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" Received request:", r.Method, r.URL.Path)
	switch r.Method {
	case http.MethodGet:
		HandleGetRequest(w, r)
	case http.MethodPost:
		//handlePostRequest(w, r)
	default:
		http.Error(w,
			fmt.Sprintf(`{"error": "REST Method '%s' not supported. Supported methods: '%s', '%s'."}`,
				r.Method, http.MethodGet, http.MethodPost),
			http.StatusMethodNotAllowed)
	}

}
