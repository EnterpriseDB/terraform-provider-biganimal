# StatusMaintenanceWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusMaintenanceWindowId** | **string** |  | 
**StartsAt** | [**PointInTime**](PointInTime.md) |  | 
**EndsAt** | [**PointInTime**](PointInTime.md) |  | 

## Methods

### NewStatusMaintenanceWindow

`func NewStatusMaintenanceWindow(statusMaintenanceWindowId string, startsAt PointInTime, endsAt PointInTime, ) *StatusMaintenanceWindow`

NewStatusMaintenanceWindow instantiates a new StatusMaintenanceWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusMaintenanceWindowWithDefaults

`func NewStatusMaintenanceWindowWithDefaults() *StatusMaintenanceWindow`

NewStatusMaintenanceWindowWithDefaults instantiates a new StatusMaintenanceWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusMaintenanceWindowId

`func (o *StatusMaintenanceWindow) GetStatusMaintenanceWindowId() string`

GetStatusMaintenanceWindowId returns the StatusMaintenanceWindowId field if non-nil, zero value otherwise.

### GetStatusMaintenanceWindowIdOk

`func (o *StatusMaintenanceWindow) GetStatusMaintenanceWindowIdOk() (*string, bool)`

GetStatusMaintenanceWindowIdOk returns a tuple with the StatusMaintenanceWindowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMaintenanceWindowId

`func (o *StatusMaintenanceWindow) SetStatusMaintenanceWindowId(v string)`

SetStatusMaintenanceWindowId sets StatusMaintenanceWindowId field to given value.


### GetStartsAt

`func (o *StatusMaintenanceWindow) GetStartsAt() PointInTime`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *StatusMaintenanceWindow) GetStartsAtOk() (*PointInTime, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *StatusMaintenanceWindow) SetStartsAt(v PointInTime)`

SetStartsAt sets StartsAt field to given value.


### GetEndsAt

`func (o *StatusMaintenanceWindow) GetEndsAt() PointInTime`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *StatusMaintenanceWindow) GetEndsAtOk() (*PointInTime, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *StatusMaintenanceWindow) SetEndsAt(v PointInTime)`

SetEndsAt sets EndsAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


