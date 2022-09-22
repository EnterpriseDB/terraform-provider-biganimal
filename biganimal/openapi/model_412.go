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

// Model412 struct for Model412
type Model412 struct {
	Error *Error401 `json:"error,omitempty"`
}

// NewModel412 instantiates a new Model412 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModel412() *Model412 {
	this := Model412{}
	return &this
}

// NewModel412WithDefaults instantiates a new Model412 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModel412WithDefaults() *Model412 {
	this := Model412{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *Model412) GetError() Error401 {
	if o == nil || o.Error == nil {
		var ret Error401
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model412) GetErrorOk() (*Error401, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *Model412) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given Error401 and assigns it to the Error field.
func (o *Model412) SetError(v Error401) {
	o.Error = &v
}

func (o Model412) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableModel412 struct {
	value *Model412
	isSet bool
}

func (v NullableModel412) Get() *Model412 {
	return v.value
}

func (v *NullableModel412) Set(val *Model412) {
	v.value = val
	v.isSet = true
}

func (v NullableModel412) IsSet() bool {
	return v.isSet
}

func (v *NullableModel412) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModel412(val *Model412) *NullableModel412 {
	return &NullableModel412{value: val, isSet: true}
}

func (v NullableModel412) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModel412) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


