# CreateCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterName** | **string** |  | 
**Password** | **string** |  | 
**PgType** | [**CreateClusterPgType**](CreateClusterPgType.md) |  | 
**PgVersion** | [**CreateClusterPgVersion**](CreateClusterPgVersion.md) |  | 
**Provider** | [**CreateClusterProvider**](CreateClusterProvider.md) |  | 
**ReadOnlyConnections** | Pointer to **bool** |  | [optional] 
**Region** | [**ActivateRegion202ResponseData**](ActivateRegion202ResponseData.md) |  | 
**InstanceType** | [**CreateClusterInstanceType**](CreateClusterInstanceType.md) |  | 
**Storage** | [**CreateClusterStorage**](CreateClusterStorage.md) |  | 
**PrivateNetworking** | **bool** |  | 
**AllowedIpRanges** | [**[]AllowedIpRange**](AllowedIpRange.md) |  | 
**PgConfig** | [**[]ClusterDetailPgConfigInner**](ClusterDetailPgConfigInner.md) |  | 
**Replicas** | Pointer to **float32** |  | [optional] 
**ClusterArchitecture** | Pointer to [**CreateClusterClusterArchitecture**](CreateClusterClusterArchitecture.md) |  | [optional] 
**BackupRetentionPeriod** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateCluster

`func NewCreateCluster(clusterName string, password string, pgType CreateClusterPgType, pgVersion CreateClusterPgVersion, provider CreateClusterProvider, region ActivateRegion202ResponseData, instanceType CreateClusterInstanceType, storage CreateClusterStorage, privateNetworking bool, allowedIpRanges []AllowedIpRange, pgConfig []ClusterDetailPgConfigInner, ) *CreateCluster`

NewCreateCluster instantiates a new CreateCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClusterWithDefaults

`func NewCreateClusterWithDefaults() *CreateCluster`

NewCreateClusterWithDefaults instantiates a new CreateCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterName

`func (o *CreateCluster) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *CreateCluster) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *CreateCluster) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetPassword

`func (o *CreateCluster) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateCluster) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateCluster) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetPgType

`func (o *CreateCluster) GetPgType() CreateClusterPgType`

GetPgType returns the PgType field if non-nil, zero value otherwise.

### GetPgTypeOk

`func (o *CreateCluster) GetPgTypeOk() (*CreateClusterPgType, bool)`

GetPgTypeOk returns a tuple with the PgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgType

`func (o *CreateCluster) SetPgType(v CreateClusterPgType)`

SetPgType sets PgType field to given value.


### GetPgVersion

`func (o *CreateCluster) GetPgVersion() CreateClusterPgVersion`

GetPgVersion returns the PgVersion field if non-nil, zero value otherwise.

### GetPgVersionOk

`func (o *CreateCluster) GetPgVersionOk() (*CreateClusterPgVersion, bool)`

GetPgVersionOk returns a tuple with the PgVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgVersion

`func (o *CreateCluster) SetPgVersion(v CreateClusterPgVersion)`

SetPgVersion sets PgVersion field to given value.


### GetProvider

`func (o *CreateCluster) GetProvider() CreateClusterProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CreateCluster) GetProviderOk() (*CreateClusterProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CreateCluster) SetProvider(v CreateClusterProvider)`

SetProvider sets Provider field to given value.


### GetReadOnlyConnections

`func (o *CreateCluster) GetReadOnlyConnections() bool`

GetReadOnlyConnections returns the ReadOnlyConnections field if non-nil, zero value otherwise.

### GetReadOnlyConnectionsOk

`func (o *CreateCluster) GetReadOnlyConnectionsOk() (*bool, bool)`

GetReadOnlyConnectionsOk returns a tuple with the ReadOnlyConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnlyConnections

`func (o *CreateCluster) SetReadOnlyConnections(v bool)`

SetReadOnlyConnections sets ReadOnlyConnections field to given value.

### HasReadOnlyConnections

`func (o *CreateCluster) HasReadOnlyConnections() bool`

HasReadOnlyConnections returns a boolean if a field has been set.

### GetRegion

`func (o *CreateCluster) GetRegion() ActivateRegion202ResponseData`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateCluster) GetRegionOk() (*ActivateRegion202ResponseData, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateCluster) SetRegion(v ActivateRegion202ResponseData)`

SetRegion sets Region field to given value.


### GetInstanceType

`func (o *CreateCluster) GetInstanceType() CreateClusterInstanceType`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *CreateCluster) GetInstanceTypeOk() (*CreateClusterInstanceType, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *CreateCluster) SetInstanceType(v CreateClusterInstanceType)`

SetInstanceType sets InstanceType field to given value.


### GetStorage

`func (o *CreateCluster) GetStorage() CreateClusterStorage`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *CreateCluster) GetStorageOk() (*CreateClusterStorage, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *CreateCluster) SetStorage(v CreateClusterStorage)`

SetStorage sets Storage field to given value.


### GetPrivateNetworking

`func (o *CreateCluster) GetPrivateNetworking() bool`

GetPrivateNetworking returns the PrivateNetworking field if non-nil, zero value otherwise.

### GetPrivateNetworkingOk

`func (o *CreateCluster) GetPrivateNetworkingOk() (*bool, bool)`

GetPrivateNetworkingOk returns a tuple with the PrivateNetworking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetworking

`func (o *CreateCluster) SetPrivateNetworking(v bool)`

SetPrivateNetworking sets PrivateNetworking field to given value.


### GetAllowedIpRanges

`func (o *CreateCluster) GetAllowedIpRanges() []AllowedIpRange`

GetAllowedIpRanges returns the AllowedIpRanges field if non-nil, zero value otherwise.

### GetAllowedIpRangesOk

`func (o *CreateCluster) GetAllowedIpRangesOk() (*[]AllowedIpRange, bool)`

GetAllowedIpRangesOk returns a tuple with the AllowedIpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedIpRanges

`func (o *CreateCluster) SetAllowedIpRanges(v []AllowedIpRange)`

SetAllowedIpRanges sets AllowedIpRanges field to given value.


### GetPgConfig

`func (o *CreateCluster) GetPgConfig() []ClusterDetailPgConfigInner`

GetPgConfig returns the PgConfig field if non-nil, zero value otherwise.

### GetPgConfigOk

`func (o *CreateCluster) GetPgConfigOk() (*[]ClusterDetailPgConfigInner, bool)`

GetPgConfigOk returns a tuple with the PgConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgConfig

`func (o *CreateCluster) SetPgConfig(v []ClusterDetailPgConfigInner)`

SetPgConfig sets PgConfig field to given value.


### GetReplicas

`func (o *CreateCluster) GetReplicas() float32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *CreateCluster) GetReplicasOk() (*float32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *CreateCluster) SetReplicas(v float32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *CreateCluster) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetClusterArchitecture

`func (o *CreateCluster) GetClusterArchitecture() CreateClusterClusterArchitecture`

GetClusterArchitecture returns the ClusterArchitecture field if non-nil, zero value otherwise.

### GetClusterArchitectureOk

`func (o *CreateCluster) GetClusterArchitectureOk() (*CreateClusterClusterArchitecture, bool)`

GetClusterArchitectureOk returns a tuple with the ClusterArchitecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterArchitecture

`func (o *CreateCluster) SetClusterArchitecture(v CreateClusterClusterArchitecture)`

SetClusterArchitecture sets ClusterArchitecture field to given value.

### HasClusterArchitecture

`func (o *CreateCluster) HasClusterArchitecture() bool`

HasClusterArchitecture returns a boolean if a field has been set.

### GetBackupRetentionPeriod

`func (o *CreateCluster) GetBackupRetentionPeriod() string`

GetBackupRetentionPeriod returns the BackupRetentionPeriod field if non-nil, zero value otherwise.

### GetBackupRetentionPeriodOk

`func (o *CreateCluster) GetBackupRetentionPeriodOk() (*string, bool)`

GetBackupRetentionPeriodOk returns a tuple with the BackupRetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRetentionPeriod

`func (o *CreateCluster) SetBackupRetentionPeriod(v string)`

SetBackupRetentionPeriod sets BackupRetentionPeriod field to given value.

### HasBackupRetentionPeriod

`func (o *CreateCluster) HasBackupRetentionPeriod() bool`

HasBackupRetentionPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


