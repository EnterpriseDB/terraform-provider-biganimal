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

// ListPermissions200Response struct for ListPermissions200Response
type ListPermissions200Response struct {
	Data []string `json:"data"`
}

// NewListPermissions200Response instantiates a new ListPermissions200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPermissions200Response(data []string) *ListPermissions200Response {
	this := ListPermissions200Response{}
	this.Data = data
	return &this
}

// NewListPermissions200ResponseWithDefaults instantiates a new ListPermissions200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPermissions200ResponseWithDefaults() *ListPermissions200Response {
	this := ListPermissions200Response{}
	return &this
}

// GetData returns the Data field value
func (o *ListPermissions200Response) GetData() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListPermissions200Response) GetDataOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListPermissions200Response) SetData(v []string) {
	o.Data = v
}

func (o ListPermissions200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableListPermissions200Response struct {
	value *ListPermissions200Response
	isSet bool
}

func (v NullableListPermissions200Response) Get() *ListPermissions200Response {
	return v.value
}

func (v *NullableListPermissions200Response) Set(val *ListPermissions200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListPermissions200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListPermissions200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPermissions200Response(val *ListPermissions200Response) *NullableListPermissions200Response {
	return &NullableListPermissions200Response{value: val, isSet: true}
}

func (v NullableListPermissions200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPermissions200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


