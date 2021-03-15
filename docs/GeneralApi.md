# \GeneralApi

All URIs are relative to *https://v2.namsor.com/NamSorAPIv2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NameType**](GeneralApi.md#NameType) | **Get** /api2/json/nameType/{properNoun} | Infer the likely type of a proper noun (personal name, brand name, place name etc.)
[**NameTypeBatch**](GeneralApi.md#NameTypeBatch) | **Post** /api2/json/nameTypeBatch | Infer the likely common type of up to 100 proper nouns (personal name, brand name, place name etc.)
[**NameTypeGeo**](GeneralApi.md#NameTypeGeo) | **Get** /api2/json/nameTypeGeo/{properNoun}/{countryIso2} | Infer the likely type of a proper noun (personal name, brand name, place name etc.)
[**NameTypeGeoBatch**](GeneralApi.md#NameTypeGeoBatch) | **Post** /api2/json/nameTypeGeoBatch | Infer the likely common type of up to 100 proper nouns (personal name, brand name, place name etc.)


# **NameType**
> ProperNounCategorizedOut NameType(ctx, properNoun)
Infer the likely type of a proper noun (personal name, brand name, place name etc.)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **properNoun** | **string**|  | 

### Return type

[**ProperNounCategorizedOut**](ProperNounCategorizedOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NameTypeBatch**
> BatchProperNounCategorizedOut NameTypeBatch(ctx, optional)
Infer the likely common type of up to 100 proper nouns (personal name, brand name, place name etc.)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NameTypeBatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NameTypeBatchOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchNameIn** | [**optional.Interface of BatchNameIn**](BatchNameIn.md)| A list of proper names | 

### Return type

[**BatchProperNounCategorizedOut**](BatchProperNounCategorizedOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NameTypeGeo**
> ProperNounCategorizedOut NameTypeGeo(ctx, properNoun, countryIso2)
Infer the likely type of a proper noun (personal name, brand name, place name etc.)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **properNoun** | **string**|  | 
  **countryIso2** | **string**|  | 

### Return type

[**ProperNounCategorizedOut**](ProperNounCategorizedOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NameTypeGeoBatch**
> BatchProperNounCategorizedOut NameTypeGeoBatch(ctx, optional)
Infer the likely common type of up to 100 proper nouns (personal name, brand name, place name etc.)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NameTypeGeoBatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NameTypeGeoBatchOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchNameGeoIn** | [**optional.Interface of BatchNameGeoIn**](BatchNameGeoIn.md)| A list of proper names | 

### Return type

[**BatchProperNounCategorizedOut**](BatchProperNounCategorizedOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

