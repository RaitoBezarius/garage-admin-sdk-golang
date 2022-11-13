# KeyInfoBucketsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**GlobalAliases** | Pointer to **[]string** |  | [optional] 
**LocalAliases** | Pointer to **[]string** |  | [optional] 
**Permissions** | Pointer to [**KeyInfoBucketsInnerPermissions**](KeyInfoBucketsInnerPermissions.md) |  | [optional] 

## Methods

### NewKeyInfoBucketsInner

`func NewKeyInfoBucketsInner() *KeyInfoBucketsInner`

NewKeyInfoBucketsInner instantiates a new KeyInfoBucketsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyInfoBucketsInnerWithDefaults

`func NewKeyInfoBucketsInnerWithDefaults() *KeyInfoBucketsInner`

NewKeyInfoBucketsInnerWithDefaults instantiates a new KeyInfoBucketsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyInfoBucketsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyInfoBucketsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyInfoBucketsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KeyInfoBucketsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGlobalAliases

`func (o *KeyInfoBucketsInner) GetGlobalAliases() []string`

GetGlobalAliases returns the GlobalAliases field if non-nil, zero value otherwise.

### GetGlobalAliasesOk

`func (o *KeyInfoBucketsInner) GetGlobalAliasesOk() (*[]string, bool)`

GetGlobalAliasesOk returns a tuple with the GlobalAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalAliases

`func (o *KeyInfoBucketsInner) SetGlobalAliases(v []string)`

SetGlobalAliases sets GlobalAliases field to given value.

### HasGlobalAliases

`func (o *KeyInfoBucketsInner) HasGlobalAliases() bool`

HasGlobalAliases returns a boolean if a field has been set.

### GetLocalAliases

`func (o *KeyInfoBucketsInner) GetLocalAliases() []string`

GetLocalAliases returns the LocalAliases field if non-nil, zero value otherwise.

### GetLocalAliasesOk

`func (o *KeyInfoBucketsInner) GetLocalAliasesOk() (*[]string, bool)`

GetLocalAliasesOk returns a tuple with the LocalAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAliases

`func (o *KeyInfoBucketsInner) SetLocalAliases(v []string)`

SetLocalAliases sets LocalAliases field to given value.

### HasLocalAliases

`func (o *KeyInfoBucketsInner) HasLocalAliases() bool`

HasLocalAliases returns a boolean if a field has been set.

### GetPermissions

`func (o *KeyInfoBucketsInner) GetPermissions() KeyInfoBucketsInnerPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *KeyInfoBucketsInner) GetPermissionsOk() (*KeyInfoBucketsInnerPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *KeyInfoBucketsInner) SetPermissions(v KeyInfoBucketsInnerPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *KeyInfoBucketsInner) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


