# ListApiKeysRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** |  | [optional] [default to 25]
**Offset** | Pointer to **int32** |  | [optional] 
**OrderBy** | Pointer to **string** |  | [optional] 
**Search** | Pointer to **string** |  | [optional] 
**SortBy** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewListApiKeysRequest

`func NewListApiKeysRequest() *ListApiKeysRequest`

NewListApiKeysRequest instantiates a new ListApiKeysRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListApiKeysRequestWithDefaults

`func NewListApiKeysRequestWithDefaults() *ListApiKeysRequest`

NewListApiKeysRequestWithDefaults instantiates a new ListApiKeysRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *ListApiKeysRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListApiKeysRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListApiKeysRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListApiKeysRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ListApiKeysRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListApiKeysRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListApiKeysRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListApiKeysRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOrderBy

`func (o *ListApiKeysRequest) GetOrderBy() string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *ListApiKeysRequest) GetOrderByOk() (*string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *ListApiKeysRequest) SetOrderBy(v string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *ListApiKeysRequest) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetSearch

`func (o *ListApiKeysRequest) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *ListApiKeysRequest) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *ListApiKeysRequest) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *ListApiKeysRequest) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### GetSortBy

`func (o *ListApiKeysRequest) GetSortBy() string`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *ListApiKeysRequest) GetSortByOk() (*string, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *ListApiKeysRequest) SetSortBy(v string)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *ListApiKeysRequest) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.

### GetType

`func (o *ListApiKeysRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListApiKeysRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListApiKeysRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListApiKeysRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


