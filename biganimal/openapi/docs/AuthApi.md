# \AuthApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBigAnimalTokenByRawToken**](AuthApi.md#CreateBigAnimalTokenByRawToken) | **Post** /auth/token | 



## CreateBigAnimalTokenByRawToken

> AuthToken CreateBigAnimalTokenByRawToken(ctx).AuthToken(authToken).Execute()





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
    authToken := *openapiclient.NewAuthToken("Token_example") // AuthToken | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.CreateBigAnimalTokenByRawToken(context.Background()).AuthToken(authToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.CreateBigAnimalTokenByRawToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBigAnimalTokenByRawToken`: AuthToken
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.CreateBigAnimalTokenByRawToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBigAnimalTokenByRawTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authToken** | [**AuthToken**](AuthToken.md) |  | 

### Return type

[**AuthToken**](AuthToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

