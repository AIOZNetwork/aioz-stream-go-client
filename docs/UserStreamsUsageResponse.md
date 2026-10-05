# UserStreamsUsageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Streams** | Pointer to [**[]StreamUsageSummary**](StreamUsageSummary.md) |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewUserStreamsUsageResponse

`func NewUserStreamsUsageResponse() *UserStreamsUsageResponse`

NewUserStreamsUsageResponse instantiates a new UserStreamsUsageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserStreamsUsageResponseWithDefaults

`func NewUserStreamsUsageResponseWithDefaults() *UserStreamsUsageResponse`

NewUserStreamsUsageResponseWithDefaults instantiates a new UserStreamsUsageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreams

`func (o *UserStreamsUsageResponse) GetStreams() []StreamUsageSummary`

GetStreams returns the Streams field if non-nil, zero value otherwise.

### GetStreamsOk

`func (o *UserStreamsUsageResponse) GetStreamsOk() (*[]StreamUsageSummary, bool)`

GetStreamsOk returns a tuple with the Streams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreams

`func (o *UserStreamsUsageResponse) SetStreams(v []StreamUsageSummary)`

SetStreams sets Streams field to given value.

### HasStreams

`func (o *UserStreamsUsageResponse) HasStreams() bool`

HasStreams returns a boolean if a field has been set.

### GetTotal

`func (o *UserStreamsUsageResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *UserStreamsUsageResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *UserStreamsUsageResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *UserStreamsUsageResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


