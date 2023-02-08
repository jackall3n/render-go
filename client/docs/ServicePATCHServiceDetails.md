# ServicePATCHServiceDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildCommand** | Pointer to **string** |  | [optional] 
**PublishPath** | Pointer to **string** |  | [optional] 
**PullRequestPreviewsEnabled** | Pointer to **string** |  | [optional] 
**EnvSpecificDetails** | Pointer to [**WebServiceDetailsPATCHEnvSpecificDetails**](WebServiceDetailsPATCHEnvSpecificDetails.md) |  | [optional] 
**HealthCheckPath** | Pointer to **string** |  | [optional] 
**NumInstances** | Pointer to **int32** |  | [optional] 
**Plan** | Pointer to **string** |  | [optional] 
**Schedule** | Pointer to **string** |  | [optional] 

## Methods

### NewServicePATCHServiceDetails

`func NewServicePATCHServiceDetails() *ServicePATCHServiceDetails`

NewServicePATCHServiceDetails instantiates a new ServicePATCHServiceDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePATCHServiceDetailsWithDefaults

`func NewServicePATCHServiceDetailsWithDefaults() *ServicePATCHServiceDetails`

NewServicePATCHServiceDetailsWithDefaults instantiates a new ServicePATCHServiceDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildCommand

`func (o *ServicePATCHServiceDetails) GetBuildCommand() string`

GetBuildCommand returns the BuildCommand field if non-nil, zero value otherwise.

### GetBuildCommandOk

`func (o *ServicePATCHServiceDetails) GetBuildCommandOk() (*string, bool)`

GetBuildCommandOk returns a tuple with the BuildCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildCommand

`func (o *ServicePATCHServiceDetails) SetBuildCommand(v string)`

SetBuildCommand sets BuildCommand field to given value.

### HasBuildCommand

`func (o *ServicePATCHServiceDetails) HasBuildCommand() bool`

HasBuildCommand returns a boolean if a field has been set.

### GetPublishPath

`func (o *ServicePATCHServiceDetails) GetPublishPath() string`

GetPublishPath returns the PublishPath field if non-nil, zero value otherwise.

### GetPublishPathOk

`func (o *ServicePATCHServiceDetails) GetPublishPathOk() (*string, bool)`

GetPublishPathOk returns a tuple with the PublishPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishPath

`func (o *ServicePATCHServiceDetails) SetPublishPath(v string)`

SetPublishPath sets PublishPath field to given value.

### HasPublishPath

`func (o *ServicePATCHServiceDetails) HasPublishPath() bool`

HasPublishPath returns a boolean if a field has been set.

### GetPullRequestPreviewsEnabled

`func (o *ServicePATCHServiceDetails) GetPullRequestPreviewsEnabled() string`

GetPullRequestPreviewsEnabled returns the PullRequestPreviewsEnabled field if non-nil, zero value otherwise.

### GetPullRequestPreviewsEnabledOk

`func (o *ServicePATCHServiceDetails) GetPullRequestPreviewsEnabledOk() (*string, bool)`

GetPullRequestPreviewsEnabledOk returns a tuple with the PullRequestPreviewsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestPreviewsEnabled

`func (o *ServicePATCHServiceDetails) SetPullRequestPreviewsEnabled(v string)`

SetPullRequestPreviewsEnabled sets PullRequestPreviewsEnabled field to given value.

### HasPullRequestPreviewsEnabled

`func (o *ServicePATCHServiceDetails) HasPullRequestPreviewsEnabled() bool`

HasPullRequestPreviewsEnabled returns a boolean if a field has been set.

### GetEnvSpecificDetails

`func (o *ServicePATCHServiceDetails) GetEnvSpecificDetails() WebServiceDetailsPATCHEnvSpecificDetails`

GetEnvSpecificDetails returns the EnvSpecificDetails field if non-nil, zero value otherwise.

### GetEnvSpecificDetailsOk

`func (o *ServicePATCHServiceDetails) GetEnvSpecificDetailsOk() (*WebServiceDetailsPATCHEnvSpecificDetails, bool)`

GetEnvSpecificDetailsOk returns a tuple with the EnvSpecificDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvSpecificDetails

`func (o *ServicePATCHServiceDetails) SetEnvSpecificDetails(v WebServiceDetailsPATCHEnvSpecificDetails)`

SetEnvSpecificDetails sets EnvSpecificDetails field to given value.

### HasEnvSpecificDetails

`func (o *ServicePATCHServiceDetails) HasEnvSpecificDetails() bool`

HasEnvSpecificDetails returns a boolean if a field has been set.

### GetHealthCheckPath

`func (o *ServicePATCHServiceDetails) GetHealthCheckPath() string`

GetHealthCheckPath returns the HealthCheckPath field if non-nil, zero value otherwise.

### GetHealthCheckPathOk

`func (o *ServicePATCHServiceDetails) GetHealthCheckPathOk() (*string, bool)`

GetHealthCheckPathOk returns a tuple with the HealthCheckPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckPath

`func (o *ServicePATCHServiceDetails) SetHealthCheckPath(v string)`

SetHealthCheckPath sets HealthCheckPath field to given value.

### HasHealthCheckPath

`func (o *ServicePATCHServiceDetails) HasHealthCheckPath() bool`

HasHealthCheckPath returns a boolean if a field has been set.

### GetNumInstances

`func (o *ServicePATCHServiceDetails) GetNumInstances() int32`

GetNumInstances returns the NumInstances field if non-nil, zero value otherwise.

### GetNumInstancesOk

`func (o *ServicePATCHServiceDetails) GetNumInstancesOk() (*int32, bool)`

GetNumInstancesOk returns a tuple with the NumInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumInstances

`func (o *ServicePATCHServiceDetails) SetNumInstances(v int32)`

SetNumInstances sets NumInstances field to given value.

### HasNumInstances

`func (o *ServicePATCHServiceDetails) HasNumInstances() bool`

HasNumInstances returns a boolean if a field has been set.

### GetPlan

`func (o *ServicePATCHServiceDetails) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *ServicePATCHServiceDetails) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *ServicePATCHServiceDetails) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *ServicePATCHServiceDetails) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetSchedule

`func (o *ServicePATCHServiceDetails) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ServicePATCHServiceDetails) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ServicePATCHServiceDetails) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *ServicePATCHServiceDetails) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


