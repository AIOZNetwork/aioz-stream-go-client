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

// CreatePlaylistRequest struct for CreatePlaylistRequest
type CreatePlaylistRequest struct {
	Metadata *[]Metadata `json:"metadata,omitempty"`
	Name     *string     `json:"name,omitempty"`
	Tags     *[]string   `json:"tags,omitempty"`
}

// NewCreatePlaylistRequest instantiates a new CreatePlaylistRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePlaylistRequest() *CreatePlaylistRequest {
	this := CreatePlaylistRequest{}
	return &this
}

// NewCreatePlaylistRequestWithDefaults instantiates a new CreatePlaylistRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePlaylistRequestWithDefaults() *CreatePlaylistRequest {
	this := CreatePlaylistRequest{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CreatePlaylistRequest) GetMetadata() []Metadata {
	if o == nil || o.Metadata == nil {
		var ret []Metadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePlaylistRequest) GetMetadataOk() (*[]Metadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CreatePlaylistRequest) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []Metadata and assigns it to the Metadata field.
func (o *CreatePlaylistRequest) SetMetadata(v []Metadata) {
	o.Metadata = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreatePlaylistRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePlaylistRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreatePlaylistRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreatePlaylistRequest) SetName(v string) {
	o.Name = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreatePlaylistRequest) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePlaylistRequest) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreatePlaylistRequest) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *CreatePlaylistRequest) SetTags(v []string) {
	o.Tags = &v
}

type NullableCreatePlaylistRequest struct {
	value *CreatePlaylistRequest
	isSet bool
}

func (v NullableCreatePlaylistRequest) Get() *CreatePlaylistRequest {
	return v.value
}

func (v *NullableCreatePlaylistRequest) Set(val *CreatePlaylistRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePlaylistRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePlaylistRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePlaylistRequest(val *CreatePlaylistRequest) *NullableCreatePlaylistRequest {
	return &NullableCreatePlaylistRequest{value: val, isSet: true}
}
