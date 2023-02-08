# CreateDeployRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClearCache** | Pointer to **string** | Defaults to \&quot;do_not_clear\&quot; | [optional] [default to "do_not_clear"]

## Methods

### NewCreateDeployRequest

`func NewCreateDeployRequest() *CreateDeployRequest`

NewCreateDeployRequest instantiates a new CreateDeployRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDeployRequestWithDefaults

`func NewCreateDeployRequestWithDefaults() *CreateDeployRequest`

NewCreateDeployRequestWithDefaults instantiates a new CreateDeployRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClearCache

`func (o *CreateDeployRequest) GetClearCache() string`

GetClearCache returns the ClearCache field if non-nil, zero value otherwise.

### GetClearCacheOk

`func (o *CreateDeployRequest) GetClearCacheOk() (*string, bool)`

GetClearCacheOk returns a tuple with the ClearCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearCache

`func (o *CreateDeployRequest) SetClearCache(v string)`

SetClearCache sets ClearCache field to given value.

### HasClearCache

`func (o *CreateDeployRequest) HasClearCache() bool`

HasClearCache returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


