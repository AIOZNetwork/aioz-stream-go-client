# GetAggregatedMetricsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**MetricsContext**](MetricsContext.md) |  | [optional] 
**Data** | Pointer to **float32** |  | [optional] 

## Methods

### NewGetAggregatedMetricsResponse

`func NewGetAggregatedMetricsResponse() *GetAggregatedMetricsResponse`

NewGetAggregatedMetricsResponse instantiates a new GetAggregatedMetricsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAggregatedMetricsResponseWithDefaults

`func NewGetAggregatedMetricsResponseWithDefaults() *GetAggregatedMetricsResponse`

NewGetAggregatedMetricsResponseWithDefaults instantiates a new GetAggregatedMetricsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *GetAggregatedMetricsResponse) GetContext() MetricsContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *GetAggregatedMetricsResponse) GetContextOk() (*MetricsContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *GetAggregatedMetricsResponse) SetContext(v MetricsContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *GetAggregatedMetricsResponse) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetData

`func (o *GetAggregatedMetricsResponse) GetData() float32`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAggregatedMetricsResponse) GetDataOk() (*float32, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAggregatedMetricsResponse) SetData(v float32)`

SetData sets Data field to given value.

### HasData

`func (o *GetAggregatedMetricsResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


