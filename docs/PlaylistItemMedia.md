# PlaylistItemMedia

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Captions** | Pointer to [**[]MediaCaption**](MediaCaption.md) |  | [optional] 
**Chapters** | Pointer to [**[]MediaChapter**](MediaChapter.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to **float32** |  | [optional] 
**HlsUrl** | Pointer to **string** |  | [optional] 
**Qualities** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**ThumbnailUrl** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewPlaylistItemMedia

`func NewPlaylistItemMedia() *PlaylistItemMedia`

NewPlaylistItemMedia instantiates a new PlaylistItemMedia object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaylistItemMediaWithDefaults

`func NewPlaylistItemMediaWithDefaults() *PlaylistItemMedia`

NewPlaylistItemMediaWithDefaults instantiates a new PlaylistItemMedia object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaptions

`func (o *PlaylistItemMedia) GetCaptions() []MediaCaption`

GetCaptions returns the Captions field if non-nil, zero value otherwise.

### GetCaptionsOk

`func (o *PlaylistItemMedia) GetCaptionsOk() (*[]MediaCaption, bool)`

GetCaptionsOk returns a tuple with the Captions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptions

`func (o *PlaylistItemMedia) SetCaptions(v []MediaCaption)`

SetCaptions sets Captions field to given value.

### HasCaptions

`func (o *PlaylistItemMedia) HasCaptions() bool`

HasCaptions returns a boolean if a field has been set.

### GetChapters

`func (o *PlaylistItemMedia) GetChapters() []MediaChapter`

GetChapters returns the Chapters field if non-nil, zero value otherwise.

### GetChaptersOk

`func (o *PlaylistItemMedia) GetChaptersOk() (*[]MediaChapter, bool)`

GetChaptersOk returns a tuple with the Chapters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChapters

`func (o *PlaylistItemMedia) SetChapters(v []MediaChapter)`

SetChapters sets Chapters field to given value.

### HasChapters

`func (o *PlaylistItemMedia) HasChapters() bool`

HasChapters returns a boolean if a field has been set.

### GetDescription

`func (o *PlaylistItemMedia) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlaylistItemMedia) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PlaylistItemMedia) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PlaylistItemMedia) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDuration

`func (o *PlaylistItemMedia) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *PlaylistItemMedia) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *PlaylistItemMedia) SetDuration(v float32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *PlaylistItemMedia) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetHlsUrl

`func (o *PlaylistItemMedia) GetHlsUrl() string`

GetHlsUrl returns the HlsUrl field if non-nil, zero value otherwise.

### GetHlsUrlOk

`func (o *PlaylistItemMedia) GetHlsUrlOk() (*string, bool)`

GetHlsUrlOk returns a tuple with the HlsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHlsUrl

`func (o *PlaylistItemMedia) SetHlsUrl(v string)`

SetHlsUrl sets HlsUrl field to given value.

### HasHlsUrl

`func (o *PlaylistItemMedia) HasHlsUrl() bool`

HasHlsUrl returns a boolean if a field has been set.

### GetQualities

`func (o *PlaylistItemMedia) GetQualities() string`

GetQualities returns the Qualities field if non-nil, zero value otherwise.

### GetQualitiesOk

`func (o *PlaylistItemMedia) GetQualitiesOk() (*string, bool)`

GetQualitiesOk returns a tuple with the Qualities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualities

`func (o *PlaylistItemMedia) SetQualities(v string)`

SetQualities sets Qualities field to given value.

### HasQualities

`func (o *PlaylistItemMedia) HasQualities() bool`

HasQualities returns a boolean if a field has been set.

### GetSize

`func (o *PlaylistItemMedia) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PlaylistItemMedia) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PlaylistItemMedia) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *PlaylistItemMedia) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetThumbnailUrl

`func (o *PlaylistItemMedia) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *PlaylistItemMedia) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *PlaylistItemMedia) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *PlaylistItemMedia) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### GetTitle

`func (o *PlaylistItemMedia) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PlaylistItemMedia) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PlaylistItemMedia) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PlaylistItemMedia) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


