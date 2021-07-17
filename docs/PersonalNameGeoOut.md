# PersonalNameGeoOut

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Script** | **string** |  | [optional] 
**Id** | **string** |  | [optional] 
**Name** | **string** | The input name. | [optional] 
**Score** | **float64** | Higher score is better, but score is not normalized. Use calibratedProbability if available.  | [optional] 
**Country** | **string** | Most likely country  | [optional] 
**CountryAlt** | **string** | Second best alternative : country  | [optional] 
**Region** | **string** | Most likely region (based on country ISO2 code) | [optional] 
**TopRegion** | **string** | Most likely top region (based on country ISO2 code) | [optional] 
**SubRegion** | **string** | Most likely sub region (based on country ISO2 code) | [optional] 
**CountriesTop** | **[]string** | List countries (top 10) | [optional] 
**ProbabilityCalibrated** | **float64** | The calibrated probability for country to have been guessed correctly. | [optional] 
**ProbabilityAltCalibrated** | **float64** | The calibrated probability for country OR countryAlt to have been guessed correctly. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


