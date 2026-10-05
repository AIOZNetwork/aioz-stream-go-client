# \Webhook

All URIs are relative to https://api.aiozstream.network/api

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](Webhook.md#Create) | **Post** /webhooks | Create a webhook
[**Get**](Webhook.md#Get) | **Get** /webhooks/{id} | Get a webhook
[**Update**](Webhook.md#Update) | **Patch** /webhooks/{id} | Update a webhook
[**Delete**](Webhook.md#Delete) | **Delete** /webhooks/{id} | Delete a webhook
[**List**](Webhook.md#List) | **Get** /webhooks | List webhooks
[**Check**](Webhook.md#Check) | **Post** /webhooks/check/{id} | Send a test event



## Create

> Create(writeWebhookRequest WriteWebhookRequest) (*WebhookResponse, error)

> CreateWithContext(ctx context.Context, writeWebhookRequest WriteWebhookRequest) (*WebhookResponse, error)


Create a webhook



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
        
    writeWebhookRequest := *aiozstreamsdk.NewWriteWebhookRequest() // WriteWebhookRequest | Webhook

    
    res, err := client.Webhook.Create(writeWebhookRequest)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Webhook.Create``: %v\n", err)
    }
    // response from `Create`: WebhookResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Webhook.Create`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**writeWebhookRequest** | [**WriteWebhookRequest**](WriteWebhookRequest.md) | Webhook | 

### Return type

[**WebhookResponse**](WebhookResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> Get(id string) (*WebhookResponse, error)

> GetWithContext(ctx context.Context, id string) (*WebhookResponse, error)


Get a webhook



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
        
    id := "id_example" // string | Webhook ID

    
    res, err := client.Webhook.Get(id)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Webhook.Get``: %v\n", err)
    }
    // response from `Get`: WebhookResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Webhook.Get`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Webhook ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**WebhookResponse**](WebhookResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Update(id string, writeWebhookRequest WriteWebhookRequest) (*ResponseSuccess, error)

> UpdateWithContext(ctx context.Context, id string, writeWebhookRequest WriteWebhookRequest) (*ResponseSuccess, error)


Update a webhook



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
        
    id := "id_example" // string | Webhook ID
    writeWebhookRequest := *aiozstreamsdk.NewWriteWebhookRequest() // WriteWebhookRequest | Fields to change

    
    res, err := client.Webhook.Update(id, writeWebhookRequest)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Webhook.Update``: %v\n", err)
    }
    // response from `Update`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Webhook.Update`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Webhook ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**writeWebhookRequest** | [**WriteWebhookRequest**](WriteWebhookRequest.md) | Fields to change | 

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(id string) (*ResponseSuccess, error)

> DeleteWithContext(ctx context.Context, id string) (*ResponseSuccess, error)


Delete a webhook



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
        
    id := "id_example" // string | Webhook ID

    
    res, err := client.Webhook.Delete(id)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Webhook.Delete``: %v\n", err)
    }
    // response from `Delete`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Webhook.Delete`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Webhook ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> List(r WebhookApiListRequest) (*ListWebhooksResponse, error)


> ListWithContext(ctx context.Context, r WebhookApiListRequest) (*ListWebhooksResponse, error)



List webhooks



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
    req := aiozstreamsdk.WebhookApiListRequest{}
    
    req.EncodingFailed(true) // bool | 
    req.EncodingFinished(true) // bool | 
    req.EncodingStarted(true) // bool | 
    req.FileReceived(true) // bool | 
    req.Limit(int64(789)) // int64 |  (default to 25)
    req.Offset(int64(789)) // int64 | 
    req.OrderBy("orderBy_example") // string | 
    req.PartialFinished(true) // bool | 
    req.Search("search_example") // string | 
    req.SortBy("sortBy_example") // string | 

    res, err := client.Webhook.List(req)
    

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Webhook.List``: %v\n", err)
    }
    // response from `List`: ListWebhooksResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Webhook.List`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**encodingFailed** | **bool** |  | 
**encodingFinished** | **bool** |  | 
**encodingStarted** | **bool** |  | 
**fileReceived** | **bool** |  | 
**limit** | **int64** |  | [default to 25]
**offset** | **int64** |  | 
**orderBy** | **string** |  | 
**partialFinished** | **bool** |  | 
**search** | **string** |  | 
**sortBy** | **string** |  | 

### Return type

[**ListWebhooksResponse**](ListWebhooksResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Check

> Check(id string) (*ResponseSuccess, error)

> CheckWithContext(ctx context.Context, id string) (*ResponseSuccess, error)


Send a test event



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
        
    id := "id_example" // string | Webhook ID

    
    res, err := client.Webhook.Check(id)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Webhook.Check``: %v\n", err)
    }
    // response from `Check`: ResponseSuccess
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `Webhook.Check`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Webhook ID | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ResponseSuccess**](ResponseSuccess.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

