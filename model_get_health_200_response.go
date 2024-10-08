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

// GetHealth200Response struct for GetHealth200Response
type GetHealth200Response struct {
	Status string `json:"status"`
	KnownNodes int64 `json:"knownNodes"`
	ConnectedNodes int64 `json:"connectedNodes"`
	StorageNodes int64 `json:"storageNodes"`
	StorageNodesOk int64 `json:"storageNodesOk"`
	Partitions int64 `json:"partitions"`
	PartitionsQuorum int64 `json:"partitionsQuorum"`
	PartitionsAllOk int64 `json:"partitionsAllOk"`
}

// NewGetHealth200Response instantiates a new GetHealth200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetHealth200Response(status string, knownNodes int64, connectedNodes int64, storageNodes int64, storageNodesOk int64, partitions int64, partitionsQuorum int64, partitionsAllOk int64) *GetHealth200Response {
	this := GetHealth200Response{}
	this.Status = status
	this.KnownNodes = knownNodes
	this.ConnectedNodes = connectedNodes
	this.StorageNodes = storageNodes
	this.StorageNodesOk = storageNodesOk
	this.Partitions = partitions
	this.PartitionsQuorum = partitionsQuorum
	this.PartitionsAllOk = partitionsAllOk
	return &this
}

// NewGetHealth200ResponseWithDefaults instantiates a new GetHealth200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetHealth200ResponseWithDefaults() *GetHealth200Response {
	this := GetHealth200Response{}
	return &this
}

// GetStatus returns the Status field value
func (o *GetHealth200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetHealth200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetHealth200Response) SetStatus(v string) {
	o.Status = v
}

// GetKnownNodes returns the KnownNodes field value
func (o *GetHealth200Response) GetKnownNodes() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.KnownNodes
}

// GetKnownNodesOk returns a tuple with the KnownNodes field value
// and a boolean to check if the value has been set.
func (o *GetHealth200Response) GetKnownNodesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KnownNodes, true
}

// SetKnownNodes sets field value
func (o *GetHealth200Response) SetKnownNodes(v int64) {
	o.KnownNodes = v
}

// GetConnectedNodes returns the ConnectedNodes field value
func (o *GetHealth200Response) GetConnectedNodes() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ConnectedNodes
}

// GetConnectedNodesOk returns a tuple with the ConnectedNodes field value
// and a boolean to check if the value has been set.
func (o *GetHealth200Response) GetConnectedNodesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectedNodes, true
}

// SetConnectedNodes sets field value
func (o *GetHealth200Response) SetConnectedNodes(v int64) {
	o.ConnectedNodes = v
}

// GetStorageNodes returns the StorageNodes field value
func (o *GetHealth200Response) GetStorageNodesAll() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StorageNodes
}

// GetStorageNodesOk returns a tuple with the StorageNodes field value
// and a boolean to check if the value has been set.
func (o *GetHealth200Response) GetStorageNodesAllOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageNodes, true
}

// SetStorageNodes sets field value
func (o *GetHealth200Response) SetStorageNodes(v int64) {
	o.StorageNodes = v
}

// GetStorageNodesOk returns the StorageNodesOk field value
func (o *GetHealth200Response) GetStorageNodesOk() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StorageNodesOk
}

// GetStorageNodesOkOk returns a tuple with the StorageNodesOk field value
// and a boolean to check if the value has been set.
func (o *GetHealth200Response) GetStorageNodesOkOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageNodesOk, true
}

// SetStorageNodesOk sets field value
func (o *GetHealth200Response) SetStorageNodesOk(v int64) {
	o.StorageNodesOk = v
}

// GetPartitions returns the Partitions field value
func (o *GetHealth200Response) GetPartitions() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Partitions
}

// GetPartitionsOk returns a tuple with the Partitions field value
// and a boolean to check if the value has been set.
func (o *GetHealth200Response) GetPartitionsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Partitions, true
}

// SetPartitions sets field value
func (o *GetHealth200Response) SetPartitions(v int64) {
	o.Partitions = v
}

// GetPartitionsQuorum returns the PartitionsQuorum field value
func (o *GetHealth200Response) GetPartitionsQuorum() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PartitionsQuorum
}

// GetPartitionsQuorumOk returns a tuple with the PartitionsQuorum field value
// and a boolean to check if the value has been set.
func (o *GetHealth200Response) GetPartitionsQuorumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PartitionsQuorum, true
}

// SetPartitionsQuorum sets field value
func (o *GetHealth200Response) SetPartitionsQuorum(v int64) {
	o.PartitionsQuorum = v
}

// GetPartitionsAllOk returns the PartitionsAllOk field value
func (o *GetHealth200Response) GetPartitionsAllOk() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PartitionsAllOk
}

// GetPartitionsAllOkOk returns a tuple with the PartitionsAllOk field value
// and a boolean to check if the value has been set.
func (o *GetHealth200Response) GetPartitionsAllOkOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PartitionsAllOk, true
}

// SetPartitionsAllOk sets field value
func (o *GetHealth200Response) SetPartitionsAllOk(v int64) {
	o.PartitionsAllOk = v
}

func (o GetHealth200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["knownNodes"] = o.KnownNodes
	}
	if true {
		toSerialize["connectedNodes"] = o.ConnectedNodes
	}
	if true {
		toSerialize["storageNodes"] = o.StorageNodes
	}
	if true {
		toSerialize["storageNodesOk"] = o.StorageNodesOk
	}
	if true {
		toSerialize["partitions"] = o.Partitions
	}
	if true {
		toSerialize["partitionsQuorum"] = o.PartitionsQuorum
	}
	if true {
		toSerialize["partitionsAllOk"] = o.PartitionsAllOk
	}
	return json.Marshal(toSerialize)
}

type NullableGetHealth200Response struct {
	value *GetHealth200Response
	isSet bool
}

func (v NullableGetHealth200Response) Get() *GetHealth200Response {
	return v.value
}

func (v *NullableGetHealth200Response) Set(val *GetHealth200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetHealth200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetHealth200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetHealth200Response(val *GetHealth200Response) *NullableGetHealth200Response {
	return &NullableGetHealth200Response{value: val, isSet: true}
}

func (v NullableGetHealth200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetHealth200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


