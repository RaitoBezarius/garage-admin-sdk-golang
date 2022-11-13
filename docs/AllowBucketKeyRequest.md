# AllowBucketKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketId** | **string** |  | 
**AccessKeyId** | **string** |  | 
**Permissions** | [**AllowBucketKeyRequestPermissions**](AllowBucketKeyRequestPermissions.md) |  | 

## Methods

### NewAllowBucketKeyRequest

`func NewAllowBucketKeyRequest(bucketId string, accessKeyId string, permissions AllowBucketKeyRequestPermissions, ) *AllowBucketKeyRequest`

NewAllowBucketKeyRequest instantiates a new AllowBucketKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllowBucketKeyRequestWithDefaults

`func NewAllowBucketKeyRequestWithDefaults() *AllowBucketKeyRequest`

NewAllowBucketKeyRequestWithDefaults instantiates a new AllowBucketKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketId

`func (o *AllowBucketKeyRequest) GetBucketId() string`

GetBucketId returns the BucketId field if non-nil, zero value otherwise.

### GetBucketIdOk

`func (o *AllowBucketKeyRequest) GetBucketIdOk() (*string, bool)`

GetBucketIdOk returns a tuple with the BucketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketId

`func (o *AllowBucketKeyRequest) SetBucketId(v string)`

SetBucketId sets BucketId field to given value.


### GetAccessKeyId

`func (o *AllowBucketKeyRequest) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *AllowBucketKeyRequest) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *AllowBucketKeyRequest) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.


### GetPermissions

`func (o *AllowBucketKeyRequest) GetPermissions() AllowBucketKeyRequestPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *AllowBucketKeyRequest) GetPermissionsOk() (*AllowBucketKeyRequestPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *AllowBucketKeyRequest) SetPermissions(v AllowBucketKeyRequestPermissions)`

SetPermissions sets Permissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


