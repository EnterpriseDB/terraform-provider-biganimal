# CreateClusterClusterArchitecture

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterArchitectureId** | **string** |  | 
**Nodes** | **float32** |  | 
**Params** | Pointer to [**[]ClusterDetailPgConfigInner**](ClusterDetailPgConfigInner.md) |  | [optional] 

## Methods

### NewCreateClusterClusterArchitecture

`func NewCreateClusterClusterArchitecture(clusterArchitectureId string, nodes float32, ) *CreateClusterClusterArchitecture`

NewCreateClusterClusterArchitecture instantiates a new CreateClusterClusterArchitecture object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClusterClusterArchitectureWithDefaults

`func NewCreateClusterClusterArchitectureWithDefaults() *CreateClusterClusterArchitecture`

NewCreateClusterClusterArchitectureWithDefaults instantiates a new CreateClusterClusterArchitecture object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterArchitectureId

`func (o *CreateClusterClusterArchitecture) GetClusterArchitectureId() string`

GetClusterArchitectureId returns the ClusterArchitectureId field if non-nil, zero value otherwise.

### GetClusterArchitectureIdOk

`func (o *CreateClusterClusterArchitecture) GetClusterArchitectureIdOk() (*string, bool)`

GetClusterArchitectureIdOk returns a tuple with the ClusterArchitectureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterArchitectureId

`func (o *CreateClusterClusterArchitecture) SetClusterArchitectureId(v string)`

SetClusterArchitectureId sets ClusterArchitectureId field to given value.


### GetNodes

`func (o *CreateClusterClusterArchitecture) GetNodes() float32`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *CreateClusterClusterArchitecture) GetNodesOk() (*float32, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *CreateClusterClusterArchitecture) SetNodes(v float32)`

SetNodes sets Nodes field to given value.


### GetParams

`func (o *CreateClusterClusterArchitecture) GetParams() []ClusterDetailPgConfigInner`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *CreateClusterClusterArchitecture) GetParamsOk() (*[]ClusterDetailPgConfigInner, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *CreateClusterClusterArchitecture) SetParams(v []ClusterDetailPgConfigInner)`

SetParams sets Params field to given value.

### HasParams

`func (o *CreateClusterClusterArchitecture) HasParams() bool`

HasParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


