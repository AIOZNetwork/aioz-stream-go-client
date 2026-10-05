# MetricsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterBy** | Pointer to [**MetricFilter**](MetricFilter.md) |  | [optional] 
**From** | Pointer to **int64** |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] [default to 25]
**Offset** | Pointer to **int64** |  | [optional] 
**OrderBy** | Pointer to **string** |  | [optional] 
**SortBy** | Pointer to **string** |  | [optional] 
**SumOthers** | Pointer to **bool** |  | [optional] 
**To** | Pointer to **int64** |  | [optional] 

## Methods

### NewMetricsRequest

`func NewMetricsRequest() *MetricsRequest`

NewMetricsRequest instantiates a new MetricsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsRequestWithDefaults

`func NewMetricsRequestWithDefaults() *MetricsRequest`

NewMetricsRequestWithDefaults instantiates a new MetricsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterBy

`func (o *MetricsRequest) GetFilterBy() MetricFilter`

GetFilterBy returns the FilterBy field if non-nil, zero value otherwise.

### GetFilterByOk

`func (o *MetricsRequest) GetFilterByOk() (*MetricFilter, bool)`

GetFilterByOk returns a tuple with the FilterBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterBy

`func (o *MetricsRequest) SetFilterBy(v MetricFilter)`

SetFilterBy sets FilterBy field to given value.

### HasFilterBy

`func (o *MetricsRequest) HasFilterBy() bool`

HasFilterBy returns a boolean if a field has been set.

### GetFrom

`func (o *MetricsRequest) GetFrom() int64`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *MetricsRequest) GetFromOk() (*int64, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *MetricsRequest) SetFrom(v int64)`

SetFrom sets From field to given value.

### HasFrom

`func (o *MetricsRequest) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetLimit

`func (o *MetricsRequest) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *MetricsRequest) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *MetricsRequest) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *MetricsRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *MetricsRequest) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *MetricsRequest) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *MetricsRequest) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *MetricsRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOrderBy

`func (o *MetricsRequest) GetOrderBy() string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *MetricsRequest) GetOrderByOk() (*string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *MetricsRequest) SetOrderBy(v string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *MetricsRequest) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetSortBy

`func (o *MetricsRequest) GetSortBy() string`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *MetricsRequest) GetSortByOk() (*string, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *MetricsRequest) SetSortBy(v string)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *MetricsRequest) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.

### GetSumOthers

`func (o *MetricsRequest) GetSumOthers() bool`

GetSumOthers returns the SumOthers field if non-nil, zero value otherwise.

### GetSumOthersOk

`func (o *MetricsRequest) GetSumOthersOk() (*bool, bool)`

GetSumOthersOk returns a tuple with the SumOthers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumOthers

`func (o *MetricsRequest) SetSumOthers(v bool)`

SetSumOthers sets SumOthers field to given value.

### HasSumOthers

`func (o *MetricsRequest) HasSumOthers() bool`

HasSumOthers returns a boolean if a field has been set.

### GetTo

`func (o *MetricsRequest) GetTo() int64`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *MetricsRequest) GetToOk() (*int64, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *MetricsRequest) SetTo(v int64)`

SetTo sets To field to given value.

### HasTo

`func (o *MetricsRequest) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


