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

type BatchPersonalNameIn struct {
	PersonalNames []PersonalNameIn `json:"personalNames,omitempty"`
}
