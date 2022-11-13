# \BucketApi

All URIs are relative to *http://localhost:3903/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllowBucketKey**](BucketApi.md#AllowBucketKey) | **Post** /bucket/allow | Allow key
[**CreateBucket**](BucketApi.md#CreateBucket) | **Post** /bucket | Create a bucket
[**DeleteBucket**](BucketApi.md#DeleteBucket) | **Delete** /bucket?id&#x3D;{bucket_id} | Delete a bucket
[**DeleteBucketGlobalAlias**](BucketApi.md#DeleteBucketGlobalAlias) | **Delete** /bucket/alias/global | Delete a global alias
[**DeleteBucketLocalAlias**](BucketApi.md#DeleteBucketLocalAlias) | **Delete** /bucket/alias/local | Delete a local alias
[**DenyBucketKey**](BucketApi.md#DenyBucketKey) | **Post** /bucket/deny | Deny key
[**FindBucketInfo**](BucketApi.md#FindBucketInfo) | **Get** /bucket?globalAlias&#x3D;{alias} | Find a bucket
[**GetBucketInfo**](BucketApi.md#GetBucketInfo) | **Get** /bucket?id&#x3D;{bucket_id} | Get a bucket
[**ListBuckets**](BucketApi.md#ListBuckets) | **Get** /bucket | List all buckets
[**PutBucketGlobalAlias**](BucketApi.md#PutBucketGlobalAlias) | **Put** /bucket/alias/global | Add a global alias
[**PutBucketLocalAlias**](BucketApi.md#PutBucketLocalAlias) | **Put** /bucket/alias/local | Add a local alias
[**UpdateBucket**](BucketApi.md#UpdateBucket) | **Put** /bucket?id&#x3D;{bucket_id} | Update a bucket



## AllowBucketKey

> BucketInfo AllowBucketKey(ctx).AllowBucketKeyRequest(allowBucketKeyRequest).Execute()

Allow key



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
    allowBucketKeyRequest := *openapiclient.NewAllowBucketKeyRequest("e6a14cd6a27f48684579ec6b381c078ab11697e6bc8513b72b2f5307e25fff9b", "GK31c2f218a2e44f485b94239e", *openapiclient.NewAllowBucketKeyRequestPermissions(true, true, true)) // AllowBucketKeyRequest | Aliases to put on the new bucket 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.AllowBucketKey(context.Background()).AllowBucketKeyRequest(allowBucketKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.AllowBucketKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AllowBucketKey`: BucketInfo
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.AllowBucketKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAllowBucketKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **allowBucketKeyRequest** | [**AllowBucketKeyRequest**](AllowBucketKeyRequest.md) | Aliases to put on the new bucket  | 

### Return type

[**BucketInfo**](BucketInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBucket

> BucketInfo CreateBucket(ctx).CreateBucketRequest(createBucketRequest).Execute()

Create a bucket



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
    createBucketRequest := *openapiclient.NewCreateBucketRequest() // CreateBucketRequest | Aliases to put on the new bucket 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.CreateBucket(context.Background()).CreateBucketRequest(createBucketRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.CreateBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBucket`: BucketInfo
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.CreateBucket`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBucketRequest** | [**CreateBucketRequest**](CreateBucketRequest.md) | Aliases to put on the new bucket  | 

### Return type

[**BucketInfo**](BucketInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBucket

> DeleteBucket(ctx, bucketId).Execute()

Delete a bucket



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
    bucketId := "b4018dc61b27ccb5c64ec1b24f53454bbbd180697c758c4d47a22a8921864a87" // string | The exact bucket identifier, a 32 bytes hexadecimal string

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.DeleteBucket(context.Background(), bucketId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.DeleteBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **string** | The exact bucket identifier, a 32 bytes hexadecimal string | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBucketGlobalAlias

> BucketInfo DeleteBucketGlobalAlias(ctx).Id(id).Alias(alias).Execute()

Delete a global alias



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
    id := "e6a14cd6a27f48684579ec6b381c078ab11697e6bc8513b72b2f5307e25fff9b" // string | 
    alias := "my_documents" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.DeleteBucketGlobalAlias(context.Background()).Id(id).Alias(alias).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.DeleteBucketGlobalAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBucketGlobalAlias`: BucketInfo
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.DeleteBucketGlobalAlias`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBucketGlobalAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **alias** | **string** |  | 

### Return type

[**BucketInfo**](BucketInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBucketLocalAlias

> BucketInfo DeleteBucketLocalAlias(ctx).Id(id).AccessKeyId(accessKeyId).Alias(alias).Execute()

Delete a local alias



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
    id := "e6a14cd6a27f48684579ec6b381c078ab11697e6bc8513b72b2f5307e25fff9b" // string | 
    accessKeyId := "GK31c2f218a2e44f485b94239e" // string | 
    alias := "my_documents" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.DeleteBucketLocalAlias(context.Background()).Id(id).AccessKeyId(accessKeyId).Alias(alias).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.DeleteBucketLocalAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBucketLocalAlias`: BucketInfo
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.DeleteBucketLocalAlias`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBucketLocalAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **accessKeyId** | **string** |  | 
 **alias** | **string** |  | 

### Return type

[**BucketInfo**](BucketInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DenyBucketKey

> BucketInfo DenyBucketKey(ctx).AllowBucketKeyRequest(allowBucketKeyRequest).Execute()

Deny key



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
    allowBucketKeyRequest := *openapiclient.NewAllowBucketKeyRequest("e6a14cd6a27f48684579ec6b381c078ab11697e6bc8513b72b2f5307e25fff9b", "GK31c2f218a2e44f485b94239e", *openapiclient.NewAllowBucketKeyRequestPermissions(true, true, true)) // AllowBucketKeyRequest | Aliases to put on the new bucket 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.DenyBucketKey(context.Background()).AllowBucketKeyRequest(allowBucketKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.DenyBucketKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DenyBucketKey`: BucketInfo
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.DenyBucketKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDenyBucketKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **allowBucketKeyRequest** | [**AllowBucketKeyRequest**](AllowBucketKeyRequest.md) | Aliases to put on the new bucket  | 

### Return type

[**BucketInfo**](BucketInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindBucketInfo

> BucketInfo FindBucketInfo(ctx, alias).Execute()

Find a bucket



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
    alias := "my_documents" // string | The exact global alias of one of the existing buckets

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.FindBucketInfo(context.Background(), alias).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.FindBucketInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindBucketInfo`: BucketInfo
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.FindBucketInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alias** | **string** | The exact global alias of one of the existing buckets | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindBucketInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BucketInfo**](BucketInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBucketInfo

> BucketInfo GetBucketInfo(ctx, bucketId).Execute()

Get a bucket



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
    bucketId := "b4018dc61b27ccb5c64ec1b24f53454bbbd180697c758c4d47a22a8921864a87" // string | The exact bucket identifier, a 32 bytes hexadecimal string

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.GetBucketInfo(context.Background(), bucketId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.GetBucketInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBucketInfo`: BucketInfo
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.GetBucketInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **string** | The exact bucket identifier, a 32 bytes hexadecimal string | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBucketInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BucketInfo**](BucketInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBuckets

> []ListBuckets200ResponseInner ListBuckets(ctx).Execute()

List all buckets



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
    resp, r, err := apiClient.BucketApi.ListBuckets(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.ListBuckets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuckets`: []ListBuckets200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.ListBuckets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBucketsRequest struct via the builder pattern


### Return type

[**[]ListBuckets200ResponseInner**](ListBuckets200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutBucketGlobalAlias

> BucketInfo PutBucketGlobalAlias(ctx).Id(id).Alias(alias).Execute()

Add a global alias



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
    id := "e6a14cd6a27f48684579ec6b381c078ab11697e6bc8513b72b2f5307e25fff9b" // string | 
    alias := "my_documents" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.PutBucketGlobalAlias(context.Background()).Id(id).Alias(alias).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.PutBucketGlobalAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutBucketGlobalAlias`: BucketInfo
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.PutBucketGlobalAlias`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutBucketGlobalAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **alias** | **string** |  | 

### Return type

[**BucketInfo**](BucketInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutBucketLocalAlias

> BucketInfo PutBucketLocalAlias(ctx).Id(id).AccessKeyId(accessKeyId).Alias(alias).Execute()

Add a local alias



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
    id := "e6a14cd6a27f48684579ec6b381c078ab11697e6bc8513b72b2f5307e25fff9b" // string | 
    accessKeyId := "GK31c2f218a2e44f485b94239e" // string | 
    alias := "my_documents" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.PutBucketLocalAlias(context.Background()).Id(id).AccessKeyId(accessKeyId).Alias(alias).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.PutBucketLocalAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutBucketLocalAlias`: BucketInfo
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.PutBucketLocalAlias`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutBucketLocalAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **accessKeyId** | **string** |  | 
 **alias** | **string** |  | 

### Return type

[**BucketInfo**](BucketInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBucket

> BucketInfo UpdateBucket(ctx, bucketId).UpdateBucketRequest(updateBucketRequest).Execute()

Update a bucket



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
    bucketId := "b4018dc61b27ccb5c64ec1b24f53454bbbd180697c758c4d47a22a8921864a87" // string | The exact bucket identifier, a 32 bytes hexadecimal string
    updateBucketRequest := *openapiclient.NewUpdateBucketRequest() // UpdateBucketRequest | Requested changes on the bucket. Both root fields are optionals. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.UpdateBucket(context.Background(), bucketId).UpdateBucketRequest(updateBucketRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.UpdateBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBucket`: BucketInfo
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.UpdateBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **string** | The exact bucket identifier, a 32 bytes hexadecimal string | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateBucketRequest** | [**UpdateBucketRequest**](UpdateBucketRequest.md) | Requested changes on the bucket. Both root fields are optionals.  | 

### Return type

[**BucketInfo**](BucketInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

