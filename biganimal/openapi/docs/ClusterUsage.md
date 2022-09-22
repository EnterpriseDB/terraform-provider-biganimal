# ClusterUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** |  | 
**ClusterName** | **string** |  | 
**PgType** | [**ClusterUsagePgType**](ClusterUsagePgType.md) |  | 
**CloudProvider** | [**ClusterUsageCloudProvider**](ClusterUsageCloudProvider.md) |  | 
**VcpuHours** | **float32** |  | 
**ClusterArchitecture** | Pointer to [**ClusterUsageClusterArchitecture**](ClusterUsageClusterArchitecture.md) |  | [optional] 
**ClusterCreationTime** | Pointer to [**PointInTime**](PointInTime.md) |  | [optional] 

## Methods

### NewClusterUsage

`func NewClusterUsage(clusterId string, clusterName string, pgType ClusterUsagePgType, cloudProvider ClusterUsageCloudProvider, vcpuHours float32, ) *ClusterUsage`

NewClusterUsage instantiates a new ClusterUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterUsageWithDefaults

`func NewClusterUsageWithDefaults() *ClusterUsage`

NewClusterUsageWithDefaults instantiates a new ClusterUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *ClusterUsage) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ClusterUsage) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ClusterUsage) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetClusterName

`func (o *ClusterUsage) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *ClusterUsage) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *ClusterUsage) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetPgType

`func (o *ClusterUsage) GetPgType() ClusterUsagePgType`

GetPgType returns the PgType field if non-nil, zero value otherwise.

### GetPgTypeOk

`func (o *ClusterUsage) GetPgTypeOk() (*ClusterUsagePgType, bool)`

GetPgTypeOk returns a tuple with the PgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgType

`func (o *ClusterUsage) SetPgType(v ClusterUsagePgType)`

SetPgType sets PgType field to given value.


### GetCloudProvider

`func (o *ClusterUsage) GetCloudProvider() ClusterUsageCloudProvider`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ClusterUsage) GetCloudProviderOk() (*ClusterUsageCloudProvider, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ClusterUsage) SetCloudProvider(v ClusterUsageCloudProvider)`

SetCloudProvider sets CloudProvider field to given value.


### GetVcpuHours

`func (o *ClusterUsage) GetVcpuHours() float32`

GetVcpuHours returns the VcpuHours field if non-nil, zero value otherwise.

### GetVcpuHoursOk

`func (o *ClusterUsage) GetVcpuHoursOk() (*float32, bool)`

GetVcpuHoursOk returns a tuple with the VcpuHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpuHours

`func (o *ClusterUsage) SetVcpuHours(v float32)`

SetVcpuHours sets VcpuHours field to given value.


### GetClusterArchitecture

`func (o *ClusterUsage) GetClusterArchitecture() ClusterUsageClusterArchitecture`

GetClusterArchitecture returns the ClusterArchitecture field if non-nil, zero value otherwise.

### GetClusterArchitectureOk

`func (o *ClusterUsage) GetClusterArchitectureOk() (*ClusterUsageClusterArchitecture, bool)`

GetClusterArchitectureOk returns a tuple with the ClusterArchitecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterArchitecture

`func (o *ClusterUsage) SetClusterArchitecture(v ClusterUsageClusterArchitecture)`

SetClusterArchitecture sets ClusterArchitecture field to given value.

### HasClusterArchitecture

`func (o *ClusterUsage) HasClusterArchitecture() bool`

HasClusterArchitecture returns a boolean if a field has been set.

### GetClusterCreationTime

`func (o *ClusterUsage) GetClusterCreationTime() PointInTime`

GetClusterCreationTime returns the ClusterCreationTime field if non-nil, zero value otherwise.

### GetClusterCreationTimeOk

`func (o *ClusterUsage) GetClusterCreationTimeOk() (*PointInTime, bool)`

GetClusterCreationTimeOk returns a tuple with the ClusterCreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterCreationTime

`func (o *ClusterUsage) SetClusterCreationTime(v PointInTime)`

SetClusterCreationTime sets ClusterCreationTime field to given value.

### HasClusterCreationTime

`func (o *ClusterUsage) HasClusterCreationTime() bool`

HasClusterCreationTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


