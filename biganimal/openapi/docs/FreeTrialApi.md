# \FreeTrialApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountTrialUsage**](FreeTrialApi.md#GetAccountTrialUsage) | **Get** /account/trial-usage | 



## GetAccountTrialUsage

> GetAccountTrialUsage200Response GetAccountTrialUsage(ctx).Execute()





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
    resp, r, err := apiClient.FreeTrialApi.GetAccountTrialUsage(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FreeTrialApi.GetAccountTrialUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountTrialUsage`: GetAccountTrialUsage200Response
    fmt.Fprintf(os.Stdout, "Response from `FreeTrialApi.GetAccountTrialUsage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountTrialUsageRequest struct via the builder pattern


### Return type

[**GetAccountTrialUsage200Response**](GetAccountTrialUsage200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

