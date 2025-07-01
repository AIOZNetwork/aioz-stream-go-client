# GetVideoListData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**Videos** | Pointer to [**[]Media**](Media.md) |  | [optional] 

## Methods

### NewGetVideoListData

`func NewGetVideoListData() *GetVideoListData`

NewGetVideoListData instantiates a new GetVideoListData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVideoListDataWithDefaults

`func NewGetVideoListDataWithDefaults() *GetVideoListData`

NewGetVideoListDataWithDefaults instantiates a new GetVideoListData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *GetVideoListData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetVideoListData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetVideoListData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetVideoListData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetVideos

`func (o *GetVideoListData) GetVideos() []Media`

GetVideos returns the Videos field if non-nil, zero value otherwise.

### GetVideosOk

`func (o *GetVideoListData) GetVideosOk() (*[]Media, bool)`

GetVideosOk returns a tuple with the Videos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideos

`func (o *GetVideoListData) SetVideos(v []Media)`

SetVideos sets Videos field to given value.

### HasVideos

`func (o *GetVideoListData) HasVideos() bool`

HasVideos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


