# \Playlist

All URIs are relative to https://api.aiozstream.network/api

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](Playlist.md#Create) | **Post** /playlists/create | Create a playlist
[**Get**](Playlist.md#Get) | **Get** /playlists/{id} | Get a playlist
[**Update**](Playlist.md#Update) | **Patch** /playlists/{id} | Update a playlist
[**Delete**](Playlist.md#Delete) | **Delete** /playlists/{id} | Delete a playlist
[**List**](Playlist.md#List) | **Post** /playlists | List playlists
[**DeleteThumbnail**](Playlist.md#DeleteThumbnail) | **Delete** /playlists/{id}/thumbnail | Delete a playlist thumbnail
[**AddMedia**](Playlist.md#AddMedia) | **Post** /playlists/{id}/items | Add media to playlists
[**GetPublic**](Playlist.md#GetPublic) | **Get** /playlists/{id}/player.json | Get a playlist for the player
[**MoveItem**](Playlist.md#MoveItem) | **Put** /playlists/{id}/items | Reorder a playlist
[**RemoveMedia**](Playlist.md#RemoveMedia) | **Delete** /playlists/{id}/items/{item_id} | Remove an item from playlists



## Create

> Create(createPlaylistRequest CreatePlaylistRequest) (*PlaylistResponse, error)

> CreateWithContext(ctx context.Context, createPlaylistRequest CreatePlaylistRequest) (*PlaylistResponse, error)


Create a playlist



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
        
    createPlaylistRequest := *aiozstreamsdk.NewCreatePlaylistRequest() // CreatePlaylistRequest | Playlist

    
    res, err := client.Playlist.Create(createPlaylistRequest)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Playlist.Create``: %v\n", err)
    }
    // response from `Create`: PlaylistResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Playlist.Create`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**createPlaylistRequest** | [**CreatePlaylistRequest**](CreatePlaylistRequest.md) | Playlist | 

### Return type

[**PlaylistResponse**](PlaylistResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> Get(id string, r PlaylistApiGetRequest) (*PlaylistResponse, error)


> GetWithContext(ctx context.Context, id string, r PlaylistApiGetRequest) (*PlaylistResponse, error)



Get a playlist



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
    req := aiozstreamsdk.PlaylistApiGetRequest{}
    
    req.Id("id_example") // string | Playlist ID
    req.OrderBy("orderBy_example") // string | 
    req.Search("search_example") // string | 
    req.SortBy("sortBy_example") // string | 

    res, err := client.Playlist.Get(id string, req)
    

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Playlist.Get``: %v\n", err)
    }
    // response from `Get`: PlaylistResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Playlist.Get`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Playlist ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**orderBy** | **string** |  | 
**search** | **string** |  | 
**sortBy** | **string** |  | 

### Return type

[**PlaylistResponse**](PlaylistResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> UpdateFile(id string) (*ResponseSuccess, error)
> Update(id string, fileName string, fileReader io.Reader)
> UpdateFileWithContext(ctx context.Context, id string) (*ResponseSuccess, error)
> UpdateWithContext(ctx context.Context, id string, fileName string, fileReader io.Reader)

Update a playlist



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
        
    id := "id_example" // string | Playlist ID
    file := os.NewFile(1234, "some_file") // *os.File | New thumbnail
    name := "name_example" // string | New name
    tags := []string{"Inner_example"} // []string | New tags, one field per tag

    
    res, err := client.Playlist.UpdateFile(id)

    // you can also use a Reader instead of a File:
    // we recommend using Reader instead!
    // client.Playlist.Update(id, fileName, fileReader)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Playlist.Update``: %v\n", err)
    }
    // response from `Update`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Playlist.Update`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Playlist ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**file** | ***os.File** | New thumbnail | 
**name** | **string** | New name | 
**tags** | **[]string** | New tags, one field per tag | 

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(id string) (*ResponseSuccess, error)

> DeleteWithContext(ctx context.Context, id string) (*ResponseSuccess, error)


Delete a playlist



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
        
    id := "id_example" // string | Playlist ID

    
    res, err := client.Playlist.Delete(id)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Playlist.Delete``: %v\n", err)
    }
    // response from `Delete`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Playlist.Delete`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Playlist ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> List(listPlaylistsRequest ListPlaylistsRequest) (*ListPlaylistsResponse, error)

> ListWithContext(ctx context.Context, listPlaylistsRequest ListPlaylistsRequest) (*ListPlaylistsResponse, error)


List playlists



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
        
    listPlaylistsRequest := *aiozstreamsdk.NewListPlaylistsRequest() // ListPlaylistsRequest | Filter and paging

    
    res, err := client.Playlist.List(listPlaylistsRequest)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Playlist.List``: %v\n", err)
    }
    // response from `List`: ListPlaylistsResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Playlist.List`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**listPlaylistsRequest** | [**ListPlaylistsRequest**](ListPlaylistsRequest.md) | Filter and paging | 

### Return type

[**ListPlaylistsResponse**](ListPlaylistsResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteThumbnail

> DeleteThumbnail(id string) (*ResponseSuccess, error)

> DeleteThumbnailWithContext(ctx context.Context, id string) (*ResponseSuccess, error)


Delete a playlist thumbnail



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
        
    id := "id_example" // string | Playlist ID

    
    res, err := client.Playlist.DeleteThumbnail(id)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Playlist.DeleteThumbnail``: %v\n", err)
    }
    // response from `DeleteThumbnail`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Playlist.DeleteThumbnail`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Playlist ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddMedia

> AddMedia(id string, addMediaRequest AddMediaRequest) (*ResponseSuccess, error)

> AddMediaWithContext(ctx context.Context, id string, addMediaRequest AddMediaRequest) (*ResponseSuccess, error)


Add media to playlists



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
        
    id := "id_example" // string | Playlist ID
    addMediaRequest := *aiozstreamsdk.NewAddMediaRequest() // AddMediaRequest | Media and playlists

    
    res, err := client.Playlist.AddMedia(id, addMediaRequest)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Playlist.AddMedia``: %v\n", err)
    }
    // response from `AddMedia`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Playlist.AddMedia`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Playlist ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**addMediaRequest** | [**AddMediaRequest**](AddMediaRequest.md) | Media and playlists | 

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublic

> GetPublic(id string) (*ResponseSuccess, error)

> GetPublicWithContext(ctx context.Context, id string) (*ResponseSuccess, error)


Get a playlist for the player



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
        
    id := "id_example" // string | Playlist ID

    
    res, err := client.Playlist.GetPublic(id)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Playlist.GetPublic``: %v\n", err)
    }
    // response from `GetPublic`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Playlist.GetPublic`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Playlist ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveItem

> MoveItem(id string, moveItemRequest MoveItemRequest) (*ResponseSuccess, error)

> MoveItemWithContext(ctx context.Context, id string, moveItemRequest MoveItemRequest) (*ResponseSuccess, error)


Reorder a playlist



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
        
    id := "id_example" // string | Playlist ID
    moveItemRequest := *aiozstreamsdk.NewMoveItemRequest() // MoveItemRequest | Where to move it

    
    res, err := client.Playlist.MoveItem(id, moveItemRequest)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Playlist.MoveItem``: %v\n", err)
    }
    // response from `MoveItem`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Playlist.MoveItem`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Playlist ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**moveItemRequest** | [**MoveItemRequest**](MoveItemRequest.md) | Where to move it | 

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveMedia

> RemoveMedia(id string, itemId string, removeMediaRequest RemoveMediaRequest) (*ResponseSuccess, error)

> RemoveMediaWithContext(ctx context.Context, id string, itemId string, removeMediaRequest RemoveMediaRequest) (*ResponseSuccess, error)


Remove an item from playlists



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
        
    id := "id_example" // string | Playlist ID
    itemId := "itemId_example" // string | Playlist item ID
    removeMediaRequest := *aiozstreamsdk.NewRemoveMediaRequest() // RemoveMediaRequest | Other playlists

    
    res, err := client.Playlist.RemoveMedia(id, itemId, removeMediaRequest)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Playlist.RemoveMedia``: %v\n", err)
    }
    // response from `RemoveMedia`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Playlist.RemoveMedia`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Playlist ID | 
**itemId** | **string** | Playlist item ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**removeMediaRequest** | [**RemoveMediaRequest**](RemoveMediaRequest.md) | Other playlists | 

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

