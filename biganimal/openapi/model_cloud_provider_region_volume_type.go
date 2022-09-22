/*
BigAnimal

BigAnimal REST API v2 <br /><br /> Please visit [API v2 Changelog page](/api/docs/v2migration.html) for information about migrating from API v1. 

API version: 2.5.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// CloudProviderRegionVolumeType struct for CloudProviderRegionVolumeType
type CloudProviderRegionVolumeType struct {
	VolumeTypeId string `json:"volumeTypeId"`
	VolumeTypeName string `json:"volumeTypeName"`
	StorageClass string `json:"storageClass"`
	EnabledInRegion *bool `json:"enabledInRegion,omitempty"`
	SupportedInstanceFamilyNames []string `json:"supportedInstanceFamilyNames,omitempty"`
}

// NewCloudProviderRegionVolumeType instantiates a new CloudProviderRegionVolumeType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudProviderRegionVolumeType(volumeTypeId string, volumeTypeName string, storageClass string) *CloudProviderRegionVolumeType {
	this := CloudProviderRegionVolumeType{}
	this.VolumeTypeId = volumeTypeId
	this.VolumeTypeName = volumeTypeName
	this.StorageClass = storageClass
	return &this
}

// NewCloudProviderRegionVolumeTypeWithDefaults instantiates a new CloudProviderRegionVolumeType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudProviderRegionVolumeTypeWithDefaults() *CloudProviderRegionVolumeType {
	this := CloudProviderRegionVolumeType{}
	return &this
}

// GetVolumeTypeId returns the VolumeTypeId field value
func (o *CloudProviderRegionVolumeType) GetVolumeTypeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VolumeTypeId
}

// GetVolumeTypeIdOk returns a tuple with the VolumeTypeId field value
// and a boolean to check if the value has been set.
func (o *CloudProviderRegionVolumeType) GetVolumeTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VolumeTypeId, true
}

// SetVolumeTypeId sets field value
func (o *CloudProviderRegionVolumeType) SetVolumeTypeId(v string) {
	o.VolumeTypeId = v
}

// GetVolumeTypeName returns the VolumeTypeName field value
func (o *CloudProviderRegionVolumeType) GetVolumeTypeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VolumeTypeName
}

// GetVolumeTypeNameOk returns a tuple with the VolumeTypeName field value
// and a boolean to check if the value has been set.
func (o *CloudProviderRegionVolumeType) GetVolumeTypeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VolumeTypeName, true
}

// SetVolumeTypeName sets field value
func (o *CloudProviderRegionVolumeType) SetVolumeTypeName(v string) {
	o.VolumeTypeName = v
}

// GetStorageClass returns the StorageClass field value
func (o *CloudProviderRegionVolumeType) GetStorageClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StorageClass
}

// GetStorageClassOk returns a tuple with the StorageClass field value
// and a boolean to check if the value has been set.
func (o *CloudProviderRegionVolumeType) GetStorageClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageClass, true
}

// SetStorageClass sets field value
func (o *CloudProviderRegionVolumeType) SetStorageClass(v string) {
	o.StorageClass = v
}

// GetEnabledInRegion returns the EnabledInRegion field value if set, zero value otherwise.
func (o *CloudProviderRegionVolumeType) GetEnabledInRegion() bool {
	if o == nil || o.EnabledInRegion == nil {
		var ret bool
		return ret
	}
	return *o.EnabledInRegion
}

// GetEnabledInRegionOk returns a tuple with the EnabledInRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderRegionVolumeType) GetEnabledInRegionOk() (*bool, bool) {
	if o == nil || o.EnabledInRegion == nil {
		return nil, false
	}
	return o.EnabledInRegion, true
}

// HasEnabledInRegion returns a boolean if a field has been set.
func (o *CloudProviderRegionVolumeType) HasEnabledInRegion() bool {
	if o != nil && o.EnabledInRegion != nil {
		return true
	}

	return false
}

// SetEnabledInRegion gets a reference to the given bool and assigns it to the EnabledInRegion field.
func (o *CloudProviderRegionVolumeType) SetEnabledInRegion(v bool) {
	o.EnabledInRegion = &v
}

// GetSupportedInstanceFamilyNames returns the SupportedInstanceFamilyNames field value if set, zero value otherwise.
func (o *CloudProviderRegionVolumeType) GetSupportedInstanceFamilyNames() []string {
	if o == nil || o.SupportedInstanceFamilyNames == nil {
		var ret []string
		return ret
	}
	return o.SupportedInstanceFamilyNames
}

// GetSupportedInstanceFamilyNamesOk returns a tuple with the SupportedInstanceFamilyNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderRegionVolumeType) GetSupportedInstanceFamilyNamesOk() ([]string, bool) {
	if o == nil || o.SupportedInstanceFamilyNames == nil {
		return nil, false
	}
	return o.SupportedInstanceFamilyNames, true
}

// HasSupportedInstanceFamilyNames returns a boolean if a field has been set.
func (o *CloudProviderRegionVolumeType) HasSupportedInstanceFamilyNames() bool {
	if o != nil && o.SupportedInstanceFamilyNames != nil {
		return true
	}

	return false
}

// SetSupportedInstanceFamilyNames gets a reference to the given []string and assigns it to the SupportedInstanceFamilyNames field.
func (o *CloudProviderRegionVolumeType) SetSupportedInstanceFamilyNames(v []string) {
	o.SupportedInstanceFamilyNames = v
}

func (o CloudProviderRegionVolumeType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["volumeTypeId"] = o.VolumeTypeId
	}
	if true {
		toSerialize["volumeTypeName"] = o.VolumeTypeName
	}
	if true {
		toSerialize["storageClass"] = o.StorageClass
	}
	if o.EnabledInRegion != nil {
		toSerialize["enabledInRegion"] = o.EnabledInRegion
	}
	if o.SupportedInstanceFamilyNames != nil {
		toSerialize["supportedInstanceFamilyNames"] = o.SupportedInstanceFamilyNames
	}
	return json.Marshal(toSerialize)
}

type NullableCloudProviderRegionVolumeType struct {
	value *CloudProviderRegionVolumeType
	isSet bool
}

func (v NullableCloudProviderRegionVolumeType) Get() *CloudProviderRegionVolumeType {
	return v.value
}

func (v *NullableCloudProviderRegionVolumeType) Set(val *CloudProviderRegionVolumeType) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudProviderRegionVolumeType) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudProviderRegionVolumeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudProviderRegionVolumeType(val *CloudProviderRegionVolumeType) *NullableCloudProviderRegionVolumeType {
	return &NullableCloudProviderRegionVolumeType{value: val, isSet: true}
}

func (v NullableCloudProviderRegionVolumeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudProviderRegionVolumeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


