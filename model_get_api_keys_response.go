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

// GetApiKeysResponse struct for GetApiKeysResponse
type GetApiKeysResponse struct {
	Data   *GetApiKeysData `json:"data,omitempty"`
	Status *string         `json:"status,omitempty"`
}

// NewGetApiKeysResponse instantiates a new GetApiKeysResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetApiKeysResponse() *GetApiKeysResponse {
	this := GetApiKeysResponse{}
	return &this
}

// NewGetApiKeysResponseWithDefaults instantiates a new GetApiKeysResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetApiKeysResponseWithDefaults() *GetApiKeysResponse {
	this := GetApiKeysResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetApiKeysResponse) GetData() GetApiKeysData {
	if o == nil || o.Data == nil {
		var ret GetApiKeysData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApiKeysResponse) GetDataOk() (*GetApiKeysData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetApiKeysResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GetApiKeysData and assigns it to the Data field.
func (o *GetApiKeysResponse) SetData(v GetApiKeysData) {
	o.Data = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetApiKeysResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApiKeysResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetApiKeysResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetApiKeysResponse) SetStatus(v string) {
	o.Status = &v
}

type NullableGetApiKeysResponse struct {
	value *GetApiKeysResponse
	isSet bool
}

func (v NullableGetApiKeysResponse) Get() *GetApiKeysResponse {
	return v.value
}

func (v *NullableGetApiKeysResponse) Set(val *GetApiKeysResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetApiKeysResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetApiKeysResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetApiKeysResponse(val *GetApiKeysResponse) *NullableGetApiKeysResponse {
	return &NullableGetApiKeysResponse{value: val, isSet: true}
}
