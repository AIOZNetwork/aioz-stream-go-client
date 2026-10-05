# ListPlaylistsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Playlists** | Pointer to [**[]Playlist**](Playlist.md) |  | [optional] 
**Query** | Pointer to [**ListPlaylistsRequest**](ListPlaylistsRequest.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewListPlaylistsData

`func NewListPlaylistsData() *ListPlaylistsData`

NewListPlaylistsData instantiates a new ListPlaylistsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPlaylistsDataWithDefaults

`func NewListPlaylistsDataWithDefaults() *ListPlaylistsData`

NewListPlaylistsDataWithDefaults instantiates a new ListPlaylistsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlaylists

`func (o *ListPlaylistsData) GetPlaylists() []Playlist`

GetPlaylists returns the Playlists field if non-nil, zero value otherwise.

### GetPlaylistsOk

`func (o *ListPlaylistsData) GetPlaylistsOk() (*[]Playlist, bool)`

GetPlaylistsOk returns a tuple with the Playlists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylists

`func (o *ListPlaylistsData) SetPlaylists(v []Playlist)`

SetPlaylists sets Playlists field to given value.

### HasPlaylists

`func (o *ListPlaylistsData) HasPlaylists() bool`

HasPlaylists returns a boolean if a field has been set.

### GetQuery

`func (o *ListPlaylistsData) GetQuery() ListPlaylistsRequest`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ListPlaylistsData) GetQueryOk() (*ListPlaylistsRequest, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ListPlaylistsData) SetQuery(v ListPlaylistsRequest)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *ListPlaylistsData) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetTotal

`func (o *ListPlaylistsData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListPlaylistsData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListPlaylistsData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ListPlaylistsData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


