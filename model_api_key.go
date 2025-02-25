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

// ApiKey struct for ApiKey
type ApiKey struct {
	CreatedAt *string `json:"created_at,omitempty"`
	ExpiredAt *string `json:"expired_at,omitempty"`
	Id *string `json:"id,omitempty"`
	LastUsedAt *string `json:"last_used_at,omitempty"`
	Name *string `json:"name,omitempty"`
	PublicKey *string `json:"public_key,omitempty"`
	Secret *string `json:"secret,omitempty"`
	TruncatedSecret *string `json:"truncated_secret,omitempty"`
	Ttl *string `json:"ttl,omitempty"`
	Type *string `json:"type,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// NewApiKey instantiates a new ApiKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiKey() *ApiKey {
	this := ApiKey{}
	return &this
}

// NewApiKeyWithDefaults instantiates a new ApiKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiKeyWithDefaults() *ApiKey {
	this := ApiKey{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ApiKey) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKey) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ApiKey) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ApiKey) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetExpiredAt returns the ExpiredAt field value if set, zero value otherwise.
func (o *ApiKey) GetExpiredAt() string {
	if o == nil || o.ExpiredAt == nil {
		var ret string
		return ret
	}
	return *o.ExpiredAt
}

// GetExpiredAtOk returns a tuple with the ExpiredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKey) GetExpiredAtOk() (*string, bool) {
	if o == nil || o.ExpiredAt == nil {
		return nil, false
	}
	return o.ExpiredAt, true
}

// HasExpiredAt returns a boolean if a field has been set.
func (o *ApiKey) HasExpiredAt() bool {
	if o != nil && o.ExpiredAt != nil {
		return true
	}

	return false
}

// SetExpiredAt gets a reference to the given string and assigns it to the ExpiredAt field.
func (o *ApiKey) SetExpiredAt(v string) {
	o.ExpiredAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiKey) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKey) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiKey) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiKey) SetId(v string) {
	o.Id = &v
}

// GetLastUsedAt returns the LastUsedAt field value if set, zero value otherwise.
func (o *ApiKey) GetLastUsedAt() string {
	if o == nil || o.LastUsedAt == nil {
		var ret string
		return ret
	}
	return *o.LastUsedAt
}

// GetLastUsedAtOk returns a tuple with the LastUsedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKey) GetLastUsedAtOk() (*string, bool) {
	if o == nil || o.LastUsedAt == nil {
		return nil, false
	}
	return o.LastUsedAt, true
}

// HasLastUsedAt returns a boolean if a field has been set.
func (o *ApiKey) HasLastUsedAt() bool {
	if o != nil && o.LastUsedAt != nil {
		return true
	}

	return false
}

// SetLastUsedAt gets a reference to the given string and assigns it to the LastUsedAt field.
func (o *ApiKey) SetLastUsedAt(v string) {
	o.LastUsedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiKey) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKey) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiKey) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiKey) SetName(v string) {
	o.Name = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *ApiKey) GetPublicKey() string {
	if o == nil || o.PublicKey == nil {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKey) GetPublicKeyOk() (*string, bool) {
	if o == nil || o.PublicKey == nil {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *ApiKey) HasPublicKey() bool {
	if o != nil && o.PublicKey != nil {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *ApiKey) SetPublicKey(v string) {
	o.PublicKey = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *ApiKey) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKey) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *ApiKey) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *ApiKey) SetSecret(v string) {
	o.Secret = &v
}

// GetTruncatedSecret returns the TruncatedSecret field value if set, zero value otherwise.
func (o *ApiKey) GetTruncatedSecret() string {
	if o == nil || o.TruncatedSecret == nil {
		var ret string
		return ret
	}
	return *o.TruncatedSecret
}

// GetTruncatedSecretOk returns a tuple with the TruncatedSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKey) GetTruncatedSecretOk() (*string, bool) {
	if o == nil || o.TruncatedSecret == nil {
		return nil, false
	}
	return o.TruncatedSecret, true
}

// HasTruncatedSecret returns a boolean if a field has been set.
func (o *ApiKey) HasTruncatedSecret() bool {
	if o != nil && o.TruncatedSecret != nil {
		return true
	}

	return false
}

// SetTruncatedSecret gets a reference to the given string and assigns it to the TruncatedSecret field.
func (o *ApiKey) SetTruncatedSecret(v string) {
	o.TruncatedSecret = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *ApiKey) GetTtl() string {
	if o == nil || o.Ttl == nil {
		var ret string
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKey) GetTtlOk() (*string, bool) {
	if o == nil || o.Ttl == nil {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *ApiKey) HasTtl() bool {
	if o != nil && o.Ttl != nil {
		return true
	}

	return false
}

// SetTtl gets a reference to the given string and assigns it to the Ttl field.
func (o *ApiKey) SetTtl(v string) {
	o.Ttl = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ApiKey) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKey) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApiKey) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ApiKey) SetType(v string) {
	o.Type = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ApiKey) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKey) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ApiKey) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ApiKey) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}


type NullableApiKey struct {
	value *ApiKey
	isSet bool
}

func (v NullableApiKey) Get() *ApiKey {
	return v.value
}

func (v *NullableApiKey) Set(val *ApiKey) {
	v.value = val
	v.isSet = true
}

func (v NullableApiKey) IsSet() bool {
	return v.isSet
}

func (v *NullableApiKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiKey(val *ApiKey) *NullableApiKey {
	return &NullableApiKey{value: val, isSet: true}
}


