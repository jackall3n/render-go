# GetRoutes200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Routes** | Pointer to [**Route**](Route.md) |  | [optional] 
**Cursor** | Pointer to **string** |  | [optional] 

## Methods

### NewGetRoutes200ResponseInner

`func NewGetRoutes200ResponseInner() *GetRoutes200ResponseInner`

NewGetRoutes200ResponseInner instantiates a new GetRoutes200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRoutes200ResponseInnerWithDefaults

`func NewGetRoutes200ResponseInnerWithDefaults() *GetRoutes200ResponseInner`

NewGetRoutes200ResponseInnerWithDefaults instantiates a new GetRoutes200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoutes

`func (o *GetRoutes200ResponseInner) GetRoutes() Route`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *GetRoutes200ResponseInner) GetRoutesOk() (*Route, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *GetRoutes200ResponseInner) SetRoutes(v Route)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *GetRoutes200ResponseInner) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetCursor

`func (o *GetRoutes200ResponseInner) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *GetRoutes200ResponseInner) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *GetRoutes200ResponseInner) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *GetRoutes200ResponseInner) HasCursor() bool`

HasCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


