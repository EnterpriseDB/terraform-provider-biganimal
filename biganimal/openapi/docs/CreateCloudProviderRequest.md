# CreateCloudProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | **string** |  | 
**RoleArn** | **string** |  | 
**PolicyArn** | Pointer to **string** |  | [optional] 
**ClientId** | **string** |  | 
**ClientSecret** | **string** |  | 
**SubscriptionId** | **string** |  | 
**TenantId** | **string** |  | 

## Methods

### NewCreateCloudProviderRequest

`func NewCreateCloudProviderRequest(externalId string, roleArn string, clientId string, clientSecret string, subscriptionId string, tenantId string, ) *CreateCloudProviderRequest`

NewCreateCloudProviderRequest instantiates a new CreateCloudProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCloudProviderRequestWithDefaults

`func NewCreateCloudProviderRequestWithDefaults() *CreateCloudProviderRequest`

NewCreateCloudProviderRequestWithDefaults instantiates a new CreateCloudProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *CreateCloudProviderRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CreateCloudProviderRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CreateCloudProviderRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetRoleArn

`func (o *CreateCloudProviderRequest) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *CreateCloudProviderRequest) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *CreateCloudProviderRequest) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.


### GetPolicyArn

`func (o *CreateCloudProviderRequest) GetPolicyArn() string`

GetPolicyArn returns the PolicyArn field if non-nil, zero value otherwise.

### GetPolicyArnOk

`func (o *CreateCloudProviderRequest) GetPolicyArnOk() (*string, bool)`

GetPolicyArnOk returns a tuple with the PolicyArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyArn

`func (o *CreateCloudProviderRequest) SetPolicyArn(v string)`

SetPolicyArn sets PolicyArn field to given value.

### HasPolicyArn

`func (o *CreateCloudProviderRequest) HasPolicyArn() bool`

HasPolicyArn returns a boolean if a field has been set.

### GetClientId

`func (o *CreateCloudProviderRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CreateCloudProviderRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CreateCloudProviderRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *CreateCloudProviderRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *CreateCloudProviderRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *CreateCloudProviderRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetSubscriptionId

`func (o *CreateCloudProviderRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *CreateCloudProviderRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *CreateCloudProviderRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetTenantId

`func (o *CreateCloudProviderRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CreateCloudProviderRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CreateCloudProviderRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


