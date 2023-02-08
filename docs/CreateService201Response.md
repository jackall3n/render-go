# CreateService201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | Pointer to [**Service**](Service.md) |  | [optional] 
**DeployId** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateService201Response

`func NewCreateService201Response() *CreateService201Response`

NewCreateService201Response instantiates a new CreateService201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateService201ResponseWithDefaults

`func NewCreateService201ResponseWithDefaults() *CreateService201Response`

NewCreateService201ResponseWithDefaults instantiates a new CreateService201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *CreateService201Response) GetService() Service`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CreateService201Response) GetServiceOk() (*Service, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CreateService201Response) SetService(v Service)`

SetService sets Service field to given value.

### HasService

`func (o *CreateService201Response) HasService() bool`

HasService returns a boolean if a field has been set.

### GetDeployId

`func (o *CreateService201Response) GetDeployId() string`

GetDeployId returns the DeployId field if non-nil, zero value otherwise.

### GetDeployIdOk

`func (o *CreateService201Response) GetDeployIdOk() (*string, bool)`

GetDeployIdOk returns a tuple with the DeployId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployId

`func (o *CreateService201Response) SetDeployId(v string)`

SetDeployId sets DeployId field to given value.

### HasDeployId

`func (o *CreateService201Response) HasDeployId() bool`

HasDeployId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


