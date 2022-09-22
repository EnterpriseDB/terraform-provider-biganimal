# GetConnectionList200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] 
**BrokerConnectionId** | Pointer to **string** |  | [optional] 
**BrokerConnectionName** | Pointer to **string** |  | [optional] 
**Strategy** | Pointer to **string** |  | [optional] 
**BrokerConnOptions** | Pointer to [**GetConnectionList200ResponseBrokerConnOptions**](GetConnectionList200ResponseBrokerConnOptions.md) |  | [optional] 
**BrokerConfigInfo** | Pointer to [**GetConnectionList200ResponseBrokerConfigInfo**](GetConnectionList200ResponseBrokerConfigInfo.md) |  | [optional] 
**Validated** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetConnectionList200Response

`func NewGetConnectionList200Response() *GetConnectionList200Response`

NewGetConnectionList200Response instantiates a new GetConnectionList200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConnectionList200ResponseWithDefaults

`func NewGetConnectionList200ResponseWithDefaults() *GetConnectionList200Response`

NewGetConnectionList200ResponseWithDefaults instantiates a new GetConnectionList200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetConnectionList200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetConnectionList200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetConnectionList200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetConnectionList200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *GetConnectionList200Response) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *GetConnectionList200Response) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *GetConnectionList200Response) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *GetConnectionList200Response) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetBrokerConnectionId

`func (o *GetConnectionList200Response) GetBrokerConnectionId() string`

GetBrokerConnectionId returns the BrokerConnectionId field if non-nil, zero value otherwise.

### GetBrokerConnectionIdOk

`func (o *GetConnectionList200Response) GetBrokerConnectionIdOk() (*string, bool)`

GetBrokerConnectionIdOk returns a tuple with the BrokerConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerConnectionId

`func (o *GetConnectionList200Response) SetBrokerConnectionId(v string)`

SetBrokerConnectionId sets BrokerConnectionId field to given value.

### HasBrokerConnectionId

`func (o *GetConnectionList200Response) HasBrokerConnectionId() bool`

HasBrokerConnectionId returns a boolean if a field has been set.

### GetBrokerConnectionName

`func (o *GetConnectionList200Response) GetBrokerConnectionName() string`

GetBrokerConnectionName returns the BrokerConnectionName field if non-nil, zero value otherwise.

### GetBrokerConnectionNameOk

`func (o *GetConnectionList200Response) GetBrokerConnectionNameOk() (*string, bool)`

GetBrokerConnectionNameOk returns a tuple with the BrokerConnectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerConnectionName

`func (o *GetConnectionList200Response) SetBrokerConnectionName(v string)`

SetBrokerConnectionName sets BrokerConnectionName field to given value.

### HasBrokerConnectionName

`func (o *GetConnectionList200Response) HasBrokerConnectionName() bool`

HasBrokerConnectionName returns a boolean if a field has been set.

### GetStrategy

`func (o *GetConnectionList200Response) GetStrategy() string`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *GetConnectionList200Response) GetStrategyOk() (*string, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *GetConnectionList200Response) SetStrategy(v string)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *GetConnectionList200Response) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### GetBrokerConnOptions

`func (o *GetConnectionList200Response) GetBrokerConnOptions() GetConnectionList200ResponseBrokerConnOptions`

GetBrokerConnOptions returns the BrokerConnOptions field if non-nil, zero value otherwise.

### GetBrokerConnOptionsOk

`func (o *GetConnectionList200Response) GetBrokerConnOptionsOk() (*GetConnectionList200ResponseBrokerConnOptions, bool)`

GetBrokerConnOptionsOk returns a tuple with the BrokerConnOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerConnOptions

`func (o *GetConnectionList200Response) SetBrokerConnOptions(v GetConnectionList200ResponseBrokerConnOptions)`

SetBrokerConnOptions sets BrokerConnOptions field to given value.

### HasBrokerConnOptions

`func (o *GetConnectionList200Response) HasBrokerConnOptions() bool`

HasBrokerConnOptions returns a boolean if a field has been set.

### GetBrokerConfigInfo

`func (o *GetConnectionList200Response) GetBrokerConfigInfo() GetConnectionList200ResponseBrokerConfigInfo`

GetBrokerConfigInfo returns the BrokerConfigInfo field if non-nil, zero value otherwise.

### GetBrokerConfigInfoOk

`func (o *GetConnectionList200Response) GetBrokerConfigInfoOk() (*GetConnectionList200ResponseBrokerConfigInfo, bool)`

GetBrokerConfigInfoOk returns a tuple with the BrokerConfigInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerConfigInfo

`func (o *GetConnectionList200Response) SetBrokerConfigInfo(v GetConnectionList200ResponseBrokerConfigInfo)`

SetBrokerConfigInfo sets BrokerConfigInfo field to given value.

### HasBrokerConfigInfo

`func (o *GetConnectionList200Response) HasBrokerConfigInfo() bool`

HasBrokerConfigInfo returns a boolean if a field has been set.

### GetValidated

`func (o *GetConnectionList200Response) GetValidated() bool`

GetValidated returns the Validated field if non-nil, zero value otherwise.

### GetValidatedOk

`func (o *GetConnectionList200Response) GetValidatedOk() (*bool, bool)`

GetValidatedOk returns a tuple with the Validated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidated

`func (o *GetConnectionList200Response) SetValidated(v bool)`

SetValidated sets Validated field to given value.

### HasValidated

`func (o *GetConnectionList200Response) HasValidated() bool`

HasValidated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


