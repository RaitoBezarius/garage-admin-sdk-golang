# KeyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**AccessKeyId** | Pointer to **string** |  | [optional] 
**SecretAccessKey** | Pointer to **NullableString** |  | [optional] 
**Permissions** | Pointer to [**KeyInfoPermissions**](KeyInfoPermissions.md) |  | [optional] 
**Buckets** | Pointer to [**[]KeyInfoBucketsInner**](KeyInfoBucketsInner.md) |  | [optional] 

## Methods

### NewKeyInfo

`func NewKeyInfo() *KeyInfo`

NewKeyInfo instantiates a new KeyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyInfoWithDefaults

`func NewKeyInfoWithDefaults() *KeyInfo`

NewKeyInfoWithDefaults instantiates a new KeyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KeyInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KeyInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccessKeyId

`func (o *KeyInfo) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *KeyInfo) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *KeyInfo) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *KeyInfo) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### GetSecretAccessKey

`func (o *KeyInfo) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *KeyInfo) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *KeyInfo) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.

### HasSecretAccessKey

`func (o *KeyInfo) HasSecretAccessKey() bool`

HasSecretAccessKey returns a boolean if a field has been set.

### SetSecretAccessKeyNil

`func (o *KeyInfo) SetSecretAccessKeyNil(b bool)`

 SetSecretAccessKeyNil sets the value for SecretAccessKey to be an explicit nil

### UnsetSecretAccessKey
`func (o *KeyInfo) UnsetSecretAccessKey()`

UnsetSecretAccessKey ensures that no value is present for SecretAccessKey, not even an explicit nil
### GetPermissions

`func (o *KeyInfo) GetPermissions() KeyInfoPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *KeyInfo) GetPermissionsOk() (*KeyInfoPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *KeyInfo) SetPermissions(v KeyInfoPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *KeyInfo) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetBuckets

`func (o *KeyInfo) GetBuckets() []KeyInfoBucketsInner`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *KeyInfo) GetBucketsOk() (*[]KeyInfoBucketsInner, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *KeyInfo) SetBuckets(v []KeyInfoBucketsInner)`

SetBuckets sets Buckets field to given value.

### HasBuckets

`func (o *KeyInfo) HasBuckets() bool`

HasBuckets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


