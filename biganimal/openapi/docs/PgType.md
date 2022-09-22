# PgType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PgTypeId** | **string** |  | 
**PgTypeName** | **string** |  | 
**SupportedClusterArchitectureIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPgType

`func NewPgType(pgTypeId string, pgTypeName string, ) *PgType`

NewPgType instantiates a new PgType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPgTypeWithDefaults

`func NewPgTypeWithDefaults() *PgType`

NewPgTypeWithDefaults instantiates a new PgType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPgTypeId

`func (o *PgType) GetPgTypeId() string`

GetPgTypeId returns the PgTypeId field if non-nil, zero value otherwise.

### GetPgTypeIdOk

`func (o *PgType) GetPgTypeIdOk() (*string, bool)`

GetPgTypeIdOk returns a tuple with the PgTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgTypeId

`func (o *PgType) SetPgTypeId(v string)`

SetPgTypeId sets PgTypeId field to given value.


### GetPgTypeName

`func (o *PgType) GetPgTypeName() string`

GetPgTypeName returns the PgTypeName field if non-nil, zero value otherwise.

### GetPgTypeNameOk

`func (o *PgType) GetPgTypeNameOk() (*string, bool)`

GetPgTypeNameOk returns a tuple with the PgTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgTypeName

`func (o *PgType) SetPgTypeName(v string)`

SetPgTypeName sets PgTypeName field to given value.


### GetSupportedClusterArchitectureIds

`func (o *PgType) GetSupportedClusterArchitectureIds() []string`

GetSupportedClusterArchitectureIds returns the SupportedClusterArchitectureIds field if non-nil, zero value otherwise.

### GetSupportedClusterArchitectureIdsOk

`func (o *PgType) GetSupportedClusterArchitectureIdsOk() (*[]string, bool)`

GetSupportedClusterArchitectureIdsOk returns a tuple with the SupportedClusterArchitectureIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedClusterArchitectureIds

`func (o *PgType) SetSupportedClusterArchitectureIds(v []string)`

SetSupportedClusterArchitectureIds sets SupportedClusterArchitectureIds field to given value.

### HasSupportedClusterArchitectureIds

`func (o *PgType) HasSupportedClusterArchitectureIds() bool`

HasSupportedClusterArchitectureIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


