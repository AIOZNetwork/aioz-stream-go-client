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

// Asset struct for Asset
type Asset struct {
	Logo          *string `json:"logo,omitempty"`
	LogoImageLink *string `json:"logo_image_link,omitempty"`
	LogoLink      *string `json:"logo_link,omitempty"`
}

// NewAsset instantiates a new Asset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAsset() *Asset {
	this := Asset{}
	return &this
}

// NewAssetWithDefaults instantiates a new Asset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetWithDefaults() *Asset {
	this := Asset{}
	return &this
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *Asset) GetLogo() string {
	if o == nil || o.Logo == nil {
		var ret string
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetLogoOk() (*string, bool) {
	if o == nil || o.Logo == nil {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *Asset) HasLogo() bool {
	if o != nil && o.Logo != nil {
		return true
	}

	return false
}

// SetLogo gets a reference to the given string and assigns it to the Logo field.
func (o *Asset) SetLogo(v string) {
	o.Logo = &v
}

// GetLogoImageLink returns the LogoImageLink field value if set, zero value otherwise.
func (o *Asset) GetLogoImageLink() string {
	if o == nil || o.LogoImageLink == nil {
		var ret string
		return ret
	}
	return *o.LogoImageLink
}

// GetLogoImageLinkOk returns a tuple with the LogoImageLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetLogoImageLinkOk() (*string, bool) {
	if o == nil || o.LogoImageLink == nil {
		return nil, false
	}
	return o.LogoImageLink, true
}

// HasLogoImageLink returns a boolean if a field has been set.
func (o *Asset) HasLogoImageLink() bool {
	if o != nil && o.LogoImageLink != nil {
		return true
	}

	return false
}

// SetLogoImageLink gets a reference to the given string and assigns it to the LogoImageLink field.
func (o *Asset) SetLogoImageLink(v string) {
	o.LogoImageLink = &v
}

// GetLogoLink returns the LogoLink field value if set, zero value otherwise.
func (o *Asset) GetLogoLink() string {
	if o == nil || o.LogoLink == nil {
		var ret string
		return ret
	}
	return *o.LogoLink
}

// GetLogoLinkOk returns a tuple with the LogoLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetLogoLinkOk() (*string, bool) {
	if o == nil || o.LogoLink == nil {
		return nil, false
	}
	return o.LogoLink, true
}

// HasLogoLink returns a boolean if a field has been set.
func (o *Asset) HasLogoLink() bool {
	if o != nil && o.LogoLink != nil {
		return true
	}

	return false
}

// SetLogoLink gets a reference to the given string and assigns it to the LogoLink field.
func (o *Asset) SetLogoLink(v string) {
	o.LogoLink = &v
}

type NullableAsset struct {
	value *Asset
	isSet bool
}

func (v NullableAsset) Get() *Asset {
	return v.value
}

func (v *NullableAsset) Set(val *Asset) {
	v.value = val
	v.isSet = true
}

func (v NullableAsset) IsSet() bool {
	return v.isSet
}

func (v *NullableAsset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsset(val *Asset) *NullableAsset {
	return &NullableAsset{value: val, isSet: true}
}
