# StreamUsageSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataTransferred** | Pointer to **float32** |  | [optional] 
**EndedAt** | Pointer to **string** |  | [optional] 
**LiveStreamKeyId** | Pointer to **string** |  | [optional] 
**RenditionBreakdown** | Pointer to [**[]RenditionUsage**](RenditionUsage.md) |  | [optional] 
**StartedAt** | Pointer to **string** |  | [optional] 
**StreamId** | Pointer to **string** |  | [optional] 
**TotalDurationMin** | Pointer to **float32** |  | [optional] 
**TotalDurationMs** | Pointer to **int64** |  | [optional] 
**TotalSegments** | Pointer to **int64** |  | [optional] 
**TotalStorageBytes** | Pointer to **int64** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewStreamUsageSummary

`func NewStreamUsageSummary() *StreamUsageSummary`

NewStreamUsageSummary instantiates a new StreamUsageSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamUsageSummaryWithDefaults

`func NewStreamUsageSummaryWithDefaults() *StreamUsageSummary`

NewStreamUsageSummaryWithDefaults instantiates a new StreamUsageSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataTransferred

`func (o *StreamUsageSummary) GetDataTransferred() float32`

GetDataTransferred returns the DataTransferred field if non-nil, zero value otherwise.

### GetDataTransferredOk

`func (o *StreamUsageSummary) GetDataTransferredOk() (*float32, bool)`

GetDataTransferredOk returns a tuple with the DataTransferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTransferred

`func (o *StreamUsageSummary) SetDataTransferred(v float32)`

SetDataTransferred sets DataTransferred field to given value.

### HasDataTransferred

`func (o *StreamUsageSummary) HasDataTransferred() bool`

HasDataTransferred returns a boolean if a field has been set.

### GetEndedAt

`func (o *StreamUsageSummary) GetEndedAt() string`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *StreamUsageSummary) GetEndedAtOk() (*string, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *StreamUsageSummary) SetEndedAt(v string)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *StreamUsageSummary) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetLiveStreamKeyId

`func (o *StreamUsageSummary) GetLiveStreamKeyId() string`

GetLiveStreamKeyId returns the LiveStreamKeyId field if non-nil, zero value otherwise.

### GetLiveStreamKeyIdOk

`func (o *StreamUsageSummary) GetLiveStreamKeyIdOk() (*string, bool)`

GetLiveStreamKeyIdOk returns a tuple with the LiveStreamKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveStreamKeyId

`func (o *StreamUsageSummary) SetLiveStreamKeyId(v string)`

SetLiveStreamKeyId sets LiveStreamKeyId field to given value.

### HasLiveStreamKeyId

`func (o *StreamUsageSummary) HasLiveStreamKeyId() bool`

HasLiveStreamKeyId returns a boolean if a field has been set.

### GetRenditionBreakdown

`func (o *StreamUsageSummary) GetRenditionBreakdown() []RenditionUsage`

GetRenditionBreakdown returns the RenditionBreakdown field if non-nil, zero value otherwise.

### GetRenditionBreakdownOk

`func (o *StreamUsageSummary) GetRenditionBreakdownOk() (*[]RenditionUsage, bool)`

GetRenditionBreakdownOk returns a tuple with the RenditionBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenditionBreakdown

`func (o *StreamUsageSummary) SetRenditionBreakdown(v []RenditionUsage)`

SetRenditionBreakdown sets RenditionBreakdown field to given value.

### HasRenditionBreakdown

`func (o *StreamUsageSummary) HasRenditionBreakdown() bool`

HasRenditionBreakdown returns a boolean if a field has been set.

### GetStartedAt

`func (o *StreamUsageSummary) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *StreamUsageSummary) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *StreamUsageSummary) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *StreamUsageSummary) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStreamId

`func (o *StreamUsageSummary) GetStreamId() string`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *StreamUsageSummary) GetStreamIdOk() (*string, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *StreamUsageSummary) SetStreamId(v string)`

SetStreamId sets StreamId field to given value.

### HasStreamId

`func (o *StreamUsageSummary) HasStreamId() bool`

HasStreamId returns a boolean if a field has been set.

### GetTotalDurationMin

`func (o *StreamUsageSummary) GetTotalDurationMin() float32`

GetTotalDurationMin returns the TotalDurationMin field if non-nil, zero value otherwise.

### GetTotalDurationMinOk

`func (o *StreamUsageSummary) GetTotalDurationMinOk() (*float32, bool)`

GetTotalDurationMinOk returns a tuple with the TotalDurationMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDurationMin

`func (o *StreamUsageSummary) SetTotalDurationMin(v float32)`

SetTotalDurationMin sets TotalDurationMin field to given value.

### HasTotalDurationMin

`func (o *StreamUsageSummary) HasTotalDurationMin() bool`

HasTotalDurationMin returns a boolean if a field has been set.

### GetTotalDurationMs

`func (o *StreamUsageSummary) GetTotalDurationMs() int64`

GetTotalDurationMs returns the TotalDurationMs field if non-nil, zero value otherwise.

### GetTotalDurationMsOk

`func (o *StreamUsageSummary) GetTotalDurationMsOk() (*int64, bool)`

GetTotalDurationMsOk returns a tuple with the TotalDurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDurationMs

`func (o *StreamUsageSummary) SetTotalDurationMs(v int64)`

SetTotalDurationMs sets TotalDurationMs field to given value.

### HasTotalDurationMs

`func (o *StreamUsageSummary) HasTotalDurationMs() bool`

HasTotalDurationMs returns a boolean if a field has been set.

### GetTotalSegments

`func (o *StreamUsageSummary) GetTotalSegments() int64`

GetTotalSegments returns the TotalSegments field if non-nil, zero value otherwise.

### GetTotalSegmentsOk

`func (o *StreamUsageSummary) GetTotalSegmentsOk() (*int64, bool)`

GetTotalSegmentsOk returns a tuple with the TotalSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSegments

`func (o *StreamUsageSummary) SetTotalSegments(v int64)`

SetTotalSegments sets TotalSegments field to given value.

### HasTotalSegments

`func (o *StreamUsageSummary) HasTotalSegments() bool`

HasTotalSegments returns a boolean if a field has been set.

### GetTotalStorageBytes

`func (o *StreamUsageSummary) GetTotalStorageBytes() int64`

GetTotalStorageBytes returns the TotalStorageBytes field if non-nil, zero value otherwise.

### GetTotalStorageBytesOk

`func (o *StreamUsageSummary) GetTotalStorageBytesOk() (*int64, bool)`

GetTotalStorageBytesOk returns a tuple with the TotalStorageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStorageBytes

`func (o *StreamUsageSummary) SetTotalStorageBytes(v int64)`

SetTotalStorageBytes sets TotalStorageBytes field to given value.

### HasTotalStorageBytes

`func (o *StreamUsageSummary) HasTotalStorageBytes() bool`

HasTotalStorageBytes returns a boolean if a field has been set.

### GetUserId

`func (o *StreamUsageSummary) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *StreamUsageSummary) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *StreamUsageSummary) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *StreamUsageSummary) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


