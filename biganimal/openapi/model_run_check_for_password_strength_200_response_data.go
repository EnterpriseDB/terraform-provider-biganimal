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

// RunCheckForPasswordStrength200ResponseData struct for RunCheckForPasswordStrength200ResponseData
type RunCheckForPasswordStrength200ResponseData struct {
	Score float32 `json:"score"`
	Feedback RunCheckForPasswordStrength200ResponseDataFeedback `json:"feedback"`
}

// NewRunCheckForPasswordStrength200ResponseData instantiates a new RunCheckForPasswordStrength200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunCheckForPasswordStrength200ResponseData(score float32, feedback RunCheckForPasswordStrength200ResponseDataFeedback) *RunCheckForPasswordStrength200ResponseData {
	this := RunCheckForPasswordStrength200ResponseData{}
	this.Score = score
	this.Feedback = feedback
	return &this
}

// NewRunCheckForPasswordStrength200ResponseDataWithDefaults instantiates a new RunCheckForPasswordStrength200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunCheckForPasswordStrength200ResponseDataWithDefaults() *RunCheckForPasswordStrength200ResponseData {
	this := RunCheckForPasswordStrength200ResponseData{}
	return &this
}

// GetScore returns the Score field value
func (o *RunCheckForPasswordStrength200ResponseData) GetScore() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Score
}

// GetScoreOk returns a tuple with the Score field value
// and a boolean to check if the value has been set.
func (o *RunCheckForPasswordStrength200ResponseData) GetScoreOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Score, true
}

// SetScore sets field value
func (o *RunCheckForPasswordStrength200ResponseData) SetScore(v float32) {
	o.Score = v
}

// GetFeedback returns the Feedback field value
func (o *RunCheckForPasswordStrength200ResponseData) GetFeedback() RunCheckForPasswordStrength200ResponseDataFeedback {
	if o == nil {
		var ret RunCheckForPasswordStrength200ResponseDataFeedback
		return ret
	}

	return o.Feedback
}

// GetFeedbackOk returns a tuple with the Feedback field value
// and a boolean to check if the value has been set.
func (o *RunCheckForPasswordStrength200ResponseData) GetFeedbackOk() (*RunCheckForPasswordStrength200ResponseDataFeedback, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Feedback, true
}

// SetFeedback sets field value
func (o *RunCheckForPasswordStrength200ResponseData) SetFeedback(v RunCheckForPasswordStrength200ResponseDataFeedback) {
	o.Feedback = v
}

func (o RunCheckForPasswordStrength200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["score"] = o.Score
	}
	if true {
		toSerialize["feedback"] = o.Feedback
	}
	return json.Marshal(toSerialize)
}

type NullableRunCheckForPasswordStrength200ResponseData struct {
	value *RunCheckForPasswordStrength200ResponseData
	isSet bool
}

func (v NullableRunCheckForPasswordStrength200ResponseData) Get() *RunCheckForPasswordStrength200ResponseData {
	return v.value
}

func (v *NullableRunCheckForPasswordStrength200ResponseData) Set(val *RunCheckForPasswordStrength200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableRunCheckForPasswordStrength200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableRunCheckForPasswordStrength200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunCheckForPasswordStrength200ResponseData(val *RunCheckForPasswordStrength200ResponseData) *NullableRunCheckForPasswordStrength200ResponseData {
	return &NullableRunCheckForPasswordStrength200ResponseData{value: val, isSet: true}
}

func (v NullableRunCheckForPasswordStrength200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunCheckForPasswordStrength200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


