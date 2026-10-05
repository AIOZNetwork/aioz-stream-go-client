# \LiveStream

All URIs are relative to https://api.aiozstream.network/api

Method | HTTP request | Description
------------- | ------------- | -------------
[**UploadThumbnail**](LiveStream.md#UploadThumbnail) | **Post** /live_streams/{id}/thumbnail | Upload live stream media thumbnail
[**DeleteThumbnail**](LiveStream.md#DeleteThumbnail) | **Delete** /live_streams/{id}/thumbnail | Delete live stream media thumbnail
[**AddLiveStreamMulticasts**](LiveStream.md#AddLiveStreamMulticasts) | **Post** /live_streams/multicast/{stream_key} | Add live stream multicast
[**CreateLiveStreamKey**](LiveStream.md#CreateLiveStreamKey) | **Post** /live_streams | Create live stream key
[**CreateStreaming**](LiveStream.md#CreateStreaming) | **Post** /live_streams/{id}/streamings | Create a new live stream media
[**DeleteLiveStreamKey**](LiveStream.md#DeleteLiveStreamKey) | **Delete** /live_streams/{id} | Delete live stream key
[**DeleteLiveStreamMulticast**](LiveStream.md#DeleteLiveStreamMulticast) | **Delete** /live_streams/multicast/{stream_key} | Delete live stream multicast
[**GetLiveStreamKey**](LiveStream.md#GetLiveStreamKey) | **Get** /live_streams/{id} | Get live stream key
[**GetLiveStreamKeys**](LiveStream.md#GetLiveStreamKeys) | **Get** /live_streams | Get live stream key list
[**GetLiveStreamMedia**](LiveStream.md#GetLiveStreamMedia) | **Get** /live_streams/{id}/media | Get live stream media
[**GetLiveStreamMedias**](LiveStream.md#GetLiveStreamMedias) | **Post** /live_streams/{id}/media | Get live stream media
[**GetLiveStreamMulticastByStreamKey**](LiveStream.md#GetLiveStreamMulticastByStreamKey) | **Get** /live_streams/multicast/{stream_key} | Get live stream multicast by stream key
[**GetLiveStreamPlayerInfo**](LiveStream.md#GetLiveStreamPlayerInfo) | **Get** /live_streams/player/{id}/media | Get live stream media public
[**GetLiveStreamStatisticByStreamMediaId**](LiveStream.md#GetLiveStreamStatisticByStreamMediaId) | **Get** /live_streams/statistic/{stream_media_id} | Get live stream statistic by stream media id
[**GetLiveStreamUsage**](LiveStream.md#GetLiveStreamUsage) | **Get** /live_streams/usage/{stream_id} | Get usage details for a specific live stream
[**GetStreaming**](LiveStream.md#GetStreaming) | **Get** /live_streams/{id}/streamings/{stream_id} | Get live stream media streaming
[**GetStreamings**](LiveStream.md#GetStreamings) | **Get** /live_streams/{id}/streamings | Get live stream media streamings
[**GetUserStreamsUsageDetail**](LiveStream.md#GetUserStreamsUsageDetail) | **Get** /live_streams/usage/streams | Get paginated list of streams with usage details for current user
[**UpdateLiveStreamKey**](LiveStream.md#UpdateLiveStreamKey) | **Put** /live_streams/{id} | Update live stream key
[**UpdateLiveStreamMedia**](LiveStream.md#UpdateLiveStreamMedia) | **Put** /live_streams/{id}/streamings | Update live stream media



## UploadThumbnail

> UploadThumbnailFile(id string, file *os.File) (*ResponseSuccess, error)
> UploadThumbnail(id string, fileName string, fileReader io.Reader)
> UploadThumbnailFileWithContext(ctx context.Context, id string, file *os.File) (*ResponseSuccess, error)
> UploadThumbnailWithContext(ctx context.Context, id string, fileName string, fileReader io.Reader)

Upload live stream media thumbnail

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
        
    id := "id_example" // string | live stream media's id
    file := os.NewFile(1234, "some_file") // *os.File | file media to be uploaded

    
    res, err := client.LiveStream.UploadThumbnailFile(id, file)

    // you can also use a Reader instead of a File:
    // we recommend using Reader instead!
    // client.LiveStream.UploadThumbnail(id, fileName, fileReader)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStream.UploadThumbnail``: %v\n", err)
    }
    // response from `UploadThumbnail`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `LiveStream.UploadThumbnail`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | live stream media&#39;s id | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**file** | ***os.File** | file media to be uploaded | 

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteThumbnail

> DeleteThumbnail(id string) (*ResponseSuccess, error)

> DeleteThumbnailWithContext(ctx context.Context, id string) (*ResponseSuccess, error)


Delete live stream media thumbnail

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
        
    id := "id_example" // string | live stream media's id

    
    res, err := client.LiveStream.DeleteThumbnail(id)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStream.DeleteThumbnail``: %v\n", err)
    }
    // response from `DeleteThumbnail`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `LiveStream.DeleteThumbnail`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | live stream media&#39;s id | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddLiveStreamMulticasts

> AddLiveStreamMulticasts(streamKey string, upsertLiveStreamMulticastInput UpsertLiveStreamMulticastInput) (*GetLiveStreamMulticastResponse, error)

> AddLiveStreamMulticastsWithContext(ctx context.Context, streamKey string, upsertLiveStreamMulticastInput UpsertLiveStreamMulticastInput) (*GetLiveStreamMulticastResponse, error)


Add live stream multicast



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
        
    streamKey := "streamKey_example" // string | Live stream key. Use uuid
    upsertLiveStreamMulticastInput := *aiozstreamsdk.NewUpsertLiveStreamMulticastInput() // UpsertLiveStreamMulticastInput | data

    
    res, err := client.LiveStream.AddLiveStreamMulticasts(streamKey, upsertLiveStreamMulticastInput)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStream.AddLiveStreamMulticasts``: %v\n", err)
    }
    // response from `AddLiveStreamMulticasts`: GetLiveStreamMulticastResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `LiveStream.AddLiveStreamMulticasts`")
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
**upsertLiveStreamMulticastInput** | [**UpsertLiveStreamMulticastInput**](UpsertLiveStreamMulticastInput.md) | data | 

### Return type

[**GetLiveStreamMulticastResponse**](GetLiveStreamMulticastResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLiveStreamKey

> CreateLiveStreamKey(createLiveStreamKeyRequest CreateLiveStreamKeyRequest) (*CreateLiveStreamKeyResponse, error)

> CreateLiveStreamKeyWithContext(ctx context.Context, createLiveStreamKeyRequest CreateLiveStreamKeyRequest) (*CreateLiveStreamKeyResponse, error)


Create live stream key



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
        
    createLiveStreamKeyRequest := *aiozstreamsdk.NewCreateLiveStreamKeyRequest() // CreateLiveStreamKeyRequest | CreateLiveStreamKeyRequest

    
    res, err := client.LiveStream.CreateLiveStreamKey(createLiveStreamKeyRequest)

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
**createLiveStreamKeyRequest** | [**CreateLiveStreamKeyRequest**](CreateLiveStreamKeyRequest.md) | CreateLiveStreamKeyRequest | 

### Return type

[**CreateLiveStreamKeyResponse**](CreateLiveStreamKeyResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStreaming

> CreateStreaming(id string, createStreamingRequest CreateStreamingRequest) (*CreateStreamingResponse, error)

> CreateStreamingWithContext(ctx context.Context, id string, createStreamingRequest CreateStreamingRequest) (*CreateStreamingResponse, error)


Create a new live stream media



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
        
    id := "id_example" // string | Live stream key ID
    createStreamingRequest := *aiozstreamsdk.NewCreateStreamingRequest() // CreateStreamingRequest | CreateStreamingRequest

    
    res, err := client.LiveStream.CreateStreaming(id, createStreamingRequest)

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
**createStreamingRequest** | [**CreateStreamingRequest**](CreateStreamingRequest.md) | CreateStreamingRequest | 

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
    aiozstreamsdk "github.com/AIOZNetwork/aioz-stream-go-client/v3"
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


## DeleteLiveStreamMulticast

> DeleteLiveStreamMulticast(streamKey string) (*ResponseSuccess, error)

> DeleteLiveStreamMulticastWithContext(ctx context.Context, streamKey string) (*ResponseSuccess, error)


Delete live stream multicast



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
        
    streamKey := "streamKey_example" // string | Live stream key. UUID string format

    
    res, err := client.LiveStream.DeleteLiveStreamMulticast(streamKey)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStream.DeleteLiveStreamMulticast``: %v\n", err)
    }
    // response from `DeleteLiveStreamMulticast`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `LiveStream.DeleteLiveStreamMulticast`")
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
    aiozstreamsdk "github.com/AIOZNetwork/aioz-stream-go-client/v3"
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
    aiozstreamsdk "github.com/AIOZNetwork/aioz-stream-go-client/v3"
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
    req.Offset(int64(789)) // int64 | offset, allowed values greater than or equal to 0.
    req.Limit(int64(789)) // int64 | results per page.
    req.Type_("type__example") // string | type of media. Enums(audio, video) default(video). (default to "video")

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
**offset** | **int64** | offset, allowed values greater than or equal to 0. | 
**limit** | **int64** | results per page. | 
**type_** | **string** | type of media. Enums(audio, video) default(video). | [default to &quot;video&quot;]

### Return type

[**GetLiveStreamKeysListResponse**](GetLiveStreamKeysListResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLiveStreamMedia

> GetLiveStreamMedia(id string) (*GetLiveStreamMediaResponse, error)

> GetLiveStreamMediaWithContext(ctx context.Context, id string) (*GetLiveStreamMediaResponse, error)


Get live stream media



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
        
    id := "id_example" // string | Live stream media ID

    
    res, err := client.LiveStream.GetLiveStreamMedia(id)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStream.GetLiveStreamMedia``: %v\n", err)
    }
    // response from `GetLiveStreamMedia`: GetLiveStreamMediaResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `LiveStream.GetLiveStreamMedia`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Live stream media ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**GetLiveStreamMediaResponse**](GetLiveStreamMediaResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLiveStreamMedias

> GetLiveStreamMedias(id string, getLiveStreamMediasRequest GetLiveStreamMediasRequest) (*GetLiveStreamMediasResponse, error)

> GetLiveStreamMediasWithContext(ctx context.Context, id string, getLiveStreamMediasRequest GetLiveStreamMediasRequest) (*GetLiveStreamMediasResponse, error)


Get live stream media



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
        
    id := "id_example" // string | Live stream key ID
    getLiveStreamMediasRequest := *aiozstreamsdk.NewGetLiveStreamMediasRequest() // GetLiveStreamMediasRequest | data

    
    res, err := client.LiveStream.GetLiveStreamMedias(id, getLiveStreamMediasRequest)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStream.GetLiveStreamMedias``: %v\n", err)
    }
    // response from `GetLiveStreamMedias`: GetLiveStreamMediasResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `LiveStream.GetLiveStreamMedias`")
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
**getLiveStreamMediasRequest** | [**GetLiveStreamMediasRequest**](GetLiveStreamMediasRequest.md) | data | 

### Return type

[**GetLiveStreamMediasResponse**](GetLiveStreamMediasResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLiveStreamMulticastByStreamKey

> GetLiveStreamMulticastByStreamKey(streamKey string) (*GetLiveStreamMulticastResponse, error)

> GetLiveStreamMulticastByStreamKeyWithContext(ctx context.Context, streamKey string) (*GetLiveStreamMulticastResponse, error)


Get live stream multicast by stream key



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
        
    streamKey := "streamKey_example" // string | Live stream key. UUID string format

    
    res, err := client.LiveStream.GetLiveStreamMulticastByStreamKey(streamKey)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStream.GetLiveStreamMulticastByStreamKey``: %v\n", err)
    }
    // response from `GetLiveStreamMulticastByStreamKey`: GetLiveStreamMulticastResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `LiveStream.GetLiveStreamMulticastByStreamKey`")
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


## GetLiveStreamPlayerInfo

> GetLiveStreamPlayerInfo(id string) (*GetLiveStreamMediaPublicResponse, error)

> GetLiveStreamPlayerInfoWithContext(ctx context.Context, id string) (*GetLiveStreamMediaPublicResponse, error)


Get live stream media public



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
        
    id := "id_example" // string | Live stream key ID

    
    res, err := client.LiveStream.GetLiveStreamPlayerInfo(id)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStream.GetLiveStreamPlayerInfo``: %v\n", err)
    }
    // response from `GetLiveStreamPlayerInfo`: GetLiveStreamMediaPublicResponse
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

[**GetLiveStreamMediaPublicResponse**](GetLiveStreamMediaPublicResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLiveStreamStatisticByStreamMediaId

> GetLiveStreamStatisticByStreamMediaId(streamMediaId string) (*GetLiveStreamStatisticResponse, error)

> GetLiveStreamStatisticByStreamMediaIdWithContext(ctx context.Context, streamMediaId string) (*GetLiveStreamStatisticResponse, error)


Get live stream statistic by stream media id



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
        
    streamMediaId := "streamMediaId_example" // string | Live stream media ID

    
    res, err := client.LiveStream.GetLiveStreamStatisticByStreamMediaId(streamMediaId)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStream.GetLiveStreamStatisticByStreamMediaId``: %v\n", err)
    }
    // response from `GetLiveStreamStatisticByStreamMediaId`: GetLiveStreamStatisticResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `LiveStream.GetLiveStreamStatisticByStreamMediaId`")
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


## GetLiveStreamUsage

> GetLiveStreamUsage(streamId string) (*GetStreamUsageResponse, error)

> GetLiveStreamUsageWithContext(ctx context.Context, streamId string) (*GetStreamUsageResponse, error)


Get usage details for a specific live stream



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
        
    streamId := "streamId_example" // string | Stream ID (LiveStreamMedia ID)

    
    res, err := client.LiveStream.GetLiveStreamUsage(streamId)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStream.GetLiveStreamUsage``: %v\n", err)
    }
    // response from `GetLiveStreamUsage`: GetStreamUsageResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `LiveStream.GetLiveStreamUsage`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**streamId** | **string** | Stream ID (LiveStreamMedia ID) | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**GetStreamUsageResponse**](GetStreamUsageResponse.md)

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
    aiozstreamsdk "github.com/AIOZNetwork/aioz-stream-go-client/v3"
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

> GetStreamings(id string, r LiveStreamApiGetStreamingsRequest) (*GetStreamingsResponse, error)


> GetStreamingsWithContext(ctx context.Context, id string, r LiveStreamApiGetStreamingsRequest) (*GetStreamingsResponse, error)



Get live stream media streamings



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
    req := aiozstreamsdk.LiveStreamApiGetStreamingsRequest{}
    
    req.Id("id_example") // string | Live stream key ID
    req.Search("search_example") // string | Search

    res, err := client.LiveStream.GetStreamings(id string, req)
    

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
**search** | **string** | Search | 

### Return type

[**GetStreamingsResponse**](GetStreamingsResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserStreamsUsageDetail

> GetUserStreamsUsageDetail(r LiveStreamApiGetUserStreamsUsageDetailRequest) (*GetUserStreamsUsageDetailResponse, error)


> GetUserStreamsUsageDetailWithContext(ctx context.Context, r LiveStreamApiGetUserStreamsUsageDetailRequest) (*GetUserStreamsUsageDetailResponse, error)



Get paginated list of streams with usage details for current user



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
    req := aiozstreamsdk.LiveStreamApiGetUserStreamsUsageDetailRequest{}
    
    req.From(int64(789)) // int64 | Start unix timestamp (seconds)
    req.To(int64(789)) // int64 | End unix timestamp (seconds)
    req.Offset(int64(789)) // int64 | Offset for pagination
    req.Limit(int64(789)) // int64 | Limit for pagination

    res, err := client.LiveStream.GetUserStreamsUsageDetail(req)
    

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStream.GetUserStreamsUsageDetail``: %v\n", err)
    }
    // response from `GetUserStreamsUsageDetail`: GetUserStreamsUsageDetailResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `LiveStream.GetUserStreamsUsageDetail`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**from** | **int64** | Start unix timestamp (seconds) | 
**to** | **int64** | End unix timestamp (seconds) | 
**offset** | **int64** | Offset for pagination | 
**limit** | **int64** | Limit for pagination | 

### Return type

[**GetUserStreamsUsageDetailResponse**](GetUserStreamsUsageDetailResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLiveStreamKey

> UpdateLiveStreamKey(id string, updateLiveStreamKeyRequest UpdateLiveStreamKeyRequest) (*UpdateLiveStreamKeyResponse, error)

> UpdateLiveStreamKeyWithContext(ctx context.Context, id string, updateLiveStreamKeyRequest UpdateLiveStreamKeyRequest) (*UpdateLiveStreamKeyResponse, error)


Update live stream key



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
        
    id := "id_example" // string | Live stream key ID
    updateLiveStreamKeyRequest := *aiozstreamsdk.NewUpdateLiveStreamKeyRequest() // UpdateLiveStreamKeyRequest | UpdateLiveStreamKeyRequest

    
    res, err := client.LiveStream.UpdateLiveStreamKey(id, updateLiveStreamKeyRequest)

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
**updateLiveStreamKeyRequest** | [**UpdateLiveStreamKeyRequest**](UpdateLiveStreamKeyRequest.md) | UpdateLiveStreamKeyRequest | 

### Return type

[**UpdateLiveStreamKeyResponse**](UpdateLiveStreamKeyResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLiveStreamMedia

> UpdateLiveStreamMedia(id string, updateLiveStreamMediaRequest UpdateLiveStreamMediaRequest) (*ResponseSuccess, error)

> UpdateLiveStreamMediaWithContext(ctx context.Context, id string, updateLiveStreamMediaRequest UpdateLiveStreamMediaRequest) (*ResponseSuccess, error)


Update live stream media



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
        
    id := "id_example" // string | Live stream key ID
    updateLiveStreamMediaRequest := *aiozstreamsdk.NewUpdateLiveStreamMediaRequest() // UpdateLiveStreamMediaRequest | data

    
    res, err := client.LiveStream.UpdateLiveStreamMedia(id, updateLiveStreamMediaRequest)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStream.UpdateLiveStreamMedia``: %v\n", err)
    }
    // response from `UpdateLiveStreamMedia`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `LiveStream.UpdateLiveStreamMedia`")
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
**updateLiveStreamMediaRequest** | [**UpdateLiveStreamMediaRequest**](UpdateLiveStreamMediaRequest.md) | data | 

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

