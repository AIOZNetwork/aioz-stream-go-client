# ListThemesData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlayerThemes** | Pointer to [**[]PlayerTheme**](PlayerTheme.md) |  | [optional] 
**Query** | Pointer to [**ListThemesRequest**](ListThemesRequest.md) |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewListThemesData

`func NewListThemesData() *ListThemesData`

NewListThemesData instantiates a new ListThemesData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListThemesDataWithDefaults

`func NewListThemesDataWithDefaults() *ListThemesData`

NewListThemesDataWithDefaults instantiates a new ListThemesData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlayerThemes

`func (o *ListThemesData) GetPlayerThemes() []PlayerTheme`

GetPlayerThemes returns the PlayerThemes field if non-nil, zero value otherwise.

### GetPlayerThemesOk

`func (o *ListThemesData) GetPlayerThemesOk() (*[]PlayerTheme, bool)`

GetPlayerThemesOk returns a tuple with the PlayerThemes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerThemes

`func (o *ListThemesData) SetPlayerThemes(v []PlayerTheme)`

SetPlayerThemes sets PlayerThemes field to given value.

### HasPlayerThemes

`func (o *ListThemesData) HasPlayerThemes() bool`

HasPlayerThemes returns a boolean if a field has been set.

### GetQuery

`func (o *ListThemesData) GetQuery() ListThemesRequest`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ListThemesData) GetQueryOk() (*ListThemesRequest, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ListThemesData) SetQuery(v ListThemesRequest)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *ListThemesData) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetTotal

`func (o *ListThemesData) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListThemesData) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListThemesData) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ListThemesData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


