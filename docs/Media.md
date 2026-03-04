# \Media

All URIs are relative to https://api.aiozstream.network/api

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](Media.md#Create) | **Post** /media/create | Create media object
[**Update**](Media.md#Update) | **Patch** /media/{id} | update media info
[**Delete**](Media.md#Delete) | **Delete** /media/{id} | Delete media
[**UploadThumbnail**](Media.md#UploadThumbnail) | **Post** /media/{id}/thumbnail | Upload media thumbnail
[**DeleteThumbnail**](Media.md#DeleteThumbnail) | **Delete** /media/{id}/thumbnail | Delete media thumbnail
[**CreateCaption**](Media.md#CreateCaption) | **Post** /media/{id}/captions/{lan} | Create a new media caption
[**DeleteCaption**](Media.md#DeleteCaption) | **Delete** /media/{id}/captions/{lan} | Delete a media caption
[**GetCaptions**](Media.md#GetCaptions) | **Get** /media/{id}/captions | Get media captions
[**GetCost**](Media.md#GetCost) | **Get** /media/cost | get media transcoding cost
[**GetDetail**](Media.md#GetDetail) | **Get** /media/{id} | get media detail
[**GetMediaList**](Media.md#GetMediaList) | **Post** /media | Get user videos list
[**GetMediaPlayerInfo**](Media.md#GetMediaPlayerInfo) | **Get** /media/{id}/player.json | Get media object
[**SetDefaultCaption**](Media.md#SetDefaultCaption) | **Patch** /media/{id}/captions/{lan} | Set default caption
[**UploadMediaComplete**](Media.md#UploadMediaComplete) | **Get** /media/{id}/complete | Get upload media when complete
[**UploadPart**](Media.md#UploadPart) | **Post** /media/{id}/part | Upload part of media



## Create

> Create(request CreateMediaRequest) (*CreateMediaResponse, error)

> CreateWithContext(ctx context.Context, request CreateMediaRequest) (*CreateMediaResponse, error)


Create media object



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
        
    request := *aiozstreamsdk.NewCreateMediaRequest() // CreateMediaRequest | media's info

    
    res, err := client.Media.Create(request)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Media.Create``: %v\n", err)
    }
    // response from `Create`: CreateMediaResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Media.Create`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**request** | [**CreateMediaRequest**](CreateMediaRequest.md) | media&#39;s info | 

### Return type

[**CreateMediaResponse**](CreateMediaResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Update(id string, input UpdateMediaInfoRequest) (*ResponseSuccess, error)

> UpdateWithContext(ctx context.Context, id string, input UpdateMediaInfoRequest) (*ResponseSuccess, error)


update media info

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
        
    id := "id_example" // string | media's id
    input := *aiozstreamsdk.NewUpdateMediaInfoRequest() // UpdateMediaInfoRequest | input

    
    res, err := client.Media.Update(id, input)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Media.Update``: %v\n", err)
    }
    // response from `Update`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Media.Update`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | media&#39;s id | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**input** | [**UpdateMediaInfoRequest**](UpdateMediaInfoRequest.md) | input | 

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(id string) (*ResponseSuccess, error)

> DeleteWithContext(ctx context.Context, id string) (*ResponseSuccess, error)


Delete media



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
        
    id := "id_example" // string | Media ID

    
    res, err := client.Media.Delete(id)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Media.Delete``: %v\n", err)
    }
    // response from `Delete`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Media.Delete`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Media ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadThumbnail

> UploadThumbnailFile(id string, file *os.File) (*ResponseSuccess, error)
> UploadThumbnail(id string, fileName string, fileReader io.Reader)
> UploadThumbnailFileWithContext(ctx context.Context, id string, file *os.File) (*ResponseSuccess, error)
> UploadThumbnailWithContext(ctx context.Context, id string, fileName string, fileReader io.Reader)

Upload media thumbnail

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
        
    id := "id_example" // string | media's id
    file := os.NewFile(1234, "some_file") // *os.File | file media to be uploaded

    
    res, err := client.Media.UploadThumbnailFile(id, file)

    // you can also use a Reader instead of a File:
    // we recommend using Reader instead!
    // client.Media.UploadThumbnail(id, fileName, fileReader)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Media.UploadThumbnail``: %v\n", err)
    }
    // response from `UploadThumbnail`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Media.UploadThumbnail`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | media&#39;s id | 

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


Delete media thumbnail

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
        
    id := "id_example" // string | media's id

    
    res, err := client.Media.DeleteThumbnail(id)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Media.DeleteThumbnail``: %v\n", err)
    }
    // response from `DeleteThumbnail`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Media.DeleteThumbnail`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | media&#39;s id | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCaption

> CreateCaptionFile(id string, lan string, file *os.File) (*CreateMediaCaptionResponse, error)
> CreateCaption(id string, lan string, fileName string, fileReader io.Reader)
> CreateCaptionFileWithContext(ctx context.Context, id string, lan string, file *os.File) (*CreateMediaCaptionResponse, error)
> CreateCaptionWithContext(ctx context.Context, id string, lan string, fileName string, fileReader io.Reader)

Create a new media caption



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
        
    id := "id_example" // string | Media ID
    lan := "lan_example" // string | Language
    file := os.NewFile(1234, "some_file") // *os.File | VTT File

    
    res, err := client.Media.CreateCaptionFile(id, lan, file)

    // you can also use a Reader instead of a File:
    // we recommend using Reader instead!
    // client.Media.CreateCaption(id, lan, fileName, fileReader)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Media.CreateCaption``: %v\n", err)
    }
    // response from `CreateCaption`: CreateMediaCaptionResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Media.CreateCaption`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Media ID | 
**lan** | **string** | Language | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**file** | ***os.File** | VTT File | 

### Return type

[**CreateMediaCaptionResponse**](CreateMediaCaptionResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCaption

> DeleteCaption(id string, lan string) (*ResponseSuccess, error)

> DeleteCaptionWithContext(ctx context.Context, id string, lan string) (*ResponseSuccess, error)


Delete a media caption



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
        
    id := "id_example" // string | Media ID
    lan := "lan_example" // string | Language

    
    res, err := client.Media.DeleteCaption(id, lan)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Media.DeleteCaption``: %v\n", err)
    }
    // response from `DeleteCaption`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Media.DeleteCaption`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Media ID | 
**lan** | **string** | Language | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCaptions

> GetCaptions(id string, r MediaApiGetCaptionsRequest) (*GetMediaCaptionsResponse, error)


> GetCaptionsWithContext(ctx context.Context, id string, r MediaApiGetCaptionsRequest) (*GetMediaCaptionsResponse, error)



Get media captions



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
    req := aiozstreamsdk.MediaApiGetCaptionsRequest{}
    
    req.Id("id_example") // string | Media ID
    req.Offset(int32(56)) // int32 | offset, allowed values greater than or equal to 0. Default(0) (default to 0)
    req.Limit(int32(56)) // int32 | results per page. Allowed values 1-100, default is 25 (default to 25)

    res, err := client.Media.GetCaptions(id string, req)
    

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Media.GetCaptions``: %v\n", err)
    }
    // response from `GetCaptions`: GetMediaCaptionsResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Media.GetCaptions`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Media ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**offset** | **int32** | offset, allowed values greater than or equal to 0. Default(0) | [default to 0]
**limit** | **int32** | results per page. Allowed values 1-100, default is 25 | [default to 25]

### Return type

[**GetMediaCaptionsResponse**](GetMediaCaptionsResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCost

> GetCost(qualities string, type_ string, duration float32) (*GetTranscodeCostResponse, error)

> GetCostWithContext(ctx context.Context, qualities string, type_ string, duration float32) (*GetTranscodeCostResponse, error)


get media transcoding cost



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
        
    qualities := "qualities_example" // string | media's qualities
    type_ := "type__example" // string | media's type
    duration := float32(8.14) // float32 | media's duration

    
    res, err := client.Media.GetCost(qualities, type_, duration)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Media.GetCost``: %v\n", err)
    }
    // response from `GetCost`: GetTranscodeCostResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Media.GetCost`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**qualities** | **string** | media&#39;s qualities | 
**type_** | **string** | media&#39;s type | 
**duration** | **float32** | media&#39;s duration | 

### Return type

[**GetTranscodeCostResponse**](GetTranscodeCostResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDetail

> GetDetail(id string) (*GetMediaDetailResponse, error)

> GetDetailWithContext(ctx context.Context, id string) (*GetMediaDetailResponse, error)


get media detail



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
        
    id := "id_example" // string | mediav's id

    
    res, err := client.Media.GetDetail(id)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Media.GetDetail``: %v\n", err)
    }
    // response from `GetDetail`: GetMediaDetailResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Media.GetDetail`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | mediav&#39;s id | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**GetMediaDetailResponse**](GetMediaDetailResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMediaList

> GetMediaList(request GetMediaListRequest) (*GetMediaListResponse, error)

> GetMediaListWithContext(ctx context.Context, request GetMediaListRequest) (*GetMediaListResponse, error)


Get user videos list



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
        
    request := *aiozstreamsdk.NewGetMediaListRequest() // GetMediaListRequest | video's info

    
    res, err := client.Media.GetMediaList(request)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Media.GetMediaList``: %v\n", err)
    }
    // response from `GetMediaList`: GetMediaListResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Media.GetMediaList`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**request** | [**GetMediaListRequest**](GetMediaListRequest.md) | video&#39;s info | 

### Return type

[**GetMediaListResponse**](GetMediaListResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMediaPlayerInfo

> GetMediaPlayerInfo(id string, r MediaApiGetMediaPlayerInfoRequest) (*GetMediaPlayerInfoResponse, error)


> GetMediaPlayerInfoWithContext(ctx context.Context, id string, r MediaApiGetMediaPlayerInfoRequest) (*GetMediaPlayerInfoResponse, error)



Get media object



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
    req := aiozstreamsdk.MediaApiGetMediaPlayerInfoRequest{}
    
    req.Id("id_example") // string | media ID
    req.Token("token_example") // string | Token

    res, err := client.Media.GetMediaPlayerInfo(id string, req)
    

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Media.GetMediaPlayerInfo``: %v\n", err)
    }
    // response from `GetMediaPlayerInfo`: GetMediaPlayerInfoResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Media.GetMediaPlayerInfo`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | media ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**token** | **string** | Token | 

### Return type

[**GetMediaPlayerInfoResponse**](GetMediaPlayerInfoResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDefaultCaption

> SetDefaultCaption(id string, lan string) (*ResponseSuccess, error)

> SetDefaultCaptionWithContext(ctx context.Context, id string, lan string) (*ResponseSuccess, error)


Set default caption



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
        
    id := "id_example" // string | Media ID
    lan := "lan_example" // string | Language

    
    res, err := client.Media.SetDefaultCaption(id, lan)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Media.SetDefaultCaption``: %v\n", err)
    }
    // response from `SetDefaultCaption`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Media.SetDefaultCaption`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Media ID | 
**lan** | **string** | Language | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadMediaComplete

> UploadMediaComplete(id string) (*ResponseSuccess, error)

> UploadMediaCompleteWithContext(ctx context.Context, id string) (*ResponseSuccess, error)


Get upload media when complete



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
        
    id := "id_example" // string | media's id

    
    res, err := client.Media.UploadMediaComplete(id)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Media.UploadMediaComplete``: %v\n", err)
    }
    // response from `UploadMediaComplete`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Media.UploadMediaComplete`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | media&#39;s id | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadPart

> UploadPartFile(id string, file *os.File) (*ResponseSuccess, error)
> UploadPart(id string, fileName string, fileReader io.Reader, fileSize int64)
> UploadPartFileWithContext(ctx context.Context, id string, file *os.File) (*ResponseSuccess, error)
> UploadPartWithContext(ctx context.Context, id string, fileName string, fileReader io.Reader, fileSize int64)

Upload part of media



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
        
    id := "id_example" // string | video's id
    file := os.NewFile(1234, "some_file") // *os.File | File media to be uploaded
    hash := "hash_example" // string | Md5 hash of part
    index := "index_example" // string | Index of the part

    
    res, err := client.Media.UploadPartFile(id, file)

    // you can also use a Reader instead of a File:
    // we recommend using Reader instead!
    // client.Media.UploadPart(id, fileName, fileReader, fileSize)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Media.UploadPart``: %v\n", err)
    }
    // response from `UploadPart`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Media.UploadPart`")
    fmt.Println(string(newJsonString))
}
```


### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | video&#39;s id | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**file** | ***os.File** | File media to be uploaded | 
**hash** | **string** | Md5 hash of part | 
**index** | **string** | Index of the part | 

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

