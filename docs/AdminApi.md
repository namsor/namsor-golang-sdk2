# \AdminApi

All URIs are relative to *https://v2.namsor.com/NamSorAPIv2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Anonymize**](AdminApi.md#Anonymize) | **Get** /api2/json/anonymize/{source}/{anonymized} | Activate/deactivate anonymization for a source.
[**ApiStatus**](AdminApi.md#ApiStatus) | **Get** /api2/json/apiStatus | Prints the current status of the classifiers. A classifier name in apiStatus corresponds to a service name in apiServices.
[**ApiUsage**](AdminApi.md#ApiUsage) | **Get** /api2/json/apiUsage | Print current API usage.
[**ApiUsageHistory**](AdminApi.md#ApiUsageHistory) | **Get** /api2/json/apiUsageHistory | Print historical API usage.
[**ApiUsageHistoryAggregate**](AdminApi.md#ApiUsageHistoryAggregate) | **Get** /api2/json/apiUsageHistoryAggregate | Print historical API usage (in an aggregated view, by service, by day/hour/min).
[**AvailableServices**](AdminApi.md#AvailableServices) | **Get** /api2/json/apiServices | List of classification services and usage cost in Units per classification (default is 1&#x3D;ONE Unit). Some API endpoints (ex. Corridor) combine multiple classifiers.
[**Disable**](AdminApi.md#Disable) | **Get** /api2/json/disable/{source}/{disabled} | Activate/deactivate an API Key.
[**Learnable**](AdminApi.md#Learnable) | **Get** /api2/json/learnable/{source}/{learnable} | Activate/deactivate learning from a source.
[**SoftwareVersion**](AdminApi.md#SoftwareVersion) | **Get** /api2/json/softwareVersion | Get the current software version
[**TaxonomyClasses**](AdminApi.md#TaxonomyClasses) | **Get** /api2/json/taxonomyClasses/{classifierName} | Print the taxonomy classes valid for the given classifier.


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
> ApiClassifiersStatusOut ApiStatus(ctx, )
Prints the current status of the classifiers. A classifier name in apiStatus corresponds to a service name in apiServices.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiClassifiersStatusOut**](APIClassifiersStatusOut.md)

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
> ApiUsageHistoryOut ApiUsageHistory(ctx, )
Print historical API usage.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiUsageHistoryOut**](APIUsageHistoryOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiUsageHistoryAggregate**
> ApiUsageAggregatedOut ApiUsageHistoryAggregate(ctx, )
Print historical API usage (in an aggregated view, by service, by day/hour/min).

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiUsageAggregatedOut**](APIUsageAggregatedOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AvailableServices**
> ApiServicesOut AvailableServices(ctx, )
List of classification services and usage cost in Units per classification (default is 1=ONE Unit). Some API endpoints (ex. Corridor) combine multiple classifiers.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiServicesOut**](APIServicesOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Disable**
> Disable(ctx, source, disabled)
Activate/deactivate an API Key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **source** | **string**| The API Key to set as enabled/disabled. | 
  **disabled** | **bool**|  | 

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
  **source** | **string**| The API Key to set as learnable/non learnable. | 
  **learnable** | **bool**|  | 

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

# **TaxonomyClasses**
> ApiClassifierTaxonomyOut TaxonomyClasses(ctx, classifierName)
Print the taxonomy classes valid for the given classifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **classifierName** | **string**|  | 

### Return type

[**ApiClassifierTaxonomyOut**](APIClassifierTaxonomyOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

