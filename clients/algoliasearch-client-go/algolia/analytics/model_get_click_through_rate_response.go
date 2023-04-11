// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package analytics

import (
	"encoding/json"
	"fmt"
)

// GetClickThroughRateResponse struct for GetClickThroughRateResponse
type GetClickThroughRateResponse struct {
	// The click-through rate.
	Rate float64 `json:"rate"`
	// The number of click event.
	ClickCount int32 `json:"clickCount"`
	// The number of tracked search click.
	TrackedSearchCount int32 `json:"trackedSearchCount"`
	// A list of click-through rate events with their date.
	Dates []ClickThroughRateEvent `json:"dates"`
}

// NewGetClickThroughRateResponse instantiates a new GetClickThroughRateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetClickThroughRateResponse(rate float64, clickCount int32, trackedSearchCount int32, dates []ClickThroughRateEvent) *GetClickThroughRateResponse {
	this := &GetClickThroughRateResponse{}
	this.Rate = rate
	this.ClickCount = clickCount
	this.TrackedSearchCount = trackedSearchCount
	this.Dates = dates
	return this
}

// NewGetClickThroughRateResponseWithDefaults instantiates a new GetClickThroughRateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetClickThroughRateResponseWithDefaults() *GetClickThroughRateResponse {
	this := &GetClickThroughRateResponse{}
	return this
}

// GetRate returns the Rate field value
func (o *GetClickThroughRateResponse) GetRate() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Rate
}

// GetRateOk returns a tuple with the Rate field value
// and a boolean to check if the value has been set.
func (o *GetClickThroughRateResponse) GetRateOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rate, true
}

// SetRate sets field value
func (o *GetClickThroughRateResponse) SetRate(v float64) {
	o.Rate = v
}

// GetClickCount returns the ClickCount field value
func (o *GetClickThroughRateResponse) GetClickCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ClickCount
}

// GetClickCountOk returns a tuple with the ClickCount field value
// and a boolean to check if the value has been set.
func (o *GetClickThroughRateResponse) GetClickCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClickCount, true
}

// SetClickCount sets field value
func (o *GetClickThroughRateResponse) SetClickCount(v int32) {
	o.ClickCount = v
}

// GetTrackedSearchCount returns the TrackedSearchCount field value
func (o *GetClickThroughRateResponse) GetTrackedSearchCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TrackedSearchCount
}

// GetTrackedSearchCountOk returns a tuple with the TrackedSearchCount field value
// and a boolean to check if the value has been set.
func (o *GetClickThroughRateResponse) GetTrackedSearchCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrackedSearchCount, true
}

// SetTrackedSearchCount sets field value
func (o *GetClickThroughRateResponse) SetTrackedSearchCount(v int32) {
	o.TrackedSearchCount = v
}

// GetDates returns the Dates field value
func (o *GetClickThroughRateResponse) GetDates() []ClickThroughRateEvent {
	if o == nil {
		var ret []ClickThroughRateEvent
		return ret
	}

	return o.Dates
}

// GetDatesOk returns a tuple with the Dates field value
// and a boolean to check if the value has been set.
func (o *GetClickThroughRateResponse) GetDatesOk() ([]ClickThroughRateEvent, bool) {
	if o == nil {
		return nil, false
	}
	return o.Dates, true
}

// SetDates sets field value
func (o *GetClickThroughRateResponse) SetDates(v []ClickThroughRateEvent) {
	o.Dates = v
}

func (o GetClickThroughRateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["rate"] = o.Rate
	}
	if true {
		toSerialize["clickCount"] = o.ClickCount
	}
	if true {
		toSerialize["trackedSearchCount"] = o.TrackedSearchCount
	}
	if true {
		toSerialize["dates"] = o.Dates
	}
	return json.Marshal(toSerialize)
}

func (o GetClickThroughRateResponse) String() string {
	out := ""
	out += fmt.Sprintf("  rate=%v\n", o.Rate)
	out += fmt.Sprintf("  clickCount=%v\n", o.ClickCount)
	out += fmt.Sprintf("  trackedSearchCount=%v\n", o.TrackedSearchCount)
	out += fmt.Sprintf("  dates=%v\n", o.Dates)
	return fmt.Sprintf("GetClickThroughRateResponse {\n%s}", out)
}

type NullableGetClickThroughRateResponse struct {
	value *GetClickThroughRateResponse
	isSet bool
}

func (v NullableGetClickThroughRateResponse) Get() *GetClickThroughRateResponse {
	return v.value
}

func (v *NullableGetClickThroughRateResponse) Set(val *GetClickThroughRateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetClickThroughRateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetClickThroughRateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetClickThroughRateResponse(val *GetClickThroughRateResponse) *NullableGetClickThroughRateResponse {
	return &NullableGetClickThroughRateResponse{value: val, isSet: true}
}

func (v NullableGetClickThroughRateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetClickThroughRateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}