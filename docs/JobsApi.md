# \JobsApi

All URIs are relative to *https://api.render.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetJob**](JobsApi.md#GetJob) | **Get** /services/{serviceId}/jobs/{jobId} | Retrieve job
[**ListJob**](JobsApi.md#ListJob) | **Get** /services/{serviceId}/jobs | List jobs
[**PostJob**](JobsApi.md#PostJob) | **Post** /services/{serviceId}/jobs | Create job



## GetJob

> Job GetJob(ctx, serviceId, jobId).Execute()

Retrieve job

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
    jobId := "jobId_example" // string | The ID of the job

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsApi.GetJob(context.Background(), serviceId, jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.GetJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJob`: Job
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.GetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service | 
**jobId** | **string** | The ID of the job | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Job**](Job.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListJob

> []ListJob200ResponseInner ListJob(ctx, serviceId).Cursor(cursor).Limit(limit).Status(status).CreatedBefore(createdBefore).CreatedAfter(createdAfter).StartedBefore(startedBefore).StartedAfter(startedAfter).FinishedBefore(finishedBefore).FinishedAfter(finishedAfter).Execute()

List jobs

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
    status := []string{"Status_example"} // []string | Filter for the status of the job (`pending`, `running`, `succeeded`, or `failed`) (optional)
    createdBefore := time.Now() // time.Time | Filter for jobs created before a certain time (specified as an ISO 8601 timestamp) (optional)
    createdAfter := time.Now() // time.Time | Filter for jobs created after a certain time (specified as an ISO 8601 timestamp) (optional)
    startedBefore := time.Now() // time.Time | Filter for jobs started before a certain time (specified as an ISO 8601 timestamp) (optional)
    startedAfter := time.Now() // time.Time | Filter for jobs started after a certain time (specified as an ISO 8601 timestamp) (optional)
    finishedBefore := time.Now() // time.Time | Filter for jobs finished before a certain time (specified as an ISO 8601 timestamp) (optional)
    finishedAfter := time.Now() // time.Time | Filter for jobs finished after a certain time (specified as an ISO 8601 timestamp) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsApi.ListJob(context.Background(), serviceId).Cursor(cursor).Limit(limit).Status(status).CreatedBefore(createdBefore).CreatedAfter(createdAfter).StartedBefore(startedBefore).StartedAfter(startedAfter).FinishedBefore(finishedBefore).FinishedAfter(finishedAfter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.ListJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListJob`: []ListJob200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.ListJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service | 

### Other Parameters

Other parameters are passed through a pointer to a apiListJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Cursor to begin retrieving entries for this query | 
 **limit** | **float32** | Max number of items that can be returned | [default to 20]
 **status** | **[]string** | Filter for the status of the job (&#x60;pending&#x60;, &#x60;running&#x60;, &#x60;succeeded&#x60;, or &#x60;failed&#x60;) | 
 **createdBefore** | **time.Time** | Filter for jobs created before a certain time (specified as an ISO 8601 timestamp) | 
 **createdAfter** | **time.Time** | Filter for jobs created after a certain time (specified as an ISO 8601 timestamp) | 
 **startedBefore** | **time.Time** | Filter for jobs started before a certain time (specified as an ISO 8601 timestamp) | 
 **startedAfter** | **time.Time** | Filter for jobs started after a certain time (specified as an ISO 8601 timestamp) | 
 **finishedBefore** | **time.Time** | Filter for jobs finished before a certain time (specified as an ISO 8601 timestamp) | 
 **finishedAfter** | **time.Time** | Filter for jobs finished after a certain time (specified as an ISO 8601 timestamp) | 

### Return type

[**[]ListJob200ResponseInner**](ListJob200ResponseInner.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostJob

> Job PostJob(ctx, serviceId).PostJobRequest(postJobRequest).Execute()

Create job

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
    postJobRequest := *openapiclient.NewPostJobRequest("StartCommand_example") // PostJobRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsApi.PostJob(context.Background(), serviceId).PostJobRequest(postJobRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.PostJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostJob`: Job
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.PostJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postJobRequest** | [**PostJobRequest**](PostJobRequest.md) |  | 

### Return type

[**Job**](Job.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

