# ApiKeyOut

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | **string** | The user API Key. | [optional] 
**UserId** | **string** | The user identifier. | [optional] 
**Admin** | **bool** | The API Key has admin role. | [optional] 
**Vetted** | **bool** | The API Key is vetted (assumed truthful) for machine learning. | [optional] 
**Learnable** | **bool** | The API Key is learnable (without assuming truthfulness) for machine learning. Set learnable&#x3D;false and anonymized&#x3D;true for highest privacy (ie. to forget data as it&#39;s processed). | [optional] 
**Anonymized** | **bool** | The API Key is anonymized (using SHA-252 digest for logging). Set learnable&#x3D;false and anonymized&#x3D;true for highest privacy (ie. to forget data as it&#39;s processed). | [optional] 
**Partner** | **bool** | The API Key has partner role. | [optional] 
**Striped** | **bool** | The API Key is associated to a valid Stripe account. | [optional] 
**Corporate** | **bool** | The API Key has role corporate (ex SWIFT payments instead of Stripe payments). | [optional] 
**Disabled** | **bool** | The API Key is temporarily or permanently disabled. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


