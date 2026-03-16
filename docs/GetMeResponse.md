# GetMeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**GetUserData**](GetUserData.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewGetMeResponse

`func NewGetMeResponse() *GetMeResponse`

NewGetMeResponse instantiates a new GetMeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMeResponseWithDefaults

`func NewGetMeResponseWithDefaults() *GetMeResponse`

NewGetMeResponseWithDefaults instantiates a new GetMeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetMeResponse) GetData() GetUserData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetMeResponse) GetDataOk() (*GetUserData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetMeResponse) SetData(v GetUserData)`

SetData sets Data field to given value.

### HasData

`func (o *GetMeResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *GetMeResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetMeResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetMeResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetMeResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


