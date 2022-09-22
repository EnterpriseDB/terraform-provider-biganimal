# \UtilsApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RunCheckForPasswordStrength**](UtilsApi.md#RunCheckForPasswordStrength) | **Put** /utils/password-strength | 



## RunCheckForPasswordStrength

> RunCheckForPasswordStrength200Response RunCheckForPasswordStrength(ctx).RunCheckForPasswordStrengthRequest(runCheckForPasswordStrengthRequest).Execute()





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
    runCheckForPasswordStrengthRequest := *openapiclient.NewRunCheckForPasswordStrengthRequest("Password_example") // RunCheckForPasswordStrengthRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UtilsApi.RunCheckForPasswordStrength(context.Background()).RunCheckForPasswordStrengthRequest(runCheckForPasswordStrengthRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UtilsApi.RunCheckForPasswordStrength``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RunCheckForPasswordStrength`: RunCheckForPasswordStrength200Response
    fmt.Fprintf(os.Stdout, "Response from `UtilsApi.RunCheckForPasswordStrength`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRunCheckForPasswordStrengthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **runCheckForPasswordStrengthRequest** | [**RunCheckForPasswordStrengthRequest**](RunCheckForPasswordStrengthRequest.md) |  | 

### Return type

[**RunCheckForPasswordStrength200Response**](RunCheckForPasswordStrength200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

