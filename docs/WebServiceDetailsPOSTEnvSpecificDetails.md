# WebServiceDetailsPOSTEnvSpecificDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DockerCommand** | Pointer to **string** |  | [optional] 
**DockerContext** | Pointer to **string** |  | [optional] 
**DockerfilePath** | Pointer to **string** | Defaults to \&quot;./Dockerfile\&quot; | [optional] [default to "./Dockerfile"]
**BuildCommand** | **string** |  | 
**StartCommand** | **string** |  | 

## Methods

### NewWebServiceDetailsPOSTEnvSpecificDetails

`func NewWebServiceDetailsPOSTEnvSpecificDetails(buildCommand string, startCommand string, ) *WebServiceDetailsPOSTEnvSpecificDetails`

NewWebServiceDetailsPOSTEnvSpecificDetails instantiates a new WebServiceDetailsPOSTEnvSpecificDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebServiceDetailsPOSTEnvSpecificDetailsWithDefaults

`func NewWebServiceDetailsPOSTEnvSpecificDetailsWithDefaults() *WebServiceDetailsPOSTEnvSpecificDetails`

NewWebServiceDetailsPOSTEnvSpecificDetailsWithDefaults instantiates a new WebServiceDetailsPOSTEnvSpecificDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDockerCommand

`func (o *WebServiceDetailsPOSTEnvSpecificDetails) GetDockerCommand() string`

GetDockerCommand returns the DockerCommand field if non-nil, zero value otherwise.

### GetDockerCommandOk

`func (o *WebServiceDetailsPOSTEnvSpecificDetails) GetDockerCommandOk() (*string, bool)`

GetDockerCommandOk returns a tuple with the DockerCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerCommand

`func (o *WebServiceDetailsPOSTEnvSpecificDetails) SetDockerCommand(v string)`

SetDockerCommand sets DockerCommand field to given value.

### HasDockerCommand

`func (o *WebServiceDetailsPOSTEnvSpecificDetails) HasDockerCommand() bool`

HasDockerCommand returns a boolean if a field has been set.

### GetDockerContext

`func (o *WebServiceDetailsPOSTEnvSpecificDetails) GetDockerContext() string`

GetDockerContext returns the DockerContext field if non-nil, zero value otherwise.

### GetDockerContextOk

`func (o *WebServiceDetailsPOSTEnvSpecificDetails) GetDockerContextOk() (*string, bool)`

GetDockerContextOk returns a tuple with the DockerContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerContext

`func (o *WebServiceDetailsPOSTEnvSpecificDetails) SetDockerContext(v string)`

SetDockerContext sets DockerContext field to given value.

### HasDockerContext

`func (o *WebServiceDetailsPOSTEnvSpecificDetails) HasDockerContext() bool`

HasDockerContext returns a boolean if a field has been set.

### GetDockerfilePath

`func (o *WebServiceDetailsPOSTEnvSpecificDetails) GetDockerfilePath() string`

GetDockerfilePath returns the DockerfilePath field if non-nil, zero value otherwise.

### GetDockerfilePathOk

`func (o *WebServiceDetailsPOSTEnvSpecificDetails) GetDockerfilePathOk() (*string, bool)`

GetDockerfilePathOk returns a tuple with the DockerfilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfilePath

`func (o *WebServiceDetailsPOSTEnvSpecificDetails) SetDockerfilePath(v string)`

SetDockerfilePath sets DockerfilePath field to given value.

### HasDockerfilePath

`func (o *WebServiceDetailsPOSTEnvSpecificDetails) HasDockerfilePath() bool`

HasDockerfilePath returns a boolean if a field has been set.

### GetBuildCommand

`func (o *WebServiceDetailsPOSTEnvSpecificDetails) GetBuildCommand() string`

GetBuildCommand returns the BuildCommand field if non-nil, zero value otherwise.

### GetBuildCommandOk

`func (o *WebServiceDetailsPOSTEnvSpecificDetails) GetBuildCommandOk() (*string, bool)`

GetBuildCommandOk returns a tuple with the BuildCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildCommand

`func (o *WebServiceDetailsPOSTEnvSpecificDetails) SetBuildCommand(v string)`

SetBuildCommand sets BuildCommand field to given value.


### GetStartCommand

`func (o *WebServiceDetailsPOSTEnvSpecificDetails) GetStartCommand() string`

GetStartCommand returns the StartCommand field if non-nil, zero value otherwise.

### GetStartCommandOk

`func (o *WebServiceDetailsPOSTEnvSpecificDetails) GetStartCommandOk() (*string, bool)`

GetStartCommandOk returns a tuple with the StartCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartCommand

`func (o *WebServiceDetailsPOSTEnvSpecificDetails) SetStartCommand(v string)`

SetStartCommand sets StartCommand field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


