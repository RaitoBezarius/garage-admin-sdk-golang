# NodeRoleChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Remove** | **bool** |  | 
**Zone** | **string** |  | 
**Capacity** | **NullableInt32** |  | 
**Tags** | **[]string** |  | 

## Methods

### NewNodeRoleChange

`func NewNodeRoleChange(remove bool, zone string, capacity NullableInt32, tags []string, ) *NodeRoleChange`

NewNodeRoleChange instantiates a new NodeRoleChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeRoleChangeWithDefaults

`func NewNodeRoleChangeWithDefaults() *NodeRoleChange`

NewNodeRoleChangeWithDefaults instantiates a new NodeRoleChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NodeRoleChange) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NodeRoleChange) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NodeRoleChange) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NodeRoleChange) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRemove

`func (o *NodeRoleChange) GetRemove() bool`

GetRemove returns the Remove field if non-nil, zero value otherwise.

### GetRemoveOk

`func (o *NodeRoleChange) GetRemoveOk() (*bool, bool)`

GetRemoveOk returns a tuple with the Remove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemove

`func (o *NodeRoleChange) SetRemove(v bool)`

SetRemove sets Remove field to given value.


### GetZone

`func (o *NodeRoleChange) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *NodeRoleChange) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *NodeRoleChange) SetZone(v string)`

SetZone sets Zone field to given value.


### GetCapacity

`func (o *NodeRoleChange) GetCapacity() int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *NodeRoleChange) GetCapacityOk() (*int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *NodeRoleChange) SetCapacity(v int32)`

SetCapacity sets Capacity field to given value.


### SetCapacityNil

`func (o *NodeRoleChange) SetCapacityNil(b bool)`

 SetCapacityNil sets the value for Capacity to be an explicit nil

### UnsetCapacity
`func (o *NodeRoleChange) UnsetCapacity()`

UnsetCapacity ensures that no value is present for Capacity, not even an explicit nil
### GetTags

`func (o *NodeRoleChange) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *NodeRoleChange) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *NodeRoleChange) SetTags(v []string)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


