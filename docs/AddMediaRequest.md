# AddMediaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MediaId** | Pointer to **string** |  | [optional] 
**OptionMediaIds** | Pointer to **[]string** |  | [optional] 
**OptionPlaylists** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAddMediaRequest

`func NewAddMediaRequest() *AddMediaRequest`

NewAddMediaRequest instantiates a new AddMediaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddMediaRequestWithDefaults

`func NewAddMediaRequestWithDefaults() *AddMediaRequest`

NewAddMediaRequestWithDefaults instantiates a new AddMediaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMediaId

`func (o *AddMediaRequest) GetMediaId() string`

GetMediaId returns the MediaId field if non-nil, zero value otherwise.

### GetMediaIdOk

`func (o *AddMediaRequest) GetMediaIdOk() (*string, bool)`

GetMediaIdOk returns a tuple with the MediaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaId

`func (o *AddMediaRequest) SetMediaId(v string)`

SetMediaId sets MediaId field to given value.

### HasMediaId

`func (o *AddMediaRequest) HasMediaId() bool`

HasMediaId returns a boolean if a field has been set.

### GetOptionMediaIds

`func (o *AddMediaRequest) GetOptionMediaIds() []string`

GetOptionMediaIds returns the OptionMediaIds field if non-nil, zero value otherwise.

### GetOptionMediaIdsOk

`func (o *AddMediaRequest) GetOptionMediaIdsOk() (*[]string, bool)`

GetOptionMediaIdsOk returns a tuple with the OptionMediaIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionMediaIds

`func (o *AddMediaRequest) SetOptionMediaIds(v []string)`

SetOptionMediaIds sets OptionMediaIds field to given value.

### HasOptionMediaIds

`func (o *AddMediaRequest) HasOptionMediaIds() bool`

HasOptionMediaIds returns a boolean if a field has been set.

### GetOptionPlaylists

`func (o *AddMediaRequest) GetOptionPlaylists() []string`

GetOptionPlaylists returns the OptionPlaylists field if non-nil, zero value otherwise.

### GetOptionPlaylistsOk

`func (o *AddMediaRequest) GetOptionPlaylistsOk() (*[]string, bool)`

GetOptionPlaylistsOk returns a tuple with the OptionPlaylists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionPlaylists

`func (o *AddMediaRequest) SetOptionPlaylists(v []string)`

SetOptionPlaylists sets OptionPlaylists field to given value.

### HasOptionPlaylists

`func (o *AddMediaRequest) HasOptionPlaylists() bool`

HasOptionPlaylists returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


