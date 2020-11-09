# \GeneralApi

All URIs are relative to *https://v2.namsor.com/NamSorAPIv2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NameType**](GeneralApi.md#NameType) | **Get** /api2/json/nameType/{properNoun} | Infer the likely type of a proper noun (personal name, brand name, place name etc.)
[**NameType1**](GeneralApi.md#NameType1) | **Get** /api2/json/nameType/{properNoun}/{countryIso2} | Infer the likely type of a proper noun (personal name, brand name, place name etc.)


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

# **NameType1**
> ProperNounCategorizedOut NameType1(ctx, properNoun, countryIso2)
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

