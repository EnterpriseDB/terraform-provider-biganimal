# \UsersApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUser**](UsersApi.md#DeleteUser) | **Delete** /users/{userId} | 
[**GetUser**](UsersApi.md#GetUser) | **Get** /users/{userId} | 
[**GetUserInfo**](UsersApi.md#GetUserInfo) | **Get** /user-info | 
[**GetUsers**](UsersApi.md#GetUsers) | **Get** /users | 
[**ListRoleUsers**](UsersApi.md#ListRoleUsers) | **Get** /roles/{roleId}/users | 
[**ListUserRoles**](UsersApi.md#ListUserRoles) | **Get** /users/{userId}/roles | 
[**UpdateUserRoles**](UsersApi.md#UpdateUserRoles) | **Put** /users/{userId}/roles | 



## DeleteUser

> map[string]interface{} DeleteUser(ctx, userId).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.DeleteUser(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.DeleteUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteUser`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.DeleteUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


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


## GetUser

> GetUser200Response GetUser(ctx, userId).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.GetUser(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUser`: GetUser200Response
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetUser200Response**](GetUser200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserInfo

> GetUserInfo200Response GetUserInfo(ctx).Execute()





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
    resp, r, err := apiClient.UsersApi.GetUserInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetUserInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserInfo`: GetUserInfo200Response
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetUserInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserInfoRequest struct via the builder pattern


### Return type

[**GetUserInfo200Response**](GetUserInfo200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsers

> ListRoleUsers200Response GetUsers(ctx).Q(q).Sort(sort).Execute()





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
    resp, r, err := apiClient.UsersApi.GetUsers(context.Background()).Q(q).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsers`: ListRoleUsers200Response
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersRequest struct via the builder pattern


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
    resp, r, err := apiClient.UsersApi.ListRoleUsers(context.Background(), roleId).Q(q).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ListRoleUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRoleUsers`: ListRoleUsers200Response
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.ListRoleUsers`: %v\n", resp)
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
    resp, r, err := apiClient.UsersApi.ListUserRoles(context.Background(), userId).Q(q).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ListUserRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserRoles`: GetRoles200Response
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.ListUserRoles`: %v\n", resp)
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
    resp, r, err := apiClient.UsersApi.UpdateUserRoles(context.Background(), userId).UserRoles(userRoles).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UpdateUserRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUserRoles`: GetUser200Response
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UpdateUserRoles`: %v\n", resp)
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

