# ListBuckets200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**GlobalAliases** | Pointer to **[]string** |  | [optional] 
**LocalAliases** | Pointer to [**[]ListBuckets200ResponseInnerLocalAliasesInner**](ListBuckets200ResponseInnerLocalAliasesInner.md) |  | [optional] 

## Methods

### NewListBuckets200ResponseInner

`func NewListBuckets200ResponseInner(id string, ) *ListBuckets200ResponseInner`

NewListBuckets200ResponseInner instantiates a new ListBuckets200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBuckets200ResponseInnerWithDefaults

`func NewListBuckets200ResponseInnerWithDefaults() *ListBuckets200ResponseInner`

NewListBuckets200ResponseInnerWithDefaults instantiates a new ListBuckets200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListBuckets200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListBuckets200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListBuckets200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetGlobalAliases

`func (o *ListBuckets200ResponseInner) GetGlobalAliases() []string`

GetGlobalAliases returns the GlobalAliases field if non-nil, zero value otherwise.

### GetGlobalAliasesOk

`func (o *ListBuckets200ResponseInner) GetGlobalAliasesOk() (*[]string, bool)`

GetGlobalAliasesOk returns a tuple with the GlobalAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalAliases

`func (o *ListBuckets200ResponseInner) SetGlobalAliases(v []string)`

SetGlobalAliases sets GlobalAliases field to given value.

### HasGlobalAliases

`func (o *ListBuckets200ResponseInner) HasGlobalAliases() bool`

HasGlobalAliases returns a boolean if a field has been set.

### GetLocalAliases

`func (o *ListBuckets200ResponseInner) GetLocalAliases() []ListBuckets200ResponseInnerLocalAliasesInner`

GetLocalAliases returns the LocalAliases field if non-nil, zero value otherwise.

### GetLocalAliasesOk

`func (o *ListBuckets200ResponseInner) GetLocalAliasesOk() (*[]ListBuckets200ResponseInnerLocalAliasesInner, bool)`

GetLocalAliasesOk returns a tuple with the LocalAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAliases

`func (o *ListBuckets200ResponseInner) SetLocalAliases(v []ListBuckets200ResponseInnerLocalAliasesInner)`

SetLocalAliases sets LocalAliases field to given value.

### HasLocalAliases

`func (o *ListBuckets200ResponseInner) HasLocalAliases() bool`

HasLocalAliases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


