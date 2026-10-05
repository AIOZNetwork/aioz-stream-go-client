# MetricsPageData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**MetricsContext**](MetricsContext.md) |  | [optional] 
**Data** | Pointer to [**[]MetricItem**](MetricItem.md) |  | [optional] 
**Query** | Pointer to [**MetricsRequest**](MetricsRequest.md) |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewMetricsPageData

`func NewMetricsPageData() *MetricsPageData`

NewMetricsPageData instantiates a new MetricsPageData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsPageDataWithDefaults

`func NewMetricsPageDataWithDefaults() *MetricsPageData`

NewMetricsPageDataWithDefaults instantiates a new MetricsPageData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *MetricsPageData) GetContext() MetricsContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *MetricsPageData) GetContextOk() (*MetricsContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *MetricsPageData) SetContext(v MetricsContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *MetricsPageData) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetData

`func (o *MetricsPageData) GetData() []MetricItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MetricsPageData) GetDataOk() (*[]MetricItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MetricsPageData) SetData(v []MetricItem)`

SetData sets Data field to given value.

### HasData

`func (o *MetricsPageData) HasData() bool`

HasData returns a boolean if a field has been set.

### GetQuery

`func (o *MetricsPageData) GetQuery() MetricsRequest`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *MetricsPageData) GetQueryOk() (*MetricsRequest, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *MetricsPageData) SetQuery(v MetricsRequest)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *MetricsPageData) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetTotal

`func (o *MetricsPageData) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MetricsPageData) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MetricsPageData) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *MetricsPageData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


