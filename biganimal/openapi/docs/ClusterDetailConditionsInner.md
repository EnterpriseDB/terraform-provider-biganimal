# ClusterDetailConditionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionStatus** | Pointer to **string** |  | [optional] 
**LastTransitionTime** | Pointer to [**PointInTime**](PointInTime.md) |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewClusterDetailConditionsInner

`func NewClusterDetailConditionsInner() *ClusterDetailConditionsInner`

NewClusterDetailConditionsInner instantiates a new ClusterDetailConditionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterDetailConditionsInnerWithDefaults

`func NewClusterDetailConditionsInnerWithDefaults() *ClusterDetailConditionsInner`

NewClusterDetailConditionsInnerWithDefaults instantiates a new ClusterDetailConditionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionStatus

`func (o *ClusterDetailConditionsInner) GetConditionStatus() string`

GetConditionStatus returns the ConditionStatus field if non-nil, zero value otherwise.

### GetConditionStatusOk

`func (o *ClusterDetailConditionsInner) GetConditionStatusOk() (*string, bool)`

GetConditionStatusOk returns a tuple with the ConditionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionStatus

`func (o *ClusterDetailConditionsInner) SetConditionStatus(v string)`

SetConditionStatus sets ConditionStatus field to given value.

### HasConditionStatus

`func (o *ClusterDetailConditionsInner) HasConditionStatus() bool`

HasConditionStatus returns a boolean if a field has been set.

### GetLastTransitionTime

`func (o *ClusterDetailConditionsInner) GetLastTransitionTime() PointInTime`

GetLastTransitionTime returns the LastTransitionTime field if non-nil, zero value otherwise.

### GetLastTransitionTimeOk

`func (o *ClusterDetailConditionsInner) GetLastTransitionTimeOk() (*PointInTime, bool)`

GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionTime

`func (o *ClusterDetailConditionsInner) SetLastTransitionTime(v PointInTime)`

SetLastTransitionTime sets LastTransitionTime field to given value.

### HasLastTransitionTime

`func (o *ClusterDetailConditionsInner) HasLastTransitionTime() bool`

HasLastTransitionTime returns a boolean if a field has been set.

### GetReason

`func (o *ClusterDetailConditionsInner) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ClusterDetailConditionsInner) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ClusterDetailConditionsInner) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ClusterDetailConditionsInner) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetMessage

`func (o *ClusterDetailConditionsInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ClusterDetailConditionsInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ClusterDetailConditionsInner) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ClusterDetailConditionsInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetType

`func (o *ClusterDetailConditionsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClusterDetailConditionsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClusterDetailConditionsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ClusterDetailConditionsInner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


