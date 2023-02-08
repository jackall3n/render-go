# PrivateServiceDetailsPOST

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disk** | Pointer to [**WebServiceDetailsPOSTDisk**](WebServiceDetailsPOSTDisk.md) |  | [optional] 
**Env** | [**ServiceEnv**](ServiceEnv.md) |  | 
**EnvSpecificDetails** | Pointer to [**WebServiceDetailsPOSTEnvSpecificDetails**](WebServiceDetailsPOSTEnvSpecificDetails.md) |  | [optional] 
**NumInstances** | Pointer to **int32** | Defaults to 1 | [optional] [default to 1]
**Plan** | Pointer to **string** | Defaults to \&quot;starter\&quot; | [optional] [default to "starter"]
**PullRequestPreviewsEnabled** | Pointer to **string** | Defaults to \&quot;no\&quot; | [optional] [default to "no"]
**Region** | Pointer to [**Region**](Region.md) |  | [optional] 

## Methods

### NewPrivateServiceDetailsPOST

`func NewPrivateServiceDetailsPOST(env ServiceEnv, ) *PrivateServiceDetailsPOST`

NewPrivateServiceDetailsPOST instantiates a new PrivateServiceDetailsPOST object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateServiceDetailsPOSTWithDefaults

`func NewPrivateServiceDetailsPOSTWithDefaults() *PrivateServiceDetailsPOST`

NewPrivateServiceDetailsPOSTWithDefaults instantiates a new PrivateServiceDetailsPOST object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisk

`func (o *PrivateServiceDetailsPOST) GetDisk() WebServiceDetailsPOSTDisk`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *PrivateServiceDetailsPOST) GetDiskOk() (*WebServiceDetailsPOSTDisk, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *PrivateServiceDetailsPOST) SetDisk(v WebServiceDetailsPOSTDisk)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *PrivateServiceDetailsPOST) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetEnv

`func (o *PrivateServiceDetailsPOST) GetEnv() ServiceEnv`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *PrivateServiceDetailsPOST) GetEnvOk() (*ServiceEnv, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *PrivateServiceDetailsPOST) SetEnv(v ServiceEnv)`

SetEnv sets Env field to given value.


### GetEnvSpecificDetails

`func (o *PrivateServiceDetailsPOST) GetEnvSpecificDetails() WebServiceDetailsPOSTEnvSpecificDetails`

GetEnvSpecificDetails returns the EnvSpecificDetails field if non-nil, zero value otherwise.

### GetEnvSpecificDetailsOk

`func (o *PrivateServiceDetailsPOST) GetEnvSpecificDetailsOk() (*WebServiceDetailsPOSTEnvSpecificDetails, bool)`

GetEnvSpecificDetailsOk returns a tuple with the EnvSpecificDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvSpecificDetails

`func (o *PrivateServiceDetailsPOST) SetEnvSpecificDetails(v WebServiceDetailsPOSTEnvSpecificDetails)`

SetEnvSpecificDetails sets EnvSpecificDetails field to given value.

### HasEnvSpecificDetails

`func (o *PrivateServiceDetailsPOST) HasEnvSpecificDetails() bool`

HasEnvSpecificDetails returns a boolean if a field has been set.

### GetNumInstances

`func (o *PrivateServiceDetailsPOST) GetNumInstances() int32`

GetNumInstances returns the NumInstances field if non-nil, zero value otherwise.

### GetNumInstancesOk

`func (o *PrivateServiceDetailsPOST) GetNumInstancesOk() (*int32, bool)`

GetNumInstancesOk returns a tuple with the NumInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumInstances

`func (o *PrivateServiceDetailsPOST) SetNumInstances(v int32)`

SetNumInstances sets NumInstances field to given value.

### HasNumInstances

`func (o *PrivateServiceDetailsPOST) HasNumInstances() bool`

HasNumInstances returns a boolean if a field has been set.

### GetPlan

`func (o *PrivateServiceDetailsPOST) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *PrivateServiceDetailsPOST) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *PrivateServiceDetailsPOST) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *PrivateServiceDetailsPOST) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetPullRequestPreviewsEnabled

`func (o *PrivateServiceDetailsPOST) GetPullRequestPreviewsEnabled() string`

GetPullRequestPreviewsEnabled returns the PullRequestPreviewsEnabled field if non-nil, zero value otherwise.

### GetPullRequestPreviewsEnabledOk

`func (o *PrivateServiceDetailsPOST) GetPullRequestPreviewsEnabledOk() (*string, bool)`

GetPullRequestPreviewsEnabledOk returns a tuple with the PullRequestPreviewsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestPreviewsEnabled

`func (o *PrivateServiceDetailsPOST) SetPullRequestPreviewsEnabled(v string)`

SetPullRequestPreviewsEnabled sets PullRequestPreviewsEnabled field to given value.

### HasPullRequestPreviewsEnabled

`func (o *PrivateServiceDetailsPOST) HasPullRequestPreviewsEnabled() bool`

HasPullRequestPreviewsEnabled returns a boolean if a field has been set.

### GetRegion

`func (o *PrivateServiceDetailsPOST) GetRegion() Region`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PrivateServiceDetailsPOST) GetRegionOk() (*Region, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PrivateServiceDetailsPOST) SetRegion(v Region)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *PrivateServiceDetailsPOST) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


