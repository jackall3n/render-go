# StaticSiteDetailsPOST

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildCommand** | Pointer to **string** |  | [optional] 
**Headers** | Pointer to [**[]Header**](Header.md) |  | [optional] 
**PublishPath** | Pointer to **string** | Defaults to \&quot;public\&quot; | [optional] [default to "public"]
**PullRequestPreviewsEnabled** | Pointer to **string** | Defaults to \&quot;no\&quot; | [optional] [default to "false"]
**Routes** | Pointer to [**[]Route**](Route.md) |  | [optional] 

## Methods

### NewStaticSiteDetailsPOST

`func NewStaticSiteDetailsPOST() *StaticSiteDetailsPOST`

NewStaticSiteDetailsPOST instantiates a new StaticSiteDetailsPOST object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticSiteDetailsPOSTWithDefaults

`func NewStaticSiteDetailsPOSTWithDefaults() *StaticSiteDetailsPOST`

NewStaticSiteDetailsPOSTWithDefaults instantiates a new StaticSiteDetailsPOST object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildCommand

`func (o *StaticSiteDetailsPOST) GetBuildCommand() string`

GetBuildCommand returns the BuildCommand field if non-nil, zero value otherwise.

### GetBuildCommandOk

`func (o *StaticSiteDetailsPOST) GetBuildCommandOk() (*string, bool)`

GetBuildCommandOk returns a tuple with the BuildCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildCommand

`func (o *StaticSiteDetailsPOST) SetBuildCommand(v string)`

SetBuildCommand sets BuildCommand field to given value.

### HasBuildCommand

`func (o *StaticSiteDetailsPOST) HasBuildCommand() bool`

HasBuildCommand returns a boolean if a field has been set.

### GetHeaders

`func (o *StaticSiteDetailsPOST) GetHeaders() []Header`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *StaticSiteDetailsPOST) GetHeadersOk() (*[]Header, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *StaticSiteDetailsPOST) SetHeaders(v []Header)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *StaticSiteDetailsPOST) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetPublishPath

`func (o *StaticSiteDetailsPOST) GetPublishPath() string`

GetPublishPath returns the PublishPath field if non-nil, zero value otherwise.

### GetPublishPathOk

`func (o *StaticSiteDetailsPOST) GetPublishPathOk() (*string, bool)`

GetPublishPathOk returns a tuple with the PublishPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishPath

`func (o *StaticSiteDetailsPOST) SetPublishPath(v string)`

SetPublishPath sets PublishPath field to given value.

### HasPublishPath

`func (o *StaticSiteDetailsPOST) HasPublishPath() bool`

HasPublishPath returns a boolean if a field has been set.

### GetPullRequestPreviewsEnabled

`func (o *StaticSiteDetailsPOST) GetPullRequestPreviewsEnabled() string`

GetPullRequestPreviewsEnabled returns the PullRequestPreviewsEnabled field if non-nil, zero value otherwise.

### GetPullRequestPreviewsEnabledOk

`func (o *StaticSiteDetailsPOST) GetPullRequestPreviewsEnabledOk() (*string, bool)`

GetPullRequestPreviewsEnabledOk returns a tuple with the PullRequestPreviewsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestPreviewsEnabled

`func (o *StaticSiteDetailsPOST) SetPullRequestPreviewsEnabled(v string)`

SetPullRequestPreviewsEnabled sets PullRequestPreviewsEnabled field to given value.

### HasPullRequestPreviewsEnabled

`func (o *StaticSiteDetailsPOST) HasPullRequestPreviewsEnabled() bool`

HasPullRequestPreviewsEnabled returns a boolean if a field has been set.

### GetRoutes

`func (o *StaticSiteDetailsPOST) GetRoutes() []Route`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *StaticSiteDetailsPOST) GetRoutesOk() (*[]Route, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *StaticSiteDetailsPOST) SetRoutes(v []Route)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *StaticSiteDetailsPOST) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


