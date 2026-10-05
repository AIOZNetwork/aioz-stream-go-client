# \Analytics

All URIs are relative to https://api.aiozstream.network/api

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAggregatedMetrics**](Analytics.md#GetAggregatedMetrics) | **Post** /analytics/metrics/data/{metric}/{aggregation} | Aggregate one metric
[**GetBreakdownMetrics**](Analytics.md#GetBreakdownMetrics) | **Post** /analytics/metrics/bucket/{metric}/{breakdown} | Bucket one metric by dimension
[**GetDataUsage**](Analytics.md#GetDataUsage) | **Get** /analytics/data | Delivery volume over time
[**GetOvertimeMetrics**](Analytics.md#GetOvertimeMetrics) | **Post** /analytics/metrics/timeseries/{metric}/{interval} | Bucket one metric by time



## GetAggregatedMetrics

> GetAggregatedMetrics(metric string, aggregation string, metricsRequest MetricsRequest) (*AggregatedMetricsResponse, error)

> GetAggregatedMetricsWithContext(ctx context.Context, metric string, aggregation string, metricsRequest MetricsRequest) (*AggregatedMetricsResponse, error)


Aggregate one metric



### Example

```go
package main

import (
    "context"
    "fmt"
    "encoding/json"
    "os"
    aiozstreamsdk "github.com/AIOZNetwork/aioz-stream-go-client/v3"
)

func main() {
    // create a new client
    apiCreds := aiozstreamsdk.AuthCredentials{
		SecretKey: "YOUR_SECRET_KEY",
		PublicKey: "YOUR_PUBLIC_KEY",
    }
    client := aiozstreamsdk.ClientBuilder(apiCreds).Build()
        
    metric := "metric_example" // string | Metric
    aggregation := "aggregation_example" // string | Aggregation
    metricsRequest := *aiozstreamsdk.NewMetricsRequest() // MetricsRequest | Window and filter

    
    res, err := client.Analytics.GetAggregatedMetrics(metric, aggregation, metricsRequest)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Analytics.GetAggregatedMetrics``: %v\n", err)
    }
    // response from `GetAggregatedMetrics`: AggregatedMetricsResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Analytics.GetAggregatedMetrics`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**metric** | **string** | Metric | 
**aggregation** | **string** | Aggregation | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**metricsRequest** | [**MetricsRequest**](MetricsRequest.md) | Window and filter | 

### Return type

[**AggregatedMetricsResponse**](AggregatedMetricsResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBreakdownMetrics

> GetBreakdownMetrics(metric string, breakdown string, metricsRequest MetricsRequest) (*MetricsPageResponse, error)

> GetBreakdownMetricsWithContext(ctx context.Context, metric string, breakdown string, metricsRequest MetricsRequest) (*MetricsPageResponse, error)


Bucket one metric by dimension



### Example

```go
package main

import (
    "context"
    "fmt"
    "encoding/json"
    "os"
    aiozstreamsdk "github.com/AIOZNetwork/aioz-stream-go-client/v3"
)

func main() {
    // create a new client
    apiCreds := aiozstreamsdk.AuthCredentials{
		SecretKey: "YOUR_SECRET_KEY",
		PublicKey: "YOUR_PUBLIC_KEY",
    }
    client := aiozstreamsdk.ClientBuilder(apiCreds).Build()
        
    metric := "metric_example" // string | Metric
    breakdown := "breakdown_example" // string | Dimension
    metricsRequest := *aiozstreamsdk.NewMetricsRequest() // MetricsRequest | Window, filter and paging

    
    res, err := client.Analytics.GetBreakdownMetrics(metric, breakdown, metricsRequest)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Analytics.GetBreakdownMetrics``: %v\n", err)
    }
    // response from `GetBreakdownMetrics`: MetricsPageResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Analytics.GetBreakdownMetrics`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**metric** | **string** | Metric | 
**breakdown** | **string** | Dimension | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**metricsRequest** | [**MetricsRequest**](MetricsRequest.md) | Window, filter and paging | 

### Return type

[**MetricsPageResponse**](MetricsPageResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataUsage

> GetDataUsage(r AnalyticsApiGetDataUsageRequest) (*DataUsageResponse, error)


> GetDataUsageWithContext(ctx context.Context, r AnalyticsApiGetDataUsageRequest) (*DataUsageResponse, error)



Delivery volume over time



### Example

```go
package main

import (
    "context"
    "fmt"
    "encoding/json"
    "os"
    aiozstreamsdk "github.com/AIOZNetwork/aioz-stream-go-client/v3"
)

func main() {
    // create a new client
    apiCreds := aiozstreamsdk.AuthCredentials{
		SecretKey: "YOUR_SECRET_KEY",
		PublicKey: "YOUR_PUBLIC_KEY",
    }
    client := aiozstreamsdk.ClientBuilder(apiCreds).Build()
    req := aiozstreamsdk.AnalyticsApiGetDataUsageRequest{}
    
    req.From(int64(789)) // int64 | 
    req.Interval("interval_example") // string | 
    req.Limit(int64(789)) // int64 |  (default to 25)
    req.Offset(int64(789)) // int64 | 
    req.To(int64(789)) // int64 | 

    res, err := client.Analytics.GetDataUsage(req)
    

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Analytics.GetDataUsage``: %v\n", err)
    }
    // response from `GetDataUsage`: DataUsageResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Analytics.GetDataUsage`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**from** | **int64** |  | 
**interval** | **string** |  | 
**limit** | **int64** |  | [default to 25]
**offset** | **int64** |  | 
**to** | **int64** |  | 

### Return type

[**DataUsageResponse**](DataUsageResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOvertimeMetrics

> GetOvertimeMetrics(metric string, interval string, metricsRequest MetricsRequest) (*MetricsPageResponse, error)

> GetOvertimeMetricsWithContext(ctx context.Context, metric string, interval string, metricsRequest MetricsRequest) (*MetricsPageResponse, error)


Bucket one metric by time



### Example

```go
package main

import (
    "context"
    "fmt"
    "encoding/json"
    "os"
    aiozstreamsdk "github.com/AIOZNetwork/aioz-stream-go-client/v3"
)

func main() {
    // create a new client
    apiCreds := aiozstreamsdk.AuthCredentials{
		SecretKey: "YOUR_SECRET_KEY",
		PublicKey: "YOUR_PUBLIC_KEY",
    }
    client := aiozstreamsdk.ClientBuilder(apiCreds).Build()
        
    metric := "metric_example" // string | Metric
    interval := "interval_example" // string | Bucket size
    metricsRequest := *aiozstreamsdk.NewMetricsRequest() // MetricsRequest | Window, filter and paging

    
    res, err := client.Analytics.GetOvertimeMetrics(metric, interval, metricsRequest)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Analytics.GetOvertimeMetrics``: %v\n", err)
    }
    // response from `GetOvertimeMetrics`: MetricsPageResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Analytics.GetOvertimeMetrics`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**metric** | **string** | Metric | 
**interval** | **string** | Bucket size | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**metricsRequest** | [**MetricsRequest**](MetricsRequest.md) | Window, filter and paging | 

### Return type

[**MetricsPageResponse**](MetricsPageResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

