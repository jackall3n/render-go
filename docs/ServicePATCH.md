# ServicePATCH

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoDeploy** | Pointer to **string** |  | [optional] 
**Branch** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ServiceDetails** | Pointer to [**ServicePATCHServiceDetails**](ServicePATCHServiceDetails.md) |  | [optional] 

## Methods

### NewServicePATCH

`func NewServicePATCH() *ServicePATCH`

NewServicePATCH instantiates a new ServicePATCH object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePATCHWithDefaults

`func NewServicePATCHWithDefaults() *ServicePATCH`

NewServicePATCHWithDefaults instantiates a new ServicePATCH object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoDeploy

`func (o *ServicePATCH) GetAutoDeploy() string`

GetAutoDeploy returns the AutoDeploy field if non-nil, zero value otherwise.

### GetAutoDeployOk

`func (o *ServicePATCH) GetAutoDeployOk() (*string, bool)`

GetAutoDeployOk returns a tuple with the AutoDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDeploy

`func (o *ServicePATCH) SetAutoDeploy(v string)`

SetAutoDeploy sets AutoDeploy field to given value.

### HasAutoDeploy

`func (o *ServicePATCH) HasAutoDeploy() bool`

HasAutoDeploy returns a boolean if a field has been set.

### GetBranch

`func (o *ServicePATCH) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *ServicePATCH) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *ServicePATCH) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *ServicePATCH) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetName

`func (o *ServicePATCH) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServicePATCH) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServicePATCH) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServicePATCH) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServiceDetails

`func (o *ServicePATCH) GetServiceDetails() ServicePATCHServiceDetails`

GetServiceDetails returns the ServiceDetails field if non-nil, zero value otherwise.

### GetServiceDetailsOk

`func (o *ServicePATCH) GetServiceDetailsOk() (*ServicePATCHServiceDetails, bool)`

GetServiceDetailsOk returns a tuple with the ServiceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDetails

`func (o *ServicePATCH) SetServiceDetails(v ServicePATCHServiceDetails)`

SetServiceDetails sets ServiceDetails field to given value.

### HasServiceDetails

`func (o *ServicePATCH) HasServiceDetails() bool`

HasServiceDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


