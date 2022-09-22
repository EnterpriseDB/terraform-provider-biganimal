/*
BigAnimal

BigAnimal REST API v2 <br /><br /> Please visit [API v2 Changelog page](/api/docs/v2migration.html) for information about migrating from API v1.

API version: 2.5.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// CreateCloudProviderRequest - struct for CreateCloudProviderRequest
type CreateCloudProviderRequest struct {
	RegisterAws *RegisterAws
	RegisterAzure *RegisterAzure
}

// RegisterAwsAsCreateCloudProviderRequest is a convenience function that returns RegisterAws wrapped in CreateCloudProviderRequest
func RegisterAwsAsCreateCloudProviderRequest(v *RegisterAws) CreateCloudProviderRequest {
	return CreateCloudProviderRequest{
		RegisterAws: v,
	}
}

// RegisterAzureAsCreateCloudProviderRequest is a convenience function that returns RegisterAzure wrapped in CreateCloudProviderRequest
func RegisterAzureAsCreateCloudProviderRequest(v *RegisterAzure) CreateCloudProviderRequest {
	return CreateCloudProviderRequest{
		RegisterAzure: v,
	}
}

// A wrapper for strict JSON decoding
func newStrictDecoder(data []byte) *json.Decoder {
	dec := json.NewDecoder(bytes.NewBuffer(data))
	dec.DisallowUnknownFields()
	return dec
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateCloudProviderRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into RegisterAws
	err = newStrictDecoder(data).Decode(&dst.RegisterAws)
	if err == nil {
		jsonRegisterAws, _ := json.Marshal(dst.RegisterAws)
		if string(jsonRegisterAws) == "{}" { // empty struct
			dst.RegisterAws = nil
		} else {
			match++
		}
	} else {
		dst.RegisterAws = nil
	}

	// try to unmarshal data into RegisterAzure
	err = newStrictDecoder(data).Decode(&dst.RegisterAzure)
	if err == nil {
		jsonRegisterAzure, _ := json.Marshal(dst.RegisterAzure)
		if string(jsonRegisterAzure) == "{}" { // empty struct
			dst.RegisterAzure = nil
		} else {
			match++
		}
	} else {
		dst.RegisterAzure = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.RegisterAws = nil
		dst.RegisterAzure = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(CreateCloudProviderRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(CreateCloudProviderRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateCloudProviderRequest) MarshalJSON() ([]byte, error) {
	if src.RegisterAws != nil {
		return json.Marshal(&src.RegisterAws)
	}

	if src.RegisterAzure != nil {
		return json.Marshal(&src.RegisterAzure)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateCloudProviderRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.RegisterAws != nil {
		return obj.RegisterAws
	}

	if obj.RegisterAzure != nil {
		return obj.RegisterAzure
	}

	// all schemas are nil
	return nil
}

type NullableCreateCloudProviderRequest struct {
	value *CreateCloudProviderRequest
	isSet bool
}

func (v NullableCreateCloudProviderRequest) Get() *CreateCloudProviderRequest {
	return v.value
}

func (v *NullableCreateCloudProviderRequest) Set(val *CreateCloudProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCloudProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCloudProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCloudProviderRequest(val *CreateCloudProviderRequest) *NullableCreateCloudProviderRequest {
	return &NullableCreateCloudProviderRequest{value: val, isSet: true}
}

func (v NullableCreateCloudProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCloudProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


