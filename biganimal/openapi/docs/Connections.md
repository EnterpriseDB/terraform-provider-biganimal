# Connections

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | **string** |  | 
**Strategy** | **string** |  | 
**BrokerConnOptions** | [**ConnectionsBrokerConnOptions**](ConnectionsBrokerConnOptions.md) |  | 

## Methods

### NewConnections

`func NewConnections(connectionId string, strategy string, brokerConnOptions ConnectionsBrokerConnOptions, ) *Connections`

NewConnections instantiates a new Connections object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionsWithDefaults

`func NewConnectionsWithDefaults() *Connections`

NewConnectionsWithDefaults instantiates a new Connections object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *Connections) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *Connections) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *Connections) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetStrategy

`func (o *Connections) GetStrategy() string`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *Connections) GetStrategyOk() (*string, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *Connections) SetStrategy(v string)`

SetStrategy sets Strategy field to given value.


### GetBrokerConnOptions

`func (o *Connections) GetBrokerConnOptions() ConnectionsBrokerConnOptions`

GetBrokerConnOptions returns the BrokerConnOptions field if non-nil, zero value otherwise.

### GetBrokerConnOptionsOk

`func (o *Connections) GetBrokerConnOptionsOk() (*ConnectionsBrokerConnOptions, bool)`

GetBrokerConnOptionsOk returns a tuple with the BrokerConnOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerConnOptions

`func (o *Connections) SetBrokerConnOptions(v ConnectionsBrokerConnOptions)`

SetBrokerConnOptions sets BrokerConnOptions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


