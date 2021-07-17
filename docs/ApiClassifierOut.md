# ApiClassifierOut

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassifierName** | **string** | The classifier name | [optional] 
**Serving** | **bool** | True if the classifier is serving requests (has reached minimal learning, is not shutting down) | [optional] 
**Learning** | **bool** | True if the classifier is learning | [optional] 
**ShuttingDown** | **bool** | True if the classifier is shutting down | [optional] 
**ProbabilityCalibrated** | **bool** | True if the classifier has finished the initial learning and calibrated probabilities (meanwhile, during initial learning, probabilities will be equal to -1) | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


