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

// ResponseError struct for ResponseError
type ResponseError struct {
	Message *string `json:"message,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewResponseError instantiates a new ResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseError() *ResponseError {
	this := ResponseError{}
	return &this
}

// NewResponseErrorWithDefaults instantiates a new ResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseErrorWithDefaults() *ResponseError {
	this := ResponseError{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ResponseError) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseError) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ResponseError) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ResponseError) SetMessage(v string) {
	o.Message = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ResponseError) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseError) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ResponseError) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ResponseError) SetStatus(v string) {
	o.Status = &v
}


type NullableResponseError struct {
	value *ResponseError
	isSet bool
}

func (v NullableResponseError) Get() *ResponseError {
	return v.value
}

func (v *NullableResponseError) Set(val *ResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseError(val *ResponseError) *NullableResponseError {
	return &NullableResponseError{value: val, isSet: true}
}


