# \ApiKey

All URIs are relative to https://api.aiozstream.network/api

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](ApiKey.md#Create) | **Post** /api_keys | Create API key
[**Update**](ApiKey.md#Update) | **Patch** /api_keys/{id} | Rename an API key
[**Delete**](ApiKey.md#Delete) | **Delete** /api_keys/{id} | Delete an API key
[**List**](ApiKey.md#List) | **Get** /api_keys | List API keys



## Create

> Create(createApiKeyRequest CreateApiKeyRequest) (*CreateApiKeyResponse, error)

> CreateWithContext(ctx context.Context, createApiKeyRequest CreateApiKeyRequest) (*CreateApiKeyResponse, error)


Create API key



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
        
    createApiKeyRequest := *aiozstreamsdk.NewCreateApiKeyRequest() // CreateApiKeyRequest | api key's data

    
    res, err := client.ApiKey.Create(createApiKeyRequest)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiKey.Create``: %v\n", err)
    }
    // response from `Create`: CreateApiKeyResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `ApiKey.Create`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**createApiKeyRequest** | [**CreateApiKeyRequest**](CreateApiKeyRequest.md) | api key&#39;s data | 

### Return type

[**CreateApiKeyResponse**](CreateApiKeyResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Update(id string, renameApiKeyRequest RenameApiKeyRequest) (*ResponseSuccess, error)

> UpdateWithContext(ctx context.Context, id string, renameApiKeyRequest RenameApiKeyRequest) (*ResponseSuccess, error)


Rename an API key



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
        
    id := "id_example" // string | API key ID
    renameApiKeyRequest := *aiozstreamsdk.NewRenameApiKeyRequest() // RenameApiKeyRequest | new name

    
    res, err := client.ApiKey.Update(id, renameApiKeyRequest)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiKey.Update``: %v\n", err)
    }
    // response from `Update`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `ApiKey.Update`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | API key ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**renameApiKeyRequest** | [**RenameApiKeyRequest**](RenameApiKeyRequest.md) | new name | 

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(id string) (*ResponseSuccess, error)

> DeleteWithContext(ctx context.Context, id string) (*ResponseSuccess, error)


Delete an API key



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
        
    id := "id_example" // string | API key ID

    
    res, err := client.ApiKey.Delete(id)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiKey.Delete``: %v\n", err)
    }
    // response from `Delete`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `ApiKey.Delete`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | API key ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> List(r ApiKeyApiListRequest) (*ListApiKeysResponse, error)


> ListWithContext(ctx context.Context, r ApiKeyApiListRequest) (*ListApiKeysResponse, error)



List API keys



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
    req := aiozstreamsdk.ApiKeyApiListRequest{}
    
    req.Limit(int32(56)) // int32 |  (default to 25)
    req.Offset(int32(56)) // int32 | 
    req.OrderBy("orderBy_example") // string | 
    req.Search("search_example") // string | 
    req.SortBy("sortBy_example") // string | 
    req.Type_("type__example") // string | 

    res, err := client.ApiKey.List(req)
    

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiKey.List``: %v\n", err)
    }
    // response from `List`: ListApiKeysResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `ApiKey.List`")
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
**type_** | **string** |  | 

### Return type

[**ListApiKeysResponse**](ListApiKeysResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

