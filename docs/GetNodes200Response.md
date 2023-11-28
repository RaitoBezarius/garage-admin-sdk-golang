# GetNodes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Node** | **string** |  | 
**GarageVersion** | **string** |  | 
**GarageFeatures** | **[]string** |  | 
**RustVersion** | **string** |  | 
**DbEngine** | **string** |  | 
**KnownNodes** | [**[]NodeNetworkInfo**](NodeNetworkInfo.md) |  | 
**Layout** | [**ClusterLayout**](ClusterLayout.md) |  | 

## Methods

### NewGetNodes200Response

`func NewGetNodes200Response(node string, garageVersion string, garageFeatures []string, rustVersion string, dbEngine string, knownNodes []NodeNetworkInfo, layout ClusterLayout, ) *GetNodes200Response`

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


### GetGarageFeatures

`func (o *GetNodes200Response) GetGarageFeatures() []string`

GetGarageFeatures returns the GarageFeatures field if non-nil, zero value otherwise.

### GetGarageFeaturesOk

`func (o *GetNodes200Response) GetGarageFeaturesOk() (*[]string, bool)`

GetGarageFeaturesOk returns a tuple with the GarageFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGarageFeatures

`func (o *GetNodes200Response) SetGarageFeatures(v []string)`

SetGarageFeatures sets GarageFeatures field to given value.


### GetRustVersion

`func (o *GetNodes200Response) GetRustVersion() string`

GetRustVersion returns the RustVersion field if non-nil, zero value otherwise.

### GetRustVersionOk

`func (o *GetNodes200Response) GetRustVersionOk() (*string, bool)`

GetRustVersionOk returns a tuple with the RustVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRustVersion

`func (o *GetNodes200Response) SetRustVersion(v string)`

SetRustVersion sets RustVersion field to given value.


### GetDbEngine

`func (o *GetNodes200Response) GetDbEngine() string`

GetDbEngine returns the DbEngine field if non-nil, zero value otherwise.

### GetDbEngineOk

`func (o *GetNodes200Response) GetDbEngineOk() (*string, bool)`

GetDbEngineOk returns a tuple with the DbEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbEngine

`func (o *GetNodes200Response) SetDbEngine(v string)`

SetDbEngine sets DbEngine field to given value.


### GetKnownNodes

`func (o *GetNodes200Response) GetKnownNodes() []NodeNetworkInfo`

GetKnownNodes returns the KnownNodes field if non-nil, zero value otherwise.

### GetKnownNodesOk

`func (o *GetNodes200Response) GetKnownNodesOk() (*[]NodeNetworkInfo, bool)`

GetKnownNodesOk returns a tuple with the KnownNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnownNodes

`func (o *GetNodes200Response) SetKnownNodes(v []NodeNetworkInfo)`

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


