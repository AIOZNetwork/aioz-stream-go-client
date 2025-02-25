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

// PlaylistItemVideo struct for PlaylistItemVideo
type PlaylistItemVideo struct {
	Chapters *[]VideoChapter `json:"chapters,omitempty"`
	Duration *float32 `json:"duration,omitempty"`
	HlsUrl *string `json:"hls_url,omitempty"`
	Qualities *string `json:"qualities,omitempty"`
	ThumbnailUrl *string `json:"thumbnail_url,omitempty"`
	Title *string `json:"title,omitempty"`
}

// NewPlaylistItemVideo instantiates a new PlaylistItemVideo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaylistItemVideo() *PlaylistItemVideo {
	this := PlaylistItemVideo{}
	return &this
}

// NewPlaylistItemVideoWithDefaults instantiates a new PlaylistItemVideo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaylistItemVideoWithDefaults() *PlaylistItemVideo {
	this := PlaylistItemVideo{}
	return &this
}

// GetChapters returns the Chapters field value if set, zero value otherwise.
func (o *PlaylistItemVideo) GetChapters() []VideoChapter {
	if o == nil || o.Chapters == nil {
		var ret []VideoChapter
		return ret
	}
	return *o.Chapters
}

// GetChaptersOk returns a tuple with the Chapters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaylistItemVideo) GetChaptersOk() (*[]VideoChapter, bool) {
	if o == nil || o.Chapters == nil {
		return nil, false
	}
	return o.Chapters, true
}

// HasChapters returns a boolean if a field has been set.
func (o *PlaylistItemVideo) HasChapters() bool {
	if o != nil && o.Chapters != nil {
		return true
	}

	return false
}

// SetChapters gets a reference to the given []VideoChapter and assigns it to the Chapters field.
func (o *PlaylistItemVideo) SetChapters(v []VideoChapter) {
	o.Chapters = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *PlaylistItemVideo) GetDuration() float32 {
	if o == nil || o.Duration == nil {
		var ret float32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaylistItemVideo) GetDurationOk() (*float32, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *PlaylistItemVideo) HasDuration() bool {
	if o != nil && o.Duration != nil {
		return true
	}

	return false
}

// SetDuration gets a reference to the given float32 and assigns it to the Duration field.
func (o *PlaylistItemVideo) SetDuration(v float32) {
	o.Duration = &v
}

// GetHlsUrl returns the HlsUrl field value if set, zero value otherwise.
func (o *PlaylistItemVideo) GetHlsUrl() string {
	if o == nil || o.HlsUrl == nil {
		var ret string
		return ret
	}
	return *o.HlsUrl
}

// GetHlsUrlOk returns a tuple with the HlsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaylistItemVideo) GetHlsUrlOk() (*string, bool) {
	if o == nil || o.HlsUrl == nil {
		return nil, false
	}
	return o.HlsUrl, true
}

// HasHlsUrl returns a boolean if a field has been set.
func (o *PlaylistItemVideo) HasHlsUrl() bool {
	if o != nil && o.HlsUrl != nil {
		return true
	}

	return false
}

// SetHlsUrl gets a reference to the given string and assigns it to the HlsUrl field.
func (o *PlaylistItemVideo) SetHlsUrl(v string) {
	o.HlsUrl = &v
}

// GetQualities returns the Qualities field value if set, zero value otherwise.
func (o *PlaylistItemVideo) GetQualities() string {
	if o == nil || o.Qualities == nil {
		var ret string
		return ret
	}
	return *o.Qualities
}

// GetQualitiesOk returns a tuple with the Qualities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaylistItemVideo) GetQualitiesOk() (*string, bool) {
	if o == nil || o.Qualities == nil {
		return nil, false
	}
	return o.Qualities, true
}

// HasQualities returns a boolean if a field has been set.
func (o *PlaylistItemVideo) HasQualities() bool {
	if o != nil && o.Qualities != nil {
		return true
	}

	return false
}

// SetQualities gets a reference to the given string and assigns it to the Qualities field.
func (o *PlaylistItemVideo) SetQualities(v string) {
	o.Qualities = &v
}

// GetThumbnailUrl returns the ThumbnailUrl field value if set, zero value otherwise.
func (o *PlaylistItemVideo) GetThumbnailUrl() string {
	if o == nil || o.ThumbnailUrl == nil {
		var ret string
		return ret
	}
	return *o.ThumbnailUrl
}

// GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaylistItemVideo) GetThumbnailUrlOk() (*string, bool) {
	if o == nil || o.ThumbnailUrl == nil {
		return nil, false
	}
	return o.ThumbnailUrl, true
}

// HasThumbnailUrl returns a boolean if a field has been set.
func (o *PlaylistItemVideo) HasThumbnailUrl() bool {
	if o != nil && o.ThumbnailUrl != nil {
		return true
	}

	return false
}

// SetThumbnailUrl gets a reference to the given string and assigns it to the ThumbnailUrl field.
func (o *PlaylistItemVideo) SetThumbnailUrl(v string) {
	o.ThumbnailUrl = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *PlaylistItemVideo) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaylistItemVideo) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *PlaylistItemVideo) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *PlaylistItemVideo) SetTitle(v string) {
	o.Title = &v
}


type NullablePlaylistItemVideo struct {
	value *PlaylistItemVideo
	isSet bool
}

func (v NullablePlaylistItemVideo) Get() *PlaylistItemVideo {
	return v.value
}

func (v *NullablePlaylistItemVideo) Set(val *PlaylistItemVideo) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaylistItemVideo) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaylistItemVideo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaylistItemVideo(val *PlaylistItemVideo) *NullablePlaylistItemVideo {
	return &NullablePlaylistItemVideo{value: val, isSet: true}
}


