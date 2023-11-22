# ImportKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** |  | 
**AccessKeyId** | **string** |  | 
**SecretAccessKey** | **string** |  | 

## Methods

### NewImportKeyRequest

`func NewImportKeyRequest(name NullableString, accessKeyId string, secretAccessKey string, ) *ImportKeyRequest`

NewImportKeyRequest instantiates a new ImportKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportKeyRequestWithDefaults

`func NewImportKeyRequestWithDefaults() *ImportKeyRequest`

NewImportKeyRequestWithDefaults instantiates a new ImportKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ImportKeyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImportKeyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImportKeyRequest) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *ImportKeyRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ImportKeyRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAccessKeyId

`func (o *ImportKeyRequest) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *ImportKeyRequest) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *ImportKeyRequest) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.


### GetSecretAccessKey

`func (o *ImportKeyRequest) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *ImportKeyRequest) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *ImportKeyRequest) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


