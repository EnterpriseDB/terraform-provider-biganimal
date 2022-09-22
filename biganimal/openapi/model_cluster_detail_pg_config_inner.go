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

// ClusterDetailPgConfigInner struct for ClusterDetailPgConfigInner
type ClusterDetailPgConfigInner struct {
	Name *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewClusterDetailPgConfigInner instantiates a new ClusterDetailPgConfigInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterDetailPgConfigInner() *ClusterDetailPgConfigInner {
	this := ClusterDetailPgConfigInner{}
	return &this
}

// NewClusterDetailPgConfigInnerWithDefaults instantiates a new ClusterDetailPgConfigInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterDetailPgConfigInnerWithDefaults() *ClusterDetailPgConfigInner {
	this := ClusterDetailPgConfigInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ClusterDetailPgConfigInner) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDetailPgConfigInner) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ClusterDetailPgConfigInner) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ClusterDetailPgConfigInner) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ClusterDetailPgConfigInner) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDetailPgConfigInner) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ClusterDetailPgConfigInner) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ClusterDetailPgConfigInner) SetValue(v string) {
	o.Value = &v
}

func (o ClusterDetailPgConfigInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableClusterDetailPgConfigInner struct {
	value *ClusterDetailPgConfigInner
	isSet bool
}

func (v NullableClusterDetailPgConfigInner) Get() *ClusterDetailPgConfigInner {
	return v.value
}

func (v *NullableClusterDetailPgConfigInner) Set(val *ClusterDetailPgConfigInner) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterDetailPgConfigInner) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterDetailPgConfigInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterDetailPgConfigInner(val *ClusterDetailPgConfigInner) *NullableClusterDetailPgConfigInner {
	return &NullableClusterDetailPgConfigInner{value: val, isSet: true}
}

func (v NullableClusterDetailPgConfigInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterDetailPgConfigInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


