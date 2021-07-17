# ApiPeriodUsageOut

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscription** | [**ApiPlanSubscriptionOut**](APIPlanSubscriptionOut.md) |  | [optional] 
**BillingPeriod** | [**ApiBillingPeriodUsageOut**](APIBillingPeriodUsageOut.md) |  | [optional] 
**OverageExclTax** | **float64** | Overage amount including any tax. | [optional] 
**OverageInclTax** | **float64** | Overage amount including tax (if applicable). | [optional] 
**OverageCurrency** | **string** | Currency of the overage amount. | [optional] 
**OverageQuantity** | **int64** | Quantity above monthly quota of the current subscritpion, in units. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


