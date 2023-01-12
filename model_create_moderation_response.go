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

// CreateModerationResponse struct for CreateModerationResponse
type CreateModerationResponse struct {
	Id string `json:"id"`
	Model string `json:"model"`
	Results []CreateModerationResponseResultsInner `json:"results"`
}

// NewCreateModerationResponse instantiates a new CreateModerationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateModerationResponse(id string, model string, results []CreateModerationResponseResultsInner) *CreateModerationResponse {
	this := CreateModerationResponse{}
	this.Id = id
	this.Model = model
	this.Results = results
	return &this
}

// NewCreateModerationResponseWithDefaults instantiates a new CreateModerationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateModerationResponseWithDefaults() *CreateModerationResponse {
	this := CreateModerationResponse{}
	return &this
}

// GetId returns the Id field value
func (o *CreateModerationResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreateModerationResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CreateModerationResponse) SetId(v string) {
	o.Id = v
}

// GetModel returns the Model field value
func (o *CreateModerationResponse) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *CreateModerationResponse) GetModelOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *CreateModerationResponse) SetModel(v string) {
	o.Model = v
}

// GetResults returns the Results field value
func (o *CreateModerationResponse) GetResults() []CreateModerationResponseResultsInner {
	if o == nil {
		var ret []CreateModerationResponseResultsInner
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *CreateModerationResponse) GetResultsOk() ([]CreateModerationResponseResultsInner, bool) {
	if o == nil {
    return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *CreateModerationResponse) SetResults(v []CreateModerationResponseResultsInner) {
	o.Results = v
}

func (o CreateModerationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["model"] = o.Model
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableCreateModerationResponse struct {
	value *CreateModerationResponse
	isSet bool
}

func (v NullableCreateModerationResponse) Get() *CreateModerationResponse {
	return v.value
}

func (v *NullableCreateModerationResponse) Set(val *CreateModerationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateModerationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateModerationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateModerationResponse(val *CreateModerationResponse) *NullableCreateModerationResponse {
	return &NullableCreateModerationResponse{value: val, isSet: true}
}

func (v NullableCreateModerationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateModerationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

