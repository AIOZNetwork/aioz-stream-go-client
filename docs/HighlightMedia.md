# HighlightMedia

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**DeletedAt** | Pointer to [**DeletedAt**](DeletedAt.md) |  | [optional] 
**EditorProjectId** | Pointer to **string** |  | [optional] 
**FileMd5** | Pointer to **string** |  | [optional] 
**FileName** | Pointer to **string** |  | [optional] 
**FileSize** | Pointer to **int32** |  | [optional] 
**HighlightChunks** | Pointer to [**[]HighlightChunk**](HighlightChunk.md) |  | [optional] 
**HighlightClips** | Pointer to [**[]HighlightClip**](HighlightClip.md) |  | [optional] 
**HighlightManifests** | Pointer to [**[]HighlightManifest**](HighlightManifest.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**MediaType** | Pointer to [**MediaType**](MediaType.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**MimeType** | Pointer to **string** |  | [optional] 
**ObjUrl** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**HighlightMediaStatus**](HighlightMediaStatus.md) |  | [optional] 
**TotalChunk** | Pointer to **int32** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewHighlightMedia

`func NewHighlightMedia() *HighlightMedia`

NewHighlightMedia instantiates a new HighlightMedia object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHighlightMediaWithDefaults

`func NewHighlightMediaWithDefaults() *HighlightMedia`

NewHighlightMediaWithDefaults instantiates a new HighlightMedia object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *HighlightMedia) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HighlightMedia) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HighlightMedia) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *HighlightMedia) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *HighlightMedia) GetDeletedAt() DeletedAt`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *HighlightMedia) GetDeletedAtOk() (*DeletedAt, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *HighlightMedia) SetDeletedAt(v DeletedAt)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *HighlightMedia) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetEditorProjectId

`func (o *HighlightMedia) GetEditorProjectId() string`

GetEditorProjectId returns the EditorProjectId field if non-nil, zero value otherwise.

### GetEditorProjectIdOk

`func (o *HighlightMedia) GetEditorProjectIdOk() (*string, bool)`

GetEditorProjectIdOk returns a tuple with the EditorProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditorProjectId

`func (o *HighlightMedia) SetEditorProjectId(v string)`

SetEditorProjectId sets EditorProjectId field to given value.

### HasEditorProjectId

`func (o *HighlightMedia) HasEditorProjectId() bool`

HasEditorProjectId returns a boolean if a field has been set.

### GetFileMd5

`func (o *HighlightMedia) GetFileMd5() string`

GetFileMd5 returns the FileMd5 field if non-nil, zero value otherwise.

### GetFileMd5Ok

`func (o *HighlightMedia) GetFileMd5Ok() (*string, bool)`

GetFileMd5Ok returns a tuple with the FileMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileMd5

`func (o *HighlightMedia) SetFileMd5(v string)`

SetFileMd5 sets FileMd5 field to given value.

### HasFileMd5

`func (o *HighlightMedia) HasFileMd5() bool`

HasFileMd5 returns a boolean if a field has been set.

### GetFileName

`func (o *HighlightMedia) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *HighlightMedia) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *HighlightMedia) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *HighlightMedia) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFileSize

`func (o *HighlightMedia) GetFileSize() int32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *HighlightMedia) GetFileSizeOk() (*int32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *HighlightMedia) SetFileSize(v int32)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *HighlightMedia) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetHighlightChunks

`func (o *HighlightMedia) GetHighlightChunks() []HighlightChunk`

GetHighlightChunks returns the HighlightChunks field if non-nil, zero value otherwise.

### GetHighlightChunksOk

`func (o *HighlightMedia) GetHighlightChunksOk() (*[]HighlightChunk, bool)`

GetHighlightChunksOk returns a tuple with the HighlightChunks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlightChunks

`func (o *HighlightMedia) SetHighlightChunks(v []HighlightChunk)`

SetHighlightChunks sets HighlightChunks field to given value.

### HasHighlightChunks

`func (o *HighlightMedia) HasHighlightChunks() bool`

HasHighlightChunks returns a boolean if a field has been set.

### GetHighlightClips

`func (o *HighlightMedia) GetHighlightClips() []HighlightClip`

GetHighlightClips returns the HighlightClips field if non-nil, zero value otherwise.

### GetHighlightClipsOk

`func (o *HighlightMedia) GetHighlightClipsOk() (*[]HighlightClip, bool)`

GetHighlightClipsOk returns a tuple with the HighlightClips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlightClips

`func (o *HighlightMedia) SetHighlightClips(v []HighlightClip)`

SetHighlightClips sets HighlightClips field to given value.

### HasHighlightClips

`func (o *HighlightMedia) HasHighlightClips() bool`

HasHighlightClips returns a boolean if a field has been set.

### GetHighlightManifests

`func (o *HighlightMedia) GetHighlightManifests() []HighlightManifest`

GetHighlightManifests returns the HighlightManifests field if non-nil, zero value otherwise.

### GetHighlightManifestsOk

`func (o *HighlightMedia) GetHighlightManifestsOk() (*[]HighlightManifest, bool)`

GetHighlightManifestsOk returns a tuple with the HighlightManifests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlightManifests

`func (o *HighlightMedia) SetHighlightManifests(v []HighlightManifest)`

SetHighlightManifests sets HighlightManifests field to given value.

### HasHighlightManifests

`func (o *HighlightMedia) HasHighlightManifests() bool`

HasHighlightManifests returns a boolean if a field has been set.

### GetId

`func (o *HighlightMedia) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HighlightMedia) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HighlightMedia) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HighlightMedia) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMediaType

`func (o *HighlightMedia) GetMediaType() MediaType`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *HighlightMedia) GetMediaTypeOk() (*MediaType, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *HighlightMedia) SetMediaType(v MediaType)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *HighlightMedia) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### GetMetadata

`func (o *HighlightMedia) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HighlightMedia) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HighlightMedia) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *HighlightMedia) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMimeType

`func (o *HighlightMedia) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *HighlightMedia) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *HighlightMedia) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *HighlightMedia) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetObjUrl

`func (o *HighlightMedia) GetObjUrl() string`

GetObjUrl returns the ObjUrl field if non-nil, zero value otherwise.

### GetObjUrlOk

`func (o *HighlightMedia) GetObjUrlOk() (*string, bool)`

GetObjUrlOk returns a tuple with the ObjUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjUrl

`func (o *HighlightMedia) SetObjUrl(v string)`

SetObjUrl sets ObjUrl field to given value.

### HasObjUrl

`func (o *HighlightMedia) HasObjUrl() bool`

HasObjUrl returns a boolean if a field has been set.

### GetStatus

`func (o *HighlightMedia) GetStatus() HighlightMediaStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HighlightMedia) GetStatusOk() (*HighlightMediaStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HighlightMedia) SetStatus(v HighlightMediaStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HighlightMedia) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotalChunk

`func (o *HighlightMedia) GetTotalChunk() int32`

GetTotalChunk returns the TotalChunk field if non-nil, zero value otherwise.

### GetTotalChunkOk

`func (o *HighlightMedia) GetTotalChunkOk() (*int32, bool)`

GetTotalChunkOk returns a tuple with the TotalChunk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalChunk

`func (o *HighlightMedia) SetTotalChunk(v int32)`

SetTotalChunk sets TotalChunk field to given value.

### HasTotalChunk

`func (o *HighlightMedia) HasTotalChunk() bool`

HasTotalChunk returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *HighlightMedia) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *HighlightMedia) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *HighlightMedia) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *HighlightMedia) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


