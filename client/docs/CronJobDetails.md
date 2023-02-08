# CronJobDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Env** | Pointer to [**ServiceEnv**](ServiceEnv.md) |  | [optional] 
**EnvSpecificDetails** | Pointer to [**WebServiceDetailsEnvSpecificDetails**](WebServiceDetailsEnvSpecificDetails.md) |  | [optional] 
**LastSuccessfulRunAt** | Pointer to **time.Time** |  | [optional] 
**Plan** | Pointer to **string** |  | [optional] 
**Region** | Pointer to [**Region**](Region.md) |  | [optional] 
**Schedule** | Pointer to **string** |  | [optional] 

## Methods

### NewCronJobDetails

`func NewCronJobDetails() *CronJobDetails`

NewCronJobDetails instantiates a new CronJobDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCronJobDetailsWithDefaults

`func NewCronJobDetailsWithDefaults() *CronJobDetails`

NewCronJobDetailsWithDefaults instantiates a new CronJobDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnv

`func (o *CronJobDetails) GetEnv() ServiceEnv`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *CronJobDetails) GetEnvOk() (*ServiceEnv, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *CronJobDetails) SetEnv(v ServiceEnv)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *CronJobDetails) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetEnvSpecificDetails

`func (o *CronJobDetails) GetEnvSpecificDetails() WebServiceDetailsEnvSpecificDetails`

GetEnvSpecificDetails returns the EnvSpecificDetails field if non-nil, zero value otherwise.

### GetEnvSpecificDetailsOk

`func (o *CronJobDetails) GetEnvSpecificDetailsOk() (*WebServiceDetailsEnvSpecificDetails, bool)`

GetEnvSpecificDetailsOk returns a tuple with the EnvSpecificDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvSpecificDetails

`func (o *CronJobDetails) SetEnvSpecificDetails(v WebServiceDetailsEnvSpecificDetails)`

SetEnvSpecificDetails sets EnvSpecificDetails field to given value.

### HasEnvSpecificDetails

`func (o *CronJobDetails) HasEnvSpecificDetails() bool`

HasEnvSpecificDetails returns a boolean if a field has been set.

### GetLastSuccessfulRunAt

`func (o *CronJobDetails) GetLastSuccessfulRunAt() time.Time`

GetLastSuccessfulRunAt returns the LastSuccessfulRunAt field if non-nil, zero value otherwise.

### GetLastSuccessfulRunAtOk

`func (o *CronJobDetails) GetLastSuccessfulRunAtOk() (*time.Time, bool)`

GetLastSuccessfulRunAtOk returns a tuple with the LastSuccessfulRunAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessfulRunAt

`func (o *CronJobDetails) SetLastSuccessfulRunAt(v time.Time)`

SetLastSuccessfulRunAt sets LastSuccessfulRunAt field to given value.

### HasLastSuccessfulRunAt

`func (o *CronJobDetails) HasLastSuccessfulRunAt() bool`

HasLastSuccessfulRunAt returns a boolean if a field has been set.

### GetPlan

`func (o *CronJobDetails) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *CronJobDetails) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *CronJobDetails) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *CronJobDetails) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetRegion

`func (o *CronJobDetails) GetRegion() Region`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CronJobDetails) GetRegionOk() (*Region, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CronJobDetails) SetRegion(v Region)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CronJobDetails) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSchedule

`func (o *CronJobDetails) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CronJobDetails) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CronJobDetails) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *CronJobDetails) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


