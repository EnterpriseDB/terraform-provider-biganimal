# GetConnectionList200ResponseBrokerConfigInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Saml** | Pointer to [**GetConnectionList200ResponseBrokerConfigInfoSaml**](GetConnectionList200ResponseBrokerConfigInfoSaml.md) |  | [optional] 
**Oidc** | Pointer to [**GetConnectionList200ResponseBrokerConfigInfoOidc**](GetConnectionList200ResponseBrokerConfigInfoOidc.md) |  | [optional] 

## Methods

### NewGetConnectionList200ResponseBrokerConfigInfo

`func NewGetConnectionList200ResponseBrokerConfigInfo() *GetConnectionList200ResponseBrokerConfigInfo`

NewGetConnectionList200ResponseBrokerConfigInfo instantiates a new GetConnectionList200ResponseBrokerConfigInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConnectionList200ResponseBrokerConfigInfoWithDefaults

`func NewGetConnectionList200ResponseBrokerConfigInfoWithDefaults() *GetConnectionList200ResponseBrokerConfigInfo`

NewGetConnectionList200ResponseBrokerConfigInfoWithDefaults instantiates a new GetConnectionList200ResponseBrokerConfigInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSaml

`func (o *GetConnectionList200ResponseBrokerConfigInfo) GetSaml() GetConnectionList200ResponseBrokerConfigInfoSaml`

GetSaml returns the Saml field if non-nil, zero value otherwise.

### GetSamlOk

`func (o *GetConnectionList200ResponseBrokerConfigInfo) GetSamlOk() (*GetConnectionList200ResponseBrokerConfigInfoSaml, bool)`

GetSamlOk returns a tuple with the Saml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml

`func (o *GetConnectionList200ResponseBrokerConfigInfo) SetSaml(v GetConnectionList200ResponseBrokerConfigInfoSaml)`

SetSaml sets Saml field to given value.

### HasSaml

`func (o *GetConnectionList200ResponseBrokerConfigInfo) HasSaml() bool`

HasSaml returns a boolean if a field has been set.

### GetOidc

`func (o *GetConnectionList200ResponseBrokerConfigInfo) GetOidc() GetConnectionList200ResponseBrokerConfigInfoOidc`

GetOidc returns the Oidc field if non-nil, zero value otherwise.

### GetOidcOk

`func (o *GetConnectionList200ResponseBrokerConfigInfo) GetOidcOk() (*GetConnectionList200ResponseBrokerConfigInfoOidc, bool)`

GetOidcOk returns a tuple with the Oidc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidc

`func (o *GetConnectionList200ResponseBrokerConfigInfo) SetOidc(v GetConnectionList200ResponseBrokerConfigInfoOidc)`

SetOidc sets Oidc field to given value.

### HasOidc

`func (o *GetConnectionList200ResponseBrokerConfigInfo) HasOidc() bool`

HasOidc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


