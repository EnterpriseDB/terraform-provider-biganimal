# ClusterDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** |  | 
**ClusterName** | **string** |  | 
**ClusterArchitecture** | Pointer to [**ClusterDetailClusterArchitecture**](ClusterDetailClusterArchitecture.md) |  | [optional] 
**PrivateNetworking** | **bool** |  | 
**BackupRetentionPeriod** | **string** |  | 
**Phase** | **string** |  | 
**Replicas** | Pointer to **float32** |  | [optional] 
**CreatedAt** | Pointer to [**PointInTime**](PointInTime.md) |  | [optional] 
**DeletedAt** | Pointer to [**PointInTime**](PointInTime.md) |  | [optional] 
**ExpiredAt** | Pointer to [**PointInTime**](PointInTime.md) |  | [optional] 
**FirstRecoverabilityPointAt** | Pointer to [**PointInTime**](PointInTime.md) |  | [optional] 
**AllowedIpRanges** | [**[]AllowedIpRange**](AllowedIpRange.md) |  | 
**PgConfig** | Pointer to [**[]ClusterDetailPgConfigInner**](ClusterDetailPgConfigInner.md) |  | [optional] 
**PgType** | Pointer to [**PgType**](PgType.md) |  | [optional] 
**PgVersion** | Pointer to [**PgVersion**](PgVersion.md) |  | [optional] 
**Provider** | Pointer to [**CloudProvider**](CloudProvider.md) |  | [optional] 
**ReadOnlyConnections** | Pointer to **bool** |  | [optional] 
**Region** | Pointer to [**CloudProviderRegion**](CloudProviderRegion.md) |  | [optional] 
**ResizingPvc** | **[]string** |  | 
**Conditions** | [**[]ClusterDetailConditionsInner**](ClusterDetailConditionsInner.md) |  | 
**InstanceType** | Pointer to [**CloudProviderRegionInstanceType**](CloudProviderRegionInstanceType.md) |  | [optional] 
**Storage** | Pointer to [**ClusterDetailStorage**](ClusterDetailStorage.md) |  | [optional] 
**EvaluatedPgConfig** | Pointer to [**[]ClusterDetailPgConfigInner**](ClusterDetailPgConfigInner.md) |  | [optional] 

## Methods

### NewClusterDetail

`func NewClusterDetail(clusterId string, clusterName string, privateNetworking bool, backupRetentionPeriod string, phase string, allowedIpRanges []AllowedIpRange, resizingPvc []string, conditions []ClusterDetailConditionsInner, ) *ClusterDetail`

NewClusterDetail instantiates a new ClusterDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterDetailWithDefaults

`func NewClusterDetailWithDefaults() *ClusterDetail`

NewClusterDetailWithDefaults instantiates a new ClusterDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *ClusterDetail) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ClusterDetail) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ClusterDetail) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetClusterName

`func (o *ClusterDetail) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *ClusterDetail) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *ClusterDetail) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetClusterArchitecture

`func (o *ClusterDetail) GetClusterArchitecture() ClusterDetailClusterArchitecture`

GetClusterArchitecture returns the ClusterArchitecture field if non-nil, zero value otherwise.

### GetClusterArchitectureOk

`func (o *ClusterDetail) GetClusterArchitectureOk() (*ClusterDetailClusterArchitecture, bool)`

GetClusterArchitectureOk returns a tuple with the ClusterArchitecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterArchitecture

`func (o *ClusterDetail) SetClusterArchitecture(v ClusterDetailClusterArchitecture)`

SetClusterArchitecture sets ClusterArchitecture field to given value.

### HasClusterArchitecture

`func (o *ClusterDetail) HasClusterArchitecture() bool`

HasClusterArchitecture returns a boolean if a field has been set.

### GetPrivateNetworking

`func (o *ClusterDetail) GetPrivateNetworking() bool`

GetPrivateNetworking returns the PrivateNetworking field if non-nil, zero value otherwise.

### GetPrivateNetworkingOk

`func (o *ClusterDetail) GetPrivateNetworkingOk() (*bool, bool)`

GetPrivateNetworkingOk returns a tuple with the PrivateNetworking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetworking

`func (o *ClusterDetail) SetPrivateNetworking(v bool)`

SetPrivateNetworking sets PrivateNetworking field to given value.


### GetBackupRetentionPeriod

`func (o *ClusterDetail) GetBackupRetentionPeriod() string`

GetBackupRetentionPeriod returns the BackupRetentionPeriod field if non-nil, zero value otherwise.

### GetBackupRetentionPeriodOk

`func (o *ClusterDetail) GetBackupRetentionPeriodOk() (*string, bool)`

GetBackupRetentionPeriodOk returns a tuple with the BackupRetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRetentionPeriod

`func (o *ClusterDetail) SetBackupRetentionPeriod(v string)`

SetBackupRetentionPeriod sets BackupRetentionPeriod field to given value.


### GetPhase

`func (o *ClusterDetail) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *ClusterDetail) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *ClusterDetail) SetPhase(v string)`

SetPhase sets Phase field to given value.


### GetReplicas

`func (o *ClusterDetail) GetReplicas() float32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *ClusterDetail) GetReplicasOk() (*float32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *ClusterDetail) SetReplicas(v float32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *ClusterDetail) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ClusterDetail) GetCreatedAt() PointInTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ClusterDetail) GetCreatedAtOk() (*PointInTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ClusterDetail) SetCreatedAt(v PointInTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ClusterDetail) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ClusterDetail) GetDeletedAt() PointInTime`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ClusterDetail) GetDeletedAtOk() (*PointInTime, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ClusterDetail) SetDeletedAt(v PointInTime)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ClusterDetail) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetExpiredAt

`func (o *ClusterDetail) GetExpiredAt() PointInTime`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *ClusterDetail) GetExpiredAtOk() (*PointInTime, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *ClusterDetail) SetExpiredAt(v PointInTime)`

SetExpiredAt sets ExpiredAt field to given value.

### HasExpiredAt

`func (o *ClusterDetail) HasExpiredAt() bool`

HasExpiredAt returns a boolean if a field has been set.

### GetFirstRecoverabilityPointAt

`func (o *ClusterDetail) GetFirstRecoverabilityPointAt() PointInTime`

GetFirstRecoverabilityPointAt returns the FirstRecoverabilityPointAt field if non-nil, zero value otherwise.

### GetFirstRecoverabilityPointAtOk

`func (o *ClusterDetail) GetFirstRecoverabilityPointAtOk() (*PointInTime, bool)`

GetFirstRecoverabilityPointAtOk returns a tuple with the FirstRecoverabilityPointAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstRecoverabilityPointAt

`func (o *ClusterDetail) SetFirstRecoverabilityPointAt(v PointInTime)`

SetFirstRecoverabilityPointAt sets FirstRecoverabilityPointAt field to given value.

### HasFirstRecoverabilityPointAt

`func (o *ClusterDetail) HasFirstRecoverabilityPointAt() bool`

HasFirstRecoverabilityPointAt returns a boolean if a field has been set.

### GetAllowedIpRanges

`func (o *ClusterDetail) GetAllowedIpRanges() []AllowedIpRange`

GetAllowedIpRanges returns the AllowedIpRanges field if non-nil, zero value otherwise.

### GetAllowedIpRangesOk

`func (o *ClusterDetail) GetAllowedIpRangesOk() (*[]AllowedIpRange, bool)`

GetAllowedIpRangesOk returns a tuple with the AllowedIpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedIpRanges

`func (o *ClusterDetail) SetAllowedIpRanges(v []AllowedIpRange)`

SetAllowedIpRanges sets AllowedIpRanges field to given value.


### GetPgConfig

`func (o *ClusterDetail) GetPgConfig() []ClusterDetailPgConfigInner`

GetPgConfig returns the PgConfig field if non-nil, zero value otherwise.

### GetPgConfigOk

`func (o *ClusterDetail) GetPgConfigOk() (*[]ClusterDetailPgConfigInner, bool)`

GetPgConfigOk returns a tuple with the PgConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgConfig

`func (o *ClusterDetail) SetPgConfig(v []ClusterDetailPgConfigInner)`

SetPgConfig sets PgConfig field to given value.

### HasPgConfig

`func (o *ClusterDetail) HasPgConfig() bool`

HasPgConfig returns a boolean if a field has been set.

### GetPgType

`func (o *ClusterDetail) GetPgType() PgType`

GetPgType returns the PgType field if non-nil, zero value otherwise.

### GetPgTypeOk

`func (o *ClusterDetail) GetPgTypeOk() (*PgType, bool)`

GetPgTypeOk returns a tuple with the PgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgType

`func (o *ClusterDetail) SetPgType(v PgType)`

SetPgType sets PgType field to given value.

### HasPgType

`func (o *ClusterDetail) HasPgType() bool`

HasPgType returns a boolean if a field has been set.

### GetPgVersion

`func (o *ClusterDetail) GetPgVersion() PgVersion`

GetPgVersion returns the PgVersion field if non-nil, zero value otherwise.

### GetPgVersionOk

`func (o *ClusterDetail) GetPgVersionOk() (*PgVersion, bool)`

GetPgVersionOk returns a tuple with the PgVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgVersion

`func (o *ClusterDetail) SetPgVersion(v PgVersion)`

SetPgVersion sets PgVersion field to given value.

### HasPgVersion

`func (o *ClusterDetail) HasPgVersion() bool`

HasPgVersion returns a boolean if a field has been set.

### GetProvider

`func (o *ClusterDetail) GetProvider() CloudProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ClusterDetail) GetProviderOk() (*CloudProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ClusterDetail) SetProvider(v CloudProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ClusterDetail) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetReadOnlyConnections

`func (o *ClusterDetail) GetReadOnlyConnections() bool`

GetReadOnlyConnections returns the ReadOnlyConnections field if non-nil, zero value otherwise.

### GetReadOnlyConnectionsOk

`func (o *ClusterDetail) GetReadOnlyConnectionsOk() (*bool, bool)`

GetReadOnlyConnectionsOk returns a tuple with the ReadOnlyConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnlyConnections

`func (o *ClusterDetail) SetReadOnlyConnections(v bool)`

SetReadOnlyConnections sets ReadOnlyConnections field to given value.

### HasReadOnlyConnections

`func (o *ClusterDetail) HasReadOnlyConnections() bool`

HasReadOnlyConnections returns a boolean if a field has been set.

### GetRegion

`func (o *ClusterDetail) GetRegion() CloudProviderRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ClusterDetail) GetRegionOk() (*CloudProviderRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ClusterDetail) SetRegion(v CloudProviderRegion)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ClusterDetail) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetResizingPvc

`func (o *ClusterDetail) GetResizingPvc() []string`

GetResizingPvc returns the ResizingPvc field if non-nil, zero value otherwise.

### GetResizingPvcOk

`func (o *ClusterDetail) GetResizingPvcOk() (*[]string, bool)`

GetResizingPvcOk returns a tuple with the ResizingPvc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResizingPvc

`func (o *ClusterDetail) SetResizingPvc(v []string)`

SetResizingPvc sets ResizingPvc field to given value.


### GetConditions

`func (o *ClusterDetail) GetConditions() []ClusterDetailConditionsInner`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *ClusterDetail) GetConditionsOk() (*[]ClusterDetailConditionsInner, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *ClusterDetail) SetConditions(v []ClusterDetailConditionsInner)`

SetConditions sets Conditions field to given value.


### GetInstanceType

`func (o *ClusterDetail) GetInstanceType() CloudProviderRegionInstanceType`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *ClusterDetail) GetInstanceTypeOk() (*CloudProviderRegionInstanceType, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *ClusterDetail) SetInstanceType(v CloudProviderRegionInstanceType)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *ClusterDetail) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetStorage

`func (o *ClusterDetail) GetStorage() ClusterDetailStorage`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *ClusterDetail) GetStorageOk() (*ClusterDetailStorage, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *ClusterDetail) SetStorage(v ClusterDetailStorage)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *ClusterDetail) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetEvaluatedPgConfig

`func (o *ClusterDetail) GetEvaluatedPgConfig() []ClusterDetailPgConfigInner`

GetEvaluatedPgConfig returns the EvaluatedPgConfig field if non-nil, zero value otherwise.

### GetEvaluatedPgConfigOk

`func (o *ClusterDetail) GetEvaluatedPgConfigOk() (*[]ClusterDetailPgConfigInner, bool)`

GetEvaluatedPgConfigOk returns a tuple with the EvaluatedPgConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluatedPgConfig

`func (o *ClusterDetail) SetEvaluatedPgConfig(v []ClusterDetailPgConfigInner)`

SetEvaluatedPgConfig sets EvaluatedPgConfig field to given value.

### HasEvaluatedPgConfig

`func (o *ClusterDetail) HasEvaluatedPgConfig() bool`

HasEvaluatedPgConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


