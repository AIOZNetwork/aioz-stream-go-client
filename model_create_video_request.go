/*
 * Aioz Stream API
 *
 * Aioz Stream Service
 *
 * API version: 1.0
 * Contact: support@swagger.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aiozstreamsdk

import (
//"encoding/json"
)

// CreateVideoRequest struct for CreateVideoRequest
type CreateVideoRequest struct {
	// Description of the video
	Description *string `json:"description,omitempty"`
	// // Is panoramic video IsPanoramic *bool `json:\"is_panoramic\" form:\"is_panoramic\"` Is public video
	IsPublic *bool `json:"is_public,omitempty"`
	// Metadata of the video (key-value pair, max: 50 items, key max length: 255, value max length: 255)
	Metadata *[]Metadata `json:"metadata,omitempty"`
	// Qualities of the video (default: 1080p, 720p,  360p, allow:2160p, 1440p, 1080p, 720p,  360p, 240p, 144p)
	Qualities *[]string `json:"qualities,omitempty"`
	// Tags of the video (max: 50 items, max length: 255)
	Tags *[]string `json:"tags,omitempty"`
	// Title of the video
	Title *string `json:"title,omitempty"`
	// Video thumbnailConfig
	Watermark *VideoWatermark `json:"watermark,omitempty"`
}

// NewCreateVideoRequest instantiates a new CreateVideoRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateVideoRequest() *CreateVideoRequest {
	this := CreateVideoRequest{}
	return &this
}

// NewCreateVideoRequestWithDefaults instantiates a new CreateVideoRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateVideoRequestWithDefaults() *CreateVideoRequest {
	this := CreateVideoRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateVideoRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVideoRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateVideoRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateVideoRequest) SetDescription(v string) {
	o.Description = &v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *CreateVideoRequest) GetIsPublic() bool {
	if o == nil || o.IsPublic == nil {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVideoRequest) GetIsPublicOk() (*bool, bool) {
	if o == nil || o.IsPublic == nil {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *CreateVideoRequest) HasIsPublic() bool {
	if o != nil && o.IsPublic != nil {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *CreateVideoRequest) SetIsPublic(v bool) {
	o.IsPublic = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CreateVideoRequest) GetMetadata() []Metadata {
	if o == nil || o.Metadata == nil {
		var ret []Metadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVideoRequest) GetMetadataOk() (*[]Metadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CreateVideoRequest) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []Metadata and assigns it to the Metadata field.
func (o *CreateVideoRequest) SetMetadata(v []Metadata) {
	o.Metadata = &v
}

// GetQualities returns the Qualities field value if set, zero value otherwise.
func (o *CreateVideoRequest) GetQualities() []string {
	if o == nil || o.Qualities == nil {
		var ret []string
		return ret
	}
	return *o.Qualities
}

// GetQualitiesOk returns a tuple with the Qualities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVideoRequest) GetQualitiesOk() (*[]string, bool) {
	if o == nil || o.Qualities == nil {
		return nil, false
	}
	return o.Qualities, true
}

// HasQualities returns a boolean if a field has been set.
func (o *CreateVideoRequest) HasQualities() bool {
	if o != nil && o.Qualities != nil {
		return true
	}

	return false
}

// SetQualities gets a reference to the given []string and assigns it to the Qualities field.
func (o *CreateVideoRequest) SetQualities(v []string) {
	o.Qualities = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreateVideoRequest) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVideoRequest) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateVideoRequest) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *CreateVideoRequest) SetTags(v []string) {
	o.Tags = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *CreateVideoRequest) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVideoRequest) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *CreateVideoRequest) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *CreateVideoRequest) SetTitle(v string) {
	o.Title = &v
}

// GetWatermark returns the Watermark field value if set, zero value otherwise.
func (o *CreateVideoRequest) GetWatermark() VideoWatermark {
	if o == nil || o.Watermark == nil {
		var ret VideoWatermark
		return ret
	}
	return *o.Watermark
}

// GetWatermarkOk returns a tuple with the Watermark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVideoRequest) GetWatermarkOk() (*VideoWatermark, bool) {
	if o == nil || o.Watermark == nil {
		return nil, false
	}
	return o.Watermark, true
}

// HasWatermark returns a boolean if a field has been set.
func (o *CreateVideoRequest) HasWatermark() bool {
	if o != nil && o.Watermark != nil {
		return true
	}

	return false
}

// SetWatermark gets a reference to the given VideoWatermark and assigns it to the Watermark field.
func (o *CreateVideoRequest) SetWatermark(v VideoWatermark) {
	o.Watermark = &v
}

type NullableCreateVideoRequest struct {
	value *CreateVideoRequest
	isSet bool
}

func (v NullableCreateVideoRequest) Get() *CreateVideoRequest {
	return v.value
}

func (v *NullableCreateVideoRequest) Set(val *CreateVideoRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateVideoRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateVideoRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateVideoRequest(val *CreateVideoRequest) *NullableCreateVideoRequest {
	return &NullableCreateVideoRequest{value: val, isSet: true}
}
