# \User

All URIs are relative to https://api.aiozstream.network/api

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMe**](User.md#GetMe) | **Get** /user/me | Get the current account



## GetMe

> GetMe() (*MeResponse, error)

> GetMeWithContext(ctx context.Context, ) (*MeResponse, error)


Get the current account



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
        

    
    res, err := client.User.GetMe()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `User.GetMe``: %v\n", err)
    }
    // response from `GetMe`: MeResponse
    newJsonString, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println("Response from `User.GetMe`")
    fmt.Println(string(newJsonString))
}
```
### Path Parameters

This endpoint does not need any parameter.

### Other Parameters



### Return type

[**MeResponse**](MeResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

