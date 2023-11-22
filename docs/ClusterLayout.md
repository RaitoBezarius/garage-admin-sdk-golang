# ClusterLayout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **int32** |  | 
**Roles** | [**[]NodeClusterInfo**](NodeClusterInfo.md) |  | 
**StagedRoleChanges** | [**[]NodeRoleChange**](NodeRoleChange.md) |  | 

## Methods

### NewClusterLayout

`func NewClusterLayout(version int32, roles []NodeClusterInfo, stagedRoleChanges []NodeRoleChange, ) *ClusterLayout`

NewClusterLayout instantiates a new ClusterLayout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterLayoutWithDefaults

`func NewClusterLayoutWithDefaults() *ClusterLayout`

NewClusterLayoutWithDefaults instantiates a new ClusterLayout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *ClusterLayout) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ClusterLayout) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ClusterLayout) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetRoles

`func (o *ClusterLayout) GetRoles() []NodeClusterInfo`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ClusterLayout) GetRolesOk() (*[]NodeClusterInfo, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ClusterLayout) SetRoles(v []NodeClusterInfo)`

SetRoles sets Roles field to given value.


### GetStagedRoleChanges

`func (o *ClusterLayout) GetStagedRoleChanges() []NodeRoleChange`

GetStagedRoleChanges returns the StagedRoleChanges field if non-nil, zero value otherwise.

### GetStagedRoleChangesOk

`func (o *ClusterLayout) GetStagedRoleChangesOk() (*[]NodeRoleChange, bool)`

GetStagedRoleChangesOk returns a tuple with the StagedRoleChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStagedRoleChanges

`func (o *ClusterLayout) SetStagedRoleChanges(v []NodeRoleChange)`

SetStagedRoleChanges sets StagedRoleChanges field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


