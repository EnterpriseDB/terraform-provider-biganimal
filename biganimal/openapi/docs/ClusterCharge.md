# ClusterCharge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** |  | 
**VcpuHours** | **float32** |  | 

## Methods

### NewClusterCharge

`func NewClusterCharge(clusterId string, vcpuHours float32, ) *ClusterCharge`

NewClusterCharge instantiates a new ClusterCharge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterChargeWithDefaults

`func NewClusterChargeWithDefaults() *ClusterCharge`

NewClusterChargeWithDefaults instantiates a new ClusterCharge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *ClusterCharge) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ClusterCharge) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ClusterCharge) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetVcpuHours

`func (o *ClusterCharge) GetVcpuHours() float32`

GetVcpuHours returns the VcpuHours field if non-nil, zero value otherwise.

### GetVcpuHoursOk

`func (o *ClusterCharge) GetVcpuHoursOk() (*float32, bool)`

GetVcpuHoursOk returns a tuple with the VcpuHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpuHours

`func (o *ClusterCharge) SetVcpuHours(v float32)`

SetVcpuHours sets VcpuHours field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


