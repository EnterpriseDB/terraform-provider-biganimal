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

// GetAccountClustersOverview200Response struct for GetAccountClustersOverview200Response
type GetAccountClustersOverview200Response struct {
	Data []ClusterEstate `json:"data"`
}

// NewGetAccountClustersOverview200Response instantiates a new GetAccountClustersOverview200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccountClustersOverview200Response(data []ClusterEstate) *GetAccountClustersOverview200Response {
	this := GetAccountClustersOverview200Response{}
	this.Data = data
	return &this
}

// NewGetAccountClustersOverview200ResponseWithDefaults instantiates a new GetAccountClustersOverview200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccountClustersOverview200ResponseWithDefaults() *GetAccountClustersOverview200Response {
	this := GetAccountClustersOverview200Response{}
	return &this
}

// GetData returns the Data field value
func (o *GetAccountClustersOverview200Response) GetData() []ClusterEstate {
	if o == nil {
		var ret []ClusterEstate
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetAccountClustersOverview200Response) GetDataOk() ([]ClusterEstate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetAccountClustersOverview200Response) SetData(v []ClusterEstate) {
	o.Data = v
}

func (o GetAccountClustersOverview200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetAccountClustersOverview200Response struct {
	value *GetAccountClustersOverview200Response
	isSet bool
}

func (v NullableGetAccountClustersOverview200Response) Get() *GetAccountClustersOverview200Response {
	return v.value
}

func (v *NullableGetAccountClustersOverview200Response) Set(val *GetAccountClustersOverview200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccountClustersOverview200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccountClustersOverview200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccountClustersOverview200Response(val *GetAccountClustersOverview200Response) *NullableGetAccountClustersOverview200Response {
	return &NullableGetAccountClustersOverview200Response{value: val, isSet: true}
}

func (v NullableGetAccountClustersOverview200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccountClustersOverview200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


