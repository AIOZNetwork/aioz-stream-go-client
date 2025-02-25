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

// CreateApiKeyResponse struct for CreateApiKeyResponse
type CreateApiKeyResponse struct {
	Data *CreateApiKeyData `json:"data,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewCreateApiKeyResponse instantiates a new CreateApiKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateApiKeyResponse() *CreateApiKeyResponse {
	this := CreateApiKeyResponse{}
	return &this
}

// NewCreateApiKeyResponseWithDefaults instantiates a new CreateApiKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateApiKeyResponseWithDefaults() *CreateApiKeyResponse {
	this := CreateApiKeyResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateApiKeyResponse) GetData() CreateApiKeyData {
	if o == nil || o.Data == nil {
		var ret CreateApiKeyData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateApiKeyResponse) GetDataOk() (*CreateApiKeyData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateApiKeyResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given CreateApiKeyData and assigns it to the Data field.
func (o *CreateApiKeyResponse) SetData(v CreateApiKeyData) {
	o.Data = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CreateApiKeyResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateApiKeyResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CreateApiKeyResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CreateApiKeyResponse) SetStatus(v string) {
	o.Status = &v
}


type NullableCreateApiKeyResponse struct {
	value *CreateApiKeyResponse
	isSet bool
}

func (v NullableCreateApiKeyResponse) Get() *CreateApiKeyResponse {
	return v.value
}

func (v *NullableCreateApiKeyResponse) Set(val *CreateApiKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateApiKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateApiKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateApiKeyResponse(val *CreateApiKeyResponse) *NullableCreateApiKeyResponse {
	return &NullableCreateApiKeyResponse{value: val, isSet: true}
}


