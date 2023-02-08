# CronJobDetailsPOST

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Env** | [**ServiceEnv**](ServiceEnv.md) |  | 
**EnvSpecificDetails** | Pointer to [**WebServiceDetailsPOSTEnvSpecificDetails**](WebServiceDetailsPOSTEnvSpecificDetails.md) |  | [optional] 
**Plan** | Pointer to **string** | Defaults to \&quot;starter\&quot; | [optional] [default to "starter"]
**Region** | Pointer to [**Region**](Region.md) |  | [optional] 
**Schedule** | **string** |  | 

## Methods

### NewCronJobDetailsPOST

`func NewCronJobDetailsPOST(env ServiceEnv, schedule string, ) *CronJobDetailsPOST`

NewCronJobDetailsPOST instantiates a new CronJobDetailsPOST object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCronJobDetailsPOSTWithDefaults

`func NewCronJobDetailsPOSTWithDefaults() *CronJobDetailsPOST`

NewCronJobDetailsPOSTWithDefaults instantiates a new CronJobDetailsPOST object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnv

`func (o *CronJobDetailsPOST) GetEnv() ServiceEnv`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *CronJobDetailsPOST) GetEnvOk() (*ServiceEnv, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *CronJobDetailsPOST) SetEnv(v ServiceEnv)`

SetEnv sets Env field to given value.


### GetEnvSpecificDetails

`func (o *CronJobDetailsPOST) GetEnvSpecificDetails() WebServiceDetailsPOSTEnvSpecificDetails`

GetEnvSpecificDetails returns the EnvSpecificDetails field if non-nil, zero value otherwise.

### GetEnvSpecificDetailsOk

`func (o *CronJobDetailsPOST) GetEnvSpecificDetailsOk() (*WebServiceDetailsPOSTEnvSpecificDetails, bool)`

GetEnvSpecificDetailsOk returns a tuple with the EnvSpecificDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvSpecificDetails

`func (o *CronJobDetailsPOST) SetEnvSpecificDetails(v WebServiceDetailsPOSTEnvSpecificDetails)`

SetEnvSpecificDetails sets EnvSpecificDetails field to given value.

### HasEnvSpecificDetails

`func (o *CronJobDetailsPOST) HasEnvSpecificDetails() bool`

HasEnvSpecificDetails returns a boolean if a field has been set.

### GetPlan

`func (o *CronJobDetailsPOST) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *CronJobDetailsPOST) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *CronJobDetailsPOST) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *CronJobDetailsPOST) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetRegion

`func (o *CronJobDetailsPOST) GetRegion() Region`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CronJobDetailsPOST) GetRegionOk() (*Region, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CronJobDetailsPOST) SetRegion(v Region)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CronJobDetailsPOST) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSchedule

`func (o *CronJobDetailsPOST) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CronJobDetailsPOST) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CronJobDetailsPOST) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


