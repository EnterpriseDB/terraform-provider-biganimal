# AllowedIpRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CidrBlock** | **string** |  | 
**Description** | **string** |  | 

## Methods

### NewAllowedIpRange

`func NewAllowedIpRange(cidrBlock string, description string, ) *AllowedIpRange`

NewAllowedIpRange instantiates a new AllowedIpRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllowedIpRangeWithDefaults

`func NewAllowedIpRangeWithDefaults() *AllowedIpRange`

NewAllowedIpRangeWithDefaults instantiates a new AllowedIpRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidrBlock

`func (o *AllowedIpRange) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *AllowedIpRange) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *AllowedIpRange) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.


### GetDescription

`func (o *AllowedIpRange) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AllowedIpRange) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AllowedIpRange) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


