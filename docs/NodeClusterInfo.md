# NodeClusterInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Zone** | **string** |  | 
**Capacity** | **NullableInt32** |  | 
**Tags** | **[]string** | User defined tags, put whatever makes sense for you, these tags are not interpreted by Garage  | 

## Methods

### NewNodeClusterInfo

`func NewNodeClusterInfo(zone string, capacity NullableInt32, tags []string, ) *NodeClusterInfo`

NewNodeClusterInfo instantiates a new NodeClusterInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeClusterInfoWithDefaults

`func NewNodeClusterInfoWithDefaults() *NodeClusterInfo`

NewNodeClusterInfoWithDefaults instantiates a new NodeClusterInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZone

`func (o *NodeClusterInfo) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *NodeClusterInfo) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *NodeClusterInfo) SetZone(v string)`

SetZone sets Zone field to given value.


### GetCapacity

`func (o *NodeClusterInfo) GetCapacity() int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *NodeClusterInfo) GetCapacityOk() (*int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *NodeClusterInfo) SetCapacity(v int32)`

SetCapacity sets Capacity field to given value.


### SetCapacityNil

`func (o *NodeClusterInfo) SetCapacityNil(b bool)`

 SetCapacityNil sets the value for Capacity to be an explicit nil

### UnsetCapacity
`func (o *NodeClusterInfo) UnsetCapacity()`

UnsetCapacity ensures that no value is present for Capacity, not even an explicit nil
### GetTags

`func (o *NodeClusterInfo) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *NodeClusterInfo) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *NodeClusterInfo) SetTags(v []string)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


