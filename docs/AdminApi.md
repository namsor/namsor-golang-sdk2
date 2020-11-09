# \AdminApi

All URIs are relative to *https://v2.namsor.com/NamSorAPIv2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCredits**](AdminApi.md#AddCredits) | **Get** /api2/json/addCredits/{apiKey}/{usageCredits}/{userMessage} | Add usage credits to an API Key.
[**Anonymize**](AdminApi.md#Anonymize) | **Get** /api2/json/anonymize/{source}/{anonymized} | Activate/deactivate anonymization for a source.
[**ApiStatus**](AdminApi.md#ApiStatus) | **Get** /api2/json/apiStatus | Prints the current status of the classifiers.
[**ApiUsage**](AdminApi.md#ApiUsage) | **Get** /api2/json/apiUsage | Print current API usage.
[**ApiUsageHistory**](AdminApi.md#ApiUsageHistory) | **Get** /api2/json/apiUsageHistory | Print historical API usage.
[**ApiUsageHistoryAggregate**](AdminApi.md#ApiUsageHistoryAggregate) | **Get** /api2/json/apiUsageHistoryAggregate | Print historical API usage (in an aggregated view, by service, by day/hour/min).
[**AvailablePlans**](AdminApi.md#AvailablePlans) | **Get** /api2/json/availablePlans/{token} | List all available plans in the user&#39;s preferred currency.
[**AvailablePlans1**](AdminApi.md#AvailablePlans1) | **Get** /api2/json/availablePlans | List all available plans in the default currency (usd).
[**AvailableServices**](AdminApi.md#AvailableServices) | **Get** /api2/json/apiServices | List of API services and usage cost in Units (default is 1&#x3D;ONE Unit).
[**BillingCurrencies**](AdminApi.md#BillingCurrencies) | **Get** /api2/json/billingCurrencies | List possible currency options for billing (USD, EUR, GBP, ...)
[**BillingHistory**](AdminApi.md#BillingHistory) | **Get** /api2/json/billingHistory/{token} | Read the history billing information (invoices paid via Stripe or manually).
[**BillingInfo**](AdminApi.md#BillingInfo) | **Get** /api2/json/billingInfo/{token} | Read the billing information (company name, address, phone, vat ID)
[**Charge**](AdminApi.md#Charge) | **Post** /api2/json/charge | Create a Stripe Customer, based on a payment card token (from secure StripeJS) and email.
[**CorporateKey**](AdminApi.md#CorporateKey) | **Get** /api2/json/corporateKey/{apiKey}/{corporate} | Setting an API Key to a corporate status.
[**DebugLevel**](AdminApi.md#DebugLevel) | **Get** /api2/json/debugLevel/{logger}/{level} | Update debug level for a classifier
[**Flush**](AdminApi.md#Flush) | **Get** /api2/json/flush | Flush counters.
[**InvalidateCache**](AdminApi.md#InvalidateCache) | **Get** /api2/json/invalidateCache | Invalidate system caches.
[**Learnable**](AdminApi.md#Learnable) | **Get** /api2/json/learnable/{source}/{learnable} | Activate/deactivate learning from a source.
[**NamsorCounter**](AdminApi.md#NamsorCounter) | **Get** /api2/json/namsorCounter | Get the overall API counter
[**PaymentInfo**](AdminApi.md#PaymentInfo) | **Get** /api2/json/paymentInfo/{token} | Get the Stripe payment information associated with the current google auth session token.
[**ProcureKey**](AdminApi.md#ProcureKey) | **Get** /api2/json/procureKey/{token} | Procure an API Key (sent via Email), based on an auth token. Keep your API Key secret.
[**RedeployUI**](AdminApi.md#RedeployUI) | **Get** /api2/json/redeployUI/{live} | Redeploy UI from current dev branch.
[**RedeployUI1**](AdminApi.md#RedeployUI1) | **Get** /api2/json/redeployUI | Redeploy UI from current dev branch.
[**RemoveUserAccount**](AdminApi.md#RemoveUserAccount) | **Get** /api2/json/removeUserAccount/{token} | Remove the user account.
[**RemoveUserAccountOnBehalf**](AdminApi.md#RemoveUserAccountOnBehalf) | **Get** /api2/json/removeUserAccountOnBehalf/{apiKey} | Remove (on behalf) a user account.
[**Shutdown**](AdminApi.md#Shutdown) | **Get** /api2/json/shutdown | Stop learning and shutdown system.
[**SoftwareVersion**](AdminApi.md#SoftwareVersion) | **Get** /api2/json/softwareVersion | Get the current software version
[**SourceStats**](AdminApi.md#SourceStats) | **Get** /api2/json/sourceStats/{source} | Print basic source statistics.
[**Stats**](AdminApi.md#Stats) | **Get** /api2/json/stats | Print basic system statistics.
[**StripeConnect**](AdminApi.md#StripeConnect) | **Get** /api2/json/stripeConnect | Connects a Stripe Account.
[**SubscribePlan**](AdminApi.md#SubscribePlan) | **Get** /api2/json/subscribePlan/{planName}/{token} | Subscribe to a give API plan, using the user&#39;s preferred or default currency.
[**SubscribePlanOnBehalf**](AdminApi.md#SubscribePlanOnBehalf) | **Get** /api2/json/subscribePlanOnBehalf/{planName}/{apiKey} | Subscribe to a give API plan, using the user&#39;s preferred or default currency (admin only).
[**TaxonomyClasses**](AdminApi.md#TaxonomyClasses) | **Get** /api2/json/taxonomyClasses/{classifierName} | Print the taxonomy classes valid for the given classifier.
[**UpdateBillingInfo**](AdminApi.md#UpdateBillingInfo) | **Post** /api2/json/updateBillingInfo/{token} | Sets or update the billing information (company name, address, phone, vat ID)
[**UpdateLimit**](AdminApi.md#UpdateLimit) | **Get** /api2/json/updateLimit/{usageLimit}/{hardOrSoft}/{token} | Modifies the hard/soft limit on the API plan&#39;s overages (default is 0$ soft limit).
[**UpdatePaymentDefault**](AdminApi.md#UpdatePaymentDefault) | **Get** /api2/json/updatePaymentDefault/{defautSourceId}/{token} | Update the default Stripe card associated with the current google auth session token.
[**UserInfo**](AdminApi.md#UserInfo) | **Get** /api2/json/userInfo/{token} | Get the user profile associated with the current google auth session token.
[**VerifyEmail**](AdminApi.md#VerifyEmail) | **Get** /api2/json/verifyEmail/{emailToken} | Verifies an email, based on token sent to that email
[**VerifyRemoveEmail**](AdminApi.md#VerifyRemoveEmail) | **Get** /api2/json/verifyRemoveEmail/{emailToken} | Verifies an email, based on token sent to that email
[**Vet**](AdminApi.md#Vet) | **Get** /api2/json/vetting/{source}/{vetted} | Vetting of a source.


# **AddCredits**
> SystemMetricsOut AddCredits(ctx, apiKey, usageCredits, userMessage)
Add usage credits to an API Key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**|  | 
  **usageCredits** | **int64**|  | 
  **userMessage** | **string**|  | 

### Return type

[**SystemMetricsOut**](SystemMetricsOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Anonymize**
> Anonymize(ctx, source, anonymized)
Activate/deactivate anonymization for a source.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **source** | **string**|  | 
  **anonymized** | **bool**|  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStatus**
> ApiPlansOut ApiStatus(ctx, )
Prints the current status of the classifiers.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiPlansOut**](APIPlansOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiUsage**
> ApiPeriodUsageOut ApiUsage(ctx, )
Print current API usage.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiPeriodUsageOut**](APIPeriodUsageOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiUsageHistory**
> ApiPeriodUsageOut ApiUsageHistory(ctx, )
Print historical API usage.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiPeriodUsageOut**](APIPeriodUsageOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiUsageHistoryAggregate**
> ApiPeriodUsageOut ApiUsageHistoryAggregate(ctx, )
Print historical API usage (in an aggregated view, by service, by day/hour/min).

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiPeriodUsageOut**](APIPeriodUsageOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AvailablePlans**
> ApiPlansOut AvailablePlans(ctx, token)
List all available plans in the user's preferred currency.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**|  | 

### Return type

[**ApiPlansOut**](APIPlansOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AvailablePlans1**
> ApiPlansOut AvailablePlans1(ctx, )
List all available plans in the default currency (usd).

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiPlansOut**](APIPlansOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AvailableServices**
> ApiPlansOut AvailableServices(ctx, )
List of API services and usage cost in Units (default is 1=ONE Unit).

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiPlansOut**](APIPlansOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BillingCurrencies**
> CurrenciesOut BillingCurrencies(ctx, )
List possible currency options for billing (USD, EUR, GBP, ...)

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CurrenciesOut**](CurrenciesOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BillingHistory**
> BillingHistoryOut BillingHistory(ctx, token)
Read the history billing information (invoices paid via Stripe or manually).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**|  | 

### Return type

[**BillingHistoryOut**](BillingHistoryOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BillingInfo**
> BillingInfoInOut BillingInfo(ctx, token)
Read the billing information (company name, address, phone, vat ID)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**|  | 

### Return type

[**BillingInfoInOut**](BillingInfoInOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Charge**
> ApiKeyOut Charge(ctx, optional)
Create a Stripe Customer, based on a payment card token (from secure StripeJS) and email.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ChargeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChargeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject** | [**optional.Interface of InlineObject**](InlineObject.md)|  | 

### Return type

[**ApiKeyOut**](APIKeyOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CorporateKey**
> CorporateKey(ctx, apiKey, corporate)
Setting an API Key to a corporate status.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**|  | 
  **corporate** | **bool**|  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DebugLevel**
> DebugLevel(ctx, logger, level)
Update debug level for a classifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logger** | **string**|  | 
  **level** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Flush**
> Flush(ctx, )
Flush counters.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InvalidateCache**
> InvalidateCache(ctx, )
Invalidate system caches.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Learnable**
> Learnable(ctx, source, learnable)
Activate/deactivate learning from a source.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **source** | **string**|  | 
  **learnable** | **bool**|  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NamsorCounter**
> SoftwareVersionOut NamsorCounter(ctx, )
Get the overall API counter

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SoftwareVersionOut**](SoftwareVersionOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PaymentInfo**
> ApiKeyOut PaymentInfo(ctx, token)
Get the Stripe payment information associated with the current google auth session token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**|  | 

### Return type

[**ApiKeyOut**](APIKeyOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProcureKey**
> ApiKeyOut ProcureKey(ctx, token)
Procure an API Key (sent via Email), based on an auth token. Keep your API Key secret.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**|  | 

### Return type

[**ApiKeyOut**](APIKeyOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RedeployUI**
> RedeployUI(ctx, live)
Redeploy UI from current dev branch.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **live** | **bool**|  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RedeployUI1**
> RedeployUI1(ctx, )
Redeploy UI from current dev branch.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveUserAccount**
> ApiPlanSubscriptionOut RemoveUserAccount(ctx, token)
Remove the user account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**|  | 

### Return type

[**ApiPlanSubscriptionOut**](APIPlanSubscriptionOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveUserAccountOnBehalf**
> ApiPlanSubscriptionOut RemoveUserAccountOnBehalf(ctx, apiKey)
Remove (on behalf) a user account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**|  | 

### Return type

[**ApiPlanSubscriptionOut**](APIPlanSubscriptionOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Shutdown**
> Shutdown(ctx, )
Stop learning and shutdown system.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SoftwareVersion**
> SoftwareVersionOut SoftwareVersion(ctx, )
Get the current software version

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SoftwareVersionOut**](SoftwareVersionOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SourceStats**
> SystemMetricsOut SourceStats(ctx, source)
Print basic source statistics.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **source** | **string**|  | 

### Return type

[**SystemMetricsOut**](SystemMetricsOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Stats**
> SystemMetricsOut Stats(ctx, )
Print basic system statistics.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SystemMetricsOut**](SystemMetricsOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StripeConnect**
> StripeConnect(ctx, optional)
Connects a Stripe Account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StripeConnectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StripeConnectOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scope** | **optional.String**|  | 
 **code** | **optional.String**|  | 
 **error_** | **optional.String**|  | 
 **errorDescription** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubscribePlan**
> ApiPlanSubscriptionOut SubscribePlan(ctx, planName, token)
Subscribe to a give API plan, using the user's preferred or default currency.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planName** | **string**|  | 
  **token** | **string**|  | 

### Return type

[**ApiPlanSubscriptionOut**](APIPlanSubscriptionOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubscribePlanOnBehalf**
> ApiPlanSubscriptionOut SubscribePlanOnBehalf(ctx, planName, apiKey)
Subscribe to a give API plan, using the user's preferred or default currency (admin only).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planName** | **string**|  | 
  **apiKey** | **string**|  | 

### Return type

[**ApiPlanSubscriptionOut**](APIPlanSubscriptionOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TaxonomyClasses**
> ApiPlansOut TaxonomyClasses(ctx, classifierName)
Print the taxonomy classes valid for the given classifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **classifierName** | **string**|  | 

### Return type

[**ApiPlansOut**](APIPlansOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBillingInfo**
> BillingInfoInOut UpdateBillingInfo(ctx, token, optional)
Sets or update the billing information (company name, address, phone, vat ID)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**|  | 
 **optional** | ***UpdateBillingInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateBillingInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **billingInfoInOut** | [**optional.Interface of BillingInfoInOut**](BillingInfoInOut.md)|  | 

### Return type

[**BillingInfoInOut**](BillingInfoInOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLimit**
> ApiPeriodUsageOut UpdateLimit(ctx, usageLimit, hardOrSoft, token)
Modifies the hard/soft limit on the API plan's overages (default is 0$ soft limit).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **usageLimit** | **int32**|  | 
  **hardOrSoft** | **bool**|  | 
  **token** | **string**|  | 

### Return type

[**ApiPeriodUsageOut**](APIPeriodUsageOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePaymentDefault**
> ApiKeyOut UpdatePaymentDefault(ctx, defautSourceId, token)
Update the default Stripe card associated with the current google auth session token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **defautSourceId** | **string**|  | 
  **token** | **string**|  | 

### Return type

[**ApiKeyOut**](APIKeyOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserInfo**
> ApiKeyOut UserInfo(ctx, token)
Get the user profile associated with the current google auth session token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**|  | 

### Return type

[**ApiKeyOut**](APIKeyOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerifyEmail**
> ApiKeyOut VerifyEmail(ctx, emailToken)
Verifies an email, based on token sent to that email

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **emailToken** | **string**|  | 

### Return type

[**ApiKeyOut**](APIKeyOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerifyRemoveEmail**
> ApiKeyOut VerifyRemoveEmail(ctx, emailToken)
Verifies an email, based on token sent to that email

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **emailToken** | **string**|  | 

### Return type

[**ApiKeyOut**](APIKeyOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Vet**
> Vet(ctx, source, vetted)
Vetting of a source.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **source** | **string**|  | 
  **vetted** | **bool**|  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

