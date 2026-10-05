# DataUsageData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]DataUsage**](DataUsage.md) |  | [optional] 
**Query** | Pointer to [**DataUsageRequest**](DataUsageRequest.md) |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewDataUsageData

`func NewDataUsageData() *DataUsageData`

NewDataUsageData instantiates a new DataUsageData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataUsageDataWithDefaults

`func NewDataUsageDataWithDefaults() *DataUsageData`

NewDataUsageDataWithDefaults instantiates a new DataUsageData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DataUsageData) GetData() []DataUsage`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DataUsageData) GetDataOk() (*[]DataUsage, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DataUsageData) SetData(v []DataUsage)`

SetData sets Data field to given value.

### HasData

`func (o *DataUsageData) HasData() bool`

HasData returns a boolean if a field has been set.

### GetQuery

`func (o *DataUsageData) GetQuery() DataUsageRequest`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *DataUsageData) GetQueryOk() (*DataUsageRequest, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *DataUsageData) SetQuery(v DataUsageRequest)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *DataUsageData) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetTotal

`func (o *DataUsageData) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *DataUsageData) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *DataUsageData) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *DataUsageData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


