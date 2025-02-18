package country

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
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
