/*
OpenAI API

APIs for sampling from and fine-tuning language models

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi-go

import (
	"encoding/json"
)

// CreateModerationRequest struct for CreateModerationRequest
type CreateModerationRequest struct {
	Input CreateModerationRequestInput `json:"input"`
	// Two content moderations models are available: `text-moderation-stable` and `text-moderation-latest`.  The default is `text-moderation-latest` which will be automatically upgraded over time. This ensures you are always using our most accurate model. If you use `text-moderation-stable`, we will provide advanced notice before updating the model. Accuracy of `text-moderation-stable` may be slightly lower than for `text-moderation-latest`. 
	Model *string `json:"model,omitempty"`
}

// NewCreateModerationRequest instantiates a new CreateModerationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateModerationRequest(input CreateModerationRequestInput) *CreateModerationRequest {
	this := CreateModerationRequest{}
	this.Input = input
	var model string = "text-moderation-latest"
	this.Model = &model
	return &this
}

// NewCreateModerationRequestWithDefaults instantiates a new CreateModerationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateModerationRequestWithDefaults() *CreateModerationRequest {
	this := CreateModerationRequest{}
	var model string = "text-moderation-latest"
	this.Model = &model
	return &this
}

// GetInput returns the Input field value
func (o *CreateModerationRequest) GetInput() CreateModerationRequestInput {
	if o == nil {
		var ret CreateModerationRequestInput
		return ret
	}

	return o.Input
}

// GetInputOk returns a tuple with the Input field value
// and a boolean to check if the value has been set.
func (o *CreateModerationRequest) GetInputOk() (*CreateModerationRequestInput, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Input, true
}

// SetInput sets field value
func (o *CreateModerationRequest) SetInput(v CreateModerationRequestInput) {
	o.Input = v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *CreateModerationRequest) GetModel() string {
	if o == nil || isNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateModerationRequest) GetModelOk() (*string, bool) {
	if o == nil || isNil(o.Model) {
    return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *CreateModerationRequest) HasModel() bool {
	if o != nil && !isNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *CreateModerationRequest) SetModel(v string) {
	o.Model = &v
}

func (o CreateModerationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["input"] = o.Input
	}
	if !isNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	return json.Marshal(toSerialize)
}

type NullableCreateModerationRequest struct {
	value *CreateModerationRequest
	isSet bool
}

func (v NullableCreateModerationRequest) Get() *CreateModerationRequest {
	return v.value
}

func (v *NullableCreateModerationRequest) Set(val *CreateModerationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateModerationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateModerationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateModerationRequest(val *CreateModerationRequest) *NullableCreateModerationRequest {
	return &NullableCreateModerationRequest{value: val, isSet: true}
}

func (v NullableCreateModerationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateModerationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


