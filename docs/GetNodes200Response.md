# GetNodes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Node** | **string** |  | 
**GarageVersion** | **string** |  | 
**KnownNodes** | [**map[string]NodeNetworkInfo**](NodeNetworkInfo.md) |  | 
**Layout** | [**ClusterLayout**](ClusterLayout.md) |  | 

## Methods

### NewGetNodes200Response

`func NewGetNodes200Response(node string, garageVersion string, knownNodes map[string]NodeNetworkInfo, layout ClusterLayout, ) *GetNodes200Response`

NewGetNodes200Response instantiates a new GetNodes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNodes200ResponseWithDefaults

`func NewGetNodes200ResponseWithDefaults() *GetNodes200Response`

NewGetNodes200ResponseWithDefaults instantiates a new GetNodes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNode

`func (o *GetNodes200Response) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *GetNodes200Response) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *GetNodes200Response) SetNode(v string)`

SetNode sets Node field to given value.


### GetGarageVersion

`func (o *GetNodes200Response) GetGarageVersion() string`

GetGarageVersion returns the GarageVersion field if non-nil, zero value otherwise.

### GetGarageVersionOk

`func (o *GetNodes200Response) GetGarageVersionOk() (*string, bool)`

GetGarageVersionOk returns a tuple with the GarageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGarageVersion

`func (o *GetNodes200Response) SetGarageVersion(v string)`

SetGarageVersion sets GarageVersion field to given value.


### GetKnownNodes

`func (o *GetNodes200Response) GetKnownNodes() map[string]NodeNetworkInfo`

GetKnownNodes returns the KnownNodes field if non-nil, zero value otherwise.

### GetKnownNodesOk

`func (o *GetNodes200Response) GetKnownNodesOk() (*map[string]NodeNetworkInfo, bool)`

GetKnownNodesOk returns a tuple with the KnownNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnownNodes

`func (o *GetNodes200Response) SetKnownNodes(v map[string]NodeNetworkInfo)`

SetKnownNodes sets KnownNodes field to given value.


### GetLayout

`func (o *GetNodes200Response) GetLayout() ClusterLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *GetNodes200Response) GetLayoutOk() (*ClusterLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *GetNodes200Response) SetLayout(v ClusterLayout)`

SetLayout sets Layout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


