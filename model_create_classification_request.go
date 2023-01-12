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

// CreateClassificationRequest struct for CreateClassificationRequest
type CreateClassificationRequest struct {
	// ID of the model to use. You can use the [List models](/docs/api-reference/models/list) API to see all of your available models, or see our [Model overview](/docs/models/overview) for descriptions of them.
	Model string `json:"model"`
	// Query to be classified.
	Query string `json:"query"`
	// A list of examples with labels, in the following format:  `[[\"The movie is so interesting.\", \"Positive\"], [\"It is quite boring.\", \"Negative\"], ...]`  All the label strings will be normalized to be capitalized.  You should specify either `examples` or `file`, but not both. 
	Examples [][]string `json:"examples,omitempty"`
	// The ID of the uploaded file that contains training examples. See [upload file](/docs/api-reference/files/upload) for how to upload a file of the desired format and purpose.  You should specify either `examples` or `file`, but not both. 
	File NullableString `json:"file,omitempty"`
	// The set of categories being classified. If not specified, candidate labels will be automatically collected from the examples you provide. All the label strings will be normalized to be capitalized.
	Labels []string `json:"labels,omitempty"`
	// ID of the model to use for [Search](/docs/api-reference/searches/create). You can select one of `ada`, `babbage`, `curie`, or `davinci`.
	SearchModel NullableString `json:"search_model,omitempty"`
	// What sampling `temperature` to use. Higher values mean the model will take more risks. Try 0.9 for more creative applications, and 0 (argmax sampling) for ones with a well-defined answer.
	Temperature NullableFloat32 `json:"temperature,omitempty"`
	// Include the log probabilities on the `logprobs` most likely tokens, as well the chosen tokens. For example, if `logprobs` is 5, the API will return a list of the 5 most likely tokens. The API will always return the `logprob` of the sampled token, so there may be up to `logprobs+1` elements in the response.  The maximum value for `logprobs` is 5. If you need more than this, please contact us through our [Help center](https://help.openai.com) and describe your use case.  When `logprobs` is set, `completion` will be automatically added into `expand` to get the logprobs. 
	Logprobs NullableInt32 `json:"logprobs,omitempty"`
	// The maximum number of examples to be ranked by [Search](/docs/api-reference/searches/create) when using `file`. Setting it to a higher value leads to improved accuracy but with increased latency and cost.
	MaxExamples NullableInt32 `json:"max_examples,omitempty"`
	// Modify the likelihood of specified tokens appearing in the completion.  Accepts a json object that maps tokens (specified by their token ID in the GPT tokenizer) to an associated bias value from -100 to 100. You can use this [tokenizer tool](/tokenizer?view=bpe) (which works for both GPT-2 and GPT-3) to convert text to token IDs. Mathematically, the bias is added to the logits generated by the model prior to sampling. The exact effect will vary per model, but values between -1 and 1 should decrease or increase likelihood of selection; values like -100 or 100 should result in a ban or exclusive selection of the relevant token.  As an example, you can pass `{\"50256\": -100}` to prevent the <|endoftext|> token from being generated. 
	LogitBias map[string]interface{} `json:"logit_bias,omitempty"`
	// If set to `true`, the returned JSON will include a \"prompt\" field containing the final prompt that was used to request a completion. This is mainly useful for debugging purposes.
	ReturnPrompt NullableBool `json:"return_prompt,omitempty"`
	// A special boolean flag for showing metadata. If set to `true`, each document entry in the returned JSON will contain a \"metadata\" field.  This flag only takes effect when `file` is set. 
	ReturnMetadata NullableBool `json:"return_metadata,omitempty"`
	// If an object name is in the list, we provide the full information of the object; otherwise, we only provide the object ID. Currently we support `completion` and `file` objects for expansion.
	Expand []interface{} `json:"expand,omitempty"`
	// A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse. [Learn more](/docs/guides/safety-best-practices/end-user-ids). 
	User *string `json:"user,omitempty"`
}

// NewCreateClassificationRequest instantiates a new CreateClassificationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateClassificationRequest(model string, query string) *CreateClassificationRequest {
	this := CreateClassificationRequest{}
	this.Model = model
	this.Query = query
	var searchModel string = "ada"
	this.SearchModel = *NewNullableString(&searchModel)
	var temperature float32 = 0
	this.Temperature = *NewNullableFloat32(&temperature)
	var maxExamples int32 = 200
	this.MaxExamples = *NewNullableInt32(&maxExamples)
	var returnPrompt bool = false
	this.ReturnPrompt = *NewNullableBool(&returnPrompt)
	var returnMetadata bool = false
	this.ReturnMetadata = *NewNullableBool(&returnMetadata)
	return &this
}

// NewCreateClassificationRequestWithDefaults instantiates a new CreateClassificationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateClassificationRequestWithDefaults() *CreateClassificationRequest {
	this := CreateClassificationRequest{}
	var searchModel string = "ada"
	this.SearchModel = *NewNullableString(&searchModel)
	var temperature float32 = 0
	this.Temperature = *NewNullableFloat32(&temperature)
	var maxExamples int32 = 200
	this.MaxExamples = *NewNullableInt32(&maxExamples)
	var returnPrompt bool = false
	this.ReturnPrompt = *NewNullableBool(&returnPrompt)
	var returnMetadata bool = false
	this.ReturnMetadata = *NewNullableBool(&returnMetadata)
	return &this
}

// GetModel returns the Model field value
func (o *CreateClassificationRequest) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *CreateClassificationRequest) GetModelOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *CreateClassificationRequest) SetModel(v string) {
	o.Model = v
}

// GetQuery returns the Query field value
func (o *CreateClassificationRequest) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *CreateClassificationRequest) GetQueryOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *CreateClassificationRequest) SetQuery(v string) {
	o.Query = v
}

// GetExamples returns the Examples field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateClassificationRequest) GetExamples() [][]string {
	if o == nil {
		var ret [][]string
		return ret
	}
	return o.Examples
}

// GetExamplesOk returns a tuple with the Examples field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateClassificationRequest) GetExamplesOk() ([][]string, bool) {
	if o == nil || isNil(o.Examples) {
    return nil, false
	}
	return o.Examples, true
}

// HasExamples returns a boolean if a field has been set.
func (o *CreateClassificationRequest) HasExamples() bool {
	if o != nil && isNil(o.Examples) {
		return true
	}

	return false
}

// SetExamples gets a reference to the given [][]string and assigns it to the Examples field.
func (o *CreateClassificationRequest) SetExamples(v [][]string) {
	o.Examples = v
}

// GetFile returns the File field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateClassificationRequest) GetFile() string {
	if o == nil || isNil(o.File.Get()) {
		var ret string
		return ret
	}
	return *o.File.Get()
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateClassificationRequest) GetFileOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.File.Get(), o.File.IsSet()
}

// HasFile returns a boolean if a field has been set.
func (o *CreateClassificationRequest) HasFile() bool {
	if o != nil && o.File.IsSet() {
		return true
	}

	return false
}

// SetFile gets a reference to the given NullableString and assigns it to the File field.
func (o *CreateClassificationRequest) SetFile(v string) {
	o.File.Set(&v)
}
// SetFileNil sets the value for File to be an explicit nil
func (o *CreateClassificationRequest) SetFileNil() {
	o.File.Set(nil)
}

// UnsetFile ensures that no value is present for File, not even an explicit nil
func (o *CreateClassificationRequest) UnsetFile() {
	o.File.Unset()
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateClassificationRequest) GetLabels() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateClassificationRequest) GetLabelsOk() ([]string, bool) {
	if o == nil || isNil(o.Labels) {
    return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *CreateClassificationRequest) HasLabels() bool {
	if o != nil && isNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *CreateClassificationRequest) SetLabels(v []string) {
	o.Labels = v
}

// GetSearchModel returns the SearchModel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateClassificationRequest) GetSearchModel() string {
	if o == nil || isNil(o.SearchModel.Get()) {
		var ret string
		return ret
	}
	return *o.SearchModel.Get()
}

// GetSearchModelOk returns a tuple with the SearchModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateClassificationRequest) GetSearchModelOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.SearchModel.Get(), o.SearchModel.IsSet()
}

// HasSearchModel returns a boolean if a field has been set.
func (o *CreateClassificationRequest) HasSearchModel() bool {
	if o != nil && o.SearchModel.IsSet() {
		return true
	}

	return false
}

// SetSearchModel gets a reference to the given NullableString and assigns it to the SearchModel field.
func (o *CreateClassificationRequest) SetSearchModel(v string) {
	o.SearchModel.Set(&v)
}
// SetSearchModelNil sets the value for SearchModel to be an explicit nil
func (o *CreateClassificationRequest) SetSearchModelNil() {
	o.SearchModel.Set(nil)
}

// UnsetSearchModel ensures that no value is present for SearchModel, not even an explicit nil
func (o *CreateClassificationRequest) UnsetSearchModel() {
	o.SearchModel.Unset()
}

// GetTemperature returns the Temperature field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateClassificationRequest) GetTemperature() float32 {
	if o == nil || isNil(o.Temperature.Get()) {
		var ret float32
		return ret
	}
	return *o.Temperature.Get()
}

// GetTemperatureOk returns a tuple with the Temperature field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateClassificationRequest) GetTemperatureOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return o.Temperature.Get(), o.Temperature.IsSet()
}

// HasTemperature returns a boolean if a field has been set.
func (o *CreateClassificationRequest) HasTemperature() bool {
	if o != nil && o.Temperature.IsSet() {
		return true
	}

	return false
}

// SetTemperature gets a reference to the given NullableFloat32 and assigns it to the Temperature field.
func (o *CreateClassificationRequest) SetTemperature(v float32) {
	o.Temperature.Set(&v)
}
// SetTemperatureNil sets the value for Temperature to be an explicit nil
func (o *CreateClassificationRequest) SetTemperatureNil() {
	o.Temperature.Set(nil)
}

// UnsetTemperature ensures that no value is present for Temperature, not even an explicit nil
func (o *CreateClassificationRequest) UnsetTemperature() {
	o.Temperature.Unset()
}

// GetLogprobs returns the Logprobs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateClassificationRequest) GetLogprobs() int32 {
	if o == nil || isNil(o.Logprobs.Get()) {
		var ret int32
		return ret
	}
	return *o.Logprobs.Get()
}

// GetLogprobsOk returns a tuple with the Logprobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateClassificationRequest) GetLogprobsOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.Logprobs.Get(), o.Logprobs.IsSet()
}

// HasLogprobs returns a boolean if a field has been set.
func (o *CreateClassificationRequest) HasLogprobs() bool {
	if o != nil && o.Logprobs.IsSet() {
		return true
	}

	return false
}

// SetLogprobs gets a reference to the given NullableInt32 and assigns it to the Logprobs field.
func (o *CreateClassificationRequest) SetLogprobs(v int32) {
	o.Logprobs.Set(&v)
}
// SetLogprobsNil sets the value for Logprobs to be an explicit nil
func (o *CreateClassificationRequest) SetLogprobsNil() {
	o.Logprobs.Set(nil)
}

// UnsetLogprobs ensures that no value is present for Logprobs, not even an explicit nil
func (o *CreateClassificationRequest) UnsetLogprobs() {
	o.Logprobs.Unset()
}

// GetMaxExamples returns the MaxExamples field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateClassificationRequest) GetMaxExamples() int32 {
	if o == nil || isNil(o.MaxExamples.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxExamples.Get()
}

// GetMaxExamplesOk returns a tuple with the MaxExamples field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateClassificationRequest) GetMaxExamplesOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.MaxExamples.Get(), o.MaxExamples.IsSet()
}

// HasMaxExamples returns a boolean if a field has been set.
func (o *CreateClassificationRequest) HasMaxExamples() bool {
	if o != nil && o.MaxExamples.IsSet() {
		return true
	}

	return false
}

// SetMaxExamples gets a reference to the given NullableInt32 and assigns it to the MaxExamples field.
func (o *CreateClassificationRequest) SetMaxExamples(v int32) {
	o.MaxExamples.Set(&v)
}
// SetMaxExamplesNil sets the value for MaxExamples to be an explicit nil
func (o *CreateClassificationRequest) SetMaxExamplesNil() {
	o.MaxExamples.Set(nil)
}

// UnsetMaxExamples ensures that no value is present for MaxExamples, not even an explicit nil
func (o *CreateClassificationRequest) UnsetMaxExamples() {
	o.MaxExamples.Unset()
}

// GetLogitBias returns the LogitBias field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateClassificationRequest) GetLogitBias() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.LogitBias
}

// GetLogitBiasOk returns a tuple with the LogitBias field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateClassificationRequest) GetLogitBiasOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.LogitBias) {
    return map[string]interface{}{}, false
	}
	return o.LogitBias, true
}

// HasLogitBias returns a boolean if a field has been set.
func (o *CreateClassificationRequest) HasLogitBias() bool {
	if o != nil && isNil(o.LogitBias) {
		return true
	}

	return false
}

// SetLogitBias gets a reference to the given map[string]interface{} and assigns it to the LogitBias field.
func (o *CreateClassificationRequest) SetLogitBias(v map[string]interface{}) {
	o.LogitBias = v
}

// GetReturnPrompt returns the ReturnPrompt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateClassificationRequest) GetReturnPrompt() bool {
	if o == nil || isNil(o.ReturnPrompt.Get()) {
		var ret bool
		return ret
	}
	return *o.ReturnPrompt.Get()
}

// GetReturnPromptOk returns a tuple with the ReturnPrompt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateClassificationRequest) GetReturnPromptOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.ReturnPrompt.Get(), o.ReturnPrompt.IsSet()
}

// HasReturnPrompt returns a boolean if a field has been set.
func (o *CreateClassificationRequest) HasReturnPrompt() bool {
	if o != nil && o.ReturnPrompt.IsSet() {
		return true
	}

	return false
}

// SetReturnPrompt gets a reference to the given NullableBool and assigns it to the ReturnPrompt field.
func (o *CreateClassificationRequest) SetReturnPrompt(v bool) {
	o.ReturnPrompt.Set(&v)
}
// SetReturnPromptNil sets the value for ReturnPrompt to be an explicit nil
func (o *CreateClassificationRequest) SetReturnPromptNil() {
	o.ReturnPrompt.Set(nil)
}

// UnsetReturnPrompt ensures that no value is present for ReturnPrompt, not even an explicit nil
func (o *CreateClassificationRequest) UnsetReturnPrompt() {
	o.ReturnPrompt.Unset()
}

// GetReturnMetadata returns the ReturnMetadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateClassificationRequest) GetReturnMetadata() bool {
	if o == nil || isNil(o.ReturnMetadata.Get()) {
		var ret bool
		return ret
	}
	return *o.ReturnMetadata.Get()
}

// GetReturnMetadataOk returns a tuple with the ReturnMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateClassificationRequest) GetReturnMetadataOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.ReturnMetadata.Get(), o.ReturnMetadata.IsSet()
}

// HasReturnMetadata returns a boolean if a field has been set.
func (o *CreateClassificationRequest) HasReturnMetadata() bool {
	if o != nil && o.ReturnMetadata.IsSet() {
		return true
	}

	return false
}

// SetReturnMetadata gets a reference to the given NullableBool and assigns it to the ReturnMetadata field.
func (o *CreateClassificationRequest) SetReturnMetadata(v bool) {
	o.ReturnMetadata.Set(&v)
}
// SetReturnMetadataNil sets the value for ReturnMetadata to be an explicit nil
func (o *CreateClassificationRequest) SetReturnMetadataNil() {
	o.ReturnMetadata.Set(nil)
}

// UnsetReturnMetadata ensures that no value is present for ReturnMetadata, not even an explicit nil
func (o *CreateClassificationRequest) UnsetReturnMetadata() {
	o.ReturnMetadata.Unset()
}

// GetExpand returns the Expand field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateClassificationRequest) GetExpand() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}
	return o.Expand
}

// GetExpandOk returns a tuple with the Expand field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateClassificationRequest) GetExpandOk() ([]interface{}, bool) {
	if o == nil || isNil(o.Expand) {
    return nil, false
	}
	return o.Expand, true
}

// HasExpand returns a boolean if a field has been set.
func (o *CreateClassificationRequest) HasExpand() bool {
	if o != nil && isNil(o.Expand) {
		return true
	}

	return false
}

// SetExpand gets a reference to the given []interface{} and assigns it to the Expand field.
func (o *CreateClassificationRequest) SetExpand(v []interface{}) {
	o.Expand = v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *CreateClassificationRequest) GetUser() string {
	if o == nil || isNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClassificationRequest) GetUserOk() (*string, bool) {
	if o == nil || isNil(o.User) {
    return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *CreateClassificationRequest) HasUser() bool {
	if o != nil && !isNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *CreateClassificationRequest) SetUser(v string) {
	o.User = &v
}

func (o CreateClassificationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["model"] = o.Model
	}
	if true {
		toSerialize["query"] = o.Query
	}
	if o.Examples != nil {
		toSerialize["examples"] = o.Examples
	}
	if o.File.IsSet() {
		toSerialize["file"] = o.File.Get()
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.SearchModel.IsSet() {
		toSerialize["search_model"] = o.SearchModel.Get()
	}
	if o.Temperature.IsSet() {
		toSerialize["temperature"] = o.Temperature.Get()
	}
	if o.Logprobs.IsSet() {
		toSerialize["logprobs"] = o.Logprobs.Get()
	}
	if o.MaxExamples.IsSet() {
		toSerialize["max_examples"] = o.MaxExamples.Get()
	}
	if o.LogitBias != nil {
		toSerialize["logit_bias"] = o.LogitBias
	}
	if o.ReturnPrompt.IsSet() {
		toSerialize["return_prompt"] = o.ReturnPrompt.Get()
	}
	if o.ReturnMetadata.IsSet() {
		toSerialize["return_metadata"] = o.ReturnMetadata.Get()
	}
	if o.Expand != nil {
		toSerialize["expand"] = o.Expand
	}
	if !isNil(o.User) {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableCreateClassificationRequest struct {
	value *CreateClassificationRequest
	isSet bool
}

func (v NullableCreateClassificationRequest) Get() *CreateClassificationRequest {
	return v.value
}

func (v *NullableCreateClassificationRequest) Set(val *CreateClassificationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateClassificationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateClassificationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateClassificationRequest(val *CreateClassificationRequest) *NullableCreateClassificationRequest {
	return &NullableCreateClassificationRequest{value: val, isSet: true}
}

func (v NullableCreateClassificationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateClassificationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


