# \OrganizationsApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountClustersOverview**](OrganizationsApi.md#GetAccountClustersOverview) | **Get** /account/clusters | 
[**GetAccountNews**](OrganizationsApi.md#GetAccountNews) | **Get** /account/news | 
[**GetAccountNotifications**](OrganizationsApi.md#GetAccountNotifications) | **Get** /account/notifications | 
[**GetOrganizationSettings**](OrganizationsApi.md#GetOrganizationSettings) | **Get** /account/settings | 
[**UpdateOrganizationSettings**](OrganizationsApi.md#UpdateOrganizationSettings) | **Put** /account/settings | 



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
    resp, r, err := apiClient.OrganizationsApi.GetAccountClustersOverview(context.Background()).Q(q).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetAccountClustersOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountClustersOverview`: GetAccountClustersOverview200Response
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetAccountClustersOverview`: %v\n", resp)
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


## GetAccountNews

> GetAccountNews200Response GetAccountNews(ctx).Q(q).Sort(sort).Execute()





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
    resp, r, err := apiClient.OrganizationsApi.GetAccountNews(context.Background()).Q(q).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetAccountNews``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountNews`: GetAccountNews200Response
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetAccountNews`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountNewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Fulltext search | 
 **sort** | **string** | Sort by property of the object. Use dot notation eg. \&quot;state.city\&quot; for sorting by nested property. Default sort is ASC. Prepend with \&quot;-\&quot; for DESC sort. For numeric sorting use \&quot;size|numeric:true\&quot;. First part is property path, second part is optional parameter of the sort. Supported options are \&quot;numeric:true|false\&quot; for numeric sorting and \&quot;sensitivity:base\&quot; to compare base of the string only: a ≠ b, a &#x3D; á, a &#x3D; A. | 

### Return type

[**GetAccountNews200Response**](GetAccountNews200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountNotifications

> GetAccountNotifications200Response GetAccountNotifications(ctx).Q(q).Sort(sort).Execute()





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
    resp, r, err := apiClient.OrganizationsApi.GetAccountNotifications(context.Background()).Q(q).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetAccountNotifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountNotifications`: GetAccountNotifications200Response
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetAccountNotifications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Fulltext search | 
 **sort** | **string** | Sort by property of the object. Use dot notation eg. \&quot;state.city\&quot; for sorting by nested property. Default sort is ASC. Prepend with \&quot;-\&quot; for DESC sort. For numeric sorting use \&quot;size|numeric:true\&quot;. First part is property path, second part is optional parameter of the sort. Supported options are \&quot;numeric:true|false\&quot; for numeric sorting and \&quot;sensitivity:base\&quot; to compare base of the string only: a ≠ b, a &#x3D; á, a &#x3D; A. | 

### Return type

[**GetAccountNotifications200Response**](GetAccountNotifications200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSettings

> GetOrganizationSettings200Response GetOrganizationSettings(ctx).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSettings`: GetOrganizationSettings200Response
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSettingsRequest struct via the builder pattern


### Return type

[**GetOrganizationSettings200Response**](GetOrganizationSettings200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationSettings

> GetOrganizationSettings200Response UpdateOrganizationSettings(ctx).RequestBody(requestBody).Execute()





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
    requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.UpdateOrganizationSettings(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UpdateOrganizationSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationSettings`: GetOrganizationSettings200Response
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.UpdateOrganizationSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]interface{}** |  | 

### Return type

[**GetOrganizationSettings200Response**](GetOrganizationSettings200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

