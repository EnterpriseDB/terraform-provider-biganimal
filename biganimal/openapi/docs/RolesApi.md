# \RolesApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignRolePermissions**](RolesApi.md#AssignRolePermissions) | **Put** /roles/{roleId}/permissions | 
[**DeleteRole**](RolesApi.md#DeleteRole) | **Delete** /roles/{roleId} | 
[**GetRole**](RolesApi.md#GetRole) | **Get** /roles/{roleId} | 
[**GetRoles**](RolesApi.md#GetRoles) | **Get** /roles | 
[**ListRolePermissions**](RolesApi.md#ListRolePermissions) | **Get** /roles/{roleId}/permissions | 
[**ListRoleUsers**](RolesApi.md#ListRoleUsers) | **Get** /roles/{roleId}/users | 
[**ListUserRoles**](RolesApi.md#ListUserRoles) | **Get** /users/{userId}/roles | 
[**ResetRole**](RolesApi.md#ResetRole) | **Put** /roles/{roleId} | 
[**UpdateUserRoles**](RolesApi.md#UpdateUserRoles) | **Put** /users/{userId}/roles | 



## AssignRolePermissions

> AssignRolePermissions200Response AssignRolePermissions(ctx, roleId).RolePermissions(rolePermissions).Execute()





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
    roleId := "role_id_1" // string | 
    rolePermissions := *openapiclient.NewRolePermissions([]string{"PermissionIds_example"}) // RolePermissions | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.AssignRolePermissions(context.Background(), roleId).RolePermissions(rolePermissions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.AssignRolePermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignRolePermissions`: AssignRolePermissions200Response
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.AssignRolePermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignRolePermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rolePermissions** | [**RolePermissions**](RolePermissions.md) |  | 

### Return type

[**AssignRolePermissions200Response**](AssignRolePermissions200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRole

> map[string]interface{} DeleteRole(ctx, roleId).Execute()





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
    roleId := "role_id_1" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.DeleteRole(context.Background(), roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.DeleteRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRole`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.DeleteRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRole

> GetRole200Response GetRole(ctx, roleId).Execute()





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
    roleId := "role_id_1" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.GetRole(context.Background(), roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.GetRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRole`: GetRole200Response
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.GetRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRole200Response**](GetRole200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoles

> GetRoles200Response GetRoles(ctx).Q(q).Sort(sort).Execute()





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
    resp, r, err := apiClient.RolesApi.GetRoles(context.Background()).Q(q).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.GetRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoles`: GetRoles200Response
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.GetRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Fulltext search | 
 **sort** | **string** | Sort by property of the object. Use dot notation eg. \&quot;state.city\&quot; for sorting by nested property. Default sort is ASC. Prepend with \&quot;-\&quot; for DESC sort. For numeric sorting use \&quot;size|numeric:true\&quot;. First part is property path, second part is optional parameter of the sort. Supported options are \&quot;numeric:true|false\&quot; for numeric sorting and \&quot;sensitivity:base\&quot; to compare base of the string only: a ≠ b, a &#x3D; á, a &#x3D; A. | 

### Return type

[**GetRoles200Response**](GetRoles200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRolePermissions

> ListPermissions200Response ListRolePermissions(ctx, roleId).Execute()





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
    roleId := "role_id_1" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.ListRolePermissions(context.Background(), roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.ListRolePermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRolePermissions`: ListPermissions200Response
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.ListRolePermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRolePermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListPermissions200Response**](ListPermissions200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoleUsers

> ListRoleUsers200Response ListRoleUsers(ctx, roleId).Q(q).Sort(sort).Execute()





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
    roleId := "role_id_1" // string | 
    q := "John" // string | Fulltext search (optional)
    sort := "firstName" // string | Sort by property of the object. Use dot notation eg. \"state.city\" for sorting by nested property. Default sort is ASC. Prepend with \"-\" for DESC sort. For numeric sorting use \"size|numeric:true\". First part is property path, second part is optional parameter of the sort. Supported options are \"numeric:true|false\" for numeric sorting and \"sensitivity:base\" to compare base of the string only: a ≠ b, a = á, a = A. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.ListRoleUsers(context.Background(), roleId).Q(q).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.ListRoleUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRoleUsers`: ListRoleUsers200Response
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.ListRoleUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRoleUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **string** | Fulltext search | 
 **sort** | **string** | Sort by property of the object. Use dot notation eg. \&quot;state.city\&quot; for sorting by nested property. Default sort is ASC. Prepend with \&quot;-\&quot; for DESC sort. For numeric sorting use \&quot;size|numeric:true\&quot;. First part is property path, second part is optional parameter of the sort. Supported options are \&quot;numeric:true|false\&quot; for numeric sorting and \&quot;sensitivity:base\&quot; to compare base of the string only: a ≠ b, a &#x3D; á, a &#x3D; A. | 

### Return type

[**ListRoleUsers200Response**](ListRoleUsers200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserRoles

> GetRoles200Response ListUserRoles(ctx, userId).Q(q).Sort(sort).Execute()





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
    userId := "user_id_1" // string | 
    q := "John" // string | Fulltext search (optional)
    sort := "firstName" // string | Sort by property of the object. Use dot notation eg. \"state.city\" for sorting by nested property. Default sort is ASC. Prepend with \"-\" for DESC sort. For numeric sorting use \"size|numeric:true\". First part is property path, second part is optional parameter of the sort. Supported options are \"numeric:true|false\" for numeric sorting and \"sensitivity:base\" to compare base of the string only: a ≠ b, a = á, a = A. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.ListUserRoles(context.Background(), userId).Q(q).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.ListUserRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserRoles`: GetRoles200Response
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.ListUserRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **string** | Fulltext search | 
 **sort** | **string** | Sort by property of the object. Use dot notation eg. \&quot;state.city\&quot; for sorting by nested property. Default sort is ASC. Prepend with \&quot;-\&quot; for DESC sort. For numeric sorting use \&quot;size|numeric:true\&quot;. First part is property path, second part is optional parameter of the sort. Supported options are \&quot;numeric:true|false\&quot; for numeric sorting and \&quot;sensitivity:base\&quot; to compare base of the string only: a ≠ b, a &#x3D; á, a &#x3D; A. | 

### Return type

[**GetRoles200Response**](GetRoles200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetRole

> GetRole200Response ResetRole(ctx, roleId).Role(role).Execute()





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
    roleId := "role_id_1" // string | 
    role := *openapiclient.NewRole("RoleName_example", "RoleId_example", "RoleDescription_example") // Role | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.ResetRole(context.Background(), roleId).Role(role).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.ResetRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetRole`: GetRole200Response
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.ResetRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **role** | [**Role**](Role.md) |  | 

### Return type

[**GetRole200Response**](GetRole200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserRoles

> GetUser200Response UpdateUserRoles(ctx, userId).UserRoles(userRoles).Execute()





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
    userId := "user_id_1" // string | 
    userRoles := *openapiclient.NewUserRoles([]string{"RoleIds_example"}) // UserRoles | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.UpdateUserRoles(context.Background(), userId).UserRoles(userRoles).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.UpdateUserRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUserRoles`: GetUser200Response
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.UpdateUserRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userRoles** | [**UserRoles**](UserRoles.md) |  | 

### Return type

[**GetUser200Response**](GetUser200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

