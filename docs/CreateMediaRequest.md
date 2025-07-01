# CreateMediaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to [**[]Metadata**](Metadata.md) |  | [optional] 
**Qualities** | Pointer to [**[]QualityConfig**](QualityConfig.md) |  | [optional] 
**SegmentDuration** | Pointer to **int32** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Watermark** | Pointer to [**VideoWatermark**](VideoWatermark.md) |  | [optional] 

## Methods

### NewCreateMediaRequest

`func NewCreateMediaRequest() *CreateMediaRequest`

NewCreateMediaRequest instantiates a new CreateMediaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMediaRequestWithDefaults

`func NewCreateMediaRequestWithDefaults() *CreateMediaRequest`

NewCreateMediaRequestWithDefaults instantiates a new CreateMediaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateMediaRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateMediaRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateMediaRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateMediaRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsPublic

`func (o *CreateMediaRequest) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *CreateMediaRequest) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *CreateMediaRequest) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *CreateMediaRequest) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateMediaRequest) GetMetadata() []Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateMediaRequest) GetMetadataOk() (*[]Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateMediaRequest) SetMetadata(v []Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateMediaRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetQualities

`func (o *CreateMediaRequest) GetQualities() []QualityConfig`

GetQualities returns the Qualities field if non-nil, zero value otherwise.

### GetQualitiesOk

`func (o *CreateMediaRequest) GetQualitiesOk() (*[]QualityConfig, bool)`

GetQualitiesOk returns a tuple with the Qualities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualities

`func (o *CreateMediaRequest) SetQualities(v []QualityConfig)`

SetQualities sets Qualities field to given value.

### HasQualities

`func (o *CreateMediaRequest) HasQualities() bool`

HasQualities returns a boolean if a field has been set.

### GetSegmentDuration

`func (o *CreateMediaRequest) GetSegmentDuration() int32`

GetSegmentDuration returns the SegmentDuration field if non-nil, zero value otherwise.

### GetSegmentDurationOk

`func (o *CreateMediaRequest) GetSegmentDurationOk() (*int32, bool)`

GetSegmentDurationOk returns a tuple with the SegmentDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentDuration

`func (o *CreateMediaRequest) SetSegmentDuration(v int32)`

SetSegmentDuration sets SegmentDuration field to given value.

### HasSegmentDuration

`func (o *CreateMediaRequest) HasSegmentDuration() bool`

HasSegmentDuration returns a boolean if a field has been set.

### GetTags

`func (o *CreateMediaRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateMediaRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateMediaRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateMediaRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTitle

`func (o *CreateMediaRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateMediaRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateMediaRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreateMediaRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetWatermark

`func (o *CreateMediaRequest) GetWatermark() VideoWatermark`

GetWatermark returns the Watermark field if non-nil, zero value otherwise.

### GetWatermarkOk

`func (o *CreateMediaRequest) GetWatermarkOk() (*VideoWatermark, bool)`

GetWatermarkOk returns a tuple with the Watermark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatermark

`func (o *CreateMediaRequest) SetWatermark(v VideoWatermark)`

SetWatermark sets Watermark field to given value.

### HasWatermark

`func (o *CreateMediaRequest) HasWatermark() bool`

HasWatermark returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


