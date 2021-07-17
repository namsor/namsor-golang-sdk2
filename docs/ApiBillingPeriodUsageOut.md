# ApiBillingPeriodUsageOut

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | **string** | User API Key. | [optional] 
**SubscriptionStarted** | **int64** | Datetime when the user subscribed to the plan. | [optional] 
**PeriodStarted** | **int64** | Datetime when the the plan&#39;s current period started. | [optional] 
**PeriodEnded** | **int64** | Datetime when the the plan&#39;s current period endend. | [optional] 
**StripeCurrentPeriodEnd** | **int64** | Datetime when the the plan&#39;s current period endend (in Stripe). Internal and Stripe periodicity should ~coincide. | [optional] 
**StripeCurrentPeriodStart** | **int64** | Datetime when the the plan&#39;s current period started (in Stripe). Internal and Stripe periodicity should ~coincide. | [optional] 
**BillingStatus** | **string** | Current period billing status ex OPEN. | [optional] 
**Usage** | **int64** | Current period usage in units (NB some API endpoints use more than one unit). | [optional] 
**SoftLimit** | **int64** | Current period soft limit (reaching the limit sends an email notification). | [optional] 
**HardLimit** | **int64** | Current period hard limit (reaching the limit sends an email notification and blocks the API Key). | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


