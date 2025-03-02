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

// UpdateLiveStreamKeyData struct for UpdateLiveStreamKeyData
type UpdateLiveStreamKeyData struct {
	CreatedAt *string `json:"created_at,omitempty"`
	Id        *string `json:"id,omitempty"`
	Name      *string `json:"name,omitempty"`
	RtmpUrl   *string `json:"rtmp_url,omitempty"`
	Save      *bool   `json:"save,omitempty"`
	StreamKey *string `json:"stream_key,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	UserId    *string `json:"user_id,omitempty"`
}

// NewUpdateLiveStreamKeyData instantiates a new UpdateLiveStreamKeyData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateLiveStreamKeyData() *UpdateLiveStreamKeyData {
	this := UpdateLiveStreamKeyData{}
	return &this
}

// NewUpdateLiveStreamKeyDataWithDefaults instantiates a new UpdateLiveStreamKeyData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateLiveStreamKeyDataWithDefaults() *UpdateLiveStreamKeyData {
	this := UpdateLiveStreamKeyData{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *UpdateLiveStreamKeyData) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLiveStreamKeyData) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *UpdateLiveStreamKeyData) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *UpdateLiveStreamKeyData) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdateLiveStreamKeyData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLiveStreamKeyData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdateLiveStreamKeyData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UpdateLiveStreamKeyData) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateLiveStreamKeyData) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLiveStreamKeyData) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateLiveStreamKeyData) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateLiveStreamKeyData) SetName(v string) {
	o.Name = &v
}

// GetRtmpUrl returns the RtmpUrl field value if set, zero value otherwise.
func (o *UpdateLiveStreamKeyData) GetRtmpUrl() string {
	if o == nil || o.RtmpUrl == nil {
		var ret string
		return ret
	}
	return *o.RtmpUrl
}

// GetRtmpUrlOk returns a tuple with the RtmpUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLiveStreamKeyData) GetRtmpUrlOk() (*string, bool) {
	if o == nil || o.RtmpUrl == nil {
		return nil, false
	}
	return o.RtmpUrl, true
}

// HasRtmpUrl returns a boolean if a field has been set.
func (o *UpdateLiveStreamKeyData) HasRtmpUrl() bool {
	if o != nil && o.RtmpUrl != nil {
		return true
	}

	return false
}

// SetRtmpUrl gets a reference to the given string and assigns it to the RtmpUrl field.
func (o *UpdateLiveStreamKeyData) SetRtmpUrl(v string) {
	o.RtmpUrl = &v
}

// GetSave returns the Save field value if set, zero value otherwise.
func (o *UpdateLiveStreamKeyData) GetSave() bool {
	if o == nil || o.Save == nil {
		var ret bool
		return ret
	}
	return *o.Save
}

// GetSaveOk returns a tuple with the Save field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLiveStreamKeyData) GetSaveOk() (*bool, bool) {
	if o == nil || o.Save == nil {
		return nil, false
	}
	return o.Save, true
}

// HasSave returns a boolean if a field has been set.
func (o *UpdateLiveStreamKeyData) HasSave() bool {
	if o != nil && o.Save != nil {
		return true
	}

	return false
}

// SetSave gets a reference to the given bool and assigns it to the Save field.
func (o *UpdateLiveStreamKeyData) SetSave(v bool) {
	o.Save = &v
}

// GetStreamKey returns the StreamKey field value if set, zero value otherwise.
func (o *UpdateLiveStreamKeyData) GetStreamKey() string {
	if o == nil || o.StreamKey == nil {
		var ret string
		return ret
	}
	return *o.StreamKey
}

// GetStreamKeyOk returns a tuple with the StreamKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLiveStreamKeyData) GetStreamKeyOk() (*string, bool) {
	if o == nil || o.StreamKey == nil {
		return nil, false
	}
	return o.StreamKey, true
}

// HasStreamKey returns a boolean if a field has been set.
func (o *UpdateLiveStreamKeyData) HasStreamKey() bool {
	if o != nil && o.StreamKey != nil {
		return true
	}

	return false
}

// SetStreamKey gets a reference to the given string and assigns it to the StreamKey field.
func (o *UpdateLiveStreamKeyData) SetStreamKey(v string) {
	o.StreamKey = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *UpdateLiveStreamKeyData) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLiveStreamKeyData) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *UpdateLiveStreamKeyData) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *UpdateLiveStreamKeyData) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UpdateLiveStreamKeyData) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLiveStreamKeyData) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UpdateLiveStreamKeyData) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *UpdateLiveStreamKeyData) SetUserId(v string) {
	o.UserId = &v
}

type NullableUpdateLiveStreamKeyData struct {
	value *UpdateLiveStreamKeyData
	isSet bool
}

func (v NullableUpdateLiveStreamKeyData) Get() *UpdateLiveStreamKeyData {
	return v.value
}

func (v *NullableUpdateLiveStreamKeyData) Set(val *UpdateLiveStreamKeyData) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateLiveStreamKeyData) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateLiveStreamKeyData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateLiveStreamKeyData(val *UpdateLiveStreamKeyData) *NullableUpdateLiveStreamKeyData {
	return &NullableUpdateLiveStreamKeyData{value: val, isSet: true}
}
