<!--<documentation_excluded>-->
<h1 align="center">AIOZ Stream Go client</h1>

AIOZ Stream is the video infrastructure for product builders. Lightning fast video APIs for integrating, scaling, and managing on-demand & low latency live streaming features in your app.

## Project description

AIOZ Stream's Go client streamlines the coding process. Chunking files is handled for you, as is pagination and refreshing your tokens.

## Getting started

### Installation
```bash
go get github.com/AIOZNetwork/aioz-stream-go-client/v3
```


### Code sample

For a more advanced usage you can checkout the rest of the documentation in the [docs directory](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs)

```golang

import (
	"context"
	"fmt"
	"os"
 
	aiozstreamsdk "github.com/AIOZNetwork/aioz-stream-go-client/v3"
)
 
func main() {
    // Connect to production environment
    publicKey := "YOUR_PUBLIC_KEY" // Replace with your public key
    secretKey := "YOUR_SECRET_KEY" // Replace with your actual API secret key
	apiCreds := aiozstreamsdk.AuthCredentials{
		PublicKey: publicKey,
		SecretKey: secretKey,
	}
    client := aiozstreamsdk.ClientBuilder(apiCreds).Build()
 
    // Create a video object
	title := "Sample Video Title"
	videoData := aiozstreamsdk.CreateVideoRequest{
		Title: &title,
	}
	createResult, err := client.Video.Create(videoData)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating video: %v\n", err)
		return
	}
 
    videoId := createResult.Data.Id // Get the video ID from the response
 
    // Open the video file
    videoFile, err := os.Open("./path/to/video.mp4")
    if err != nil {
        fmt.Println("Error opening video file:", err)
        return
    }
    defer videoFile.Close() // Close the file after use
 
    fileInfo, err := videoFile.Stat()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error getting file info: %v\n", err)
        return
    }
 
    fileSize := fileInfo.Size()
    fileName := fileInfo.Name()
 
    // Option 1: Use client upload with videoId
	err = client.UploadVideo(context.Background(), *videoId, fileName, videoFile, fileSize)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error uploading video with client: %v\n", err)
		return
	}
 
    // Option 2: Upload parts yourself
    // This example is commented out as you already used option 1
	//_, err = client.Video.UploadPart(*videoId, nil, nil, "./path/to/video.mp4", videoFile, fileInfo.Size())
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Error uploading video part: %v\n", err)
	//	return
	//}
	//
	//success, err := client.Video.UploadVideoComplete(*videoId)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Error completing video upload: %v\n", err)
	//	return
	//}
	//
	//jsonString, err := json.MarshalIndent(success, "", "  ")
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Error marshalling response: %v\n", err)
	//	return
	//}
	//fmt.Println(string(jsonString))
    fmt.Println("Video uploaded successfully!")
}
```

## Documentation

### API endpoints

All urls are relative to https://api.aiozstream.network/api


#### Analytics


##### Retrieve an instance of the Analytics API:
```go
secretKey := "YOUR_SECRET_KEY" // Replace with your actual secret key
publicKey := "YOUR_PUBLIC_KEY" // Replace with your public key
apiCreds := aiozstreamsdk.AuthCredentials{
	PublicKey: publicKey,
	SecretKey: secretKey,
}
client := aiozstreamsdk.ClientBuilder(apiCreds).Build()
analyticsApi := client.Analytics
```

##### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAggregatedMetrics**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Analytics.md#GetAggregatedMetrics) | **Post** `/analytics/metrics/data/{metric}/{aggregation}` | Aggregate one metric
[**GetBreakdownMetrics**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Analytics.md#GetBreakdownMetrics) | **Post** `/analytics/metrics/bucket/{metric}/{breakdown}` | Bucket one metric by dimension
[**GetDataUsage**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Analytics.md#GetDataUsage) | **Get** `/analytics/data` | Delivery volume over time
[**GetOvertimeMetrics**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Analytics.md#GetOvertimeMetrics) | **Post** `/analytics/metrics/timeseries/{metric}/{interval}` | Bucket one metric by time


#### ApiKey


##### Retrieve an instance of the ApiKey API:
```go
secretKey := "YOUR_SECRET_KEY" // Replace with your actual secret key
publicKey := "YOUR_PUBLIC_KEY" // Replace with your public key
apiCreds := aiozstreamsdk.AuthCredentials{
	PublicKey: publicKey,
	SecretKey: secretKey,
}
client := aiozstreamsdk.ClientBuilder(apiCreds).Build()
apiKeyApi := client.ApiKey
```

##### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/ApiKey.md#Create) | **Post** `/api_keys` | Create API key
[**Update**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/ApiKey.md#Update) | **Patch** `/api_keys/{id}` | Rename an API key
[**Delete**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/ApiKey.md#Delete) | **Delete** `/api_keys/{id}` | Delete an API key
[**List**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/ApiKey.md#List) | **Get** `/api_keys` | List API keys


#### LiveStream


##### Retrieve an instance of the LiveStream API:
```go
secretKey := "YOUR_SECRET_KEY" // Replace with your actual secret key
publicKey := "YOUR_PUBLIC_KEY" // Replace with your public key
apiCreds := aiozstreamsdk.AuthCredentials{
	PublicKey: publicKey,
	SecretKey: secretKey,
}
client := aiozstreamsdk.ClientBuilder(apiCreds).Build()
liveStreamApi := client.LiveStream
```

##### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
[**UploadThumbnail**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStream.md#UploadThumbnail) | **Post** `/live_streams/{id}/thumbnail` | Upload live stream media thumbnail
[**DeleteThumbnail**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStream.md#DeleteThumbnail) | **Delete** `/live_streams/{id}/thumbnail` | Delete live stream media thumbnail
[**AddLiveStreamMulticasts**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStream.md#AddLiveStreamMulticasts) | **Post** `/live_streams/multicast/{stream_key}` | Add live stream multicast
[**CreateLiveStreamKey**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStream.md#CreateLiveStreamKey) | **Post** `/live_streams` | Create live stream key
[**CreateStreaming**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStream.md#CreateStreaming) | **Post** `/live_streams/{id}/streamings` | Create a new live stream media
[**DeleteLiveStreamKey**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStream.md#DeleteLiveStreamKey) | **Delete** `/live_streams/{id}` | Delete live stream key
[**DeleteLiveStreamMulticast**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStream.md#DeleteLiveStreamMulticast) | **Delete** `/live_streams/multicast/{stream_key}` | Delete live stream multicast
[**GetLiveStreamKey**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStream.md#GetLiveStreamKey) | **Get** `/live_streams/{id}` | Get live stream key
[**GetLiveStreamKeys**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStream.md#GetLiveStreamKeys) | **Get** `/live_streams` | Get live stream key list
[**GetLiveStreamMedia**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStream.md#GetLiveStreamMedia) | **Get** `/live_streams/{id}/media` | Get live stream media
[**GetLiveStreamMedias**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStream.md#GetLiveStreamMedias) | **Post** `/live_streams/{id}/media` | Get live stream media
[**GetLiveStreamMulticastByStreamKey**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStream.md#GetLiveStreamMulticastByStreamKey) | **Get** `/live_streams/multicast/{stream_key}` | Get live stream multicast by stream key
[**GetLiveStreamPlayerInfo**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStream.md#GetLiveStreamPlayerInfo) | **Get** `/live_streams/player/{id}/media` | Get live stream media public
[**GetLiveStreamStatisticByStreamMediaId**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStream.md#GetLiveStreamStatisticByStreamMediaId) | **Get** `/live_streams/statistic/{stream_media_id}` | Get live stream statistic by stream media id
[**GetLiveStreamUsage**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStream.md#GetLiveStreamUsage) | **Get** `/live_streams/usage/{stream_id}` | Get usage details for a specific live stream
[**GetStreaming**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStream.md#GetStreaming) | **Get** `/live_streams/{id}/streamings/{stream_id}` | Get live stream media streaming
[**GetStreamings**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStream.md#GetStreamings) | **Get** `/live_streams/{id}/streamings` | Get live stream media streamings
[**GetUserStreamsUsageDetail**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStream.md#GetUserStreamsUsageDetail) | **Get** `/live_streams/usage/streams` | Get paginated list of streams with usage details for current user
[**UpdateLiveStreamKey**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStream.md#UpdateLiveStreamKey) | **Put** `/live_streams/{id}` | Update live stream key
[**UpdateLiveStreamMedia**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStream.md#UpdateLiveStreamMedia) | **Put** `/live_streams/{id}/streamings` | Update live stream media


#### Media


##### Retrieve an instance of the Media API:
```go
secretKey := "YOUR_SECRET_KEY" // Replace with your actual secret key
publicKey := "YOUR_PUBLIC_KEY" // Replace with your public key
apiCreds := aiozstreamsdk.AuthCredentials{
	PublicKey: publicKey,
	SecretKey: secretKey,
}
client := aiozstreamsdk.ClientBuilder(apiCreds).Build()
mediaApi := client.Media
```

##### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Media.md#Create) | **Post** `/media/create` | Create media object
[**Update**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Media.md#Update) | **Patch** `/media/{id}` | update media info
[**Delete**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Media.md#Delete) | **Delete** `/media/{id}` | Delete media
[**UploadThumbnail**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Media.md#UploadThumbnail) | **Post** `/media/{id}/thumbnail` | Upload media thumbnail
[**DeleteThumbnail**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Media.md#DeleteThumbnail) | **Delete** `/media/{id}/thumbnail` | Delete media thumbnail
[**CreateCaption**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Media.md#CreateCaption) | **Post** `/media/{id}/captions/{lan}` | Create a new media caption
[**DeleteCaption**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Media.md#DeleteCaption) | **Delete** `/media/{id}/captions/{lan}` | Delete a media caption
[**GetCaptions**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Media.md#GetCaptions) | **Get** `/media/{id}/captions` | Get media captions
[**GetCost**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Media.md#GetCost) | **Get** `/media/cost` | get media transcoding cost
[**GetDetail**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Media.md#GetDetail) | **Get** `/media/{id}` | get media detail
[**GetMediaList**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Media.md#GetMediaList) | **Post** `/media` | Get user media list
[**GetMediaPlayerInfo**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Media.md#GetMediaPlayerInfo) | **Get** `/media/{id}/player.json` | Get media player info
[**SetDefaultCaption**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Media.md#SetDefaultCaption) | **Patch** `/media/{id}/captions/{lan}` | Set the default caption
[**UploadMediaComplete**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Media.md#UploadMediaComplete) | **Get** `/media/{id}/complete` | Get upload media when complete
[**UploadPart**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Media.md#UploadPart) | **Post** `/media/{id}/part` | Upload part of media


#### MediaChapter


##### Retrieve an instance of the MediaChapter API:
```go
secretKey := "YOUR_SECRET_KEY" // Replace with your actual secret key
publicKey := "YOUR_PUBLIC_KEY" // Replace with your public key
apiCreds := aiozstreamsdk.AuthCredentials{
	PublicKey: publicKey,
	SecretKey: secretKey,
}
client := aiozstreamsdk.ClientBuilder(apiCreds).Build()
mediaChapterApi := client.MediaChapter
```

##### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/MediaChapter.md#Create) | **Post** `/media/{id}/chapters/{lan}` | Create a media chapter
[**Get**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/MediaChapter.md#Get) | **Get** `/media/{id}/chapters` | Get media chapters
[**Delete**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/MediaChapter.md#Delete) | **Delete** `/media/{id}/chapters/{lan}` | Delete a media chapter


#### Players


##### Retrieve an instance of the Players API:
```go
secretKey := "YOUR_SECRET_KEY" // Replace with your actual secret key
publicKey := "YOUR_PUBLIC_KEY" // Replace with your public key
apiCreds := aiozstreamsdk.AuthCredentials{
	PublicKey: publicKey,
	SecretKey: secretKey,
}
client := aiozstreamsdk.ClientBuilder(apiCreds).Build()
playersApi := client.Players
```

##### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Players.md#Create) | **Post** `/players` | Create a player theme
[**Get**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Players.md#Get) | **Get** `/players/{id}` | Get a player theme
[**Update**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Players.md#Update) | **Patch** `/players/{id}` | Update a player theme
[**Delete**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Players.md#Delete) | **Delete** `/players/{id}` | Delete a player theme
[**List**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Players.md#List) | **Get** `/players` | List player themes
[**UploadLogo**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Players.md#UploadLogo) | **Post** `/players/{id}/logo` | Upload a player theme logo
[**DeleteLogo**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Players.md#DeleteLogo) | **Delete** `/players/{id}/logo` | Delete a player theme logo
[**Attach**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Players.md#Attach) | **Post** `/players/add-player` | Add a player theme to a media
[**Detach**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Players.md#Detach) | **Post** `/players/remove-player` | Remove a player theme from a media


#### Playlist


##### Retrieve an instance of the Playlist API:
```go
secretKey := "YOUR_SECRET_KEY" // Replace with your actual secret key
publicKey := "YOUR_PUBLIC_KEY" // Replace with your public key
apiCreds := aiozstreamsdk.AuthCredentials{
	PublicKey: publicKey,
	SecretKey: secretKey,
}
client := aiozstreamsdk.ClientBuilder(apiCreds).Build()
playlistApi := client.Playlist
```

##### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Playlist.md#Create) | **Post** `/playlists/create` | Create a playlist
[**Get**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Playlist.md#Get) | **Get** `/playlists/{id}` | Get a playlist
[**Update**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Playlist.md#Update) | **Patch** `/playlists/{id}` | Update a playlist
[**Delete**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Playlist.md#Delete) | **Delete** `/playlists/{id}` | Delete a playlist
[**List**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Playlist.md#List) | **Post** `/playlists` | List playlists
[**DeleteThumbnail**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Playlist.md#DeleteThumbnail) | **Delete** `/playlists/{id}/thumbnail` | Delete a playlist thumbnail
[**AddMedia**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Playlist.md#AddMedia) | **Post** `/playlists/{id}/items` | Add media to playlists
[**GetPublic**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Playlist.md#GetPublic) | **Get** `/playlists/{id}/player.json` | Get a playlist for the player
[**MoveItem**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Playlist.md#MoveItem) | **Put** `/playlists/{id}/items` | Reorder a playlist
[**RemoveMedia**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Playlist.md#RemoveMedia) | **Delete** `/playlists/{id}/items/{item_id}` | Remove an item from playlists


#### User


##### Retrieve an instance of the User API:
```go
secretKey := "YOUR_SECRET_KEY" // Replace with your actual secret key
publicKey := "YOUR_PUBLIC_KEY" // Replace with your public key
apiCreds := aiozstreamsdk.AuthCredentials{
	PublicKey: publicKey,
	SecretKey: secretKey,
}
client := aiozstreamsdk.ClientBuilder(apiCreds).Build()
userApi := client.User
```

##### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMe**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/User.md#GetMe) | **Get** `/user/me` | Get the current account


#### Webhook


##### Retrieve an instance of the Webhook API:
```go
secretKey := "YOUR_SECRET_KEY" // Replace with your actual secret key
publicKey := "YOUR_PUBLIC_KEY" // Replace with your public key
apiCreds := aiozstreamsdk.AuthCredentials{
	PublicKey: publicKey,
	SecretKey: secretKey,
}
client := aiozstreamsdk.ClientBuilder(apiCreds).Build()
webhookApi := client.Webhook
```

##### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Webhook.md#Create) | **Post** `/webhooks` | Create a webhook
[**Get**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Webhook.md#Get) | **Get** `/webhooks/{id}` | Get a webhook
[**Update**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Webhook.md#Update) | **Patch** `/webhooks/{id}` | Update a webhook
[**Delete**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Webhook.md#Delete) | **Delete** `/webhooks/{id}` | Delete a webhook
[**List**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Webhook.md#List) | **Get** `/webhooks` | List webhooks
[**Check**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Webhook.md#Check) | **Post** `/webhooks/check/{id}` | Send a test event




### Models

 - [AddMediaRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/AddMediaRequest.md)
 - [AggregatedMetricsData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/AggregatedMetricsData.md)
 - [AggregatedMetricsResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/AggregatedMetricsResponse.md)
 - [ApiKey](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/ApiKey.md)
 - [Asset](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Asset.md)
 - [AttachThemeRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/AttachThemeRequest.md)
 - [AudioConfig](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/AudioConfig.md)
 - [Controls](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Controls.md)
 - [CreateApiKeyData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/CreateApiKeyData.md)
 - [CreateApiKeyRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/CreateApiKeyRequest.md)
 - [CreateApiKeyResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/CreateApiKeyResponse.md)
 - [CreateLiveStreamKeyRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/CreateLiveStreamKeyRequest.md)
 - [CreateLiveStreamKeyResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/CreateLiveStreamKeyResponse.md)
 - [CreateMediaCaptionData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/CreateMediaCaptionData.md)
 - [CreateMediaCaptionResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/CreateMediaCaptionResponse.md)
 - [CreateMediaChapterData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/CreateMediaChapterData.md)
 - [CreateMediaChapterResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/CreateMediaChapterResponse.md)
 - [CreateMediaRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/CreateMediaRequest.md)
 - [CreateMediaResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/CreateMediaResponse.md)
 - [CreatePlaylistRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/CreatePlaylistRequest.md)
 - [CreateStreamingRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/CreateStreamingRequest.md)
 - [CreateStreamingResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/CreateStreamingResponse.md)
 - [DataUsage](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/DataUsage.md)
 - [DataUsageData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/DataUsageData.md)
 - [DataUsageRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/DataUsageRequest.md)
 - [DataUsageResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/DataUsageResponse.md)
 - [DeletedAt](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/DeletedAt.md)
 - [EditorProject](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/EditorProject.md)
 - [ErrorBody](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/ErrorBody.md)
 - [GetLiveStreamKeyData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetLiveStreamKeyData.md)
 - [GetLiveStreamKeyResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetLiveStreamKeyResponse.md)
 - [GetLiveStreamKeysListData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetLiveStreamKeysListData.md)
 - [GetLiveStreamKeysListResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetLiveStreamKeysListResponse.md)
 - [GetLiveStreamMediaPublicResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetLiveStreamMediaPublicResponse.md)
 - [GetLiveStreamMediaResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetLiveStreamMediaResponse.md)
 - [GetLiveStreamMediasRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetLiveStreamMediasRequest.md)
 - [GetLiveStreamMediasResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetLiveStreamMediasResponse.md)
 - [GetLiveStreamMulticastResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetLiveStreamMulticastResponse.md)
 - [GetLiveStreamStatisticResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetLiveStreamStatisticResponse.md)
 - [GetMediaCaptionsData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetMediaCaptionsData.md)
 - [GetMediaCaptionsResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetMediaCaptionsResponse.md)
 - [GetMediaChaptersData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetMediaChaptersData.md)
 - [GetMediaChaptersResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetMediaChaptersResponse.md)
 - [GetMediaDetailResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetMediaDetailResponse.md)
 - [GetMediaListData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetMediaListData.md)
 - [GetMediaListRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetMediaListRequest.md)
 - [GetMediaListResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetMediaListResponse.md)
 - [GetMediaPlayerInfoResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetMediaPlayerInfoResponse.md)
 - [GetStreamUsageResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetStreamUsageResponse.md)
 - [GetStreamingResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetStreamingResponse.md)
 - [GetStreamingsResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetStreamingsResponse.md)
 - [GetTranscodeCostData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetTranscodeCostData.md)
 - [GetTranscodeCostResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetTranscodeCostResponse.md)
 - [GetUserStreamsUsageDetailResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetUserStreamsUsageDetailResponse.md)
 - [HighlightChunk](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/HighlightChunk.md)
 - [HighlightChunkStatus](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/HighlightChunkStatus.md)
 - [HighlightClip](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/HighlightClip.md)
 - [HighlightManifest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/HighlightManifest.md)
 - [HighlightMedia](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/HighlightMedia.md)
 - [HighlightMediaStatus](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/HighlightMediaStatus.md)
 - [ListApiKeysData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/ListApiKeysData.md)
 - [ListApiKeysRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/ListApiKeysRequest.md)
 - [ListApiKeysResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/ListApiKeysResponse.md)
 - [ListPlaylistsData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/ListPlaylistsData.md)
 - [ListPlaylistsRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/ListPlaylistsRequest.md)
 - [ListPlaylistsResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/ListPlaylistsResponse.md)
 - [ListThemesData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/ListThemesData.md)
 - [ListThemesRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/ListThemesRequest.md)
 - [ListThemesResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/ListThemesResponse.md)
 - [ListWebhooksData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/ListWebhooksData.md)
 - [ListWebhooksRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/ListWebhooksRequest.md)
 - [ListWebhooksResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/ListWebhooksResponse.md)
 - [LiveStreamAssets](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStreamAssets.md)
 - [LiveStreamKeyData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStreamKeyData.md)
 - [LiveStreamMediaData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStreamMediaData.md)
 - [LiveStreamMediaResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStreamMediaResponse.md)
 - [LiveStreamMediasResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStreamMediasResponse.md)
 - [LiveStreamMulticast](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStreamMulticast.md)
 - [LiveStreamStatisticResp](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStreamStatisticResp.md)
 - [ManifestType](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/ManifestType.md)
 - [MeData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/MeData.md)
 - [MeResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/MeResponse.md)
 - [MediaAssets](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/MediaAssets.md)
 - [MediaCaption](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/MediaCaption.md)
 - [MediaChapter](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/MediaChapter.md)
 - [MediaObject](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/MediaObject.md)
 - [MediaSummary](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/MediaSummary.md)
 - [MediaType](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/MediaType.md)
 - [MediaWatermark](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/MediaWatermark.md)
 - [Metadata](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Metadata.md)
 - [MetricFilter](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/MetricFilter.md)
 - [MetricItem](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/MetricItem.md)
 - [MetricsContext](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/MetricsContext.md)
 - [MetricsPageData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/MetricsPageData.md)
 - [MetricsPageResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/MetricsPageResponse.md)
 - [MetricsRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/MetricsRequest.md)
 - [MoveItemRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/MoveItemRequest.md)
 - [PlayerTheme](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/PlayerTheme.md)
 - [PlayerThemeInput](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/PlayerThemeInput.md)
 - [Playlist](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Playlist.md)
 - [PlaylistData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/PlaylistData.md)
 - [PlaylistItem](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/PlaylistItem.md)
 - [PlaylistItemMedia](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/PlaylistItemMedia.md)
 - [PlaylistResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/PlaylistResponse.md)
 - [QualityConfig](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/QualityConfig.md)
 - [QualityObject](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/QualityObject.md)
 - [RemoveMediaRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/RemoveMediaRequest.md)
 - [RenameApiKeyRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/RenameApiKeyRequest.md)
 - [RenditionUsage](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/RenditionUsage.md)
 - [ResponseError](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/ResponseError.md)
 - [ResponseSuccess](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/ResponseSuccess.md)
 - [SetDefaultCaptionRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/SetDefaultCaptionRequest.md)
 - [StreamUsageSummary](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/StreamUsageSummary.md)
 - [Theme](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Theme.md)
 - [ThemeData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/ThemeData.md)
 - [ThemeResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/ThemeResponse.md)
 - [TimeFrame](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/TimeFrame.md)
 - [UpdateLiveStreamKeyData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/UpdateLiveStreamKeyData.md)
 - [UpdateLiveStreamKeyRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/UpdateLiveStreamKeyRequest.md)
 - [UpdateLiveStreamKeyResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/UpdateLiveStreamKeyResponse.md)
 - [UpdateLiveStreamMediaRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/UpdateLiveStreamMediaRequest.md)
 - [UpdateMediaInfoRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/UpdateMediaInfoRequest.md)
 - [UpsertLiveStreamMulticastInput](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/UpsertLiveStreamMulticastInput.md)
 - [User](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/User.md)
 - [UserStreamsUsageResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/UserStreamsUsageResponse.md)
 - [VideoConfig](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/VideoConfig.md)
 - [VideoCropInfo](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/VideoCropInfo.md)
 - [Webhook](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Webhook.md)
 - [WebhookData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/WebhookData.md)
 - [WebhookResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/WebhookResponse.md)
 - [WriteWebhookRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/WriteWebhookRequest.md)



## Have you gotten use from this API client?

Please take a moment to leave a star on the client ⭐

This helps other users to find the clients and also helps us understand which clients are most popular. Thank you!

