/*
Garage Administration API v0+garage-v0.8.0

Administrate your Garage cluster programatically, including status, layout, keys, buckets, and maintainance tasks.  *Disclaimer: The API is not stable yet, hence its v0 tag. The API can change at any time, and changes can include breaking backward compatibility. Read the changelog and upgrade your scripts before upgrading. Additionnaly, this specification is very early stage and can contain bugs, especially on error return codes/types that are not tested yet. Do not expect a well finished and polished product!* 

API version: v0.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package garage

import (
	"encoding/json"
)

// AllowBucketKeyRequestPermissions struct for AllowBucketKeyRequestPermissions
type AllowBucketKeyRequestPermissions struct {
	Read bool `json:"read"`
	Write bool `json:"write"`
	Owner bool `json:"owner"`
}

// NewAllowBucketKeyRequestPermissions instantiates a new AllowBucketKeyRequestPermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllowBucketKeyRequestPermissions(read bool, write bool, owner bool) *AllowBucketKeyRequestPermissions {
	this := AllowBucketKeyRequestPermissions{}
	this.Read = read
	this.Write = write
	this.Owner = owner
	return &this
}

// NewAllowBucketKeyRequestPermissionsWithDefaults instantiates a new AllowBucketKeyRequestPermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllowBucketKeyRequestPermissionsWithDefaults() *AllowBucketKeyRequestPermissions {
	this := AllowBucketKeyRequestPermissions{}
	return &this
}

// GetRead returns the Read field value
func (o *AllowBucketKeyRequestPermissions) GetRead() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Read
}

// GetReadOk returns a tuple with the Read field value
// and a boolean to check if the value has been set.
func (o *AllowBucketKeyRequestPermissions) GetReadOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Read, true
}

// SetRead sets field value
func (o *AllowBucketKeyRequestPermissions) SetRead(v bool) {
	o.Read = v
}

// GetWrite returns the Write field value
func (o *AllowBucketKeyRequestPermissions) GetWrite() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Write
}

// GetWriteOk returns a tuple with the Write field value
// and a boolean to check if the value has been set.
func (o *AllowBucketKeyRequestPermissions) GetWriteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Write, true
}

// SetWrite sets field value
func (o *AllowBucketKeyRequestPermissions) SetWrite(v bool) {
	o.Write = v
}

// GetOwner returns the Owner field value
func (o *AllowBucketKeyRequestPermissions) GetOwner() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *AllowBucketKeyRequestPermissions) GetOwnerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value
func (o *AllowBucketKeyRequestPermissions) SetOwner(v bool) {
	o.Owner = v
}

func (o AllowBucketKeyRequestPermissions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["read"] = o.Read
	}
	if true {
		toSerialize["write"] = o.Write
	}
	if true {
		toSerialize["owner"] = o.Owner
	}
	return json.Marshal(toSerialize)
}

type NullableAllowBucketKeyRequestPermissions struct {
	value *AllowBucketKeyRequestPermissions
	isSet bool
}

func (v NullableAllowBucketKeyRequestPermissions) Get() *AllowBucketKeyRequestPermissions {
	return v.value
}

func (v *NullableAllowBucketKeyRequestPermissions) Set(val *AllowBucketKeyRequestPermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableAllowBucketKeyRequestPermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableAllowBucketKeyRequestPermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllowBucketKeyRequestPermissions(val *AllowBucketKeyRequestPermissions) *NullableAllowBucketKeyRequestPermissions {
	return &NullableAllowBucketKeyRequestPermissions{value: val, isSet: true}
}

func (v NullableAllowBucketKeyRequestPermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllowBucketKeyRequestPermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


