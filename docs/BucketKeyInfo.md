# BucketKeyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeyId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to [**CreateBucketRequestLocalAliasAllow**](CreateBucketRequestLocalAliasAllow.md) |  | [optional] 
**BucketLocalAliases** | Pointer to **[]string** |  | [optional] 

## Methods

### NewBucketKeyInfo

`func NewBucketKeyInfo() *BucketKeyInfo`

NewBucketKeyInfo instantiates a new BucketKeyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketKeyInfoWithDefaults

`func NewBucketKeyInfoWithDefaults() *BucketKeyInfo`

NewBucketKeyInfoWithDefaults instantiates a new BucketKeyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKeyId

`func (o *BucketKeyInfo) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *BucketKeyInfo) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *BucketKeyInfo) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *BucketKeyInfo) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### GetName

`func (o *BucketKeyInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BucketKeyInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BucketKeyInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BucketKeyInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPermissions

`func (o *BucketKeyInfo) GetPermissions() CreateBucketRequestLocalAliasAllow`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *BucketKeyInfo) GetPermissionsOk() (*CreateBucketRequestLocalAliasAllow, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *BucketKeyInfo) SetPermissions(v CreateBucketRequestLocalAliasAllow)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *BucketKeyInfo) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetBucketLocalAliases

`func (o *BucketKeyInfo) GetBucketLocalAliases() []string`

GetBucketLocalAliases returns the BucketLocalAliases field if non-nil, zero value otherwise.

### GetBucketLocalAliasesOk

`func (o *BucketKeyInfo) GetBucketLocalAliasesOk() (*[]string, bool)`

GetBucketLocalAliasesOk returns a tuple with the BucketLocalAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketLocalAliases

`func (o *BucketKeyInfo) SetBucketLocalAliases(v []string)`

SetBucketLocalAliases sets BucketLocalAliases field to given value.

### HasBucketLocalAliases

`func (o *BucketKeyInfo) HasBucketLocalAliases() bool`

HasBucketLocalAliases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


