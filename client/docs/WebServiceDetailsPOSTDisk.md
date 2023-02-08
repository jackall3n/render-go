# WebServiceDetailsPOSTDisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**MountPath** | **string** |  | 
**SizeGB** | Pointer to **int32** | Defaults to 1 | [optional] [default to 1]

## Methods

### NewWebServiceDetailsPOSTDisk

`func NewWebServiceDetailsPOSTDisk(name string, mountPath string, ) *WebServiceDetailsPOSTDisk`

NewWebServiceDetailsPOSTDisk instantiates a new WebServiceDetailsPOSTDisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebServiceDetailsPOSTDiskWithDefaults

`func NewWebServiceDetailsPOSTDiskWithDefaults() *WebServiceDetailsPOSTDisk`

NewWebServiceDetailsPOSTDiskWithDefaults instantiates a new WebServiceDetailsPOSTDisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WebServiceDetailsPOSTDisk) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebServiceDetailsPOSTDisk) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebServiceDetailsPOSTDisk) SetName(v string)`

SetName sets Name field to given value.


### GetMountPath

`func (o *WebServiceDetailsPOSTDisk) GetMountPath() string`

GetMountPath returns the MountPath field if non-nil, zero value otherwise.

### GetMountPathOk

`func (o *WebServiceDetailsPOSTDisk) GetMountPathOk() (*string, bool)`

GetMountPathOk returns a tuple with the MountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPath

`func (o *WebServiceDetailsPOSTDisk) SetMountPath(v string)`

SetMountPath sets MountPath field to given value.


### GetSizeGB

`func (o *WebServiceDetailsPOSTDisk) GetSizeGB() int32`

GetSizeGB returns the SizeGB field if non-nil, zero value otherwise.

### GetSizeGBOk

`func (o *WebServiceDetailsPOSTDisk) GetSizeGBOk() (*int32, bool)`

GetSizeGBOk returns a tuple with the SizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeGB

`func (o *WebServiceDetailsPOSTDisk) SetSizeGB(v int32)`

SetSizeGB sets SizeGB field to given value.

### HasSizeGB

`func (o *WebServiceDetailsPOSTDisk) HasSizeGB() bool`

HasSizeGB returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


