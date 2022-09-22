# RestoreCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupRetentionPeriod** | Pointer to **string** |  | [optional] 
**SelectedRestorePointInTime** | Pointer to **string** |  | [optional] 
**ClusterName** | **string** |  | 
**Password** | **string** |  | 
**Region** | [**ActivateRegion202ResponseData**](ActivateRegion202ResponseData.md) |  | 
**InstanceType** | [**CreateClusterInstanceType**](CreateClusterInstanceType.md) |  | 
**Storage** | [**CreateClusterStorage**](CreateClusterStorage.md) |  | 
**AllowedIpRanges** | [**[]AllowedIpRange**](AllowedIpRange.md) |  | 
**PgConfig** | [**[]ClusterDetailPgConfigInner**](ClusterDetailPgConfigInner.md) |  | 
**Replicas** | Pointer to **float32** |  | [optional] 
**ReadOnlyConnections** | Pointer to **bool** |  | [optional] 
**ClusterArchitecture** | Pointer to [**CreateClusterClusterArchitecture**](CreateClusterClusterArchitecture.md) |  | [optional] 

## Methods

### NewRestoreCluster

`func NewRestoreCluster(clusterName string, password string, region ActivateRegion202ResponseData, instanceType CreateClusterInstanceType, storage CreateClusterStorage, allowedIpRanges []AllowedIpRange, pgConfig []ClusterDetailPgConfigInner, ) *RestoreCluster`

NewRestoreCluster instantiates a new RestoreCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreClusterWithDefaults

`func NewRestoreClusterWithDefaults() *RestoreCluster`

NewRestoreClusterWithDefaults instantiates a new RestoreCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupRetentionPeriod

`func (o *RestoreCluster) GetBackupRetentionPeriod() string`

GetBackupRetentionPeriod returns the BackupRetentionPeriod field if non-nil, zero value otherwise.

### GetBackupRetentionPeriodOk

`func (o *RestoreCluster) GetBackupRetentionPeriodOk() (*string, bool)`

GetBackupRetentionPeriodOk returns a tuple with the BackupRetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRetentionPeriod

`func (o *RestoreCluster) SetBackupRetentionPeriod(v string)`

SetBackupRetentionPeriod sets BackupRetentionPeriod field to given value.

### HasBackupRetentionPeriod

`func (o *RestoreCluster) HasBackupRetentionPeriod() bool`

HasBackupRetentionPeriod returns a boolean if a field has been set.

### GetSelectedRestorePointInTime

`func (o *RestoreCluster) GetSelectedRestorePointInTime() string`

GetSelectedRestorePointInTime returns the SelectedRestorePointInTime field if non-nil, zero value otherwise.

### GetSelectedRestorePointInTimeOk

`func (o *RestoreCluster) GetSelectedRestorePointInTimeOk() (*string, bool)`

GetSelectedRestorePointInTimeOk returns a tuple with the SelectedRestorePointInTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedRestorePointInTime

`func (o *RestoreCluster) SetSelectedRestorePointInTime(v string)`

SetSelectedRestorePointInTime sets SelectedRestorePointInTime field to given value.

### HasSelectedRestorePointInTime

`func (o *RestoreCluster) HasSelectedRestorePointInTime() bool`

HasSelectedRestorePointInTime returns a boolean if a field has been set.

### GetClusterName

`func (o *RestoreCluster) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *RestoreCluster) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *RestoreCluster) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetPassword

`func (o *RestoreCluster) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RestoreCluster) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RestoreCluster) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetRegion

`func (o *RestoreCluster) GetRegion() ActivateRegion202ResponseData`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *RestoreCluster) GetRegionOk() (*ActivateRegion202ResponseData, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *RestoreCluster) SetRegion(v ActivateRegion202ResponseData)`

SetRegion sets Region field to given value.


### GetInstanceType

`func (o *RestoreCluster) GetInstanceType() CreateClusterInstanceType`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *RestoreCluster) GetInstanceTypeOk() (*CreateClusterInstanceType, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *RestoreCluster) SetInstanceType(v CreateClusterInstanceType)`

SetInstanceType sets InstanceType field to given value.


### GetStorage

`func (o *RestoreCluster) GetStorage() CreateClusterStorage`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *RestoreCluster) GetStorageOk() (*CreateClusterStorage, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *RestoreCluster) SetStorage(v CreateClusterStorage)`

SetStorage sets Storage field to given value.


### GetAllowedIpRanges

`func (o *RestoreCluster) GetAllowedIpRanges() []AllowedIpRange`

GetAllowedIpRanges returns the AllowedIpRanges field if non-nil, zero value otherwise.

### GetAllowedIpRangesOk

`func (o *RestoreCluster) GetAllowedIpRangesOk() (*[]AllowedIpRange, bool)`

GetAllowedIpRangesOk returns a tuple with the AllowedIpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedIpRanges

`func (o *RestoreCluster) SetAllowedIpRanges(v []AllowedIpRange)`

SetAllowedIpRanges sets AllowedIpRanges field to given value.


### GetPgConfig

`func (o *RestoreCluster) GetPgConfig() []ClusterDetailPgConfigInner`

GetPgConfig returns the PgConfig field if non-nil, zero value otherwise.

### GetPgConfigOk

`func (o *RestoreCluster) GetPgConfigOk() (*[]ClusterDetailPgConfigInner, bool)`

GetPgConfigOk returns a tuple with the PgConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgConfig

`func (o *RestoreCluster) SetPgConfig(v []ClusterDetailPgConfigInner)`

SetPgConfig sets PgConfig field to given value.


### GetReplicas

`func (o *RestoreCluster) GetReplicas() float32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *RestoreCluster) GetReplicasOk() (*float32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *RestoreCluster) SetReplicas(v float32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *RestoreCluster) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetReadOnlyConnections

`func (o *RestoreCluster) GetReadOnlyConnections() bool`

GetReadOnlyConnections returns the ReadOnlyConnections field if non-nil, zero value otherwise.

### GetReadOnlyConnectionsOk

`func (o *RestoreCluster) GetReadOnlyConnectionsOk() (*bool, bool)`

GetReadOnlyConnectionsOk returns a tuple with the ReadOnlyConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnlyConnections

`func (o *RestoreCluster) SetReadOnlyConnections(v bool)`

SetReadOnlyConnections sets ReadOnlyConnections field to given value.

### HasReadOnlyConnections

`func (o *RestoreCluster) HasReadOnlyConnections() bool`

HasReadOnlyConnections returns a boolean if a field has been set.

### GetClusterArchitecture

`func (o *RestoreCluster) GetClusterArchitecture() CreateClusterClusterArchitecture`

GetClusterArchitecture returns the ClusterArchitecture field if non-nil, zero value otherwise.

### GetClusterArchitectureOk

`func (o *RestoreCluster) GetClusterArchitectureOk() (*CreateClusterClusterArchitecture, bool)`

GetClusterArchitectureOk returns a tuple with the ClusterArchitecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterArchitecture

`func (o *RestoreCluster) SetClusterArchitecture(v CreateClusterClusterArchitecture)`

SetClusterArchitecture sets ClusterArchitecture field to given value.

### HasClusterArchitecture

`func (o *RestoreCluster) HasClusterArchitecture() bool`

HasClusterArchitecture returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


