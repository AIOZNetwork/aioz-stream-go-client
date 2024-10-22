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

// GetStreamingResponse struct for GetStreamingResponse
type GetStreamingResponse struct {
	Data   *LiveStreamVideoResponse `json:"data,omitempty"`
	Status *string                  `json:"status,omitempty"`
}

// NewGetStreamingResponse instantiates a new GetStreamingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStreamingResponse() *GetStreamingResponse {
	this := GetStreamingResponse{}
	return &this
}

// NewGetStreamingResponseWithDefaults instantiates a new GetStreamingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStreamingResponseWithDefaults() *GetStreamingResponse {
	this := GetStreamingResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetStreamingResponse) GetData() LiveStreamVideoResponse {
	if o == nil || o.Data == nil {
		var ret LiveStreamVideoResponse
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStreamingResponse) GetDataOk() (*LiveStreamVideoResponse, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetStreamingResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given LiveStreamVideoResponse and assigns it to the Data field.
func (o *GetStreamingResponse) SetData(v LiveStreamVideoResponse) {
	o.Data = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetStreamingResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStreamingResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetStreamingResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetStreamingResponse) SetStatus(v string) {
	o.Status = &v
}

type NullableGetStreamingResponse struct {
	value *GetStreamingResponse
	isSet bool
}

func (v NullableGetStreamingResponse) Get() *GetStreamingResponse {
	return v.value
}

func (v *NullableGetStreamingResponse) Set(val *GetStreamingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStreamingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStreamingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStreamingResponse(val *GetStreamingResponse) *NullableGetStreamingResponse {
	return &NullableGetStreamingResponse{value: val, isSet: true}
}
