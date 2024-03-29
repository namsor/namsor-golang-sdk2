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

// The list of classifiers and versions.
type ApiClassifierOut struct {
	// The classifier name
	ClassifierName string `json:"classifierName,omitempty"`
	// True if the classifier is serving requests (has reached minimal learning, is not shutting down)
	Serving bool `json:"serving,omitempty"`
	// True if the classifier is learning
	Learning bool `json:"learning,omitempty"`
	// True if the classifier is shutting down
	ShuttingDown bool `json:"shuttingDown,omitempty"`
	// True if the classifier has finished the initial learning and calibrated probabilities (meanwhile, during initial learning, probabilities will be equal to -1)
	ProbabilityCalibrated bool `json:"probabilityCalibrated,omitempty"`
}
