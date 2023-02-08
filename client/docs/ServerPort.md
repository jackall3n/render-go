# ServerPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **int32** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 

## Methods

### NewServerPort

`func NewServerPort() *ServerPort`

NewServerPort instantiates a new ServerPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerPortWithDefaults

`func NewServerPortWithDefaults() *ServerPort`

NewServerPortWithDefaults instantiates a new ServerPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *ServerPort) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ServerPort) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ServerPort) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ServerPort) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *ServerPort) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ServerPort) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ServerPort) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ServerPort) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


