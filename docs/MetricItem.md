# MetricItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DimensionValue** | Pointer to **string** |  | [optional] 
**EmittedAt** | Pointer to **string** |  | [optional] 
**MetricValue** | Pointer to **float32** |  | [optional] 

## Methods

### NewMetricItem

`func NewMetricItem() *MetricItem`

NewMetricItem instantiates a new MetricItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricItemWithDefaults

`func NewMetricItemWithDefaults() *MetricItem`

NewMetricItemWithDefaults instantiates a new MetricItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDimensionValue

`func (o *MetricItem) GetDimensionValue() string`

GetDimensionValue returns the DimensionValue field if non-nil, zero value otherwise.

### GetDimensionValueOk

`func (o *MetricItem) GetDimensionValueOk() (*string, bool)`

GetDimensionValueOk returns a tuple with the DimensionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionValue

`func (o *MetricItem) SetDimensionValue(v string)`

SetDimensionValue sets DimensionValue field to given value.

### HasDimensionValue

`func (o *MetricItem) HasDimensionValue() bool`

HasDimensionValue returns a boolean if a field has been set.

### GetEmittedAt

`func (o *MetricItem) GetEmittedAt() string`

GetEmittedAt returns the EmittedAt field if non-nil, zero value otherwise.

### GetEmittedAtOk

`func (o *MetricItem) GetEmittedAtOk() (*string, bool)`

GetEmittedAtOk returns a tuple with the EmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmittedAt

`func (o *MetricItem) SetEmittedAt(v string)`

SetEmittedAt sets EmittedAt field to given value.

### HasEmittedAt

`func (o *MetricItem) HasEmittedAt() bool`

HasEmittedAt returns a boolean if a field has been set.

### GetMetricValue

`func (o *MetricItem) GetMetricValue() float32`

GetMetricValue returns the MetricValue field if non-nil, zero value otherwise.

### GetMetricValueOk

`func (o *MetricItem) GetMetricValueOk() (*float32, bool)`

GetMetricValueOk returns a tuple with the MetricValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricValue

`func (o *MetricItem) SetMetricValue(v float32)`

SetMetricValue sets MetricValue field to given value.

### HasMetricValue

`func (o *MetricItem) HasMetricValue() bool`

HasMetricValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


