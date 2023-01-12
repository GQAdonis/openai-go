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

// ListEnginesResponse struct for ListEnginesResponse
type ListEnginesResponse struct {
	Object string `json:"object"`
	Data []Engine `json:"data"`
}

// NewListEnginesResponse instantiates a new ListEnginesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListEnginesResponse(object string, data []Engine) *ListEnginesResponse {
	this := ListEnginesResponse{}
	this.Object = object
	this.Data = data
	return &this
}

// NewListEnginesResponseWithDefaults instantiates a new ListEnginesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListEnginesResponseWithDefaults() *ListEnginesResponse {
	this := ListEnginesResponse{}
	return &this
}

// GetObject returns the Object field value
func (o *ListEnginesResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *ListEnginesResponse) GetObjectOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *ListEnginesResponse) SetObject(v string) {
	o.Object = v
}

// GetData returns the Data field value
func (o *ListEnginesResponse) GetData() []Engine {
	if o == nil {
		var ret []Engine
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListEnginesResponse) GetDataOk() ([]Engine, bool) {
	if o == nil {
    return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListEnginesResponse) SetData(v []Engine) {
	o.Data = v
}

func (o ListEnginesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["object"] = o.Object
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableListEnginesResponse struct {
	value *ListEnginesResponse
	isSet bool
}

func (v NullableListEnginesResponse) Get() *ListEnginesResponse {
	return v.value
}

func (v *NullableListEnginesResponse) Set(val *ListEnginesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListEnginesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListEnginesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListEnginesResponse(val *ListEnginesResponse) *NullableListEnginesResponse {
	return &NullableListEnginesResponse{value: val, isSet: true}
}

func (v NullableListEnginesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListEnginesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


