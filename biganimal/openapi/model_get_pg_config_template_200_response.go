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

// GetPgConfigTemplate200Response struct for GetPgConfigTemplate200Response
type GetPgConfigTemplate200Response struct {
	Data PgConfigTemplate `json:"data"`
}

// NewGetPgConfigTemplate200Response instantiates a new GetPgConfigTemplate200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPgConfigTemplate200Response(data PgConfigTemplate) *GetPgConfigTemplate200Response {
	this := GetPgConfigTemplate200Response{}
	this.Data = data
	return &this
}

// NewGetPgConfigTemplate200ResponseWithDefaults instantiates a new GetPgConfigTemplate200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPgConfigTemplate200ResponseWithDefaults() *GetPgConfigTemplate200Response {
	this := GetPgConfigTemplate200Response{}
	return &this
}

// GetData returns the Data field value
func (o *GetPgConfigTemplate200Response) GetData() PgConfigTemplate {
	if o == nil {
		var ret PgConfigTemplate
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetPgConfigTemplate200Response) GetDataOk() (*PgConfigTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetPgConfigTemplate200Response) SetData(v PgConfigTemplate) {
	o.Data = v
}

func (o GetPgConfigTemplate200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetPgConfigTemplate200Response struct {
	value *GetPgConfigTemplate200Response
	isSet bool
}

func (v NullableGetPgConfigTemplate200Response) Get() *GetPgConfigTemplate200Response {
	return v.value
}

func (v *NullableGetPgConfigTemplate200Response) Set(val *GetPgConfigTemplate200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPgConfigTemplate200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPgConfigTemplate200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPgConfigTemplate200Response(val *GetPgConfigTemplate200Response) *NullableGetPgConfigTemplate200Response {
	return &NullableGetPgConfigTemplate200Response{value: val, isSet: true}
}

func (v NullableGetPgConfigTemplate200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPgConfigTemplate200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


