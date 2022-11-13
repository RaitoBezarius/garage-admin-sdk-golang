# CreateBucketRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GlobalAlias** | Pointer to **string** |  | [optional] 
**LocalAlias** | Pointer to [**CreateBucketRequestLocalAlias**](CreateBucketRequestLocalAlias.md) |  | [optional] 

## Methods

### NewCreateBucketRequest

`func NewCreateBucketRequest() *CreateBucketRequest`

NewCreateBucketRequest instantiates a new CreateBucketRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBucketRequestWithDefaults

`func NewCreateBucketRequestWithDefaults() *CreateBucketRequest`

NewCreateBucketRequestWithDefaults instantiates a new CreateBucketRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGlobalAlias

`func (o *CreateBucketRequest) GetGlobalAlias() string`

GetGlobalAlias returns the GlobalAlias field if non-nil, zero value otherwise.

### GetGlobalAliasOk

`func (o *CreateBucketRequest) GetGlobalAliasOk() (*string, bool)`

GetGlobalAliasOk returns a tuple with the GlobalAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalAlias

`func (o *CreateBucketRequest) SetGlobalAlias(v string)`

SetGlobalAlias sets GlobalAlias field to given value.

### HasGlobalAlias

`func (o *CreateBucketRequest) HasGlobalAlias() bool`

HasGlobalAlias returns a boolean if a field has been set.

### GetLocalAlias

`func (o *CreateBucketRequest) GetLocalAlias() CreateBucketRequestLocalAlias`

GetLocalAlias returns the LocalAlias field if non-nil, zero value otherwise.

### GetLocalAliasOk

`func (o *CreateBucketRequest) GetLocalAliasOk() (*CreateBucketRequestLocalAlias, bool)`

GetLocalAliasOk returns a tuple with the LocalAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAlias

`func (o *CreateBucketRequest) SetLocalAlias(v CreateBucketRequestLocalAlias)`

SetLocalAlias sets LocalAlias field to given value.

### HasLocalAlias

`func (o *CreateBucketRequest) HasLocalAlias() bool`

HasLocalAlias returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


