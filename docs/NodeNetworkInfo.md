# NodeNetworkInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Addr** | **string** |  | 
**IsUp** | **bool** |  | 
**LastSeenSecsAgo** | **NullableInt32** |  | 
**Hostname** | **string** |  | 

## Methods

### NewNodeNetworkInfo

`func NewNodeNetworkInfo(addr string, isUp bool, lastSeenSecsAgo NullableInt32, hostname string, ) *NodeNetworkInfo`

NewNodeNetworkInfo instantiates a new NodeNetworkInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeNetworkInfoWithDefaults

`func NewNodeNetworkInfoWithDefaults() *NodeNetworkInfo`

NewNodeNetworkInfoWithDefaults instantiates a new NodeNetworkInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NodeNetworkInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NodeNetworkInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NodeNetworkInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NodeNetworkInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAddr

`func (o *NodeNetworkInfo) GetAddr() string`

GetAddr returns the Addr field if non-nil, zero value otherwise.

### GetAddrOk

`func (o *NodeNetworkInfo) GetAddrOk() (*string, bool)`

GetAddrOk returns a tuple with the Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddr

`func (o *NodeNetworkInfo) SetAddr(v string)`

SetAddr sets Addr field to given value.


### GetIsUp

`func (o *NodeNetworkInfo) GetIsUp() bool`

GetIsUp returns the IsUp field if non-nil, zero value otherwise.

### GetIsUpOk

`func (o *NodeNetworkInfo) GetIsUpOk() (*bool, bool)`

GetIsUpOk returns a tuple with the IsUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUp

`func (o *NodeNetworkInfo) SetIsUp(v bool)`

SetIsUp sets IsUp field to given value.


### GetLastSeenSecsAgo

`func (o *NodeNetworkInfo) GetLastSeenSecsAgo() int32`

GetLastSeenSecsAgo returns the LastSeenSecsAgo field if non-nil, zero value otherwise.

### GetLastSeenSecsAgoOk

`func (o *NodeNetworkInfo) GetLastSeenSecsAgoOk() (*int32, bool)`

GetLastSeenSecsAgoOk returns a tuple with the LastSeenSecsAgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenSecsAgo

`func (o *NodeNetworkInfo) SetLastSeenSecsAgo(v int32)`

SetLastSeenSecsAgo sets LastSeenSecsAgo field to given value.


### SetLastSeenSecsAgoNil

`func (o *NodeNetworkInfo) SetLastSeenSecsAgoNil(b bool)`

 SetLastSeenSecsAgoNil sets the value for LastSeenSecsAgo to be an explicit nil

### UnsetLastSeenSecsAgo
`func (o *NodeNetworkInfo) UnsetLastSeenSecsAgo()`

UnsetLastSeenSecsAgo ensures that no value is present for LastSeenSecsAgo, not even an explicit nil
### GetHostname

`func (o *NodeNetworkInfo) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *NodeNetworkInfo) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *NodeNetworkInfo) SetHostname(v string)`

SetHostname sets Hostname field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


