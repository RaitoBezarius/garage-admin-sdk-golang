/*
Garage Administration API v0+garage-v0.9.0

Administrate your Garage cluster programatically, including status, layout, keys, buckets, and maintainance tasks.  *Disclaimer: The API is not stable yet, hence its v0 tag. The API can change at any time, and changes can include breaking backward compatibility. Read the changelog and upgrade your scripts before upgrading. Additionnaly, this specification is very early stage and can contain bugs, especially on error return codes/types that are not tested yet. Do not expect a well finished and polished product!* 

API version: v0.9.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package garage

import (
	"encoding/json"
)

// NodeRoleUpdate struct for NodeRoleUpdate
type NodeRoleUpdate struct {
	Id string `json:"id"`
	Zone string `json:"zone"`
	Capacity NullableInt64 `json:"capacity"`
	Tags []string `json:"tags"`
}

// NewNodeRoleUpdate instantiates a new NodeRoleUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeRoleUpdate(id string, zone string, capacity NullableInt64, tags []string) *NodeRoleUpdate {
	this := NodeRoleUpdate{}
	this.Id = id
	this.Zone = zone
	this.Capacity = capacity
	this.Tags = tags
	return &this
}

// NewNodeRoleUpdateWithDefaults instantiates a new NodeRoleUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeRoleUpdateWithDefaults() *NodeRoleUpdate {
	this := NodeRoleUpdate{}
	return &this
}

// GetId returns the Id field value
func (o *NodeRoleUpdate) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NodeRoleUpdate) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NodeRoleUpdate) SetId(v string) {
	o.Id = v
}

// GetZone returns the Zone field value
func (o *NodeRoleUpdate) GetZone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Zone
}

// GetZoneOk returns a tuple with the Zone field value
// and a boolean to check if the value has been set.
func (o *NodeRoleUpdate) GetZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Zone, true
}

// SetZone sets field value
func (o *NodeRoleUpdate) SetZone(v string) {
	o.Zone = v
}

// GetCapacity returns the Capacity field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *NodeRoleUpdate) GetCapacity() int64 {
	if o == nil || o.Capacity.Get() == nil {
		var ret int64
		return ret
	}

	return *o.Capacity.Get()
}

// GetCapacityOk returns a tuple with the Capacity field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NodeRoleUpdate) GetCapacityOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Capacity.Get(), o.Capacity.IsSet()
}

// SetCapacity sets field value
func (o *NodeRoleUpdate) SetCapacity(v int64) {
	o.Capacity.Set(&v)
}

// GetTags returns the Tags field value
func (o *NodeRoleUpdate) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *NodeRoleUpdate) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *NodeRoleUpdate) SetTags(v []string) {
	o.Tags = v
}

func (o NodeRoleUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["zone"] = o.Zone
	}
	if true {
		toSerialize["capacity"] = o.Capacity.Get()
	}
	if true {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableNodeRoleUpdate struct {
	value *NodeRoleUpdate
	isSet bool
}

func (v NullableNodeRoleUpdate) Get() *NodeRoleUpdate {
	return v.value
}

func (v *NullableNodeRoleUpdate) Set(val *NodeRoleUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeRoleUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeRoleUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeRoleUpdate(val *NodeRoleUpdate) *NullableNodeRoleUpdate {
	return &NullableNodeRoleUpdate{value: val, isSet: true}
}

func (v NullableNodeRoleUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeRoleUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

