/*
 * NamSor API v2
 *
 * NamSor API v2 : enpoints to process personal names (gender, cultural origin or ethnicity) in all alphabets or languages. Use GET methods for small tests, but prefer POST methods for higher throughput (batch processing of up to 100 names at a time). Need something you can't find here? We have many more features coming soon. Let us know, we'll do our best to add it! 
 *
 * API version: 2.0.13
 * Contact: contact@namsor.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package namsorapi

type PersonalNameGeoOut struct {
	Script string `json:"script,omitempty"`
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Score float64 `json:"score,omitempty"`
	Country string `json:"country,omitempty"`
	CountryAlt string `json:"countryAlt,omitempty"`
	Region string `json:"region,omitempty"`
	TopRegion string `json:"topRegion,omitempty"`
	SubRegion string `json:"subRegion,omitempty"`
	// List countries (top 10)
	CountriesTop []string `json:"countriesTop,omitempty"`
	ProbabilityCalibrated float64 `json:"probabilityCalibrated,omitempty"`
	ProbabilityAltCalibrated float64 `json:"probabilityAltCalibrated,omitempty"`
	Category string `json:"category,omitempty"`
}
