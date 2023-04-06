// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package predict

import (
	"encoding/json"
	"fmt"
)

// UpdateModelParams struct for UpdateModelParams
type UpdateModelParams struct {
	// The model’s instance name.
	Name            *string      `json:"name,omitempty"`
	ModelAttributes []string     `json:"modelAttributes,omitempty"`
	ModelStatus     *ModelStatus `json:"modelStatus,omitempty"`
}

type UpdateModelParamsOption func(f *UpdateModelParams)

func WithUpdateModelParamsName(val string) UpdateModelParamsOption {
	return func(f *UpdateModelParams) {
		f.Name = &val
	}
}

func WithUpdateModelParamsModelAttributes(val []string) UpdateModelParamsOption {
	return func(f *UpdateModelParams) {
		f.ModelAttributes = val
	}
}

func WithUpdateModelParamsModelStatus(val ModelStatus) UpdateModelParamsOption {
	return func(f *UpdateModelParams) {
		f.ModelStatus = &val
	}
}

// NewUpdateModelParams instantiates a new UpdateModelParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateModelParams(opts ...UpdateModelParamsOption) *UpdateModelParams {
	this := &UpdateModelParams{}
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewUpdateModelParamsWithDefaults instantiates a new UpdateModelParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateModelParamsWithDefaults() *UpdateModelParams {
	this := &UpdateModelParams{}
	return this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateModelParams) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateModelParams) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateModelParams) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateModelParams) SetName(v string) {
	o.Name = &v
}

// GetModelAttributes returns the ModelAttributes field value if set, zero value otherwise.
func (o *UpdateModelParams) GetModelAttributes() []string {
	if o == nil || o.ModelAttributes == nil {
		var ret []string
		return ret
	}
	return o.ModelAttributes
}

// GetModelAttributesOk returns a tuple with the ModelAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateModelParams) GetModelAttributesOk() ([]string, bool) {
	if o == nil || o.ModelAttributes == nil {
		return nil, false
	}
	return o.ModelAttributes, true
}

// HasModelAttributes returns a boolean if a field has been set.
func (o *UpdateModelParams) HasModelAttributes() bool {
	if o != nil && o.ModelAttributes != nil {
		return true
	}

	return false
}

// SetModelAttributes gets a reference to the given []string and assigns it to the ModelAttributes field.
func (o *UpdateModelParams) SetModelAttributes(v []string) {
	o.ModelAttributes = v
}

// GetModelStatus returns the ModelStatus field value if set, zero value otherwise.
func (o *UpdateModelParams) GetModelStatus() ModelStatus {
	if o == nil || o.ModelStatus == nil {
		var ret ModelStatus
		return ret
	}
	return *o.ModelStatus
}

// GetModelStatusOk returns a tuple with the ModelStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateModelParams) GetModelStatusOk() (*ModelStatus, bool) {
	if o == nil || o.ModelStatus == nil {
		return nil, false
	}
	return o.ModelStatus, true
}

// HasModelStatus returns a boolean if a field has been set.
func (o *UpdateModelParams) HasModelStatus() bool {
	if o != nil && o.ModelStatus != nil {
		return true
	}

	return false
}

// SetModelStatus gets a reference to the given ModelStatus and assigns it to the ModelStatus field.
func (o *UpdateModelParams) SetModelStatus(v ModelStatus) {
	o.ModelStatus = &v
}

func (o UpdateModelParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ModelAttributes != nil {
		toSerialize["modelAttributes"] = o.ModelAttributes
	}
	if o.ModelStatus != nil {
		toSerialize["modelStatus"] = o.ModelStatus
	}
	return json.Marshal(toSerialize)
}

func (o UpdateModelParams) String() string {
	out := ""
	out += fmt.Sprintf("  name=%v\n", o.Name)
	out += fmt.Sprintf("  modelAttributes=%v\n", o.ModelAttributes)
	out += fmt.Sprintf("  modelStatus=%v\n", o.ModelStatus)
	return fmt.Sprintf("UpdateModelParams {\n%s}", out)
}

type NullableUpdateModelParams struct {
	value *UpdateModelParams
	isSet bool
}

func (v NullableUpdateModelParams) Get() *UpdateModelParams {
	return v.value
}

func (v *NullableUpdateModelParams) Set(val *UpdateModelParams) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateModelParams) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateModelParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateModelParams(val *UpdateModelParams) *NullableUpdateModelParams {
	return &NullableUpdateModelParams{value: val, isSet: true}
}

func (v NullableUpdateModelParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateModelParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
