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
)

// StaticSiteDetails struct for StaticSiteDetails
type StaticSiteDetails struct {
	BuildCommand *string `json:"buildCommand,omitempty"`
	ParentServer *StaticSiteDetailsParentServer `json:"parentServer,omitempty"`
	PublishPath *string `json:"publishPath,omitempty"`
	PullRequestPreviewsEnabled *string `json:"pullRequestPreviewsEnabled,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewStaticSiteDetails instantiates a new StaticSiteDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStaticSiteDetails() *StaticSiteDetails {
	this := StaticSiteDetails{}
	return &this
}

// NewStaticSiteDetailsWithDefaults instantiates a new StaticSiteDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStaticSiteDetailsWithDefaults() *StaticSiteDetails {
	this := StaticSiteDetails{}
	return &this
}

// GetBuildCommand returns the BuildCommand field value if set, zero value otherwise.
func (o *StaticSiteDetails) GetBuildCommand() string {
	if o == nil || isNil(o.BuildCommand) {
		var ret string
		return ret
	}
	return *o.BuildCommand
}

// GetBuildCommandOk returns a tuple with the BuildCommand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticSiteDetails) GetBuildCommandOk() (*string, bool) {
	if o == nil || isNil(o.BuildCommand) {
    return nil, false
	}
	return o.BuildCommand, true
}

// HasBuildCommand returns a boolean if a field has been set.
func (o *StaticSiteDetails) HasBuildCommand() bool {
	if o != nil && !isNil(o.BuildCommand) {
		return true
	}

	return false
}

// SetBuildCommand gets a reference to the given string and assigns it to the BuildCommand field.
func (o *StaticSiteDetails) SetBuildCommand(v string) {
	o.BuildCommand = &v
}

// GetParentServer returns the ParentServer field value if set, zero value otherwise.
func (o *StaticSiteDetails) GetParentServer() StaticSiteDetailsParentServer {
	if o == nil || isNil(o.ParentServer) {
		var ret StaticSiteDetailsParentServer
		return ret
	}
	return *o.ParentServer
}

// GetParentServerOk returns a tuple with the ParentServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticSiteDetails) GetParentServerOk() (*StaticSiteDetailsParentServer, bool) {
	if o == nil || isNil(o.ParentServer) {
    return nil, false
	}
	return o.ParentServer, true
}

// HasParentServer returns a boolean if a field has been set.
func (o *StaticSiteDetails) HasParentServer() bool {
	if o != nil && !isNil(o.ParentServer) {
		return true
	}

	return false
}

// SetParentServer gets a reference to the given StaticSiteDetailsParentServer and assigns it to the ParentServer field.
func (o *StaticSiteDetails) SetParentServer(v StaticSiteDetailsParentServer) {
	o.ParentServer = &v
}

// GetPublishPath returns the PublishPath field value if set, zero value otherwise.
func (o *StaticSiteDetails) GetPublishPath() string {
	if o == nil || isNil(o.PublishPath) {
		var ret string
		return ret
	}
	return *o.PublishPath
}

// GetPublishPathOk returns a tuple with the PublishPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticSiteDetails) GetPublishPathOk() (*string, bool) {
	if o == nil || isNil(o.PublishPath) {
    return nil, false
	}
	return o.PublishPath, true
}

// HasPublishPath returns a boolean if a field has been set.
func (o *StaticSiteDetails) HasPublishPath() bool {
	if o != nil && !isNil(o.PublishPath) {
		return true
	}

	return false
}

// SetPublishPath gets a reference to the given string and assigns it to the PublishPath field.
func (o *StaticSiteDetails) SetPublishPath(v string) {
	o.PublishPath = &v
}

// GetPullRequestPreviewsEnabled returns the PullRequestPreviewsEnabled field value if set, zero value otherwise.
func (o *StaticSiteDetails) GetPullRequestPreviewsEnabled() string {
	if o == nil || isNil(o.PullRequestPreviewsEnabled) {
		var ret string
		return ret
	}
	return *o.PullRequestPreviewsEnabled
}

// GetPullRequestPreviewsEnabledOk returns a tuple with the PullRequestPreviewsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticSiteDetails) GetPullRequestPreviewsEnabledOk() (*string, bool) {
	if o == nil || isNil(o.PullRequestPreviewsEnabled) {
    return nil, false
	}
	return o.PullRequestPreviewsEnabled, true
}

// HasPullRequestPreviewsEnabled returns a boolean if a field has been set.
func (o *StaticSiteDetails) HasPullRequestPreviewsEnabled() bool {
	if o != nil && !isNil(o.PullRequestPreviewsEnabled) {
		return true
	}

	return false
}

// SetPullRequestPreviewsEnabled gets a reference to the given string and assigns it to the PullRequestPreviewsEnabled field.
func (o *StaticSiteDetails) SetPullRequestPreviewsEnabled(v string) {
	o.PullRequestPreviewsEnabled = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *StaticSiteDetails) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticSiteDetails) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *StaticSiteDetails) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *StaticSiteDetails) SetUrl(v string) {
	o.Url = &v
}

func (o StaticSiteDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BuildCommand) {
		toSerialize["buildCommand"] = o.BuildCommand
	}
	if !isNil(o.ParentServer) {
		toSerialize["parentServer"] = o.ParentServer
	}
	if !isNil(o.PublishPath) {
		toSerialize["publishPath"] = o.PublishPath
	}
	if !isNil(o.PullRequestPreviewsEnabled) {
		toSerialize["pullRequestPreviewsEnabled"] = o.PullRequestPreviewsEnabled
	}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableStaticSiteDetails struct {
	value *StaticSiteDetails
	isSet bool
}

func (v NullableStaticSiteDetails) Get() *StaticSiteDetails {
	return v.value
}

func (v *NullableStaticSiteDetails) Set(val *StaticSiteDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableStaticSiteDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableStaticSiteDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStaticSiteDetails(val *StaticSiteDetails) *NullableStaticSiteDetails {
	return &NullableStaticSiteDetails{value: val, isSet: true}
}

func (v NullableStaticSiteDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStaticSiteDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

