# TrialUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | **string** |  | 
**ExpireAt** | [**PointInTime**](PointInTime.md) |  | 
**CloudProviderUsages** | [**[]TrialUsageCloudProviderUsagesInner**](TrialUsageCloudProviderUsagesInner.md) |  | 

## Methods

### NewTrialUsage

`func NewTrialUsage(orgId string, expireAt PointInTime, cloudProviderUsages []TrialUsageCloudProviderUsagesInner, ) *TrialUsage`

NewTrialUsage instantiates a new TrialUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialUsageWithDefaults

`func NewTrialUsageWithDefaults() *TrialUsage`

NewTrialUsageWithDefaults instantiates a new TrialUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *TrialUsage) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *TrialUsage) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *TrialUsage) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetExpireAt

`func (o *TrialUsage) GetExpireAt() PointInTime`

GetExpireAt returns the ExpireAt field if non-nil, zero value otherwise.

### GetExpireAtOk

`func (o *TrialUsage) GetExpireAtOk() (*PointInTime, bool)`

GetExpireAtOk returns a tuple with the ExpireAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAt

`func (o *TrialUsage) SetExpireAt(v PointInTime)`

SetExpireAt sets ExpireAt field to given value.


### GetCloudProviderUsages

`func (o *TrialUsage) GetCloudProviderUsages() []TrialUsageCloudProviderUsagesInner`

GetCloudProviderUsages returns the CloudProviderUsages field if non-nil, zero value otherwise.

### GetCloudProviderUsagesOk

`func (o *TrialUsage) GetCloudProviderUsagesOk() (*[]TrialUsageCloudProviderUsagesInner, bool)`

GetCloudProviderUsagesOk returns a tuple with the CloudProviderUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderUsages

`func (o *TrialUsage) SetCloudProviderUsages(v []TrialUsageCloudProviderUsagesInner)`

SetCloudProviderUsages sets CloudProviderUsages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


