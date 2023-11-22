# \NodesApi

All URIs are relative to *http://localhost:3903/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddNode**](NodesApi.md#AddNode) | **Post** /connect | Connect target node to other Garage nodes
[**GetNodes**](NodesApi.md#GetNodes) | **Get** /status | Status of this node and other nodes in the cluster



## AddNode

> []AddNode200ResponseInner AddNode(ctx).RequestBody(requestBody).Execute()

Connect target node to other Garage nodes



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
    requestBody := []string{"Property_example"} // []string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.AddNode(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.AddNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddNode`: []AddNode200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.AddNode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** |  | 

### Return type

[**[]AddNode200ResponseInner**](AddNode200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodes

> GetNodes200Response GetNodes(ctx).Execute()

Status of this node and other nodes in the cluster



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
    resp, r, err := apiClient.NodesApi.GetNodes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.GetNodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNodes`: GetNodes200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.GetNodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodesRequest struct via the builder pattern


### Return type

[**GetNodes200Response**](GetNodes200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

