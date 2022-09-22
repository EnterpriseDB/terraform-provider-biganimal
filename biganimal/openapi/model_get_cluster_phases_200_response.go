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

// GetClusterPhases200Response struct for GetClusterPhases200Response
type GetClusterPhases200Response struct {
	Data []ClusterPhase `json:"data"`
}

// NewGetClusterPhases200Response instantiates a new GetClusterPhases200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetClusterPhases200Response(data []ClusterPhase) *GetClusterPhases200Response {
	this := GetClusterPhases200Response{}
	this.Data = data
	return &this
}

// NewGetClusterPhases200ResponseWithDefaults instantiates a new GetClusterPhases200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetClusterPhases200ResponseWithDefaults() *GetClusterPhases200Response {
	this := GetClusterPhases200Response{}
	return &this
}

// GetData returns the Data field value
func (o *GetClusterPhases200Response) GetData() []ClusterPhase {
	if o == nil {
		var ret []ClusterPhase
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetClusterPhases200Response) GetDataOk() ([]ClusterPhase, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetClusterPhases200Response) SetData(v []ClusterPhase) {
	o.Data = v
}

func (o GetClusterPhases200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetClusterPhases200Response struct {
	value *GetClusterPhases200Response
	isSet bool
}

func (v NullableGetClusterPhases200Response) Get() *GetClusterPhases200Response {
	return v.value
}

func (v *NullableGetClusterPhases200Response) Set(val *GetClusterPhases200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetClusterPhases200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetClusterPhases200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetClusterPhases200Response(val *GetClusterPhases200Response) *NullableGetClusterPhases200Response {
	return &NullableGetClusterPhases200Response{value: val, isSet: true}
}

func (v NullableGetClusterPhases200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetClusterPhases200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


