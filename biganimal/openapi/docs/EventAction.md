# EventAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **float32** |  | [optional] 
**Transform** | Pointer to **string** |  | [optional] 

## Methods

### NewEventAction

`func NewEventAction() *EventAction`

NewEventAction instantiates a new EventAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventActionWithDefaults

`func NewEventActionWithDefaults() *EventAction`

NewEventActionWithDefaults instantiates a new EventAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EventAction) GetType() float32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventAction) GetTypeOk() (*float32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventAction) SetType(v float32)`

SetType sets Type field to given value.

### HasType

`func (o *EventAction) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTransform

`func (o *EventAction) GetTransform() string`

GetTransform returns the Transform field if non-nil, zero value otherwise.

### GetTransformOk

`func (o *EventAction) GetTransformOk() (*string, bool)`

GetTransformOk returns a tuple with the Transform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransform

`func (o *EventAction) SetTransform(v string)`

SetTransform sets Transform field to given value.

### HasTransform

`func (o *EventAction) HasTransform() bool`

HasTransform returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


