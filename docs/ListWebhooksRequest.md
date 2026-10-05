# ListWebhooksRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncodingFailed** | Pointer to **bool** |  | [optional] 
**EncodingFinished** | Pointer to **bool** |  | [optional] 
**EncodingStarted** | Pointer to **bool** |  | [optional] 
**FileReceived** | Pointer to **bool** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] [default to 25]
**Offset** | Pointer to **int32** |  | [optional] 
**OrderBy** | Pointer to **string** |  | [optional] 
**PartialFinished** | Pointer to **bool** |  | [optional] 
**Search** | Pointer to **string** |  | [optional] 
**SortBy** | Pointer to **string** |  | [optional] 

## Methods

### NewListWebhooksRequest

`func NewListWebhooksRequest() *ListWebhooksRequest`

NewListWebhooksRequest instantiates a new ListWebhooksRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWebhooksRequestWithDefaults

`func NewListWebhooksRequestWithDefaults() *ListWebhooksRequest`

NewListWebhooksRequestWithDefaults instantiates a new ListWebhooksRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncodingFailed

`func (o *ListWebhooksRequest) GetEncodingFailed() bool`

GetEncodingFailed returns the EncodingFailed field if non-nil, zero value otherwise.

### GetEncodingFailedOk

`func (o *ListWebhooksRequest) GetEncodingFailedOk() (*bool, bool)`

GetEncodingFailedOk returns a tuple with the EncodingFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodingFailed

`func (o *ListWebhooksRequest) SetEncodingFailed(v bool)`

SetEncodingFailed sets EncodingFailed field to given value.

### HasEncodingFailed

`func (o *ListWebhooksRequest) HasEncodingFailed() bool`

HasEncodingFailed returns a boolean if a field has been set.

### GetEncodingFinished

`func (o *ListWebhooksRequest) GetEncodingFinished() bool`

GetEncodingFinished returns the EncodingFinished field if non-nil, zero value otherwise.

### GetEncodingFinishedOk

`func (o *ListWebhooksRequest) GetEncodingFinishedOk() (*bool, bool)`

GetEncodingFinishedOk returns a tuple with the EncodingFinished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodingFinished

`func (o *ListWebhooksRequest) SetEncodingFinished(v bool)`

SetEncodingFinished sets EncodingFinished field to given value.

### HasEncodingFinished

`func (o *ListWebhooksRequest) HasEncodingFinished() bool`

HasEncodingFinished returns a boolean if a field has been set.

### GetEncodingStarted

`func (o *ListWebhooksRequest) GetEncodingStarted() bool`

GetEncodingStarted returns the EncodingStarted field if non-nil, zero value otherwise.

### GetEncodingStartedOk

`func (o *ListWebhooksRequest) GetEncodingStartedOk() (*bool, bool)`

GetEncodingStartedOk returns a tuple with the EncodingStarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodingStarted

`func (o *ListWebhooksRequest) SetEncodingStarted(v bool)`

SetEncodingStarted sets EncodingStarted field to given value.

### HasEncodingStarted

`func (o *ListWebhooksRequest) HasEncodingStarted() bool`

HasEncodingStarted returns a boolean if a field has been set.

### GetFileReceived

`func (o *ListWebhooksRequest) GetFileReceived() bool`

GetFileReceived returns the FileReceived field if non-nil, zero value otherwise.

### GetFileReceivedOk

`func (o *ListWebhooksRequest) GetFileReceivedOk() (*bool, bool)`

GetFileReceivedOk returns a tuple with the FileReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileReceived

`func (o *ListWebhooksRequest) SetFileReceived(v bool)`

SetFileReceived sets FileReceived field to given value.

### HasFileReceived

`func (o *ListWebhooksRequest) HasFileReceived() bool`

HasFileReceived returns a boolean if a field has been set.

### GetLimit

`func (o *ListWebhooksRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListWebhooksRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListWebhooksRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListWebhooksRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ListWebhooksRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListWebhooksRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListWebhooksRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListWebhooksRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOrderBy

`func (o *ListWebhooksRequest) GetOrderBy() string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *ListWebhooksRequest) GetOrderByOk() (*string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *ListWebhooksRequest) SetOrderBy(v string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *ListWebhooksRequest) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetPartialFinished

`func (o *ListWebhooksRequest) GetPartialFinished() bool`

GetPartialFinished returns the PartialFinished field if non-nil, zero value otherwise.

### GetPartialFinishedOk

`func (o *ListWebhooksRequest) GetPartialFinishedOk() (*bool, bool)`

GetPartialFinishedOk returns a tuple with the PartialFinished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialFinished

`func (o *ListWebhooksRequest) SetPartialFinished(v bool)`

SetPartialFinished sets PartialFinished field to given value.

### HasPartialFinished

`func (o *ListWebhooksRequest) HasPartialFinished() bool`

HasPartialFinished returns a boolean if a field has been set.

### GetSearch

`func (o *ListWebhooksRequest) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *ListWebhooksRequest) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *ListWebhooksRequest) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *ListWebhooksRequest) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### GetSortBy

`func (o *ListWebhooksRequest) GetSortBy() string`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *ListWebhooksRequest) GetSortByOk() (*string, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *ListWebhooksRequest) SetSortBy(v string)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *ListWebhooksRequest) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


