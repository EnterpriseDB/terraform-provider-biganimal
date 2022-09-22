# GetConnectionList200ResponseBrokerConnOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Saml** | Pointer to [**GetConnectionList200ResponseBrokerConnOptionsSaml**](GetConnectionList200ResponseBrokerConnOptionsSaml.md) |  | [optional] 
**Oidc** | Pointer to [**GetConnectionList200ResponseBrokerConnOptionsOidc**](GetConnectionList200ResponseBrokerConnOptionsOidc.md) |  | [optional] 

## Methods

### NewGetConnectionList200ResponseBrokerConnOptions

`func NewGetConnectionList200ResponseBrokerConnOptions() *GetConnectionList200ResponseBrokerConnOptions`

NewGetConnectionList200ResponseBrokerConnOptions instantiates a new GetConnectionList200ResponseBrokerConnOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConnectionList200ResponseBrokerConnOptionsWithDefaults

`func NewGetConnectionList200ResponseBrokerConnOptionsWithDefaults() *GetConnectionList200ResponseBrokerConnOptions`

NewGetConnectionList200ResponseBrokerConnOptionsWithDefaults instantiates a new GetConnectionList200ResponseBrokerConnOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSaml

`func (o *GetConnectionList200ResponseBrokerConnOptions) GetSaml() GetConnectionList200ResponseBrokerConnOptionsSaml`

GetSaml returns the Saml field if non-nil, zero value otherwise.

### GetSamlOk

`func (o *GetConnectionList200ResponseBrokerConnOptions) GetSamlOk() (*GetConnectionList200ResponseBrokerConnOptionsSaml, bool)`

GetSamlOk returns a tuple with the Saml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml

`func (o *GetConnectionList200ResponseBrokerConnOptions) SetSaml(v GetConnectionList200ResponseBrokerConnOptionsSaml)`

SetSaml sets Saml field to given value.

### HasSaml

`func (o *GetConnectionList200ResponseBrokerConnOptions) HasSaml() bool`

HasSaml returns a boolean if a field has been set.

### GetOidc

`func (o *GetConnectionList200ResponseBrokerConnOptions) GetOidc() GetConnectionList200ResponseBrokerConnOptionsOidc`

GetOidc returns the Oidc field if non-nil, zero value otherwise.

### GetOidcOk

`func (o *GetConnectionList200ResponseBrokerConnOptions) GetOidcOk() (*GetConnectionList200ResponseBrokerConnOptionsOidc, bool)`

GetOidcOk returns a tuple with the Oidc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidc

`func (o *GetConnectionList200ResponseBrokerConnOptions) SetOidc(v GetConnectionList200ResponseBrokerConnOptionsOidc)`

SetOidc sets Oidc field to given value.

### HasOidc

`func (o *GetConnectionList200ResponseBrokerConnOptions) HasOidc() bool`

HasOidc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


