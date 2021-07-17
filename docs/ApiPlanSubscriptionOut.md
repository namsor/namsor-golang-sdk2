# ApiPlanSubscriptionOut

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | **string** | User API Key. | [optional] 
**PlanStarted** | **int64** | Datetime when the user subscribed to the current plan. | [optional] 
**PriorPlanStarted** | **int64** | Datetime when the user subscribed to the prior plan. | [optional] 
**PlanEnded** | **int64** | Datetime when the user ended the plan. | [optional] 
**TaxRate** | **float64** | Applicable tax rate for the plan. | [optional] 
**PlanName** | **string** | Current plan name. | [optional] 
**PlanBaseFeesKey** | **string** | Current plan key (as in Stripe product). | [optional] 
**PlanStatus** | **string** | Plan status. | [optional] 
**PlanQuota** | **int64** | Current plan quota in quantity of units (NB: some API use several units per name). | [optional] 
**PriceUSD** | **float64** | Current plan monthly price expressed in USD (includes a free quota). | [optional] 
**PriceOverageUSD** | **float64** | Current plan price for overages expressed in USD (extra price per unit above the free quota). | [optional] 
**Price** | **float64** | Current plan price for overages expressed in Currency (extra price per unit above the free quota). | [optional] 
**PriceOverage** | **float64** | Current plan price for overages expressed in Currency (extra price per unit above the free quota). | [optional] 
**Currency** | **string** | Current plan Currency for prices. | [optional] 
**CurrencyFactor** | **float64** | For USD, GBP, EUR - the factor is 1. | [optional] 
**StripeCustomerId** | **string** | Stripe customer identifier. | [optional] 
**StripeStatus** | **string** | Stripe status ex active. | [optional] 
**StripeSubscription** | **string** | Stripe subscription identifier. | [optional] 
**UserId** | **string** | Internal user identifier. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


