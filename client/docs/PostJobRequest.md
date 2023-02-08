# PostJobRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartCommand** | **string** |  | 
**PlanId** | Pointer to **string** |  | [optional] 

## Methods

### NewPostJobRequest

`func NewPostJobRequest(startCommand string, ) *PostJobRequest`

NewPostJobRequest instantiates a new PostJobRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostJobRequestWithDefaults

`func NewPostJobRequestWithDefaults() *PostJobRequest`

NewPostJobRequestWithDefaults instantiates a new PostJobRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartCommand

`func (o *PostJobRequest) GetStartCommand() string`

GetStartCommand returns the StartCommand field if non-nil, zero value otherwise.

### GetStartCommandOk

`func (o *PostJobRequest) GetStartCommandOk() (*string, bool)`

GetStartCommandOk returns a tuple with the StartCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartCommand

`func (o *PostJobRequest) SetStartCommand(v string)`

SetStartCommand sets StartCommand field to given value.


### GetPlanId

`func (o *PostJobRequest) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *PostJobRequest) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *PostJobRequest) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *PostJobRequest) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


