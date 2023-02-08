# \OwnersApi

All URIs are relative to *https://api.render.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOwner**](OwnersApi.md#GetOwner) | **Get** /owners/{ownerId} | Retrieve user or team
[**GetOwners**](OwnersApi.md#GetOwners) | **Get** /owners | List authorized users and teams



## GetOwner

> Owner GetOwner(ctx, ownerId).Execute()

Retrieve user or team

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ownerId := "ownerId_example" // string | The ID of the user or team

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OwnersApi.GetOwner(context.Background(), ownerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OwnersApi.GetOwner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOwner`: Owner
    fmt.Fprintf(os.Stdout, "Response from `OwnersApi.GetOwner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerId** | **string** | The ID of the user or team | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOwnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Owner**](Owner.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOwners

> []GetOwners200ResponseInner GetOwners(ctx).Name(name).Email(email).Cursor(cursor).Limit(limit).Execute()

List authorized users and teams

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := []string{"Inner_example"} // []string |  (optional)
    email := []string{"Inner_example"} // []string |  (optional)
    cursor := string(BYTE_ARRAY_DATA_HERE) // string | Cursor to begin retrieving entries for this query (optional)
    limit := float32(8.14) // float32 | Max number of items that can be returned (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OwnersApi.GetOwners(context.Background()).Name(name).Email(email).Cursor(cursor).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OwnersApi.GetOwners``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOwners`: []GetOwners200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `OwnersApi.GetOwners`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOwnersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **[]string** |  | 
 **email** | **[]string** |  | 
 **cursor** | **string** | Cursor to begin retrieving entries for this query | 
 **limit** | **float32** | Max number of items that can be returned | [default to 20]

### Return type

[**[]GetOwners200ResponseInner**](GetOwners200ResponseInner.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

