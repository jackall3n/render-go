# CronJobDetailsPATCH

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvSpecificDetails** | Pointer to [**WebServiceDetailsPATCHEnvSpecificDetails**](WebServiceDetailsPATCHEnvSpecificDetails.md) |  | [optional] 
**Plan** | Pointer to **string** |  | [optional] 
**Schedule** | Pointer to **string** |  | [optional] 

## Methods

### NewCronJobDetailsPATCH

`func NewCronJobDetailsPATCH() *CronJobDetailsPATCH`

NewCronJobDetailsPATCH instantiates a new CronJobDetailsPATCH object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCronJobDetailsPATCHWithDefaults

`func NewCronJobDetailsPATCHWithDefaults() *CronJobDetailsPATCH`

NewCronJobDetailsPATCHWithDefaults instantiates a new CronJobDetailsPATCH object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvSpecificDetails

`func (o *CronJobDetailsPATCH) GetEnvSpecificDetails() WebServiceDetailsPATCHEnvSpecificDetails`

GetEnvSpecificDetails returns the EnvSpecificDetails field if non-nil, zero value otherwise.

### GetEnvSpecificDetailsOk

`func (o *CronJobDetailsPATCH) GetEnvSpecificDetailsOk() (*WebServiceDetailsPATCHEnvSpecificDetails, bool)`

GetEnvSpecificDetailsOk returns a tuple with the EnvSpecificDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvSpecificDetails

`func (o *CronJobDetailsPATCH) SetEnvSpecificDetails(v WebServiceDetailsPATCHEnvSpecificDetails)`

SetEnvSpecificDetails sets EnvSpecificDetails field to given value.

### HasEnvSpecificDetails

`func (o *CronJobDetailsPATCH) HasEnvSpecificDetails() bool`

HasEnvSpecificDetails returns a boolean if a field has been set.

### GetPlan

`func (o *CronJobDetailsPATCH) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *CronJobDetailsPATCH) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *CronJobDetailsPATCH) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *CronJobDetailsPATCH) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetSchedule

`func (o *CronJobDetailsPATCH) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CronJobDetailsPATCH) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CronJobDetailsPATCH) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *CronJobDetailsPATCH) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


