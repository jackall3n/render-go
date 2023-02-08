# CustomDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DomainType** | Pointer to **string** |  | [optional] 
**PublicSuffix** | Pointer to **string** |  | [optional] 
**RedirectForName** | Pointer to **string** |  | [optional] 
**VerificationStatus** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Server** | Pointer to [**StaticSiteDetailsParentServer**](StaticSiteDetailsParentServer.md) |  | [optional] 

## Methods

### NewCustomDomain

`func NewCustomDomain() *CustomDomain`

NewCustomDomain instantiates a new CustomDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomDomainWithDefaults

`func NewCustomDomainWithDefaults() *CustomDomain`

NewCustomDomainWithDefaults instantiates a new CustomDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomDomain) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomDomain) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomDomain) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomDomain) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CustomDomain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomDomain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomDomain) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomDomain) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDomainType

`func (o *CustomDomain) GetDomainType() string`

GetDomainType returns the DomainType field if non-nil, zero value otherwise.

### GetDomainTypeOk

`func (o *CustomDomain) GetDomainTypeOk() (*string, bool)`

GetDomainTypeOk returns a tuple with the DomainType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainType

`func (o *CustomDomain) SetDomainType(v string)`

SetDomainType sets DomainType field to given value.

### HasDomainType

`func (o *CustomDomain) HasDomainType() bool`

HasDomainType returns a boolean if a field has been set.

### GetPublicSuffix

`func (o *CustomDomain) GetPublicSuffix() string`

GetPublicSuffix returns the PublicSuffix field if non-nil, zero value otherwise.

### GetPublicSuffixOk

`func (o *CustomDomain) GetPublicSuffixOk() (*string, bool)`

GetPublicSuffixOk returns a tuple with the PublicSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicSuffix

`func (o *CustomDomain) SetPublicSuffix(v string)`

SetPublicSuffix sets PublicSuffix field to given value.

### HasPublicSuffix

`func (o *CustomDomain) HasPublicSuffix() bool`

HasPublicSuffix returns a boolean if a field has been set.

### GetRedirectForName

`func (o *CustomDomain) GetRedirectForName() string`

GetRedirectForName returns the RedirectForName field if non-nil, zero value otherwise.

### GetRedirectForNameOk

`func (o *CustomDomain) GetRedirectForNameOk() (*string, bool)`

GetRedirectForNameOk returns a tuple with the RedirectForName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectForName

`func (o *CustomDomain) SetRedirectForName(v string)`

SetRedirectForName sets RedirectForName field to given value.

### HasRedirectForName

`func (o *CustomDomain) HasRedirectForName() bool`

HasRedirectForName returns a boolean if a field has been set.

### GetVerificationStatus

`func (o *CustomDomain) GetVerificationStatus() string`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *CustomDomain) GetVerificationStatusOk() (*string, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *CustomDomain) SetVerificationStatus(v string)`

SetVerificationStatus sets VerificationStatus field to given value.

### HasVerificationStatus

`func (o *CustomDomain) HasVerificationStatus() bool`

HasVerificationStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CustomDomain) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CustomDomain) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CustomDomain) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CustomDomain) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetServer

`func (o *CustomDomain) GetServer() StaticSiteDetailsParentServer`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *CustomDomain) GetServerOk() (*StaticSiteDetailsParentServer, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *CustomDomain) SetServer(v StaticSiteDetailsParentServer)`

SetServer sets Server field to given value.

### HasServer

`func (o *CustomDomain) HasServer() bool`

HasServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


