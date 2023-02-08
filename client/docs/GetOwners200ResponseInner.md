# GetOwners200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | Pointer to [**Owner**](Owner.md) |  | [optional] 
**Cursor** | Pointer to **string** |  | [optional] 

## Methods

### NewGetOwners200ResponseInner

`func NewGetOwners200ResponseInner() *GetOwners200ResponseInner`

NewGetOwners200ResponseInner instantiates a new GetOwners200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOwners200ResponseInnerWithDefaults

`func NewGetOwners200ResponseInnerWithDefaults() *GetOwners200ResponseInner`

NewGetOwners200ResponseInnerWithDefaults instantiates a new GetOwners200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwner

`func (o *GetOwners200ResponseInner) GetOwner() Owner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GetOwners200ResponseInner) GetOwnerOk() (*Owner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GetOwners200ResponseInner) SetOwner(v Owner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GetOwners200ResponseInner) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetCursor

`func (o *GetOwners200ResponseInner) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *GetOwners200ResponseInner) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *GetOwners200ResponseInner) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *GetOwners200ResponseInner) HasCursor() bool`

HasCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


