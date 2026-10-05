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
 - [ApiKey](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/ApiKey.md)
 - [Asset](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Asset.md)
 - [AttachThemeRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/AttachThemeRequest.md)
 - [AudioConfig](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/AudioConfig.md)
 - [Controls](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Controls.md)
 - [CreateApiKeyData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/CreateApiKeyData.md)
 - [CreateApiKeyRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/CreateApiKeyRequest.md)
 - [CreateApiKeyResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/CreateApiKeyResponse.md)
 - [CreateMediaCaptionData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/CreateMediaCaptionData.md)
 - [CreateMediaCaptionResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/CreateMediaCaptionResponse.md)
 - [CreateMediaChapterData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/CreateMediaChapterData.md)
 - [CreateMediaChapterResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/CreateMediaChapterResponse.md)
 - [CreateMediaRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/CreateMediaRequest.md)
 - [CreateMediaResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/CreateMediaResponse.md)
 - [CreatePlaylistRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/CreatePlaylistRequest.md)
 - [DeletedAt](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/DeletedAt.md)
 - [EditorProject](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/EditorProject.md)
 - [ErrorBody](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/ErrorBody.md)
 - [GetMediaCaptionsData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetMediaCaptionsData.md)
 - [GetMediaCaptionsResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetMediaCaptionsResponse.md)
 - [GetMediaChaptersData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetMediaChaptersData.md)
 - [GetMediaChaptersResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetMediaChaptersResponse.md)
 - [GetMediaDetailResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetMediaDetailResponse.md)
 - [GetMediaListData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetMediaListData.md)
 - [GetMediaListRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetMediaListRequest.md)
 - [GetMediaListResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetMediaListResponse.md)
 - [GetMediaPlayerInfoResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetMediaPlayerInfoResponse.md)
 - [GetTranscodeCostData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetTranscodeCostData.md)
 - [GetTranscodeCostResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/GetTranscodeCostResponse.md)
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
 - [ManifestType](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/ManifestType.md)
 - [MediaAssets](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/MediaAssets.md)
 - [MediaCaption](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/MediaCaption.md)
 - [MediaChapter](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/MediaChapter.md)
 - [MediaObject](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/MediaObject.md)
 - [MediaSummary](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/MediaSummary.md)
 - [MediaType](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/MediaType.md)
 - [MediaWatermark](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/MediaWatermark.md)
 - [Metadata](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Metadata.md)
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
 - [ResponseError](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/ResponseError.md)
 - [ResponseSuccess](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/ResponseSuccess.md)
 - [SetDefaultCaptionRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/SetDefaultCaptionRequest.md)
 - [Theme](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Theme.md)
 - [ThemeData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/ThemeData.md)
 - [ThemeResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/ThemeResponse.md)
 - [UpdateMediaInfoRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/UpdateMediaInfoRequest.md)
 - [User](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/User.md)
 - [VideoConfig](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/VideoConfig.md)
 - [VideoCropInfo](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/VideoCropInfo.md)
 - [Webhook](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/Webhook.md)
 - [WebhookData](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/WebhookData.md)
 - [WebhookResponse](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/WebhookResponse.md)
 - [WriteWebhookRequest](https://github.com/AIOZNetwork/aioz-stream-go-client/blob/main/docs/WriteWebhookRequest.md)



## Have you gotten use from this API client?

Please take a moment to leave a star on the client ⭐

This helps other users to find the clients and also helps us understand which clients are most popular. Thank you!

