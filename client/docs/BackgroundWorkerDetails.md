# BackgroundWorkerDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disk** | Pointer to [**StaticSiteDetailsParentServer**](StaticSiteDetailsParentServer.md) |  | [optional] 
**Env** | Pointer to [**ServiceEnv**](ServiceEnv.md) |  | [optional] 
**EnvSpecificDetails** | Pointer to [**WebServiceDetailsEnvSpecificDetails**](WebServiceDetailsEnvSpecificDetails.md) |  | [optional] 
**NumInstances** | Pointer to **int32** |  | [optional] 
**ParentServer** | Pointer to [**StaticSiteDetailsParentServer**](StaticSiteDetailsParentServer.md) |  | [optional] 
**Plan** | Pointer to **string** |  | [optional] 
**PullRequestPreviewsEnabled** | Pointer to **string** |  | [optional] 
**Region** | Pointer to [**Region**](Region.md) |  | [optional] 

## Methods

### NewBackgroundWorkerDetails

`func NewBackgroundWorkerDetails() *BackgroundWorkerDetails`

NewBackgroundWorkerDetails instantiates a new BackgroundWorkerDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackgroundWorkerDetailsWithDefaults

`func NewBackgroundWorkerDetailsWithDefaults() *BackgroundWorkerDetails`

NewBackgroundWorkerDetailsWithDefaults instantiates a new BackgroundWorkerDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisk

`func (o *BackgroundWorkerDetails) GetDisk() StaticSiteDetailsParentServer`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *BackgroundWorkerDetails) GetDiskOk() (*StaticSiteDetailsParentServer, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *BackgroundWorkerDetails) SetDisk(v StaticSiteDetailsParentServer)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *BackgroundWorkerDetails) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetEnv

`func (o *BackgroundWorkerDetails) GetEnv() ServiceEnv`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *BackgroundWorkerDetails) GetEnvOk() (*ServiceEnv, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *BackgroundWorkerDetails) SetEnv(v ServiceEnv)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *BackgroundWorkerDetails) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetEnvSpecificDetails

`func (o *BackgroundWorkerDetails) GetEnvSpecificDetails() WebServiceDetailsEnvSpecificDetails`

GetEnvSpecificDetails returns the EnvSpecificDetails field if non-nil, zero value otherwise.

### GetEnvSpecificDetailsOk

`func (o *BackgroundWorkerDetails) GetEnvSpecificDetailsOk() (*WebServiceDetailsEnvSpecificDetails, bool)`

GetEnvSpecificDetailsOk returns a tuple with the EnvSpecificDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvSpecificDetails

`func (o *BackgroundWorkerDetails) SetEnvSpecificDetails(v WebServiceDetailsEnvSpecificDetails)`

SetEnvSpecificDetails sets EnvSpecificDetails field to given value.

### HasEnvSpecificDetails

`func (o *BackgroundWorkerDetails) HasEnvSpecificDetails() bool`

HasEnvSpecificDetails returns a boolean if a field has been set.

### GetNumInstances

`func (o *BackgroundWorkerDetails) GetNumInstances() int32`

GetNumInstances returns the NumInstances field if non-nil, zero value otherwise.

### GetNumInstancesOk

`func (o *BackgroundWorkerDetails) GetNumInstancesOk() (*int32, bool)`

GetNumInstancesOk returns a tuple with the NumInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumInstances

`func (o *BackgroundWorkerDetails) SetNumInstances(v int32)`

SetNumInstances sets NumInstances field to given value.

### HasNumInstances

`func (o *BackgroundWorkerDetails) HasNumInstances() bool`

HasNumInstances returns a boolean if a field has been set.

### GetParentServer

`func (o *BackgroundWorkerDetails) GetParentServer() StaticSiteDetailsParentServer`

GetParentServer returns the ParentServer field if non-nil, zero value otherwise.

### GetParentServerOk

`func (o *BackgroundWorkerDetails) GetParentServerOk() (*StaticSiteDetailsParentServer, bool)`

GetParentServerOk returns a tuple with the ParentServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentServer

`func (o *BackgroundWorkerDetails) SetParentServer(v StaticSiteDetailsParentServer)`

SetParentServer sets ParentServer field to given value.

### HasParentServer

`func (o *BackgroundWorkerDetails) HasParentServer() bool`

HasParentServer returns a boolean if a field has been set.

### GetPlan

`func (o *BackgroundWorkerDetails) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *BackgroundWorkerDetails) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *BackgroundWorkerDetails) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *BackgroundWorkerDetails) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetPullRequestPreviewsEnabled

`func (o *BackgroundWorkerDetails) GetPullRequestPreviewsEnabled() string`

GetPullRequestPreviewsEnabled returns the PullRequestPreviewsEnabled field if non-nil, zero value otherwise.

### GetPullRequestPreviewsEnabledOk

`func (o *BackgroundWorkerDetails) GetPullRequestPreviewsEnabledOk() (*string, bool)`

GetPullRequestPreviewsEnabledOk returns a tuple with the PullRequestPreviewsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestPreviewsEnabled

`func (o *BackgroundWorkerDetails) SetPullRequestPreviewsEnabled(v string)`

SetPullRequestPreviewsEnabled sets PullRequestPreviewsEnabled field to given value.

### HasPullRequestPreviewsEnabled

`func (o *BackgroundWorkerDetails) HasPullRequestPreviewsEnabled() bool`

HasPullRequestPreviewsEnabled returns a boolean if a field has been set.

### GetRegion

`func (o *BackgroundWorkerDetails) GetRegion() Region`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *BackgroundWorkerDetails) GetRegionOk() (*Region, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *BackgroundWorkerDetails) SetRegion(v Region)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *BackgroundWorkerDetails) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

