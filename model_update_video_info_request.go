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

// UpdateVideoInfoRequest struct for UpdateVideoInfoRequest
type UpdateVideoInfoRequest struct {
	// Description of the video
	Description *string `json:"description,omitempty"`
	// Video's publish status
	IsPublic *bool `json:"is_public,omitempty"`
	// Video's metadata
	Metadata *[]Metadata `json:"metadata,omitempty"`
	// Video player 's id
	PlayerId *string `json:"player_id,omitempty"`
	// Video's tags
	Tags *[]string `json:"tags,omitempty"`
	// Title of the video
	Title *string `json:"title,omitempty"`
}

// NewUpdateVideoInfoRequest instantiates a new UpdateVideoInfoRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateVideoInfoRequest() *UpdateVideoInfoRequest {
	this := UpdateVideoInfoRequest{}
	return &this
}

// NewUpdateVideoInfoRequestWithDefaults instantiates a new UpdateVideoInfoRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateVideoInfoRequestWithDefaults() *UpdateVideoInfoRequest {
	this := UpdateVideoInfoRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateVideoInfoRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVideoInfoRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateVideoInfoRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateVideoInfoRequest) SetDescription(v string) {
	o.Description = &v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *UpdateVideoInfoRequest) GetIsPublic() bool {
	if o == nil || o.IsPublic == nil {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVideoInfoRequest) GetIsPublicOk() (*bool, bool) {
	if o == nil || o.IsPublic == nil {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *UpdateVideoInfoRequest) HasIsPublic() bool {
	if o != nil && o.IsPublic != nil {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *UpdateVideoInfoRequest) SetIsPublic(v bool) {
	o.IsPublic = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UpdateVideoInfoRequest) GetMetadata() []Metadata {
	if o == nil || o.Metadata == nil {
		var ret []Metadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVideoInfoRequest) GetMetadataOk() (*[]Metadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UpdateVideoInfoRequest) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []Metadata and assigns it to the Metadata field.
func (o *UpdateVideoInfoRequest) SetMetadata(v []Metadata) {
	o.Metadata = &v
}

// GetPlayerId returns the PlayerId field value if set, zero value otherwise.
func (o *UpdateVideoInfoRequest) GetPlayerId() string {
	if o == nil || o.PlayerId == nil {
		var ret string
		return ret
	}
	return *o.PlayerId
}

// GetPlayerIdOk returns a tuple with the PlayerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVideoInfoRequest) GetPlayerIdOk() (*string, bool) {
	if o == nil || o.PlayerId == nil {
		return nil, false
	}
	return o.PlayerId, true
}

// HasPlayerId returns a boolean if a field has been set.
func (o *UpdateVideoInfoRequest) HasPlayerId() bool {
	if o != nil && o.PlayerId != nil {
		return true
	}

	return false
}

// SetPlayerId gets a reference to the given string and assigns it to the PlayerId field.
func (o *UpdateVideoInfoRequest) SetPlayerId(v string) {
	o.PlayerId = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *UpdateVideoInfoRequest) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVideoInfoRequest) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UpdateVideoInfoRequest) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *UpdateVideoInfoRequest) SetTags(v []string) {
	o.Tags = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *UpdateVideoInfoRequest) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVideoInfoRequest) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *UpdateVideoInfoRequest) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *UpdateVideoInfoRequest) SetTitle(v string) {
	o.Title = &v
}

type NullableUpdateVideoInfoRequest struct {
	value *UpdateVideoInfoRequest
	isSet bool
}

func (v NullableUpdateVideoInfoRequest) Get() *UpdateVideoInfoRequest {
	return v.value
}

func (v *NullableUpdateVideoInfoRequest) Set(val *UpdateVideoInfoRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateVideoInfoRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateVideoInfoRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateVideoInfoRequest(val *UpdateVideoInfoRequest) *NullableUpdateVideoInfoRequest {
	return &NullableUpdateVideoInfoRequest{value: val, isSet: true}
}
