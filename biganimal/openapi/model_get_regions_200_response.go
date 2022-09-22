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

// GetRegions200Response struct for GetRegions200Response
type GetRegions200Response struct {
	Data []Region `json:"data"`
}

// NewGetRegions200Response instantiates a new GetRegions200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRegions200Response(data []Region) *GetRegions200Response {
	this := GetRegions200Response{}
	this.Data = data
	return &this
}

// NewGetRegions200ResponseWithDefaults instantiates a new GetRegions200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRegions200ResponseWithDefaults() *GetRegions200Response {
	this := GetRegions200Response{}
	return &this
}

// GetData returns the Data field value
func (o *GetRegions200Response) GetData() []Region {
	if o == nil {
		var ret []Region
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetRegions200Response) GetDataOk() ([]Region, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetRegions200Response) SetData(v []Region) {
	o.Data = v
}

func (o GetRegions200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetRegions200Response struct {
	value *GetRegions200Response
	isSet bool
}

func (v NullableGetRegions200Response) Get() *GetRegions200Response {
	return v.value
}

func (v *NullableGetRegions200Response) Set(val *GetRegions200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRegions200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRegions200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRegions200Response(val *GetRegions200Response) *NullableGetRegions200Response {
	return &NullableGetRegions200Response{value: val, isSet: true}
}

func (v NullableGetRegions200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRegions200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


