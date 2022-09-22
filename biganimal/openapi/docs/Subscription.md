# Subscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrincipalName** | **string** |  | 
**PrincipalSecret** | **string** |  | 
**AzSubId** | **string** |  | 
**AzResourceGroup** | Pointer to **string** |  | [optional] 
**OrgName** | **string** |  | 

## Methods

### NewSubscription

`func NewSubscription(principalName string, principalSecret string, azSubId string, orgName string, ) *Subscription`

NewSubscription instantiates a new Subscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionWithDefaults

`func NewSubscriptionWithDefaults() *Subscription`

NewSubscriptionWithDefaults instantiates a new Subscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipalName

`func (o *Subscription) GetPrincipalName() string`

GetPrincipalName returns the PrincipalName field if non-nil, zero value otherwise.

### GetPrincipalNameOk

`func (o *Subscription) GetPrincipalNameOk() (*string, bool)`

GetPrincipalNameOk returns a tuple with the PrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalName

`func (o *Subscription) SetPrincipalName(v string)`

SetPrincipalName sets PrincipalName field to given value.


### GetPrincipalSecret

`func (o *Subscription) GetPrincipalSecret() string`

GetPrincipalSecret returns the PrincipalSecret field if non-nil, zero value otherwise.

### GetPrincipalSecretOk

`func (o *Subscription) GetPrincipalSecretOk() (*string, bool)`

GetPrincipalSecretOk returns a tuple with the PrincipalSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalSecret

`func (o *Subscription) SetPrincipalSecret(v string)`

SetPrincipalSecret sets PrincipalSecret field to given value.


### GetAzSubId

`func (o *Subscription) GetAzSubId() string`

GetAzSubId returns the AzSubId field if non-nil, zero value otherwise.

### GetAzSubIdOk

`func (o *Subscription) GetAzSubIdOk() (*string, bool)`

GetAzSubIdOk returns a tuple with the AzSubId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzSubId

`func (o *Subscription) SetAzSubId(v string)`

SetAzSubId sets AzSubId field to given value.


### GetAzResourceGroup

`func (o *Subscription) GetAzResourceGroup() string`

GetAzResourceGroup returns the AzResourceGroup field if non-nil, zero value otherwise.

### GetAzResourceGroupOk

`func (o *Subscription) GetAzResourceGroupOk() (*string, bool)`

GetAzResourceGroupOk returns a tuple with the AzResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzResourceGroup

`func (o *Subscription) SetAzResourceGroup(v string)`

SetAzResourceGroup sets AzResourceGroup field to given value.

### HasAzResourceGroup

`func (o *Subscription) HasAzResourceGroup() bool`

HasAzResourceGroup returns a boolean if a field has been set.

### GetOrgName

`func (o *Subscription) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *Subscription) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *Subscription) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


