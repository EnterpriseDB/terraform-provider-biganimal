# \BillingApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBilling**](BillingApi.md#GetBilling) | **Get** /billing | 
[**GetUsage**](BillingApi.md#GetUsage) | **Get** /usage | 
[**GetUsageReportFile**](BillingApi.md#GetUsageReportFile) | **Get** /usage-csv | 



## GetBilling

> GetBilling200Response GetBilling(ctx).StartAt(startAt).EndAt(endAt).Execute()





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
    startAt := "2021-01-15T12:30:40Z" // string | 
    endAt := "2021-02-15T12:30:40Z" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingApi.GetBilling(context.Background()).StartAt(startAt).EndAt(endAt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.GetBilling``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBilling`: GetBilling200Response
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.GetBilling`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startAt** | **string** |  | 
 **endAt** | **string** |  | 

### Return type

[**GetBilling200Response**](GetBilling200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsage

> GetUsage200Response GetUsage(ctx).StartAt(startAt).EndAt(endAt).CloudProviderIds(cloudProviderIds).PgTypeIds(pgTypeIds).ClusterArchitecturesIds(clusterArchitecturesIds).Q(q).Sort(sort).Execute()





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
    startAt := "2021-01-15T12:30:40Z" // string | 
    endAt := "2021-02-15T12:30:40Z" // string | 
    cloudProviderIds := []string{"Inner_example"} // []string | Array of cloud provider IDs (optional)
    pgTypeIds := []string{"Inner_example"} // []string | Array of pgType IDs (optional)
    clusterArchitecturesIds := []string{"Inner_example"} // []string | Array of cloud cluster architectures (optional)
    q := "John" // string | Fulltext search (optional)
    sort := "firstName" // string | Sort by property of the object. Use dot notation eg. \"state.city\" for sorting by nested property. Default sort is ASC. Prepend with \"-\" for DESC sort. For numeric sorting use \"size|numeric:true\". First part is property path, second part is optional parameter of the sort. Supported options are \"numeric:true|false\" for numeric sorting and \"sensitivity:base\" to compare base of the string only: a ≠ b, a = á, a = A. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingApi.GetUsage(context.Background()).StartAt(startAt).EndAt(endAt).CloudProviderIds(cloudProviderIds).PgTypeIds(pgTypeIds).ClusterArchitecturesIds(clusterArchitecturesIds).Q(q).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.GetUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsage`: GetUsage200Response
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.GetUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startAt** | **string** |  | 
 **endAt** | **string** |  | 
 **cloudProviderIds** | **[]string** | Array of cloud provider IDs | 
 **pgTypeIds** | **[]string** | Array of pgType IDs | 
 **clusterArchitecturesIds** | **[]string** | Array of cloud cluster architectures | 
 **q** | **string** | Fulltext search | 
 **sort** | **string** | Sort by property of the object. Use dot notation eg. \&quot;state.city\&quot; for sorting by nested property. Default sort is ASC. Prepend with \&quot;-\&quot; for DESC sort. For numeric sorting use \&quot;size|numeric:true\&quot;. First part is property path, second part is optional parameter of the sort. Supported options are \&quot;numeric:true|false\&quot; for numeric sorting and \&quot;sensitivity:base\&quot; to compare base of the string only: a ≠ b, a &#x3D; á, a &#x3D; A. | 

### Return type

[**GetUsage200Response**](GetUsage200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageReportFile

> string GetUsageReportFile(ctx).StartAt(startAt).EndAt(endAt).CloudProviderIds(cloudProviderIds).ClusterArchitecturesIds(clusterArchitecturesIds).PgTypeIds(pgTypeIds).Q(q).Sort(sort).ClusterNameTitle(clusterNameTitle).PgTypeTitle(pgTypeTitle).CloudProviderTitle(cloudProviderTitle).ClusterArchitectureTitle(clusterArchitectureTitle).VcpuHoursTitle(vcpuHoursTitle).DeletedClusterPrefix(deletedClusterPrefix).Execute()





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
    startAt := "2021-01-15T12:30:40Z" // string | 
    endAt := "2021-02-15T12:30:40Z" // string | 
    cloudProviderIds := []string{"Inner_example"} // []string | Array of cloud provider IDs (optional)
    clusterArchitecturesIds := []string{"Inner_example"} // []string | Array of cloud cluster architectures (optional)
    pgTypeIds := []string{"Inner_example"} // []string | Array of pgType IDs (optional)
    q := "John" // string | Fulltext search (optional)
    sort := "firstName" // string | Sort by property of the object. Use dot notation eg. \"state.city\" for sorting by nested property. Default sort is ASC. Prepend with \"-\" for DESC sort. For numeric sorting use \"size|numeric:true\". First part is property path, second part is optional parameter of the sort. Supported options are \"numeric:true|false\" for numeric sorting and \"sensitivity:base\" to compare base of the string only: a ≠ b, a = á, a = A. (optional)
    clusterNameTitle := "Cluster Name" // string | The cluster name columm title. The default value is Cluster Name. (optional)
    pgTypeTitle := "Database Type" // string | The database type column title. The default value is Database Type. (optional)
    cloudProviderTitle := "Cloud Provider" // string | The cloud provider column title. The default value is Cloud Provider. (optional)
    clusterArchitectureTitle := "Single" // string | The cluster architecture title. The default value is Cluster Architecture. (optional)
    vcpuHoursTitle := "vCPU Hours" // string | The usage of vCPU hour column title. The default value is vCPU Hours. (optional)
    deletedClusterPrefix := "Deleted cluster" // string | If the cluster is deleted, in Cluster Name column, it will show the deleted prefix with cluster id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingApi.GetUsageReportFile(context.Background()).StartAt(startAt).EndAt(endAt).CloudProviderIds(cloudProviderIds).ClusterArchitecturesIds(clusterArchitecturesIds).PgTypeIds(pgTypeIds).Q(q).Sort(sort).ClusterNameTitle(clusterNameTitle).PgTypeTitle(pgTypeTitle).CloudProviderTitle(cloudProviderTitle).ClusterArchitectureTitle(clusterArchitectureTitle).VcpuHoursTitle(vcpuHoursTitle).DeletedClusterPrefix(deletedClusterPrefix).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.GetUsageReportFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageReportFile`: string
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.GetUsageReportFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageReportFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startAt** | **string** |  | 
 **endAt** | **string** |  | 
 **cloudProviderIds** | **[]string** | Array of cloud provider IDs | 
 **clusterArchitecturesIds** | **[]string** | Array of cloud cluster architectures | 
 **pgTypeIds** | **[]string** | Array of pgType IDs | 
 **q** | **string** | Fulltext search | 
 **sort** | **string** | Sort by property of the object. Use dot notation eg. \&quot;state.city\&quot; for sorting by nested property. Default sort is ASC. Prepend with \&quot;-\&quot; for DESC sort. For numeric sorting use \&quot;size|numeric:true\&quot;. First part is property path, second part is optional parameter of the sort. Supported options are \&quot;numeric:true|false\&quot; for numeric sorting and \&quot;sensitivity:base\&quot; to compare base of the string only: a ≠ b, a &#x3D; á, a &#x3D; A. | 
 **clusterNameTitle** | **string** | The cluster name columm title. The default value is Cluster Name. | 
 **pgTypeTitle** | **string** | The database type column title. The default value is Database Type. | 
 **cloudProviderTitle** | **string** | The cloud provider column title. The default value is Cloud Provider. | 
 **clusterArchitectureTitle** | **string** | The cluster architecture title. The default value is Cluster Architecture. | 
 **vcpuHoursTitle** | **string** | The usage of vCPU hour column title. The default value is vCPU Hours. | 
 **deletedClusterPrefix** | **string** | If the cluster is deleted, in Cluster Name column, it will show the deleted prefix with cluster id. | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

