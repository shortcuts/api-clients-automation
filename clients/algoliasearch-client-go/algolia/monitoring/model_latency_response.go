// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package monitoring

import (
	"encoding/json"
	"fmt"
)

// LatencyResponse struct for LatencyResponse
type LatencyResponse struct {
	Metrics *LatencyResponseMetrics `json:"metrics,omitempty"`
}

type LatencyResponseOption func(f *LatencyResponse)

func WithLatencyResponseMetrics(val LatencyResponseMetrics) LatencyResponseOption {
	return func(f *LatencyResponse) {
		f.Metrics = &val
	}
}

// NewLatencyResponse instantiates a new LatencyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLatencyResponse(opts ...LatencyResponseOption) *LatencyResponse {
	this := &LatencyResponse{}
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewLatencyResponseWithDefaults instantiates a new LatencyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLatencyResponseWithDefaults() *LatencyResponse {
	this := &LatencyResponse{}
	return this
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *LatencyResponse) GetMetrics() LatencyResponseMetrics {
	if o == nil || o.Metrics == nil {
		var ret LatencyResponseMetrics
		return ret
	}
	return *o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LatencyResponse) GetMetricsOk() (*LatencyResponseMetrics, bool) {
	if o == nil || o.Metrics == nil {
		return nil, false
	}
	return o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *LatencyResponse) HasMetrics() bool {
	if o != nil && o.Metrics != nil {
		return true
	}

	return false
}

// SetMetrics gets a reference to the given LatencyResponseMetrics and assigns it to the Metrics field.
func (o *LatencyResponse) SetMetrics(v LatencyResponseMetrics) {
	o.Metrics = &v
}

func (o LatencyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Metrics != nil {
		toSerialize["metrics"] = o.Metrics
	}
	return json.Marshal(toSerialize)
}

func (o LatencyResponse) String() string {
	out := ""
	out += fmt.Sprintf("  metrics=%v\n", o.Metrics)
	return fmt.Sprintf("LatencyResponse {\n%s}", out)
}

type NullableLatencyResponse struct {
	value *LatencyResponse
	isSet bool
}

func (v NullableLatencyResponse) Get() *LatencyResponse {
	return v.value
}

func (v *NullableLatencyResponse) Set(val *LatencyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLatencyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLatencyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLatencyResponse(val *LatencyResponse) *NullableLatencyResponse {
	return &NullableLatencyResponse{value: val, isSet: true}
}

func (v NullableLatencyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLatencyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}