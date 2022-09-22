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

// RunSearchForEventsRequest struct for RunSearchForEventsRequest
type RunSearchForEventsRequest struct {
	Paging RunSearchForEventsRequestPaging `json:"paging"`
	StartAt string `json:"startAt"`
	EndAt string `json:"endAt"`
	Order *string `json:"order,omitempty"`
	User *RunSearchForEventsRequestUser `json:"user,omitempty"`
	Resource *RunSearchForEventsRequestResource `json:"resource,omitempty"`
}

// NewRunSearchForEventsRequest instantiates a new RunSearchForEventsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunSearchForEventsRequest(paging RunSearchForEventsRequestPaging, startAt string, endAt string) *RunSearchForEventsRequest {
	this := RunSearchForEventsRequest{}
	this.Paging = paging
	this.StartAt = startAt
	this.EndAt = endAt
	return &this
}

// NewRunSearchForEventsRequestWithDefaults instantiates a new RunSearchForEventsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunSearchForEventsRequestWithDefaults() *RunSearchForEventsRequest {
	this := RunSearchForEventsRequest{}
	return &this
}

// GetPaging returns the Paging field value
func (o *RunSearchForEventsRequest) GetPaging() RunSearchForEventsRequestPaging {
	if o == nil {
		var ret RunSearchForEventsRequestPaging
		return ret
	}

	return o.Paging
}

// GetPagingOk returns a tuple with the Paging field value
// and a boolean to check if the value has been set.
func (o *RunSearchForEventsRequest) GetPagingOk() (*RunSearchForEventsRequestPaging, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Paging, true
}

// SetPaging sets field value
func (o *RunSearchForEventsRequest) SetPaging(v RunSearchForEventsRequestPaging) {
	o.Paging = v
}

// GetStartAt returns the StartAt field value
func (o *RunSearchForEventsRequest) GetStartAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartAt
}

// GetStartAtOk returns a tuple with the StartAt field value
// and a boolean to check if the value has been set.
func (o *RunSearchForEventsRequest) GetStartAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartAt, true
}

// SetStartAt sets field value
func (o *RunSearchForEventsRequest) SetStartAt(v string) {
	o.StartAt = v
}

// GetEndAt returns the EndAt field value
func (o *RunSearchForEventsRequest) GetEndAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndAt
}

// GetEndAtOk returns a tuple with the EndAt field value
// and a boolean to check if the value has been set.
func (o *RunSearchForEventsRequest) GetEndAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndAt, true
}

// SetEndAt sets field value
func (o *RunSearchForEventsRequest) SetEndAt(v string) {
	o.EndAt = v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *RunSearchForEventsRequest) GetOrder() string {
	if o == nil || o.Order == nil {
		var ret string
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunSearchForEventsRequest) GetOrderOk() (*string, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *RunSearchForEventsRequest) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given string and assigns it to the Order field.
func (o *RunSearchForEventsRequest) SetOrder(v string) {
	o.Order = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *RunSearchForEventsRequest) GetUser() RunSearchForEventsRequestUser {
	if o == nil || o.User == nil {
		var ret RunSearchForEventsRequestUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunSearchForEventsRequest) GetUserOk() (*RunSearchForEventsRequestUser, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *RunSearchForEventsRequest) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given RunSearchForEventsRequestUser and assigns it to the User field.
func (o *RunSearchForEventsRequest) SetUser(v RunSearchForEventsRequestUser) {
	o.User = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *RunSearchForEventsRequest) GetResource() RunSearchForEventsRequestResource {
	if o == nil || o.Resource == nil {
		var ret RunSearchForEventsRequestResource
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunSearchForEventsRequest) GetResourceOk() (*RunSearchForEventsRequestResource, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *RunSearchForEventsRequest) HasResource() bool {
	if o != nil && o.Resource != nil {
		return true
	}

	return false
}

// SetResource gets a reference to the given RunSearchForEventsRequestResource and assigns it to the Resource field.
func (o *RunSearchForEventsRequest) SetResource(v RunSearchForEventsRequestResource) {
	o.Resource = &v
}

func (o RunSearchForEventsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["paging"] = o.Paging
	}
	if true {
		toSerialize["startAt"] = o.StartAt
	}
	if true {
		toSerialize["endAt"] = o.EndAt
	}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Resource != nil {
		toSerialize["resource"] = o.Resource
	}
	return json.Marshal(toSerialize)
}

type NullableRunSearchForEventsRequest struct {
	value *RunSearchForEventsRequest
	isSet bool
}

func (v NullableRunSearchForEventsRequest) Get() *RunSearchForEventsRequest {
	return v.value
}

func (v *NullableRunSearchForEventsRequest) Set(val *RunSearchForEventsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRunSearchForEventsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRunSearchForEventsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunSearchForEventsRequest(val *RunSearchForEventsRequest) *NullableRunSearchForEventsRequest {
	return &NullableRunSearchForEventsRequest{value: val, isSet: true}
}

func (v NullableRunSearchForEventsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunSearchForEventsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


