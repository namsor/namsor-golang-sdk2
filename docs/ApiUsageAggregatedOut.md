# ApiUsageAggregatedOut

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeUnit** | **string** | Time unit is DAY, WEEK or MONTH depending on prior usage | [optional] 
**PeriodStart** | **int64** | Start datetime of the reporting period | [optional] 
**PeriodEnd** | **int64** | End datetime of the reporting period | [optional] 
**TotalUsage** | **int64** | Total usage in the current period | [optional] 
**HistoryTruncated** | **bool** | If the history was truncaded due to data limit | [optional] 
**Data** | [**[][]int32**](array.md) | Data points : usage per DAY, WEEK or MONTH and per apiService | [optional] 
**ColHeaders** | **[]string** | apiServices as column headers  | [optional] 
**RowHeaders** | **[]string** | dates as row headers  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


