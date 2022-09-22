# Usage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Usages** | [**[]ClusterUsage**](ClusterUsage.md) |  | 
**Totals** | [**UsageTotals**](UsageTotals.md) |  | 

## Methods

### NewUsage

`func NewUsage(usages []ClusterUsage, totals UsageTotals, ) *Usage`

NewUsage instantiates a new Usage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageWithDefaults

`func NewUsageWithDefaults() *Usage`

NewUsageWithDefaults instantiates a new Usage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsages

`func (o *Usage) GetUsages() []ClusterUsage`

GetUsages returns the Usages field if non-nil, zero value otherwise.

### GetUsagesOk

`func (o *Usage) GetUsagesOk() (*[]ClusterUsage, bool)`

GetUsagesOk returns a tuple with the Usages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsages

`func (o *Usage) SetUsages(v []ClusterUsage)`

SetUsages sets Usages field to given value.


### GetTotals

`func (o *Usage) GetTotals() UsageTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *Usage) GetTotalsOk() (*UsageTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *Usage) SetTotals(v UsageTotals)`

SetTotals sets Totals field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


