# GetHeaders200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Headers** | Pointer to [**Header**](Header.md) |  | [optional] 
**Cursor** | Pointer to **string** |  | [optional] 

## Methods

### NewGetHeaders200ResponseInner

`func NewGetHeaders200ResponseInner() *GetHeaders200ResponseInner`

NewGetHeaders200ResponseInner instantiates a new GetHeaders200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetHeaders200ResponseInnerWithDefaults

`func NewGetHeaders200ResponseInnerWithDefaults() *GetHeaders200ResponseInner`

NewGetHeaders200ResponseInnerWithDefaults instantiates a new GetHeaders200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeaders

`func (o *GetHeaders200ResponseInner) GetHeaders() Header`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *GetHeaders200ResponseInner) GetHeadersOk() (*Header, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *GetHeaders200ResponseInner) SetHeaders(v Header)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *GetHeaders200ResponseInner) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetCursor

`func (o *GetHeaders200ResponseInner) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *GetHeaders200ResponseInner) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *GetHeaders200ResponseInner) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *GetHeaders200ResponseInner) HasCursor() bool`

HasCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


