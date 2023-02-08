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

// WebServiceDetailsPOSTDisk struct for WebServiceDetailsPOSTDisk
type WebServiceDetailsPOSTDisk struct {
	Name string `json:"name"`
	MountPath string `json:"mountPath"`
	// Defaults to 1
	SizeGB *int32 `json:"sizeGB,omitempty"`
}

// NewWebServiceDetailsPOSTDisk instantiates a new WebServiceDetailsPOSTDisk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebServiceDetailsPOSTDisk(name string, mountPath string) *WebServiceDetailsPOSTDisk {
	this := WebServiceDetailsPOSTDisk{}
	this.Name = name
	this.MountPath = mountPath
	var sizeGB int32 = 1
	this.SizeGB = &sizeGB
	return &this
}

// NewWebServiceDetailsPOSTDiskWithDefaults instantiates a new WebServiceDetailsPOSTDisk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebServiceDetailsPOSTDiskWithDefaults() *WebServiceDetailsPOSTDisk {
	this := WebServiceDetailsPOSTDisk{}
	var sizeGB int32 = 1
	this.SizeGB = &sizeGB
	return &this
}

// GetName returns the Name field value
func (o *WebServiceDetailsPOSTDisk) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WebServiceDetailsPOSTDisk) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WebServiceDetailsPOSTDisk) SetName(v string) {
	o.Name = v
}

// GetMountPath returns the MountPath field value
func (o *WebServiceDetailsPOSTDisk) GetMountPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MountPath
}

// GetMountPathOk returns a tuple with the MountPath field value
// and a boolean to check if the value has been set.
func (o *WebServiceDetailsPOSTDisk) GetMountPathOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MountPath, true
}

// SetMountPath sets field value
func (o *WebServiceDetailsPOSTDisk) SetMountPath(v string) {
	o.MountPath = v
}

// GetSizeGB returns the SizeGB field value if set, zero value otherwise.
func (o *WebServiceDetailsPOSTDisk) GetSizeGB() int32 {
	if o == nil || isNil(o.SizeGB) {
		var ret int32
		return ret
	}
	return *o.SizeGB
}

// GetSizeGBOk returns a tuple with the SizeGB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebServiceDetailsPOSTDisk) GetSizeGBOk() (*int32, bool) {
	if o == nil || isNil(o.SizeGB) {
    return nil, false
	}
	return o.SizeGB, true
}

// HasSizeGB returns a boolean if a field has been set.
func (o *WebServiceDetailsPOSTDisk) HasSizeGB() bool {
	if o != nil && !isNil(o.SizeGB) {
		return true
	}

	return false
}

// SetSizeGB gets a reference to the given int32 and assigns it to the SizeGB field.
func (o *WebServiceDetailsPOSTDisk) SetSizeGB(v int32) {
	o.SizeGB = &v
}

func (o WebServiceDetailsPOSTDisk) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["mountPath"] = o.MountPath
	}
	if !isNil(o.SizeGB) {
		toSerialize["sizeGB"] = o.SizeGB
	}
	return json.Marshal(toSerialize)
}

type NullableWebServiceDetailsPOSTDisk struct {
	value *WebServiceDetailsPOSTDisk
	isSet bool
}

func (v NullableWebServiceDetailsPOSTDisk) Get() *WebServiceDetailsPOSTDisk {
	return v.value
}

func (v *NullableWebServiceDetailsPOSTDisk) Set(val *WebServiceDetailsPOSTDisk) {
	v.value = val
	v.isSet = true
}

func (v NullableWebServiceDetailsPOSTDisk) IsSet() bool {
	return v.isSet
}

func (v *NullableWebServiceDetailsPOSTDisk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebServiceDetailsPOSTDisk(val *WebServiceDetailsPOSTDisk) *NullableWebServiceDetailsPOSTDisk {
	return &NullableWebServiceDetailsPOSTDisk{value: val, isSet: true}
}

func (v NullableWebServiceDetailsPOSTDisk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebServiceDetailsPOSTDisk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

