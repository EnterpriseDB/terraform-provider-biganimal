# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**EventAction**](EventAction.md) |  | [optional] 
**User** | Pointer to [**RunSearchForEventsRequestUser**](RunSearchForEventsRequestUser.md) |  | [optional] 
**Resource** | Pointer to [**EventResource**](EventResource.md) |  | [optional] 
**IpAddress** | **string** |  | 
**SessionId** | **string** |  | 
**CreatedAt** | Pointer to [**PointInTime**](PointInTime.md) |  | [optional] 

## Methods

### NewEvent

`func NewEvent(ipAddress string, sessionId string, ) *Event`

NewEvent instantiates a new Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventWithDefaults

`func NewEventWithDefaults() *Event`

NewEventWithDefaults instantiates a new Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *Event) GetAction() EventAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *Event) GetActionOk() (*EventAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *Event) SetAction(v EventAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *Event) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetUser

`func (o *Event) GetUser() RunSearchForEventsRequestUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Event) GetUserOk() (*RunSearchForEventsRequestUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Event) SetUser(v RunSearchForEventsRequestUser)`

SetUser sets User field to given value.

### HasUser

`func (o *Event) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetResource

`func (o *Event) GetResource() EventResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *Event) GetResourceOk() (*EventResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *Event) SetResource(v EventResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *Event) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetIpAddress

`func (o *Event) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *Event) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *Event) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### GetSessionId

`func (o *Event) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *Event) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *Event) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetCreatedAt

`func (o *Event) GetCreatedAt() PointInTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Event) GetCreatedAtOk() (*PointInTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Event) SetCreatedAt(v PointInTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Event) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


