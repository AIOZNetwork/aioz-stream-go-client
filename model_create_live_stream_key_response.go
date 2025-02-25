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

// CreateLiveStreamKeyResponse struct for CreateLiveStreamKeyResponse
type CreateLiveStreamKeyResponse struct {
	Data *LiveStreamKeyData `json:"data,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewCreateLiveStreamKeyResponse instantiates a new CreateLiveStreamKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLiveStreamKeyResponse() *CreateLiveStreamKeyResponse {
	this := CreateLiveStreamKeyResponse{}
	return &this
}

// NewCreateLiveStreamKeyResponseWithDefaults instantiates a new CreateLiveStreamKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLiveStreamKeyResponseWithDefaults() *CreateLiveStreamKeyResponse {
	this := CreateLiveStreamKeyResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateLiveStreamKeyResponse) GetData() LiveStreamKeyData {
	if o == nil || o.Data == nil {
		var ret LiveStreamKeyData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLiveStreamKeyResponse) GetDataOk() (*LiveStreamKeyData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateLiveStreamKeyResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given LiveStreamKeyData and assigns it to the Data field.
func (o *CreateLiveStreamKeyResponse) SetData(v LiveStreamKeyData) {
	o.Data = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CreateLiveStreamKeyResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLiveStreamKeyResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CreateLiveStreamKeyResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CreateLiveStreamKeyResponse) SetStatus(v string) {
	o.Status = &v
}


type NullableCreateLiveStreamKeyResponse struct {
	value *CreateLiveStreamKeyResponse
	isSet bool
}

func (v NullableCreateLiveStreamKeyResponse) Get() *CreateLiveStreamKeyResponse {
	return v.value
}

func (v *NullableCreateLiveStreamKeyResponse) Set(val *CreateLiveStreamKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLiveStreamKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLiveStreamKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLiveStreamKeyResponse(val *CreateLiveStreamKeyResponse) *NullableCreateLiveStreamKeyResponse {
	return &NullableCreateLiveStreamKeyResponse{value: val, isSet: true}
}


