# \LiveStream

All URIs are relative to https://api.aiozstream.network/api

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMulticast**](LiveStream.md#AddMulticast) | **Post** /live_streams/multicast/{stream_key} | Add live stream multicast
[**CreateLiveStreamKey**](LiveStream.md#CreateLiveStreamKey) | **Post** /live_streams | Create live stream key
[**CreateStreaming**](LiveStream.md#CreateStreaming) | **Post** /live_streams/{id}/streamings | Create a new live stream media
[**DeleteLiveStreamKey**](LiveStream.md#DeleteLiveStreamKey) | **Delete** /live_streams/{id} | Delete live stream key
[**DeleteMulticast**](LiveStream.md#DeleteMulticast) | **Delete** /live_streams/multicast/{stream_key} | Delete live stream multicast
[**DeleteStreaming**](LiveStream.md#DeleteStreaming) | **Delete** /live_streams/{id}/streamings/{stream_id} | Delete live stream video
[**GetLiveStreamKey**](LiveStream.md#GetLiveStreamKey) | **Get** /live_streams/{id} | Get live stream key
[**GetLiveStreamKeys**](LiveStream.md#GetLiveStreamKeys) | **Get** /live_streams | Get live stream key list
[**GetLiveStreamPlayerInfo**](LiveStream.md#GetLiveStreamPlayerInfo) | **Get** /live_streams/player/{id}/videos | Get live stream video public
[**GetLiveStreamVideo**](LiveStream.md#GetLiveStreamVideo) | **Get** /live_streams/{id}/video | Get live stream video
[**GetMedias**](LiveStream.md#GetMedias) | **Post** /live_streams/{id}/videos | Get live stream media
[**GetMulticastByStreamKey**](LiveStream.md#GetMulticastByStreamKey) | **Get** /live_streams/multicast/{stream_key} | Get live stream multicast by stream key
[**GetStatisticByStreamMediaId**](LiveStream.md#GetStatisticByStreamMediaId) | **Get** /live_streams/statistic/{stream_media_id} | Get live stream statistic by stream media id
[**GetStreaming**](LiveStream.md#GetStreaming) | **Get** /live_streams/{id}/streamings/{stream_id} | Get live stream media streaming
[**GetStreamings**](LiveStream.md#GetStreamings) | **Get** /live_streams/{id}/streamings | Get live stream media streamings
[**UpdateLiveStreamKey**](LiveStream.md#UpdateLiveStreamKey) | **Put** /live_streams/{id} | Update live stream key
[**UpdateMedia**](LiveStream.md#UpdateMedia) | **Put** /live_streams/{id}/streamings | Update live stream media



## AddMulticast

> AddMulticast(streamKey string, data UpsertLiveStreamMulticastInput) (*GetLiveStreamMulticastResponse, error)

> AddMulticastWithContext(ctx context.Context, streamKey string, data UpsertLiveStreamMulticastInput) (*GetLiveStreamMulticastResponse, error)


Add live stream multicast



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
        
    streamKey := "streamKey_example" // string | Live stream key. Use uuid
    data := *aiozstreamsdk.NewUpsertLiveStreamMulticastInput() // UpsertLiveStreamMulticastInput | data

    
    res, err := client.LiveStream.AddMulticast(streamKey, data)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStream.AddMulticast``: %v\n", err)
    }
    // response from `AddMulticast`: GetLiveStreamMulticastResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `LiveStream.AddMulticast`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**streamKey** | **string** | Live stream key. Use uuid | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**data** | [**UpsertLiveStreamMulticastInput**](UpsertLiveStreamMulticastInput.md) | data | 

### Return type

[**GetLiveStreamMulticastResponse**](GetLiveStreamMulticastResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLiveStreamKey

> CreateLiveStreamKey(input CreateLiveStreamKeyRequest) (*CreateLiveStreamKeyResponse, error)

> CreateLiveStreamKeyWithContext(ctx context.Context, input CreateLiveStreamKeyRequest) (*CreateLiveStreamKeyResponse, error)


Create live stream key



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
        
    input := *aiozstreamsdk.NewCreateLiveStreamKeyRequest() // CreateLiveStreamKeyRequest | CreateLiveStreamKeyRequest

    
    res, err := client.LiveStream.CreateLiveStreamKey(input)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStream.CreateLiveStreamKey``: %v\n", err)
    }
    // response from `CreateLiveStreamKey`: CreateLiveStreamKeyResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `LiveStream.CreateLiveStreamKey`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**input** | [**CreateLiveStreamKeyRequest**](CreateLiveStreamKeyRequest.md) | CreateLiveStreamKeyRequest | 

### Return type

[**CreateLiveStreamKeyResponse**](CreateLiveStreamKeyResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStreaming

> CreateStreaming(id string, input CreateStreamingRequest) (*CreateStreamingResponse, error)

> CreateStreamingWithContext(ctx context.Context, id string, input CreateStreamingRequest) (*CreateStreamingResponse, error)


Create a new live stream media



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
        
    id := "id_example" // string | Live stream key ID
    input := *aiozstreamsdk.NewCreateStreamingRequest() // CreateStreamingRequest | CreateStreamingRequest

    
    res, err := client.LiveStream.CreateStreaming(id, input)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStream.CreateStreaming``: %v\n", err)
    }
    // response from `CreateStreaming`: CreateStreamingResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `LiveStream.CreateStreaming`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Live stream key ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**input** | [**CreateStreamingRequest**](CreateStreamingRequest.md) | CreateStreamingRequest | 

### Return type

[**CreateStreamingResponse**](CreateStreamingResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLiveStreamKey

> DeleteLiveStreamKey(id string) (*ResponseSuccess, error)

> DeleteLiveStreamKeyWithContext(ctx context.Context, id string) (*ResponseSuccess, error)


Delete live stream key



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
        
    id := "id_example" // string | Live stream key ID

    
    res, err := client.LiveStream.DeleteLiveStreamKey(id)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStream.DeleteLiveStreamKey``: %v\n", err)
    }
    // response from `DeleteLiveStreamKey`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `LiveStream.DeleteLiveStreamKey`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Live stream key ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMulticast

> DeleteMulticast(streamKey string) (*ResponseSuccess, error)

> DeleteMulticastWithContext(ctx context.Context, streamKey string) (*ResponseSuccess, error)


Delete live stream multicast



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
        
    streamKey := "streamKey_example" // string | Live stream key. UUID string format

    
    res, err := client.LiveStream.DeleteMulticast(streamKey)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStream.DeleteMulticast``: %v\n", err)
    }
    // response from `DeleteMulticast`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `LiveStream.DeleteMulticast`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**streamKey** | **string** | Live stream key. UUID string format | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStreaming

> DeleteStreaming(id string, streamId string) (*ResponseSuccess, error)

> DeleteStreamingWithContext(ctx context.Context, id string, streamId string) (*ResponseSuccess, error)


Delete live stream video



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
        
    id := "id_example" // string | Live stream key ID
    streamId := "streamId_example" // string | Streaming ID

    
    res, err := client.LiveStream.DeleteStreaming(id, streamId)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStream.DeleteStreaming``: %v\n", err)
    }
    // response from `DeleteStreaming`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `LiveStream.DeleteStreaming`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Live stream key ID | 
**streamId** | **string** | Streaming ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLiveStreamKey

> GetLiveStreamKey(id string) (*GetLiveStreamKeyResponse, error)

> GetLiveStreamKeyWithContext(ctx context.Context, id string) (*GetLiveStreamKeyResponse, error)


Get live stream key



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
        
    id := "id_example" // string | ID

    
    res, err := client.LiveStream.GetLiveStreamKey(id)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStream.GetLiveStreamKey``: %v\n", err)
    }
    // response from `GetLiveStreamKey`: GetLiveStreamKeyResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `LiveStream.GetLiveStreamKey`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**GetLiveStreamKeyResponse**](GetLiveStreamKeyResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLiveStreamKeys

> GetLiveStreamKeys(r LiveStreamApiGetLiveStreamKeysRequest) (*GetLiveStreamKeysListResponse, error)


> GetLiveStreamKeysWithContext(ctx context.Context, r LiveStreamApiGetLiveStreamKeysRequest) (*GetLiveStreamKeysListResponse, error)



Get live stream key list



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
    req := aiozstreamsdk.LiveStreamApiGetLiveStreamKeysRequest{}
    
    req.Search("search_example") // string | only support search by name
    req.SortBy("sortBy_example") // string | sort by (default to "created_at")
    req.OrderBy("orderBy_example") // string | allowed: asc, desc. Default: asc (default to "asc")
    req.Offset(int32(56)) // int32 | offset, allowed values greater than or equal to 0.
    req.Limit(int32(56)) // int32 | results per page.

    res, err := client.LiveStream.GetLiveStreamKeys(req)
    

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStream.GetLiveStreamKeys``: %v\n", err)
    }
    // response from `GetLiveStreamKeys`: GetLiveStreamKeysListResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `LiveStream.GetLiveStreamKeys`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**search** | **string** | only support search by name | 
**sortBy** | **string** | sort by | [default to &quot;created_at&quot;]
**orderBy** | **string** | allowed: asc, desc. Default: asc | [default to &quot;asc&quot;]
**offset** | **int32** | offset, allowed values greater than or equal to 0. | 
**limit** | **int32** | results per page. | 

### Return type

[**GetLiveStreamKeysListResponse**](GetLiveStreamKeysListResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLiveStreamPlayerInfo

> GetLiveStreamPlayerInfo(id string) (*GetLiveStreamVideoPublicResponse, error)

> GetLiveStreamPlayerInfoWithContext(ctx context.Context, id string) (*GetLiveStreamVideoPublicResponse, error)


Get live stream video public



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
        
    id := "id_example" // string | Live stream key ID

    
    res, err := client.LiveStream.GetLiveStreamPlayerInfo(id)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStream.GetLiveStreamPlayerInfo``: %v\n", err)
    }
    // response from `GetLiveStreamPlayerInfo`: GetLiveStreamVideoPublicResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `LiveStream.GetLiveStreamPlayerInfo`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Live stream key ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**GetLiveStreamVideoPublicResponse**](GetLiveStreamVideoPublicResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLiveStreamVideo

> GetLiveStreamVideo(id string) (*GetLiveStreamVideoResponse, error)

> GetLiveStreamVideoWithContext(ctx context.Context, id string) (*GetLiveStreamVideoResponse, error)


Get live stream video



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
        
    id := "id_example" // string | Live stream video ID

    
    res, err := client.LiveStream.GetLiveStreamVideo(id)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStream.GetLiveStreamVideo``: %v\n", err)
    }
    // response from `GetLiveStreamVideo`: GetLiveStreamVideoResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `LiveStream.GetLiveStreamVideo`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Live stream video ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**GetLiveStreamVideoResponse**](GetLiveStreamVideoResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMedias

> GetMedias(id string, data GetLiveStreamMediasRequest) (*GetLiveStreamMediasResponse, error)

> GetMediasWithContext(ctx context.Context, id string, data GetLiveStreamMediasRequest) (*GetLiveStreamMediasResponse, error)


Get live stream media



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
        
    id := "id_example" // string | Live stream key ID
    data := *aiozstreamsdk.NewGetLiveStreamMediasRequest() // GetLiveStreamMediasRequest | data

    
    res, err := client.LiveStream.GetMedias(id, data)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStream.GetMedias``: %v\n", err)
    }
    // response from `GetMedias`: GetLiveStreamMediasResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `LiveStream.GetMedias`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Live stream key ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**data** | [**GetLiveStreamMediasRequest**](GetLiveStreamMediasRequest.md) | data | 

### Return type

[**GetLiveStreamMediasResponse**](GetLiveStreamMediasResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMulticastByStreamKey

> GetMulticastByStreamKey(streamKey string) (*GetLiveStreamMulticastResponse, error)

> GetMulticastByStreamKeyWithContext(ctx context.Context, streamKey string) (*GetLiveStreamMulticastResponse, error)


Get live stream multicast by stream key



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
        
    streamKey := "streamKey_example" // string | Live stream key. UUID string format

    
    res, err := client.LiveStream.GetMulticastByStreamKey(streamKey)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStream.GetMulticastByStreamKey``: %v\n", err)
    }
    // response from `GetMulticastByStreamKey`: GetLiveStreamMulticastResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `LiveStream.GetMulticastByStreamKey`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**streamKey** | **string** | Live stream key. UUID string format | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**GetLiveStreamMulticastResponse**](GetLiveStreamMulticastResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatisticByStreamMediaId

> GetStatisticByStreamMediaId(streamMediaId string) (*GetLiveStreamStatisticResponse, error)

> GetStatisticByStreamMediaIdWithContext(ctx context.Context, streamMediaId string) (*GetLiveStreamStatisticResponse, error)


Get live stream statistic by stream media id



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
        
    streamMediaId := "streamMediaId_example" // string | Live stream media ID

    
    res, err := client.LiveStream.GetStatisticByStreamMediaId(streamMediaId)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStream.GetStatisticByStreamMediaId``: %v\n", err)
    }
    // response from `GetStatisticByStreamMediaId`: GetLiveStreamStatisticResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `LiveStream.GetStatisticByStreamMediaId`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**streamMediaId** | **string** | Live stream media ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**GetLiveStreamStatisticResponse**](GetLiveStreamStatisticResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStreaming

> GetStreaming(id string, streamId string) (*GetStreamingResponse, error)

> GetStreamingWithContext(ctx context.Context, id string, streamId string) (*GetStreamingResponse, error)


Get live stream media streaming



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
        
    id := "id_example" // string | Live stream key ID
    streamId := "streamId_example" // string | Stream ID

    
    res, err := client.LiveStream.GetStreaming(id, streamId)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStream.GetStreaming``: %v\n", err)
    }
    // response from `GetStreaming`: GetStreamingResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `LiveStream.GetStreaming`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Live stream key ID | 
**streamId** | **string** | Stream ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**GetStreamingResponse**](GetStreamingResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStreamings

> GetStreamings(id string) (*GetStreamingsResponse, error)

> GetStreamingsWithContext(ctx context.Context, id string) (*GetStreamingsResponse, error)


Get live stream media streamings



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
        
    id := "id_example" // string | Live stream key ID

    
    res, err := client.LiveStream.GetStreamings(id)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStream.GetStreamings``: %v\n", err)
    }
    // response from `GetStreamings`: GetStreamingsResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `LiveStream.GetStreamings`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Live stream key ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**GetStreamingsResponse**](GetStreamingsResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLiveStreamKey

> UpdateLiveStreamKey(id string, input UpdateLiveStreamKeyRequest) (*UpdateLiveStreamKeyResponse, error)

> UpdateLiveStreamKeyWithContext(ctx context.Context, id string, input UpdateLiveStreamKeyRequest) (*UpdateLiveStreamKeyResponse, error)


Update live stream key



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
        
    id := "id_example" // string | Live stream key ID
    input := *aiozstreamsdk.NewUpdateLiveStreamKeyRequest() // UpdateLiveStreamKeyRequest | UpdateLiveStreamKeyRequest

    
    res, err := client.LiveStream.UpdateLiveStreamKey(id, input)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStream.UpdateLiveStreamKey``: %v\n", err)
    }
    // response from `UpdateLiveStreamKey`: UpdateLiveStreamKeyResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `LiveStream.UpdateLiveStreamKey`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Live stream key ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**input** | [**UpdateLiveStreamKeyRequest**](UpdateLiveStreamKeyRequest.md) | UpdateLiveStreamKeyRequest | 

### Return type

[**UpdateLiveStreamKeyResponse**](UpdateLiveStreamKeyResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMedia

> UpdateMedia(id string, data UpdateLiveStreamMediaRequest) (*ResponseSuccess, error)

> UpdateMediaWithContext(ctx context.Context, id string, data UpdateLiveStreamMediaRequest) (*ResponseSuccess, error)


Update live stream media



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
        
    id := "id_example" // string | Live stream key ID
    data := *aiozstreamsdk.NewUpdateLiveStreamMediaRequest() // UpdateLiveStreamMediaRequest | data

    
    res, err := client.LiveStream.UpdateMedia(id, data)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStream.UpdateMedia``: %v\n", err)
    }
    // response from `UpdateMedia`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `LiveStream.UpdateMedia`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Live stream key ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**data** | [**UpdateLiveStreamMediaRequest**](UpdateLiveStreamMediaRequest.md) | data | 

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

