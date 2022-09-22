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

// GetAccountNotifications200Response struct for GetAccountNotifications200Response
type GetAccountNotifications200Response struct {
	Data []AccountNotification `json:"data"`
}

// NewGetAccountNotifications200Response instantiates a new GetAccountNotifications200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccountNotifications200Response(data []AccountNotification) *GetAccountNotifications200Response {
	this := GetAccountNotifications200Response{}
	this.Data = data
	return &this
}

// NewGetAccountNotifications200ResponseWithDefaults instantiates a new GetAccountNotifications200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccountNotifications200ResponseWithDefaults() *GetAccountNotifications200Response {
	this := GetAccountNotifications200Response{}
	return &this
}

// GetData returns the Data field value
func (o *GetAccountNotifications200Response) GetData() []AccountNotification {
	if o == nil {
		var ret []AccountNotification
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetAccountNotifications200Response) GetDataOk() ([]AccountNotification, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetAccountNotifications200Response) SetData(v []AccountNotification) {
	o.Data = v
}

func (o GetAccountNotifications200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetAccountNotifications200Response struct {
	value *GetAccountNotifications200Response
	isSet bool
}

func (v NullableGetAccountNotifications200Response) Get() *GetAccountNotifications200Response {
	return v.value
}

func (v *NullableGetAccountNotifications200Response) Set(val *GetAccountNotifications200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccountNotifications200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccountNotifications200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccountNotifications200Response(val *GetAccountNotifications200Response) *NullableGetAccountNotifications200Response {
	return &NullableGetAccountNotifications200Response{value: val, isSet: true}
}

func (v NullableGetAccountNotifications200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccountNotifications200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


