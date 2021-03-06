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

type ApiBillingPeriodUsageOut struct {
	ApiKey string `json:"apiKey,omitempty"`
	SubscriptionStarted int64 `json:"subscriptionStarted,omitempty"`
	PeriodStarted int64 `json:"periodStarted,omitempty"`
	PeriodEnded int64 `json:"periodEnded,omitempty"`
	StripeCurrentPeriodEnd int64 `json:"stripeCurrentPeriodEnd,omitempty"`
	StripeCurrentPeriodStart int64 `json:"stripeCurrentPeriodStart,omitempty"`
	BillingStatus string `json:"billingStatus,omitempty"`
	Usage int64 `json:"usage,omitempty"`
	SoftLimit int64 `json:"softLimit,omitempty"`
	HardLimit int64 `json:"hardLimit,omitempty"`
}
