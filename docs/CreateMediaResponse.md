# CreateMediaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**Media**](Media.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateMediaResponse

`func NewCreateMediaResponse() *CreateMediaResponse`

NewCreateMediaResponse instantiates a new CreateMediaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMediaResponseWithDefaults

`func NewCreateMediaResponseWithDefaults() *CreateMediaResponse`

NewCreateMediaResponseWithDefaults instantiates a new CreateMediaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CreateMediaResponse) GetData() Media`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateMediaResponse) GetDataOk() (*Media, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateMediaResponse) SetData(v Media)`

SetData sets Data field to given value.

### HasData

`func (o *CreateMediaResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *CreateMediaResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateMediaResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateMediaResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateMediaResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


