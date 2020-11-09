/*
 * NamSor API v2
 *
 * NamSor API v2 : enpoints to process personal names (gender, cultural origin or ethnicity) in all alphabets or languages. Use GET methods for small tests, but prefer POST methods for higher throughput (batch processing of up to 100 names at a time). Need something you can't find here? We have many more features coming soon. Let us know, we'll do our best to add it! 
 *
 * API version: 2.0.11
 * Contact: contact@namsor.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// Represents the output of inferring the LIKELY country of Origin from a personal name.
type FirstLastNameOriginedOut struct {
	Id string `json:"id,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName string `json:"lastName,omitempty"`
	// Most likely country of Origin
	CountryOrigin string `json:"countryOrigin,omitempty"`
	// Second best alternative : country of Origin
	CountryOriginAlt string `json:"countryOriginAlt,omitempty"`
	// List countries of Origin (top 10)
	CountriesOriginTop []string `json:"countriesOriginTop,omitempty"`
	// Compatibility to NamSor_v1 Origin score value
	Score float64 `json:"score,omitempty"`
	// Most likely region of Origin (based on countryOrigin ISO2 code)
	RegionOrigin string `json:"regionOrigin,omitempty"`
	// Most likely region of Origin (based on countryOrigin ISO2 code)
	TopRegionOrigin string `json:"topRegionOrigin,omitempty"`
	// Most likely region of Origin (based on countryOrigin ISO2 code)
	SubRegionOrigin string `json:"subRegionOrigin,omitempty"`
	ProbabilityCalibrated float64 `json:"probabilityCalibrated,omitempty"`
	ProbabilityAltCalibrated float64 `json:"probabilityAltCalibrated,omitempty"`
}