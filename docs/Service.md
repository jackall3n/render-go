# Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**AutoDeploy** | Pointer to **string** |  | [optional] 
**Branch** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NotifyOnFail** | Pointer to [**NotifySetting**](NotifySetting.md) |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**Repo** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Suspended** | Pointer to **string** |  | [optional] 
**Suspenders** | Pointer to [**[]SuspenderType**](SuspenderType.md) |  | [optional] 
**Type** | Pointer to [**ServiceType**](ServiceType.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**ServiceDetails** | Pointer to [**ServiceServiceDetails**](ServiceServiceDetails.md) |  | [optional] 

## Methods

### NewService

`func NewService() *Service`

NewService instantiates a new Service object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceWithDefaults

`func NewServiceWithDefaults() *Service`

NewServiceWithDefaults instantiates a new Service object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Service) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Service) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Service) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Service) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAutoDeploy

`func (o *Service) GetAutoDeploy() string`

GetAutoDeploy returns the AutoDeploy field if non-nil, zero value otherwise.

### GetAutoDeployOk

`func (o *Service) GetAutoDeployOk() (*string, bool)`

GetAutoDeployOk returns a tuple with the AutoDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDeploy

`func (o *Service) SetAutoDeploy(v string)`

SetAutoDeploy sets AutoDeploy field to given value.

### HasAutoDeploy

`func (o *Service) HasAutoDeploy() bool`

HasAutoDeploy returns a boolean if a field has been set.

### GetBranch

`func (o *Service) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *Service) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *Service) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *Service) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Service) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Service) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Service) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Service) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetName

`func (o *Service) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Service) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Service) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Service) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotifyOnFail

`func (o *Service) GetNotifyOnFail() NotifySetting`

GetNotifyOnFail returns the NotifyOnFail field if non-nil, zero value otherwise.

### GetNotifyOnFailOk

`func (o *Service) GetNotifyOnFailOk() (*NotifySetting, bool)`

GetNotifyOnFailOk returns a tuple with the NotifyOnFail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyOnFail

`func (o *Service) SetNotifyOnFail(v NotifySetting)`

SetNotifyOnFail sets NotifyOnFail field to given value.

### HasNotifyOnFail

`func (o *Service) HasNotifyOnFail() bool`

HasNotifyOnFail returns a boolean if a field has been set.

### GetOwnerId

`func (o *Service) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *Service) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *Service) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *Service) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetRepo

`func (o *Service) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *Service) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *Service) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *Service) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetSlug

`func (o *Service) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Service) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Service) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *Service) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetSuspended

`func (o *Service) GetSuspended() string`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *Service) GetSuspendedOk() (*string, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *Service) SetSuspended(v string)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *Service) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### GetSuspenders

`func (o *Service) GetSuspenders() []SuspenderType`

GetSuspenders returns the Suspenders field if non-nil, zero value otherwise.

### GetSuspendersOk

`func (o *Service) GetSuspendersOk() (*[]SuspenderType, bool)`

GetSuspendersOk returns a tuple with the Suspenders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspenders

`func (o *Service) SetSuspenders(v []SuspenderType)`

SetSuspenders sets Suspenders field to given value.

### HasSuspenders

`func (o *Service) HasSuspenders() bool`

HasSuspenders returns a boolean if a field has been set.

### GetType

`func (o *Service) GetType() ServiceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Service) GetTypeOk() (*ServiceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Service) SetType(v ServiceType)`

SetType sets Type field to given value.

### HasType

`func (o *Service) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Service) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Service) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Service) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Service) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetServiceDetails

`func (o *Service) GetServiceDetails() ServiceServiceDetails`

GetServiceDetails returns the ServiceDetails field if non-nil, zero value otherwise.

### GetServiceDetailsOk

`func (o *Service) GetServiceDetailsOk() (*ServiceServiceDetails, bool)`

GetServiceDetailsOk returns a tuple with the ServiceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDetails

`func (o *Service) SetServiceDetails(v ServiceServiceDetails)`

SetServiceDetails sets ServiceDetails field to given value.

### HasServiceDetails

`func (o *Service) HasServiceDetails() bool`

HasServiceDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


