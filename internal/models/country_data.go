package models

// CountryData struct for storing country information
type CountryData struct {
	Name struct {
		Common string `json:"common"` // Common name of the country
	} `json:"name"`

	Continents []string          `json:"continents"` // Continents the country belongs to
	Population int64             `json:"population"` // Population of the country
	Languages  map[string]string `json:"languages"`  // Languages spoken with their abbreviations
	Borders    []string          `json:"borders"`    // Borders of country (list of country codes)

	Flags struct {
		PNG string `json:"png"` // Flags URL (PNG format)
	} `json:"flags"`

	Capital []string `json:"capital"` // Capital cities of the country
}

/* Example API Response:
{
    "name": {"common": "Norway"},
    "continents": ["Europe"],
    "population": 4700000,
    "languages": {"nno": "Norwegian Nynorsk", "nob": "Norwegian Bokmål", "smi": "Sami"},
    "borders": ["FIN", "SWE", "RUS"],
    "flags": {
        "png": "https://flagcdn.com/w320/no.png",
    },
    "capital": ["Oslo"],
}
*/
