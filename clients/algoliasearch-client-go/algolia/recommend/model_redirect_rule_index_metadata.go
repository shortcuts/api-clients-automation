// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package recommend

import (
	"encoding/json"
	"fmt"
)

// RedirectRuleIndexMetadata struct for RedirectRuleIndexMetadata
type RedirectRuleIndexMetadata struct {
	// Source index for the redirect rule.
	Source string `json:"source"`
	// Destination index for the redirect rule.
	Dest string `json:"dest"`
	// Reason for the redirect rule.
	Reason string `json:"reason"`
	// Status for the redirect rule.
	Succeed bool                          `json:"succeed"`
	Data    RedirectRuleIndexMetadataData `json:"data"`
}

// NewRedirectRuleIndexMetadata instantiates a new RedirectRuleIndexMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRedirectRuleIndexMetadata(source string, dest string, reason string, succeed bool, data RedirectRuleIndexMetadataData) *RedirectRuleIndexMetadata {
	this := &RedirectRuleIndexMetadata{}
	this.Source = source
	this.Dest = dest
	this.Reason = reason
	this.Succeed = succeed
	this.Data = data
	return this
}

// NewRedirectRuleIndexMetadataWithDefaults instantiates a new RedirectRuleIndexMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRedirectRuleIndexMetadataWithDefaults() *RedirectRuleIndexMetadata {
	this := &RedirectRuleIndexMetadata{}
	return this
}

// GetSource returns the Source field value
func (o *RedirectRuleIndexMetadata) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *RedirectRuleIndexMetadata) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *RedirectRuleIndexMetadata) SetSource(v string) {
	o.Source = v
}

// GetDest returns the Dest field value
func (o *RedirectRuleIndexMetadata) GetDest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dest
}

// GetDestOk returns a tuple with the Dest field value
// and a boolean to check if the value has been set.
func (o *RedirectRuleIndexMetadata) GetDestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dest, true
}

// SetDest sets field value
func (o *RedirectRuleIndexMetadata) SetDest(v string) {
	o.Dest = v
}

// GetReason returns the Reason field value
func (o *RedirectRuleIndexMetadata) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *RedirectRuleIndexMetadata) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *RedirectRuleIndexMetadata) SetReason(v string) {
	o.Reason = v
}

// GetSucceed returns the Succeed field value
func (o *RedirectRuleIndexMetadata) GetSucceed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Succeed
}

// GetSucceedOk returns a tuple with the Succeed field value
// and a boolean to check if the value has been set.
func (o *RedirectRuleIndexMetadata) GetSucceedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Succeed, true
}

// SetSucceed sets field value
func (o *RedirectRuleIndexMetadata) SetSucceed(v bool) {
	o.Succeed = v
}

// GetData returns the Data field value
func (o *RedirectRuleIndexMetadata) GetData() RedirectRuleIndexMetadataData {
	if o == nil {
		var ret RedirectRuleIndexMetadataData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *RedirectRuleIndexMetadata) GetDataOk() (*RedirectRuleIndexMetadataData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *RedirectRuleIndexMetadata) SetData(v RedirectRuleIndexMetadataData) {
	o.Data = v
}

func (o RedirectRuleIndexMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["dest"] = o.Dest
	}
	if true {
		toSerialize["reason"] = o.Reason
	}
	if true {
		toSerialize["succeed"] = o.Succeed
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

func (o RedirectRuleIndexMetadata) String() string {
	out := ""
	out += fmt.Sprintf("  source=%v\n", o.Source)
	out += fmt.Sprintf("  dest=%v\n", o.Dest)
	out += fmt.Sprintf("  reason=%v\n", o.Reason)
	out += fmt.Sprintf("  succeed=%v\n", o.Succeed)
	out += fmt.Sprintf("  data=%v\n", o.Data)
	return fmt.Sprintf("RedirectRuleIndexMetadata {\n%s}", out)
}

type NullableRedirectRuleIndexMetadata struct {
	value *RedirectRuleIndexMetadata
	isSet bool
}

func (v NullableRedirectRuleIndexMetadata) Get() *RedirectRuleIndexMetadata {
	return v.value
}

func (v *NullableRedirectRuleIndexMetadata) Set(val *RedirectRuleIndexMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableRedirectRuleIndexMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableRedirectRuleIndexMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedirectRuleIndexMetadata(val *RedirectRuleIndexMetadata) *NullableRedirectRuleIndexMetadata {
	return &NullableRedirectRuleIndexMetadata{value: val, isSet: true}
}

func (v NullableRedirectRuleIndexMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedirectRuleIndexMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}