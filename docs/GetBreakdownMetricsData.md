# GetBreakdownMetricsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**MetricsContext**](MetricsContext.md) |  | [optional] 
**Data** | Pointer to [**[]MetricItem**](MetricItem.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetBreakdownMetricsData

`func NewGetBreakdownMetricsData() *GetBreakdownMetricsData`

NewGetBreakdownMetricsData instantiates a new GetBreakdownMetricsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBreakdownMetricsDataWithDefaults

`func NewGetBreakdownMetricsDataWithDefaults() *GetBreakdownMetricsData`

NewGetBreakdownMetricsDataWithDefaults instantiates a new GetBreakdownMetricsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *GetBreakdownMetricsData) GetContext() MetricsContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *GetBreakdownMetricsData) GetContextOk() (*MetricsContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *GetBreakdownMetricsData) SetContext(v MetricsContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *GetBreakdownMetricsData) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetData

`func (o *GetBreakdownMetricsData) GetData() []MetricItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetBreakdownMetricsData) GetDataOk() (*[]MetricItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetBreakdownMetricsData) SetData(v []MetricItem)`

SetData sets Data field to given value.

### HasData

`func (o *GetBreakdownMetricsData) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotal

`func (o *GetBreakdownMetricsData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetBreakdownMetricsData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetBreakdownMetricsData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetBreakdownMetricsData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


