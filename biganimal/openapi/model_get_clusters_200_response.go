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

// GetClusters200Response struct for GetClusters200Response
type GetClusters200Response struct {
	Data []ClusterDetail `json:"data"`
}

// NewGetClusters200Response instantiates a new GetClusters200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetClusters200Response(data []ClusterDetail) *GetClusters200Response {
	this := GetClusters200Response{}
	this.Data = data
	return &this
}

// NewGetClusters200ResponseWithDefaults instantiates a new GetClusters200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetClusters200ResponseWithDefaults() *GetClusters200Response {
	this := GetClusters200Response{}
	return &this
}

// GetData returns the Data field value
func (o *GetClusters200Response) GetData() []ClusterDetail {
	if o == nil {
		var ret []ClusterDetail
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetClusters200Response) GetDataOk() ([]ClusterDetail, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetClusters200Response) SetData(v []ClusterDetail) {
	o.Data = v
}

func (o GetClusters200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetClusters200Response struct {
	value *GetClusters200Response
	isSet bool
}

func (v NullableGetClusters200Response) Get() *GetClusters200Response {
	return v.value
}

func (v *NullableGetClusters200Response) Set(val *GetClusters200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetClusters200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetClusters200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetClusters200Response(val *GetClusters200Response) *NullableGetClusters200Response {
	return &NullableGetClusters200Response{value: val, isSet: true}
}

func (v NullableGetClusters200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetClusters200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


