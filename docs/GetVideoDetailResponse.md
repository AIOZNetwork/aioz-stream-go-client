# GetVideoDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**Media**](Media.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewGetVideoDetailResponse

`func NewGetVideoDetailResponse() *GetVideoDetailResponse`

NewGetVideoDetailResponse instantiates a new GetVideoDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVideoDetailResponseWithDefaults

`func NewGetVideoDetailResponseWithDefaults() *GetVideoDetailResponse`

NewGetVideoDetailResponseWithDefaults instantiates a new GetVideoDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetVideoDetailResponse) GetData() Media`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetVideoDetailResponse) GetDataOk() (*Media, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetVideoDetailResponse) SetData(v Media)`

SetData sets Data field to given value.

### HasData

`func (o *GetVideoDetailResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *GetVideoDetailResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetVideoDetailResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetVideoDetailResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetVideoDetailResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


