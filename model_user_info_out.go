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

type UserInfoOut struct {
	Uid string `json:"uid,omitempty"`
	Email string `json:"email,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	EmailVerified bool `json:"emailVerified,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	PhotoUrl string `json:"photoUrl,omitempty"`
	Disabled bool `json:"disabled,omitempty"`
	FirstKnownIpAddress string `json:"firstKnownIpAddress,omitempty"`
	ProviderId string `json:"providerId,omitempty"`
	TimeStamp int64 `json:"timeStamp,omitempty"`
	VerifyToken string `json:"verifyToken,omitempty"`
	ApiKey string `json:"apiKey,omitempty"`
	StripePerishableKey string `json:"stripePerishableKey,omitempty"`
	StripeCustomerId string `json:"stripeCustomerId,omitempty"`
	OtherInfos []UserInfoOut `json:"otherInfos,omitempty"`
}