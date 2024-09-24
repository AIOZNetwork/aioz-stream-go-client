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

// VideoWatermark struct for VideoWatermark
type VideoWatermark struct {
	Bottom  *string `json:"bottom,omitempty"`
	Height  *string `json:"height,omitempty"`
	Id      *string `json:"id,omitempty"`
	Left    *string `json:"left,omitempty"`
	Opacity *string `json:"opacity,omitempty"`
	Right   *string `json:"right,omitempty"`
	Top     *string `json:"top,omitempty"`
	Width   *string `json:"width,omitempty"`
}

// NewVideoWatermark instantiates a new VideoWatermark object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoWatermark() *VideoWatermark {
	this := VideoWatermark{}
	return &this
}

// NewVideoWatermarkWithDefaults instantiates a new VideoWatermark object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoWatermarkWithDefaults() *VideoWatermark {
	this := VideoWatermark{}
	return &this
}

// GetBottom returns the Bottom field value if set, zero value otherwise.
func (o *VideoWatermark) GetBottom() string {
	if o == nil || o.Bottom == nil {
		var ret string
		return ret
	}
	return *o.Bottom
}

// GetBottomOk returns a tuple with the Bottom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoWatermark) GetBottomOk() (*string, bool) {
	if o == nil || o.Bottom == nil {
		return nil, false
	}
	return o.Bottom, true
}

// HasBottom returns a boolean if a field has been set.
func (o *VideoWatermark) HasBottom() bool {
	if o != nil && o.Bottom != nil {
		return true
	}

	return false
}

// SetBottom gets a reference to the given string and assigns it to the Bottom field.
func (o *VideoWatermark) SetBottom(v string) {
	o.Bottom = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *VideoWatermark) GetHeight() string {
	if o == nil || o.Height == nil {
		var ret string
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoWatermark) GetHeightOk() (*string, bool) {
	if o == nil || o.Height == nil {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *VideoWatermark) HasHeight() bool {
	if o != nil && o.Height != nil {
		return true
	}

	return false
}

// SetHeight gets a reference to the given string and assigns it to the Height field.
func (o *VideoWatermark) SetHeight(v string) {
	o.Height = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VideoWatermark) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoWatermark) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VideoWatermark) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VideoWatermark) SetId(v string) {
	o.Id = &v
}

// GetLeft returns the Left field value if set, zero value otherwise.
func (o *VideoWatermark) GetLeft() string {
	if o == nil || o.Left == nil {
		var ret string
		return ret
	}
	return *o.Left
}

// GetLeftOk returns a tuple with the Left field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoWatermark) GetLeftOk() (*string, bool) {
	if o == nil || o.Left == nil {
		return nil, false
	}
	return o.Left, true
}

// HasLeft returns a boolean if a field has been set.
func (o *VideoWatermark) HasLeft() bool {
	if o != nil && o.Left != nil {
		return true
	}

	return false
}

// SetLeft gets a reference to the given string and assigns it to the Left field.
func (o *VideoWatermark) SetLeft(v string) {
	o.Left = &v
}

// GetOpacity returns the Opacity field value if set, zero value otherwise.
func (o *VideoWatermark) GetOpacity() string {
	if o == nil || o.Opacity == nil {
		var ret string
		return ret
	}
	return *o.Opacity
}

// GetOpacityOk returns a tuple with the Opacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoWatermark) GetOpacityOk() (*string, bool) {
	if o == nil || o.Opacity == nil {
		return nil, false
	}
	return o.Opacity, true
}

// HasOpacity returns a boolean if a field has been set.
func (o *VideoWatermark) HasOpacity() bool {
	if o != nil && o.Opacity != nil {
		return true
	}

	return false
}

// SetOpacity gets a reference to the given string and assigns it to the Opacity field.
func (o *VideoWatermark) SetOpacity(v string) {
	o.Opacity = &v
}

// GetRight returns the Right field value if set, zero value otherwise.
func (o *VideoWatermark) GetRight() string {
	if o == nil || o.Right == nil {
		var ret string
		return ret
	}
	return *o.Right
}

// GetRightOk returns a tuple with the Right field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoWatermark) GetRightOk() (*string, bool) {
	if o == nil || o.Right == nil {
		return nil, false
	}
	return o.Right, true
}

// HasRight returns a boolean if a field has been set.
func (o *VideoWatermark) HasRight() bool {
	if o != nil && o.Right != nil {
		return true
	}

	return false
}

// SetRight gets a reference to the given string and assigns it to the Right field.
func (o *VideoWatermark) SetRight(v string) {
	o.Right = &v
}

// GetTop returns the Top field value if set, zero value otherwise.
func (o *VideoWatermark) GetTop() string {
	if o == nil || o.Top == nil {
		var ret string
		return ret
	}
	return *o.Top
}

// GetTopOk returns a tuple with the Top field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoWatermark) GetTopOk() (*string, bool) {
	if o == nil || o.Top == nil {
		return nil, false
	}
	return o.Top, true
}

// HasTop returns a boolean if a field has been set.
func (o *VideoWatermark) HasTop() bool {
	if o != nil && o.Top != nil {
		return true
	}

	return false
}

// SetTop gets a reference to the given string and assigns it to the Top field.
func (o *VideoWatermark) SetTop(v string) {
	o.Top = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *VideoWatermark) GetWidth() string {
	if o == nil || o.Width == nil {
		var ret string
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoWatermark) GetWidthOk() (*string, bool) {
	if o == nil || o.Width == nil {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *VideoWatermark) HasWidth() bool {
	if o != nil && o.Width != nil {
		return true
	}

	return false
}

// SetWidth gets a reference to the given string and assigns it to the Width field.
func (o *VideoWatermark) SetWidth(v string) {
	o.Width = &v
}

type NullableVideoWatermark struct {
	value *VideoWatermark
	isSet bool
}

func (v NullableVideoWatermark) Get() *VideoWatermark {
	return v.value
}

func (v *NullableVideoWatermark) Set(val *VideoWatermark) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoWatermark) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoWatermark) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoWatermark(val *VideoWatermark) *NullableVideoWatermark {
	return &NullableVideoWatermark{value: val, isSet: true}
}
