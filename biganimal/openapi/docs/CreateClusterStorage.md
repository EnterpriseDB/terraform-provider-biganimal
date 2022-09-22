# CreateClusterStorage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VolumePropertiesId** | **string** |  | 
**VolumeTypeId** | **string** |  | 
**Iops** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **string** |  | [optional] 
**Throughput** | Pointer to **string** | Unused | [optional] 

## Methods

### NewCreateClusterStorage

`func NewCreateClusterStorage(volumePropertiesId string, volumeTypeId string, ) *CreateClusterStorage`

NewCreateClusterStorage instantiates a new CreateClusterStorage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClusterStorageWithDefaults

`func NewCreateClusterStorageWithDefaults() *CreateClusterStorage`

NewCreateClusterStorageWithDefaults instantiates a new CreateClusterStorage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolumePropertiesId

`func (o *CreateClusterStorage) GetVolumePropertiesId() string`

GetVolumePropertiesId returns the VolumePropertiesId field if non-nil, zero value otherwise.

### GetVolumePropertiesIdOk

`func (o *CreateClusterStorage) GetVolumePropertiesIdOk() (*string, bool)`

GetVolumePropertiesIdOk returns a tuple with the VolumePropertiesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumePropertiesId

`func (o *CreateClusterStorage) SetVolumePropertiesId(v string)`

SetVolumePropertiesId sets VolumePropertiesId field to given value.


### GetVolumeTypeId

`func (o *CreateClusterStorage) GetVolumeTypeId() string`

GetVolumeTypeId returns the VolumeTypeId field if non-nil, zero value otherwise.

### GetVolumeTypeIdOk

`func (o *CreateClusterStorage) GetVolumeTypeIdOk() (*string, bool)`

GetVolumeTypeIdOk returns a tuple with the VolumeTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeTypeId

`func (o *CreateClusterStorage) SetVolumeTypeId(v string)`

SetVolumeTypeId sets VolumeTypeId field to given value.


### GetIops

`func (o *CreateClusterStorage) GetIops() string`

GetIops returns the Iops field if non-nil, zero value otherwise.

### GetIopsOk

`func (o *CreateClusterStorage) GetIopsOk() (*string, bool)`

GetIopsOk returns a tuple with the Iops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIops

`func (o *CreateClusterStorage) SetIops(v string)`

SetIops sets Iops field to given value.

### HasIops

`func (o *CreateClusterStorage) HasIops() bool`

HasIops returns a boolean if a field has been set.

### GetSize

`func (o *CreateClusterStorage) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateClusterStorage) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateClusterStorage) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *CreateClusterStorage) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetThroughput

`func (o *CreateClusterStorage) GetThroughput() string`

GetThroughput returns the Throughput field if non-nil, zero value otherwise.

### GetThroughputOk

`func (o *CreateClusterStorage) GetThroughputOk() (*string, bool)`

GetThroughputOk returns a tuple with the Throughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughput

`func (o *CreateClusterStorage) SetThroughput(v string)`

SetThroughput sets Throughput field to given value.

### HasThroughput

`func (o *CreateClusterStorage) HasThroughput() bool`

HasThroughput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


