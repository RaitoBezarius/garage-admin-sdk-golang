# UpdateBucketRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebsiteAccess** | Pointer to [**UpdateBucketRequestWebsiteAccess**](UpdateBucketRequestWebsiteAccess.md) |  | [optional] 
**Quotas** | Pointer to [**UpdateBucketRequestQuotas**](UpdateBucketRequestQuotas.md) |  | [optional] 

## Methods

### NewUpdateBucketRequest

`func NewUpdateBucketRequest() *UpdateBucketRequest`

NewUpdateBucketRequest instantiates a new UpdateBucketRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBucketRequestWithDefaults

`func NewUpdateBucketRequestWithDefaults() *UpdateBucketRequest`

NewUpdateBucketRequestWithDefaults instantiates a new UpdateBucketRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebsiteAccess

`func (o *UpdateBucketRequest) GetWebsiteAccess() UpdateBucketRequestWebsiteAccess`

GetWebsiteAccess returns the WebsiteAccess field if non-nil, zero value otherwise.

### GetWebsiteAccessOk

`func (o *UpdateBucketRequest) GetWebsiteAccessOk() (*UpdateBucketRequestWebsiteAccess, bool)`

GetWebsiteAccessOk returns a tuple with the WebsiteAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteAccess

`func (o *UpdateBucketRequest) SetWebsiteAccess(v UpdateBucketRequestWebsiteAccess)`

SetWebsiteAccess sets WebsiteAccess field to given value.

### HasWebsiteAccess

`func (o *UpdateBucketRequest) HasWebsiteAccess() bool`

HasWebsiteAccess returns a boolean if a field has been set.

### GetQuotas

`func (o *UpdateBucketRequest) GetQuotas() UpdateBucketRequestQuotas`

GetQuotas returns the Quotas field if non-nil, zero value otherwise.

### GetQuotasOk

`func (o *UpdateBucketRequest) GetQuotasOk() (*UpdateBucketRequestQuotas, bool)`

GetQuotasOk returns a tuple with the Quotas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotas

`func (o *UpdateBucketRequest) SetQuotas(v UpdateBucketRequestQuotas)`

SetQuotas sets Quotas field to given value.

### HasQuotas

`func (o *UpdateBucketRequest) HasQuotas() bool`

HasQuotas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


