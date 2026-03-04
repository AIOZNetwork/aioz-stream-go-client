# GetMediaChaptersData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MediaChapters** | Pointer to [**[]MediaChapter**](MediaChapter.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetMediaChaptersData

`func NewGetMediaChaptersData() *GetMediaChaptersData`

NewGetMediaChaptersData instantiates a new GetMediaChaptersData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMediaChaptersDataWithDefaults

`func NewGetMediaChaptersDataWithDefaults() *GetMediaChaptersData`

NewGetMediaChaptersDataWithDefaults instantiates a new GetMediaChaptersData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMediaChapters

`func (o *GetMediaChaptersData) GetMediaChapters() []MediaChapter`

GetMediaChapters returns the MediaChapters field if non-nil, zero value otherwise.

### GetMediaChaptersOk

`func (o *GetMediaChaptersData) GetMediaChaptersOk() (*[]MediaChapter, bool)`

GetMediaChaptersOk returns a tuple with the MediaChapters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaChapters

`func (o *GetMediaChaptersData) SetMediaChapters(v []MediaChapter)`

SetMediaChapters sets MediaChapters field to given value.

### HasMediaChapters

`func (o *GetMediaChaptersData) HasMediaChapters() bool`

HasMediaChapters returns a boolean if a field has been set.

### GetTotal

`func (o *GetMediaChaptersData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetMediaChaptersData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetMediaChaptersData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetMediaChaptersData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


