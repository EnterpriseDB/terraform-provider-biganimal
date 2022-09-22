# IntegrationsAzureEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The operation ID in Azure, will call with Azure operation GET API | [optional] 
**ActivityId** | Pointer to **string** | Provided by Azure. We store it for future usage | [optional] 
**SubscriptionId** | Pointer to **string** | The GUID identifier for the SaaS resource which status changes | [optional] 
**PublisherId** | Pointer to **string** | A unique string identifier for each publisher | [optional] 
**OfferId** | Pointer to **string** | A unique string identifier for each offer | [optional] 
**PlanId** | Pointer to **string** | The most up-to-date plan ID | [optional] 
**Quantity** | Pointer to **int32** | The most up-to-date number of seats, can be empty if not relevant | [optional] 
**TimeStamp** | Pointer to **string** | The UTC time when the webhook was called | [optional] 
**Action** | Pointer to **string** | The operation the webhook notifies about | [optional] 
**Status** | Pointer to **string** | Can be either InProgress or Success | [optional] 

## Methods

### NewIntegrationsAzureEvent

`func NewIntegrationsAzureEvent() *IntegrationsAzureEvent`

NewIntegrationsAzureEvent instantiates a new IntegrationsAzureEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationsAzureEventWithDefaults

`func NewIntegrationsAzureEventWithDefaults() *IntegrationsAzureEvent`

NewIntegrationsAzureEventWithDefaults instantiates a new IntegrationsAzureEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntegrationsAzureEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntegrationsAzureEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntegrationsAzureEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IntegrationsAzureEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActivityId

`func (o *IntegrationsAzureEvent) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *IntegrationsAzureEvent) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *IntegrationsAzureEvent) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *IntegrationsAzureEvent) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *IntegrationsAzureEvent) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *IntegrationsAzureEvent) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *IntegrationsAzureEvent) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *IntegrationsAzureEvent) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetPublisherId

`func (o *IntegrationsAzureEvent) GetPublisherId() string`

GetPublisherId returns the PublisherId field if non-nil, zero value otherwise.

### GetPublisherIdOk

`func (o *IntegrationsAzureEvent) GetPublisherIdOk() (*string, bool)`

GetPublisherIdOk returns a tuple with the PublisherId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisherId

`func (o *IntegrationsAzureEvent) SetPublisherId(v string)`

SetPublisherId sets PublisherId field to given value.

### HasPublisherId

`func (o *IntegrationsAzureEvent) HasPublisherId() bool`

HasPublisherId returns a boolean if a field has been set.

### GetOfferId

`func (o *IntegrationsAzureEvent) GetOfferId() string`

GetOfferId returns the OfferId field if non-nil, zero value otherwise.

### GetOfferIdOk

`func (o *IntegrationsAzureEvent) GetOfferIdOk() (*string, bool)`

GetOfferIdOk returns a tuple with the OfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferId

`func (o *IntegrationsAzureEvent) SetOfferId(v string)`

SetOfferId sets OfferId field to given value.

### HasOfferId

`func (o *IntegrationsAzureEvent) HasOfferId() bool`

HasOfferId returns a boolean if a field has been set.

### GetPlanId

`func (o *IntegrationsAzureEvent) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *IntegrationsAzureEvent) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *IntegrationsAzureEvent) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *IntegrationsAzureEvent) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetQuantity

`func (o *IntegrationsAzureEvent) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *IntegrationsAzureEvent) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *IntegrationsAzureEvent) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *IntegrationsAzureEvent) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetTimeStamp

`func (o *IntegrationsAzureEvent) GetTimeStamp() string`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *IntegrationsAzureEvent) GetTimeStampOk() (*string, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *IntegrationsAzureEvent) SetTimeStamp(v string)`

SetTimeStamp sets TimeStamp field to given value.

### HasTimeStamp

`func (o *IntegrationsAzureEvent) HasTimeStamp() bool`

HasTimeStamp returns a boolean if a field has been set.

### GetAction

`func (o *IntegrationsAzureEvent) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *IntegrationsAzureEvent) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *IntegrationsAzureEvent) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *IntegrationsAzureEvent) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetStatus

`func (o *IntegrationsAzureEvent) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IntegrationsAzureEvent) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IntegrationsAzureEvent) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IntegrationsAzureEvent) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


