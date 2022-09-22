# ConnectionsBrokerConnOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Saml** | Pointer to [**ConnectionsBrokerConnOptionsSaml**](ConnectionsBrokerConnOptionsSaml.md) |  | [optional] 
**Oidc** | Pointer to [**ConnectionsBrokerConnOptionsOidc**](ConnectionsBrokerConnOptionsOidc.md) |  | [optional] 

## Methods

### NewConnectionsBrokerConnOptions

`func NewConnectionsBrokerConnOptions() *ConnectionsBrokerConnOptions`

NewConnectionsBrokerConnOptions instantiates a new ConnectionsBrokerConnOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionsBrokerConnOptionsWithDefaults

`func NewConnectionsBrokerConnOptionsWithDefaults() *ConnectionsBrokerConnOptions`

NewConnectionsBrokerConnOptionsWithDefaults instantiates a new ConnectionsBrokerConnOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSaml

`func (o *ConnectionsBrokerConnOptions) GetSaml() ConnectionsBrokerConnOptionsSaml`

GetSaml returns the Saml field if non-nil, zero value otherwise.

### GetSamlOk

`func (o *ConnectionsBrokerConnOptions) GetSamlOk() (*ConnectionsBrokerConnOptionsSaml, bool)`

GetSamlOk returns a tuple with the Saml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml

`func (o *ConnectionsBrokerConnOptions) SetSaml(v ConnectionsBrokerConnOptionsSaml)`

SetSaml sets Saml field to given value.

### HasSaml

`func (o *ConnectionsBrokerConnOptions) HasSaml() bool`

HasSaml returns a boolean if a field has been set.

### GetOidc

`func (o *ConnectionsBrokerConnOptions) GetOidc() ConnectionsBrokerConnOptionsOidc`

GetOidc returns the Oidc field if non-nil, zero value otherwise.

### GetOidcOk

`func (o *ConnectionsBrokerConnOptions) GetOidcOk() (*ConnectionsBrokerConnOptionsOidc, bool)`

GetOidcOk returns a tuple with the Oidc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidc

`func (o *ConnectionsBrokerConnOptions) SetOidc(v ConnectionsBrokerConnOptionsOidc)`

SetOidc sets Oidc field to given value.

### HasOidc

`func (o *ConnectionsBrokerConnOptions) HasOidc() bool`

HasOidc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


