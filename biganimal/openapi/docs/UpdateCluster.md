# UpdateCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterName** | **string** |  | 
**Password** | Pointer to **string** |  | [optional] 
**InstanceType** | [**CreateClusterInstanceType**](CreateClusterInstanceType.md) |  | 
**ReadOnlyConnections** | Pointer to **bool** |  | [optional] 
**Storage** | [**CreateClusterStorage**](CreateClusterStorage.md) |  | 
**PrivateNetworking** | **bool** |  | 
**AllowedIpRanges** | [**[]AllowedIpRange**](AllowedIpRange.md) |  | 
**PgConfig** | [**[]ClusterDetailPgConfigInner**](ClusterDetailPgConfigInner.md) |  | 
**Replicas** | Pointer to **float32** |  | [optional] 
**ClusterArchitecture** | Pointer to [**CreateClusterClusterArchitecture**](CreateClusterClusterArchitecture.md) |  | [optional] 
**BackupRetentionPeriod** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateCluster

`func NewUpdateCluster(clusterName string, instanceType CreateClusterInstanceType, storage CreateClusterStorage, privateNetworking bool, allowedIpRanges []AllowedIpRange, pgConfig []ClusterDetailPgConfigInner, ) *UpdateCluster`

NewUpdateCluster instantiates a new UpdateCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateClusterWithDefaults

`func NewUpdateClusterWithDefaults() *UpdateCluster`

NewUpdateClusterWithDefaults instantiates a new UpdateCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterName

`func (o *UpdateCluster) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *UpdateCluster) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *UpdateCluster) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetPassword

`func (o *UpdateCluster) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateCluster) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateCluster) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateCluster) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetInstanceType

`func (o *UpdateCluster) GetInstanceType() CreateClusterInstanceType`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *UpdateCluster) GetInstanceTypeOk() (*CreateClusterInstanceType, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *UpdateCluster) SetInstanceType(v CreateClusterInstanceType)`

SetInstanceType sets InstanceType field to given value.


### GetReadOnlyConnections

`func (o *UpdateCluster) GetReadOnlyConnections() bool`

GetReadOnlyConnections returns the ReadOnlyConnections field if non-nil, zero value otherwise.

### GetReadOnlyConnectionsOk

`func (o *UpdateCluster) GetReadOnlyConnectionsOk() (*bool, bool)`

GetReadOnlyConnectionsOk returns a tuple with the ReadOnlyConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnlyConnections

`func (o *UpdateCluster) SetReadOnlyConnections(v bool)`

SetReadOnlyConnections sets ReadOnlyConnections field to given value.

### HasReadOnlyConnections

`func (o *UpdateCluster) HasReadOnlyConnections() bool`

HasReadOnlyConnections returns a boolean if a field has been set.

### GetStorage

`func (o *UpdateCluster) GetStorage() CreateClusterStorage`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *UpdateCluster) GetStorageOk() (*CreateClusterStorage, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *UpdateCluster) SetStorage(v CreateClusterStorage)`

SetStorage sets Storage field to given value.


### GetPrivateNetworking

`func (o *UpdateCluster) GetPrivateNetworking() bool`

GetPrivateNetworking returns the PrivateNetworking field if non-nil, zero value otherwise.

### GetPrivateNetworkingOk

`func (o *UpdateCluster) GetPrivateNetworkingOk() (*bool, bool)`

GetPrivateNetworkingOk returns a tuple with the PrivateNetworking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetworking

`func (o *UpdateCluster) SetPrivateNetworking(v bool)`

SetPrivateNetworking sets PrivateNetworking field to given value.


### GetAllowedIpRanges

`func (o *UpdateCluster) GetAllowedIpRanges() []AllowedIpRange`

GetAllowedIpRanges returns the AllowedIpRanges field if non-nil, zero value otherwise.

### GetAllowedIpRangesOk

`func (o *UpdateCluster) GetAllowedIpRangesOk() (*[]AllowedIpRange, bool)`

GetAllowedIpRangesOk returns a tuple with the AllowedIpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedIpRanges

`func (o *UpdateCluster) SetAllowedIpRanges(v []AllowedIpRange)`

SetAllowedIpRanges sets AllowedIpRanges field to given value.


### GetPgConfig

`func (o *UpdateCluster) GetPgConfig() []ClusterDetailPgConfigInner`

GetPgConfig returns the PgConfig field if non-nil, zero value otherwise.

### GetPgConfigOk

`func (o *UpdateCluster) GetPgConfigOk() (*[]ClusterDetailPgConfigInner, bool)`

GetPgConfigOk returns a tuple with the PgConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgConfig

`func (o *UpdateCluster) SetPgConfig(v []ClusterDetailPgConfigInner)`

SetPgConfig sets PgConfig field to given value.


### GetReplicas

`func (o *UpdateCluster) GetReplicas() float32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *UpdateCluster) GetReplicasOk() (*float32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *UpdateCluster) SetReplicas(v float32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *UpdateCluster) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetClusterArchitecture

`func (o *UpdateCluster) GetClusterArchitecture() CreateClusterClusterArchitecture`

GetClusterArchitecture returns the ClusterArchitecture field if non-nil, zero value otherwise.

### GetClusterArchitectureOk

`func (o *UpdateCluster) GetClusterArchitectureOk() (*CreateClusterClusterArchitecture, bool)`

GetClusterArchitectureOk returns a tuple with the ClusterArchitecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterArchitecture

`func (o *UpdateCluster) SetClusterArchitecture(v CreateClusterClusterArchitecture)`

SetClusterArchitecture sets ClusterArchitecture field to given value.

### HasClusterArchitecture

`func (o *UpdateCluster) HasClusterArchitecture() bool`

HasClusterArchitecture returns a boolean if a field has been set.

### GetBackupRetentionPeriod

`func (o *UpdateCluster) GetBackupRetentionPeriod() string`

GetBackupRetentionPeriod returns the BackupRetentionPeriod field if non-nil, zero value otherwise.

### GetBackupRetentionPeriodOk

`func (o *UpdateCluster) GetBackupRetentionPeriodOk() (*string, bool)`

GetBackupRetentionPeriodOk returns a tuple with the BackupRetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRetentionPeriod

`func (o *UpdateCluster) SetBackupRetentionPeriod(v string)`

SetBackupRetentionPeriod sets BackupRetentionPeriod field to given value.

### HasBackupRetentionPeriod

`func (o *UpdateCluster) HasBackupRetentionPeriod() bool`

HasBackupRetentionPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


