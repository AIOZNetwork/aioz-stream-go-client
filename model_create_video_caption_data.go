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

// CreateVideoCaptionData struct for CreateVideoCaptionData
type CreateVideoCaptionData struct {
	VideoCaption *VideoCaption `json:"video_caption,omitempty"`
}

// NewCreateVideoCaptionData instantiates a new CreateVideoCaptionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateVideoCaptionData() *CreateVideoCaptionData {
	this := CreateVideoCaptionData{}
	return &this
}

// NewCreateVideoCaptionDataWithDefaults instantiates a new CreateVideoCaptionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateVideoCaptionDataWithDefaults() *CreateVideoCaptionData {
	this := CreateVideoCaptionData{}
	return &this
}

// GetVideoCaption returns the VideoCaption field value if set, zero value otherwise.
func (o *CreateVideoCaptionData) GetVideoCaption() VideoCaption {
	if o == nil || o.VideoCaption == nil {
		var ret VideoCaption
		return ret
	}
	return *o.VideoCaption
}

// GetVideoCaptionOk returns a tuple with the VideoCaption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVideoCaptionData) GetVideoCaptionOk() (*VideoCaption, bool) {
	if o == nil || o.VideoCaption == nil {
		return nil, false
	}
	return o.VideoCaption, true
}

// HasVideoCaption returns a boolean if a field has been set.
func (o *CreateVideoCaptionData) HasVideoCaption() bool {
	if o != nil && o.VideoCaption != nil {
		return true
	}

	return false
}

// SetVideoCaption gets a reference to the given VideoCaption and assigns it to the VideoCaption field.
func (o *CreateVideoCaptionData) SetVideoCaption(v VideoCaption) {
	o.VideoCaption = &v
}

type NullableCreateVideoCaptionData struct {
	value *CreateVideoCaptionData
	isSet bool
}

func (v NullableCreateVideoCaptionData) Get() *CreateVideoCaptionData {
	return v.value
}

func (v *NullableCreateVideoCaptionData) Set(val *CreateVideoCaptionData) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateVideoCaptionData) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateVideoCaptionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateVideoCaptionData(val *CreateVideoCaptionData) *NullableCreateVideoCaptionData {
	return &NullableCreateVideoCaptionData{value: val, isSet: true}
}
