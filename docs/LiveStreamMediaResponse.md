# LiveStreamMediaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assets** | Pointer to [**LiveStreamAssets**](LiveStreamAssets.md) |  | [optional] 
**AudioBitrate** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CurrentView** | Pointer to **int32** |  | [optional] 
**Duration** | Pointer to **int32** |  | [optional] 
**FrameRate** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LiveStreamKeyId** | Pointer to **string** |  | [optional] 
**Qualities** | Pointer to **[]string** |  | [optional] 
**Save** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**TotalView** | Pointer to **int32** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**Video** | Pointer to [**Media**](Media.md) |  | [optional] 

## Methods

### NewLiveStreamMediaResponse

`func NewLiveStreamMediaResponse() *LiveStreamMediaResponse`

NewLiveStreamMediaResponse instantiates a new LiveStreamMediaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveStreamMediaResponseWithDefaults

`func NewLiveStreamMediaResponseWithDefaults() *LiveStreamMediaResponse`

NewLiveStreamMediaResponseWithDefaults instantiates a new LiveStreamMediaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssets

`func (o *LiveStreamMediaResponse) GetAssets() LiveStreamAssets`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *LiveStreamMediaResponse) GetAssetsOk() (*LiveStreamAssets, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *LiveStreamMediaResponse) SetAssets(v LiveStreamAssets)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *LiveStreamMediaResponse) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetAudioBitrate

`func (o *LiveStreamMediaResponse) GetAudioBitrate() int32`

GetAudioBitrate returns the AudioBitrate field if non-nil, zero value otherwise.

### GetAudioBitrateOk

`func (o *LiveStreamMediaResponse) GetAudioBitrateOk() (*int32, bool)`

GetAudioBitrateOk returns a tuple with the AudioBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioBitrate

`func (o *LiveStreamMediaResponse) SetAudioBitrate(v int32)`

SetAudioBitrate sets AudioBitrate field to given value.

### HasAudioBitrate

`func (o *LiveStreamMediaResponse) HasAudioBitrate() bool`

HasAudioBitrate returns a boolean if a field has been set.

### GetCreatedAt

`func (o *LiveStreamMediaResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LiveStreamMediaResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LiveStreamMediaResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LiveStreamMediaResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrentView

`func (o *LiveStreamMediaResponse) GetCurrentView() int32`

GetCurrentView returns the CurrentView field if non-nil, zero value otherwise.

### GetCurrentViewOk

`func (o *LiveStreamMediaResponse) GetCurrentViewOk() (*int32, bool)`

GetCurrentViewOk returns a tuple with the CurrentView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentView

`func (o *LiveStreamMediaResponse) SetCurrentView(v int32)`

SetCurrentView sets CurrentView field to given value.

### HasCurrentView

`func (o *LiveStreamMediaResponse) HasCurrentView() bool`

HasCurrentView returns a boolean if a field has been set.

### GetDuration

`func (o *LiveStreamMediaResponse) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *LiveStreamMediaResponse) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *LiveStreamMediaResponse) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *LiveStreamMediaResponse) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetFrameRate

`func (o *LiveStreamMediaResponse) GetFrameRate() int32`

GetFrameRate returns the FrameRate field if non-nil, zero value otherwise.

### GetFrameRateOk

`func (o *LiveStreamMediaResponse) GetFrameRateOk() (*int32, bool)`

GetFrameRateOk returns a tuple with the FrameRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameRate

`func (o *LiveStreamMediaResponse) SetFrameRate(v int32)`

SetFrameRate sets FrameRate field to given value.

### HasFrameRate

`func (o *LiveStreamMediaResponse) HasFrameRate() bool`

HasFrameRate returns a boolean if a field has been set.

### GetId

`func (o *LiveStreamMediaResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LiveStreamMediaResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LiveStreamMediaResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LiveStreamMediaResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLiveStreamKeyId

`func (o *LiveStreamMediaResponse) GetLiveStreamKeyId() string`

GetLiveStreamKeyId returns the LiveStreamKeyId field if non-nil, zero value otherwise.

### GetLiveStreamKeyIdOk

`func (o *LiveStreamMediaResponse) GetLiveStreamKeyIdOk() (*string, bool)`

GetLiveStreamKeyIdOk returns a tuple with the LiveStreamKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveStreamKeyId

`func (o *LiveStreamMediaResponse) SetLiveStreamKeyId(v string)`

SetLiveStreamKeyId sets LiveStreamKeyId field to given value.

### HasLiveStreamKeyId

`func (o *LiveStreamMediaResponse) HasLiveStreamKeyId() bool`

HasLiveStreamKeyId returns a boolean if a field has been set.

### GetQualities

`func (o *LiveStreamMediaResponse) GetQualities() []string`

GetQualities returns the Qualities field if non-nil, zero value otherwise.

### GetQualitiesOk

`func (o *LiveStreamMediaResponse) GetQualitiesOk() (*[]string, bool)`

GetQualitiesOk returns a tuple with the Qualities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualities

`func (o *LiveStreamMediaResponse) SetQualities(v []string)`

SetQualities sets Qualities field to given value.

### HasQualities

`func (o *LiveStreamMediaResponse) HasQualities() bool`

HasQualities returns a boolean if a field has been set.

### GetSave

`func (o *LiveStreamMediaResponse) GetSave() bool`

GetSave returns the Save field if non-nil, zero value otherwise.

### GetSaveOk

`func (o *LiveStreamMediaResponse) GetSaveOk() (*bool, bool)`

GetSaveOk returns a tuple with the Save field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSave

`func (o *LiveStreamMediaResponse) SetSave(v bool)`

SetSave sets Save field to given value.

### HasSave

`func (o *LiveStreamMediaResponse) HasSave() bool`

HasSave returns a boolean if a field has been set.

### GetStatus

`func (o *LiveStreamMediaResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LiveStreamMediaResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LiveStreamMediaResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LiveStreamMediaResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *LiveStreamMediaResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *LiveStreamMediaResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *LiveStreamMediaResponse) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *LiveStreamMediaResponse) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTotalView

`func (o *LiveStreamMediaResponse) GetTotalView() int32`

GetTotalView returns the TotalView field if non-nil, zero value otherwise.

### GetTotalViewOk

`func (o *LiveStreamMediaResponse) GetTotalViewOk() (*int32, bool)`

GetTotalViewOk returns a tuple with the TotalView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalView

`func (o *LiveStreamMediaResponse) SetTotalView(v int32)`

SetTotalView sets TotalView field to given value.

### HasTotalView

`func (o *LiveStreamMediaResponse) HasTotalView() bool`

HasTotalView returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *LiveStreamMediaResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LiveStreamMediaResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LiveStreamMediaResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *LiveStreamMediaResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserId

`func (o *LiveStreamMediaResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *LiveStreamMediaResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *LiveStreamMediaResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *LiveStreamMediaResponse) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetVideo

`func (o *LiveStreamMediaResponse) GetVideo() Media`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *LiveStreamMediaResponse) GetVideoOk() (*Media, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *LiveStreamMediaResponse) SetVideo(v Media)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *LiveStreamMediaResponse) HasVideo() bool`

HasVideo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


