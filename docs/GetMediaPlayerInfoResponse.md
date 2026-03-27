# GetMediaPlayerInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assets** | Pointer to [**MediaAssets**](MediaAssets.md) |  | [optional] 
**Captions** | Pointer to [**[]MediaCaption**](MediaCaption.md) |  | [optional] 
**Chapters** | Pointer to [**[]MediaChapter**](MediaChapter.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to **float32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsMp4** | Pointer to **bool** |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to [**[]Metadata**](Metadata.md) |  | [optional] 
**PlayerTheme** | Pointer to [**PlayerTheme**](PlayerTheme.md) |  | [optional] 
**PlayerThemeId** | Pointer to **string** |  | [optional] 
**Qualities** | Pointer to [**[]QualityObject**](QualityObject.md) |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewGetMediaPlayerInfoResponse

`func NewGetMediaPlayerInfoResponse() *GetMediaPlayerInfoResponse`

NewGetMediaPlayerInfoResponse instantiates a new GetMediaPlayerInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMediaPlayerInfoResponseWithDefaults

`func NewGetMediaPlayerInfoResponseWithDefaults() *GetMediaPlayerInfoResponse`

NewGetMediaPlayerInfoResponseWithDefaults instantiates a new GetMediaPlayerInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssets

`func (o *GetMediaPlayerInfoResponse) GetAssets() MediaAssets`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *GetMediaPlayerInfoResponse) GetAssetsOk() (*MediaAssets, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *GetMediaPlayerInfoResponse) SetAssets(v MediaAssets)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *GetMediaPlayerInfoResponse) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetCaptions

`func (o *GetMediaPlayerInfoResponse) GetCaptions() []MediaCaption`

GetCaptions returns the Captions field if non-nil, zero value otherwise.

### GetCaptionsOk

`func (o *GetMediaPlayerInfoResponse) GetCaptionsOk() (*[]MediaCaption, bool)`

GetCaptionsOk returns a tuple with the Captions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptions

`func (o *GetMediaPlayerInfoResponse) SetCaptions(v []MediaCaption)`

SetCaptions sets Captions field to given value.

### HasCaptions

`func (o *GetMediaPlayerInfoResponse) HasCaptions() bool`

HasCaptions returns a boolean if a field has been set.

### GetChapters

`func (o *GetMediaPlayerInfoResponse) GetChapters() []MediaChapter`

GetChapters returns the Chapters field if non-nil, zero value otherwise.

### GetChaptersOk

`func (o *GetMediaPlayerInfoResponse) GetChaptersOk() (*[]MediaChapter, bool)`

GetChaptersOk returns a tuple with the Chapters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChapters

`func (o *GetMediaPlayerInfoResponse) SetChapters(v []MediaChapter)`

SetChapters sets Chapters field to given value.

### HasChapters

`func (o *GetMediaPlayerInfoResponse) HasChapters() bool`

HasChapters returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetMediaPlayerInfoResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetMediaPlayerInfoResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetMediaPlayerInfoResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetMediaPlayerInfoResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *GetMediaPlayerInfoResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetMediaPlayerInfoResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetMediaPlayerInfoResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetMediaPlayerInfoResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDuration

`func (o *GetMediaPlayerInfoResponse) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *GetMediaPlayerInfoResponse) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *GetMediaPlayerInfoResponse) SetDuration(v float32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *GetMediaPlayerInfoResponse) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetId

`func (o *GetMediaPlayerInfoResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetMediaPlayerInfoResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetMediaPlayerInfoResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetMediaPlayerInfoResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsMp4

`func (o *GetMediaPlayerInfoResponse) GetIsMp4() bool`

GetIsMp4 returns the IsMp4 field if non-nil, zero value otherwise.

### GetIsMp4Ok

`func (o *GetMediaPlayerInfoResponse) GetIsMp4Ok() (*bool, bool)`

GetIsMp4Ok returns a tuple with the IsMp4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMp4

`func (o *GetMediaPlayerInfoResponse) SetIsMp4(v bool)`

SetIsMp4 sets IsMp4 field to given value.

### HasIsMp4

`func (o *GetMediaPlayerInfoResponse) HasIsMp4() bool`

HasIsMp4 returns a boolean if a field has been set.

### GetIsPublic

`func (o *GetMediaPlayerInfoResponse) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *GetMediaPlayerInfoResponse) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *GetMediaPlayerInfoResponse) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *GetMediaPlayerInfoResponse) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetMetadata

`func (o *GetMediaPlayerInfoResponse) GetMetadata() []Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetMediaPlayerInfoResponse) GetMetadataOk() (*[]Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetMediaPlayerInfoResponse) SetMetadata(v []Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetMediaPlayerInfoResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPlayerTheme

`func (o *GetMediaPlayerInfoResponse) GetPlayerTheme() PlayerTheme`

GetPlayerTheme returns the PlayerTheme field if non-nil, zero value otherwise.

### GetPlayerThemeOk

`func (o *GetMediaPlayerInfoResponse) GetPlayerThemeOk() (*PlayerTheme, bool)`

GetPlayerThemeOk returns a tuple with the PlayerTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerTheme

`func (o *GetMediaPlayerInfoResponse) SetPlayerTheme(v PlayerTheme)`

SetPlayerTheme sets PlayerTheme field to given value.

### HasPlayerTheme

`func (o *GetMediaPlayerInfoResponse) HasPlayerTheme() bool`

HasPlayerTheme returns a boolean if a field has been set.

### GetPlayerThemeId

`func (o *GetMediaPlayerInfoResponse) GetPlayerThemeId() string`

GetPlayerThemeId returns the PlayerThemeId field if non-nil, zero value otherwise.

### GetPlayerThemeIdOk

`func (o *GetMediaPlayerInfoResponse) GetPlayerThemeIdOk() (*string, bool)`

GetPlayerThemeIdOk returns a tuple with the PlayerThemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerThemeId

`func (o *GetMediaPlayerInfoResponse) SetPlayerThemeId(v string)`

SetPlayerThemeId sets PlayerThemeId field to given value.

### HasPlayerThemeId

`func (o *GetMediaPlayerInfoResponse) HasPlayerThemeId() bool`

HasPlayerThemeId returns a boolean if a field has been set.

### GetQualities

`func (o *GetMediaPlayerInfoResponse) GetQualities() []QualityObject`

GetQualities returns the Qualities field if non-nil, zero value otherwise.

### GetQualitiesOk

`func (o *GetMediaPlayerInfoResponse) GetQualitiesOk() (*[]QualityObject, bool)`

GetQualitiesOk returns a tuple with the Qualities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualities

`func (o *GetMediaPlayerInfoResponse) SetQualities(v []QualityObject)`

SetQualities sets Qualities field to given value.

### HasQualities

`func (o *GetMediaPlayerInfoResponse) HasQualities() bool`

HasQualities returns a boolean if a field has been set.

### GetSize

`func (o *GetMediaPlayerInfoResponse) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetMediaPlayerInfoResponse) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetMediaPlayerInfoResponse) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *GetMediaPlayerInfoResponse) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStatus

`func (o *GetMediaPlayerInfoResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetMediaPlayerInfoResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetMediaPlayerInfoResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetMediaPlayerInfoResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *GetMediaPlayerInfoResponse) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetMediaPlayerInfoResponse) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetMediaPlayerInfoResponse) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetMediaPlayerInfoResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTitle

`func (o *GetMediaPlayerInfoResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetMediaPlayerInfoResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetMediaPlayerInfoResponse) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GetMediaPlayerInfoResponse) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetMediaPlayerInfoResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetMediaPlayerInfoResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetMediaPlayerInfoResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetMediaPlayerInfoResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserId

`func (o *GetMediaPlayerInfoResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetMediaPlayerInfoResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetMediaPlayerInfoResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *GetMediaPlayerInfoResponse) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


