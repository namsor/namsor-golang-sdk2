# FirstLastNameDiasporaedOut

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Script** | **string** |  | [optional] 
**Id** | **string** |  | [optional] 
**FirstName** | **string** | The first name (also known as given name) | [optional] 
**LastName** | **string** | The last name (also known as family name, or surname) | [optional] 
**Score** | **float64** | Compatibility to NamSor_v1 Diaspora score value. Higher score is better, but score is not normalized. Use calibratedProbability if available.  | [optional] 
**EthnicityAlt** | **string** | The second best alternative ethnicity | [optional] 
**Ethnicity** | **string** | The most likely ethnicity | [optional] 
**Lifted** | **bool** | Indicates if the output ethnicity is based on machine learning only, or further lifted as a known fact by a country-specific rule. Let us know if you believe ethnicity is incorrect on a specific case where lifted is true. | [optional] 
**CountryIso2** | **string** | From input data, the countryIso2 of geographic context (US,CA etc.) | [optional] 
**EthnicitiesTop** | **[]string** | List most likely ethnicities (top 10) | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


