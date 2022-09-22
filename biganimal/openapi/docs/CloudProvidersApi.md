# \CloudProvidersApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateRegion**](CloudProvidersApi.md#ActivateRegion) | **Post** /cloud-providers/{cloudProviderId}/regions/{regionId}/activate | 
[**CreateCloudProvider**](CloudProvidersApi.md#CreateCloudProvider) | **Post** /cloud-providers/{cloudProviderId}/register | 
[**DeleteRegion**](CloudProvidersApi.md#DeleteRegion) | **Post** /cloud-providers/{cloudProviderId}/regions/{regionId}/delete | 
[**GetCloudProviderRegionInstanceTypes**](CloudProvidersApi.md#GetCloudProviderRegionInstanceTypes) | **Get** /cloud-providers/{cloudProviderId}/regions/{regionId}/instance-types | 
[**GetCloudProviderRegionVolumeTypes**](CloudProvidersApi.md#GetCloudProviderRegionVolumeTypes) | **Get** /cloud-providers/{cloudProviderId}/regions/{regionId}/volume-types | 
[**GetCloudProviderRegions**](CloudProvidersApi.md#GetCloudProviderRegions) | **Get** /cloud-providers/{cloudProviderId}/regions | 
[**GetCloudProviders**](CloudProvidersApi.md#GetCloudProviders) | **Get** /cloud-providers | 
[**GetRegions**](CloudProvidersApi.md#GetRegions) | **Get** /regions | 
[**GetRegisteredCloudProvider**](CloudProvidersApi.md#GetRegisteredCloudProvider) | **Get** /cloud-providers/{cloudProviderId}/register | 
[**GetVolumeProperties**](CloudProvidersApi.md#GetVolumeProperties) | **Get** /cloud-providers/{cloudProviderId}/regions/{regionId}/volume-types/{volumeTypeId}/volume-properties | 
[**SuspendRegion**](CloudProvidersApi.md#SuspendRegion) | **Post** /cloud-providers/{cloudProviderId}/regions/{regionId}/suspend | 



## ActivateRegion

> ActivateRegion202Response ActivateRegion(ctx, cloudProviderId, regionId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cloudProviderId := "azure" // string | 
    regionId := "us-east-1" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudProvidersApi.ActivateRegion(context.Background(), cloudProviderId, regionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProvidersApi.ActivateRegion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivateRegion`: ActivateRegion202Response
    fmt.Fprintf(os.Stdout, "Response from `CloudProvidersApi.ActivateRegion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProviderId** | **string** |  | 
**regionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ActivateRegion202Response**](ActivateRegion202Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCloudProvider

> CreateCloudProvider200Response CreateCloudProvider(ctx, cloudProviderId).CreateCloudProviderRequest(createCloudProviderRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cloudProviderId := "azure" // string | 
    createCloudProviderRequest := openapiclient.createCloudProvider_request{RegisterAws: openapiclient.NewRegisterAws("ExternalId_example", "RoleArn_example")} // CreateCloudProviderRequest | Register Request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudProvidersApi.CreateCloudProvider(context.Background(), cloudProviderId).CreateCloudProviderRequest(createCloudProviderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProvidersApi.CreateCloudProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCloudProvider`: CreateCloudProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `CloudProvidersApi.CreateCloudProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProviderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCloudProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createCloudProviderRequest** | [**CreateCloudProviderRequest**](CreateCloudProviderRequest.md) | Register Request | 

### Return type

[**CreateCloudProvider200Response**](CreateCloudProvider200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRegion

> ActivateRegion202Response DeleteRegion(ctx, cloudProviderId, regionId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cloudProviderId := "azure" // string | 
    regionId := "us-east-1" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudProvidersApi.DeleteRegion(context.Background(), cloudProviderId, regionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProvidersApi.DeleteRegion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRegion`: ActivateRegion202Response
    fmt.Fprintf(os.Stdout, "Response from `CloudProvidersApi.DeleteRegion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProviderId** | **string** |  | 
**regionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ActivateRegion202Response**](ActivateRegion202Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudProviderRegionInstanceTypes

> GetCloudProviderRegionInstanceTypes200Response GetCloudProviderRegionInstanceTypes(ctx, cloudProviderId, regionId).Q(q).Sort(sort).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cloudProviderId := "azure" // string | 
    regionId := "us-east-1" // string | 
    q := "John" // string | Fulltext search (optional)
    sort := "firstName" // string | Sort by property of the object. Use dot notation eg. \"state.city\" for sorting by nested property. Default sort is ASC. Prepend with \"-\" for DESC sort. For numeric sorting use \"size|numeric:true\". First part is property path, second part is optional parameter of the sort. Supported options are \"numeric:true|false\" for numeric sorting and \"sensitivity:base\" to compare base of the string only: a ≠ b, a = á, a = A. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudProvidersApi.GetCloudProviderRegionInstanceTypes(context.Background(), cloudProviderId, regionId).Q(q).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProvidersApi.GetCloudProviderRegionInstanceTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudProviderRegionInstanceTypes`: GetCloudProviderRegionInstanceTypes200Response
    fmt.Fprintf(os.Stdout, "Response from `CloudProvidersApi.GetCloudProviderRegionInstanceTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProviderId** | **string** |  | 
**regionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudProviderRegionInstanceTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **q** | **string** | Fulltext search | 
 **sort** | **string** | Sort by property of the object. Use dot notation eg. \&quot;state.city\&quot; for sorting by nested property. Default sort is ASC. Prepend with \&quot;-\&quot; for DESC sort. For numeric sorting use \&quot;size|numeric:true\&quot;. First part is property path, second part is optional parameter of the sort. Supported options are \&quot;numeric:true|false\&quot; for numeric sorting and \&quot;sensitivity:base\&quot; to compare base of the string only: a ≠ b, a &#x3D; á, a &#x3D; A. | 

### Return type

[**GetCloudProviderRegionInstanceTypes200Response**](GetCloudProviderRegionInstanceTypes200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudProviderRegionVolumeTypes

> GetCloudProviderRegionVolumeTypes200Response GetCloudProviderRegionVolumeTypes(ctx, cloudProviderId, regionId).Q(q).Sort(sort).InstanceFamilyNames(instanceFamilyNames).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cloudProviderId := "azure" // string | 
    regionId := "us-east-1" // string | 
    q := "John" // string | Fulltext search (optional)
    sort := "firstName" // string | Sort by property of the object. Use dot notation eg. \"state.city\" for sorting by nested property. Default sort is ASC. Prepend with \"-\" for DESC sort. For numeric sorting use \"size|numeric:true\". First part is property path, second part is optional parameter of the sort. Supported options are \"numeric:true|false\" for numeric sorting and \"sensitivity:base\" to compare base of the string only: a ≠ b, a = á, a = A. (optional)
    instanceFamilyNames := []string{"Inner_example"} // []string | Array of instanceType family names (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudProvidersApi.GetCloudProviderRegionVolumeTypes(context.Background(), cloudProviderId, regionId).Q(q).Sort(sort).InstanceFamilyNames(instanceFamilyNames).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProvidersApi.GetCloudProviderRegionVolumeTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudProviderRegionVolumeTypes`: GetCloudProviderRegionVolumeTypes200Response
    fmt.Fprintf(os.Stdout, "Response from `CloudProvidersApi.GetCloudProviderRegionVolumeTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProviderId** | **string** |  | 
**regionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudProviderRegionVolumeTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **q** | **string** | Fulltext search | 
 **sort** | **string** | Sort by property of the object. Use dot notation eg. \&quot;state.city\&quot; for sorting by nested property. Default sort is ASC. Prepend with \&quot;-\&quot; for DESC sort. For numeric sorting use \&quot;size|numeric:true\&quot;. First part is property path, second part is optional parameter of the sort. Supported options are \&quot;numeric:true|false\&quot; for numeric sorting and \&quot;sensitivity:base\&quot; to compare base of the string only: a ≠ b, a &#x3D; á, a &#x3D; A. | 
 **instanceFamilyNames** | **[]string** | Array of instanceType family names | 

### Return type

[**GetCloudProviderRegionVolumeTypes200Response**](GetCloudProviderRegionVolumeTypes200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudProviderRegions

> GetCloudProviderRegions200Response GetCloudProviderRegions(ctx, cloudProviderId).Q(q).Sort(sort).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cloudProviderId := "azure" // string | 
    q := "John" // string | Fulltext search (optional)
    sort := "firstName" // string | Sort by property of the object. Use dot notation eg. \"state.city\" for sorting by nested property. Default sort is ASC. Prepend with \"-\" for DESC sort. For numeric sorting use \"size|numeric:true\". First part is property path, second part is optional parameter of the sort. Supported options are \"numeric:true|false\" for numeric sorting and \"sensitivity:base\" to compare base of the string only: a ≠ b, a = á, a = A. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudProvidersApi.GetCloudProviderRegions(context.Background(), cloudProviderId).Q(q).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProvidersApi.GetCloudProviderRegions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudProviderRegions`: GetCloudProviderRegions200Response
    fmt.Fprintf(os.Stdout, "Response from `CloudProvidersApi.GetCloudProviderRegions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProviderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudProviderRegionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **string** | Fulltext search | 
 **sort** | **string** | Sort by property of the object. Use dot notation eg. \&quot;state.city\&quot; for sorting by nested property. Default sort is ASC. Prepend with \&quot;-\&quot; for DESC sort. For numeric sorting use \&quot;size|numeric:true\&quot;. First part is property path, second part is optional parameter of the sort. Supported options are \&quot;numeric:true|false\&quot; for numeric sorting and \&quot;sensitivity:base\&quot; to compare base of the string only: a ≠ b, a &#x3D; á, a &#x3D; A. | 

### Return type

[**GetCloudProviderRegions200Response**](GetCloudProviderRegions200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudProviders

> GetCloudProviders200Response GetCloudProviders(ctx).Q(q).Sort(sort).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    q := "John" // string | Fulltext search (optional)
    sort := "firstName" // string | Sort by property of the object. Use dot notation eg. \"state.city\" for sorting by nested property. Default sort is ASC. Prepend with \"-\" for DESC sort. For numeric sorting use \"size|numeric:true\". First part is property path, second part is optional parameter of the sort. Supported options are \"numeric:true|false\" for numeric sorting and \"sensitivity:base\" to compare base of the string only: a ≠ b, a = á, a = A. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudProvidersApi.GetCloudProviders(context.Background()).Q(q).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProvidersApi.GetCloudProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudProviders`: GetCloudProviders200Response
    fmt.Fprintf(os.Stdout, "Response from `CloudProvidersApi.GetCloudProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Fulltext search | 
 **sort** | **string** | Sort by property of the object. Use dot notation eg. \&quot;state.city\&quot; for sorting by nested property. Default sort is ASC. Prepend with \&quot;-\&quot; for DESC sort. For numeric sorting use \&quot;size|numeric:true\&quot;. First part is property path, second part is optional parameter of the sort. Supported options are \&quot;numeric:true|false\&quot; for numeric sorting and \&quot;sensitivity:base\&quot; to compare base of the string only: a ≠ b, a &#x3D; á, a &#x3D; A. | 

### Return type

[**GetCloudProviders200Response**](GetCloudProviders200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegions

> GetRegions200Response GetRegions(ctx).CloudProviderIds(cloudProviderIds).Continents(continents).Q(q).Sort(sort).Status(status).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cloudProviderIds := []string{"Inner_example"} // []string | Array of cloud provider IDs (optional)
    continents := []string{"Inner_example"} // []string | Array of continents (optional)
    q := "John" // string | Fulltext search (optional)
    sort := "firstName" // string | Sort by property of the object. Use dot notation eg. \"state.city\" for sorting by nested property. Default sort is ASC. Prepend with \"-\" for DESC sort. For numeric sorting use \"size|numeric:true\". First part is property path, second part is optional parameter of the sort. Supported options are \"numeric:true|false\" for numeric sorting and \"sensitivity:base\" to compare base of the string only: a ≠ b, a = á, a = A. (optional)
    status := []string{"Inner_example"} // []string | Array of status (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudProvidersApi.GetRegions(context.Background()).CloudProviderIds(cloudProviderIds).Continents(continents).Q(q).Sort(sort).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProvidersApi.GetRegions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegions`: GetRegions200Response
    fmt.Fprintf(os.Stdout, "Response from `CloudProvidersApi.GetRegions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRegionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudProviderIds** | **[]string** | Array of cloud provider IDs | 
 **continents** | **[]string** | Array of continents | 
 **q** | **string** | Fulltext search | 
 **sort** | **string** | Sort by property of the object. Use dot notation eg. \&quot;state.city\&quot; for sorting by nested property. Default sort is ASC. Prepend with \&quot;-\&quot; for DESC sort. For numeric sorting use \&quot;size|numeric:true\&quot;. First part is property path, second part is optional parameter of the sort. Supported options are \&quot;numeric:true|false\&quot; for numeric sorting and \&quot;sensitivity:base\&quot; to compare base of the string only: a ≠ b, a &#x3D; á, a &#x3D; A. | 
 **status** | **[]string** | Array of status | 

### Return type

[**GetRegions200Response**](GetRegions200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegisteredCloudProvider

> GetRegisteredCloudProvider200Response GetRegisteredCloudProvider(ctx, cloudProviderId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cloudProviderId := "azure" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudProvidersApi.GetRegisteredCloudProvider(context.Background(), cloudProviderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProvidersApi.GetRegisteredCloudProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegisteredCloudProvider`: GetRegisteredCloudProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `CloudProvidersApi.GetRegisteredCloudProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProviderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegisteredCloudProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRegisteredCloudProvider200Response**](GetRegisteredCloudProvider200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVolumeProperties

> GetVolumeProperties200Response GetVolumeProperties(ctx, cloudProviderId, regionId, volumeTypeId).Q(q).Sort(sort).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cloudProviderId := "azure" // string | 
    regionId := "us-east-1" // string | 
    volumeTypeId := "azurepremiumstorage" // string | 
    q := "John" // string | Fulltext search (optional)
    sort := "firstName" // string | Sort by property of the object. Use dot notation eg. \"state.city\" for sorting by nested property. Default sort is ASC. Prepend with \"-\" for DESC sort. For numeric sorting use \"size|numeric:true\". First part is property path, second part is optional parameter of the sort. Supported options are \"numeric:true|false\" for numeric sorting and \"sensitivity:base\" to compare base of the string only: a ≠ b, a = á, a = A. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudProvidersApi.GetVolumeProperties(context.Background(), cloudProviderId, regionId, volumeTypeId).Q(q).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProvidersApi.GetVolumeProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVolumeProperties`: GetVolumeProperties200Response
    fmt.Fprintf(os.Stdout, "Response from `CloudProvidersApi.GetVolumeProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProviderId** | **string** |  | 
**regionId** | **string** |  | 
**volumeTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVolumePropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **q** | **string** | Fulltext search | 
 **sort** | **string** | Sort by property of the object. Use dot notation eg. \&quot;state.city\&quot; for sorting by nested property. Default sort is ASC. Prepend with \&quot;-\&quot; for DESC sort. For numeric sorting use \&quot;size|numeric:true\&quot;. First part is property path, second part is optional parameter of the sort. Supported options are \&quot;numeric:true|false\&quot; for numeric sorting and \&quot;sensitivity:base\&quot; to compare base of the string only: a ≠ b, a &#x3D; á, a &#x3D; A. | 

### Return type

[**GetVolumeProperties200Response**](GetVolumeProperties200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuspendRegion

> ActivateRegion202Response SuspendRegion(ctx, cloudProviderId, regionId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cloudProviderId := "azure" // string | 
    regionId := "us-east-1" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudProvidersApi.SuspendRegion(context.Background(), cloudProviderId, regionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProvidersApi.SuspendRegion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SuspendRegion`: ActivateRegion202Response
    fmt.Fprintf(os.Stdout, "Response from `CloudProvidersApi.SuspendRegion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProviderId** | **string** |  | 
**regionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSuspendRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ActivateRegion202Response**](ActivateRegion202Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

