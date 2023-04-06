// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package predict

import (
	"encoding/json"
	"fmt"
)

// DeleteUserProfileResponse struct for DeleteUserProfileResponse
type DeleteUserProfileResponse struct {
	// The ID of the user that was deleted.
	User string `json:"user"`
	// The time the same user ID will be imported again when the data is ingested.
	DeletedUntil string `json:"deletedUntil"`
}

// NewDeleteUserProfileResponse instantiates a new DeleteUserProfileResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteUserProfileResponse(user string, deletedUntil string) *DeleteUserProfileResponse {
	this := &DeleteUserProfileResponse{}
	this.User = user
	this.DeletedUntil = deletedUntil
	return this
}

// NewDeleteUserProfileResponseWithDefaults instantiates a new DeleteUserProfileResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteUserProfileResponseWithDefaults() *DeleteUserProfileResponse {
	this := &DeleteUserProfileResponse{}
	return this
}

// GetUser returns the User field value
func (o *DeleteUserProfileResponse) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *DeleteUserProfileResponse) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *DeleteUserProfileResponse) SetUser(v string) {
	o.User = v
}

// GetDeletedUntil returns the DeletedUntil field value
func (o *DeleteUserProfileResponse) GetDeletedUntil() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeletedUntil
}

// GetDeletedUntilOk returns a tuple with the DeletedUntil field value
// and a boolean to check if the value has been set.
func (o *DeleteUserProfileResponse) GetDeletedUntilOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeletedUntil, true
}

// SetDeletedUntil sets field value
func (o *DeleteUserProfileResponse) SetDeletedUntil(v string) {
	o.DeletedUntil = v
}

func (o DeleteUserProfileResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["user"] = o.User
	}
	if true {
		toSerialize["deletedUntil"] = o.DeletedUntil
	}
	return json.Marshal(toSerialize)
}

func (o DeleteUserProfileResponse) String() string {
	out := ""
	out += fmt.Sprintf("  user=%v\n", o.User)
	out += fmt.Sprintf("  deletedUntil=%v\n", o.DeletedUntil)
	return fmt.Sprintf("DeleteUserProfileResponse {\n%s}", out)
}

type NullableDeleteUserProfileResponse struct {
	value *DeleteUserProfileResponse
	isSet bool
}

func (v NullableDeleteUserProfileResponse) Get() *DeleteUserProfileResponse {
	return v.value
}

func (v *NullableDeleteUserProfileResponse) Set(val *DeleteUserProfileResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteUserProfileResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteUserProfileResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteUserProfileResponse(val *DeleteUserProfileResponse) *NullableDeleteUserProfileResponse {
	return &NullableDeleteUserProfileResponse{value: val, isSet: true}
}

func (v NullableDeleteUserProfileResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteUserProfileResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
