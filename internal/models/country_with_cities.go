package models

// CountryWithCities struct for storing both country and city data
type CountryWithCities struct {
	Name       string            `json:"name"`       // Name of the country
	Continents []string          `json:"continents"` // Continents the country belongs to
	Population int64             `json:"population"` // Population of the country
	Languages  map[string]string `json:"languages"`  // Languages spoken with and there abbreviations
	Borders    []string          `json:"borders"`    // Borders of the country
	Flag       string            `json:"flag"`       // Flag png
	Capital    []string          `json:"capital"`    // Capital city
	Cities     []string          `json:"cities"`     // Cities of the country
}

/* Example Endpoint Response:
/*
{
"name": "Eritrea",
"continents": [
"Africa"
],
"population": 5352000,
"languages": {
"ara": "Arabic",
"eng": "English",
"tir": "Tigrinya"
},
"borders": [
"DJI",
"ETH",
"SDN"
],
"flag": "https://flagcdn.com/w320/er.png",
"capital": [
"Asmara"
],
"cities": [
"Anseba",
"Debub",
"Debubawi K'eyih Bahri",
"Gash Barka",
]
}*/
