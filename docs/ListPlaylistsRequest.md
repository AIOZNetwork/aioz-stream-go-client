# ListPlaylistsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** |  | [optional] [default to 25]
**Metadata** | Pointer to [**[]Metadata**](Metadata.md) |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**OrderBy** | Pointer to **string** |  | [optional] 
**PlaylistType** | Pointer to **string** |  | [optional] 
**Search** | Pointer to **string** |  | [optional] 
**SortBy** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewListPlaylistsRequest

`func NewListPlaylistsRequest() *ListPlaylistsRequest`

NewListPlaylistsRequest instantiates a new ListPlaylistsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPlaylistsRequestWithDefaults

`func NewListPlaylistsRequestWithDefaults() *ListPlaylistsRequest`

NewListPlaylistsRequestWithDefaults instantiates a new ListPlaylistsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *ListPlaylistsRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListPlaylistsRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListPlaylistsRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListPlaylistsRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetMetadata

`func (o *ListPlaylistsRequest) GetMetadata() []Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ListPlaylistsRequest) GetMetadataOk() (*[]Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ListPlaylistsRequest) SetMetadata(v []Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ListPlaylistsRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOffset

`func (o *ListPlaylistsRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListPlaylistsRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListPlaylistsRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListPlaylistsRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOrderBy

`func (o *ListPlaylistsRequest) GetOrderBy() string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *ListPlaylistsRequest) GetOrderByOk() (*string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *ListPlaylistsRequest) SetOrderBy(v string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *ListPlaylistsRequest) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetPlaylistType

`func (o *ListPlaylistsRequest) GetPlaylistType() string`

GetPlaylistType returns the PlaylistType field if non-nil, zero value otherwise.

### GetPlaylistTypeOk

`func (o *ListPlaylistsRequest) GetPlaylistTypeOk() (*string, bool)`

GetPlaylistTypeOk returns a tuple with the PlaylistType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistType

`func (o *ListPlaylistsRequest) SetPlaylistType(v string)`

SetPlaylistType sets PlaylistType field to given value.

### HasPlaylistType

`func (o *ListPlaylistsRequest) HasPlaylistType() bool`

HasPlaylistType returns a boolean if a field has been set.

### GetSearch

`func (o *ListPlaylistsRequest) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *ListPlaylistsRequest) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *ListPlaylistsRequest) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *ListPlaylistsRequest) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### GetSortBy

`func (o *ListPlaylistsRequest) GetSortBy() string`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *ListPlaylistsRequest) GetSortByOk() (*string, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *ListPlaylistsRequest) SetSortBy(v string)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *ListPlaylistsRequest) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.

### GetTags

`func (o *ListPlaylistsRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ListPlaylistsRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ListPlaylistsRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ListPlaylistsRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


