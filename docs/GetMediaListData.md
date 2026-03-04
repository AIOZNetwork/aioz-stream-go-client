# GetMediaListData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Media** | Pointer to [**[]Media**](Media.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetMediaListData

`func NewGetMediaListData() *GetMediaListData`

NewGetMediaListData instantiates a new GetMediaListData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMediaListDataWithDefaults

`func NewGetMediaListDataWithDefaults() *GetMediaListData`

NewGetMediaListDataWithDefaults instantiates a new GetMediaListData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMedia

`func (o *GetMediaListData) GetMedia() []Media`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *GetMediaListData) GetMediaOk() (*[]Media, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *GetMediaListData) SetMedia(v []Media)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *GetMediaListData) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetTotal

`func (o *GetMediaListData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetMediaListData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetMediaListData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetMediaListData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


