# \ServicesApi

All URIs are relative to *https://api.render.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateService**](ServicesApi.md#CreateService) | **Post** /services | Create service
[**DeleteService**](ServicesApi.md#DeleteService) | **Delete** /services/{serviceId} | Delete service
[**GetEnvVarsForService**](ServicesApi.md#GetEnvVarsForService) | **Get** /services/{serviceId}/env-vars | Retrieve environment variables
[**GetHeaders**](ServicesApi.md#GetHeaders) | **Get** /services/{serviceId}/headers | Retrieve headers
[**GetRoutes**](ServicesApi.md#GetRoutes) | **Get** /services/{serviceId}/routes | Retrieve redirect and rewrite rules
[**GetService**](ServicesApi.md#GetService) | **Get** /services/{serviceId} | Retrieve service
[**GetServices**](ServicesApi.md#GetServices) | **Get** /services | List services
[**ResumeService**](ServicesApi.md#ResumeService) | **Post** /services/{serviceId}/resume | Resume service
[**ScaleService**](ServicesApi.md#ScaleService) | **Post** /services/{serviceId}/scale | Scale service to desired number of instances
[**SuspendService**](ServicesApi.md#SuspendService) | **Post** /services/{serviceId}/suspend | Suspend service
[**UpdateEnvVarsForService**](ServicesApi.md#UpdateEnvVarsForService) | **Put** /services/{serviceId}/env-vars | Update environment variables
[**UpdateService**](ServicesApi.md#UpdateService) | **Patch** /services/{serviceId} | Update service



## CreateService

> CreateService201Response CreateService(ctx).ServicePOST(servicePOST).Execute()

Create service

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
    servicePOST := *openapiclient.NewServicePOST(openapiclient.serviceType("static_site"), "Name_example", "OwnerId_example", "https://github.com/render-examples/flask-hello-world") // ServicePOST | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServicesApi.CreateService(context.Background()).ServicePOST(servicePOST).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.CreateService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateService`: CreateService201Response
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.CreateService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **servicePOST** | [**ServicePOST**](ServicePOST.md) |  | 

### Return type

[**CreateService201Response**](CreateService201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteService

> DeleteService(ctx, serviceId).Execute()

Delete service

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServicesApi.DeleteService(context.Background(), serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.DeleteService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceRequest struct via the builder pattern


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


## GetEnvVarsForService

> []GetEnvVarsForService200ResponseInner GetEnvVarsForService(ctx, serviceId).Cursor(cursor).Limit(limit).Execute()

Retrieve environment variables

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
    cursor := string(BYTE_ARRAY_DATA_HERE) // string | Cursor to begin retrieving entries for this query (optional)
    limit := float32(8.14) // float32 | Max number of items that can be returned (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServicesApi.GetEnvVarsForService(context.Background(), serviceId).Cursor(cursor).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.GetEnvVarsForService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvVarsForService`: []GetEnvVarsForService200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.GetEnvVarsForService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvVarsForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Cursor to begin retrieving entries for this query | 
 **limit** | **float32** | Max number of items that can be returned | [default to 20]

### Return type

[**[]GetEnvVarsForService200ResponseInner**](GetEnvVarsForService200ResponseInner.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHeaders

> []GetHeaders200ResponseInner GetHeaders(ctx, serviceId).Path(path).Name(name).Value(value).Cursor(cursor).Limit(limit).Execute()

Retrieve headers

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
    path := []string{"Inner_example"} // []string | Filter for specific paths that headers apply to (optional)
    name := []string{"Inner_example"} // []string | Filter for header names (optional)
    value := []string{"Inner_example"} // []string | Filter for header values (optional)
    cursor := string(BYTE_ARRAY_DATA_HERE) // string | Cursor to begin retrieving entries for this query (optional)
    limit := float32(8.14) // float32 | Max number of items that can be returned (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServicesApi.GetHeaders(context.Background(), serviceId).Path(path).Name(name).Value(value).Cursor(cursor).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.GetHeaders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHeaders`: []GetHeaders200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.GetHeaders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHeadersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **path** | **[]string** | Filter for specific paths that headers apply to | 
 **name** | **[]string** | Filter for header names | 
 **value** | **[]string** | Filter for header values | 
 **cursor** | **string** | Cursor to begin retrieving entries for this query | 
 **limit** | **float32** | Max number of items that can be returned | [default to 20]

### Return type

[**[]GetHeaders200ResponseInner**](GetHeaders200ResponseInner.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoutes

> []GetRoutes200ResponseInner GetRoutes(ctx, serviceId).Type_(type_).Source(source).Destination(destination).Cursor(cursor).Limit(limit).Execute()

Retrieve redirect and rewrite rules

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
    type_ := []string{"Type_example"} // []string | Filter for the type of route rule (optional)
    source := []string{"Inner_example"} // []string | Filter for the source path of the route (optional)
    destination := []string{"Inner_example"} // []string | Filter for the destination path of the route (optional)
    cursor := string(BYTE_ARRAY_DATA_HERE) // string | Cursor to begin retrieving entries for this query (optional)
    limit := float32(8.14) // float32 | Max number of items that can be returned (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServicesApi.GetRoutes(context.Background(), serviceId).Type_(type_).Source(source).Destination(destination).Cursor(cursor).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.GetRoutes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoutes`: []GetRoutes200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.GetRoutes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **[]string** | Filter for the type of route rule | 
 **source** | **[]string** | Filter for the source path of the route | 
 **destination** | **[]string** | Filter for the destination path of the route | 
 **cursor** | **string** | Cursor to begin retrieving entries for this query | 
 **limit** | **float32** | Max number of items that can be returned | [default to 20]

### Return type

[**[]GetRoutes200ResponseInner**](GetRoutes200ResponseInner.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetService

> Service GetService(ctx, serviceId).Execute()

Retrieve service

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServicesApi.GetService(context.Background(), serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.GetService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetService`: Service
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.GetService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Service**](Service.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServices

> []GetServices200ResponseInner GetServices(ctx).Name(name).Type_(type_).Env(env).Region(region).Suspended(suspended).CreatedBefore(createdBefore).CreatedAfter(createdAfter).UpdatedBefore(updatedBefore).UpdatedAfter(updatedAfter).OwnerId(ownerId).Cursor(cursor).Limit(limit).Execute()

List services

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
    name := []string{"Inner_example"} // []string | Filter for the names of services (optional)
    type_ := []openapiclient.ServiceType{openapiclient.serviceType("static_site")} // []ServiceType | Filter for types of services (optional)
    env := []openapiclient.ServiceEnv{openapiclient.serviceEnv("docker")} // []ServiceEnv | Filter for environments of services (optional)
    region := []openapiclient.Region{openapiclient.region("oregon")} // []Region | Filter for regions of services (optional)
    suspended := []string{"Suspended_example"} // []string | Filter services based on whether they're suspended or not suspended (optional)
    createdBefore := time.Now() // time.Time | Filter for services created before a certain time (specified as an ISO 8601 timestamp) (optional)
    createdAfter := time.Now() // time.Time | Filter for services created after a certain time (specified as an ISO 8601 timestamp) (optional)
    updatedBefore := time.Now() // time.Time | Filter for services updated before a certain time (specified as an ISO 8601 timestamp) (optional)
    updatedAfter := time.Now() // time.Time | Filter for services updated after a certain time (specified as an ISO 8601 timestamp) (optional)
    ownerId := []string{"Inner_example"} // []string | The ID of the owner (team or personal user) whose resources should be returned (optional)
    cursor := string(BYTE_ARRAY_DATA_HERE) // string | Cursor to begin retrieving entries for this query (optional)
    limit := float32(8.14) // float32 | Max number of items that can be returned (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServicesApi.GetServices(context.Background()).Name(name).Type_(type_).Env(env).Region(region).Suspended(suspended).CreatedBefore(createdBefore).CreatedAfter(createdAfter).UpdatedBefore(updatedBefore).UpdatedAfter(updatedAfter).OwnerId(ownerId).Cursor(cursor).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.GetServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServices`: []GetServices200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.GetServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **[]string** | Filter for the names of services | 
 **type_** | [**[]ServiceType**](ServiceType.md) | Filter for types of services | 
 **env** | [**[]ServiceEnv**](ServiceEnv.md) | Filter for environments of services | 
 **region** | [**[]Region**](Region.md) | Filter for regions of services | 
 **suspended** | **[]string** | Filter services based on whether they&#39;re suspended or not suspended | 
 **createdBefore** | **time.Time** | Filter for services created before a certain time (specified as an ISO 8601 timestamp) | 
 **createdAfter** | **time.Time** | Filter for services created after a certain time (specified as an ISO 8601 timestamp) | 
 **updatedBefore** | **time.Time** | Filter for services updated before a certain time (specified as an ISO 8601 timestamp) | 
 **updatedAfter** | **time.Time** | Filter for services updated after a certain time (specified as an ISO 8601 timestamp) | 
 **ownerId** | **[]string** | The ID of the owner (team or personal user) whose resources should be returned | 
 **cursor** | **string** | Cursor to begin retrieving entries for this query | 
 **limit** | **float32** | Max number of items that can be returned | [default to 20]

### Return type

[**[]GetServices200ResponseInner**](GetServices200ResponseInner.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResumeService

> ResumeService(ctx, serviceId).Execute()

Resume service

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServicesApi.ResumeService(context.Background(), serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.ResumeService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeServiceRequest struct via the builder pattern


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


## ScaleService

> Deploy ScaleService(ctx, serviceId).ScaleServiceRequest(scaleServiceRequest).Execute()

Scale service to desired number of instances

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
    scaleServiceRequest := *openapiclient.NewScaleServiceRequest(int32(3)) // ScaleServiceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServicesApi.ScaleService(context.Background(), serviceId).ScaleServiceRequest(scaleServiceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.ScaleService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScaleService`: Deploy
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.ScaleService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service | 

### Other Parameters

Other parameters are passed through a pointer to a apiScaleServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scaleServiceRequest** | [**ScaleServiceRequest**](ScaleServiceRequest.md) |  | 

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


## SuspendService

> SuspendService(ctx, serviceId).Execute()

Suspend service

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServicesApi.SuspendService(context.Background(), serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.SuspendService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service | 

### Other Parameters

Other parameters are passed through a pointer to a apiSuspendServiceRequest struct via the builder pattern


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


## UpdateEnvVarsForService

> []GetEnvVarsForService200ResponseInner UpdateEnvVarsForService(ctx, serviceId).UpdateEnvVarsForServiceRequestInner(updateEnvVarsForServiceRequestInner).Execute()

Update environment variables

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
    updateEnvVarsForServiceRequestInner := []openapiclient.UpdateEnvVarsForServiceRequestInner{openapiclient.update_env_vars_for_service_request_inner{EnvVarKeyGenerateValue: openapiclient.NewEnvVarKeyGenerateValue("Key_example", "GenerateValue_example")}} // []UpdateEnvVarsForServiceRequestInner | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServicesApi.UpdateEnvVarsForService(context.Background(), serviceId).UpdateEnvVarsForServiceRequestInner(updateEnvVarsForServiceRequestInner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.UpdateEnvVarsForService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEnvVarsForService`: []GetEnvVarsForService200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.UpdateEnvVarsForService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEnvVarsForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateEnvVarsForServiceRequestInner** | [**[]UpdateEnvVarsForServiceRequestInner**](UpdateEnvVarsForServiceRequestInner.md) |  | 

### Return type

[**[]GetEnvVarsForService200ResponseInner**](GetEnvVarsForService200ResponseInner.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateService

> Service UpdateService(ctx, serviceId).ServicePATCH(servicePATCH).Execute()

Update service

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
    servicePATCH := *openapiclient.NewServicePATCH() // ServicePATCH | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServicesApi.UpdateService(context.Background(), serviceId).ServicePATCH(servicePATCH).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.UpdateService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateService`: Service
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.UpdateService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **servicePATCH** | [**ServicePATCH**](ServicePATCH.md) |  | 

### Return type

[**Service**](Service.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

