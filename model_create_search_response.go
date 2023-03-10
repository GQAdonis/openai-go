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

// CreateSearchResponse struct for CreateSearchResponse
type CreateSearchResponse struct {
	Object *string `json:"object,omitempty"`
	Model *string `json:"model,omitempty"`
	Data []CreateSearchResponseDataInner `json:"data,omitempty"`
}

// NewCreateSearchResponse instantiates a new CreateSearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSearchResponse() *CreateSearchResponse {
	this := CreateSearchResponse{}
	return &this
}

// NewCreateSearchResponseWithDefaults instantiates a new CreateSearchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSearchResponseWithDefaults() *CreateSearchResponse {
	this := CreateSearchResponse{}
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *CreateSearchResponse) GetObject() string {
	if o == nil || isNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSearchResponse) GetObjectOk() (*string, bool) {
	if o == nil || isNil(o.Object) {
    return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *CreateSearchResponse) HasObject() bool {
	if o != nil && !isNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *CreateSearchResponse) SetObject(v string) {
	o.Object = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *CreateSearchResponse) GetModel() string {
	if o == nil || isNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSearchResponse) GetModelOk() (*string, bool) {
	if o == nil || isNil(o.Model) {
    return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *CreateSearchResponse) HasModel() bool {
	if o != nil && !isNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *CreateSearchResponse) SetModel(v string) {
	o.Model = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateSearchResponse) GetData() []CreateSearchResponseDataInner {
	if o == nil || isNil(o.Data) {
		var ret []CreateSearchResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSearchResponse) GetDataOk() ([]CreateSearchResponseDataInner, bool) {
	if o == nil || isNil(o.Data) {
    return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateSearchResponse) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []CreateSearchResponseDataInner and assigns it to the Data field.
func (o *CreateSearchResponse) SetData(v []CreateSearchResponseDataInner) {
	o.Data = v
}

func (o CreateSearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if !isNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCreateSearchResponse struct {
	value *CreateSearchResponse
	isSet bool
}

func (v NullableCreateSearchResponse) Get() *CreateSearchResponse {
	return v.value
}

func (v *NullableCreateSearchResponse) Set(val *CreateSearchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSearchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSearchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSearchResponse(val *CreateSearchResponse) *NullableCreateSearchResponse {
	return &NullableCreateSearchResponse{value: val, isSet: true}
}

func (v NullableCreateSearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSearchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


