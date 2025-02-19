package models

// Json response struct
type CountryInfo struct {
	Country    string            `json:"name"`
	Continents []string          `json:"continents"`
	Population int64             `json:"population"`
	Languages  map[string]string `json:"languages"`
	Borders    []string          `json:"borders"`
	Flag       string            `json:"flag"`
	Capital    string            `json:"capital"`
	Cities     []string          `json:"cities"`
}

/*{
"name": "Norway",
"continents": ["Europe"],
"population": 4700000,
"languages": {"nno":"Norwegian Nynorsk","nob":"Norwegian Bokmål","smi":"Sami"},
"borders": ["FIN","SWE","RUS"],
"flag": "https://flagcdn.com/w320/no.png",
"capital": "Oslo",
"cities": ["Abelvaer","Adalsbruk","Adland","Agotnes",...]
}
*/
