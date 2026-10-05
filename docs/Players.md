# \Players

All URIs are relative to https://api.aiozstream.network/api

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](Players.md#Create) | **Post** /players | Create a player theme
[**Get**](Players.md#Get) | **Get** /players/{id} | Get a player theme
[**Update**](Players.md#Update) | **Patch** /players/{id} | Update a player theme
[**Delete**](Players.md#Delete) | **Delete** /players/{id} | Delete a player theme
[**List**](Players.md#List) | **Get** /players | List player themes
[**UploadLogo**](Players.md#UploadLogo) | **Post** /players/{id}/logo | Upload a player theme logo
[**DeleteLogo**](Players.md#DeleteLogo) | **Delete** /players/{id}/logo | Delete a player theme logo
[**Attach**](Players.md#Attach) | **Post** /players/add-player | Add a player theme to a media
[**Detach**](Players.md#Detach) | **Post** /players/remove-player | Remove a player theme from a media



## Create

> Create(playerThemeInput PlayerThemeInput) (*ThemeResponse, error)

> CreateWithContext(ctx context.Context, playerThemeInput PlayerThemeInput) (*ThemeResponse, error)


Create a player theme



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
        
    playerThemeInput := *aiozstreamsdk.NewPlayerThemeInput() // PlayerThemeInput | Player theme

    
    res, err := client.Players.Create(playerThemeInput)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Players.Create``: %v\n", err)
    }
    // response from `Create`: ThemeResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Players.Create`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**playerThemeInput** | [**PlayerThemeInput**](PlayerThemeInput.md) | Player theme | 

### Return type

[**ThemeResponse**](ThemeResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> Get(id string) (*ThemeResponse, error)

> GetWithContext(ctx context.Context, id string) (*ThemeResponse, error)


Get a player theme



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
        
    id := "id_example" // string | Player theme ID

    
    res, err := client.Players.Get(id)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Players.Get``: %v\n", err)
    }
    // response from `Get`: ThemeResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Players.Get`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Player theme ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ThemeResponse**](ThemeResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Update(id string, playerThemeInput PlayerThemeInput) (*ThemeResponse, error)

> UpdateWithContext(ctx context.Context, id string, playerThemeInput PlayerThemeInput) (*ThemeResponse, error)


Update a player theme



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
        
    id := "id_example" // string | Player theme ID
    playerThemeInput := *aiozstreamsdk.NewPlayerThemeInput() // PlayerThemeInput | Fields to change

    
    res, err := client.Players.Update(id, playerThemeInput)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Players.Update``: %v\n", err)
    }
    // response from `Update`: ThemeResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Players.Update`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Player theme ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**playerThemeInput** | [**PlayerThemeInput**](PlayerThemeInput.md) | Fields to change | 

### Return type

[**ThemeResponse**](ThemeResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(id string) (*ResponseSuccess, error)

> DeleteWithContext(ctx context.Context, id string) (*ResponseSuccess, error)


Delete a player theme



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
        
    id := "id_example" // string | Player theme ID

    
    res, err := client.Players.Delete(id)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Players.Delete``: %v\n", err)
    }
    // response from `Delete`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Players.Delete`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Player theme ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> List(r PlayersApiListRequest) (*ListThemesResponse, error)


> ListWithContext(ctx context.Context, r PlayersApiListRequest) (*ListThemesResponse, error)



List player themes



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
    req := aiozstreamsdk.PlayersApiListRequest{}
    
    req.Limit(int32(56)) // int32 |  (default to 25)
    req.Offset(int32(56)) // int32 | 
    req.OrderBy("orderBy_example") // string | 
    req.Search("search_example") // string | 
    req.SortBy("sortBy_example") // string | 

    res, err := client.Players.List(req)
    

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Players.List``: %v\n", err)
    }
    // response from `List`: ListThemesResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Players.List`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**limit** | **int32** |  | [default to 25]
**offset** | **int32** |  | 
**orderBy** | **string** |  | 
**search** | **string** |  | 
**sortBy** | **string** |  | 

### Return type

[**ListThemesResponse**](ListThemesResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadLogo

> UploadLogoFile(id string, file *os.File) (*ThemeResponse, error)
> UploadLogo(id string, fileName string, fileReader io.Reader)
> UploadLogoFileWithContext(ctx context.Context, id string, file *os.File) (*ThemeResponse, error)
> UploadLogoWithContext(ctx context.Context, id string, fileName string, fileReader io.Reader)

Upload a player theme logo



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
        
    id := "id_example" // string | Player theme ID
    file := os.NewFile(1234, "some_file") // *os.File | Logo image
    link := "link_example" // string | Where clicking the logo takes the viewer

    
    res, err := client.Players.UploadLogoFile(id, file)

    // you can also use a Reader instead of a File:
    // we recommend using Reader instead!
    // client.Players.UploadLogo(id, fileName, fileReader)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Players.UploadLogo``: %v\n", err)
    }
    // response from `UploadLogo`: ThemeResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Players.UploadLogo`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Player theme ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**file** | ***os.File** | Logo image | 
**link** | **string** | Where clicking the logo takes the viewer | 

### Return type

[**ThemeResponse**](ThemeResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLogo

> DeleteLogo(id string) (*ResponseSuccess, error)

> DeleteLogoWithContext(ctx context.Context, id string) (*ResponseSuccess, error)


Delete a player theme logo



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
        
    id := "id_example" // string | Player theme ID

    
    res, err := client.Players.DeleteLogo(id)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Players.DeleteLogo``: %v\n", err)
    }
    // response from `DeleteLogo`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Players.DeleteLogo`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Player theme ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Attach

> Attach(attachThemeRequest AttachThemeRequest) (*ResponseSuccess, error)

> AttachWithContext(ctx context.Context, attachThemeRequest AttachThemeRequest) (*ResponseSuccess, error)


Add a player theme to a media



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
        
    attachThemeRequest := *aiozstreamsdk.NewAttachThemeRequest() // AttachThemeRequest | Media and theme

    
    res, err := client.Players.Attach(attachThemeRequest)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Players.Attach``: %v\n", err)
    }
    // response from `Attach`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Players.Attach`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**attachThemeRequest** | [**AttachThemeRequest**](AttachThemeRequest.md) | Media and theme | 

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Detach

> Detach(attachThemeRequest AttachThemeRequest) (*ResponseSuccess, error)

> DetachWithContext(ctx context.Context, attachThemeRequest AttachThemeRequest) (*ResponseSuccess, error)


Remove a player theme from a media



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
        
    attachThemeRequest := *aiozstreamsdk.NewAttachThemeRequest() // AttachThemeRequest | Media and theme

    
    res, err := client.Players.Detach(attachThemeRequest)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Players.Detach``: %v\n", err)
    }
    // response from `Detach`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Players.Detach`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**attachThemeRequest** | [**AttachThemeRequest**](AttachThemeRequest.md) | Media and theme | 

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

