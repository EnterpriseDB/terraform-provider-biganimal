# Error400

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **float32** |  |
**Message** | **string** |  |
**Errors** | Pointer to [**[]Error400ErrorsInner**](Error400ErrorsInner.md) |  | [optional]
**Reference** | **string** |  |
**Source** | Pointer to **string** |  | [optional]

## Methods

### NewError400

`func NewError400(status float32, message string, reference string, ) *Error400`

NewError400 instantiates a new Error400 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewError400WithDefaults

`func NewError400WithDefaults() *Error400`

NewError400WithDefaults instantiates a new Error400 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Error400) GetStatus() float32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Error400) GetStatusOk() (*float32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Error400) SetStatus(v float32)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *Error400) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Error400) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Error400) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetErrors

`func (o *Error400) GetErrors() []Error400ErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Error400) GetErrorsOk() (*[]Error400ErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Error400) SetErrors(v []Error400ErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *Error400) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetReference

`func (o *Error400) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *Error400) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *Error400) SetReference(v string)`

SetReference sets Reference field to given value.


### GetSource

`func (o *Error400) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Error400) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Error400) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *Error400) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


