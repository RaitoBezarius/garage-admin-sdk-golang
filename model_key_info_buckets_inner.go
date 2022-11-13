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

// KeyInfoBucketsInner struct for KeyInfoBucketsInner
type KeyInfoBucketsInner struct {
	Id *string `json:"id,omitempty"`
	GlobalAliases []string `json:"globalAliases,omitempty"`
	LocalAliases []string `json:"localAliases,omitempty"`
	Permissions *KeyInfoBucketsInnerPermissions `json:"permissions,omitempty"`
}

// NewKeyInfoBucketsInner instantiates a new KeyInfoBucketsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyInfoBucketsInner() *KeyInfoBucketsInner {
	this := KeyInfoBucketsInner{}
	return &this
}

// NewKeyInfoBucketsInnerWithDefaults instantiates a new KeyInfoBucketsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyInfoBucketsInnerWithDefaults() *KeyInfoBucketsInner {
	this := KeyInfoBucketsInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KeyInfoBucketsInner) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyInfoBucketsInner) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KeyInfoBucketsInner) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *KeyInfoBucketsInner) SetId(v string) {
	o.Id = &v
}

// GetGlobalAliases returns the GlobalAliases field value if set, zero value otherwise.
func (o *KeyInfoBucketsInner) GetGlobalAliases() []string {
	if o == nil || o.GlobalAliases == nil {
		var ret []string
		return ret
	}
	return o.GlobalAliases
}

// GetGlobalAliasesOk returns a tuple with the GlobalAliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyInfoBucketsInner) GetGlobalAliasesOk() ([]string, bool) {
	if o == nil || o.GlobalAliases == nil {
		return nil, false
	}
	return o.GlobalAliases, true
}

// HasGlobalAliases returns a boolean if a field has been set.
func (o *KeyInfoBucketsInner) HasGlobalAliases() bool {
	if o != nil && o.GlobalAliases != nil {
		return true
	}

	return false
}

// SetGlobalAliases gets a reference to the given []string and assigns it to the GlobalAliases field.
func (o *KeyInfoBucketsInner) SetGlobalAliases(v []string) {
	o.GlobalAliases = v
}

// GetLocalAliases returns the LocalAliases field value if set, zero value otherwise.
func (o *KeyInfoBucketsInner) GetLocalAliases() []string {
	if o == nil || o.LocalAliases == nil {
		var ret []string
		return ret
	}
	return o.LocalAliases
}

// GetLocalAliasesOk returns a tuple with the LocalAliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyInfoBucketsInner) GetLocalAliasesOk() ([]string, bool) {
	if o == nil || o.LocalAliases == nil {
		return nil, false
	}
	return o.LocalAliases, true
}

// HasLocalAliases returns a boolean if a field has been set.
func (o *KeyInfoBucketsInner) HasLocalAliases() bool {
	if o != nil && o.LocalAliases != nil {
		return true
	}

	return false
}

// SetLocalAliases gets a reference to the given []string and assigns it to the LocalAliases field.
func (o *KeyInfoBucketsInner) SetLocalAliases(v []string) {
	o.LocalAliases = v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *KeyInfoBucketsInner) GetPermissions() KeyInfoBucketsInnerPermissions {
	if o == nil || o.Permissions == nil {
		var ret KeyInfoBucketsInnerPermissions
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyInfoBucketsInner) GetPermissionsOk() (*KeyInfoBucketsInnerPermissions, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *KeyInfoBucketsInner) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given KeyInfoBucketsInnerPermissions and assigns it to the Permissions field.
func (o *KeyInfoBucketsInner) SetPermissions(v KeyInfoBucketsInnerPermissions) {
	o.Permissions = &v
}

func (o KeyInfoBucketsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.GlobalAliases != nil {
		toSerialize["globalAliases"] = o.GlobalAliases
	}
	if o.LocalAliases != nil {
		toSerialize["localAliases"] = o.LocalAliases
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	return json.Marshal(toSerialize)
}

type NullableKeyInfoBucketsInner struct {
	value *KeyInfoBucketsInner
	isSet bool
}

func (v NullableKeyInfoBucketsInner) Get() *KeyInfoBucketsInner {
	return v.value
}

func (v *NullableKeyInfoBucketsInner) Set(val *KeyInfoBucketsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyInfoBucketsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyInfoBucketsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyInfoBucketsInner(val *KeyInfoBucketsInner) *NullableKeyInfoBucketsInner {
	return &NullableKeyInfoBucketsInner{value: val, isSet: true}
}

func (v NullableKeyInfoBucketsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyInfoBucketsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


