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

// Simple metrics on a classifier, for a given expected class
type ExpectedClassMetricsOut struct {
	ClassifierName string `json:"classifierName,omitempty"`
	ExpectedClass string `json:"expectedClass,omitempty"`
	AiEstimateTotal int64 `json:"aiEstimateTotal,omitempty"`
	AiEstimatePrecision float64 `json:"aiEstimatePrecision,omitempty"`
	AiEstimateRecall float64 `json:"aiEstimateRecall,omitempty"`
	AiLearnTotal int64 `json:"aiLearnTotal,omitempty"`
}
