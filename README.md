<!--<documentation_excluded>-->
<h1 align="center">AIOZ Stream Go client</h1>

AIOZ Stream is the video infrastructure for product builders. Lightning fast video APIs for integrating, scaling, and managing on-demand & low latency live streaming features in your app.

## Project description

AIOZ Stream's Go client streamlines the coding process. Chunking files is handled for you, as is pagination and refreshing your tokens.

## Getting started

### Installation
```bash
go get github.com/AIOZNetwork/aioz-stream-go-client
```


### Code sample

For a more advanced usage you can checkout the rest of the documentation in the [docs directory](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs)

```golang

import (
	"context"
	"fmt"
	"os"
 
	aiozstreamsdk "github.com/AIOZNetwork/aioz-stream-go-client"
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
[**GetAggregatedMetrics**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Analytics.md#GetAggregatedMetrics) | **Post** `/analytics/metrics/data/{metric}/{aggregation}` | Get aggregated metrics
[**GetBreakdownMetrics**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Analytics.md#GetBreakdownMetrics) | **Post** `/analytics/metrics/bucket/{metric}/{breakdown}` | Get breakdown metrics
[**GetDataUsage**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Analytics.md#GetDataUsage) | **Get** `/analytics/data` | Get data usage
[**GetOvertimeMetrics**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Analytics.md#GetOvertimeMetrics) | **Post** `/analytics/metrics/timeseries/{metric}/{interval}` | Get overtime metrics
[**GetStatisticMedias**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Analytics.md#GetStatisticMedias) | **Get** `/analytics/media` | Get statistic media


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
[**Update**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/ApiKey.md#Update) | **Patch** `/api_keys/{id}` | Rename api key
[**Delete**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/ApiKey.md#Delete) | **Delete** `/api_keys/{id}` | Delete API key
[**List**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/ApiKey.md#List) | **Get** `/api_keys` | Get list API keys


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
[**AddMulticast**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStream.md#AddMulticast) | **Post** `/live_streams/multicast/{stream_key}` | Add live stream multicast
[**CreateLiveStreamKey**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStream.md#CreateLiveStreamKey) | **Post** `/live_streams` | Create live stream key
[**CreateStreaming**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStream.md#CreateStreaming) | **Post** `/live_streams/{id}/streamings` | Create a new live stream media
[**DeleteLiveStreamKey**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStream.md#DeleteLiveStreamKey) | **Delete** `/live_streams/{id}` | Delete live stream key
[**DeleteLiveStreamMulticast**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStream.md#DeleteLiveStreamMulticast) | **Delete** `/live_streams/multicast/{stream_key}` | Delete live stream multicast
[**DeleteStreaming**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStream.md#DeleteStreaming) | **Delete** `/live_streams/{id}/streamings/{stream_id}` | Delete live stream video
[**GetLiveStreamKey**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStream.md#GetLiveStreamKey) | **Get** `/live_streams/{id}` | Get live stream key
[**GetLiveStreamKeys**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStream.md#GetLiveStreamKeys) | **Get** `/live_streams` | Get live stream key list
[**GetLiveStreamMulticastByStreamKey**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStream.md#GetLiveStreamMulticastByStreamKey) | **Get** `/live_streams/multicast/{stream_key}` | Get live stream multicast by stream key
[**GetLiveStreamPlayerInfo**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStream.md#GetLiveStreamPlayerInfo) | **Get** `/live_streams/player/{id}/media` | Get live stream video public
[**GetLiveStreamStatisticByStreamMediaId**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStream.md#GetLiveStreamStatisticByStreamMediaId) | **Get** `/live_streams/statistic/{stream_media_id}` | Get live stream statistic by stream media id
[**GetLiveStreamVideo**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStream.md#GetLiveStreamVideo) | **Get** `/live_streams/{id}/video` | Get live stream video
[**GetMedias**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStream.md#GetMedias) | **Post** `/live_streams/{id}/media` | Get live stream media
[**GetStreaming**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStream.md#GetStreaming) | **Get** `/live_streams/{id}/streamings/{stream_id}` | Get live stream media streaming
[**GetStreamings**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStream.md#GetStreamings) | **Get** `/live_streams/{id}/streamings` | Get live stream media streamings
[**UpdateLiveStreamKey**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStream.md#UpdateLiveStreamKey) | **Put** `/live_streams/{id}` | Update live stream key
[**UpdateMedia**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStream.md#UpdateMedia) | **Put** `/live_streams/{id}/streamings` | Update live stream media


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
[**GetMediaList**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Media.md#GetMediaList) | **Post** `/media` | Get user videos list
[**GetMediaPlayerInfo**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Media.md#GetMediaPlayerInfo) | **Get** `/media/{id}/player.json` | Get media object
[**SetDefaultCaption**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Media.md#SetDefaultCaption) | **Patch** `/media/{id}/captions/{lan}` | Set default caption
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
[**Delete**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/MediaChapter.md#Delete) | **Delete** `/media/{id}/chapters/{lan}` | Delete a video chapter


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
[**Get**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Players.md#Get) | **Get** `/players/{id}` | Get a player theme by ID
[**Update**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Players.md#Update) | **Patch** `/players/{id}` | Update a player theme by ID
[**Delete**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Players.md#Delete) | **Delete** `/players/{id}` | Delete a player theme by ID
[**List**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Players.md#List) | **Get** `/players` | List all player themes
[**UploadLogo**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Players.md#UploadLogo) | **Post** `/players/{id}/logo` | Upload a logo for a player theme by ID
[**DeleteLogo**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Players.md#DeleteLogo) | **Delete** `/players/{id}/logo` | Delete a logo for a player theme by ID
[**AddPlayer**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Players.md#AddPlayer) | **Post** `/players/add-player` | Add a player theme to a video
[**RemovePlayer**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Players.md#RemovePlayer) | **Post** `/players/remove-player` | Remove a player theme from a video


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
[**AddVideoToPlaylist**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Playlist.md#AddVideoToPlaylist) | **Post** `/playlists/{id}/items` | Add a video to a playlist
[**CreatePlaylist**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Playlist.md#CreatePlaylist) | **Post** `/playlists/create` | Create a playlist
[**DeletePlaylistById**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Playlist.md#DeletePlaylistById) | **Delete** `/playlists/{id}` | Delete a playlist by ID
[**DeletePlaylistThumbnail**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Playlist.md#DeletePlaylistThumbnail) | **Delete** `/playlists/{id}/thumbnail` | Delete a playlist thumbnail
[**GetPlaylistById**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Playlist.md#GetPlaylistById) | **Get** `/playlists/{id}` | Get playlist by ID
[**GetPlaylistPublicInfo**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Playlist.md#GetPlaylistPublicInfo) | **Get** `/playlists/{id}/player.json` | Get a playlist public
[**GetPlaylists**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Playlist.md#GetPlaylists) | **Post** `/playlists` | Get user&#39;s playlists
[**MoveVideoInPlaylist**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Playlist.md#MoveVideoInPlaylist) | **Put** `/playlists/{id}/items` | Move a video in a playlist
[**RemoveMediaFromPlaylist**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Playlist.md#RemoveMediaFromPlaylist) | **Delete** `/playlists/{id}/items/{item_id}` | Remove a media from a playlist
[**UpdatePlaylist**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Playlist.md#UpdatePlaylist) | **Patch** `/playlists/{id}` | Update a playlist


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
[**GetMe**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/User.md#GetMe) | **Get** `/user/me` | Get me


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
[**Create**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Webhook.md#Create) | **Post** `/webhooks` | Create webhook
[**Get**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Webhook.md#Get) | **Get** `/webhooks/{id}` | Get user&#39;s webhook by id
[**Update**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Webhook.md#Update) | **Patch** `/webhooks/{id}` | Update event webhook
[**Delete**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Webhook.md#Delete) | **Delete** `/webhooks/{id}` | Delete webhook
[**List**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Webhook.md#List) | **Get** `/webhooks` | Get list webhooks
[**Check**](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Webhook.md#Check) | **Post** `/webhooks/check/{id}` | Check webhook by id




### Models

 - [AddMediaToPlaylistRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/AddMediaToPlaylistRequest.md)
 - [AddPlayerThemesToVideoRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/AddPlayerThemesToVideoRequest.md)
 - [ApiKey](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/ApiKey.md)
 - [Asset](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Asset.md)
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
 - [CreatePlayerThemeRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/CreatePlayerThemeRequest.md)
 - [CreatePlayerThemesData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/CreatePlayerThemesData.md)
 - [CreatePlayerThemesResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/CreatePlayerThemesResponse.md)
 - [CreatePlaylistData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/CreatePlaylistData.md)
 - [CreatePlaylistRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/CreatePlaylistRequest.md)
 - [CreatePlaylistResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/CreatePlaylistResponse.md)
 - [CreateStreamingRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/CreateStreamingRequest.md)
 - [CreateStreamingResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/CreateStreamingResponse.md)
 - [CreateWebhookData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/CreateWebhookData.md)
 - [CreateWebhookRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/CreateWebhookRequest.md)
 - [CreateWebhookResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/CreateWebhookResponse.md)
 - [DataUsage](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/DataUsage.md)
 - [GetAggregatedMetricsData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetAggregatedMetricsData.md)
 - [GetAggregatedMetricsResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetAggregatedMetricsResponse.md)
 - [GetApiKeysData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetApiKeysData.md)
 - [GetApiKeysResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetApiKeysResponse.md)
 - [GetBreakdownMetricsData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetBreakdownMetricsData.md)
 - [GetBreakdownMetricsRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetBreakdownMetricsRequest.md)
 - [GetBreakdownMetricsResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetBreakdownMetricsResponse.md)
 - [GetDataUsageData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetDataUsageData.md)
 - [GetDataUsageResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetDataUsageResponse.md)
 - [GetLiveStreamKeyData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetLiveStreamKeyData.md)
 - [GetLiveStreamKeyResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetLiveStreamKeyResponse.md)
 - [GetLiveStreamKeysListData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetLiveStreamKeysListData.md)
 - [GetLiveStreamKeysListResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetLiveStreamKeysListResponse.md)
 - [GetLiveStreamMediasRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetLiveStreamMediasRequest.md)
 - [GetLiveStreamMediasResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetLiveStreamMediasResponse.md)
 - [GetLiveStreamMulticastResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetLiveStreamMulticastResponse.md)
 - [GetLiveStreamStatisticResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetLiveStreamStatisticResponse.md)
 - [GetLiveStreamVideoPublicResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetLiveStreamVideoPublicResponse.md)
 - [GetLiveStreamVideoResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetLiveStreamVideoResponse.md)
 - [GetMeResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetMeResponse.md)
 - [GetMediaCaptionsData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetMediaCaptionsData.md)
 - [GetMediaCaptionsResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetMediaCaptionsResponse.md)
 - [GetMediaChaptersData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetMediaChaptersData.md)
 - [GetMediaChaptersResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetMediaChaptersResponse.md)
 - [GetMediaDetailResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetMediaDetailResponse.md)
 - [GetMediaListData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetMediaListData.md)
 - [GetMediaListRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetMediaListRequest.md)
 - [GetMediaListResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetMediaListResponse.md)
 - [GetMediaPlayerInfoResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetMediaPlayerInfoResponse.md)
 - [GetOvertimeMetricsData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetOvertimeMetricsData.md)
 - [GetOvertimeMetricsResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetOvertimeMetricsResponse.md)
 - [GetPlayerThemeByIdData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetPlayerThemeByIdData.md)
 - [GetPlayerThemeByIdResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetPlayerThemeByIdResponse.md)
 - [GetPlayerThemeData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetPlayerThemeData.md)
 - [GetPlayerThemeResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetPlayerThemeResponse.md)
 - [GetPlaylistByIdData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetPlaylistByIdData.md)
 - [GetPlaylistByIdResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetPlaylistByIdResponse.md)
 - [GetPlaylistListData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetPlaylistListData.md)
 - [GetPlaylistListRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetPlaylistListRequest.md)
 - [GetPlaylistListResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetPlaylistListResponse.md)
 - [GetStatisticMediasData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetStatisticMediasData.md)
 - [GetStatisticMediasResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetStatisticMediasResponse.md)
 - [GetStreamingResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetStreamingResponse.md)
 - [GetStreamingsResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetStreamingsResponse.md)
 - [GetTranscodeCostData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetTranscodeCostData.md)
 - [GetTranscodeCostResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetTranscodeCostResponse.md)
 - [GetUserData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetUserData.md)
 - [GetUserWebhookData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetUserWebhookData.md)
 - [GetUserWebhookResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetUserWebhookResponse.md)
 - [GetWebhooksListData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetWebhooksListData.md)
 - [GetWebhooksListResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetWebhooksListResponse.md)
 - [InternalControllersGetAggreatedMetricsMetricsRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/InternalControllersGetAggreatedMetricsMetricsRequest.md)
 - [LiveStreamAssets](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStreamAssets.md)
 - [LiveStreamKeyData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStreamKeyData.md)
 - [LiveStreamMediaData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStreamMediaData.md)
 - [LiveStreamMediaResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStreamMediaResponse.md)
 - [LiveStreamMediasResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStreamMediasResponse.md)
 - [LiveStreamMulticast](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStreamMulticast.md)
 - [LiveStreamStatisticResp](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/LiveStreamStatisticResp.md)
 - [Media](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Media.md)
 - [MediaAssets](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/MediaAssets.md)
 - [MediaCaption](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/MediaCaption.md)
 - [MediaChapter](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/MediaChapter.md)
 - [Metadata](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Metadata.md)
 - [MetricFilter](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/MetricFilter.md)
 - [MetricItem](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/MetricItem.md)
 - [MetricsContext](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/MetricsContext.md)
 - [MoveVideoInPlaylistRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/MoveVideoInPlaylistRequest.md)
 - [PlayerTheme](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/PlayerTheme.md)
 - [Playlist](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Playlist.md)
 - [PlaylistItem](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/PlaylistItem.md)
 - [PlaylistItemMedia](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/PlaylistItemMedia.md)
 - [PublicPlaylistObject](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/PublicPlaylistObject.md)
 - [QualityConfig](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/QualityConfig.md)
 - [QualityObject](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/QualityObject.md)
 - [RemoveMediasFromPlaylistRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/RemoveMediasFromPlaylistRequest.md)
 - [RemovePlayerThemesFromMediaRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/RemovePlayerThemesFromMediaRequest.md)
 - [RenameAPIKeyRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/RenameAPIKeyRequest.md)
 - [RequestCreateCaption](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/RequestCreateCaption.md)
 - [ResponseError](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/ResponseError.md)
 - [ResponseSuccess](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/ResponseSuccess.md)
 - [Theme](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Theme.md)
 - [TimeFrame](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/TimeFrame.md)
 - [UpdateLiveStreamKeyData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/UpdateLiveStreamKeyData.md)
 - [UpdateLiveStreamKeyRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/UpdateLiveStreamKeyRequest.md)
 - [UpdateLiveStreamKeyResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/UpdateLiveStreamKeyResponse.md)
 - [UpdateLiveStreamMediaRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/UpdateLiveStreamMediaRequest.md)
 - [UpdateMediaInfoRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/UpdateMediaInfoRequest.md)
 - [UpdatePlayerThemeRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/UpdatePlayerThemeRequest.md)
 - [UpdatePlayerThemeResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/UpdatePlayerThemeResponse.md)
 - [UpdateWebhookRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/UpdateWebhookRequest.md)
 - [UploadLogoByIdResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/UploadLogoByIdResponse.md)
 - [UpsertLiveStreamMulticastInput](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/UpsertLiveStreamMulticastInput.md)
 - [User](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/User.md)
 - [VideoConfig](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/VideoConfig.md)
 - [VideoWatermark](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/VideoWatermark.md)
 - [Webhook](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Webhook.md)



## Have you gotten use from this API client?

Please take a moment to leave a star on the client ⭐

This helps other users to find the clients and also helps us understand which clients are most popular. Thank you!

