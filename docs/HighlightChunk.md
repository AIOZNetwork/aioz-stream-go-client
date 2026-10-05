# HighlightChunk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**DeletedAt** | Pointer to [**DeletedAt**](DeletedAt.md) |  | [optional] 
**FileMd5** | Pointer to **string** |  | [optional] 
**FileSize** | Pointer to **int32** | FileSize is the size of the uploaded chunk in bytes. | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**MediaId** | Pointer to **string** |  | [optional] 
**ObjUrl** | Pointer to **string** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to [**HighlightChunkStatus**](HighlightChunkStatus.md) |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewHighlightChunk

`func NewHighlightChunk() *HighlightChunk`

NewHighlightChunk instantiates a new HighlightChunk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHighlightChunkWithDefaults

`func NewHighlightChunkWithDefaults() *HighlightChunk`

NewHighlightChunkWithDefaults instantiates a new HighlightChunk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *HighlightChunk) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HighlightChunk) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HighlightChunk) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *HighlightChunk) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *HighlightChunk) GetDeletedAt() DeletedAt`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *HighlightChunk) GetDeletedAtOk() (*DeletedAt, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *HighlightChunk) SetDeletedAt(v DeletedAt)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *HighlightChunk) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetFileMd5

`func (o *HighlightChunk) GetFileMd5() string`

GetFileMd5 returns the FileMd5 field if non-nil, zero value otherwise.

### GetFileMd5Ok

`func (o *HighlightChunk) GetFileMd5Ok() (*string, bool)`

GetFileMd5Ok returns a tuple with the FileMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileMd5

`func (o *HighlightChunk) SetFileMd5(v string)`

SetFileMd5 sets FileMd5 field to given value.

### HasFileMd5

`func (o *HighlightChunk) HasFileMd5() bool`

HasFileMd5 returns a boolean if a field has been set.

### GetFileSize

`func (o *HighlightChunk) GetFileSize() int32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *HighlightChunk) GetFileSizeOk() (*int32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *HighlightChunk) SetFileSize(v int32)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *HighlightChunk) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetId

`func (o *HighlightChunk) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HighlightChunk) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HighlightChunk) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HighlightChunk) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMediaId

`func (o *HighlightChunk) GetMediaId() string`

GetMediaId returns the MediaId field if non-nil, zero value otherwise.

### GetMediaIdOk

`func (o *HighlightChunk) GetMediaIdOk() (*string, bool)`

GetMediaIdOk returns a tuple with the MediaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaId

`func (o *HighlightChunk) SetMediaId(v string)`

SetMediaId sets MediaId field to given value.

### HasMediaId

`func (o *HighlightChunk) HasMediaId() bool`

HasMediaId returns a boolean if a field has been set.

### GetObjUrl

`func (o *HighlightChunk) GetObjUrl() string`

GetObjUrl returns the ObjUrl field if non-nil, zero value otherwise.

### GetObjUrlOk

`func (o *HighlightChunk) GetObjUrlOk() (*string, bool)`

GetObjUrlOk returns a tuple with the ObjUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjUrl

`func (o *HighlightChunk) SetObjUrl(v string)`

SetObjUrl sets ObjUrl field to given value.

### HasObjUrl

`func (o *HighlightChunk) HasObjUrl() bool`

HasObjUrl returns a boolean if a field has been set.

### GetOffset

`func (o *HighlightChunk) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *HighlightChunk) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *HighlightChunk) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *HighlightChunk) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetStatus

`func (o *HighlightChunk) GetStatus() HighlightChunkStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HighlightChunk) GetStatusOk() (*HighlightChunkStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HighlightChunk) SetStatus(v HighlightChunkStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HighlightChunk) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *HighlightChunk) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *HighlightChunk) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *HighlightChunk) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *HighlightChunk) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


