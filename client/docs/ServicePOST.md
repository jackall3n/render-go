# ServicePOST

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ServiceType**](ServiceType.md) |  | 
**Name** | **string** |  | 
**OwnerId** | **string** |  | 
**Repo** | **string** | Do not include the branch in the repo string. You can instead supply a &#39;branch&#39; parameter. | 
**AutoDeploy** | Pointer to **string** | Defaults to \&quot;yes\&quot; | [optional] [default to "true"]
**Branch** | Pointer to **string** | If left empty, this will fall back to the default branch of the repository | [optional] 
**EnvVars** | Pointer to [**[]UpdateEnvVarsForServiceRequestInner**](UpdateEnvVarsForServiceRequestInner.md) |  | [optional] 
**SecretFiles** | Pointer to [**[]ServicePOSTSecretFilesInner**](ServicePOSTSecretFilesInner.md) |  | [optional] 
**ServiceDetails** | Pointer to [**ServicePOSTServiceDetails**](ServicePOSTServiceDetails.md) |  | [optional] 

## Methods

### NewServicePOST

`func NewServicePOST(type_ ServiceType, name string, ownerId string, repo string, ) *ServicePOST`

NewServicePOST instantiates a new ServicePOST object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePOSTWithDefaults

`func NewServicePOSTWithDefaults() *ServicePOST`

NewServicePOSTWithDefaults instantiates a new ServicePOST object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ServicePOST) GetType() ServiceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServicePOST) GetTypeOk() (*ServiceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServicePOST) SetType(v ServiceType)`

SetType sets Type field to given value.


### GetName

`func (o *ServicePOST) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServicePOST) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServicePOST) SetName(v string)`

SetName sets Name field to given value.


### GetOwnerId

`func (o *ServicePOST) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *ServicePOST) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *ServicePOST) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetRepo

`func (o *ServicePOST) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *ServicePOST) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *ServicePOST) SetRepo(v string)`

SetRepo sets Repo field to given value.


### GetAutoDeploy

`func (o *ServicePOST) GetAutoDeploy() string`

GetAutoDeploy returns the AutoDeploy field if non-nil, zero value otherwise.

### GetAutoDeployOk

`func (o *ServicePOST) GetAutoDeployOk() (*string, bool)`

GetAutoDeployOk returns a tuple with the AutoDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDeploy

`func (o *ServicePOST) SetAutoDeploy(v string)`

SetAutoDeploy sets AutoDeploy field to given value.

### HasAutoDeploy

`func (o *ServicePOST) HasAutoDeploy() bool`

HasAutoDeploy returns a boolean if a field has been set.

### GetBranch

`func (o *ServicePOST) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *ServicePOST) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *ServicePOST) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *ServicePOST) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetEnvVars

`func (o *ServicePOST) GetEnvVars() []UpdateEnvVarsForServiceRequestInner`

GetEnvVars returns the EnvVars field if non-nil, zero value otherwise.

### GetEnvVarsOk

`func (o *ServicePOST) GetEnvVarsOk() (*[]UpdateEnvVarsForServiceRequestInner, bool)`

GetEnvVarsOk returns a tuple with the EnvVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVars

`func (o *ServicePOST) SetEnvVars(v []UpdateEnvVarsForServiceRequestInner)`

SetEnvVars sets EnvVars field to given value.

### HasEnvVars

`func (o *ServicePOST) HasEnvVars() bool`

HasEnvVars returns a boolean if a field has been set.

### GetSecretFiles

`func (o *ServicePOST) GetSecretFiles() []ServicePOSTSecretFilesInner`

GetSecretFiles returns the SecretFiles field if non-nil, zero value otherwise.

### GetSecretFilesOk

`func (o *ServicePOST) GetSecretFilesOk() (*[]ServicePOSTSecretFilesInner, bool)`

GetSecretFilesOk returns a tuple with the SecretFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretFiles

`func (o *ServicePOST) SetSecretFiles(v []ServicePOSTSecretFilesInner)`

SetSecretFiles sets SecretFiles field to given value.

### HasSecretFiles

`func (o *ServicePOST) HasSecretFiles() bool`

HasSecretFiles returns a boolean if a field has been set.

### GetServiceDetails

`func (o *ServicePOST) GetServiceDetails() ServicePOSTServiceDetails`

GetServiceDetails returns the ServiceDetails field if non-nil, zero value otherwise.

### GetServiceDetailsOk

`func (o *ServicePOST) GetServiceDetailsOk() (*ServicePOSTServiceDetails, bool)`

GetServiceDetailsOk returns a tuple with the ServiceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDetails

`func (o *ServicePOST) SetServiceDetails(v ServicePOSTServiceDetails)`

SetServiceDetails sets ServiceDetails field to given value.

### HasServiceDetails

`func (o *ServicePOST) HasServiceDetails() bool`

HasServiceDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


