/*
OpenAI API

APIs for sampling from and fine-tuning language models

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi-go

import (
	"encoding/json"
	"fmt"
)

// CreateEmbeddingRequestInput - Input text to get embeddings for, encoded as a string or array of tokens. To get embeddings for multiple inputs in a single request, pass an array of strings or array of token arrays. Each input must not exceed 8192 tokens in length. 
type CreateEmbeddingRequestInput struct {
	ArrayOfArrayOfInt32 *[][]int32
	ArrayOfInt32 *[]int32
	ArrayOfString *[]string
	String *string
}

// [][]int32AsCreateEmbeddingRequestInput is a convenience function that returns [][]int32 wrapped in CreateEmbeddingRequestInput
func ArrayOfArrayOfInt32AsCreateEmbeddingRequestInput(v *[][]int32) CreateEmbeddingRequestInput {
	return CreateEmbeddingRequestInput{
		ArrayOfArrayOfInt32: v,
	}
}

// []int32AsCreateEmbeddingRequestInput is a convenience function that returns []int32 wrapped in CreateEmbeddingRequestInput
func ArrayOfInt32AsCreateEmbeddingRequestInput(v *[]int32) CreateEmbeddingRequestInput {
	return CreateEmbeddingRequestInput{
		ArrayOfInt32: v,
	}
}

// []stringAsCreateEmbeddingRequestInput is a convenience function that returns []string wrapped in CreateEmbeddingRequestInput
func ArrayOfStringAsCreateEmbeddingRequestInput(v *[]string) CreateEmbeddingRequestInput {
	return CreateEmbeddingRequestInput{
		ArrayOfString: v,
	}
}

// stringAsCreateEmbeddingRequestInput is a convenience function that returns string wrapped in CreateEmbeddingRequestInput
func StringAsCreateEmbeddingRequestInput(v *string) CreateEmbeddingRequestInput {
	return CreateEmbeddingRequestInput{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateEmbeddingRequestInput) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ArrayOfArrayOfInt32
	err = newStrictDecoder(data).Decode(&dst.ArrayOfArrayOfInt32)
	if err == nil {
		jsonArrayOfArrayOfInt32, _ := json.Marshal(dst.ArrayOfArrayOfInt32)
		if string(jsonArrayOfArrayOfInt32) == "{}" { // empty struct
			dst.ArrayOfArrayOfInt32 = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfArrayOfInt32 = nil
	}

	// try to unmarshal data into ArrayOfInt32
	err = newStrictDecoder(data).Decode(&dst.ArrayOfInt32)
	if err == nil {
		jsonArrayOfInt32, _ := json.Marshal(dst.ArrayOfInt32)
		if string(jsonArrayOfInt32) == "{}" { // empty struct
			dst.ArrayOfInt32 = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfInt32 = nil
	}

	// try to unmarshal data into ArrayOfString
	err = newStrictDecoder(data).Decode(&dst.ArrayOfString)
	if err == nil {
		jsonArrayOfString, _ := json.Marshal(dst.ArrayOfString)
		if string(jsonArrayOfString) == "{}" { // empty struct
			dst.ArrayOfString = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfString = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ArrayOfArrayOfInt32 = nil
		dst.ArrayOfInt32 = nil
		dst.ArrayOfString = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateEmbeddingRequestInput)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateEmbeddingRequestInput)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateEmbeddingRequestInput) MarshalJSON() ([]byte, error) {
	if src.ArrayOfArrayOfInt32 != nil {
		return json.Marshal(&src.ArrayOfArrayOfInt32)
	}

	if src.ArrayOfInt32 != nil {
		return json.Marshal(&src.ArrayOfInt32)
	}

	if src.ArrayOfString != nil {
		return json.Marshal(&src.ArrayOfString)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateEmbeddingRequestInput) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ArrayOfArrayOfInt32 != nil {
		return obj.ArrayOfArrayOfInt32
	}

	if obj.ArrayOfInt32 != nil {
		return obj.ArrayOfInt32
	}

	if obj.ArrayOfString != nil {
		return obj.ArrayOfString
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableCreateEmbeddingRequestInput struct {
	value *CreateEmbeddingRequestInput
	isSet bool
}

func (v NullableCreateEmbeddingRequestInput) Get() *CreateEmbeddingRequestInput {
	return v.value
}

func (v *NullableCreateEmbeddingRequestInput) Set(val *CreateEmbeddingRequestInput) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateEmbeddingRequestInput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateEmbeddingRequestInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateEmbeddingRequestInput(val *CreateEmbeddingRequestInput) *NullableCreateEmbeddingRequestInput {
	return &NullableCreateEmbeddingRequestInput{value: val, isSet: true}
}

func (v NullableCreateEmbeddingRequestInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateEmbeddingRequestInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


