# GetStreamingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**LiveStreamVideosResponse**](LiveStreamVideosResponse.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewGetStreamingsResponse

`func NewGetStreamingsResponse() *GetStreamingsResponse`

NewGetStreamingsResponse instantiates a new GetStreamingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStreamingsResponseWithDefaults

`func NewGetStreamingsResponseWithDefaults() *GetStreamingsResponse`

NewGetStreamingsResponseWithDefaults instantiates a new GetStreamingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetStreamingsResponse) GetData() LiveStreamVideosResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetStreamingsResponse) GetDataOk() (*LiveStreamVideosResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetStreamingsResponse) SetData(v LiveStreamVideosResponse)`

SetData sets Data field to given value.

### HasData

`func (o *GetStreamingsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *GetStreamingsResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetStreamingsResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetStreamingsResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetStreamingsResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


