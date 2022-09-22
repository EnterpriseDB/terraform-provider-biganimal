# RunSearchForEventsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paging** | [**RunSearchForEventsRequestPaging**](RunSearchForEventsRequestPaging.md) |  | 
**StartAt** | **string** |  | 
**EndAt** | **string** |  | 
**Order** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**RunSearchForEventsRequestUser**](RunSearchForEventsRequestUser.md) |  | [optional] 
**Resource** | Pointer to [**RunSearchForEventsRequestResource**](RunSearchForEventsRequestResource.md) |  | [optional] 

## Methods

### NewRunSearchForEventsRequest

`func NewRunSearchForEventsRequest(paging RunSearchForEventsRequestPaging, startAt string, endAt string, ) *RunSearchForEventsRequest`

NewRunSearchForEventsRequest instantiates a new RunSearchForEventsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunSearchForEventsRequestWithDefaults

`func NewRunSearchForEventsRequestWithDefaults() *RunSearchForEventsRequest`

NewRunSearchForEventsRequestWithDefaults instantiates a new RunSearchForEventsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaging

`func (o *RunSearchForEventsRequest) GetPaging() RunSearchForEventsRequestPaging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *RunSearchForEventsRequest) GetPagingOk() (*RunSearchForEventsRequestPaging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *RunSearchForEventsRequest) SetPaging(v RunSearchForEventsRequestPaging)`

SetPaging sets Paging field to given value.


### GetStartAt

`func (o *RunSearchForEventsRequest) GetStartAt() string`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *RunSearchForEventsRequest) GetStartAtOk() (*string, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *RunSearchForEventsRequest) SetStartAt(v string)`

SetStartAt sets StartAt field to given value.


### GetEndAt

`func (o *RunSearchForEventsRequest) GetEndAt() string`

GetEndAt returns the EndAt field if non-nil, zero value otherwise.

### GetEndAtOk

`func (o *RunSearchForEventsRequest) GetEndAtOk() (*string, bool)`

GetEndAtOk returns a tuple with the EndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAt

`func (o *RunSearchForEventsRequest) SetEndAt(v string)`

SetEndAt sets EndAt field to given value.


### GetOrder

`func (o *RunSearchForEventsRequest) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *RunSearchForEventsRequest) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *RunSearchForEventsRequest) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *RunSearchForEventsRequest) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetUser

`func (o *RunSearchForEventsRequest) GetUser() RunSearchForEventsRequestUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *RunSearchForEventsRequest) GetUserOk() (*RunSearchForEventsRequestUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *RunSearchForEventsRequest) SetUser(v RunSearchForEventsRequestUser)`

SetUser sets User field to given value.

### HasUser

`func (o *RunSearchForEventsRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetResource

`func (o *RunSearchForEventsRequest) GetResource() RunSearchForEventsRequestResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *RunSearchForEventsRequest) GetResourceOk() (*RunSearchForEventsRequestResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *RunSearchForEventsRequest) SetResource(v RunSearchForEventsRequestResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *RunSearchForEventsRequest) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


