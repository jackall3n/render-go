# WebServiceDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disk** | Pointer to [**StaticSiteDetailsParentServer**](StaticSiteDetailsParentServer.md) |  | [optional] 
**Env** | Pointer to [**ServiceEnv**](ServiceEnv.md) |  | [optional] 
**EnvSpecificDetails** | Pointer to [**WebServiceDetailsEnvSpecificDetails**](WebServiceDetailsEnvSpecificDetails.md) |  | [optional] 
**HealthCheckPath** | Pointer to **string** |  | [optional] 
**NumInstances** | Pointer to **int32** |  | [optional] 
**OpenPorts** | Pointer to [**[]ServerPort**](ServerPort.md) |  | [optional] 
**ParentServer** | Pointer to [**StaticSiteDetailsParentServer**](StaticSiteDetailsParentServer.md) |  | [optional] 
**Plan** | Pointer to **string** |  | [optional] 
**PullRequestPreviewsEnabled** | Pointer to **string** |  | [optional] 
**Region** | Pointer to [**Region**](Region.md) |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewWebServiceDetails

`func NewWebServiceDetails() *WebServiceDetails`

NewWebServiceDetails instantiates a new WebServiceDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebServiceDetailsWithDefaults

`func NewWebServiceDetailsWithDefaults() *WebServiceDetails`

NewWebServiceDetailsWithDefaults instantiates a new WebServiceDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisk

`func (o *WebServiceDetails) GetDisk() StaticSiteDetailsParentServer`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *WebServiceDetails) GetDiskOk() (*StaticSiteDetailsParentServer, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *WebServiceDetails) SetDisk(v StaticSiteDetailsParentServer)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *WebServiceDetails) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetEnv

`func (o *WebServiceDetails) GetEnv() ServiceEnv`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *WebServiceDetails) GetEnvOk() (*ServiceEnv, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *WebServiceDetails) SetEnv(v ServiceEnv)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *WebServiceDetails) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetEnvSpecificDetails

`func (o *WebServiceDetails) GetEnvSpecificDetails() WebServiceDetailsEnvSpecificDetails`

GetEnvSpecificDetails returns the EnvSpecificDetails field if non-nil, zero value otherwise.

### GetEnvSpecificDetailsOk

`func (o *WebServiceDetails) GetEnvSpecificDetailsOk() (*WebServiceDetailsEnvSpecificDetails, bool)`

GetEnvSpecificDetailsOk returns a tuple with the EnvSpecificDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvSpecificDetails

`func (o *WebServiceDetails) SetEnvSpecificDetails(v WebServiceDetailsEnvSpecificDetails)`

SetEnvSpecificDetails sets EnvSpecificDetails field to given value.

### HasEnvSpecificDetails

`func (o *WebServiceDetails) HasEnvSpecificDetails() bool`

HasEnvSpecificDetails returns a boolean if a field has been set.

### GetHealthCheckPath

`func (o *WebServiceDetails) GetHealthCheckPath() string`

GetHealthCheckPath returns the HealthCheckPath field if non-nil, zero value otherwise.

### GetHealthCheckPathOk

`func (o *WebServiceDetails) GetHealthCheckPathOk() (*string, bool)`

GetHealthCheckPathOk returns a tuple with the HealthCheckPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckPath

`func (o *WebServiceDetails) SetHealthCheckPath(v string)`

SetHealthCheckPath sets HealthCheckPath field to given value.

### HasHealthCheckPath

`func (o *WebServiceDetails) HasHealthCheckPath() bool`

HasHealthCheckPath returns a boolean if a field has been set.

### GetNumInstances

`func (o *WebServiceDetails) GetNumInstances() int32`

GetNumInstances returns the NumInstances field if non-nil, zero value otherwise.

### GetNumInstancesOk

`func (o *WebServiceDetails) GetNumInstancesOk() (*int32, bool)`

GetNumInstancesOk returns a tuple with the NumInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumInstances

`func (o *WebServiceDetails) SetNumInstances(v int32)`

SetNumInstances sets NumInstances field to given value.

### HasNumInstances

`func (o *WebServiceDetails) HasNumInstances() bool`

HasNumInstances returns a boolean if a field has been set.

### GetOpenPorts

`func (o *WebServiceDetails) GetOpenPorts() []ServerPort`

GetOpenPorts returns the OpenPorts field if non-nil, zero value otherwise.

### GetOpenPortsOk

`func (o *WebServiceDetails) GetOpenPortsOk() (*[]ServerPort, bool)`

GetOpenPortsOk returns a tuple with the OpenPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenPorts

`func (o *WebServiceDetails) SetOpenPorts(v []ServerPort)`

SetOpenPorts sets OpenPorts field to given value.

### HasOpenPorts

`func (o *WebServiceDetails) HasOpenPorts() bool`

HasOpenPorts returns a boolean if a field has been set.

### GetParentServer

`func (o *WebServiceDetails) GetParentServer() StaticSiteDetailsParentServer`

GetParentServer returns the ParentServer field if non-nil, zero value otherwise.

### GetParentServerOk

`func (o *WebServiceDetails) GetParentServerOk() (*StaticSiteDetailsParentServer, bool)`

GetParentServerOk returns a tuple with the ParentServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentServer

`func (o *WebServiceDetails) SetParentServer(v StaticSiteDetailsParentServer)`

SetParentServer sets ParentServer field to given value.

### HasParentServer

`func (o *WebServiceDetails) HasParentServer() bool`

HasParentServer returns a boolean if a field has been set.

### GetPlan

`func (o *WebServiceDetails) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *WebServiceDetails) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *WebServiceDetails) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *WebServiceDetails) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetPullRequestPreviewsEnabled

`func (o *WebServiceDetails) GetPullRequestPreviewsEnabled() string`

GetPullRequestPreviewsEnabled returns the PullRequestPreviewsEnabled field if non-nil, zero value otherwise.

### GetPullRequestPreviewsEnabledOk

`func (o *WebServiceDetails) GetPullRequestPreviewsEnabledOk() (*string, bool)`

GetPullRequestPreviewsEnabledOk returns a tuple with the PullRequestPreviewsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestPreviewsEnabled

`func (o *WebServiceDetails) SetPullRequestPreviewsEnabled(v string)`

SetPullRequestPreviewsEnabled sets PullRequestPreviewsEnabled field to given value.

### HasPullRequestPreviewsEnabled

`func (o *WebServiceDetails) HasPullRequestPreviewsEnabled() bool`

HasPullRequestPreviewsEnabled returns a boolean if a field has been set.

### GetRegion

`func (o *WebServiceDetails) GetRegion() Region`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *WebServiceDetails) GetRegionOk() (*Region, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *WebServiceDetails) SetRegion(v Region)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *WebServiceDetails) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetUrl

`func (o *WebServiceDetails) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebServiceDetails) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebServiceDetails) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WebServiceDetails) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


