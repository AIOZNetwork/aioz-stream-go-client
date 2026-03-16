# GetBreakdownMetricsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**MetricsContext**](MetricsContext.md) |  | [optional] 
**Data** | Pointer to [**[]MetricItem**](MetricItem.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetBreakdownMetricsResponse

`func NewGetBreakdownMetricsResponse() *GetBreakdownMetricsResponse`

NewGetBreakdownMetricsResponse instantiates a new GetBreakdownMetricsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBreakdownMetricsResponseWithDefaults

`func NewGetBreakdownMetricsResponseWithDefaults() *GetBreakdownMetricsResponse`

NewGetBreakdownMetricsResponseWithDefaults instantiates a new GetBreakdownMetricsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *GetBreakdownMetricsResponse) GetContext() MetricsContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *GetBreakdownMetricsResponse) GetContextOk() (*MetricsContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *GetBreakdownMetricsResponse) SetContext(v MetricsContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *GetBreakdownMetricsResponse) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetData

`func (o *GetBreakdownMetricsResponse) GetData() []MetricItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetBreakdownMetricsResponse) GetDataOk() (*[]MetricItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetBreakdownMetricsResponse) SetData(v []MetricItem)`

SetData sets Data field to given value.

### HasData

`func (o *GetBreakdownMetricsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotal

`func (o *GetBreakdownMetricsResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetBreakdownMetricsResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetBreakdownMetricsResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetBreakdownMetricsResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


