# ListWebhooksData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | Pointer to [**ListWebhooksRequest**](ListWebhooksRequest.md) |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 
**Webhooks** | Pointer to [**[]Webhook**](Webhook.md) |  | [optional] 

## Methods

### NewListWebhooksData

`func NewListWebhooksData() *ListWebhooksData`

NewListWebhooksData instantiates a new ListWebhooksData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWebhooksDataWithDefaults

`func NewListWebhooksDataWithDefaults() *ListWebhooksData`

NewListWebhooksDataWithDefaults instantiates a new ListWebhooksData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *ListWebhooksData) GetQuery() ListWebhooksRequest`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ListWebhooksData) GetQueryOk() (*ListWebhooksRequest, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ListWebhooksData) SetQuery(v ListWebhooksRequest)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *ListWebhooksData) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetTotal

`func (o *ListWebhooksData) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListWebhooksData) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListWebhooksData) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ListWebhooksData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetWebhooks

`func (o *ListWebhooksData) GetWebhooks() []Webhook`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *ListWebhooksData) GetWebhooksOk() (*[]Webhook, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *ListWebhooksData) SetWebhooks(v []Webhook)`

SetWebhooks sets Webhooks field to given value.

### HasWebhooks

`func (o *ListWebhooksData) HasWebhooks() bool`

HasWebhooks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


