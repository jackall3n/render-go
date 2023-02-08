# ListJob200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Job** | Pointer to [**Job**](Job.md) |  | [optional] 
**Cursor** | Pointer to **string** |  | [optional] 

## Methods

### NewListJob200ResponseInner

`func NewListJob200ResponseInner() *ListJob200ResponseInner`

NewListJob200ResponseInner instantiates a new ListJob200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListJob200ResponseInnerWithDefaults

`func NewListJob200ResponseInnerWithDefaults() *ListJob200ResponseInner`

NewListJob200ResponseInnerWithDefaults instantiates a new ListJob200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJob

`func (o *ListJob200ResponseInner) GetJob() Job`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *ListJob200ResponseInner) GetJobOk() (*Job, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *ListJob200ResponseInner) SetJob(v Job)`

SetJob sets Job field to given value.

### HasJob

`func (o *ListJob200ResponseInner) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetCursor

`func (o *ListJob200ResponseInner) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *ListJob200ResponseInner) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *ListJob200ResponseInner) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *ListJob200ResponseInner) HasCursor() bool`

HasCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


