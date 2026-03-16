# GetOvertimeMetricsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**MetricsContext**](MetricsContext.md) |  | [optional] 
**Data** | Pointer to [**[]MetricItem**](MetricItem.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetOvertimeMetricsResponse

`func NewGetOvertimeMetricsResponse() *GetOvertimeMetricsResponse`

NewGetOvertimeMetricsResponse instantiates a new GetOvertimeMetricsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOvertimeMetricsResponseWithDefaults

`func NewGetOvertimeMetricsResponseWithDefaults() *GetOvertimeMetricsResponse`

NewGetOvertimeMetricsResponseWithDefaults instantiates a new GetOvertimeMetricsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *GetOvertimeMetricsResponse) GetContext() MetricsContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *GetOvertimeMetricsResponse) GetContextOk() (*MetricsContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *GetOvertimeMetricsResponse) SetContext(v MetricsContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *GetOvertimeMetricsResponse) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetData

`func (o *GetOvertimeMetricsResponse) GetData() []MetricItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetOvertimeMetricsResponse) GetDataOk() (*[]MetricItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetOvertimeMetricsResponse) SetData(v []MetricItem)`

SetData sets Data field to given value.

### HasData

`func (o *GetOvertimeMetricsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotal

`func (o *GetOvertimeMetricsResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetOvertimeMetricsResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetOvertimeMetricsResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetOvertimeMetricsResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


