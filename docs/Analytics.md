# \Analytics

All URIs are relative to https://api.aiozstream.network/api

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAggregatedMetrics**](Analytics.md#GetAggregatedMetrics) | **Post** /analytics/metrics/data/{metric}/{aggregation} | Get aggregated metrics
[**GetBreakdownMetrics**](Analytics.md#GetBreakdownMetrics) | **Post** /analytics/metrics/bucket/{metric}/{breakdown} | Get breakdown metrics
[**GetDataUsage**](Analytics.md#GetDataUsage) | **Get** /analytics/data | Get data usage
[**GetOvertimeMetrics**](Analytics.md#GetOvertimeMetrics) | **Post** /analytics/metrics/timeseries/{metric}/{interval} | Get overtime metrics
[**GetStatisticMedias**](Analytics.md#GetStatisticMedias) | **Get** /analytics/media | Get statistic media



## GetAggregatedMetrics

> GetAggregatedMetrics(metric string, aggregation string, request InternalControllersGetAggreatedMetricsMetricsRequest) (*GetAggregatedMetricsResponse, error)

> GetAggregatedMetricsWithContext(ctx context.Context, metric string, aggregation string, request InternalControllersGetAggreatedMetricsMetricsRequest) (*GetAggregatedMetricsResponse, error)


Get aggregated metrics



### Example

```go
package main

import (
    "context"
    "fmt"
    "encoding/json"
    "os"
    aiozstreamsdk "github.com/AIOZNetwork/aioz-stream-go-client"
)

func main() {
    // create a new client
    apiCreds := aiozstreamsdk.AuthCredentials{
		SecretKey: "YOUR_SECRET_KEY",
		PublicKey: "YOUR_PUBLIC_KEY",
    }
    client := aiozstreamsdk.ClientBuilder(apiCreds).Build()
        
    metric := "metric_example" // string | metric name
    aggregation := "aggregation_example" // string | aggregation method
    request := *aiozstreamsdk.NewInternalControllersGetAggreatedMetricsMetricsRequest() // InternalControllersGetAggreatedMetricsMetricsRequest | query parameters

    
    res, err := client.Analytics.GetAggregatedMetrics(metric, aggregation, request)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Analytics.GetAggregatedMetrics``: %v\n", err)
    }
    // response from `GetAggregatedMetrics`: GetAggregatedMetricsResponse
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
**metric** | **string** | metric name | 
**aggregation** | **string** | aggregation method | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**request** | [**InternalControllersGetAggreatedMetricsMetricsRequest**](InternalControllersGetAggreatedMetricsMetricsRequest.md) | query parameters | 

### Return type

[**GetAggregatedMetricsResponse**](GetAggregatedMetricsResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBreakdownMetrics

> GetBreakdownMetrics(metric string, breakdown string, request GetBreakdownMetricsRequest) (*GetBreakdownMetricsResponse, error)

> GetBreakdownMetricsWithContext(ctx context.Context, metric string, breakdown string, request GetBreakdownMetricsRequest) (*GetBreakdownMetricsResponse, error)


Get breakdown metrics



### Example

```go
package main

import (
    "context"
    "fmt"
    "encoding/json"
    "os"
    aiozstreamsdk "github.com/AIOZNetwork/aioz-stream-go-client"
)

func main() {
    // create a new client
    apiCreds := aiozstreamsdk.AuthCredentials{
		SecretKey: "YOUR_SECRET_KEY",
		PublicKey: "YOUR_PUBLIC_KEY",
    }
    client := aiozstreamsdk.ClientBuilder(apiCreds).Build()
        
    metric := "metric_example" // string | metric name
    breakdown := "breakdown_example" // string | breakdown dimension
    request := *aiozstreamsdk.NewGetBreakdownMetricsRequest() // GetBreakdownMetricsRequest | query parameters

    
    res, err := client.Analytics.GetBreakdownMetrics(metric, breakdown, request)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Analytics.GetBreakdownMetrics``: %v\n", err)
    }
    // response from `GetBreakdownMetrics`: GetBreakdownMetricsResponse
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
**metric** | **string** | metric name | 
**breakdown** | **string** | breakdown dimension | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**request** | [**GetBreakdownMetricsRequest**](GetBreakdownMetricsRequest.md) | query parameters | 

### Return type

[**GetBreakdownMetricsResponse**](GetBreakdownMetricsResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataUsage

> GetDataUsage(r AnalyticsApiGetDataUsageRequest) (*GetDataUsageResponse, error)


> GetDataUsageWithContext(ctx context.Context, r AnalyticsApiGetDataUsageRequest) (*GetDataUsageResponse, error)



Get data usage



### Example

```go
package main

import (
    "context"
    "fmt"
    "encoding/json"
    "os"
    aiozstreamsdk "github.com/AIOZNetwork/aioz-stream-go-client"
)

func main() {
    // create a new client
    apiCreds := aiozstreamsdk.AuthCredentials{
		SecretKey: "YOUR_SECRET_KEY",
		PublicKey: "YOUR_PUBLIC_KEY",
    }
    client := aiozstreamsdk.ClientBuilder(apiCreds).Build()
    req := aiozstreamsdk.AnalyticsApiGetDataUsageRequest{}
    
    req.From(int32(56)) // int32 | start time (unix timestamp)
    req.To(int32(56)) // int32 | end time (unix timestamp)
    req.Interval("interval_example") // string | time interval
    req.Offset(int32(56)) // int32 | offset, allowed values greater than or equal to 0. Default(0) (default to 0)
    req.Limit(int32(56)) // int32 | results per page. Allowed values 1-100, default is 25 (default to 25)

    res, err := client.Analytics.GetDataUsage(req)
    

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Analytics.GetDataUsage``: %v\n", err)
    }
    // response from `GetDataUsage`: GetDataUsageResponse
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
**from** | **int32** | start time (unix timestamp) | 
**to** | **int32** | end time (unix timestamp) | 
**interval** | **string** | time interval | 
**offset** | **int32** | offset, allowed values greater than or equal to 0. Default(0) | [default to 0]
**limit** | **int32** | results per page. Allowed values 1-100, default is 25 | [default to 25]

### Return type

[**GetDataUsageResponse**](GetDataUsageResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOvertimeMetrics

> GetOvertimeMetrics(metric string, interval string, request GetBreakdownMetricsRequest) (*GetOvertimeMetricsResponse, error)

> GetOvertimeMetricsWithContext(ctx context.Context, metric string, interval string, request GetBreakdownMetricsRequest) (*GetOvertimeMetricsResponse, error)


Get overtime metrics



### Example

```go
package main

import (
    "context"
    "fmt"
    "encoding/json"
    "os"
    aiozstreamsdk "github.com/AIOZNetwork/aioz-stream-go-client"
)

func main() {
    // create a new client
    apiCreds := aiozstreamsdk.AuthCredentials{
		SecretKey: "YOUR_SECRET_KEY",
		PublicKey: "YOUR_PUBLIC_KEY",
    }
    client := aiozstreamsdk.ClientBuilder(apiCreds).Build()
        
    metric := "metric_example" // string | metric name
    interval := "interval_example" // string | time interval
    request := *aiozstreamsdk.NewGetBreakdownMetricsRequest() // GetBreakdownMetricsRequest | query parameters

    
    res, err := client.Analytics.GetOvertimeMetrics(metric, interval, request)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Analytics.GetOvertimeMetrics``: %v\n", err)
    }
    // response from `GetOvertimeMetrics`: GetOvertimeMetricsResponse
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
**metric** | **string** | metric name | 
**interval** | **string** | time interval | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**request** | [**GetBreakdownMetricsRequest**](GetBreakdownMetricsRequest.md) | query parameters | 

### Return type

[**GetOvertimeMetricsResponse**](GetOvertimeMetricsResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatisticMedias

> GetStatisticMedias(r AnalyticsApiGetStatisticMediasRequest) (*GetStatisticMediasResponse, error)


> GetStatisticMediasWithContext(ctx context.Context, r AnalyticsApiGetStatisticMediasRequest) (*GetStatisticMediasResponse, error)



Get statistic media



### Example

```go
package main

import (
    "context"
    "fmt"
    "encoding/json"
    "os"
    aiozstreamsdk "github.com/AIOZNetwork/aioz-stream-go-client"
)

func main() {
    // create a new client
    apiCreds := aiozstreamsdk.AuthCredentials{
		SecretKey: "YOUR_SECRET_KEY",
		PublicKey: "YOUR_PUBLIC_KEY",
    }
    client := aiozstreamsdk.ClientBuilder(apiCreds).Build()
    req := aiozstreamsdk.AnalyticsApiGetStatisticMediasRequest{}
    
    req.From(int32(56)) // int32 | start time
    req.To(int32(56)) // int32 | end time
    req.Type_("type__example") // string | media type
    req.Offset(int32(56)) // int32 | offset, allowed values greater than or equal to 0. Default(0) (default to 0)
    req.Limit(int32(56)) // int32 | results per page. Allowed values 1-100, default is 25 (default to 25)

    res, err := client.Analytics.GetStatisticMedias(req)
    

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Analytics.GetStatisticMedias``: %v\n", err)
    }
    // response from `GetStatisticMedias`: GetStatisticMediasResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Analytics.GetStatisticMedias`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**from** | **int32** | start time | 
**to** | **int32** | end time | 
**type_** | **string** | media type | 
**offset** | **int32** | offset, allowed values greater than or equal to 0. Default(0) | [default to 0]
**limit** | **int32** | results per page. Allowed values 1-100, default is 25 | [default to 25]

### Return type

[**GetStatisticMediasResponse**](GetStatisticMediasResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

