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

// UpdateKeyRequestDeny struct for UpdateKeyRequestDeny
type UpdateKeyRequestDeny struct {
	CreateBucket *bool `json:"createBucket,omitempty"`
}

// NewUpdateKeyRequestDeny instantiates a new UpdateKeyRequestDeny object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateKeyRequestDeny() *UpdateKeyRequestDeny {
	this := UpdateKeyRequestDeny{}
	return &this
}

// NewUpdateKeyRequestDenyWithDefaults instantiates a new UpdateKeyRequestDeny object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateKeyRequestDenyWithDefaults() *UpdateKeyRequestDeny {
	this := UpdateKeyRequestDeny{}
	return &this
}

// GetCreateBucket returns the CreateBucket field value if set, zero value otherwise.
func (o *UpdateKeyRequestDeny) GetCreateBucket() bool {
	if o == nil || o.CreateBucket == nil {
		var ret bool
		return ret
	}
	return *o.CreateBucket
}

// GetCreateBucketOk returns a tuple with the CreateBucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateKeyRequestDeny) GetCreateBucketOk() (*bool, bool) {
	if o == nil || o.CreateBucket == nil {
		return nil, false
	}
	return o.CreateBucket, true
}

// HasCreateBucket returns a boolean if a field has been set.
func (o *UpdateKeyRequestDeny) HasCreateBucket() bool {
	if o != nil && o.CreateBucket != nil {
		return true
	}

	return false
}

// SetCreateBucket gets a reference to the given bool and assigns it to the CreateBucket field.
func (o *UpdateKeyRequestDeny) SetCreateBucket(v bool) {
	o.CreateBucket = &v
}

func (o UpdateKeyRequestDeny) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreateBucket != nil {
		toSerialize["createBucket"] = o.CreateBucket
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateKeyRequestDeny struct {
	value *UpdateKeyRequestDeny
	isSet bool
}

func (v NullableUpdateKeyRequestDeny) Get() *UpdateKeyRequestDeny {
	return v.value
}

func (v *NullableUpdateKeyRequestDeny) Set(val *UpdateKeyRequestDeny) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateKeyRequestDeny) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateKeyRequestDeny) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateKeyRequestDeny(val *UpdateKeyRequestDeny) *NullableUpdateKeyRequestDeny {
	return &NullableUpdateKeyRequestDeny{value: val, isSet: true}
}

func (v NullableUpdateKeyRequestDeny) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateKeyRequestDeny) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


