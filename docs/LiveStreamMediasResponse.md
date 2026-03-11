# LiveStreamMediasResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Media** | Pointer to [**[]LiveStreamMediaResponse**](LiveStreamMediaResponse.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewLiveStreamMediasResponse

`func NewLiveStreamMediasResponse() *LiveStreamMediasResponse`

NewLiveStreamMediasResponse instantiates a new LiveStreamMediasResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveStreamMediasResponseWithDefaults

`func NewLiveStreamMediasResponseWithDefaults() *LiveStreamMediasResponse`

NewLiveStreamMediasResponseWithDefaults instantiates a new LiveStreamMediasResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMedia

`func (o *LiveStreamMediasResponse) GetMedia() []LiveStreamMediaResponse`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *LiveStreamMediasResponse) GetMediaOk() (*[]LiveStreamMediaResponse, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *LiveStreamMediasResponse) SetMedia(v []LiveStreamMediaResponse)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *LiveStreamMediasResponse) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetTotal

`func (o *LiveStreamMediasResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *LiveStreamMediasResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *LiveStreamMediasResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *LiveStreamMediasResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


