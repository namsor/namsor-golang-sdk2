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

// Represents the output of inferring the LIKELY gender from a personal name.
type FirstLastNameGenderedOut struct {
	Script string `json:"script,omitempty"`
	Id string `json:"id,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName string `json:"lastName,omitempty"`
	// Most likely gender
	LikelyGender string `json:"likelyGender,omitempty"`
	// Compatibility to NamSor_v1 Gender Scale M[-1..U..+1]F value
	GenderScale float64 `json:"genderScale,omitempty"`
	Score float64 `json:"score,omitempty"`
	ProbabilityCalibrated float64 `json:"probabilityCalibrated,omitempty"`
	Category string `json:"category,omitempty"`
}
