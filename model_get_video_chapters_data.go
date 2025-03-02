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

// GetVideoChaptersData struct for GetVideoChaptersData
type GetVideoChaptersData struct {
	Total         *int32          `json:"total,omitempty"`
	VideoChapters *[]VideoChapter `json:"video_chapters,omitempty"`
}

// NewGetVideoChaptersData instantiates a new GetVideoChaptersData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetVideoChaptersData() *GetVideoChaptersData {
	this := GetVideoChaptersData{}
	return &this
}

// NewGetVideoChaptersDataWithDefaults instantiates a new GetVideoChaptersData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetVideoChaptersDataWithDefaults() *GetVideoChaptersData {
	this := GetVideoChaptersData{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetVideoChaptersData) GetTotal() int32 {
	if o == nil || o.Total == nil {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVideoChaptersData) GetTotalOk() (*int32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetVideoChaptersData) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *GetVideoChaptersData) SetTotal(v int32) {
	o.Total = &v
}

// GetVideoChapters returns the VideoChapters field value if set, zero value otherwise.
func (o *GetVideoChaptersData) GetVideoChapters() []VideoChapter {
	if o == nil || o.VideoChapters == nil {
		var ret []VideoChapter
		return ret
	}
	return *o.VideoChapters
}

// GetVideoChaptersOk returns a tuple with the VideoChapters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVideoChaptersData) GetVideoChaptersOk() (*[]VideoChapter, bool) {
	if o == nil || o.VideoChapters == nil {
		return nil, false
	}
	return o.VideoChapters, true
}

// HasVideoChapters returns a boolean if a field has been set.
func (o *GetVideoChaptersData) HasVideoChapters() bool {
	if o != nil && o.VideoChapters != nil {
		return true
	}

	return false
}

// SetVideoChapters gets a reference to the given []VideoChapter and assigns it to the VideoChapters field.
func (o *GetVideoChaptersData) SetVideoChapters(v []VideoChapter) {
	o.VideoChapters = &v
}

type NullableGetVideoChaptersData struct {
	value *GetVideoChaptersData
	isSet bool
}

func (v NullableGetVideoChaptersData) Get() *GetVideoChaptersData {
	return v.value
}

func (v *NullableGetVideoChaptersData) Set(val *GetVideoChaptersData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetVideoChaptersData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetVideoChaptersData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetVideoChaptersData(val *GetVideoChaptersData) *NullableGetVideoChaptersData {
	return &NullableGetVideoChaptersData{value: val, isSet: true}
}
