# UpdateKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Allow** | Pointer to [**UpdateKeyRequestAllow**](UpdateKeyRequestAllow.md) |  | [optional] 
**Deny** | Pointer to [**UpdateKeyRequestDeny**](UpdateKeyRequestDeny.md) |  | [optional] 

## Methods

### NewUpdateKeyRequest

`func NewUpdateKeyRequest() *UpdateKeyRequest`

NewUpdateKeyRequest instantiates a new UpdateKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateKeyRequestWithDefaults

`func NewUpdateKeyRequestWithDefaults() *UpdateKeyRequest`

NewUpdateKeyRequestWithDefaults instantiates a new UpdateKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateKeyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateKeyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateKeyRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateKeyRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAllow

`func (o *UpdateKeyRequest) GetAllow() UpdateKeyRequestAllow`

GetAllow returns the Allow field if non-nil, zero value otherwise.

### GetAllowOk

`func (o *UpdateKeyRequest) GetAllowOk() (*UpdateKeyRequestAllow, bool)`

GetAllowOk returns a tuple with the Allow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllow

`func (o *UpdateKeyRequest) SetAllow(v UpdateKeyRequestAllow)`

SetAllow sets Allow field to given value.

### HasAllow

`func (o *UpdateKeyRequest) HasAllow() bool`

HasAllow returns a boolean if a field has been set.

### GetDeny

`func (o *UpdateKeyRequest) GetDeny() UpdateKeyRequestDeny`

GetDeny returns the Deny field if non-nil, zero value otherwise.

### GetDenyOk

`func (o *UpdateKeyRequest) GetDenyOk() (*UpdateKeyRequestDeny, bool)`

GetDenyOk returns a tuple with the Deny field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeny

`func (o *UpdateKeyRequest) SetDeny(v UpdateKeyRequestDeny)`

SetDeny sets Deny field to given value.

### HasDeny

`func (o *UpdateKeyRequest) HasDeny() bool`

HasDeny returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


