# ErrorBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** | Reason names what failed - \&quot;media-not-found\&quot;, never \&quot;not-found\&quot; - so a client can tell two failures with the same status apart. | [optional] 
**RequestId** | Pointer to **string** | RequestId is the id this request was logged and traced under, the same value as the X-Request-Id response header. Quoting it is what makes a support conversation actionable. | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewErrorBody

`func NewErrorBody() *ErrorBody`

NewErrorBody instantiates a new ErrorBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorBodyWithDefaults

`func NewErrorBodyWithDefaults() *ErrorBody`

NewErrorBodyWithDefaults instantiates a new ErrorBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ErrorBody) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ErrorBody) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ErrorBody) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *ErrorBody) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *ErrorBody) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorBody) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorBody) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ErrorBody) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetReason

`func (o *ErrorBody) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ErrorBody) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ErrorBody) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ErrorBody) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRequestId

`func (o *ErrorBody) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ErrorBody) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ErrorBody) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ErrorBody) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetStatus

`func (o *ErrorBody) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ErrorBody) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ErrorBody) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ErrorBody) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


