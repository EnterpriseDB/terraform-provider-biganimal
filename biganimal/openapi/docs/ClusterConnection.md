# ClusterConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseName** | **string** |  | 
**PgUri** | Pointer to **string** |  | [optional] 
**Port** | **string** |  | 
**ReadOnlyPgUri** | Pointer to **string** |  | [optional] 
**ReadOnlyPort** | Pointer to **string** |  | [optional] 
**ReadOnlyServiceName** | Pointer to **string** |  | [optional] 
**ServiceName** | **string** |  | 
**Username** | **string** |  | 

## Methods

### NewClusterConnection

`func NewClusterConnection(databaseName string, port string, serviceName string, username string, ) *ClusterConnection`

NewClusterConnection instantiates a new ClusterConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterConnectionWithDefaults

`func NewClusterConnectionWithDefaults() *ClusterConnection`

NewClusterConnectionWithDefaults instantiates a new ClusterConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseName

`func (o *ClusterConnection) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *ClusterConnection) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *ClusterConnection) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.


### GetPgUri

`func (o *ClusterConnection) GetPgUri() string`

GetPgUri returns the PgUri field if non-nil, zero value otherwise.

### GetPgUriOk

`func (o *ClusterConnection) GetPgUriOk() (*string, bool)`

GetPgUriOk returns a tuple with the PgUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgUri

`func (o *ClusterConnection) SetPgUri(v string)`

SetPgUri sets PgUri field to given value.

### HasPgUri

`func (o *ClusterConnection) HasPgUri() bool`

HasPgUri returns a boolean if a field has been set.

### GetPort

`func (o *ClusterConnection) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ClusterConnection) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ClusterConnection) SetPort(v string)`

SetPort sets Port field to given value.


### GetReadOnlyPgUri

`func (o *ClusterConnection) GetReadOnlyPgUri() string`

GetReadOnlyPgUri returns the ReadOnlyPgUri field if non-nil, zero value otherwise.

### GetReadOnlyPgUriOk

`func (o *ClusterConnection) GetReadOnlyPgUriOk() (*string, bool)`

GetReadOnlyPgUriOk returns a tuple with the ReadOnlyPgUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnlyPgUri

`func (o *ClusterConnection) SetReadOnlyPgUri(v string)`

SetReadOnlyPgUri sets ReadOnlyPgUri field to given value.

### HasReadOnlyPgUri

`func (o *ClusterConnection) HasReadOnlyPgUri() bool`

HasReadOnlyPgUri returns a boolean if a field has been set.

### GetReadOnlyPort

`func (o *ClusterConnection) GetReadOnlyPort() string`

GetReadOnlyPort returns the ReadOnlyPort field if non-nil, zero value otherwise.

### GetReadOnlyPortOk

`func (o *ClusterConnection) GetReadOnlyPortOk() (*string, bool)`

GetReadOnlyPortOk returns a tuple with the ReadOnlyPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnlyPort

`func (o *ClusterConnection) SetReadOnlyPort(v string)`

SetReadOnlyPort sets ReadOnlyPort field to given value.

### HasReadOnlyPort

`func (o *ClusterConnection) HasReadOnlyPort() bool`

HasReadOnlyPort returns a boolean if a field has been set.

### GetReadOnlyServiceName

`func (o *ClusterConnection) GetReadOnlyServiceName() string`

GetReadOnlyServiceName returns the ReadOnlyServiceName field if non-nil, zero value otherwise.

### GetReadOnlyServiceNameOk

`func (o *ClusterConnection) GetReadOnlyServiceNameOk() (*string, bool)`

GetReadOnlyServiceNameOk returns a tuple with the ReadOnlyServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnlyServiceName

`func (o *ClusterConnection) SetReadOnlyServiceName(v string)`

SetReadOnlyServiceName sets ReadOnlyServiceName field to given value.

### HasReadOnlyServiceName

`func (o *ClusterConnection) HasReadOnlyServiceName() bool`

HasReadOnlyServiceName returns a boolean if a field has been set.

### GetServiceName

`func (o *ClusterConnection) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ClusterConnection) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ClusterConnection) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetUsername

`func (o *ClusterConnection) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ClusterConnection) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ClusterConnection) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


