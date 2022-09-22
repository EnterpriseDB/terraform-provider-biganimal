# PatchCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterName** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**InstanceType** | Pointer to [**CreateClusterInstanceType**](CreateClusterInstanceType.md) |  | [optional] 
**ReadOnlyConnections** | Pointer to **bool** |  | [optional] 
**Storage** | Pointer to [**PatchClusterStorage**](PatchClusterStorage.md) |  | [optional] 
**PrivateNetworking** | Pointer to **bool** |  | [optional] 
**AllowedIpRanges** | Pointer to [**[]AllowedIpRange**](AllowedIpRange.md) |  | [optional] 
**PgConfig** | Pointer to [**[]ClusterDetailPgConfigInner**](ClusterDetailPgConfigInner.md) |  | [optional] 
**Replicas** | Pointer to **float32** |  | [optional] 
**ClusterArchitecture** | Pointer to [**CreateClusterClusterArchitecture**](CreateClusterClusterArchitecture.md) |  | [optional] 
**BackupRetentionPeriod** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchCluster

`func NewPatchCluster() *PatchCluster`

NewPatchCluster instantiates a new PatchCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchClusterWithDefaults

`func NewPatchClusterWithDefaults() *PatchCluster`

NewPatchClusterWithDefaults instantiates a new PatchCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterName

`func (o *PatchCluster) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *PatchCluster) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *PatchCluster) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *PatchCluster) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetPassword

`func (o *PatchCluster) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PatchCluster) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PatchCluster) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PatchCluster) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetInstanceType

`func (o *PatchCluster) GetInstanceType() CreateClusterInstanceType`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *PatchCluster) GetInstanceTypeOk() (*CreateClusterInstanceType, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *PatchCluster) SetInstanceType(v CreateClusterInstanceType)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *PatchCluster) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetReadOnlyConnections

`func (o *PatchCluster) GetReadOnlyConnections() bool`

GetReadOnlyConnections returns the ReadOnlyConnections field if non-nil, zero value otherwise.

### GetReadOnlyConnectionsOk

`func (o *PatchCluster) GetReadOnlyConnectionsOk() (*bool, bool)`

GetReadOnlyConnectionsOk returns a tuple with the ReadOnlyConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnlyConnections

`func (o *PatchCluster) SetReadOnlyConnections(v bool)`

SetReadOnlyConnections sets ReadOnlyConnections field to given value.

### HasReadOnlyConnections

`func (o *PatchCluster) HasReadOnlyConnections() bool`

HasReadOnlyConnections returns a boolean if a field has been set.

### GetStorage

`func (o *PatchCluster) GetStorage() PatchClusterStorage`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *PatchCluster) GetStorageOk() (*PatchClusterStorage, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *PatchCluster) SetStorage(v PatchClusterStorage)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *PatchCluster) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetPrivateNetworking

`func (o *PatchCluster) GetPrivateNetworking() bool`

GetPrivateNetworking returns the PrivateNetworking field if non-nil, zero value otherwise.

### GetPrivateNetworkingOk

`func (o *PatchCluster) GetPrivateNetworkingOk() (*bool, bool)`

GetPrivateNetworkingOk returns a tuple with the PrivateNetworking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetworking

`func (o *PatchCluster) SetPrivateNetworking(v bool)`

SetPrivateNetworking sets PrivateNetworking field to given value.

### HasPrivateNetworking

`func (o *PatchCluster) HasPrivateNetworking() bool`

HasPrivateNetworking returns a boolean if a field has been set.

### GetAllowedIpRanges

`func (o *PatchCluster) GetAllowedIpRanges() []AllowedIpRange`

GetAllowedIpRanges returns the AllowedIpRanges field if non-nil, zero value otherwise.

### GetAllowedIpRangesOk

`func (o *PatchCluster) GetAllowedIpRangesOk() (*[]AllowedIpRange, bool)`

GetAllowedIpRangesOk returns a tuple with the AllowedIpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedIpRanges

`func (o *PatchCluster) SetAllowedIpRanges(v []AllowedIpRange)`

SetAllowedIpRanges sets AllowedIpRanges field to given value.

### HasAllowedIpRanges

`func (o *PatchCluster) HasAllowedIpRanges() bool`

HasAllowedIpRanges returns a boolean if a field has been set.

### GetPgConfig

`func (o *PatchCluster) GetPgConfig() []ClusterDetailPgConfigInner`

GetPgConfig returns the PgConfig field if non-nil, zero value otherwise.

### GetPgConfigOk

`func (o *PatchCluster) GetPgConfigOk() (*[]ClusterDetailPgConfigInner, bool)`

GetPgConfigOk returns a tuple with the PgConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgConfig

`func (o *PatchCluster) SetPgConfig(v []ClusterDetailPgConfigInner)`

SetPgConfig sets PgConfig field to given value.

### HasPgConfig

`func (o *PatchCluster) HasPgConfig() bool`

HasPgConfig returns a boolean if a field has been set.

### GetReplicas

`func (o *PatchCluster) GetReplicas() float32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *PatchCluster) GetReplicasOk() (*float32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *PatchCluster) SetReplicas(v float32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *PatchCluster) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetClusterArchitecture

`func (o *PatchCluster) GetClusterArchitecture() CreateClusterClusterArchitecture`

GetClusterArchitecture returns the ClusterArchitecture field if non-nil, zero value otherwise.

### GetClusterArchitectureOk

`func (o *PatchCluster) GetClusterArchitectureOk() (*CreateClusterClusterArchitecture, bool)`

GetClusterArchitectureOk returns a tuple with the ClusterArchitecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterArchitecture

`func (o *PatchCluster) SetClusterArchitecture(v CreateClusterClusterArchitecture)`

SetClusterArchitecture sets ClusterArchitecture field to given value.

### HasClusterArchitecture

`func (o *PatchCluster) HasClusterArchitecture() bool`

HasClusterArchitecture returns a boolean if a field has been set.

### GetBackupRetentionPeriod

`func (o *PatchCluster) GetBackupRetentionPeriod() string`

GetBackupRetentionPeriod returns the BackupRetentionPeriod field if non-nil, zero value otherwise.

### GetBackupRetentionPeriodOk

`func (o *PatchCluster) GetBackupRetentionPeriodOk() (*string, bool)`

GetBackupRetentionPeriodOk returns a tuple with the BackupRetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRetentionPeriod

`func (o *PatchCluster) SetBackupRetentionPeriod(v string)`

SetBackupRetentionPeriod sets BackupRetentionPeriod field to given value.

### HasBackupRetentionPeriod

`func (o *PatchCluster) HasBackupRetentionPeriod() bool`

HasBackupRetentionPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


