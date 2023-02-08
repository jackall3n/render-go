# ServiceServiceDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildCommand** | Pointer to **string** |  | [optional] 
**ParentServer** | Pointer to [**StaticSiteDetailsParentServer**](StaticSiteDetailsParentServer.md) |  | [optional] 
**PublishPath** | Pointer to **string** |  | [optional] 
**PullRequestPreviewsEnabled** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Disk** | Pointer to [**StaticSiteDetailsParentServer**](StaticSiteDetailsParentServer.md) |  | [optional] 
**Env** | Pointer to [**ServiceEnv**](ServiceEnv.md) |  | [optional] 
**EnvSpecificDetails** | Pointer to [**WebServiceDetailsEnvSpecificDetails**](WebServiceDetailsEnvSpecificDetails.md) |  | [optional] 
**HealthCheckPath** | Pointer to **string** |  | [optional] 
**NumInstances** | Pointer to **int32** |  | [optional] 
**OpenPorts** | Pointer to [**[]ServerPort**](ServerPort.md) |  | [optional] 
**Plan** | Pointer to **string** |  | [optional] 
**Region** | Pointer to [**Region**](Region.md) |  | [optional] 
**LastSuccessfulRunAt** | Pointer to **time.Time** |  | [optional] 
**Schedule** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceServiceDetails

`func NewServiceServiceDetails() *ServiceServiceDetails`

NewServiceServiceDetails instantiates a new ServiceServiceDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceServiceDetailsWithDefaults

`func NewServiceServiceDetailsWithDefaults() *ServiceServiceDetails`

NewServiceServiceDetailsWithDefaults instantiates a new ServiceServiceDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildCommand

`func (o *ServiceServiceDetails) GetBuildCommand() string`

GetBuildCommand returns the BuildCommand field if non-nil, zero value otherwise.

### GetBuildCommandOk

`func (o *ServiceServiceDetails) GetBuildCommandOk() (*string, bool)`

GetBuildCommandOk returns a tuple with the BuildCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildCommand

`func (o *ServiceServiceDetails) SetBuildCommand(v string)`

SetBuildCommand sets BuildCommand field to given value.

### HasBuildCommand

`func (o *ServiceServiceDetails) HasBuildCommand() bool`

HasBuildCommand returns a boolean if a field has been set.

### GetParentServer

`func (o *ServiceServiceDetails) GetParentServer() StaticSiteDetailsParentServer`

GetParentServer returns the ParentServer field if non-nil, zero value otherwise.

### GetParentServerOk

`func (o *ServiceServiceDetails) GetParentServerOk() (*StaticSiteDetailsParentServer, bool)`

GetParentServerOk returns a tuple with the ParentServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentServer

`func (o *ServiceServiceDetails) SetParentServer(v StaticSiteDetailsParentServer)`

SetParentServer sets ParentServer field to given value.

### HasParentServer

`func (o *ServiceServiceDetails) HasParentServer() bool`

HasParentServer returns a boolean if a field has been set.

### GetPublishPath

`func (o *ServiceServiceDetails) GetPublishPath() string`

GetPublishPath returns the PublishPath field if non-nil, zero value otherwise.

### GetPublishPathOk

`func (o *ServiceServiceDetails) GetPublishPathOk() (*string, bool)`

GetPublishPathOk returns a tuple with the PublishPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishPath

`func (o *ServiceServiceDetails) SetPublishPath(v string)`

SetPublishPath sets PublishPath field to given value.

### HasPublishPath

`func (o *ServiceServiceDetails) HasPublishPath() bool`

HasPublishPath returns a boolean if a field has been set.

### GetPullRequestPreviewsEnabled

`func (o *ServiceServiceDetails) GetPullRequestPreviewsEnabled() string`

GetPullRequestPreviewsEnabled returns the PullRequestPreviewsEnabled field if non-nil, zero value otherwise.

### GetPullRequestPreviewsEnabledOk

`func (o *ServiceServiceDetails) GetPullRequestPreviewsEnabledOk() (*string, bool)`

GetPullRequestPreviewsEnabledOk returns a tuple with the PullRequestPreviewsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestPreviewsEnabled

`func (o *ServiceServiceDetails) SetPullRequestPreviewsEnabled(v string)`

SetPullRequestPreviewsEnabled sets PullRequestPreviewsEnabled field to given value.

### HasPullRequestPreviewsEnabled

`func (o *ServiceServiceDetails) HasPullRequestPreviewsEnabled() bool`

HasPullRequestPreviewsEnabled returns a boolean if a field has been set.

### GetUrl

`func (o *ServiceServiceDetails) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ServiceServiceDetails) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ServiceServiceDetails) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ServiceServiceDetails) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisk

`func (o *ServiceServiceDetails) GetDisk() StaticSiteDetailsParentServer`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *ServiceServiceDetails) GetDiskOk() (*StaticSiteDetailsParentServer, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *ServiceServiceDetails) SetDisk(v StaticSiteDetailsParentServer)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *ServiceServiceDetails) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetEnv

`func (o *ServiceServiceDetails) GetEnv() ServiceEnv`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *ServiceServiceDetails) GetEnvOk() (*ServiceEnv, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *ServiceServiceDetails) SetEnv(v ServiceEnv)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *ServiceServiceDetails) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetEnvSpecificDetails

`func (o *ServiceServiceDetails) GetEnvSpecificDetails() WebServiceDetailsEnvSpecificDetails`

GetEnvSpecificDetails returns the EnvSpecificDetails field if non-nil, zero value otherwise.

### GetEnvSpecificDetailsOk

`func (o *ServiceServiceDetails) GetEnvSpecificDetailsOk() (*WebServiceDetailsEnvSpecificDetails, bool)`

GetEnvSpecificDetailsOk returns a tuple with the EnvSpecificDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvSpecificDetails

`func (o *ServiceServiceDetails) SetEnvSpecificDetails(v WebServiceDetailsEnvSpecificDetails)`

SetEnvSpecificDetails sets EnvSpecificDetails field to given value.

### HasEnvSpecificDetails

`func (o *ServiceServiceDetails) HasEnvSpecificDetails() bool`

HasEnvSpecificDetails returns a boolean if a field has been set.

### GetHealthCheckPath

`func (o *ServiceServiceDetails) GetHealthCheckPath() string`

GetHealthCheckPath returns the HealthCheckPath field if non-nil, zero value otherwise.

### GetHealthCheckPathOk

`func (o *ServiceServiceDetails) GetHealthCheckPathOk() (*string, bool)`

GetHealthCheckPathOk returns a tuple with the HealthCheckPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckPath

`func (o *ServiceServiceDetails) SetHealthCheckPath(v string)`

SetHealthCheckPath sets HealthCheckPath field to given value.

### HasHealthCheckPath

`func (o *ServiceServiceDetails) HasHealthCheckPath() bool`

HasHealthCheckPath returns a boolean if a field has been set.

### GetNumInstances

`func (o *ServiceServiceDetails) GetNumInstances() int32`

GetNumInstances returns the NumInstances field if non-nil, zero value otherwise.

### GetNumInstancesOk

`func (o *ServiceServiceDetails) GetNumInstancesOk() (*int32, bool)`

GetNumInstancesOk returns a tuple with the NumInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumInstances

`func (o *ServiceServiceDetails) SetNumInstances(v int32)`

SetNumInstances sets NumInstances field to given value.

### HasNumInstances

`func (o *ServiceServiceDetails) HasNumInstances() bool`

HasNumInstances returns a boolean if a field has been set.

### GetOpenPorts

`func (o *ServiceServiceDetails) GetOpenPorts() []ServerPort`

GetOpenPorts returns the OpenPorts field if non-nil, zero value otherwise.

### GetOpenPortsOk

`func (o *ServiceServiceDetails) GetOpenPortsOk() (*[]ServerPort, bool)`

GetOpenPortsOk returns a tuple with the OpenPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenPorts

`func (o *ServiceServiceDetails) SetOpenPorts(v []ServerPort)`

SetOpenPorts sets OpenPorts field to given value.

### HasOpenPorts

`func (o *ServiceServiceDetails) HasOpenPorts() bool`

HasOpenPorts returns a boolean if a field has been set.

### GetPlan

`func (o *ServiceServiceDetails) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *ServiceServiceDetails) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *ServiceServiceDetails) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *ServiceServiceDetails) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetRegion

`func (o *ServiceServiceDetails) GetRegion() Region`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ServiceServiceDetails) GetRegionOk() (*Region, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ServiceServiceDetails) SetRegion(v Region)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ServiceServiceDetails) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetLastSuccessfulRunAt

`func (o *ServiceServiceDetails) GetLastSuccessfulRunAt() time.Time`

GetLastSuccessfulRunAt returns the LastSuccessfulRunAt field if non-nil, zero value otherwise.

### GetLastSuccessfulRunAtOk

`func (o *ServiceServiceDetails) GetLastSuccessfulRunAtOk() (*time.Time, bool)`

GetLastSuccessfulRunAtOk returns a tuple with the LastSuccessfulRunAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessfulRunAt

`func (o *ServiceServiceDetails) SetLastSuccessfulRunAt(v time.Time)`

SetLastSuccessfulRunAt sets LastSuccessfulRunAt field to given value.

### HasLastSuccessfulRunAt

`func (o *ServiceServiceDetails) HasLastSuccessfulRunAt() bool`

HasLastSuccessfulRunAt returns a boolean if a field has been set.

### GetSchedule

`func (o *ServiceServiceDetails) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ServiceServiceDetails) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ServiceServiceDetails) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *ServiceServiceDetails) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


