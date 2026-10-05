# MetricsContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregation** | Pointer to **string** |  | [optional] 
**Breakdown** | Pointer to **string** |  | [optional] 
**Filter** | Pointer to [**MetricFilter**](MetricFilter.md) |  | [optional] 
**Interval** | Pointer to **string** |  | [optional] 
**Metric** | Pointer to **string** |  | [optional] 
**TimeFrame** | Pointer to [**TimeFrame**](TimeFrame.md) |  | [optional] 

## Methods

### NewMetricsContext

`func NewMetricsContext() *MetricsContext`

NewMetricsContext instantiates a new MetricsContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsContextWithDefaults

`func NewMetricsContextWithDefaults() *MetricsContext`

NewMetricsContextWithDefaults instantiates a new MetricsContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregation

`func (o *MetricsContext) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *MetricsContext) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *MetricsContext) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *MetricsContext) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### GetBreakdown

`func (o *MetricsContext) GetBreakdown() string`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *MetricsContext) GetBreakdownOk() (*string, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *MetricsContext) SetBreakdown(v string)`

SetBreakdown sets Breakdown field to given value.

### HasBreakdown

`func (o *MetricsContext) HasBreakdown() bool`

HasBreakdown returns a boolean if a field has been set.

### GetFilter

`func (o *MetricsContext) GetFilter() MetricFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *MetricsContext) GetFilterOk() (*MetricFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *MetricsContext) SetFilter(v MetricFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *MetricsContext) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetInterval

`func (o *MetricsContext) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *MetricsContext) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *MetricsContext) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *MetricsContext) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetMetric

`func (o *MetricsContext) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *MetricsContext) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *MetricsContext) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *MetricsContext) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetTimeFrame

`func (o *MetricsContext) GetTimeFrame() TimeFrame`

GetTimeFrame returns the TimeFrame field if non-nil, zero value otherwise.

### GetTimeFrameOk

`func (o *MetricsContext) GetTimeFrameOk() (*TimeFrame, bool)`

GetTimeFrameOk returns a tuple with the TimeFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFrame

`func (o *MetricsContext) SetTimeFrame(v TimeFrame)`

SetTimeFrame sets TimeFrame field to given value.

### HasTimeFrame

`func (o *MetricsContext) HasTimeFrame() bool`

HasTimeFrame returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


