# ListApiKeysData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKeys** | Pointer to [**[]ApiKey**](ApiKey.md) |  | [optional] 
**Query** | Pointer to [**ListApiKeysRequest**](ListApiKeysRequest.md) |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewListApiKeysData

`func NewListApiKeysData() *ListApiKeysData`

NewListApiKeysData instantiates a new ListApiKeysData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListApiKeysDataWithDefaults

`func NewListApiKeysDataWithDefaults() *ListApiKeysData`

NewListApiKeysDataWithDefaults instantiates a new ListApiKeysData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKeys

`func (o *ListApiKeysData) GetApiKeys() []ApiKey`

GetApiKeys returns the ApiKeys field if non-nil, zero value otherwise.

### GetApiKeysOk

`func (o *ListApiKeysData) GetApiKeysOk() (*[]ApiKey, bool)`

GetApiKeysOk returns a tuple with the ApiKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeys

`func (o *ListApiKeysData) SetApiKeys(v []ApiKey)`

SetApiKeys sets ApiKeys field to given value.

### HasApiKeys

`func (o *ListApiKeysData) HasApiKeys() bool`

HasApiKeys returns a boolean if a field has been set.

### GetQuery

`func (o *ListApiKeysData) GetQuery() ListApiKeysRequest`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ListApiKeysData) GetQueryOk() (*ListApiKeysRequest, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ListApiKeysData) SetQuery(v ListApiKeysRequest)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *ListApiKeysData) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetTotal

`func (o *ListApiKeysData) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListApiKeysData) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListApiKeysData) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ListApiKeysData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


