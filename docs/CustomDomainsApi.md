# \CustomDomainsApi

All URIs are relative to *https://api.render.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomDomain**](CustomDomainsApi.md#CreateCustomDomain) | **Post** /services/{serviceId}/custom-domains | Add custom domain
[**DeleteCustomDomain**](CustomDomainsApi.md#DeleteCustomDomain) | **Delete** /services/{serviceId}/custom-domains/{customDomainIdOrName} | Delete custom domain
[**GetCustomDomain**](CustomDomainsApi.md#GetCustomDomain) | **Get** /services/{serviceId}/custom-domains/{customDomainIdOrName} | Retrieve custom domain
[**GetCustomDomains**](CustomDomainsApi.md#GetCustomDomains) | **Get** /services/{serviceId}/custom-domains | List custom domains
[**RefreshCustomDomain**](CustomDomainsApi.md#RefreshCustomDomain) | **Post** /services/{serviceId}/custom-domains/{customDomainIdOrName}/verify | Verify DNS configuration



## CreateCustomDomain

> []CustomDomain CreateCustomDomain(ctx, serviceId).CreateCustomDomainRequest(createCustomDomainRequest).Execute()

Add custom domain

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
    serviceId := "serviceId_example" // string | The ID of the service
    createCustomDomainRequest := *openapiclient.NewCreateCustomDomainRequest("Name_example") // CreateCustomDomainRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomDomainsApi.CreateCustomDomain(context.Background(), serviceId).CreateCustomDomainRequest(createCustomDomainRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomDomainsApi.CreateCustomDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCustomDomain`: []CustomDomain
    fmt.Fprintf(os.Stdout, "Response from `CustomDomainsApi.CreateCustomDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createCustomDomainRequest** | [**CreateCustomDomainRequest**](CreateCustomDomainRequest.md) |  | 

### Return type

[**[]CustomDomain**](CustomDomain.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomDomain

> DeleteCustomDomain(ctx, serviceId, customDomainIdOrName).Execute()

Delete custom domain

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
    serviceId := "serviceId_example" // string | The ID of the service
    customDomainIdOrName := "customDomainIdOrName_example" // string | The ID or name of the custom domain

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomDomainsApi.DeleteCustomDomain(context.Background(), serviceId, customDomainIdOrName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomDomainsApi.DeleteCustomDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service | 
**customDomainIdOrName** | **string** | The ID or name of the custom domain | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomDomain

> CustomDomain GetCustomDomain(ctx, serviceId, customDomainIdOrName).Execute()

Retrieve custom domain

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
    serviceId := "serviceId_example" // string | The ID of the service
    customDomainIdOrName := "customDomainIdOrName_example" // string | The ID or name of the custom domain

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomDomainsApi.GetCustomDomain(context.Background(), serviceId, customDomainIdOrName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomDomainsApi.GetCustomDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomDomain`: CustomDomain
    fmt.Fprintf(os.Stdout, "Response from `CustomDomainsApi.GetCustomDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service | 
**customDomainIdOrName** | **string** | The ID or name of the custom domain | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CustomDomain**](CustomDomain.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomDomains

> []GetCustomDomains200ResponseInner GetCustomDomains(ctx, serviceId).Cursor(cursor).Limit(limit).Name(name).DomainType(domainType).VerificationStatus(verificationStatus).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Execute()

List custom domains

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | The ID of the service
    cursor := string(BYTE_ARRAY_DATA_HERE) // string | Cursor to begin retrieving entries for this query (optional)
    limit := float32(8.14) // float32 | Max number of items that can be returned (optional) (default to 20)
    name := []string{"Inner_example"} // []string | Filter for the names of custom domain (optional)
    domainType := "domainType_example" // string | Filter for apex or subdomains (optional)
    verificationStatus := "verificationStatus_example" // string | Filter for verified or unverified custom domains (optional)
    createdBefore := time.Now() // time.Time | Filter for custom domains created before a certain time (specified as an ISO 8601 timestamp) (optional)
    createdAfter := time.Now() // time.Time | Filter for custom domains created after a certain time (specified as an ISO 8601 timestamp) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomDomainsApi.GetCustomDomains(context.Background(), serviceId).Cursor(cursor).Limit(limit).Name(name).DomainType(domainType).VerificationStatus(verificationStatus).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomDomainsApi.GetCustomDomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomDomains`: []GetCustomDomains200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `CustomDomainsApi.GetCustomDomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Cursor to begin retrieving entries for this query | 
 **limit** | **float32** | Max number of items that can be returned | [default to 20]
 **name** | **[]string** | Filter for the names of custom domain | 
 **domainType** | **string** | Filter for apex or subdomains | 
 **verificationStatus** | **string** | Filter for verified or unverified custom domains | 
 **createdBefore** | **time.Time** | Filter for custom domains created before a certain time (specified as an ISO 8601 timestamp) | 
 **createdAfter** | **time.Time** | Filter for custom domains created after a certain time (specified as an ISO 8601 timestamp) | 

### Return type

[**[]GetCustomDomains200ResponseInner**](GetCustomDomains200ResponseInner.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshCustomDomain

> RefreshCustomDomain(ctx, serviceId, customDomainIdOrName).Execute()

Verify DNS configuration

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
    serviceId := "serviceId_example" // string | The ID of the service
    customDomainIdOrName := "customDomainIdOrName_example" // string | The ID or name of the custom domain

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomDomainsApi.RefreshCustomDomain(context.Background(), serviceId, customDomainIdOrName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomDomainsApi.RefreshCustomDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service | 
**customDomainIdOrName** | **string** | The ID or name of the custom domain | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshCustomDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

