# RunSearchForEvents200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Event**](Event.md) |  | 
**Paging** | [**RunSearchForEvents200ResponsePaging**](RunSearchForEvents200ResponsePaging.md) |  | 

## Methods

### NewRunSearchForEvents200Response

`func NewRunSearchForEvents200Response(data []Event, paging RunSearchForEvents200ResponsePaging, ) *RunSearchForEvents200Response`

NewRunSearchForEvents200Response instantiates a new RunSearchForEvents200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunSearchForEvents200ResponseWithDefaults

`func NewRunSearchForEvents200ResponseWithDefaults() *RunSearchForEvents200Response`

NewRunSearchForEvents200ResponseWithDefaults instantiates a new RunSearchForEvents200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RunSearchForEvents200Response) GetData() []Event`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RunSearchForEvents200Response) GetDataOk() (*[]Event, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RunSearchForEvents200Response) SetData(v []Event)`

SetData sets Data field to given value.


### GetPaging

`func (o *RunSearchForEvents200Response) GetPaging() RunSearchForEvents200ResponsePaging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *RunSearchForEvents200Response) GetPagingOk() (*RunSearchForEvents200ResponsePaging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *RunSearchForEvents200Response) SetPaging(v RunSearchForEvents200ResponsePaging)`

SetPaging sets Paging field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


