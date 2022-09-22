# RegisterAws

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | **string** |  | 
**RoleArn** | **string** |  | 
**PolicyArn** | Pointer to **string** |  | [optional] 

## Methods

### NewRegisterAws

`func NewRegisterAws(externalId string, roleArn string, ) *RegisterAws`

NewRegisterAws instantiates a new RegisterAws object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterAwsWithDefaults

`func NewRegisterAwsWithDefaults() *RegisterAws`

NewRegisterAwsWithDefaults instantiates a new RegisterAws object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *RegisterAws) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *RegisterAws) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *RegisterAws) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetRoleArn

`func (o *RegisterAws) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *RegisterAws) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *RegisterAws) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.


### GetPolicyArn

`func (o *RegisterAws) GetPolicyArn() string`

GetPolicyArn returns the PolicyArn field if non-nil, zero value otherwise.

### GetPolicyArnOk

`func (o *RegisterAws) GetPolicyArnOk() (*string, bool)`

GetPolicyArnOk returns a tuple with the PolicyArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyArn

`func (o *RegisterAws) SetPolicyArn(v string)`

SetPolicyArn sets PolicyArn field to given value.

### HasPolicyArn

`func (o *RegisterAws) HasPolicyArn() bool`

HasPolicyArn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


