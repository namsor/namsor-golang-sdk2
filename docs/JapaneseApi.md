# \JapaneseApi

All URIs are relative to *https://v2.namsor.com/NamSorAPIv2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenderJapaneseNameFull**](JapaneseApi.md#GenderJapaneseNameFull) | **Get** /api2/json/genderJapaneseNameFull/{japaneseName} | Infer the likely gender of a Japanese full name ex. 王晓明
[**GenderJapaneseNameFullBatch**](JapaneseApi.md#GenderJapaneseNameFullBatch) | **Post** /api2/json/genderJapaneseNameFullBatch | Infer the likely gender of up to 100 full names
[**GenderJapaneseNamePinyin**](JapaneseApi.md#GenderJapaneseNamePinyin) | **Get** /api2/json/genderJapaneseName/{japaneseSurname}/{japaneseGivenName} | Infer the likely gender of a Japanese name in LATIN (Pinyin).
[**GenderJapaneseNamePinyinBatch**](JapaneseApi.md#GenderJapaneseNamePinyinBatch) | **Post** /api2/json/genderJapaneseNameBatch | Infer the likely gender of up to 100 Japanese names in LATIN (Pinyin).
[**JapaneseNameKanjiCandidates**](JapaneseApi.md#JapaneseNameKanjiCandidates) | **Get** /api2/json/japaneseNameKanjiCandidates/{japaneseSurnameLatin}/{japaneseGivenNameLatin} | Identify japanese name candidates in KANJI, based on the romanized name ex. Yamamoto Sanae
[**JapaneseNameKanjiCandidatesBatch**](JapaneseApi.md#JapaneseNameKanjiCandidatesBatch) | **Post** /api2/json/japaneseNameKanjiCandidatesBatch | Identify japanese name candidates in KANJI, based on the romanized name (firstName &#x3D; japaneseGivenName; lastName&#x3D;japaneseSurname), ex. Yamamoto Sanae
[**JapaneseNameLatinCandidates**](JapaneseApi.md#JapaneseNameLatinCandidates) | **Get** /api2/json/japaneseNameLatinCandidates/{japaneseSurnameKanji}/{japaneseGivenNameKanji} | Romanize japanese name, based on the name in Kanji.
[**JapaneseNameLatinCandidatesBatch**](JapaneseApi.md#JapaneseNameLatinCandidatesBatch) | **Post** /api2/json/japaneseNameLatinCandidatesBatch | Romanize japanese names, based on the name in KANJI
[**JapaneseNameMatch**](JapaneseApi.md#JapaneseNameMatch) | **Get** /api2/json/japaneseNameMatch/{japaneseSurnameLatin}/{japaneseGivenNameLatin}/{japaneseName} | Return a score for matching Japanese name in KANJI ex. 山本 早苗 with a romanized name ex. Yamamoto Sanae
[**JapaneseNameMatchBatch**](JapaneseApi.md#JapaneseNameMatchBatch) | **Post** /api2/json/japaneseNameMatchBatch | Return a score for matching a list of Japanese names in KANJI ex. 山本 早苗 with romanized names ex. Yamamoto Sanae
[**JapaneseNameMatchFeedbackLoop**](JapaneseApi.md#JapaneseNameMatchFeedbackLoop) | **Get** /api2/json/japaneseNameMatchFeedbackLoop/{japaneseSurnameLatin}/{japaneseGivenNameLatin}/{japaneseName} | [CREDITS 1 UNIT] Feedback loop to better perform matching Japanese name in KANJI ex. 山本 早苗 with a romanized name ex. Yamamoto Sanae
[**ParseJapaneseName**](JapaneseApi.md#ParseJapaneseName) | **Get** /api2/json/parseJapaneseName/{japaneseName} | Infer the likely first/last name structure of a name, ex. 山本 早苗 or Yamamoto Sanae
[**ParseJapaneseNameBatch**](JapaneseApi.md#ParseJapaneseNameBatch) | **Post** /api2/json/parseJapaneseNameBatch | Infer the likely first/last name structure of a name, ex. 山本 早苗 or Yamamoto Sanae 


# **GenderJapaneseNameFull**
> PersonalNameGenderedOut GenderJapaneseNameFull(ctx, japaneseName)
Infer the likely gender of a Japanese full name ex. 王晓明

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **japaneseName** | **string**|  | 

### Return type

[**PersonalNameGenderedOut**](PersonalNameGenderedOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GenderJapaneseNameFullBatch**
> BatchPersonalNameGenderedOut GenderJapaneseNameFullBatch(ctx, optional)
Infer the likely gender of up to 100 full names

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GenderJapaneseNameFullBatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GenderJapaneseNameFullBatchOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchPersonalNameIn** | [**optional.Interface of BatchPersonalNameIn**](BatchPersonalNameIn.md)| A list of personal names | 

### Return type

[**BatchPersonalNameGenderedOut**](BatchPersonalNameGenderedOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GenderJapaneseNamePinyin**
> FirstLastNameGenderedOut GenderJapaneseNamePinyin(ctx, japaneseSurname, japaneseGivenName)
Infer the likely gender of a Japanese name in LATIN (Pinyin).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **japaneseSurname** | **string**|  | 
  **japaneseGivenName** | **string**|  | 

### Return type

[**FirstLastNameGenderedOut**](FirstLastNameGenderedOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GenderJapaneseNamePinyinBatch**
> BatchFirstLastNameGenderedOut GenderJapaneseNamePinyinBatch(ctx, optional)
Infer the likely gender of up to 100 Japanese names in LATIN (Pinyin).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GenderJapaneseNamePinyinBatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GenderJapaneseNamePinyinBatchOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchFirstLastNameIn** | [**optional.Interface of BatchFirstLastNameIn**](BatchFirstLastNameIn.md)| A list of names, with country code. | 

### Return type

[**BatchFirstLastNameGenderedOut**](BatchFirstLastNameGenderedOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **JapaneseNameKanjiCandidates**
> RomanizedNameOut JapaneseNameKanjiCandidates(ctx, japaneseSurnameLatin, japaneseGivenNameLatin)
Identify japanese name candidates in KANJI, based on the romanized name ex. Yamamoto Sanae

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **japaneseSurnameLatin** | **string**|  | 
  **japaneseGivenNameLatin** | **string**|  | 

### Return type

[**RomanizedNameOut**](RomanizedNameOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **JapaneseNameKanjiCandidatesBatch**
> BatchNameMatchCandidatesOut JapaneseNameKanjiCandidatesBatch(ctx, optional)
Identify japanese name candidates in KANJI, based on the romanized name (firstName = japaneseGivenName; lastName=japaneseSurname), ex. Yamamoto Sanae

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***JapaneseNameKanjiCandidatesBatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JapaneseNameKanjiCandidatesBatchOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchFirstLastNameIn** | [**optional.Interface of BatchFirstLastNameIn**](BatchFirstLastNameIn.md)| A list of personal japanese names in LATIN, firstName &#x3D; japaneseGivenName; lastName&#x3D;japaneseSurname | 

### Return type

[**BatchNameMatchCandidatesOut**](BatchNameMatchCandidatesOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **JapaneseNameLatinCandidates**
> RomanizedNameOut JapaneseNameLatinCandidates(ctx, japaneseSurnameKanji, japaneseGivenNameKanji)
Romanize japanese name, based on the name in Kanji.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **japaneseSurnameKanji** | **string**|  | 
  **japaneseGivenNameKanji** | **string**|  | 

### Return type

[**RomanizedNameOut**](RomanizedNameOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **JapaneseNameLatinCandidatesBatch**
> BatchNameMatchCandidatesOut JapaneseNameLatinCandidatesBatch(ctx, optional)
Romanize japanese names, based on the name in KANJI

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***JapaneseNameLatinCandidatesBatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JapaneseNameLatinCandidatesBatchOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchFirstLastNameIn** | [**optional.Interface of BatchFirstLastNameIn**](BatchFirstLastNameIn.md)| A list of personal japanese names in KANJI, firstName &#x3D; japaneseGivenName; lastName&#x3D;japaneseSurname | 

### Return type

[**BatchNameMatchCandidatesOut**](BatchNameMatchCandidatesOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **JapaneseNameMatch**
> RomanizedNameOut JapaneseNameMatch(ctx, japaneseSurnameLatin, japaneseGivenNameLatin, japaneseName)
Return a score for matching Japanese name in KANJI ex. 山本 早苗 with a romanized name ex. Yamamoto Sanae

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **japaneseSurnameLatin** | **string**|  | 
  **japaneseGivenNameLatin** | **string**|  | 
  **japaneseName** | **string**|  | 

### Return type

[**RomanizedNameOut**](RomanizedNameOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **JapaneseNameMatchBatch**
> BatchNameMatchCandidatesOut JapaneseNameMatchBatch(ctx, optional)
Return a score for matching a list of Japanese names in KANJI ex. 山本 早苗 with romanized names ex. Yamamoto Sanae

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***JapaneseNameMatchBatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JapaneseNameMatchBatchOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchFirstLastNameIn** | [**optional.Interface of BatchFirstLastNameIn**](BatchFirstLastNameIn.md)| A list of personal Japanese names in LATIN, firstName &#x3D; japaneseGivenName; lastName&#x3D;japaneseSurname | 

### Return type

[**BatchNameMatchCandidatesOut**](BatchNameMatchCandidatesOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **JapaneseNameMatchFeedbackLoop**
> RomanizedNameOut JapaneseNameMatchFeedbackLoop(ctx, japaneseSurnameLatin, japaneseGivenNameLatin, japaneseName)
[CREDITS 1 UNIT] Feedback loop to better perform matching Japanese name in KANJI ex. 山本 早苗 with a romanized name ex. Yamamoto Sanae

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **japaneseSurnameLatin** | **string**|  | 
  **japaneseGivenNameLatin** | **string**|  | 
  **japaneseName** | **string**|  | 

### Return type

[**RomanizedNameOut**](RomanizedNameOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ParseJapaneseName**
> PersonalNameParsedOut ParseJapaneseName(ctx, japaneseName)
Infer the likely first/last name structure of a name, ex. 山本 早苗 or Yamamoto Sanae

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **japaneseName** | **string**|  | 

### Return type

[**PersonalNameParsedOut**](PersonalNameParsedOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ParseJapaneseNameBatch**
> BatchPersonalNameParsedOut ParseJapaneseNameBatch(ctx, optional)
Infer the likely first/last name structure of a name, ex. 山本 早苗 or Yamamoto Sanae 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ParseJapaneseNameBatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ParseJapaneseNameBatchOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchPersonalNameIn** | [**optional.Interface of BatchPersonalNameIn**](BatchPersonalNameIn.md)| A list of personal names | 

### Return type

[**BatchPersonalNameParsedOut**](BatchPersonalNameParsedOut.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

