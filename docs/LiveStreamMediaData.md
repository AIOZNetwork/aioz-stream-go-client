# LiveStreamMediaData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assets** | Pointer to [**LiveStreamAssets**](LiveStreamAssets.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LiveStreamKeyId** | Pointer to **string** |  | [optional] 
**Media** | Pointer to [**Media**](Media.md) |  | [optional] 
**Qualities** | Pointer to **[]string** |  | [optional] 
**Save** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewLiveStreamMediaData

`func NewLiveStreamMediaData() *LiveStreamMediaData`

NewLiveStreamMediaData instantiates a new LiveStreamMediaData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveStreamMediaDataWithDefaults

`func NewLiveStreamMediaDataWithDefaults() *LiveStreamMediaData`

NewLiveStreamMediaDataWithDefaults instantiates a new LiveStreamMediaData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssets

`func (o *LiveStreamMediaData) GetAssets() LiveStreamAssets`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *LiveStreamMediaData) GetAssetsOk() (*LiveStreamAssets, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *LiveStreamMediaData) SetAssets(v LiveStreamAssets)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *LiveStreamMediaData) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetCreatedAt

`func (o *LiveStreamMediaData) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LiveStreamMediaData) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LiveStreamMediaData) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LiveStreamMediaData) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDuration

`func (o *LiveStreamMediaData) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *LiveStreamMediaData) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *LiveStreamMediaData) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *LiveStreamMediaData) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetId

`func (o *LiveStreamMediaData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LiveStreamMediaData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LiveStreamMediaData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LiveStreamMediaData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLiveStreamKeyId

`func (o *LiveStreamMediaData) GetLiveStreamKeyId() string`

GetLiveStreamKeyId returns the LiveStreamKeyId field if non-nil, zero value otherwise.

### GetLiveStreamKeyIdOk

`func (o *LiveStreamMediaData) GetLiveStreamKeyIdOk() (*string, bool)`

GetLiveStreamKeyIdOk returns a tuple with the LiveStreamKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveStreamKeyId

`func (o *LiveStreamMediaData) SetLiveStreamKeyId(v string)`

SetLiveStreamKeyId sets LiveStreamKeyId field to given value.

### HasLiveStreamKeyId

`func (o *LiveStreamMediaData) HasLiveStreamKeyId() bool`

HasLiveStreamKeyId returns a boolean if a field has been set.

### GetMedia

`func (o *LiveStreamMediaData) GetMedia() Media`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *LiveStreamMediaData) GetMediaOk() (*Media, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *LiveStreamMediaData) SetMedia(v Media)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *LiveStreamMediaData) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetQualities

`func (o *LiveStreamMediaData) GetQualities() []string`

GetQualities returns the Qualities field if non-nil, zero value otherwise.

### GetQualitiesOk

`func (o *LiveStreamMediaData) GetQualitiesOk() (*[]string, bool)`

GetQualitiesOk returns a tuple with the Qualities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualities

`func (o *LiveStreamMediaData) SetQualities(v []string)`

SetQualities sets Qualities field to given value.

### HasQualities

`func (o *LiveStreamMediaData) HasQualities() bool`

HasQualities returns a boolean if a field has been set.

### GetSave

`func (o *LiveStreamMediaData) GetSave() bool`

GetSave returns the Save field if non-nil, zero value otherwise.

### GetSaveOk

`func (o *LiveStreamMediaData) GetSaveOk() (*bool, bool)`

GetSaveOk returns a tuple with the Save field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSave

`func (o *LiveStreamMediaData) SetSave(v bool)`

SetSave sets Save field to given value.

### HasSave

`func (o *LiveStreamMediaData) HasSave() bool`

HasSave returns a boolean if a field has been set.

### GetStatus

`func (o *LiveStreamMediaData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LiveStreamMediaData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LiveStreamMediaData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LiveStreamMediaData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *LiveStreamMediaData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *LiveStreamMediaData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *LiveStreamMediaData) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *LiveStreamMediaData) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *LiveStreamMediaData) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LiveStreamMediaData) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LiveStreamMediaData) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *LiveStreamMediaData) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserId

`func (o *LiveStreamMediaData) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *LiveStreamMediaData) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *LiveStreamMediaData) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *LiveStreamMediaData) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


