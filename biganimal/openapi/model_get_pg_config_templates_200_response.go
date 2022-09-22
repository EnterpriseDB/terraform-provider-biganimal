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

// GetPgConfigTemplates200Response struct for GetPgConfigTemplates200Response
type GetPgConfigTemplates200Response struct {
	Data []PgConfigTemplate `json:"data"`
}

// NewGetPgConfigTemplates200Response instantiates a new GetPgConfigTemplates200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPgConfigTemplates200Response(data []PgConfigTemplate) *GetPgConfigTemplates200Response {
	this := GetPgConfigTemplates200Response{}
	this.Data = data
	return &this
}

// NewGetPgConfigTemplates200ResponseWithDefaults instantiates a new GetPgConfigTemplates200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPgConfigTemplates200ResponseWithDefaults() *GetPgConfigTemplates200Response {
	this := GetPgConfigTemplates200Response{}
	return &this
}

// GetData returns the Data field value
func (o *GetPgConfigTemplates200Response) GetData() []PgConfigTemplate {
	if o == nil {
		var ret []PgConfigTemplate
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetPgConfigTemplates200Response) GetDataOk() ([]PgConfigTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetPgConfigTemplates200Response) SetData(v []PgConfigTemplate) {
	o.Data = v
}

func (o GetPgConfigTemplates200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetPgConfigTemplates200Response struct {
	value *GetPgConfigTemplates200Response
	isSet bool
}

func (v NullableGetPgConfigTemplates200Response) Get() *GetPgConfigTemplates200Response {
	return v.value
}

func (v *NullableGetPgConfigTemplates200Response) Set(val *GetPgConfigTemplates200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPgConfigTemplates200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPgConfigTemplates200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPgConfigTemplates200Response(val *GetPgConfigTemplates200Response) *NullableGetPgConfigTemplates200Response {
	return &NullableGetPgConfigTemplates200Response{value: val, isSet: true}
}

func (v NullableGetPgConfigTemplates200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPgConfigTemplates200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


