# GetHealth200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**KnownNodes** | **int64** |  | 
**ConnectedNodes** | **int64** |  | 
**StorageNodes** | **int64** |  | 
**StorageNodesOk** | **int64** |  | 
**Partitions** | **int64** |  | 
**PartitionsQuorum** | **int64** |  | 
**PartitionsAllOk** | **int64** |  | 

## Methods

### NewGetHealth200Response

`func NewGetHealth200Response(status string, knownNodes int64, connectedNodes int64, storageNodes int64, storageNodesOk int64, partitions int64, partitionsQuorum int64, partitionsAllOk int64, ) *GetHealth200Response`

NewGetHealth200Response instantiates a new GetHealth200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetHealth200ResponseWithDefaults

`func NewGetHealth200ResponseWithDefaults() *GetHealth200Response`

NewGetHealth200ResponseWithDefaults instantiates a new GetHealth200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetHealth200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetHealth200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetHealth200Response) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetKnownNodes

`func (o *GetHealth200Response) GetKnownNodes() int64`

GetKnownNodes returns the KnownNodes field if non-nil, zero value otherwise.

### GetKnownNodesOk

`func (o *GetHealth200Response) GetKnownNodesOk() (*int64, bool)`

GetKnownNodesOk returns a tuple with the KnownNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnownNodes

`func (o *GetHealth200Response) SetKnownNodes(v int64)`

SetKnownNodes sets KnownNodes field to given value.


### GetConnectedNodes

`func (o *GetHealth200Response) GetConnectedNodes() int64`

GetConnectedNodes returns the ConnectedNodes field if non-nil, zero value otherwise.

### GetConnectedNodesOk

`func (o *GetHealth200Response) GetConnectedNodesOk() (*int64, bool)`

GetConnectedNodesOk returns a tuple with the ConnectedNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedNodes

`func (o *GetHealth200Response) SetConnectedNodes(v int64)`

SetConnectedNodes sets ConnectedNodes field to given value.


### GetStorageNodes

`func (o *GetHealth200Response) GetStorageNodes() int64`

GetStorageNodes returns the StorageNodes field if non-nil, zero value otherwise.

### GetStorageNodesOk

`func (o *GetHealth200Response) GetStorageNodesOk() (*int64, bool)`

GetStorageNodesOk returns a tuple with the StorageNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageNodes

`func (o *GetHealth200Response) SetStorageNodes(v int64)`

SetStorageNodes sets StorageNodes field to given value.


### GetStorageNodesOk

`func (o *GetHealth200Response) GetStorageNodesOk() int64`

GetStorageNodesOk returns the StorageNodesOk field if non-nil, zero value otherwise.

### GetStorageNodesOkOk

`func (o *GetHealth200Response) GetStorageNodesOkOk() (*int64, bool)`

GetStorageNodesOkOk returns a tuple with the StorageNodesOk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageNodesOk

`func (o *GetHealth200Response) SetStorageNodesOk(v int64)`

SetStorageNodesOk sets StorageNodesOk field to given value.


### GetPartitions

`func (o *GetHealth200Response) GetPartitions() int64`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *GetHealth200Response) GetPartitionsOk() (*int64, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *GetHealth200Response) SetPartitions(v int64)`

SetPartitions sets Partitions field to given value.


### GetPartitionsQuorum

`func (o *GetHealth200Response) GetPartitionsQuorum() int64`

GetPartitionsQuorum returns the PartitionsQuorum field if non-nil, zero value otherwise.

### GetPartitionsQuorumOk

`func (o *GetHealth200Response) GetPartitionsQuorumOk() (*int64, bool)`

GetPartitionsQuorumOk returns a tuple with the PartitionsQuorum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionsQuorum

`func (o *GetHealth200Response) SetPartitionsQuorum(v int64)`

SetPartitionsQuorum sets PartitionsQuorum field to given value.


### GetPartitionsAllOk

`func (o *GetHealth200Response) GetPartitionsAllOk() int64`

GetPartitionsAllOk returns the PartitionsAllOk field if non-nil, zero value otherwise.

### GetPartitionsAllOkOk

`func (o *GetHealth200Response) GetPartitionsAllOkOk() (*int64, bool)`

GetPartitionsAllOkOk returns a tuple with the PartitionsAllOk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionsAllOk

`func (o *GetHealth200Response) SetPartitionsAllOk(v int64)`

SetPartitionsAllOk sets PartitionsAllOk field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


