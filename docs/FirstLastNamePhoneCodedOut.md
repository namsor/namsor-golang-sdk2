# FirstLastNamePhoneCodedOut

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Script** | **string** |  | [optional] 
**Id** | **string** |  | [optional] 
**FirstName** | **string** | The first name (also known as given name) | [optional] 
**LastName** | **string** | The last name (also known as family name, or surname) | [optional] 
**InternationalPhoneNumberVerified** | **string** | The normalized phone number, verified using libphonenumber. | [optional] 
**PhoneCountryIso2Verified** | **string** | The phone ISO2 country code, verified using libphonenumber. | [optional] 
**PhoneCountryCode** | **int32** | The phone country code of the phone number, verified using libphonenumber. | [optional] 
**PhoneCountryCodeAlt** | **int32** | The best alternative phone country code of the phone number. | [optional] 
**PhoneCountryIso2** | **string** | The likely country of the phone number. | [optional] 
**PhoneCountryIso2Alt** | **string** | The best alternative country of the phone number. | [optional] 
**OriginCountryIso2** | **string** | The likely country of origin of the name. | [optional] 
**OriginCountryIso2Alt** | **string** | The best alternative country of origin of the name. | [optional] 
**PhoneNumber** | **string** | The input phone number. | [optional] 
**Verified** | **bool** | Indicates if the phone number could be positively verified using libphonenumber. | [optional] 
**Score** | **float64** | Higher score is better, but score is not normalized. Use calibratedProbability if available.  | [optional] 
**CountryIso2** | **string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


