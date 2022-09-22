# ClusterPhase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | 
**Name** | **string** |  | 
**Category** | **string** |  | 
**EffectiveAt** | Pointer to [**PointInTime**](PointInTime.md) |  | [optional] 

## Methods

### NewClusterPhase

`func NewClusterPhase(id float32, name string, category string, ) *ClusterPhase`

NewClusterPhase instantiates a new ClusterPhase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterPhaseWithDefaults

`func NewClusterPhaseWithDefaults() *ClusterPhase`

NewClusterPhaseWithDefaults instantiates a new ClusterPhase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterPhase) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterPhase) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterPhase) SetId(v float32)`

SetId sets Id field to given value.


### GetName

`func (o *ClusterPhase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterPhase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterPhase) SetName(v string)`

SetName sets Name field to given value.


### GetCategory

`func (o *ClusterPhase) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ClusterPhase) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ClusterPhase) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetEffectiveAt

`func (o *ClusterPhase) GetEffectiveAt() PointInTime`

GetEffectiveAt returns the EffectiveAt field if non-nil, zero value otherwise.

### GetEffectiveAtOk

`func (o *ClusterPhase) GetEffectiveAtOk() (*PointInTime, bool)`

GetEffectiveAtOk returns a tuple with the EffectiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveAt

`func (o *ClusterPhase) SetEffectiveAt(v PointInTime)`

SetEffectiveAt sets EffectiveAt field to given value.

### HasEffectiveAt

`func (o *ClusterPhase) HasEffectiveAt() bool`

HasEffectiveAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


