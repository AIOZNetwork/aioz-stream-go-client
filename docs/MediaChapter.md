# \MediaChapter

All URIs are relative to https://api.aiozstream.network/api

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](MediaChapter.md#Create) | **Post** /media/{id}/chapters/{lan} | Create a media chapter
[**Get**](MediaChapter.md#Get) | **Get** /media/{id}/chapters | Get media chapters
[**Delete**](MediaChapter.md#Delete) | **Delete** /media/{id}/chapters/{lan} | Delete a video chapter



## Create

> CreateFile(id string, lan string, file *os.File) (*CreateMediaChapterResponse, error)
> Create(id string, lan string, fileName string, fileReader io.Reader)
> CreateFileWithContext(ctx context.Context, id string, lan string, file *os.File) (*CreateMediaChapterResponse, error)
> CreateWithContext(ctx context.Context, id string, lan string, fileName string, fileReader io.Reader)

Create a media chapter



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

    
    res, err := client.MediaChapter.CreateFile(id, lan, file)

    // you can also use a Reader instead of a File:
    // we recommend using Reader instead!
    // client.MediaChapter.Create(id, lan, fileName, fileReader)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaChapter.Create``: %v\n", err)
    }
    // response from `Create`: CreateMediaChapterResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `MediaChapter.Create`")
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

[**CreateMediaChapterResponse**](CreateMediaChapterResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> Get(id string, r MediaChapterApiGetRequest) (*GetMediaChaptersResponse, error)


> GetWithContext(ctx context.Context, id string, r MediaChapterApiGetRequest) (*GetMediaChaptersResponse, error)



Get media chapters



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
    req := aiozstreamsdk.MediaChapterApiGetRequest{}
    
    req.Id("id_example") // string | Media ID
    req.Offset(int32(56)) // int32 | offset, allowed values greater than or equal to 0. Default(0) (default to 0)
    req.Limit(int32(56)) // int32 | results per page. Allowed values 1-100, default is 25 (default to 25)

    res, err := client.MediaChapter.Get(id string, req)
    

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaChapter.Get``: %v\n", err)
    }
    // response from `Get`: GetMediaChaptersResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `MediaChapter.Get`")
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

[**GetMediaChaptersResponse**](GetMediaChaptersResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(id string, lan string) (*ResponseSuccess, error)

> DeleteWithContext(ctx context.Context, id string, lan string) (*ResponseSuccess, error)


Delete a video chapter



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
        
    id := "id_example" // string | Video ID
    lan := "lan_example" // string | Language

    
    res, err := client.MediaChapter.Delete(id, lan)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaChapter.Delete``: %v\n", err)
    }
    // response from `Delete`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `MediaChapter.Delete`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Video ID | 
**lan** | **string** | Language | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

