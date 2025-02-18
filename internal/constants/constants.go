package constants

// Url paths as constants

// APIVersion defines the API version used in all endpoints
const APIVersion = "/v1"

// EmptyPath represents the base path for the API with no specific resource
const EmptyPath = "/"

// CountryInfoBasePath base path for country data endpoints
const CountryInfoBasePath = "/countryInfo" + APIVersion

// PopulationBasePath base path for population data endpoints
const PopulationBasePath = "/population" + APIVersion

// StatusBasePath  base path for service status endpoints
const StatusBasePath = "/status" + APIVersion

const CountryInfoEndpoint = CountryInfoBasePath + "/info"
const PopulationEndpoint = PopulationBasePath + "/population"
const ServiceStatusEndpoint = StatusBasePath
