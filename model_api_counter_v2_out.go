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

// Detailed usage as reported by the deduplicating API counter.
type ApiCounterV2Out struct {
	ApiKey ApiKeyOut `json:"apiKey,omitempty"`
	// The apiService corresponds to the classifier name.
	ApiService string `json:"apiService,omitempty"`
	// The create datetime of the counter.
	CreatedDateTime int64 `json:"createdDateTime,omitempty"`
	// The usage of the counter.
	TotalUsage int64 `json:"totalUsage,omitempty"`
	// The flush datetime of the counter.
	LastFlushedDateTime int64 `json:"lastFlushedDateTime,omitempty"`
	// The last usage datetime of the counter.
	LastUsedDateTime int64 `json:"lastUsedDateTime,omitempty"`
	// Usage of special features, such as Chinese, Japanese.
	ServiceFeaturesUsage map[string]int64 `json:"serviceFeaturesUsage,omitempty"`
}
