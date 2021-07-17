/*
 * NamSor API v2
 *
 * NamSor API v2 : enpoints to process personal names (gender, cultural origin or ethnicity) in all alphabets or languages. By default, enpoints use 1 unit per name (ex. Gender), but Ethnicity classification uses 10 to 20 units per name depending on taxonomy. Use GET methods for small tests, but prefer POST methods for higher throughput (batch processing of up to 100 names at a time). Need something you can't find here? We have many more features coming soon. Let us know, we'll do our best to add it! 
 *
 * API version: 2.0.15
 * Contact: contact@namsor.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package namsorapi

// List of API Services
type ApiServiceOut struct {
	// A service name corresponds to classifier name in apiStatus (ex. personalname_gender or personalfullname_gender)
	ServiceName string `json:"serviceName,omitempty"`
	// Groups together classifiers providing a similar service (ex. gender groups personalname_gender and personalfullname_gender)
	ServiceGroup string `json:"serviceGroup,omitempty"`
	// Indicates how many units per call this service costs (ex. the number of units per name)
	CostInUnits int32 `json:"costInUnits,omitempty"`
}
