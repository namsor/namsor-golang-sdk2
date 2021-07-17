# FirstLastNameOriginedOut

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Script** | **string** |  | [optional] 
**Id** | **string** |  | [optional] 
**FirstName** | **string** | The first name (also known as given name) | [optional] 
**LastName** | **string** | The last name (also known as family name, or surname) | [optional] 
**CountryOrigin** | **string** | Most likely country of Origin | [optional] 
**CountryOriginAlt** | **string** | Second best alternative : country of Origin | [optional] 
**CountriesOriginTop** | **[]string** | List countries of Origin (top 10) | [optional] 
**Score** | **float64** | Compatibility to NamSor_v1 Origin score value. Higher score is better, but score is not normalized. Use calibratedProbability if available.  | [optional] 
**RegionOrigin** | **string** | Most likely region of Origin (based on countryOrigin ISO2 code) | [optional] 
**TopRegionOrigin** | **string** | Most likely top region of Origin (based on countryOrigin ISO2 code) | [optional] 
**SubRegionOrigin** | **string** | Most likely sub region of Origin (based on countryOrigin ISO2 code) | [optional] 
**ProbabilityCalibrated** | **float64** | The calibrated probability for countryOrigin to have been guessed correctly. | [optional] 
**ProbabilityAltCalibrated** | **float64** | The calibrated probability for countryOrigin OR countryOriginAlt to have been guessed correctly. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


