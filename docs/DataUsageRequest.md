# DataUsageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **int64** |  | [optional] 
**Interval** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] [default to 25]
**Offset** | Pointer to **int64** |  | [optional] 
**To** | Pointer to **int64** |  | [optional] 

## Methods

### NewDataUsageRequest

`func NewDataUsageRequest() *DataUsageRequest`

NewDataUsageRequest instantiates a new DataUsageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataUsageRequestWithDefaults

`func NewDataUsageRequestWithDefaults() *DataUsageRequest`

NewDataUsageRequestWithDefaults instantiates a new DataUsageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *DataUsageRequest) GetFrom() int64`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *DataUsageRequest) GetFromOk() (*int64, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *DataUsageRequest) SetFrom(v int64)`

SetFrom sets From field to given value.

### HasFrom

`func (o *DataUsageRequest) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetInterval

`func (o *DataUsageRequest) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *DataUsageRequest) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *DataUsageRequest) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *DataUsageRequest) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetLimit

`func (o *DataUsageRequest) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *DataUsageRequest) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *DataUsageRequest) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *DataUsageRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *DataUsageRequest) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *DataUsageRequest) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *DataUsageRequest) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *DataUsageRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetTo

`func (o *DataUsageRequest) GetTo() int64`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *DataUsageRequest) GetToOk() (*int64, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *DataUsageRequest) SetTo(v int64)`

SetTo sets To field to given value.

### HasTo

`func (o *DataUsageRequest) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


