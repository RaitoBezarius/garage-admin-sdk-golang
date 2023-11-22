# NodeRoleUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Zone** | **string** |  | 
**Capacity** | **NullableInt32** |  | 
**Tags** | **[]string** |  | 

## Methods

### NewNodeRoleUpdate

`func NewNodeRoleUpdate(zone string, capacity NullableInt32, tags []string, ) *NodeRoleUpdate`

NewNodeRoleUpdate instantiates a new NodeRoleUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeRoleUpdateWithDefaults

`func NewNodeRoleUpdateWithDefaults() *NodeRoleUpdate`

NewNodeRoleUpdateWithDefaults instantiates a new NodeRoleUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NodeRoleUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NodeRoleUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NodeRoleUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NodeRoleUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetZone

`func (o *NodeRoleUpdate) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *NodeRoleUpdate) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *NodeRoleUpdate) SetZone(v string)`

SetZone sets Zone field to given value.


### GetCapacity

`func (o *NodeRoleUpdate) GetCapacity() int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *NodeRoleUpdate) GetCapacityOk() (*int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *NodeRoleUpdate) SetCapacity(v int32)`

SetCapacity sets Capacity field to given value.


### SetCapacityNil

`func (o *NodeRoleUpdate) SetCapacityNil(b bool)`

 SetCapacityNil sets the value for Capacity to be an explicit nil

### UnsetCapacity
`func (o *NodeRoleUpdate) UnsetCapacity()`

UnsetCapacity ensures that no value is present for Capacity, not even an explicit nil
### GetTags

`func (o *NodeRoleUpdate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *NodeRoleUpdate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *NodeRoleUpdate) SetTags(v []string)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


