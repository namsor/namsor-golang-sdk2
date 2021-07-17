# PersonalNameParsedOut

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Script** | **string** |  | [optional] 
**Id** | **string** |  | [optional] 
**Name** | **string** | The input name | [optional] 
**NameParserType** | **string** | Name parsing is addressed as a classification problem, for example FN1LN1 means a first then last name order. | [optional] 
**NameParserTypeAlt** | **string** | Second best alternative parsing. Name parsing is addressed as a classification problem, for example FN1LN1 means a first then last name order. | [optional] 
**FirstLastName** | [**FirstLastNameOut**](FirstLastNameOut.md) |  | [optional] 
**Score** | **float64** | Higher score is better, but score is not normalized. Use calibratedProbability if available.  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


