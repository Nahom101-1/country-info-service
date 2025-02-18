package handlers

import (
	"fmt"
	"github.com/Nahom101-1/country-info-service/internal/constants"
	"net/http"
)

/* EmptyHandler serves as a default handler for invalid root path requests.*/
func EmptyHandler(w http.ResponseWriter, r *http.Request) {

	// Ensure interpretation as HTML by client (browser)

	w.Header().Set("content-type", "text/html: charset=UTF-8")

	// Offer information for redirection to paths
	output := fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head><title>API Navigation</title></head>
		<body>
			<h2>Invalid Request</h2>
			<p>This service does not provide any functionality at the root path.</p>
			<p>Please use valid paths:</p>
			<ul>
				<li><a href="%s">Country Information</a></li>
				<li><a href="%s">Population Data</a></li>
				<li><a href="%s">Service Status</a></li>
			</ul>
		</body>
		</html>`, constants.CountryInfoBasePath, constants.PopulationBasePath, constants.StatusBasePath)

	// Write output to client
	_, err := fmt.Fprintf(w, output)
	if err != nil {
		//deal with error if any
		http.Error(w, "Error when returning output", http.StatusInternalServerError)
		return
	}

}
