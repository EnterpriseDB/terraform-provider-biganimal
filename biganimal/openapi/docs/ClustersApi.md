# \ClustersApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCluster**](ClustersApi.md#CreateCluster) | **Post** /clusters | 
[**DeleteCluster**](ClustersApi.md#DeleteCluster) | **Delete** /clusters/{clusterId} | 
[**GetAccountClustersOverview**](ClustersApi.md#GetAccountClustersOverview) | **Get** /account/clusters | 
[**GetCluster**](ClustersApi.md#GetCluster) | **Get** /clusters/{clusterId} | 
[**GetClusterConnection**](ClustersApi.md#GetClusterConnection) | **Get** /clusters/{clusterId}/connection | 
[**GetClusterPgConfig**](ClustersApi.md#GetClusterPgConfig) | **Get** /clusters/{clusterId}/pg-config-parameters | 
[**GetClusterPhases**](ClustersApi.md#GetClusterPhases) | **Get** /clusters/{clusterId}/phases | 
[**GetClusters**](ClustersApi.md#GetClusters) | **Get** /clusters | 
[**GetDeletedCluster**](ClustersApi.md#GetDeletedCluster) | **Get** /deleted-clusters/{clusterId} | 
[**GetDeletedClusters**](ClustersApi.md#GetDeletedClusters) | **Get** /deleted-clusters | 
[**PatchCluster**](ClustersApi.md#PatchCluster) | **Patch** /clusters/{clusterId} | 
[**RestoreCluster**](ClustersApi.md#RestoreCluster) | **Post** /clusters/{clusterId}/restore | 
[**RestoreDeletedCluster**](ClustersApi.md#RestoreDeletedCluster) | **Post** /deleted-clusters/{clusterId}/restore | 
[**UpdateCluster**](ClustersApi.md#UpdateCluster) | **Put** /clusters/{clusterId} | 



## CreateCluster

> CreateCluster202Response CreateCluster(ctx).CreateCluster(createCluster).Execute()





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
    createCluster := *openapiclient.NewCreateCluster("ClusterName_example", "Password_example", *openapiclient.NewCreateClusterPgType("PgTypeId_example"), *openapiclient.NewCreateClusterPgVersion("PgVersionId_example"), *openapiclient.NewCreateClusterProvider("CloudProviderId_example"), *openapiclient.NewActivateRegion202ResponseData("RegionId_example"), *openapiclient.NewCreateClusterInstanceType("InstanceTypeId_example"), *openapiclient.NewCreateClusterStorage("VolumePropertiesId_example", "VolumeTypeId_example"), false, []openapiclient.AllowedIpRange{*openapiclient.NewAllowedIpRange("CidrBlock_example", "Description_example")}, []openapiclient.ClusterDetailPgConfigInner{*openapiclient.NewClusterDetailPgConfigInner()}) // CreateCluster | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.CreateCluster(context.Background()).CreateCluster(createCluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.CreateCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCluster`: CreateCluster202Response
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.CreateCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCluster** | [**CreateCluster**](CreateCluster.md) |  | 

### Return type

[**CreateCluster202Response**](CreateCluster202Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCluster

> DeleteCluster(ctx, clusterId).Execute()





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
    clusterId := "p-c5jhbe65n1d8iosambm0" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.DeleteCluster(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.DeleteCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountClustersOverview

> GetAccountClustersOverview200Response GetAccountClustersOverview(ctx).Q(q).Sort(sort).Execute()





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
    resp, r, err := apiClient.ClustersApi.GetAccountClustersOverview(context.Background()).Q(q).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetAccountClustersOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountClustersOverview`: GetAccountClustersOverview200Response
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetAccountClustersOverview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountClustersOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Fulltext search | 
 **sort** | **string** | Sort by property of the object. Use dot notation eg. \&quot;state.city\&quot; for sorting by nested property. Default sort is ASC. Prepend with \&quot;-\&quot; for DESC sort. For numeric sorting use \&quot;size|numeric:true\&quot;. First part is property path, second part is optional parameter of the sort. Supported options are \&quot;numeric:true|false\&quot; for numeric sorting and \&quot;sensitivity:base\&quot; to compare base of the string only: a ≠ b, a &#x3D; á, a &#x3D; A. | 

### Return type

[**GetAccountClustersOverview200Response**](GetAccountClustersOverview200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCluster

> GetCluster200Response GetCluster(ctx, clusterId).Execute()





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
    clusterId := "p-c5jhbe65n1d8iosambm0" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.GetCluster(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCluster`: GetCluster200Response
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCluster200Response**](GetCluster200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterConnection

> GetClusterConnection200Response GetClusterConnection(ctx, clusterId).Execute()





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
    clusterId := "my-cluster" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.GetClusterConnection(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetClusterConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterConnection`: GetClusterConnection200Response
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetClusterConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetClusterConnection200Response**](GetClusterConnection200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterPgConfig

> GetClusterPgConfig200Response GetClusterPgConfig(ctx, clusterId).Execute()





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
    clusterId := "clusterId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.GetClusterPgConfig(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetClusterPgConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterPgConfig`: GetClusterPgConfig200Response
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetClusterPgConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterPgConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetClusterPgConfig200Response**](GetClusterPgConfig200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterPhases

> GetClusterPhases200Response GetClusterPhases(ctx, clusterId).Execute()





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
    clusterId := "my-cluster" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.GetClusterPhases(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetClusterPhases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterPhases`: GetClusterPhases200Response
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetClusterPhases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterPhasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetClusterPhases200Response**](GetClusterPhases200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusters

> GetClusters200Response GetClusters(ctx).ClusterArchitectureIds(clusterArchitectureIds).Name(name).RegionIds(regionIds).PgTypeIds(pgTypeIds).PgVersionIds(pgVersionIds).Exact(exact).Sort(sort).Execute()





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
    clusterArchitectureIds := []string{"Inner_example"} // []string | Filter clusters by cluster architecture (optional)
    name := "my-cluster-1" // string | Cluster name to filter to (optional)
    regionIds := []string{"Inner_example"} // []string | Filter clusters by region (optional)
    pgTypeIds := []string{"Inner_example"} // []string | Filter clusters by type (optional)
    pgVersionIds := []string{"Inner_example"} // []string | Filter clusters by version (optional)
    exact := true // bool | Filter for exact match on non-list parameters (optional)
    sort := "+clusterName" // string | sort clusters by property (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.GetClusters(context.Background()).ClusterArchitectureIds(clusterArchitectureIds).Name(name).RegionIds(regionIds).PgTypeIds(pgTypeIds).PgVersionIds(pgVersionIds).Exact(exact).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusters`: GetClusters200Response
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterArchitectureIds** | **[]string** | Filter clusters by cluster architecture | 
 **name** | **string** | Cluster name to filter to | 
 **regionIds** | **[]string** | Filter clusters by region | 
 **pgTypeIds** | **[]string** | Filter clusters by type | 
 **pgVersionIds** | **[]string** | Filter clusters by version | 
 **exact** | **bool** | Filter for exact match on non-list parameters | 
 **sort** | **string** | sort clusters by property | 

### Return type

[**GetClusters200Response**](GetClusters200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeletedCluster

> GetCluster200Response GetDeletedCluster(ctx, clusterId).Execute()





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
    clusterId := "p-c5jhbe65n1d8iosambm0" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.GetDeletedCluster(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetDeletedCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeletedCluster`: GetCluster200Response
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetDeletedCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeletedClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCluster200Response**](GetCluster200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeletedClusters

> GetClusters200Response GetDeletedClusters(ctx).ClusterArchitectureIds(clusterArchitectureIds).Name(name).RegionIds(regionIds).PgTypeIds(pgTypeIds).PgVersionIds(pgVersionIds).Exact(exact).Sort(sort).Execute()





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
    clusterArchitectureIds := []string{"Inner_example"} // []string | Filter clusters by cluster architecture (optional)
    name := "my-cluster-1" // string | Cluster name to filter to (optional)
    regionIds := []string{"Inner_example"} // []string | Filter clusters by region (optional)
    pgTypeIds := []string{"Inner_example"} // []string | Filter clusters by type (optional)
    pgVersionIds := []string{"Inner_example"} // []string | Filter clusters by version (optional)
    exact := true // bool | Filter for exact match on non-list parameters (optional)
    sort := "+clusterName" // string | sort clusters by property (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.GetDeletedClusters(context.Background()).ClusterArchitectureIds(clusterArchitectureIds).Name(name).RegionIds(regionIds).PgTypeIds(pgTypeIds).PgVersionIds(pgVersionIds).Exact(exact).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetDeletedClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeletedClusters`: GetClusters200Response
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetDeletedClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeletedClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterArchitectureIds** | **[]string** | Filter clusters by cluster architecture | 
 **name** | **string** | Cluster name to filter to | 
 **regionIds** | **[]string** | Filter clusters by region | 
 **pgTypeIds** | **[]string** | Filter clusters by type | 
 **pgVersionIds** | **[]string** | Filter clusters by version | 
 **exact** | **bool** | Filter for exact match on non-list parameters | 
 **sort** | **string** | sort clusters by property | 

### Return type

[**GetClusters200Response**](GetClusters200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCluster

> CreateCluster202Response PatchCluster(ctx, clusterId).PatchCluster(patchCluster).Execute()





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
    clusterId := "p-c5jhbe65n1d8iosambm0" // string | 
    patchCluster := *openapiclient.NewPatchCluster() // PatchCluster | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.PatchCluster(context.Background(), clusterId).PatchCluster(patchCluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.PatchCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchCluster`: CreateCluster202Response
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.PatchCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchCluster** | [**PatchCluster**](PatchCluster.md) |  | 

### Return type

[**CreateCluster202Response**](CreateCluster202Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreCluster

> CreateCluster202Response RestoreCluster(ctx, clusterId).RestoreCluster(restoreCluster).Execute()





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
    clusterId := "p-c5jhbe65n1d8iosambm0" // string | 
    restoreCluster := *openapiclient.NewRestoreCluster("ClusterName_example", "Password_example", *openapiclient.NewActivateRegion202ResponseData("RegionId_example"), *openapiclient.NewCreateClusterInstanceType("InstanceTypeId_example"), *openapiclient.NewCreateClusterStorage("VolumePropertiesId_example", "VolumeTypeId_example"), []openapiclient.AllowedIpRange{*openapiclient.NewAllowedIpRange("CidrBlock_example", "Description_example")}, []openapiclient.ClusterDetailPgConfigInner{*openapiclient.NewClusterDetailPgConfigInner()}) // RestoreCluster | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.RestoreCluster(context.Background(), clusterId).RestoreCluster(restoreCluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.RestoreCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestoreCluster`: CreateCluster202Response
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.RestoreCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **restoreCluster** | [**RestoreCluster**](RestoreCluster.md) |  | 

### Return type

[**CreateCluster202Response**](CreateCluster202Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreDeletedCluster

> CreateCluster202Response RestoreDeletedCluster(ctx, clusterId).RestoreCluster(restoreCluster).Execute()





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
    clusterId := "p-c5jhbe65n1d8iosambm0" // string | 
    restoreCluster := *openapiclient.NewRestoreCluster("ClusterName_example", "Password_example", *openapiclient.NewActivateRegion202ResponseData("RegionId_example"), *openapiclient.NewCreateClusterInstanceType("InstanceTypeId_example"), *openapiclient.NewCreateClusterStorage("VolumePropertiesId_example", "VolumeTypeId_example"), []openapiclient.AllowedIpRange{*openapiclient.NewAllowedIpRange("CidrBlock_example", "Description_example")}, []openapiclient.ClusterDetailPgConfigInner{*openapiclient.NewClusterDetailPgConfigInner()}) // RestoreCluster | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.RestoreDeletedCluster(context.Background(), clusterId).RestoreCluster(restoreCluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.RestoreDeletedCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestoreDeletedCluster`: CreateCluster202Response
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.RestoreDeletedCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreDeletedClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **restoreCluster** | [**RestoreCluster**](RestoreCluster.md) |  | 

### Return type

[**CreateCluster202Response**](CreateCluster202Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCluster

> CreateCluster202Response UpdateCluster(ctx, clusterId).UpdateCluster(updateCluster).Execute()





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
    clusterId := "p-c5jhbe65n1d8iosambm0" // string | 
    updateCluster := *openapiclient.NewUpdateCluster("ClusterName_example", *openapiclient.NewCreateClusterInstanceType("InstanceTypeId_example"), *openapiclient.NewCreateClusterStorage("VolumePropertiesId_example", "VolumeTypeId_example"), false, []openapiclient.AllowedIpRange{*openapiclient.NewAllowedIpRange("CidrBlock_example", "Description_example")}, []openapiclient.ClusterDetailPgConfigInner{*openapiclient.NewClusterDetailPgConfigInner()}) // UpdateCluster | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.UpdateCluster(context.Background(), clusterId).UpdateCluster(updateCluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.UpdateCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCluster`: CreateCluster202Response
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.UpdateCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCluster** | [**UpdateCluster**](UpdateCluster.md) |  | 

### Return type

[**CreateCluster202Response**](CreateCluster202Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

