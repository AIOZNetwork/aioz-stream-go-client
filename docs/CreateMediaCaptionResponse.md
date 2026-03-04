# CreateMediaCaptionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**CreateMediaCaptionData**](CreateMediaCaptionData.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateMediaCaptionResponse

`func NewCreateMediaCaptionResponse() *CreateMediaCaptionResponse`

NewCreateMediaCaptionResponse instantiates a new CreateMediaCaptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMediaCaptionResponseWithDefaults

`func NewCreateMediaCaptionResponseWithDefaults() *CreateMediaCaptionResponse`

NewCreateMediaCaptionResponseWithDefaults instantiates a new CreateMediaCaptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CreateMediaCaptionResponse) GetData() CreateMediaCaptionData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateMediaCaptionResponse) GetDataOk() (*CreateMediaCaptionData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateMediaCaptionResponse) SetData(v CreateMediaCaptionData)`

SetData sets Data field to given value.

### HasData

`func (o *CreateMediaCaptionResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *CreateMediaCaptionResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateMediaCaptionResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateMediaCaptionResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateMediaCaptionResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


