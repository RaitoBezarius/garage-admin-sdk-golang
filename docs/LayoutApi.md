# \LayoutApi

All URIs are relative to *http://localhost:3903/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLayout**](LayoutApi.md#AddLayout) | **Post** /layout | Send modifications to the cluster layout
[**ApplyLayout**](LayoutApi.md#ApplyLayout) | **Post** /layout/apply | Apply staged layout
[**GetLayout**](LayoutApi.md#GetLayout) | **Get** /layout | Details on the current and staged layout
[**RevertLayout**](LayoutApi.md#RevertLayout) | **Post** /layout/revert | Clear staged layout



## AddLayout

> ClusterLayout AddLayout(ctx).NodeRoleChange(nodeRoleChange).Execute()

Send modifications to the cluster layout



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
    nodeRoleChange := []openapiclient.NodeRoleChange{openapiclient.NodeRoleChange{NodeRoleRemove: openapiclient.NewNodeRoleRemove("6a8e08af2aab1083ebab9b22165ea8b5b9d333b60a39ecd504e85cc1f432c36f", true)}} // []NodeRoleChange | To add a new node to the layout or to change the configuration of an existing node, simply set the values you want (`zone`, `capacity`, and `tags`). To remove a node, simply pass the `remove: true` field. This logic is represented in OpenAPI with a \"One Of\" object.  Contrary to the CLI that may update only a subset of the fields capacity, zone and tags, when calling this API all of these values must be specified. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LayoutApi.AddLayout(context.Background()).NodeRoleChange(nodeRoleChange).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LayoutApi.AddLayout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddLayout`: ClusterLayout
    fmt.Fprintf(os.Stdout, "Response from `LayoutApi.AddLayout`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddLayoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nodeRoleChange** | [**[]NodeRoleChange**](NodeRoleChange.md) | To add a new node to the layout or to change the configuration of an existing node, simply set the values you want (&#x60;zone&#x60;, &#x60;capacity&#x60;, and &#x60;tags&#x60;). To remove a node, simply pass the &#x60;remove: true&#x60; field. This logic is represented in OpenAPI with a \&quot;One Of\&quot; object.  Contrary to the CLI that may update only a subset of the fields capacity, zone and tags, when calling this API all of these values must be specified.  | 

### Return type

[**ClusterLayout**](ClusterLayout.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplyLayout

> ApplyLayout200Response ApplyLayout(ctx).LayoutVersion(layoutVersion).Execute()

Apply staged layout



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
    layoutVersion := *openapiclient.NewLayoutVersion() // LayoutVersion | Similarly to the CLI, the body must include the version of the new layout that will be created, which MUST be 1 + the value of the currently existing layout in the cluster. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LayoutApi.ApplyLayout(context.Background()).LayoutVersion(layoutVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LayoutApi.ApplyLayout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplyLayout`: ApplyLayout200Response
    fmt.Fprintf(os.Stdout, "Response from `LayoutApi.ApplyLayout`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApplyLayoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **layoutVersion** | [**LayoutVersion**](LayoutVersion.md) | Similarly to the CLI, the body must include the version of the new layout that will be created, which MUST be 1 + the value of the currently existing layout in the cluster.  | 

### Return type

[**ApplyLayout200Response**](ApplyLayout200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLayout

> ClusterLayout GetLayout(ctx).Execute()

Details on the current and staged layout



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
    resp, r, err := apiClient.LayoutApi.GetLayout(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LayoutApi.GetLayout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLayout`: ClusterLayout
    fmt.Fprintf(os.Stdout, "Response from `LayoutApi.GetLayout`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLayoutRequest struct via the builder pattern


### Return type

[**ClusterLayout**](ClusterLayout.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevertLayout

> RevertLayout(ctx).LayoutVersion(layoutVersion).Execute()

Clear staged layout



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
    layoutVersion := *openapiclient.NewLayoutVersion() // LayoutVersion | Reverting the staged changes is done by incrementing the version number and clearing the contents of the staged change list. Similarly to the CLI, the body must include the incremented version number, which MUST be 1 + the value of the currently existing layout in the cluster. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LayoutApi.RevertLayout(context.Background()).LayoutVersion(layoutVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LayoutApi.RevertLayout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRevertLayoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **layoutVersion** | [**LayoutVersion**](LayoutVersion.md) | Reverting the staged changes is done by incrementing the version number and clearing the contents of the staged change list. Similarly to the CLI, the body must include the incremented version number, which MUST be 1 + the value of the currently existing layout in the cluster.  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

