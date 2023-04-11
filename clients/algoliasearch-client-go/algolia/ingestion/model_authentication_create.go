// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package ingestion

import (
	"encoding/json"
	"fmt"
)

// AuthenticationCreate The payload when creating an authentication.
type AuthenticationCreate struct {
	Type AuthenticationType `json:"type"`
	// An human readable name describing the object.
	Name     string    `json:"name"`
	Platform *Platform `json:"platform,omitempty"`
	Input    AuthInput `json:"input"`
}

type AuthenticationCreateOption func(f *AuthenticationCreate)

func WithAuthenticationCreatePlatform(val Platform) AuthenticationCreateOption {
	return func(f *AuthenticationCreate) {
		f.Platform = &val
	}
}

// NewAuthenticationCreate instantiates a new AuthenticationCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationCreate(type_ AuthenticationType, name string, input AuthInput, opts ...AuthenticationCreateOption) *AuthenticationCreate {
	this := &AuthenticationCreate{}
	this.Type = type_
	this.Name = name
	this.Input = input
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewAuthenticationCreateWithDefaults instantiates a new AuthenticationCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationCreateWithDefaults() *AuthenticationCreate {
	this := &AuthenticationCreate{}
	return this
}

// GetType returns the Type field value
func (o *AuthenticationCreate) GetType() AuthenticationType {
	if o == nil {
		var ret AuthenticationType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AuthenticationCreate) GetTypeOk() (*AuthenticationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AuthenticationCreate) SetType(v AuthenticationType) {
	o.Type = v
}

// GetName returns the Name field value
func (o *AuthenticationCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AuthenticationCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AuthenticationCreate) SetName(v string) {
	o.Name = v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *AuthenticationCreate) GetPlatform() Platform {
	if o == nil || o.Platform == nil {
		var ret Platform
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationCreate) GetPlatformOk() (*Platform, bool) {
	if o == nil || o.Platform == nil {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *AuthenticationCreate) HasPlatform() bool {
	if o != nil && o.Platform != nil {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given Platform and assigns it to the Platform field.
func (o *AuthenticationCreate) SetPlatform(v Platform) {
	o.Platform = &v
}

// GetInput returns the Input field value
func (o *AuthenticationCreate) GetInput() AuthInput {
	if o == nil {
		var ret AuthInput
		return ret
	}

	return o.Input
}

// GetInputOk returns a tuple with the Input field value
// and a boolean to check if the value has been set.
func (o *AuthenticationCreate) GetInputOk() (*AuthInput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Input, true
}

// SetInput sets field value
func (o *AuthenticationCreate) SetInput(v AuthInput) {
	o.Input = v
}

func (o AuthenticationCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Platform != nil {
		toSerialize["platform"] = o.Platform
	}
	if true {
		toSerialize["input"] = o.Input
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticationCreate) String() string {
	out := ""
	out += fmt.Sprintf("  type=%v\n", o.Type)
	out += fmt.Sprintf("  name=%v\n", o.Name)
	out += fmt.Sprintf("  platform=%v\n", o.Platform)
	out += fmt.Sprintf("  input=%v\n", o.Input)
	return fmt.Sprintf("AuthenticationCreate {\n%s}", out)
}

type NullableAuthenticationCreate struct {
	value *AuthenticationCreate
	isSet bool
}

func (v NullableAuthenticationCreate) Get() *AuthenticationCreate {
	return v.value
}

func (v *NullableAuthenticationCreate) Set(val *AuthenticationCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationCreate(val *AuthenticationCreate) *NullableAuthenticationCreate {
	return &NullableAuthenticationCreate{value: val, isSet: true}
}

func (v NullableAuthenticationCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}