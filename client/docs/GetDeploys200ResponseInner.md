# GetDeploys200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deploy** | Pointer to [**Deploy**](Deploy.md) |  | [optional] 
**Cursor** | Pointer to **string** |  | [optional] 

## Methods

### NewGetDeploys200ResponseInner

`func NewGetDeploys200ResponseInner() *GetDeploys200ResponseInner`

NewGetDeploys200ResponseInner instantiates a new GetDeploys200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeploys200ResponseInnerWithDefaults

`func NewGetDeploys200ResponseInnerWithDefaults() *GetDeploys200ResponseInner`

NewGetDeploys200ResponseInnerWithDefaults instantiates a new GetDeploys200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploy

`func (o *GetDeploys200ResponseInner) GetDeploy() Deploy`

GetDeploy returns the Deploy field if non-nil, zero value otherwise.

### GetDeployOk

`func (o *GetDeploys200ResponseInner) GetDeployOk() (*Deploy, bool)`

GetDeployOk returns a tuple with the Deploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploy

`func (o *GetDeploys200ResponseInner) SetDeploy(v Deploy)`

SetDeploy sets Deploy field to given value.

### HasDeploy

`func (o *GetDeploys200ResponseInner) HasDeploy() bool`

HasDeploy returns a boolean if a field has been set.

### GetCursor

`func (o *GetDeploys200ResponseInner) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *GetDeploys200ResponseInner) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *GetDeploys200ResponseInner) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *GetDeploys200ResponseInner) HasCursor() bool`

HasCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


