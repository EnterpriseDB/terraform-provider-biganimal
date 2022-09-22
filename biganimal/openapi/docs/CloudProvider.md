# CloudProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProviderId** | **string** |  | 
**CloudProviderName** | **string** |  | 
**Connected** | Pointer to **bool** |  | [optional] 

## Methods

### NewCloudProvider

`func NewCloudProvider(cloudProviderId string, cloudProviderName string, ) *CloudProvider`

NewCloudProvider instantiates a new CloudProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudProviderWithDefaults

`func NewCloudProviderWithDefaults() *CloudProvider`

NewCloudProviderWithDefaults instantiates a new CloudProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProviderId

`func (o *CloudProvider) GetCloudProviderId() string`

GetCloudProviderId returns the CloudProviderId field if non-nil, zero value otherwise.

### GetCloudProviderIdOk

`func (o *CloudProvider) GetCloudProviderIdOk() (*string, bool)`

GetCloudProviderIdOk returns a tuple with the CloudProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderId

`func (o *CloudProvider) SetCloudProviderId(v string)`

SetCloudProviderId sets CloudProviderId field to given value.


### GetCloudProviderName

`func (o *CloudProvider) GetCloudProviderName() string`

GetCloudProviderName returns the CloudProviderName field if non-nil, zero value otherwise.

### GetCloudProviderNameOk

`func (o *CloudProvider) GetCloudProviderNameOk() (*string, bool)`

GetCloudProviderNameOk returns a tuple with the CloudProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderName

`func (o *CloudProvider) SetCloudProviderName(v string)`

SetCloudProviderName sets CloudProviderName field to given value.


### GetConnected

`func (o *CloudProvider) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *CloudProvider) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *CloudProvider) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *CloudProvider) HasConnected() bool`

HasConnected returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


