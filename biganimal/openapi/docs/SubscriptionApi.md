# \SubscriptionApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubscription**](SubscriptionApi.md#CreateSubscription) | **Post** /register/{providerId}/sub/{subscriptionId} | 
[**GetBeneficiaryValidation**](SubscriptionApi.md#GetBeneficiaryValidation) | **Get** /register/{providerId}/sub/{subscriptionId}/validation | 
[**GetConnectionList**](SubscriptionApi.md#GetConnectionList) | **Get** /subscriptions/{subscriptionId}/connections | 
[**GetIdentityProviders**](SubscriptionApi.md#GetIdentityProviders) | **Get** /identity-providers | 
[**GetSubscriptionByCode**](SubscriptionApi.md#GetSubscriptionByCode) | **Get** /validate-subscription-code/{code} | 
[**GetSubscriptionById**](SubscriptionApi.md#GetSubscriptionById) | **Get** /subscriptions/{subscriptionId} | 
[**UpdateConnectionConfig**](SubscriptionApi.md#UpdateConnectionConfig) | **Put** /subscriptions/{subscriptionId}/connections | 
[**UpdateConnectionStatus**](SubscriptionApi.md#UpdateConnectionStatus) | **Put** /subscriptions/{subscriptionId}/connections/activate | 



## CreateSubscription

> CreateSubscription201Response CreateSubscription(ctx, providerId, subscriptionId).Body(body).Execute()





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
    providerId := "azure" // string | 
    subscriptionId := "491e0eb1-3757-4b38-dfb6-049442abaeda" // string | 
    body := Subscription(987) // Subscription | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionApi.CreateSubscription(context.Background(), providerId, subscriptionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApi.CreateSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSubscription`: CreateSubscription201Response
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionApi.CreateSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 
**subscriptionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **Subscription** |  | 

### Return type

[**CreateSubscription201Response**](CreateSubscription201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBeneficiaryValidation

> GetBeneficiaryValidation200Response GetBeneficiaryValidation(ctx, providerId, subscriptionId).Execute()





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
    providerId := "azure" // string | 
    subscriptionId := "491e0eb1-3757-4b38-dfb6-049442abaeda" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionApi.GetBeneficiaryValidation(context.Background(), providerId, subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApi.GetBeneficiaryValidation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBeneficiaryValidation`: GetBeneficiaryValidation200Response
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionApi.GetBeneficiaryValidation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 
**subscriptionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBeneficiaryValidationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetBeneficiaryValidation200Response**](GetBeneficiaryValidation200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectionList

> GetConnectionList200Response GetConnectionList(ctx, subscriptionId).Code(code).Execute()





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
    subscriptionId := "cb07adba-dbcd-453f-a562-88f9ff7e4398" // string | Subscription ID
    code := "cb07adba-dbcd-453f-a562-88f9ff7e4398" // string | Code

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionApi.GetConnectionList(context.Background(), subscriptionId).Code(code).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApi.GetConnectionList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnectionList`: GetConnectionList200Response
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionApi.GetConnectionList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **code** | **string** | Code | 

### Return type

[**GetConnectionList200Response**](GetConnectionList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityProviders

> GetConnectionList200Response GetIdentityProviders(ctx).Execute()





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
    resp, r, err := apiClient.SubscriptionApi.GetIdentityProviders(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApi.GetIdentityProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentityProviders`: GetConnectionList200Response
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionApi.GetIdentityProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityProvidersRequest struct via the builder pattern


### Return type

[**GetConnectionList200Response**](GetConnectionList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscriptionByCode

> GetSubscriptionById200Response GetSubscriptionByCode(ctx, code).Execute()





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
    code := "dCPbl7Lzaq1Ehpp2I64BTP2Z2w0tig9d" // string | Code

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionApi.GetSubscriptionByCode(context.Background(), code).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApi.GetSubscriptionByCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubscriptionByCode`: GetSubscriptionById200Response
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionApi.GetSubscriptionByCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionByCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetSubscriptionById200Response**](GetSubscriptionById200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscriptionById

> GetSubscriptionById200Response GetSubscriptionById(ctx, subscriptionId).Code(code).Execute()





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
    subscriptionId := "cb8335a3-ab39-4f5b-b95c-9d14210ce6c1" // string | Subscription ID
    code := "dCPbl7Lzaq1Ehpp2I64BTP2Z2w0tig9d" // string | Code

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionApi.GetSubscriptionById(context.Background(), subscriptionId).Code(code).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApi.GetSubscriptionById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubscriptionById`: GetSubscriptionById200Response
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionApi.GetSubscriptionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **code** | **string** | Code | 

### Return type

[**GetSubscriptionById200Response**](GetSubscriptionById200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConnectionConfig

> GetConnectionList200Response UpdateConnectionConfig(ctx, subscriptionId).Code(code).Connections(connections).Execute()





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
    subscriptionId := "cb07adba-dbcd-453f-a562-88f9ff7e4398" // string | Subscription ID
    code := "cb07adba-dbcd-453f-a562-88f9ff7e4398" // string | Code
    connections := *openapiclient.NewConnections("ConnectionId_example", "Strategy_example", *openapiclient.NewConnectionsBrokerConnOptions()) // Connections | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionApi.UpdateConnectionConfig(context.Background(), subscriptionId).Code(code).Connections(connections).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApi.UpdateConnectionConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConnectionConfig`: GetConnectionList200Response
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionApi.UpdateConnectionConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConnectionConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **code** | **string** | Code | 
 **connections** | [**Connections**](Connections.md) |  | 

### Return type

[**GetConnectionList200Response**](GetConnectionList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConnectionStatus

> map[string]interface{} UpdateConnectionStatus(ctx, subscriptionId).Code(code).Execute()





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
    subscriptionId := "cb07adba-dbcd-453f-a562-88f9ff7e4398" // string | Subscription ID
    code := "cb07adba-dbcd-453f-a562-88f9ff7e4398" // string | Code

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionApi.UpdateConnectionStatus(context.Background(), subscriptionId).Code(code).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApi.UpdateConnectionStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConnectionStatus`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionApi.UpdateConnectionStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConnectionStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **code** | **string** | Code | 

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

