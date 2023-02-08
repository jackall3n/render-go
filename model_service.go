/*
Render Public API

Manage everything about your Render services

API version: 1.0.0
Contact: support@render.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package render

import (
	"encoding/json"
	"time"
)

// Service struct for Service
type Service struct {
	Id *string `json:"id,omitempty"`
	AutoDeploy *string `json:"autoDeploy,omitempty"`
	Branch *string `json:"branch,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Name *string `json:"name,omitempty"`
	NotifyOnFail *NotifySetting `json:"notifyOnFail,omitempty"`
	OwnerId *string `json:"ownerId,omitempty"`
	Repo *string `json:"repo,omitempty"`
	Slug *string `json:"slug,omitempty"`
	Suspended *string `json:"suspended,omitempty"`
	Suspenders []SuspenderType `json:"suspenders,omitempty"`
	Type *ServiceType `json:"type,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	ServiceDetails *ServiceServiceDetails `json:"serviceDetails,omitempty"`
}

// NewService instantiates a new Service object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewService() *Service {
	this := Service{}
	return &this
}

// NewServiceWithDefaults instantiates a new Service object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceWithDefaults() *Service {
	this := Service{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Service) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Service) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Service) SetId(v string) {
	o.Id = &v
}

// GetAutoDeploy returns the AutoDeploy field value if set, zero value otherwise.
func (o *Service) GetAutoDeploy() string {
	if o == nil || isNil(o.AutoDeploy) {
		var ret string
		return ret
	}
	return *o.AutoDeploy
}

// GetAutoDeployOk returns a tuple with the AutoDeploy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetAutoDeployOk() (*string, bool) {
	if o == nil || isNil(o.AutoDeploy) {
    return nil, false
	}
	return o.AutoDeploy, true
}

// HasAutoDeploy returns a boolean if a field has been set.
func (o *Service) HasAutoDeploy() bool {
	if o != nil && !isNil(o.AutoDeploy) {
		return true
	}

	return false
}

// SetAutoDeploy gets a reference to the given string and assigns it to the AutoDeploy field.
func (o *Service) SetAutoDeploy(v string) {
	o.AutoDeploy = &v
}

// GetBranch returns the Branch field value if set, zero value otherwise.
func (o *Service) GetBranch() string {
	if o == nil || isNil(o.Branch) {
		var ret string
		return ret
	}
	return *o.Branch
}

// GetBranchOk returns a tuple with the Branch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetBranchOk() (*string, bool) {
	if o == nil || isNil(o.Branch) {
    return nil, false
	}
	return o.Branch, true
}

// HasBranch returns a boolean if a field has been set.
func (o *Service) HasBranch() bool {
	if o != nil && !isNil(o.Branch) {
		return true
	}

	return false
}

// SetBranch gets a reference to the given string and assigns it to the Branch field.
func (o *Service) SetBranch(v string) {
	o.Branch = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Service) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Service) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Service) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Service) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Service) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Service) SetName(v string) {
	o.Name = &v
}

// GetNotifyOnFail returns the NotifyOnFail field value if set, zero value otherwise.
func (o *Service) GetNotifyOnFail() NotifySetting {
	if o == nil || isNil(o.NotifyOnFail) {
		var ret NotifySetting
		return ret
	}
	return *o.NotifyOnFail
}

// GetNotifyOnFailOk returns a tuple with the NotifyOnFail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetNotifyOnFailOk() (*NotifySetting, bool) {
	if o == nil || isNil(o.NotifyOnFail) {
    return nil, false
	}
	return o.NotifyOnFail, true
}

// HasNotifyOnFail returns a boolean if a field has been set.
func (o *Service) HasNotifyOnFail() bool {
	if o != nil && !isNil(o.NotifyOnFail) {
		return true
	}

	return false
}

// SetNotifyOnFail gets a reference to the given NotifySetting and assigns it to the NotifyOnFail field.
func (o *Service) SetNotifyOnFail(v NotifySetting) {
	o.NotifyOnFail = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *Service) GetOwnerId() string {
	if o == nil || isNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetOwnerIdOk() (*string, bool) {
	if o == nil || isNil(o.OwnerId) {
    return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *Service) HasOwnerId() bool {
	if o != nil && !isNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *Service) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetRepo returns the Repo field value if set, zero value otherwise.
func (o *Service) GetRepo() string {
	if o == nil || isNil(o.Repo) {
		var ret string
		return ret
	}
	return *o.Repo
}

// GetRepoOk returns a tuple with the Repo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetRepoOk() (*string, bool) {
	if o == nil || isNil(o.Repo) {
    return nil, false
	}
	return o.Repo, true
}

// HasRepo returns a boolean if a field has been set.
func (o *Service) HasRepo() bool {
	if o != nil && !isNil(o.Repo) {
		return true
	}

	return false
}

// SetRepo gets a reference to the given string and assigns it to the Repo field.
func (o *Service) SetRepo(v string) {
	o.Repo = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *Service) GetSlug() string {
	if o == nil || isNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetSlugOk() (*string, bool) {
	if o == nil || isNil(o.Slug) {
    return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *Service) HasSlug() bool {
	if o != nil && !isNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *Service) SetSlug(v string) {
	o.Slug = &v
}

// GetSuspended returns the Suspended field value if set, zero value otherwise.
func (o *Service) GetSuspended() string {
	if o == nil || isNil(o.Suspended) {
		var ret string
		return ret
	}
	return *o.Suspended
}

// GetSuspendedOk returns a tuple with the Suspended field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetSuspendedOk() (*string, bool) {
	if o == nil || isNil(o.Suspended) {
    return nil, false
	}
	return o.Suspended, true
}

// HasSuspended returns a boolean if a field has been set.
func (o *Service) HasSuspended() bool {
	if o != nil && !isNil(o.Suspended) {
		return true
	}

	return false
}

// SetSuspended gets a reference to the given string and assigns it to the Suspended field.
func (o *Service) SetSuspended(v string) {
	o.Suspended = &v
}

// GetSuspenders returns the Suspenders field value if set, zero value otherwise.
func (o *Service) GetSuspenders() []SuspenderType {
	if o == nil || isNil(o.Suspenders) {
		var ret []SuspenderType
		return ret
	}
	return o.Suspenders
}

// GetSuspendersOk returns a tuple with the Suspenders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetSuspendersOk() ([]SuspenderType, bool) {
	if o == nil || isNil(o.Suspenders) {
    return nil, false
	}
	return o.Suspenders, true
}

// HasSuspenders returns a boolean if a field has been set.
func (o *Service) HasSuspenders() bool {
	if o != nil && !isNil(o.Suspenders) {
		return true
	}

	return false
}

// SetSuspenders gets a reference to the given []SuspenderType and assigns it to the Suspenders field.
func (o *Service) SetSuspenders(v []SuspenderType) {
	o.Suspenders = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Service) GetType() ServiceType {
	if o == nil || isNil(o.Type) {
		var ret ServiceType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetTypeOk() (*ServiceType, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Service) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ServiceType and assigns it to the Type field.
func (o *Service) SetType(v ServiceType) {
	o.Type = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Service) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Service) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Service) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetServiceDetails returns the ServiceDetails field value if set, zero value otherwise.
func (o *Service) GetServiceDetails() ServiceServiceDetails {
	if o == nil || isNil(o.ServiceDetails) {
		var ret ServiceServiceDetails
		return ret
	}
	return *o.ServiceDetails
}

// GetServiceDetailsOk returns a tuple with the ServiceDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetServiceDetailsOk() (*ServiceServiceDetails, bool) {
	if o == nil || isNil(o.ServiceDetails) {
    return nil, false
	}
	return o.ServiceDetails, true
}

// HasServiceDetails returns a boolean if a field has been set.
func (o *Service) HasServiceDetails() bool {
	if o != nil && !isNil(o.ServiceDetails) {
		return true
	}

	return false
}

// SetServiceDetails gets a reference to the given ServiceServiceDetails and assigns it to the ServiceDetails field.
func (o *Service) SetServiceDetails(v ServiceServiceDetails) {
	o.ServiceDetails = &v
}

func (o Service) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.AutoDeploy) {
		toSerialize["autoDeploy"] = o.AutoDeploy
	}
	if !isNil(o.Branch) {
		toSerialize["branch"] = o.Branch
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.NotifyOnFail) {
		toSerialize["notifyOnFail"] = o.NotifyOnFail
	}
	if !isNil(o.OwnerId) {
		toSerialize["ownerId"] = o.OwnerId
	}
	if !isNil(o.Repo) {
		toSerialize["repo"] = o.Repo
	}
	if !isNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !isNil(o.Suspended) {
		toSerialize["suspended"] = o.Suspended
	}
	if !isNil(o.Suspenders) {
		toSerialize["suspenders"] = o.Suspenders
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !isNil(o.ServiceDetails) {
		toSerialize["serviceDetails"] = o.ServiceDetails
	}
	return json.Marshal(toSerialize)
}

type NullableService struct {
	value *Service
	isSet bool
}

func (v NullableService) Get() *Service {
	return v.value
}

func (v *NullableService) Set(val *Service) {
	v.value = val
	v.isSet = true
}

func (v NullableService) IsSet() bool {
	return v.isSet
}

func (v *NullableService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableService(val *Service) *NullableService {
	return &NullableService{value: val, isSet: true}
}

func (v NullableService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


