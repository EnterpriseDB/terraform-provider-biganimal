# ClusterPgConfigParameter

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
**CurrentValue** | **string** |  | 

## Methods

### NewClusterPgConfigParameter

`func NewClusterPgConfigParameter(name string, unit string, varType string, bootValue string, resetValue string, readOnly bool, readOnlyValue string, currentValue string, ) *ClusterPgConfigParameter`

NewClusterPgConfigParameter instantiates a new ClusterPgConfigParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterPgConfigParameterWithDefaults

`func NewClusterPgConfigParameterWithDefaults() *ClusterPgConfigParameter`

NewClusterPgConfigParameterWithDefaults instantiates a new ClusterPgConfigParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClusterPgConfigParameter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterPgConfigParameter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterPgConfigParameter) SetName(v string)`

SetName sets Name field to given value.


### GetUnit

`func (o *ClusterPgConfigParameter) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ClusterPgConfigParameter) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ClusterPgConfigParameter) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetCategory

`func (o *ClusterPgConfigParameter) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ClusterPgConfigParameter) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ClusterPgConfigParameter) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ClusterPgConfigParameter) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetShortDesc

`func (o *ClusterPgConfigParameter) GetShortDesc() string`

GetShortDesc returns the ShortDesc field if non-nil, zero value otherwise.

### GetShortDescOk

`func (o *ClusterPgConfigParameter) GetShortDescOk() (*string, bool)`

GetShortDescOk returns a tuple with the ShortDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDesc

`func (o *ClusterPgConfigParameter) SetShortDesc(v string)`

SetShortDesc sets ShortDesc field to given value.

### HasShortDesc

`func (o *ClusterPgConfigParameter) HasShortDesc() bool`

HasShortDesc returns a boolean if a field has been set.

### GetExtraDesc

`func (o *ClusterPgConfigParameter) GetExtraDesc() string`

GetExtraDesc returns the ExtraDesc field if non-nil, zero value otherwise.

### GetExtraDescOk

`func (o *ClusterPgConfigParameter) GetExtraDescOk() (*string, bool)`

GetExtraDescOk returns a tuple with the ExtraDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraDesc

`func (o *ClusterPgConfigParameter) SetExtraDesc(v string)`

SetExtraDesc sets ExtraDesc field to given value.

### HasExtraDesc

`func (o *ClusterPgConfigParameter) HasExtraDesc() bool`

HasExtraDesc returns a boolean if a field has been set.

### GetContext

`func (o *ClusterPgConfigParameter) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ClusterPgConfigParameter) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ClusterPgConfigParameter) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *ClusterPgConfigParameter) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetInitDB

`func (o *ClusterPgConfigParameter) GetInitDB() bool`

GetInitDB returns the InitDB field if non-nil, zero value otherwise.

### GetInitDBOk

`func (o *ClusterPgConfigParameter) GetInitDBOk() (*bool, bool)`

GetInitDBOk returns a tuple with the InitDB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitDB

`func (o *ClusterPgConfigParameter) SetInitDB(v bool)`

SetInitDB sets InitDB field to given value.

### HasInitDB

`func (o *ClusterPgConfigParameter) HasInitDB() bool`

HasInitDB returns a boolean if a field has been set.

### GetVarType

`func (o *ClusterPgConfigParameter) GetVarType() string`

GetVarType returns the VarType field if non-nil, zero value otherwise.

### GetVarTypeOk

`func (o *ClusterPgConfigParameter) GetVarTypeOk() (*string, bool)`

GetVarTypeOk returns a tuple with the VarType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVarType

`func (o *ClusterPgConfigParameter) SetVarType(v string)`

SetVarType sets VarType field to given value.


### GetSource

`func (o *ClusterPgConfigParameter) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ClusterPgConfigParameter) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ClusterPgConfigParameter) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ClusterPgConfigParameter) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetMinValue

`func (o *ClusterPgConfigParameter) GetMinValue() string`

GetMinValue returns the MinValue field if non-nil, zero value otherwise.

### GetMinValueOk

`func (o *ClusterPgConfigParameter) GetMinValueOk() (*string, bool)`

GetMinValueOk returns a tuple with the MinValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValue

`func (o *ClusterPgConfigParameter) SetMinValue(v string)`

SetMinValue sets MinValue field to given value.

### HasMinValue

`func (o *ClusterPgConfigParameter) HasMinValue() bool`

HasMinValue returns a boolean if a field has been set.

### GetMaxValue

`func (o *ClusterPgConfigParameter) GetMaxValue() string`

GetMaxValue returns the MaxValue field if non-nil, zero value otherwise.

### GetMaxValueOk

`func (o *ClusterPgConfigParameter) GetMaxValueOk() (*string, bool)`

GetMaxValueOk returns a tuple with the MaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValue

`func (o *ClusterPgConfigParameter) SetMaxValue(v string)`

SetMaxValue sets MaxValue field to given value.

### HasMaxValue

`func (o *ClusterPgConfigParameter) HasMaxValue() bool`

HasMaxValue returns a boolean if a field has been set.

### GetEnumValues

`func (o *ClusterPgConfigParameter) GetEnumValues() []string`

GetEnumValues returns the EnumValues field if non-nil, zero value otherwise.

### GetEnumValuesOk

`func (o *ClusterPgConfigParameter) GetEnumValuesOk() (*[]string, bool)`

GetEnumValuesOk returns a tuple with the EnumValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnumValues

`func (o *ClusterPgConfigParameter) SetEnumValues(v []string)`

SetEnumValues sets EnumValues field to given value.

### HasEnumValues

`func (o *ClusterPgConfigParameter) HasEnumValues() bool`

HasEnumValues returns a boolean if a field has been set.

### GetCnpoControlled

`func (o *ClusterPgConfigParameter) GetCnpoControlled() bool`

GetCnpoControlled returns the CnpoControlled field if non-nil, zero value otherwise.

### GetCnpoControlledOk

`func (o *ClusterPgConfigParameter) GetCnpoControlledOk() (*bool, bool)`

GetCnpoControlledOk returns a tuple with the CnpoControlled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnpoControlled

`func (o *ClusterPgConfigParameter) SetCnpoControlled(v bool)`

SetCnpoControlled sets CnpoControlled field to given value.

### HasCnpoControlled

`func (o *ClusterPgConfigParameter) HasCnpoControlled() bool`

HasCnpoControlled returns a boolean if a field has been set.

### GetBootValue

`func (o *ClusterPgConfigParameter) GetBootValue() string`

GetBootValue returns the BootValue field if non-nil, zero value otherwise.

### GetBootValueOk

`func (o *ClusterPgConfigParameter) GetBootValueOk() (*string, bool)`

GetBootValueOk returns a tuple with the BootValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootValue

`func (o *ClusterPgConfigParameter) SetBootValue(v string)`

SetBootValue sets BootValue field to given value.


### GetResetValue

`func (o *ClusterPgConfigParameter) GetResetValue() string`

GetResetValue returns the ResetValue field if non-nil, zero value otherwise.

### GetResetValueOk

`func (o *ClusterPgConfigParameter) GetResetValueOk() (*string, bool)`

GetResetValueOk returns a tuple with the ResetValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetValue

`func (o *ClusterPgConfigParameter) SetResetValue(v string)`

SetResetValue sets ResetValue field to given value.


### GetReadOnly

`func (o *ClusterPgConfigParameter) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *ClusterPgConfigParameter) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *ClusterPgConfigParameter) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.


### GetReadOnlyValue

`func (o *ClusterPgConfigParameter) GetReadOnlyValue() string`

GetReadOnlyValue returns the ReadOnlyValue field if non-nil, zero value otherwise.

### GetReadOnlyValueOk

`func (o *ClusterPgConfigParameter) GetReadOnlyValueOk() (*string, bool)`

GetReadOnlyValueOk returns a tuple with the ReadOnlyValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnlyValue

`func (o *ClusterPgConfigParameter) SetReadOnlyValue(v string)`

SetReadOnlyValue sets ReadOnlyValue field to given value.


### GetCurrentValue

`func (o *ClusterPgConfigParameter) GetCurrentValue() string`

GetCurrentValue returns the CurrentValue field if non-nil, zero value otherwise.

### GetCurrentValueOk

`func (o *ClusterPgConfigParameter) GetCurrentValueOk() (*string, bool)`

GetCurrentValueOk returns a tuple with the CurrentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentValue

`func (o *ClusterPgConfigParameter) SetCurrentValue(v string)`

SetCurrentValue sets CurrentValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


