# \SocialApi

All URIs are relative to *https://v2.namsor.com/NamSorAPIv2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PhoneCode**](SocialApi.md#PhoneCode) | **Get** /api2/json/phoneCode/{firstName}/{lastName}/{phoneNumber} | [USES 11 UNITS PER NAME] Infer the likely country and phone prefix, given a personal name and formatted / unformatted phone number.
[**PhoneCodeBatch**](SocialApi.md#PhoneCodeBatch) | **Post** /api2/json/phoneCodeBatch | [USES 11 UNITS PER NAME] Infer the likely country and phone prefix, of up to 100 personal names, detecting automatically the local context given a name and formatted / unformatted phone number.
[**PhoneCodeGeo**](SocialApi.md#PhoneCodeGeo) | **Get** /api2/json/phoneCodeGeo/{firstName}/{lastName}/{phoneNumber}/{countryIso2} | [USES 11 UNITS PER NAME] Infer the likely phone prefix, given a personal name and formatted / unformatted phone number, with a local context (ISO2 country of residence).
[**PhoneCodeGeoBatch**](SocialApi.md#PhoneCodeGeoBatch) | **Post** /api2/json/phoneCodeGeoBatch | [USES 11 UNITS PER NAME] Infer the likely country and phone prefix, of up to 100 personal names, with a local context (ISO2 country of residence).
[**PhoneCodeGeoFeedbackLoop**](SocialApi.md#PhoneCodeGeoFeedbackLoop) | **Get** /api2/json/phoneCodeGeoFeedbackLoop/{firstName}/{lastName}/{phoneNumber}/{phoneNumberE164}/{countryIso2} | [CREDITS 1 UNIT] Feedback loop to better infer the likely phone prefix, given a personal name and formatted / unformatted phone number, with a local context (ISO2 country of residence).


# **PhoneCode**
> FirstLastNamePhoneCodedOut PhoneCode(ctx, firstName, lastName, phoneNumber)
[USES 11 UNITS PER NAME] Infer the likely country and phone prefix, given a personal name and formatted / unformatted phone number.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firstName** | **string**|  | 
  **lastName** | **string**|  | 
  **phoneNumber** | **string**|  | 

### Return type

[**FirstLastNamePhoneCodedOut**](FirstLastNamePhoneCodedOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PhoneCodeBatch**
> BatchFirstLastNamePhoneCodedOut PhoneCodeBatch(ctx, optional)
[USES 11 UNITS PER NAME] Infer the likely country and phone prefix, of up to 100 personal names, detecting automatically the local context given a name and formatted / unformatted phone number.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PhoneCodeBatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PhoneCodeBatchOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchFirstLastNamePhoneNumberIn** | [**optional.Interface of BatchFirstLastNamePhoneNumberIn**](BatchFirstLastNamePhoneNumberIn.md)| A list of personal names | 

### Return type

[**BatchFirstLastNamePhoneCodedOut**](BatchFirstLastNamePhoneCodedOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PhoneCodeGeo**
> FirstLastNamePhoneCodedOut PhoneCodeGeo(ctx, firstName, lastName, phoneNumber, countryIso2)
[USES 11 UNITS PER NAME] Infer the likely phone prefix, given a personal name and formatted / unformatted phone number, with a local context (ISO2 country of residence).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firstName** | **string**|  | 
  **lastName** | **string**|  | 
  **phoneNumber** | **string**|  | 
  **countryIso2** | **string**|  | 

### Return type

[**FirstLastNamePhoneCodedOut**](FirstLastNamePhoneCodedOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PhoneCodeGeoBatch**
> BatchFirstLastNamePhoneCodedOut PhoneCodeGeoBatch(ctx, optional)
[USES 11 UNITS PER NAME] Infer the likely country and phone prefix, of up to 100 personal names, with a local context (ISO2 country of residence).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PhoneCodeGeoBatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PhoneCodeGeoBatchOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchFirstLastNamePhoneNumberGeoIn** | [**optional.Interface of BatchFirstLastNamePhoneNumberGeoIn**](BatchFirstLastNamePhoneNumberGeoIn.md)| A list of personal names | 

### Return type

[**BatchFirstLastNamePhoneCodedOut**](BatchFirstLastNamePhoneCodedOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PhoneCodeGeoFeedbackLoop**
> FirstLastNamePhoneCodedOut PhoneCodeGeoFeedbackLoop(ctx, firstName, lastName, phoneNumber, phoneNumberE164, countryIso2)
[CREDITS 1 UNIT] Feedback loop to better infer the likely phone prefix, given a personal name and formatted / unformatted phone number, with a local context (ISO2 country of residence).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firstName** | **string**|  | 
  **lastName** | **string**|  | 
  **phoneNumber** | **string**|  | 
  **phoneNumberE164** | **string**|  | 
  **countryIso2** | **string**|  | 

### Return type

[**FirstLastNamePhoneCodedOut**](FirstLastNamePhoneCodedOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

