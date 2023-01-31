# BucketInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**GlobalAliases** | Pointer to **[]string** |  | [optional] 
**WebsiteAccess** | Pointer to **bool** |  | [optional] 
**WebsiteConfig** | Pointer to [**NullableBucketInfoWebsiteConfig**](BucketInfoWebsiteConfig.md) |  | [optional] 
**Keys** | Pointer to [**[]BucketKeyInfo**](BucketKeyInfo.md) |  | [optional] 
**Objects** | Pointer to **int64** |  | [optional] 
**Bytes** | Pointer to **int64** |  | [optional] 
**UnfinishedUploads** | Pointer to **int32** |  | [optional] 
**Quotas** | Pointer to [**BucketInfoQuotas**](BucketInfoQuotas.md) |  | [optional] 

## Methods

### NewBucketInfo

`func NewBucketInfo() *BucketInfo`

NewBucketInfo instantiates a new BucketInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketInfoWithDefaults

`func NewBucketInfoWithDefaults() *BucketInfo`

NewBucketInfoWithDefaults instantiates a new BucketInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BucketInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BucketInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BucketInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BucketInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGlobalAliases

`func (o *BucketInfo) GetGlobalAliases() []string`

GetGlobalAliases returns the GlobalAliases field if non-nil, zero value otherwise.

### GetGlobalAliasesOk

`func (o *BucketInfo) GetGlobalAliasesOk() (*[]string, bool)`

GetGlobalAliasesOk returns a tuple with the GlobalAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalAliases

`func (o *BucketInfo) SetGlobalAliases(v []string)`

SetGlobalAliases sets GlobalAliases field to given value.

### HasGlobalAliases

`func (o *BucketInfo) HasGlobalAliases() bool`

HasGlobalAliases returns a boolean if a field has been set.

### GetWebsiteAccess

`func (o *BucketInfo) GetWebsiteAccess() bool`

GetWebsiteAccess returns the WebsiteAccess field if non-nil, zero value otherwise.

### GetWebsiteAccessOk

`func (o *BucketInfo) GetWebsiteAccessOk() (*bool, bool)`

GetWebsiteAccessOk returns a tuple with the WebsiteAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteAccess

`func (o *BucketInfo) SetWebsiteAccess(v bool)`

SetWebsiteAccess sets WebsiteAccess field to given value.

### HasWebsiteAccess

`func (o *BucketInfo) HasWebsiteAccess() bool`

HasWebsiteAccess returns a boolean if a field has been set.

### GetWebsiteConfig

`func (o *BucketInfo) GetWebsiteConfig() BucketInfoWebsiteConfig`

GetWebsiteConfig returns the WebsiteConfig field if non-nil, zero value otherwise.

### GetWebsiteConfigOk

`func (o *BucketInfo) GetWebsiteConfigOk() (*BucketInfoWebsiteConfig, bool)`

GetWebsiteConfigOk returns a tuple with the WebsiteConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteConfig

`func (o *BucketInfo) SetWebsiteConfig(v BucketInfoWebsiteConfig)`

SetWebsiteConfig sets WebsiteConfig field to given value.

### HasWebsiteConfig

`func (o *BucketInfo) HasWebsiteConfig() bool`

HasWebsiteConfig returns a boolean if a field has been set.

### SetWebsiteConfigNil

`func (o *BucketInfo) SetWebsiteConfigNil(b bool)`

 SetWebsiteConfigNil sets the value for WebsiteConfig to be an explicit nil

### UnsetWebsiteConfig
`func (o *BucketInfo) UnsetWebsiteConfig()`

UnsetWebsiteConfig ensures that no value is present for WebsiteConfig, not even an explicit nil
### GetKeys

`func (o *BucketInfo) GetKeys() []BucketKeyInfo`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *BucketInfo) GetKeysOk() (*[]BucketKeyInfo, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *BucketInfo) SetKeys(v []BucketKeyInfo)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *BucketInfo) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetObjects

`func (o *BucketInfo) GetObjects() int64`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *BucketInfo) GetObjectsOk() (*int64, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *BucketInfo) SetObjects(v int64)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *BucketInfo) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetBytes

`func (o *BucketInfo) GetBytes() int64`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *BucketInfo) GetBytesOk() (*int64, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *BucketInfo) SetBytes(v int64)`

SetBytes sets Bytes field to given value.

### HasBytes

`func (o *BucketInfo) HasBytes() bool`

HasBytes returns a boolean if a field has been set.

### GetUnfinishedUploads

`func (o *BucketInfo) GetUnfinishedUploads() int32`

GetUnfinishedUploads returns the UnfinishedUploads field if non-nil, zero value otherwise.

### GetUnfinishedUploadsOk

`func (o *BucketInfo) GetUnfinishedUploadsOk() (*int32, bool)`

GetUnfinishedUploadsOk returns a tuple with the UnfinishedUploads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnfinishedUploads

`func (o *BucketInfo) SetUnfinishedUploads(v int32)`

SetUnfinishedUploads sets UnfinishedUploads field to given value.

### HasUnfinishedUploads

`func (o *BucketInfo) HasUnfinishedUploads() bool`

HasUnfinishedUploads returns a boolean if a field has been set.

### GetQuotas

`func (o *BucketInfo) GetQuotas() BucketInfoQuotas`

GetQuotas returns the Quotas field if non-nil, zero value otherwise.

### GetQuotasOk

`func (o *BucketInfo) GetQuotasOk() (*BucketInfoQuotas, bool)`

GetQuotasOk returns a tuple with the Quotas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotas

`func (o *BucketInfo) SetQuotas(v BucketInfoQuotas)`

SetQuotas sets Quotas field to given value.

### HasQuotas

`func (o *BucketInfo) HasQuotas() bool`

HasQuotas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


