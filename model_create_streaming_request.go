/*
 * W3STREAM API
 *
 * W3STREAM Service
 *
 * API version: 1.0
 * Contact: support@swagger.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package w3streamsdk

import (
//"encoding/json"
)

// CreateStreamingRequest struct for CreateStreamingRequest
type CreateStreamingRequest struct {
	// Qualities of the video (default: 1080p, 720p,  360p, allow:2160p, 1440p, 1080p, 720p,  360p, 240p, 144p)
	Qualities *[]string `json:"qualities,omitempty"`
	Save      *bool     `json:"save,omitempty"`
	Title     *string   `json:"title,omitempty"`
}

// NewCreateStreamingRequest instantiates a new CreateStreamingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateStreamingRequest() *CreateStreamingRequest {
	this := CreateStreamingRequest{}
	return &this
}

// NewCreateStreamingRequestWithDefaults instantiates a new CreateStreamingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateStreamingRequestWithDefaults() *CreateStreamingRequest {
	this := CreateStreamingRequest{}
	return &this
}

// GetQualities returns the Qualities field value if set, zero value otherwise.
func (o *CreateStreamingRequest) GetQualities() []string {
	if o == nil || o.Qualities == nil {
		var ret []string
		return ret
	}
	return *o.Qualities
}

// GetQualitiesOk returns a tuple with the Qualities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStreamingRequest) GetQualitiesOk() (*[]string, bool) {
	if o == nil || o.Qualities == nil {
		return nil, false
	}
	return o.Qualities, true
}

// HasQualities returns a boolean if a field has been set.
func (o *CreateStreamingRequest) HasQualities() bool {
	if o != nil && o.Qualities != nil {
		return true
	}

	return false
}

// SetQualities gets a reference to the given []string and assigns it to the Qualities field.
func (o *CreateStreamingRequest) SetQualities(v []string) {
	o.Qualities = &v
}

// GetSave returns the Save field value if set, zero value otherwise.
func (o *CreateStreamingRequest) GetSave() bool {
	if o == nil || o.Save == nil {
		var ret bool
		return ret
	}
	return *o.Save
}

// GetSaveOk returns a tuple with the Save field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStreamingRequest) GetSaveOk() (*bool, bool) {
	if o == nil || o.Save == nil {
		return nil, false
	}
	return o.Save, true
}

// HasSave returns a boolean if a field has been set.
func (o *CreateStreamingRequest) HasSave() bool {
	if o != nil && o.Save != nil {
		return true
	}

	return false
}

// SetSave gets a reference to the given bool and assigns it to the Save field.
func (o *CreateStreamingRequest) SetSave(v bool) {
	o.Save = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *CreateStreamingRequest) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStreamingRequest) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *CreateStreamingRequest) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *CreateStreamingRequest) SetTitle(v string) {
	o.Title = &v
}

type NullableCreateStreamingRequest struct {
	value *CreateStreamingRequest
	isSet bool
}

func (v NullableCreateStreamingRequest) Get() *CreateStreamingRequest {
	return v.value
}

func (v *NullableCreateStreamingRequest) Set(val *CreateStreamingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateStreamingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateStreamingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateStreamingRequest(val *CreateStreamingRequest) *NullableCreateStreamingRequest {
	return &NullableCreateStreamingRequest{value: val, isSet: true}
}
