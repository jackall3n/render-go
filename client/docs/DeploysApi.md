# \DeploysApi

All URIs are relative to *https://api.render.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeploy**](DeploysApi.md#CreateDeploy) | **Post** /services/{serviceId}/deploys | Trigger a deploy
[**GetDeploy**](DeploysApi.md#GetDeploy) | **Get** /services/{serviceId}/deploys/{deployId} | Retrieve deploy
[**GetDeploys**](DeploysApi.md#GetDeploys) | **Get** /services/{serviceId}/deploys | List deploys



## CreateDeploy

> Deploy CreateDeploy(ctx, serviceId).CreateDeployRequest(createDeployRequest).Execute()

Trigger a deploy

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
    createDeployRequest := *openapiclient.NewCreateDeployRequest() // CreateDeployRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploysApi.CreateDeploy(context.Background(), serviceId).CreateDeployRequest(createDeployRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploysApi.CreateDeploy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeploy`: Deploy
    fmt.Fprintf(os.Stdout, "Response from `DeploysApi.CreateDeploy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeployRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeployRequest** | [**CreateDeployRequest**](CreateDeployRequest.md) |  | 

### Return type

[**Deploy**](Deploy.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploy

> Deploy GetDeploy(ctx, serviceId, deployId).Execute()

Retrieve deploy

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
    deployId := "deployId_example" // string | The ID of the deploy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploysApi.GetDeploy(context.Background(), serviceId, deployId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploysApi.GetDeploy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeploy`: Deploy
    fmt.Fprintf(os.Stdout, "Response from `DeploysApi.GetDeploy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service | 
**deployId** | **string** | The ID of the deploy | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeployRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Deploy**](Deploy.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploys

> []GetDeploys200ResponseInner GetDeploys(ctx, serviceId).StartTime(startTime).EndTime(endTime).Cursor(cursor).Limit(limit).Execute()

List deploys

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
    startTime := int32(56) // int32 | Epoch/Unix timestamp of start of time range to return (optional)
    endTime := int32(56) // int32 | Epoch/Unix timestamp of end of time range to return (optional)
    cursor := string(BYTE_ARRAY_DATA_HERE) // string | Cursor to begin retrieving entries for this query (optional)
    limit := float32(8.14) // float32 | Max number of items that can be returned (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploysApi.GetDeploys(context.Background(), serviceId).StartTime(startTime).EndTime(endTime).Cursor(cursor).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploysApi.GetDeploys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeploys`: []GetDeploys200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `DeploysApi.GetDeploys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTime** | **int32** | Epoch/Unix timestamp of start of time range to return | 
 **endTime** | **int32** | Epoch/Unix timestamp of end of time range to return | 
 **cursor** | **string** | Cursor to begin retrieving entries for this query | 
 **limit** | **float32** | Max number of items that can be returned | [default to 20]

### Return type

[**[]GetDeploys200ResponseInner**](GetDeploys200ResponseInner.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

