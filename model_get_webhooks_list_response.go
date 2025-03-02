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

// GetWebhooksListResponse struct for GetWebhooksListResponse
type GetWebhooksListResponse struct {
	Data   *GetWebhooksListData `json:"data,omitempty"`
	Status *string              `json:"status,omitempty"`
}

// NewGetWebhooksListResponse instantiates a new GetWebhooksListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWebhooksListResponse() *GetWebhooksListResponse {
	this := GetWebhooksListResponse{}
	return &this
}

// NewGetWebhooksListResponseWithDefaults instantiates a new GetWebhooksListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWebhooksListResponseWithDefaults() *GetWebhooksListResponse {
	this := GetWebhooksListResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetWebhooksListResponse) GetData() GetWebhooksListData {
	if o == nil || o.Data == nil {
		var ret GetWebhooksListData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWebhooksListResponse) GetDataOk() (*GetWebhooksListData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetWebhooksListResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GetWebhooksListData and assigns it to the Data field.
func (o *GetWebhooksListResponse) SetData(v GetWebhooksListData) {
	o.Data = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetWebhooksListResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWebhooksListResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetWebhooksListResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetWebhooksListResponse) SetStatus(v string) {
	o.Status = &v
}

type NullableGetWebhooksListResponse struct {
	value *GetWebhooksListResponse
	isSet bool
}

func (v NullableGetWebhooksListResponse) Get() *GetWebhooksListResponse {
	return v.value
}

func (v *NullableGetWebhooksListResponse) Set(val *GetWebhooksListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWebhooksListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWebhooksListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWebhooksListResponse(val *GetWebhooksListResponse) *NullableGetWebhooksListResponse {
	return &NullableGetWebhooksListResponse{value: val, isSet: true}
}
