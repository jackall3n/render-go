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

// GetDeploys200ResponseInner struct for GetDeploys200ResponseInner
type GetDeploys200ResponseInner struct {
	Deploy *Deploy `json:"deploy,omitempty"`
	Cursor *string `json:"cursor,omitempty"`
}

// NewGetDeploys200ResponseInner instantiates a new GetDeploys200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeploys200ResponseInner() *GetDeploys200ResponseInner {
	this := GetDeploys200ResponseInner{}
	return &this
}

// NewGetDeploys200ResponseInnerWithDefaults instantiates a new GetDeploys200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeploys200ResponseInnerWithDefaults() *GetDeploys200ResponseInner {
	this := GetDeploys200ResponseInner{}
	return &this
}

// GetDeploy returns the Deploy field value if set, zero value otherwise.
func (o *GetDeploys200ResponseInner) GetDeploy() Deploy {
	if o == nil || isNil(o.Deploy) {
		var ret Deploy
		return ret
	}
	return *o.Deploy
}

// GetDeployOk returns a tuple with the Deploy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeploys200ResponseInner) GetDeployOk() (*Deploy, bool) {
	if o == nil || isNil(o.Deploy) {
    return nil, false
	}
	return o.Deploy, true
}

// HasDeploy returns a boolean if a field has been set.
func (o *GetDeploys200ResponseInner) HasDeploy() bool {
	if o != nil && !isNil(o.Deploy) {
		return true
	}

	return false
}

// SetDeploy gets a reference to the given Deploy and assigns it to the Deploy field.
func (o *GetDeploys200ResponseInner) SetDeploy(v Deploy) {
	o.Deploy = &v
}

// GetCursor returns the Cursor field value if set, zero value otherwise.
func (o *GetDeploys200ResponseInner) GetCursor() string {
	if o == nil || isNil(o.Cursor) {
		var ret string
		return ret
	}
	return *o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeploys200ResponseInner) GetCursorOk() (*string, bool) {
	if o == nil || isNil(o.Cursor) {
    return nil, false
	}
	return o.Cursor, true
}

// HasCursor returns a boolean if a field has been set.
func (o *GetDeploys200ResponseInner) HasCursor() bool {
	if o != nil && !isNil(o.Cursor) {
		return true
	}

	return false
}

// SetCursor gets a reference to the given string and assigns it to the Cursor field.
func (o *GetDeploys200ResponseInner) SetCursor(v string) {
	o.Cursor = &v
}

func (o GetDeploys200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Deploy) {
		toSerialize["deploy"] = o.Deploy
	}
	if !isNil(o.Cursor) {
		toSerialize["cursor"] = o.Cursor
	}
	return json.Marshal(toSerialize)
}

type NullableGetDeploys200ResponseInner struct {
	value *GetDeploys200ResponseInner
	isSet bool
}

func (v NullableGetDeploys200ResponseInner) Get() *GetDeploys200ResponseInner {
	return v.value
}

func (v *NullableGetDeploys200ResponseInner) Set(val *GetDeploys200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeploys200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeploys200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeploys200ResponseInner(val *GetDeploys200ResponseInner) *NullableGetDeploys200ResponseInner {
	return &NullableGetDeploys200ResponseInner{value: val, isSet: true}
}

func (v NullableGetDeploys200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDeploys200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

