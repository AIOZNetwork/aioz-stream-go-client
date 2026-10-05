# HighlightClip

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**DeletedAt** | Pointer to [**DeletedAt**](DeletedAt.md) |  | [optional] 
**HookTitle** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**MediaId** | Pointer to **string** |  | [optional] 
**Reasoning** | Pointer to **string** |  | [optional] 
**RelevanceScore** | Pointer to **int32** |  | [optional] 
**SourceInMs** | Pointer to **int32** |  | [optional] 
**SourceOutMs** | Pointer to **int32** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Virality** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewHighlightClip

`func NewHighlightClip() *HighlightClip`

NewHighlightClip instantiates a new HighlightClip object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHighlightClipWithDefaults

`func NewHighlightClipWithDefaults() *HighlightClip`

NewHighlightClipWithDefaults instantiates a new HighlightClip object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *HighlightClip) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HighlightClip) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HighlightClip) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *HighlightClip) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *HighlightClip) GetDeletedAt() DeletedAt`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *HighlightClip) GetDeletedAtOk() (*DeletedAt, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *HighlightClip) SetDeletedAt(v DeletedAt)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *HighlightClip) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetHookTitle

`func (o *HighlightClip) GetHookTitle() string`

GetHookTitle returns the HookTitle field if non-nil, zero value otherwise.

### GetHookTitleOk

`func (o *HighlightClip) GetHookTitleOk() (*string, bool)`

GetHookTitleOk returns a tuple with the HookTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHookTitle

`func (o *HighlightClip) SetHookTitle(v string)`

SetHookTitle sets HookTitle field to given value.

### HasHookTitle

`func (o *HighlightClip) HasHookTitle() bool`

HasHookTitle returns a boolean if a field has been set.

### GetId

`func (o *HighlightClip) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HighlightClip) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HighlightClip) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HighlightClip) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMediaId

`func (o *HighlightClip) GetMediaId() string`

GetMediaId returns the MediaId field if non-nil, zero value otherwise.

### GetMediaIdOk

`func (o *HighlightClip) GetMediaIdOk() (*string, bool)`

GetMediaIdOk returns a tuple with the MediaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaId

`func (o *HighlightClip) SetMediaId(v string)`

SetMediaId sets MediaId field to given value.

### HasMediaId

`func (o *HighlightClip) HasMediaId() bool`

HasMediaId returns a boolean if a field has been set.

### GetReasoning

`func (o *HighlightClip) GetReasoning() string`

GetReasoning returns the Reasoning field if non-nil, zero value otherwise.

### GetReasoningOk

`func (o *HighlightClip) GetReasoningOk() (*string, bool)`

GetReasoningOk returns a tuple with the Reasoning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoning

`func (o *HighlightClip) SetReasoning(v string)`

SetReasoning sets Reasoning field to given value.

### HasReasoning

`func (o *HighlightClip) HasReasoning() bool`

HasReasoning returns a boolean if a field has been set.

### GetRelevanceScore

`func (o *HighlightClip) GetRelevanceScore() int32`

GetRelevanceScore returns the RelevanceScore field if non-nil, zero value otherwise.

### GetRelevanceScoreOk

`func (o *HighlightClip) GetRelevanceScoreOk() (*int32, bool)`

GetRelevanceScoreOk returns a tuple with the RelevanceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelevanceScore

`func (o *HighlightClip) SetRelevanceScore(v int32)`

SetRelevanceScore sets RelevanceScore field to given value.

### HasRelevanceScore

`func (o *HighlightClip) HasRelevanceScore() bool`

HasRelevanceScore returns a boolean if a field has been set.

### GetSourceInMs

`func (o *HighlightClip) GetSourceInMs() int32`

GetSourceInMs returns the SourceInMs field if non-nil, zero value otherwise.

### GetSourceInMsOk

`func (o *HighlightClip) GetSourceInMsOk() (*int32, bool)`

GetSourceInMsOk returns a tuple with the SourceInMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInMs

`func (o *HighlightClip) SetSourceInMs(v int32)`

SetSourceInMs sets SourceInMs field to given value.

### HasSourceInMs

`func (o *HighlightClip) HasSourceInMs() bool`

HasSourceInMs returns a boolean if a field has been set.

### GetSourceOutMs

`func (o *HighlightClip) GetSourceOutMs() int32`

GetSourceOutMs returns the SourceOutMs field if non-nil, zero value otherwise.

### GetSourceOutMsOk

`func (o *HighlightClip) GetSourceOutMsOk() (*int32, bool)`

GetSourceOutMsOk returns a tuple with the SourceOutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOutMs

`func (o *HighlightClip) SetSourceOutMs(v int32)`

SetSourceOutMs sets SourceOutMs field to given value.

### HasSourceOutMs

`func (o *HighlightClip) HasSourceOutMs() bool`

HasSourceOutMs returns a boolean if a field has been set.

### GetText

`func (o *HighlightClip) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *HighlightClip) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *HighlightClip) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *HighlightClip) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTitle

`func (o *HighlightClip) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *HighlightClip) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *HighlightClip) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *HighlightClip) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *HighlightClip) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *HighlightClip) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *HighlightClip) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *HighlightClip) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVirality

`func (o *HighlightClip) GetVirality() map[string]interface{}`

GetVirality returns the Virality field if non-nil, zero value otherwise.

### GetViralityOk

`func (o *HighlightClip) GetViralityOk() (*map[string]interface{}, bool)`

GetViralityOk returns a tuple with the Virality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirality

`func (o *HighlightClip) SetVirality(v map[string]interface{})`

SetVirality sets Virality field to given value.

### HasVirality

`func (o *HighlightClip) HasVirality() bool`

HasVirality returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


