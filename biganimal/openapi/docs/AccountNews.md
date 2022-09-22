# AccountNews

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Url** | **string** |  | 
**CreatedAt** | [**PointInTime**](PointInTime.md) |  | 

## Methods

### NewAccountNews

`func NewAccountNews(title string, url string, createdAt PointInTime, ) *AccountNews`

NewAccountNews instantiates a new AccountNews object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountNewsWithDefaults

`func NewAccountNewsWithDefaults() *AccountNews`

NewAccountNewsWithDefaults instantiates a new AccountNews object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *AccountNews) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AccountNews) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AccountNews) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUrl

`func (o *AccountNews) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AccountNews) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AccountNews) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetCreatedAt

`func (o *AccountNews) GetCreatedAt() PointInTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccountNews) GetCreatedAtOk() (*PointInTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccountNews) SetCreatedAt(v PointInTime)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


