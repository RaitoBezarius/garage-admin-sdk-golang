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

// ListBuckets200ResponseInnerLocalAliasesInner struct for ListBuckets200ResponseInnerLocalAliasesInner
type ListBuckets200ResponseInnerLocalAliasesInner struct {
	Alias string `json:"alias"`
	AccessKeyId string `json:"accessKeyId"`
}

// NewListBuckets200ResponseInnerLocalAliasesInner instantiates a new ListBuckets200ResponseInnerLocalAliasesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListBuckets200ResponseInnerLocalAliasesInner(alias string, accessKeyId string) *ListBuckets200ResponseInnerLocalAliasesInner {
	this := ListBuckets200ResponseInnerLocalAliasesInner{}
	this.Alias = alias
	this.AccessKeyId = accessKeyId
	return &this
}

// NewListBuckets200ResponseInnerLocalAliasesInnerWithDefaults instantiates a new ListBuckets200ResponseInnerLocalAliasesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListBuckets200ResponseInnerLocalAliasesInnerWithDefaults() *ListBuckets200ResponseInnerLocalAliasesInner {
	this := ListBuckets200ResponseInnerLocalAliasesInner{}
	return &this
}

// GetAlias returns the Alias field value
func (o *ListBuckets200ResponseInnerLocalAliasesInner) GetAlias() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Alias
}

// GetAliasOk returns a tuple with the Alias field value
// and a boolean to check if the value has been set.
func (o *ListBuckets200ResponseInnerLocalAliasesInner) GetAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Alias, true
}

// SetAlias sets field value
func (o *ListBuckets200ResponseInnerLocalAliasesInner) SetAlias(v string) {
	o.Alias = v
}

// GetAccessKeyId returns the AccessKeyId field value
func (o *ListBuckets200ResponseInnerLocalAliasesInner) GetAccessKeyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessKeyId
}

// GetAccessKeyIdOk returns a tuple with the AccessKeyId field value
// and a boolean to check if the value has been set.
func (o *ListBuckets200ResponseInnerLocalAliasesInner) GetAccessKeyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessKeyId, true
}

// SetAccessKeyId sets field value
func (o *ListBuckets200ResponseInnerLocalAliasesInner) SetAccessKeyId(v string) {
	o.AccessKeyId = v
}

func (o ListBuckets200ResponseInnerLocalAliasesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["alias"] = o.Alias
	}
	if true {
		toSerialize["accessKeyId"] = o.AccessKeyId
	}
	return json.Marshal(toSerialize)
}

type NullableListBuckets200ResponseInnerLocalAliasesInner struct {
	value *ListBuckets200ResponseInnerLocalAliasesInner
	isSet bool
}

func (v NullableListBuckets200ResponseInnerLocalAliasesInner) Get() *ListBuckets200ResponseInnerLocalAliasesInner {
	return v.value
}

func (v *NullableListBuckets200ResponseInnerLocalAliasesInner) Set(val *ListBuckets200ResponseInnerLocalAliasesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListBuckets200ResponseInnerLocalAliasesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListBuckets200ResponseInnerLocalAliasesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListBuckets200ResponseInnerLocalAliasesInner(val *ListBuckets200ResponseInnerLocalAliasesInner) *NullableListBuckets200ResponseInnerLocalAliasesInner {
	return &NullableListBuckets200ResponseInnerLocalAliasesInner{value: val, isSet: true}
}

func (v NullableListBuckets200ResponseInnerLocalAliasesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListBuckets200ResponseInnerLocalAliasesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


