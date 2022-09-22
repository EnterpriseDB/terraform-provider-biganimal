# Billing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterCharges** | [**[]ClusterCharge**](ClusterCharge.md) |  | 
**TotalVcpuHours** | **float32** |  | 

## Methods

### NewBilling

`func NewBilling(clusterCharges []ClusterCharge, totalVcpuHours float32, ) *Billing`

NewBilling instantiates a new Billing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingWithDefaults

`func NewBillingWithDefaults() *Billing`

NewBillingWithDefaults instantiates a new Billing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterCharges

`func (o *Billing) GetClusterCharges() []ClusterCharge`

GetClusterCharges returns the ClusterCharges field if non-nil, zero value otherwise.

### GetClusterChargesOk

`func (o *Billing) GetClusterChargesOk() (*[]ClusterCharge, bool)`

GetClusterChargesOk returns a tuple with the ClusterCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterCharges

`func (o *Billing) SetClusterCharges(v []ClusterCharge)`

SetClusterCharges sets ClusterCharges field to given value.


### GetTotalVcpuHours

`func (o *Billing) GetTotalVcpuHours() float32`

GetTotalVcpuHours returns the TotalVcpuHours field if non-nil, zero value otherwise.

### GetTotalVcpuHoursOk

`func (o *Billing) GetTotalVcpuHoursOk() (*float32, bool)`

GetTotalVcpuHoursOk returns a tuple with the TotalVcpuHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVcpuHours

`func (o *Billing) SetTotalVcpuHours(v float32)`

SetTotalVcpuHours sets TotalVcpuHours field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


