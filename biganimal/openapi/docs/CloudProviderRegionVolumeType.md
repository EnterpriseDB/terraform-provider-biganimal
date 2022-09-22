# CloudProviderRegionVolumeType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VolumeTypeId** | **string** |  | 
**VolumeTypeName** | **string** |  | 
**StorageClass** | **string** |  | 
**EnabledInRegion** | Pointer to **bool** |  | [optional] 
**SupportedInstanceFamilyNames** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCloudProviderRegionVolumeType

`func NewCloudProviderRegionVolumeType(volumeTypeId string, volumeTypeName string, storageClass string, ) *CloudProviderRegionVolumeType`

NewCloudProviderRegionVolumeType instantiates a new CloudProviderRegionVolumeType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudProviderRegionVolumeTypeWithDefaults

`func NewCloudProviderRegionVolumeTypeWithDefaults() *CloudProviderRegionVolumeType`

NewCloudProviderRegionVolumeTypeWithDefaults instantiates a new CloudProviderRegionVolumeType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolumeTypeId

`func (o *CloudProviderRegionVolumeType) GetVolumeTypeId() string`

GetVolumeTypeId returns the VolumeTypeId field if non-nil, zero value otherwise.

### GetVolumeTypeIdOk

`func (o *CloudProviderRegionVolumeType) GetVolumeTypeIdOk() (*string, bool)`

GetVolumeTypeIdOk returns a tuple with the VolumeTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeTypeId

`func (o *CloudProviderRegionVolumeType) SetVolumeTypeId(v string)`

SetVolumeTypeId sets VolumeTypeId field to given value.


### GetVolumeTypeName

`func (o *CloudProviderRegionVolumeType) GetVolumeTypeName() string`

GetVolumeTypeName returns the VolumeTypeName field if non-nil, zero value otherwise.

### GetVolumeTypeNameOk

`func (o *CloudProviderRegionVolumeType) GetVolumeTypeNameOk() (*string, bool)`

GetVolumeTypeNameOk returns a tuple with the VolumeTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeTypeName

`func (o *CloudProviderRegionVolumeType) SetVolumeTypeName(v string)`

SetVolumeTypeName sets VolumeTypeName field to given value.


### GetStorageClass

`func (o *CloudProviderRegionVolumeType) GetStorageClass() string`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *CloudProviderRegionVolumeType) GetStorageClassOk() (*string, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *CloudProviderRegionVolumeType) SetStorageClass(v string)`

SetStorageClass sets StorageClass field to given value.


### GetEnabledInRegion

`func (o *CloudProviderRegionVolumeType) GetEnabledInRegion() bool`

GetEnabledInRegion returns the EnabledInRegion field if non-nil, zero value otherwise.

### GetEnabledInRegionOk

`func (o *CloudProviderRegionVolumeType) GetEnabledInRegionOk() (*bool, bool)`

GetEnabledInRegionOk returns a tuple with the EnabledInRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledInRegion

`func (o *CloudProviderRegionVolumeType) SetEnabledInRegion(v bool)`

SetEnabledInRegion sets EnabledInRegion field to given value.

### HasEnabledInRegion

`func (o *CloudProviderRegionVolumeType) HasEnabledInRegion() bool`

HasEnabledInRegion returns a boolean if a field has been set.

### GetSupportedInstanceFamilyNames

`func (o *CloudProviderRegionVolumeType) GetSupportedInstanceFamilyNames() []string`

GetSupportedInstanceFamilyNames returns the SupportedInstanceFamilyNames field if non-nil, zero value otherwise.

### GetSupportedInstanceFamilyNamesOk

`func (o *CloudProviderRegionVolumeType) GetSupportedInstanceFamilyNamesOk() (*[]string, bool)`

GetSupportedInstanceFamilyNamesOk returns a tuple with the SupportedInstanceFamilyNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedInstanceFamilyNames

`func (o *CloudProviderRegionVolumeType) SetSupportedInstanceFamilyNames(v []string)`

SetSupportedInstanceFamilyNames sets SupportedInstanceFamilyNames field to given value.

### HasSupportedInstanceFamilyNames

`func (o *CloudProviderRegionVolumeType) HasSupportedInstanceFamilyNames() bool`

HasSupportedInstanceFamilyNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


