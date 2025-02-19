package constants

// Url paths as constants

// APIVersion defines the API version used in all endpoints
const APIVersion = "/v1"

// EmptyPath represents the base path for the API with no specific resource
const EmptyPath = "/"

// TwoLetterCountryCode YearRange Dynamic URL Arguments
const TwoLetterCountryCode = "/{code}"
const YearRange = "/{startYear}-{endYear}"

// CountryInfoBasePath Country Info Endpoints
const CountryInfoBasePath = "/countryInfo" + APIVersion
const CountryInfoEndpoint = CountryInfoBasePath + TwoLetterCountryCode + "/info"

// PopulationBasePath Population Endpoints
const PopulationBasePath = "/population" + APIVersion
const PopulationEndpoint = PopulationBasePath + TwoLetterCountryCode + YearRange

// StatusBasePath Service Status Endpoints
const StatusBasePath = "/status" + APIVersion
const ServiceStatusEndpoint = StatusBasePath
