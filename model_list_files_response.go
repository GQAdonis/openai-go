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

// ListFilesResponse struct for ListFilesResponse
type ListFilesResponse struct {
	Object string `json:"object"`
	Data []OpenAIFile `json:"data"`
}

// NewListFilesResponse instantiates a new ListFilesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListFilesResponse(object string, data []OpenAIFile) *ListFilesResponse {
	this := ListFilesResponse{}
	this.Object = object
	this.Data = data
	return &this
}

// NewListFilesResponseWithDefaults instantiates a new ListFilesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListFilesResponseWithDefaults() *ListFilesResponse {
	this := ListFilesResponse{}
	return &this
}

// GetObject returns the Object field value
func (o *ListFilesResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *ListFilesResponse) GetObjectOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *ListFilesResponse) SetObject(v string) {
	o.Object = v
}

// GetData returns the Data field value
func (o *ListFilesResponse) GetData() []OpenAIFile {
	if o == nil {
		var ret []OpenAIFile
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListFilesResponse) GetDataOk() ([]OpenAIFile, bool) {
	if o == nil {
    return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListFilesResponse) SetData(v []OpenAIFile) {
	o.Data = v
}

func (o ListFilesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["object"] = o.Object
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableListFilesResponse struct {
	value *ListFilesResponse
	isSet bool
}

func (v NullableListFilesResponse) Get() *ListFilesResponse {
	return v.value
}

func (v *NullableListFilesResponse) Set(val *ListFilesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListFilesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListFilesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListFilesResponse(val *ListFilesResponse) *NullableListFilesResponse {
	return &NullableListFilesResponse{value: val, isSet: true}
}

func (v NullableListFilesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListFilesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

