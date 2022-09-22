# ClusterEstate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | **string** |  | 
**CloudProviderId** | **string** |  | 
**CloudProviderName** | **string** |  | 
**ClusterRegions** | [**[]ClusterEstateClusterRegionsInner**](ClusterEstateClusterRegionsInner.md) |  | 

## Methods

### NewClusterEstate

`func NewClusterEstate(organizationId string, cloudProviderId string, cloudProviderName string, clusterRegions []ClusterEstateClusterRegionsInner, ) *ClusterEstate`

NewClusterEstate instantiates a new ClusterEstate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterEstateWithDefaults

`func NewClusterEstateWithDefaults() *ClusterEstate`

NewClusterEstateWithDefaults instantiates a new ClusterEstate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *ClusterEstate) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ClusterEstate) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ClusterEstate) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetCloudProviderId

`func (o *ClusterEstate) GetCloudProviderId() string`

GetCloudProviderId returns the CloudProviderId field if non-nil, zero value otherwise.

### GetCloudProviderIdOk

`func (o *ClusterEstate) GetCloudProviderIdOk() (*string, bool)`

GetCloudProviderIdOk returns a tuple with the CloudProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderId

`func (o *ClusterEstate) SetCloudProviderId(v string)`

SetCloudProviderId sets CloudProviderId field to given value.


### GetCloudProviderName

`func (o *ClusterEstate) GetCloudProviderName() string`

GetCloudProviderName returns the CloudProviderName field if non-nil, zero value otherwise.

### GetCloudProviderNameOk

`func (o *ClusterEstate) GetCloudProviderNameOk() (*string, bool)`

GetCloudProviderNameOk returns a tuple with the CloudProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderName

`func (o *ClusterEstate) SetCloudProviderName(v string)`

SetCloudProviderName sets CloudProviderName field to given value.


### GetClusterRegions

`func (o *ClusterEstate) GetClusterRegions() []ClusterEstateClusterRegionsInner`

GetClusterRegions returns the ClusterRegions field if non-nil, zero value otherwise.

### GetClusterRegionsOk

`func (o *ClusterEstate) GetClusterRegionsOk() (*[]ClusterEstateClusterRegionsInner, bool)`

GetClusterRegionsOk returns a tuple with the ClusterRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterRegions

`func (o *ClusterEstate) SetClusterRegions(v []ClusterEstateClusterRegionsInner)`

SetClusterRegions sets ClusterRegions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


