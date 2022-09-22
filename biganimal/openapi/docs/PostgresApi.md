# \PostgresApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePgConfigTemplate**](PostgresApi.md#CreatePgConfigTemplate) | **Post** /pg-config-templates | 
[**DeletePgConfigTemplate**](PostgresApi.md#DeletePgConfigTemplate) | **Delete** /pg-config-templates/{pgConfigTemplateId} | 
[**GetPgConfigParameters**](PostgresApi.md#GetPgConfigParameters) | **Get** /pg-config-parameters | 
[**GetPgConfigTemplate**](PostgresApi.md#GetPgConfigTemplate) | **Get** /pg-config-templates/{pgConfigTemplateId} | 
[**GetPgConfigTemplates**](PostgresApi.md#GetPgConfigTemplates) | **Get** /pg-config-templates | 
[**GetPgTypes**](PostgresApi.md#GetPgTypes) | **Get** /pg-types | 
[**GetPgVersions**](PostgresApi.md#GetPgVersions) | **Get** /pg-versions | 
[**RunPgConfigEvaluation**](PostgresApi.md#RunPgConfigEvaluation) | **Put** /pg-config-evaluate | 
[**RunPgConfigTemplatesEvaluation**](PostgresApi.md#RunPgConfigTemplatesEvaluation) | **Put** /pg-config-templates/{pgConfigTemplateId}/evaluate | 
[**UpdatePgConfigTemplate**](PostgresApi.md#UpdatePgConfigTemplate) | **Put** /pg-config-templates/{pgConfigTemplateId} | 



## CreatePgConfigTemplate

> CreatePgConfigTemplate200Response CreatePgConfigTemplate(ctx).CreatePgConfigTemplate(createPgConfigTemplate).Execute()





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
    createPgConfigTemplate := *openapiclient.NewCreatePgConfigTemplate([][]string{[]string{"MachineSettings_example"}}, "PgTypeId_example", "PgVersionId_example", "TemplateDescription_example", "TemplateName_example", [][]string{[]string{"TemplateSettings_example"}}) // CreatePgConfigTemplate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PostgresApi.CreatePgConfigTemplate(context.Background()).CreatePgConfigTemplate(createPgConfigTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresApi.CreatePgConfigTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePgConfigTemplate`: CreatePgConfigTemplate200Response
    fmt.Fprintf(os.Stdout, "Response from `PostgresApi.CreatePgConfigTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePgConfigTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPgConfigTemplate** | [**CreatePgConfigTemplate**](CreatePgConfigTemplate.md) |  | 

### Return type

[**CreatePgConfigTemplate200Response**](CreatePgConfigTemplate200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePgConfigTemplate

> DeletePgConfigTemplate(ctx, pgConfigTemplateId).Execute()





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
    pgConfigTemplateId := "1a164b74-ec89-486a-b914-ccc5af38487d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PostgresApi.DeletePgConfigTemplate(context.Background(), pgConfigTemplateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresApi.DeletePgConfigTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pgConfigTemplateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePgConfigTemplateRequest struct via the builder pattern


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


## GetPgConfigParameters

> GetPgConfigParameters200Response GetPgConfigParameters(ctx).PgTypeId(pgTypeId).PgVersionId(pgVersionId).Q(q).Sort(sort).Execute()





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
    pgTypeId := "epas" // string | pgTypeId
    pgVersionId := "14" // string | pgVersionId
    q := "John" // string | Fulltext search (optional)
    sort := "firstName" // string | Sort by property of the object. Use dot notation eg. \"state.city\" for sorting by nested property. Default sort is ASC. Prepend with \"-\" for DESC sort. For numeric sorting use \"size|numeric:true\". First part is property path, second part is optional parameter of the sort. Supported options are \"numeric:true|false\" for numeric sorting and \"sensitivity:base\" to compare base of the string only: a ≠ b, a = á, a = A. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PostgresApi.GetPgConfigParameters(context.Background()).PgTypeId(pgTypeId).PgVersionId(pgVersionId).Q(q).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresApi.GetPgConfigParameters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPgConfigParameters`: GetPgConfigParameters200Response
    fmt.Fprintf(os.Stdout, "Response from `PostgresApi.GetPgConfigParameters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPgConfigParametersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pgTypeId** | **string** | pgTypeId | 
 **pgVersionId** | **string** | pgVersionId | 
 **q** | **string** | Fulltext search | 
 **sort** | **string** | Sort by property of the object. Use dot notation eg. \&quot;state.city\&quot; for sorting by nested property. Default sort is ASC. Prepend with \&quot;-\&quot; for DESC sort. For numeric sorting use \&quot;size|numeric:true\&quot;. First part is property path, second part is optional parameter of the sort. Supported options are \&quot;numeric:true|false\&quot; for numeric sorting and \&quot;sensitivity:base\&quot; to compare base of the string only: a ≠ b, a &#x3D; á, a &#x3D; A. | 

### Return type

[**GetPgConfigParameters200Response**](GetPgConfigParameters200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPgConfigTemplate

> GetPgConfigTemplate200Response GetPgConfigTemplate(ctx, pgConfigTemplateId).Execute()





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
    pgConfigTemplateId := "77693995-ee51-4a6a-ad71-b8311b666c60" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PostgresApi.GetPgConfigTemplate(context.Background(), pgConfigTemplateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresApi.GetPgConfigTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPgConfigTemplate`: GetPgConfigTemplate200Response
    fmt.Fprintf(os.Stdout, "Response from `PostgresApi.GetPgConfigTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pgConfigTemplateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPgConfigTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetPgConfigTemplate200Response**](GetPgConfigTemplate200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPgConfigTemplates

> GetPgConfigTemplates200Response GetPgConfigTemplates(ctx).Execute()





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
    resp, r, err := apiClient.PostgresApi.GetPgConfigTemplates(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresApi.GetPgConfigTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPgConfigTemplates`: GetPgConfigTemplates200Response
    fmt.Fprintf(os.Stdout, "Response from `PostgresApi.GetPgConfigTemplates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPgConfigTemplatesRequest struct via the builder pattern


### Return type

[**GetPgConfigTemplates200Response**](GetPgConfigTemplates200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPgTypes

> GetPgTypes200Response GetPgTypes(ctx).Q(q).Sort(sort).ClusterArchitectureIds(clusterArchitectureIds).Execute()





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
    clusterArchitectureIds := []string{"Inner_example"} // []string | Array of cluster architecture IDs (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PostgresApi.GetPgTypes(context.Background()).Q(q).Sort(sort).ClusterArchitectureIds(clusterArchitectureIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresApi.GetPgTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPgTypes`: GetPgTypes200Response
    fmt.Fprintf(os.Stdout, "Response from `PostgresApi.GetPgTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPgTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Fulltext search | 
 **sort** | **string** | Sort by property of the object. Use dot notation eg. \&quot;state.city\&quot; for sorting by nested property. Default sort is ASC. Prepend with \&quot;-\&quot; for DESC sort. For numeric sorting use \&quot;size|numeric:true\&quot;. First part is property path, second part is optional parameter of the sort. Supported options are \&quot;numeric:true|false\&quot; for numeric sorting and \&quot;sensitivity:base\&quot; to compare base of the string only: a ≠ b, a &#x3D; á, a &#x3D; A. | 
 **clusterArchitectureIds** | **[]string** | Array of cluster architecture IDs | 

### Return type

[**GetPgTypes200Response**](GetPgTypes200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPgVersions

> GetPgVersions200Response GetPgVersions(ctx).ClusterArchitectureIds(clusterArchitectureIds).Q(q).Sort(sort).Execute()





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
    clusterArchitectureIds := []string{"Inner_example"} // []string | Array of cluster architecture IDs (optional)
    q := "John" // string | Fulltext search (optional)
    sort := "firstName" // string | Sort by property of the object. Use dot notation eg. \"state.city\" for sorting by nested property. Default sort is ASC. Prepend with \"-\" for DESC sort. For numeric sorting use \"size|numeric:true\". First part is property path, second part is optional parameter of the sort. Supported options are \"numeric:true|false\" for numeric sorting and \"sensitivity:base\" to compare base of the string only: a ≠ b, a = á, a = A. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PostgresApi.GetPgVersions(context.Background()).ClusterArchitectureIds(clusterArchitectureIds).Q(q).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresApi.GetPgVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPgVersions`: GetPgVersions200Response
    fmt.Fprintf(os.Stdout, "Response from `PostgresApi.GetPgVersions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPgVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterArchitectureIds** | **[]string** | Array of cluster architecture IDs | 
 **q** | **string** | Fulltext search | 
 **sort** | **string** | Sort by property of the object. Use dot notation eg. \&quot;state.city\&quot; for sorting by nested property. Default sort is ASC. Prepend with \&quot;-\&quot; for DESC sort. For numeric sorting use \&quot;size|numeric:true\&quot;. First part is property path, second part is optional parameter of the sort. Supported options are \&quot;numeric:true|false\&quot; for numeric sorting and \&quot;sensitivity:base\&quot; to compare base of the string only: a ≠ b, a &#x3D; á, a &#x3D; A. | 

### Return type

[**GetPgVersions200Response**](GetPgVersions200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunPgConfigEvaluation

> RunPgConfigEvaluation200Response RunPgConfigEvaluation(ctx).EvaluatePgConfig(evaluatePgConfig).Execute()





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
    evaluatePgConfig := *openapiclient.NewEvaluatePgConfig([][]string{[]string{"MachineSettings_example"}}, [][]string{[]string{"PgConfig_example"}}, "PgTypeId_example", "PgVersionId_example") // EvaluatePgConfig | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PostgresApi.RunPgConfigEvaluation(context.Background()).EvaluatePgConfig(evaluatePgConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresApi.RunPgConfigEvaluation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RunPgConfigEvaluation`: RunPgConfigEvaluation200Response
    fmt.Fprintf(os.Stdout, "Response from `PostgresApi.RunPgConfigEvaluation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRunPgConfigEvaluationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **evaluatePgConfig** | [**EvaluatePgConfig**](EvaluatePgConfig.md) |  | 

### Return type

[**RunPgConfigEvaluation200Response**](RunPgConfigEvaluation200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunPgConfigTemplatesEvaluation

> RunPgConfigTemplatesEvaluation200Response RunPgConfigTemplatesEvaluation(ctx, pgConfigTemplateId).EvaluatePgConfigTemplate(evaluatePgConfigTemplate).Execute()





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
    pgConfigTemplateId := "77693995-ee51-4a6a-ad71-b8311b666c60" // string | 
    evaluatePgConfigTemplate := *openapiclient.NewEvaluatePgConfigTemplate([][]string{[]string{"MachineSettings_example"}}, "PgConfigTemplateId_example") // EvaluatePgConfigTemplate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PostgresApi.RunPgConfigTemplatesEvaluation(context.Background(), pgConfigTemplateId).EvaluatePgConfigTemplate(evaluatePgConfigTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresApi.RunPgConfigTemplatesEvaluation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RunPgConfigTemplatesEvaluation`: RunPgConfigTemplatesEvaluation200Response
    fmt.Fprintf(os.Stdout, "Response from `PostgresApi.RunPgConfigTemplatesEvaluation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pgConfigTemplateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRunPgConfigTemplatesEvaluationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **evaluatePgConfigTemplate** | [**EvaluatePgConfigTemplate**](EvaluatePgConfigTemplate.md) |  | 

### Return type

[**RunPgConfigTemplatesEvaluation200Response**](RunPgConfigTemplatesEvaluation200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePgConfigTemplate

> UpdatePgConfigTemplate200Response UpdatePgConfigTemplate(ctx, pgConfigTemplateId).UpdatePgConfigTemplate(updatePgConfigTemplate).Execute()





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
    pgConfigTemplateId := "77693995-ee51-4a6a-ad71-b8311b666c60" // string | 
    updatePgConfigTemplate := *openapiclient.NewUpdatePgConfigTemplate("TemplateName_example") // UpdatePgConfigTemplate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PostgresApi.UpdatePgConfigTemplate(context.Background(), pgConfigTemplateId).UpdatePgConfigTemplate(updatePgConfigTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresApi.UpdatePgConfigTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePgConfigTemplate`: UpdatePgConfigTemplate200Response
    fmt.Fprintf(os.Stdout, "Response from `PostgresApi.UpdatePgConfigTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pgConfigTemplateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePgConfigTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePgConfigTemplate** | [**UpdatePgConfigTemplate**](UpdatePgConfigTemplate.md) |  | 

### Return type

[**UpdatePgConfigTemplate200Response**](UpdatePgConfigTemplate200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

