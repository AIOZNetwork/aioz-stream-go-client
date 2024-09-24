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

// QualityObject struct for QualityObject
type QualityObject struct {
	Name   *string `json:"name,omitempty"`
	Status *string `json:"status,omitempty"`
	Type   *string `json:"type,omitempty"`
}

// NewQualityObject instantiates a new QualityObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQualityObject() *QualityObject {
	this := QualityObject{}
	return &this
}

// NewQualityObjectWithDefaults instantiates a new QualityObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQualityObjectWithDefaults() *QualityObject {
	this := QualityObject{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *QualityObject) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QualityObject) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *QualityObject) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *QualityObject) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *QualityObject) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QualityObject) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *QualityObject) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *QualityObject) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *QualityObject) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QualityObject) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *QualityObject) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *QualityObject) SetType(v string) {
	o.Type = &v
}

type NullableQualityObject struct {
	value *QualityObject
	isSet bool
}

func (v NullableQualityObject) Get() *QualityObject {
	return v.value
}

func (v *NullableQualityObject) Set(val *QualityObject) {
	v.value = val
	v.isSet = true
}

func (v NullableQualityObject) IsSet() bool {
	return v.isSet
}

func (v *NullableQualityObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQualityObject(val *QualityObject) *NullableQualityObject {
	return &NullableQualityObject{value: val, isSet: true}
}
