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

// UpdatePlayerThemeRequest struct for UpdatePlayerThemeRequest
type UpdatePlayerThemeRequest struct {
	Controls  *Controls `json:"controls,omitempty"`
	IsDefault *bool     `json:"is_default,omitempty"`
	Name      *string   `json:"name,omitempty"`
	Theme     *Theme    `json:"theme,omitempty"`
}

// NewUpdatePlayerThemeRequest instantiates a new UpdatePlayerThemeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePlayerThemeRequest() *UpdatePlayerThemeRequest {
	this := UpdatePlayerThemeRequest{}
	return &this
}

// NewUpdatePlayerThemeRequestWithDefaults instantiates a new UpdatePlayerThemeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePlayerThemeRequestWithDefaults() *UpdatePlayerThemeRequest {
	this := UpdatePlayerThemeRequest{}
	return &this
}

// GetControls returns the Controls field value if set, zero value otherwise.
func (o *UpdatePlayerThemeRequest) GetControls() Controls {
	if o == nil || o.Controls == nil {
		var ret Controls
		return ret
	}
	return *o.Controls
}

// GetControlsOk returns a tuple with the Controls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePlayerThemeRequest) GetControlsOk() (*Controls, bool) {
	if o == nil || o.Controls == nil {
		return nil, false
	}
	return o.Controls, true
}

// HasControls returns a boolean if a field has been set.
func (o *UpdatePlayerThemeRequest) HasControls() bool {
	if o != nil && o.Controls != nil {
		return true
	}

	return false
}

// SetControls gets a reference to the given Controls and assigns it to the Controls field.
func (o *UpdatePlayerThemeRequest) SetControls(v Controls) {
	o.Controls = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *UpdatePlayerThemeRequest) GetIsDefault() bool {
	if o == nil || o.IsDefault == nil {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePlayerThemeRequest) GetIsDefaultOk() (*bool, bool) {
	if o == nil || o.IsDefault == nil {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *UpdatePlayerThemeRequest) HasIsDefault() bool {
	if o != nil && o.IsDefault != nil {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *UpdatePlayerThemeRequest) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdatePlayerThemeRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePlayerThemeRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdatePlayerThemeRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdatePlayerThemeRequest) SetName(v string) {
	o.Name = &v
}

// GetTheme returns the Theme field value if set, zero value otherwise.
func (o *UpdatePlayerThemeRequest) GetTheme() Theme {
	if o == nil || o.Theme == nil {
		var ret Theme
		return ret
	}
	return *o.Theme
}

// GetThemeOk returns a tuple with the Theme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePlayerThemeRequest) GetThemeOk() (*Theme, bool) {
	if o == nil || o.Theme == nil {
		return nil, false
	}
	return o.Theme, true
}

// HasTheme returns a boolean if a field has been set.
func (o *UpdatePlayerThemeRequest) HasTheme() bool {
	if o != nil && o.Theme != nil {
		return true
	}

	return false
}

// SetTheme gets a reference to the given Theme and assigns it to the Theme field.
func (o *UpdatePlayerThemeRequest) SetTheme(v Theme) {
	o.Theme = &v
}

type NullableUpdatePlayerThemeRequest struct {
	value *UpdatePlayerThemeRequest
	isSet bool
}

func (v NullableUpdatePlayerThemeRequest) Get() *UpdatePlayerThemeRequest {
	return v.value
}

func (v *NullableUpdatePlayerThemeRequest) Set(val *UpdatePlayerThemeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePlayerThemeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePlayerThemeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePlayerThemeRequest(val *UpdatePlayerThemeRequest) *NullableUpdatePlayerThemeRequest {
	return &NullableUpdatePlayerThemeRequest{value: val, isSet: true}
}
