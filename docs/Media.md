# Media

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assets** | Pointer to [**VideoAssets**](VideoAssets.md) |  | [optional] 
**Captions** | Pointer to [**[]VideoCaption**](VideoCaption.md) |  | [optional] 
**Chapters** | Pointer to [**[]VideoChapter**](VideoChapter.md) |  | [optional] 
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
**Size** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**View** | Pointer to **int32** |  | [optional] 

## Methods

### NewMedia

`func NewMedia() *Media`

NewMedia instantiates a new Media object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaWithDefaults

`func NewMediaWithDefaults() *Media`

NewMediaWithDefaults instantiates a new Media object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssets

`func (o *Media) GetAssets() VideoAssets`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *Media) GetAssetsOk() (*VideoAssets, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *Media) SetAssets(v VideoAssets)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *Media) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetCaptions

`func (o *Media) GetCaptions() []VideoCaption`

GetCaptions returns the Captions field if non-nil, zero value otherwise.

### GetCaptionsOk

`func (o *Media) GetCaptionsOk() (*[]VideoCaption, bool)`

GetCaptionsOk returns a tuple with the Captions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptions

`func (o *Media) SetCaptions(v []VideoCaption)`

SetCaptions sets Captions field to given value.

### HasCaptions

`func (o *Media) HasCaptions() bool`

HasCaptions returns a boolean if a field has been set.

### GetChapters

`func (o *Media) GetChapters() []VideoChapter`

GetChapters returns the Chapters field if non-nil, zero value otherwise.

### GetChaptersOk

`func (o *Media) GetChaptersOk() (*[]VideoChapter, bool)`

GetChaptersOk returns a tuple with the Chapters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChapters

`func (o *Media) SetChapters(v []VideoChapter)`

SetChapters sets Chapters field to given value.

### HasChapters

`func (o *Media) HasChapters() bool`

HasChapters returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Media) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Media) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Media) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Media) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *Media) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Media) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Media) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Media) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDuration

`func (o *Media) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *Media) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *Media) SetDuration(v float32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *Media) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetId

`func (o *Media) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Media) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Media) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Media) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsMp4

`func (o *Media) GetIsMp4() bool`

GetIsMp4 returns the IsMp4 field if non-nil, zero value otherwise.

### GetIsMp4Ok

`func (o *Media) GetIsMp4Ok() (*bool, bool)`

GetIsMp4Ok returns a tuple with the IsMp4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMp4

`func (o *Media) SetIsMp4(v bool)`

SetIsMp4 sets IsMp4 field to given value.

### HasIsMp4

`func (o *Media) HasIsMp4() bool`

HasIsMp4 returns a boolean if a field has been set.

### GetIsPublic

`func (o *Media) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *Media) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *Media) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *Media) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetMetadata

`func (o *Media) GetMetadata() []Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Media) GetMetadataOk() (*[]Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Media) SetMetadata(v []Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Media) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPlayerTheme

`func (o *Media) GetPlayerTheme() PlayerTheme`

GetPlayerTheme returns the PlayerTheme field if non-nil, zero value otherwise.

### GetPlayerThemeOk

`func (o *Media) GetPlayerThemeOk() (*PlayerTheme, bool)`

GetPlayerThemeOk returns a tuple with the PlayerTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerTheme

`func (o *Media) SetPlayerTheme(v PlayerTheme)`

SetPlayerTheme sets PlayerTheme field to given value.

### HasPlayerTheme

`func (o *Media) HasPlayerTheme() bool`

HasPlayerTheme returns a boolean if a field has been set.

### GetPlayerThemeId

`func (o *Media) GetPlayerThemeId() string`

GetPlayerThemeId returns the PlayerThemeId field if non-nil, zero value otherwise.

### GetPlayerThemeIdOk

`func (o *Media) GetPlayerThemeIdOk() (*string, bool)`

GetPlayerThemeIdOk returns a tuple with the PlayerThemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerThemeId

`func (o *Media) SetPlayerThemeId(v string)`

SetPlayerThemeId sets PlayerThemeId field to given value.

### HasPlayerThemeId

`func (o *Media) HasPlayerThemeId() bool`

HasPlayerThemeId returns a boolean if a field has been set.

### GetQualities

`func (o *Media) GetQualities() []QualityObject`

GetQualities returns the Qualities field if non-nil, zero value otherwise.

### GetQualitiesOk

`func (o *Media) GetQualitiesOk() (*[]QualityObject, bool)`

GetQualitiesOk returns a tuple with the Qualities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualities

`func (o *Media) SetQualities(v []QualityObject)`

SetQualities sets Qualities field to given value.

### HasQualities

`func (o *Media) HasQualities() bool`

HasQualities returns a boolean if a field has been set.

### GetSize

`func (o *Media) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Media) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Media) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *Media) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStatus

`func (o *Media) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Media) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Media) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Media) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *Media) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Media) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Media) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Media) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTitle

`func (o *Media) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Media) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Media) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Media) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Media) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Media) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Media) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Media) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserId

`func (o *Media) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Media) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Media) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Media) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetView

`func (o *Media) GetView() int32`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *Media) GetViewOk() (*int32, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *Media) SetView(v int32)`

SetView sets View field to given value.

### HasView

`func (o *Media) HasView() bool`

HasView returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


