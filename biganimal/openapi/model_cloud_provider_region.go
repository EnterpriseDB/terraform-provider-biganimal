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

// CloudProviderRegion struct for CloudProviderRegion
type CloudProviderRegion struct {
	RegionId string `json:"regionId"`
	RegionName string `json:"regionName"`
	Status *string `json:"status,omitempty"`
	Continent *string `json:"continent,omitempty"`
}

// NewCloudProviderRegion instantiates a new CloudProviderRegion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudProviderRegion(regionId string, regionName string) *CloudProviderRegion {
	this := CloudProviderRegion{}
	this.RegionId = regionId
	this.RegionName = regionName
	return &this
}

// NewCloudProviderRegionWithDefaults instantiates a new CloudProviderRegion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudProviderRegionWithDefaults() *CloudProviderRegion {
	this := CloudProviderRegion{}
	return &this
}

// GetRegionId returns the RegionId field value
func (o *CloudProviderRegion) GetRegionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegionId
}

// GetRegionIdOk returns a tuple with the RegionId field value
// and a boolean to check if the value has been set.
func (o *CloudProviderRegion) GetRegionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegionId, true
}

// SetRegionId sets field value
func (o *CloudProviderRegion) SetRegionId(v string) {
	o.RegionId = v
}

// GetRegionName returns the RegionName field value
func (o *CloudProviderRegion) GetRegionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value
// and a boolean to check if the value has been set.
func (o *CloudProviderRegion) GetRegionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegionName, true
}

// SetRegionName sets field value
func (o *CloudProviderRegion) SetRegionName(v string) {
	o.RegionName = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CloudProviderRegion) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderRegion) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CloudProviderRegion) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CloudProviderRegion) SetStatus(v string) {
	o.Status = &v
}

// GetContinent returns the Continent field value if set, zero value otherwise.
func (o *CloudProviderRegion) GetContinent() string {
	if o == nil || o.Continent == nil {
		var ret string
		return ret
	}
	return *o.Continent
}

// GetContinentOk returns a tuple with the Continent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderRegion) GetContinentOk() (*string, bool) {
	if o == nil || o.Continent == nil {
		return nil, false
	}
	return o.Continent, true
}

// HasContinent returns a boolean if a field has been set.
func (o *CloudProviderRegion) HasContinent() bool {
	if o != nil && o.Continent != nil {
		return true
	}

	return false
}

// SetContinent gets a reference to the given string and assigns it to the Continent field.
func (o *CloudProviderRegion) SetContinent(v string) {
	o.Continent = &v
}

func (o CloudProviderRegion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["regionId"] = o.RegionId
	}
	if true {
		toSerialize["regionName"] = o.RegionName
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Continent != nil {
		toSerialize["continent"] = o.Continent
	}
	return json.Marshal(toSerialize)
}

type NullableCloudProviderRegion struct {
	value *CloudProviderRegion
	isSet bool
}

func (v NullableCloudProviderRegion) Get() *CloudProviderRegion {
	return v.value
}

func (v *NullableCloudProviderRegion) Set(val *CloudProviderRegion) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudProviderRegion) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudProviderRegion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudProviderRegion(val *CloudProviderRegion) *NullableCloudProviderRegion {
	return &NullableCloudProviderRegion{value: val, isSet: true}
}

func (v NullableCloudProviderRegion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudProviderRegion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


