# GetCustomDomains200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomDomain** | Pointer to [**CustomDomain**](CustomDomain.md) |  | [optional] 
**Cursor** | Pointer to **string** |  | [optional] 

## Methods

### NewGetCustomDomains200ResponseInner

`func NewGetCustomDomains200ResponseInner() *GetCustomDomains200ResponseInner`

NewGetCustomDomains200ResponseInner instantiates a new GetCustomDomains200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCustomDomains200ResponseInnerWithDefaults

`func NewGetCustomDomains200ResponseInnerWithDefaults() *GetCustomDomains200ResponseInner`

NewGetCustomDomains200ResponseInnerWithDefaults instantiates a new GetCustomDomains200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomDomain

`func (o *GetCustomDomains200ResponseInner) GetCustomDomain() CustomDomain`

GetCustomDomain returns the CustomDomain field if non-nil, zero value otherwise.

### GetCustomDomainOk

`func (o *GetCustomDomains200ResponseInner) GetCustomDomainOk() (*CustomDomain, bool)`

GetCustomDomainOk returns a tuple with the CustomDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDomain

`func (o *GetCustomDomains200ResponseInner) SetCustomDomain(v CustomDomain)`

SetCustomDomain sets CustomDomain field to given value.

### HasCustomDomain

`func (o *GetCustomDomains200ResponseInner) HasCustomDomain() bool`

HasCustomDomain returns a boolean if a field has been set.

### GetCursor

`func (o *GetCustomDomains200ResponseInner) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *GetCustomDomains200ResponseInner) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *GetCustomDomains200ResponseInner) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *GetCustomDomains200ResponseInner) HasCursor() bool`

HasCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


