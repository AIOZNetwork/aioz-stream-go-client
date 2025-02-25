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

// GetLiveStreamKeyResponse struct for GetLiveStreamKeyResponse
type GetLiveStreamKeyResponse struct {
	Data *GetLiveStreamKeyData `json:"data,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewGetLiveStreamKeyResponse instantiates a new GetLiveStreamKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLiveStreamKeyResponse() *GetLiveStreamKeyResponse {
	this := GetLiveStreamKeyResponse{}
	return &this
}

// NewGetLiveStreamKeyResponseWithDefaults instantiates a new GetLiveStreamKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLiveStreamKeyResponseWithDefaults() *GetLiveStreamKeyResponse {
	this := GetLiveStreamKeyResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetLiveStreamKeyResponse) GetData() GetLiveStreamKeyData {
	if o == nil || o.Data == nil {
		var ret GetLiveStreamKeyData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLiveStreamKeyResponse) GetDataOk() (*GetLiveStreamKeyData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetLiveStreamKeyResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GetLiveStreamKeyData and assigns it to the Data field.
func (o *GetLiveStreamKeyResponse) SetData(v GetLiveStreamKeyData) {
	o.Data = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetLiveStreamKeyResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLiveStreamKeyResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetLiveStreamKeyResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetLiveStreamKeyResponse) SetStatus(v string) {
	o.Status = &v
}


type NullableGetLiveStreamKeyResponse struct {
	value *GetLiveStreamKeyResponse
	isSet bool
}

func (v NullableGetLiveStreamKeyResponse) Get() *GetLiveStreamKeyResponse {
	return v.value
}

func (v *NullableGetLiveStreamKeyResponse) Set(val *GetLiveStreamKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLiveStreamKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLiveStreamKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLiveStreamKeyResponse(val *GetLiveStreamKeyResponse) *NullableGetLiveStreamKeyResponse {
	return &NullableGetLiveStreamKeyResponse{value: val, isSet: true}
}


