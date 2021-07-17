# FirstLastNameGenderedOut

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Script** | **string** |  | [optional] 
**Id** | **string** |  | [optional] 
**FirstName** | **string** | The first name (also known as given name) | [optional] 
**LastName** | **string** | The last name (also known as family name, or surname) | [optional] 
**LikelyGender** | **string** | Most likely gender | [optional] 
**GenderScale** | **float64** | Compatibility to NamSor_v1 Gender Scale M[-1..U..+1]F value | [optional] 
**Score** | **float64** | Compatibility to NamSor_v1 Gender score value. Higher score is better, but score is not normalized. Use calibratedProbability if available.  | [optional] 
**ProbabilityCalibrated** | **float64** | The calibrated probability for inferred gender to have been guessed correctly. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


