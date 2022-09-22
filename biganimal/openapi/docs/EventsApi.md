# \EventsApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListMetadata**](EventsApi.md#ListMetadata) | **Get** /events-metadata | 
[**RunSearchForEvents**](EventsApi.md#RunSearchForEvents) | **Post** /events | 



## ListMetadata

> ListMetadata200Response ListMetadata(ctx).StartAt(startAt).EndAt(endAt).Execute()





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
    resp, r, err := apiClient.EventsApi.ListMetadata(context.Background()).StartAt(startAt).EndAt(endAt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.ListMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMetadata`: ListMetadata200Response
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.ListMetadata`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startAt** | **string** |  | 
 **endAt** | **string** |  | 

### Return type

[**ListMetadata200Response**](ListMetadata200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunSearchForEvents

> RunSearchForEvents200Response RunSearchForEvents(ctx).RunSearchForEventsRequest(runSearchForEventsRequest).Execute()





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
    runSearchForEventsRequest := *openapiclient.NewRunSearchForEventsRequest(*openapiclient.NewRunSearchForEventsRequestPaging(), "StartAt_example", "EndAt_example") // RunSearchForEventsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.RunSearchForEvents(context.Background()).RunSearchForEventsRequest(runSearchForEventsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.RunSearchForEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RunSearchForEvents`: RunSearchForEvents200Response
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.RunSearchForEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRunSearchForEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **runSearchForEventsRequest** | [**RunSearchForEventsRequest**](RunSearchForEventsRequest.md) |  | 

### Return type

[**RunSearchForEvents200Response**](RunSearchForEvents200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

