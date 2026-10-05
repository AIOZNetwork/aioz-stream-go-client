# ListPlaylistsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ListPlaylistsData**](ListPlaylistsData.md) |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewListPlaylistsResponse

`func NewListPlaylistsResponse() *ListPlaylistsResponse`

NewListPlaylistsResponse instantiates a new ListPlaylistsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPlaylistsResponseWithDefaults

`func NewListPlaylistsResponseWithDefaults() *ListPlaylistsResponse`

NewListPlaylistsResponseWithDefaults instantiates a new ListPlaylistsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListPlaylistsResponse) GetData() ListPlaylistsData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListPlaylistsResponse) GetDataOk() (*ListPlaylistsData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListPlaylistsResponse) SetData(v ListPlaylistsData)`

SetData sets Data field to given value.

### HasData

`func (o *ListPlaylistsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetRequestId

`func (o *ListPlaylistsResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ListPlaylistsResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ListPlaylistsResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ListPlaylistsResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetStatus

`func (o *ListPlaylistsResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListPlaylistsResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListPlaylistsResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListPlaylistsResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


