# ServicePOSTServiceDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildCommand** | Pointer to **string** |  | [optional] 
**Headers** | Pointer to [**[]Header**](Header.md) |  | [optional] 
**PublishPath** | Pointer to **string** | Defaults to \&quot;public\&quot; | [optional] [default to "public"]
**PullRequestPreviewsEnabled** | Pointer to **string** | Defaults to \&quot;no\&quot; | [optional] [default to "no"]
**Routes** | Pointer to [**[]Route**](Route.md) |  | [optional] 
**Disk** | Pointer to [**WebServiceDetailsPOSTDisk**](WebServiceDetailsPOSTDisk.md) |  | [optional] 
**Env** | [**ServiceEnv**](ServiceEnv.md) |  | 
**EnvSpecificDetails** | Pointer to [**WebServiceDetailsPOSTEnvSpecificDetails**](WebServiceDetailsPOSTEnvSpecificDetails.md) |  | [optional] 
**HealthCheckPath** | Pointer to **string** |  | [optional] 
**NumInstances** | Pointer to **int32** | Defaults to 1 | [optional] [default to 1]
**Plan** | Pointer to **string** | Defaults to \&quot;starter\&quot; | [optional] [default to "starter"]
**Region** | Pointer to [**Region**](Region.md) |  | [optional] 
**Schedule** | **string** |  | 

## Methods

### NewServicePOSTServiceDetails

`func NewServicePOSTServiceDetails(env ServiceEnv, schedule string, ) *ServicePOSTServiceDetails`

NewServicePOSTServiceDetails instantiates a new ServicePOSTServiceDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePOSTServiceDetailsWithDefaults

`func NewServicePOSTServiceDetailsWithDefaults() *ServicePOSTServiceDetails`

NewServicePOSTServiceDetailsWithDefaults instantiates a new ServicePOSTServiceDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildCommand

`func (o *ServicePOSTServiceDetails) GetBuildCommand() string`

GetBuildCommand returns the BuildCommand field if non-nil, zero value otherwise.

### GetBuildCommandOk

`func (o *ServicePOSTServiceDetails) GetBuildCommandOk() (*string, bool)`

GetBuildCommandOk returns a tuple with the BuildCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildCommand

`func (o *ServicePOSTServiceDetails) SetBuildCommand(v string)`

SetBuildCommand sets BuildCommand field to given value.

### HasBuildCommand

`func (o *ServicePOSTServiceDetails) HasBuildCommand() bool`

HasBuildCommand returns a boolean if a field has been set.

### GetHeaders

`func (o *ServicePOSTServiceDetails) GetHeaders() []Header`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *ServicePOSTServiceDetails) GetHeadersOk() (*[]Header, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *ServicePOSTServiceDetails) SetHeaders(v []Header)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *ServicePOSTServiceDetails) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetPublishPath

`func (o *ServicePOSTServiceDetails) GetPublishPath() string`

GetPublishPath returns the PublishPath field if non-nil, zero value otherwise.

### GetPublishPathOk

`func (o *ServicePOSTServiceDetails) GetPublishPathOk() (*string, bool)`

GetPublishPathOk returns a tuple with the PublishPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishPath

`func (o *ServicePOSTServiceDetails) SetPublishPath(v string)`

SetPublishPath sets PublishPath field to given value.

### HasPublishPath

`func (o *ServicePOSTServiceDetails) HasPublishPath() bool`

HasPublishPath returns a boolean if a field has been set.

### GetPullRequestPreviewsEnabled

`func (o *ServicePOSTServiceDetails) GetPullRequestPreviewsEnabled() string`

GetPullRequestPreviewsEnabled returns the PullRequestPreviewsEnabled field if non-nil, zero value otherwise.

### GetPullRequestPreviewsEnabledOk

`func (o *ServicePOSTServiceDetails) GetPullRequestPreviewsEnabledOk() (*string, bool)`

GetPullRequestPreviewsEnabledOk returns a tuple with the PullRequestPreviewsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestPreviewsEnabled

`func (o *ServicePOSTServiceDetails) SetPullRequestPreviewsEnabled(v string)`

SetPullRequestPreviewsEnabled sets PullRequestPreviewsEnabled field to given value.

### HasPullRequestPreviewsEnabled

`func (o *ServicePOSTServiceDetails) HasPullRequestPreviewsEnabled() bool`

HasPullRequestPreviewsEnabled returns a boolean if a field has been set.

### GetRoutes

`func (o *ServicePOSTServiceDetails) GetRoutes() []Route`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *ServicePOSTServiceDetails) GetRoutesOk() (*[]Route, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *ServicePOSTServiceDetails) SetRoutes(v []Route)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *ServicePOSTServiceDetails) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetDisk

`func (o *ServicePOSTServiceDetails) GetDisk() WebServiceDetailsPOSTDisk`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *ServicePOSTServiceDetails) GetDiskOk() (*WebServiceDetailsPOSTDisk, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *ServicePOSTServiceDetails) SetDisk(v WebServiceDetailsPOSTDisk)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *ServicePOSTServiceDetails) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetEnv

`func (o *ServicePOSTServiceDetails) GetEnv() ServiceEnv`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *ServicePOSTServiceDetails) GetEnvOk() (*ServiceEnv, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *ServicePOSTServiceDetails) SetEnv(v ServiceEnv)`

SetEnv sets Env field to given value.


### GetEnvSpecificDetails

`func (o *ServicePOSTServiceDetails) GetEnvSpecificDetails() WebServiceDetailsPOSTEnvSpecificDetails`

GetEnvSpecificDetails returns the EnvSpecificDetails field if non-nil, zero value otherwise.

### GetEnvSpecificDetailsOk

`func (o *ServicePOSTServiceDetails) GetEnvSpecificDetailsOk() (*WebServiceDetailsPOSTEnvSpecificDetails, bool)`

GetEnvSpecificDetailsOk returns a tuple with the EnvSpecificDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvSpecificDetails

`func (o *ServicePOSTServiceDetails) SetEnvSpecificDetails(v WebServiceDetailsPOSTEnvSpecificDetails)`

SetEnvSpecificDetails sets EnvSpecificDetails field to given value.

### HasEnvSpecificDetails

`func (o *ServicePOSTServiceDetails) HasEnvSpecificDetails() bool`

HasEnvSpecificDetails returns a boolean if a field has been set.

### GetHealthCheckPath

`func (o *ServicePOSTServiceDetails) GetHealthCheckPath() string`

GetHealthCheckPath returns the HealthCheckPath field if non-nil, zero value otherwise.

### GetHealthCheckPathOk

`func (o *ServicePOSTServiceDetails) GetHealthCheckPathOk() (*string, bool)`

GetHealthCheckPathOk returns a tuple with the HealthCheckPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckPath

`func (o *ServicePOSTServiceDetails) SetHealthCheckPath(v string)`

SetHealthCheckPath sets HealthCheckPath field to given value.

### HasHealthCheckPath

`func (o *ServicePOSTServiceDetails) HasHealthCheckPath() bool`

HasHealthCheckPath returns a boolean if a field has been set.

### GetNumInstances

`func (o *ServicePOSTServiceDetails) GetNumInstances() int32`

GetNumInstances returns the NumInstances field if non-nil, zero value otherwise.

### GetNumInstancesOk

`func (o *ServicePOSTServiceDetails) GetNumInstancesOk() (*int32, bool)`

GetNumInstancesOk returns a tuple with the NumInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumInstances

`func (o *ServicePOSTServiceDetails) SetNumInstances(v int32)`

SetNumInstances sets NumInstances field to given value.

### HasNumInstances

`func (o *ServicePOSTServiceDetails) HasNumInstances() bool`

HasNumInstances returns a boolean if a field has been set.

### GetPlan

`func (o *ServicePOSTServiceDetails) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *ServicePOSTServiceDetails) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *ServicePOSTServiceDetails) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *ServicePOSTServiceDetails) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetRegion

`func (o *ServicePOSTServiceDetails) GetRegion() Region`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ServicePOSTServiceDetails) GetRegionOk() (*Region, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ServicePOSTServiceDetails) SetRegion(v Region)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ServicePOSTServiceDetails) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSchedule

`func (o *ServicePOSTServiceDetails) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ServicePOSTServiceDetails) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ServicePOSTServiceDetails) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


