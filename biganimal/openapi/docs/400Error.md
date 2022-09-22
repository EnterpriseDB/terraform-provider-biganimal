# 400Error

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **float32** |  | 
**Message** | **string** |  | 
**Errors** | Pointer to [**[]Model400ErrorErrorsInner**](Model400ErrorErrorsInner.md) |  | [optional] 
**Reference** | **string** |  | 
**Source** | Pointer to **string** |  | [optional] 

## Methods

### New400Error

`func New400Error(status float32, message string, reference string, ) *400Error`

New400Error instantiates a new 400Error object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### New400ErrorWithDefaults

`func New400ErrorWithDefaults() *400Error`

New400ErrorWithDefaults instantiates a new 400Error object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *400Error) GetStatus() float32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *400Error) GetStatusOk() (*float32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *400Error) SetStatus(v float32)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *400Error) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *400Error) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *400Error) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetErrors

`func (o *400Error) GetErrors() []Model400ErrorErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *400Error) GetErrorsOk() (*[]Model400ErrorErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *400Error) SetErrors(v []Model400ErrorErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *400Error) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetReference

`func (o *400Error) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *400Error) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *400Error) SetReference(v string)`

SetReference sets Reference field to given value.


### GetSource

`func (o *400Error) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *400Error) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *400Error) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *400Error) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


