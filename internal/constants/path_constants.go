package constants

// Url paths as constants

// APIVersion defines the API version used in all endpoints
const APIVersion = "/v1"

// EmptyPath represents the base path for the API with no specific resource
const EmptyPath = "/"

// TwoLetterCountryCode YearRange Dynamic URL Arguments
const TwoLetterCountryCode = "/{code}"
const YearRange = "/{startYear}/{endYear}"

// CountryInfoBasePath Country Info Base Path
const CountryInfoBasePath = "/countryinfo" + APIVersion

// CountryInfoEndpoint Country Info Endpoint
const CountryInfoEndpoint = CountryInfoBasePath + TwoLetterCountryCode + "/info"

// PopulationBasePath Population Info Base Path
const PopulationBasePath = CountryInfoBasePath + "/population"

// PopulationEndpoint Population End point
const PopulationEndpoint = PopulationBasePath + TwoLetterCountryCode + YearRange

// ServiceStatusEndpoint Service Status Endpoints
const ServiceStatusEndpoint = CountryInfoBasePath + "/status"
