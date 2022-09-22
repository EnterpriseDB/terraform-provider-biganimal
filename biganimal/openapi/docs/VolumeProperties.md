# VolumeProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VolumePropertiesId** | **string** |  | 
**VolumePropertiesName** | **string** |  | 
**MinDiskSize** | Pointer to **string** |  | [optional] 
**MaxDiskSize** | Pointer to **string** |  | [optional] 
**MinIops** | Pointer to **string** |  | [optional] 
**MaxIops** | Pointer to **string** |  | [optional] 
**DefaultIops** | Pointer to **string** |  | [optional] 
**MinThroughput** | Pointer to **string** |  | [optional] 
**MaxThroughput** | Pointer to **string** |  | [optional] 
**DefaultThroughput** | Pointer to **string** |  | [optional] 

## Methods

### NewVolumeProperties

`func NewVolumeProperties(volumePropertiesId string, volumePropertiesName string, ) *VolumeProperties`

NewVolumeProperties instantiates a new VolumeProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumePropertiesWithDefaults

`func NewVolumePropertiesWithDefaults() *VolumeProperties`

NewVolumePropertiesWithDefaults instantiates a new VolumeProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolumePropertiesId

`func (o *VolumeProperties) GetVolumePropertiesId() string`

GetVolumePropertiesId returns the VolumePropertiesId field if non-nil, zero value otherwise.

### GetVolumePropertiesIdOk

`func (o *VolumeProperties) GetVolumePropertiesIdOk() (*string, bool)`

GetVolumePropertiesIdOk returns a tuple with the VolumePropertiesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumePropertiesId

`func (o *VolumeProperties) SetVolumePropertiesId(v string)`

SetVolumePropertiesId sets VolumePropertiesId field to given value.


### GetVolumePropertiesName

`func (o *VolumeProperties) GetVolumePropertiesName() string`

GetVolumePropertiesName returns the VolumePropertiesName field if non-nil, zero value otherwise.

### GetVolumePropertiesNameOk

`func (o *VolumeProperties) GetVolumePropertiesNameOk() (*string, bool)`

GetVolumePropertiesNameOk returns a tuple with the VolumePropertiesName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumePropertiesName

`func (o *VolumeProperties) SetVolumePropertiesName(v string)`

SetVolumePropertiesName sets VolumePropertiesName field to given value.


### GetMinDiskSize

`func (o *VolumeProperties) GetMinDiskSize() string`

GetMinDiskSize returns the MinDiskSize field if non-nil, zero value otherwise.

### GetMinDiskSizeOk

`func (o *VolumeProperties) GetMinDiskSizeOk() (*string, bool)`

GetMinDiskSizeOk returns a tuple with the MinDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDiskSize

`func (o *VolumeProperties) SetMinDiskSize(v string)`

SetMinDiskSize sets MinDiskSize field to given value.

### HasMinDiskSize

`func (o *VolumeProperties) HasMinDiskSize() bool`

HasMinDiskSize returns a boolean if a field has been set.

### GetMaxDiskSize

`func (o *VolumeProperties) GetMaxDiskSize() string`

GetMaxDiskSize returns the MaxDiskSize field if non-nil, zero value otherwise.

### GetMaxDiskSizeOk

`func (o *VolumeProperties) GetMaxDiskSizeOk() (*string, bool)`

GetMaxDiskSizeOk returns a tuple with the MaxDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDiskSize

`func (o *VolumeProperties) SetMaxDiskSize(v string)`

SetMaxDiskSize sets MaxDiskSize field to given value.

### HasMaxDiskSize

`func (o *VolumeProperties) HasMaxDiskSize() bool`

HasMaxDiskSize returns a boolean if a field has been set.

### GetMinIops

`func (o *VolumeProperties) GetMinIops() string`

GetMinIops returns the MinIops field if non-nil, zero value otherwise.

### GetMinIopsOk

`func (o *VolumeProperties) GetMinIopsOk() (*string, bool)`

GetMinIopsOk returns a tuple with the MinIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinIops

`func (o *VolumeProperties) SetMinIops(v string)`

SetMinIops sets MinIops field to given value.

### HasMinIops

`func (o *VolumeProperties) HasMinIops() bool`

HasMinIops returns a boolean if a field has been set.

### GetMaxIops

`func (o *VolumeProperties) GetMaxIops() string`

GetMaxIops returns the MaxIops field if non-nil, zero value otherwise.

### GetMaxIopsOk

`func (o *VolumeProperties) GetMaxIopsOk() (*string, bool)`

GetMaxIopsOk returns a tuple with the MaxIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIops

`func (o *VolumeProperties) SetMaxIops(v string)`

SetMaxIops sets MaxIops field to given value.

### HasMaxIops

`func (o *VolumeProperties) HasMaxIops() bool`

HasMaxIops returns a boolean if a field has been set.

### GetDefaultIops

`func (o *VolumeProperties) GetDefaultIops() string`

GetDefaultIops returns the DefaultIops field if non-nil, zero value otherwise.

### GetDefaultIopsOk

`func (o *VolumeProperties) GetDefaultIopsOk() (*string, bool)`

GetDefaultIopsOk returns a tuple with the DefaultIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultIops

`func (o *VolumeProperties) SetDefaultIops(v string)`

SetDefaultIops sets DefaultIops field to given value.

### HasDefaultIops

`func (o *VolumeProperties) HasDefaultIops() bool`

HasDefaultIops returns a boolean if a field has been set.

### GetMinThroughput

`func (o *VolumeProperties) GetMinThroughput() string`

GetMinThroughput returns the MinThroughput field if non-nil, zero value otherwise.

### GetMinThroughputOk

`func (o *VolumeProperties) GetMinThroughputOk() (*string, bool)`

GetMinThroughputOk returns a tuple with the MinThroughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinThroughput

`func (o *VolumeProperties) SetMinThroughput(v string)`

SetMinThroughput sets MinThroughput field to given value.

### HasMinThroughput

`func (o *VolumeProperties) HasMinThroughput() bool`

HasMinThroughput returns a boolean if a field has been set.

### GetMaxThroughput

`func (o *VolumeProperties) GetMaxThroughput() string`

GetMaxThroughput returns the MaxThroughput field if non-nil, zero value otherwise.

### GetMaxThroughputOk

`func (o *VolumeProperties) GetMaxThroughputOk() (*string, bool)`

GetMaxThroughputOk returns a tuple with the MaxThroughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxThroughput

`func (o *VolumeProperties) SetMaxThroughput(v string)`

SetMaxThroughput sets MaxThroughput field to given value.

### HasMaxThroughput

`func (o *VolumeProperties) HasMaxThroughput() bool`

HasMaxThroughput returns a boolean if a field has been set.

### GetDefaultThroughput

`func (o *VolumeProperties) GetDefaultThroughput() string`

GetDefaultThroughput returns the DefaultThroughput field if non-nil, zero value otherwise.

### GetDefaultThroughputOk

`func (o *VolumeProperties) GetDefaultThroughputOk() (*string, bool)`

GetDefaultThroughputOk returns a tuple with the DefaultThroughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultThroughput

`func (o *VolumeProperties) SetDefaultThroughput(v string)`

SetDefaultThroughput sets DefaultThroughput field to given value.

### HasDefaultThroughput

`func (o *VolumeProperties) HasDefaultThroughput() bool`

HasDefaultThroughput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


