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

// GetTranscodeCostResponse struct for GetTranscodeCostResponse
type GetTranscodeCostResponse struct {
	Data   *GetTranscodeCostData `json:"data,omitempty"`
	Status *string               `json:"status,omitempty"`
}

// NewGetTranscodeCostResponse instantiates a new GetTranscodeCostResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTranscodeCostResponse() *GetTranscodeCostResponse {
	this := GetTranscodeCostResponse{}
	return &this
}

// NewGetTranscodeCostResponseWithDefaults instantiates a new GetTranscodeCostResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTranscodeCostResponseWithDefaults() *GetTranscodeCostResponse {
	this := GetTranscodeCostResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetTranscodeCostResponse) GetData() GetTranscodeCostData {
	if o == nil || o.Data == nil {
		var ret GetTranscodeCostData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTranscodeCostResponse) GetDataOk() (*GetTranscodeCostData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetTranscodeCostResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GetTranscodeCostData and assigns it to the Data field.
func (o *GetTranscodeCostResponse) SetData(v GetTranscodeCostData) {
	o.Data = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetTranscodeCostResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTranscodeCostResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetTranscodeCostResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetTranscodeCostResponse) SetStatus(v string) {
	o.Status = &v
}

type NullableGetTranscodeCostResponse struct {
	value *GetTranscodeCostResponse
	isSet bool
}

func (v NullableGetTranscodeCostResponse) Get() *GetTranscodeCostResponse {
	return v.value
}

func (v *NullableGetTranscodeCostResponse) Set(val *GetTranscodeCostResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTranscodeCostResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTranscodeCostResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTranscodeCostResponse(val *GetTranscodeCostResponse) *NullableGetTranscodeCostResponse {
	return &NullableGetTranscodeCostResponse{value: val, isSet: true}
}
