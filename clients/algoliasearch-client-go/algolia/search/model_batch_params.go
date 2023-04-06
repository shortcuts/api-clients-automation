// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"
)

// BatchParams The `multipleBatch` parameters.
type BatchParams struct {
	Requests []MultipleBatchRequest `json:"requests"`
}

// NewBatchParams instantiates a new BatchParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchParams(requests []MultipleBatchRequest) *BatchParams {
	this := &BatchParams{}
	this.Requests = requests
	return this
}

// NewBatchParamsWithDefaults instantiates a new BatchParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchParamsWithDefaults() *BatchParams {
	this := &BatchParams{}
	return this
}

// GetRequests returns the Requests field value
func (o *BatchParams) GetRequests() []MultipleBatchRequest {
	if o == nil {
		var ret []MultipleBatchRequest
		return ret
	}

	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value
// and a boolean to check if the value has been set.
func (o *BatchParams) GetRequestsOk() ([]MultipleBatchRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Requests, true
}

// SetRequests sets field value
func (o *BatchParams) SetRequests(v []MultipleBatchRequest) {
	o.Requests = v
}

func (o BatchParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["requests"] = o.Requests
	}
	return json.Marshal(toSerialize)
}

func (o BatchParams) String() string {
	out := ""
	out += fmt.Sprintf("  requests=%v\n", o.Requests)
	return fmt.Sprintf("BatchParams {\n%s}", out)
}

type NullableBatchParams struct {
	value *BatchParams
	isSet bool
}

func (v NullableBatchParams) Get() *BatchParams {
	return v.value
}

func (v *NullableBatchParams) Set(val *BatchParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchParams(val *BatchParams) *NullableBatchParams {
	return &NullableBatchParams{value: val, isSet: true}
}

func (v NullableBatchParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
