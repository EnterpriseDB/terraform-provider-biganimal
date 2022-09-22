# PgConfigParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Unit** | **string** |  | 
**Category** | Pointer to **string** |  | [optional] 
**ShortDesc** | Pointer to **string** |  | [optional] 
**ExtraDesc** | Pointer to **string** |  | [optional] 
**Context** | Pointer to **string** |  | [optional] 
**InitDB** | Pointer to **bool** |  | [optional] 
**VarType** | **string** |  | 
**Source** | Pointer to **string** |  | [optional] 
**MinValue** | Pointer to **string** |  | [optional] 
**MaxValue** | Pointer to **string** |  | [optional] 
**EnumValues** | Pointer to **[]string** |  | [optional] 
**CnpoControlled** | Pointer to **bool** |  | [optional] 
**BootValue** | **string** |  | 
**ResetValue** | **string** |  | 
**ReadOnly** | **bool** |  | 
**ReadOnlyValue** | **string** |  | 

## Methods

### NewPgConfigParameter

`func NewPgConfigParameter(name string, unit string, varType string, bootValue string, resetValue string, readOnly bool, readOnlyValue string, ) *PgConfigParameter`

NewPgConfigParameter instantiates a new PgConfigParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPgConfigParameterWithDefaults

`func NewPgConfigParameterWithDefaults() *PgConfigParameter`

NewPgConfigParameterWithDefaults instantiates a new PgConfigParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PgConfigParameter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PgConfigParameter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PgConfigParameter) SetName(v string)`

SetName sets Name field to given value.


### GetUnit

`func (o *PgConfigParameter) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *PgConfigParameter) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *PgConfigParameter) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetCategory

`func (o *PgConfigParameter) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *PgConfigParameter) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *PgConfigParameter) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *PgConfigParameter) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetShortDesc

`func (o *PgConfigParameter) GetShortDesc() string`

GetShortDesc returns the ShortDesc field if non-nil, zero value otherwise.

### GetShortDescOk

`func (o *PgConfigParameter) GetShortDescOk() (*string, bool)`

GetShortDescOk returns a tuple with the ShortDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDesc

`func (o *PgConfigParameter) SetShortDesc(v string)`

SetShortDesc sets ShortDesc field to given value.

### HasShortDesc

`func (o *PgConfigParameter) HasShortDesc() bool`

HasShortDesc returns a boolean if a field has been set.

### GetExtraDesc

`func (o *PgConfigParameter) GetExtraDesc() string`

GetExtraDesc returns the ExtraDesc field if non-nil, zero value otherwise.

### GetExtraDescOk

`func (o *PgConfigParameter) GetExtraDescOk() (*string, bool)`

GetExtraDescOk returns a tuple with the ExtraDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraDesc

`func (o *PgConfigParameter) SetExtraDesc(v string)`

SetExtraDesc sets ExtraDesc field to given value.

### HasExtraDesc

`func (o *PgConfigParameter) HasExtraDesc() bool`

HasExtraDesc returns a boolean if a field has been set.

### GetContext

`func (o *PgConfigParameter) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *PgConfigParameter) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *PgConfigParameter) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *PgConfigParameter) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetInitDB

`func (o *PgConfigParameter) GetInitDB() bool`

GetInitDB returns the InitDB field if non-nil, zero value otherwise.

### GetInitDBOk

`func (o *PgConfigParameter) GetInitDBOk() (*bool, bool)`

GetInitDBOk returns a tuple with the InitDB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitDB

`func (o *PgConfigParameter) SetInitDB(v bool)`

SetInitDB sets InitDB field to given value.

### HasInitDB

`func (o *PgConfigParameter) HasInitDB() bool`

HasInitDB returns a boolean if a field has been set.

### GetVarType

`func (o *PgConfigParameter) GetVarType() string`

GetVarType returns the VarType field if non-nil, zero value otherwise.

### GetVarTypeOk

`func (o *PgConfigParameter) GetVarTypeOk() (*string, bool)`

GetVarTypeOk returns a tuple with the VarType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVarType

`func (o *PgConfigParameter) SetVarType(v string)`

SetVarType sets VarType field to given value.


### GetSource

`func (o *PgConfigParameter) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PgConfigParameter) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PgConfigParameter) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *PgConfigParameter) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetMinValue

`func (o *PgConfigParameter) GetMinValue() string`

GetMinValue returns the MinValue field if non-nil, zero value otherwise.

### GetMinValueOk

`func (o *PgConfigParameter) GetMinValueOk() (*string, bool)`

GetMinValueOk returns a tuple with the MinValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValue

`func (o *PgConfigParameter) SetMinValue(v string)`

SetMinValue sets MinValue field to given value.

### HasMinValue

`func (o *PgConfigParameter) HasMinValue() bool`

HasMinValue returns a boolean if a field has been set.

### GetMaxValue

`func (o *PgConfigParameter) GetMaxValue() string`

GetMaxValue returns the MaxValue field if non-nil, zero value otherwise.

### GetMaxValueOk

`func (o *PgConfigParameter) GetMaxValueOk() (*string, bool)`

GetMaxValueOk returns a tuple with the MaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValue

`func (o *PgConfigParameter) SetMaxValue(v string)`

SetMaxValue sets MaxValue field to given value.

### HasMaxValue

`func (o *PgConfigParameter) HasMaxValue() bool`

HasMaxValue returns a boolean if a field has been set.

### GetEnumValues

`func (o *PgConfigParameter) GetEnumValues() []string`

GetEnumValues returns the EnumValues field if non-nil, zero value otherwise.

### GetEnumValuesOk

`func (o *PgConfigParameter) GetEnumValuesOk() (*[]string, bool)`

GetEnumValuesOk returns a tuple with the EnumValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnumValues

`func (o *PgConfigParameter) SetEnumValues(v []string)`

SetEnumValues sets EnumValues field to given value.

### HasEnumValues

`func (o *PgConfigParameter) HasEnumValues() bool`

HasEnumValues returns a boolean if a field has been set.

### GetCnpoControlled

`func (o *PgConfigParameter) GetCnpoControlled() bool`

GetCnpoControlled returns the CnpoControlled field if non-nil, zero value otherwise.

### GetCnpoControlledOk

`func (o *PgConfigParameter) GetCnpoControlledOk() (*bool, bool)`

GetCnpoControlledOk returns a tuple with the CnpoControlled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnpoControlled

`func (o *PgConfigParameter) SetCnpoControlled(v bool)`

SetCnpoControlled sets CnpoControlled field to given value.

### HasCnpoControlled

`func (o *PgConfigParameter) HasCnpoControlled() bool`

HasCnpoControlled returns a boolean if a field has been set.

### GetBootValue

`func (o *PgConfigParameter) GetBootValue() string`

GetBootValue returns the BootValue field if non-nil, zero value otherwise.

### GetBootValueOk

`func (o *PgConfigParameter) GetBootValueOk() (*string, bool)`

GetBootValueOk returns a tuple with the BootValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootValue

`func (o *PgConfigParameter) SetBootValue(v string)`

SetBootValue sets BootValue field to given value.


### GetResetValue

`func (o *PgConfigParameter) GetResetValue() string`

GetResetValue returns the ResetValue field if non-nil, zero value otherwise.

### GetResetValueOk

`func (o *PgConfigParameter) GetResetValueOk() (*string, bool)`

GetResetValueOk returns a tuple with the ResetValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetValue

`func (o *PgConfigParameter) SetResetValue(v string)`

SetResetValue sets ResetValue field to given value.


### GetReadOnly

`func (o *PgConfigParameter) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *PgConfigParameter) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *PgConfigParameter) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.


### GetReadOnlyValue

`func (o *PgConfigParameter) GetReadOnlyValue() string`

GetReadOnlyValue returns the ReadOnlyValue field if non-nil, zero value otherwise.

### GetReadOnlyValueOk

`func (o *PgConfigParameter) GetReadOnlyValueOk() (*string, bool)`

GetReadOnlyValueOk returns a tuple with the ReadOnlyValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnlyValue

`func (o *PgConfigParameter) SetReadOnlyValue(v string)`

SetReadOnlyValue sets ReadOnlyValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


