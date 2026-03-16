# GetBreakdownMetricsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterBy** | Pointer to [**MetricFilter**](MetricFilter.md) |  | [optional] 
**From** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**OrderBy** | Pointer to **string** |  | [optional] 
**SortBy** | Pointer to **string** |  | [optional] 
**SumOthers** | Pointer to **bool** |  | [optional] 
**To** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetBreakdownMetricsRequest

`func NewGetBreakdownMetricsRequest() *GetBreakdownMetricsRequest`

NewGetBreakdownMetricsRequest instantiates a new GetBreakdownMetricsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBreakdownMetricsRequestWithDefaults

`func NewGetBreakdownMetricsRequestWithDefaults() *GetBreakdownMetricsRequest`

NewGetBreakdownMetricsRequestWithDefaults instantiates a new GetBreakdownMetricsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterBy

`func (o *GetBreakdownMetricsRequest) GetFilterBy() MetricFilter`

GetFilterBy returns the FilterBy field if non-nil, zero value otherwise.

### GetFilterByOk

`func (o *GetBreakdownMetricsRequest) GetFilterByOk() (*MetricFilter, bool)`

GetFilterByOk returns a tuple with the FilterBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterBy

`func (o *GetBreakdownMetricsRequest) SetFilterBy(v MetricFilter)`

SetFilterBy sets FilterBy field to given value.

### HasFilterBy

`func (o *GetBreakdownMetricsRequest) HasFilterBy() bool`

HasFilterBy returns a boolean if a field has been set.

### GetFrom

`func (o *GetBreakdownMetricsRequest) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *GetBreakdownMetricsRequest) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *GetBreakdownMetricsRequest) SetFrom(v int32)`

SetFrom sets From field to given value.

### HasFrom

`func (o *GetBreakdownMetricsRequest) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetLimit

`func (o *GetBreakdownMetricsRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetBreakdownMetricsRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetBreakdownMetricsRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *GetBreakdownMetricsRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *GetBreakdownMetricsRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GetBreakdownMetricsRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GetBreakdownMetricsRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *GetBreakdownMetricsRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOrderBy

`func (o *GetBreakdownMetricsRequest) GetOrderBy() string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *GetBreakdownMetricsRequest) GetOrderByOk() (*string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *GetBreakdownMetricsRequest) SetOrderBy(v string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *GetBreakdownMetricsRequest) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetSortBy

`func (o *GetBreakdownMetricsRequest) GetSortBy() string`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *GetBreakdownMetricsRequest) GetSortByOk() (*string, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *GetBreakdownMetricsRequest) SetSortBy(v string)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *GetBreakdownMetricsRequest) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.

### GetSumOthers

`func (o *GetBreakdownMetricsRequest) GetSumOthers() bool`

GetSumOthers returns the SumOthers field if non-nil, zero value otherwise.

### GetSumOthersOk

`func (o *GetBreakdownMetricsRequest) GetSumOthersOk() (*bool, bool)`

GetSumOthersOk returns a tuple with the SumOthers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumOthers

`func (o *GetBreakdownMetricsRequest) SetSumOthers(v bool)`

SetSumOthers sets SumOthers field to given value.

### HasSumOthers

`func (o *GetBreakdownMetricsRequest) HasSumOthers() bool`

HasSumOthers returns a boolean if a field has been set.

### GetTo

`func (o *GetBreakdownMetricsRequest) GetTo() int32`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *GetBreakdownMetricsRequest) GetToOk() (*int32, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *GetBreakdownMetricsRequest) SetTo(v int32)`

SetTo sets To field to given value.

### HasTo

`func (o *GetBreakdownMetricsRequest) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


