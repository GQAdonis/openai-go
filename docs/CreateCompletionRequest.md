# CreateCompletionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** | ID of the model to use. You can use the [List models](/docs/api-reference/models/list) API to see all of your available models, or see our [Model overview](/docs/models/overview) for descriptions of them. | 
**Prompt** | Pointer to [**NullableCreateCompletionRequestPrompt**](CreateCompletionRequestPrompt.md) |  | [optional] [default to <|endoftext|>]
**Suffix** | Pointer to **NullableString** | The suffix that comes after a completion of inserted text. | [optional] 
**MaxTokens** | Pointer to **NullableInt32** | The maximum number of [tokens](/tokenizer) to generate in the completion.  The token count of your prompt plus &#x60;max_tokens&#x60; cannot exceed the model&#39;s context length. Most models have a context length of 2048 tokens (except for the newest models, which support 4096).  | [optional] [default to 16]
**Temperature** | Pointer to **NullableFloat32** | What [sampling temperature](https://towardsdatascience.com/how-to-sample-from-language-models-682bceb97277) to use. Higher values means the model will take more risks. Try 0.9 for more creative applications, and 0 (argmax sampling) for ones with a well-defined answer.  We generally recommend altering this or &#x60;top_p&#x60; but not both.  | [optional] [default to 1]
**TopP** | Pointer to **NullableFloat32** | An alternative to sampling with temperature, called nucleus sampling, where the model considers the results of the tokens with top_p probability mass. So 0.1 means only the tokens comprising the top 10% probability mass are considered.  We generally recommend altering this or &#x60;temperature&#x60; but not both.  | [optional] [default to 1]
**N** | Pointer to **NullableInt32** | How many completions to generate for each prompt.  **Note:** Because this parameter generates many completions, it can quickly consume your token quota. Use carefully and ensure that you have reasonable settings for &#x60;max_tokens&#x60; and &#x60;stop&#x60;.  | [optional] [default to 1]
**Stream** | Pointer to **NullableBool** | Whether to stream back partial progress. If set, tokens will be sent as data-only [server-sent events](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events#Event_stream_format) as they become available, with the stream terminated by a &#x60;data: [DONE]&#x60; message.  | [optional] [default to false]
**Logprobs** | Pointer to **NullableInt32** | Include the log probabilities on the &#x60;logprobs&#x60; most likely tokens, as well the chosen tokens. For example, if &#x60;logprobs&#x60; is 5, the API will return a list of the 5 most likely tokens. The API will always return the &#x60;logprob&#x60; of the sampled token, so there may be up to &#x60;logprobs+1&#x60; elements in the response.  The maximum value for &#x60;logprobs&#x60; is 5. If you need more than this, please contact us through our [Help center](https://help.openai.com) and describe your use case.  | [optional] 
**Echo** | Pointer to **NullableBool** | Echo back the prompt in addition to the completion  | [optional] [default to false]
**Stop** | Pointer to [**NullableCreateCompletionRequestStop**](CreateCompletionRequestStop.md) |  | [optional] [default to null]
**PresencePenalty** | Pointer to **NullableFloat32** | Number between -2.0 and 2.0. Positive values penalize new tokens based on whether they appear in the text so far, increasing the model&#39;s likelihood to talk about new topics.  [See more information about frequency and presence penalties.](/docs/api-reference/parameter-details)  | [optional] [default to 0]
**FrequencyPenalty** | Pointer to **NullableFloat32** | Number between -2.0 and 2.0. Positive values penalize new tokens based on their existing frequency in the text so far, decreasing the model&#39;s likelihood to repeat the same line verbatim.  [See more information about frequency and presence penalties.](/docs/api-reference/parameter-details)  | [optional] [default to 0]
**BestOf** | Pointer to **NullableInt32** | Generates &#x60;best_of&#x60; completions server-side and returns the \&quot;best\&quot; (the one with the highest log probability per token). Results cannot be streamed.  When used with &#x60;n&#x60;, &#x60;best_of&#x60; controls the number of candidate completions and &#x60;n&#x60; specifies how many to return ??? &#x60;best_of&#x60; must be greater than &#x60;n&#x60;.  **Note:** Because this parameter generates many completions, it can quickly consume your token quota. Use carefully and ensure that you have reasonable settings for &#x60;max_tokens&#x60; and &#x60;stop&#x60;.  | [optional] [default to 1]
**LogitBias** | Pointer to **map[string]interface{}** | Modify the likelihood of specified tokens appearing in the completion.  Accepts a json object that maps tokens (specified by their token ID in the GPT tokenizer) to an associated bias value from -100 to 100. You can use this [tokenizer tool](/tokenizer?view&#x3D;bpe) (which works for both GPT-2 and GPT-3) to convert text to token IDs. Mathematically, the bias is added to the logits generated by the model prior to sampling. The exact effect will vary per model, but values between -1 and 1 should decrease or increase likelihood of selection; values like -100 or 100 should result in a ban or exclusive selection of the relevant token.  As an example, you can pass &#x60;{\&quot;50256\&quot;: -100}&#x60; to prevent the &lt;|endoftext|&gt; token from being generated.  | [optional] 
**User** | Pointer to **string** | A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse. [Learn more](/docs/guides/safety-best-practices/end-user-ids).  | [optional] 

## Methods

### NewCreateCompletionRequest

`func NewCreateCompletionRequest(model string, ) *CreateCompletionRequest`

NewCreateCompletionRequest instantiates a new CreateCompletionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCompletionRequestWithDefaults

`func NewCreateCompletionRequestWithDefaults() *CreateCompletionRequest`

NewCreateCompletionRequestWithDefaults instantiates a new CreateCompletionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *CreateCompletionRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CreateCompletionRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CreateCompletionRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetPrompt

`func (o *CreateCompletionRequest) GetPrompt() CreateCompletionRequestPrompt`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *CreateCompletionRequest) GetPromptOk() (*CreateCompletionRequestPrompt, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *CreateCompletionRequest) SetPrompt(v CreateCompletionRequestPrompt)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *CreateCompletionRequest) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### SetPromptNil

`func (o *CreateCompletionRequest) SetPromptNil(b bool)`

 SetPromptNil sets the value for Prompt to be an explicit nil

### UnsetPrompt
`func (o *CreateCompletionRequest) UnsetPrompt()`

UnsetPrompt ensures that no value is present for Prompt, not even an explicit nil
### GetSuffix

`func (o *CreateCompletionRequest) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *CreateCompletionRequest) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *CreateCompletionRequest) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *CreateCompletionRequest) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### SetSuffixNil

`func (o *CreateCompletionRequest) SetSuffixNil(b bool)`

 SetSuffixNil sets the value for Suffix to be an explicit nil

### UnsetSuffix
`func (o *CreateCompletionRequest) UnsetSuffix()`

UnsetSuffix ensures that no value is present for Suffix, not even an explicit nil
### GetMaxTokens

`func (o *CreateCompletionRequest) GetMaxTokens() int32`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *CreateCompletionRequest) GetMaxTokensOk() (*int32, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *CreateCompletionRequest) SetMaxTokens(v int32)`

SetMaxTokens sets MaxTokens field to given value.

### HasMaxTokens

`func (o *CreateCompletionRequest) HasMaxTokens() bool`

HasMaxTokens returns a boolean if a field has been set.

### SetMaxTokensNil

`func (o *CreateCompletionRequest) SetMaxTokensNil(b bool)`

 SetMaxTokensNil sets the value for MaxTokens to be an explicit nil

### UnsetMaxTokens
`func (o *CreateCompletionRequest) UnsetMaxTokens()`

UnsetMaxTokens ensures that no value is present for MaxTokens, not even an explicit nil
### GetTemperature

`func (o *CreateCompletionRequest) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *CreateCompletionRequest) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *CreateCompletionRequest) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *CreateCompletionRequest) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### SetTemperatureNil

`func (o *CreateCompletionRequest) SetTemperatureNil(b bool)`

 SetTemperatureNil sets the value for Temperature to be an explicit nil

### UnsetTemperature
`func (o *CreateCompletionRequest) UnsetTemperature()`

UnsetTemperature ensures that no value is present for Temperature, not even an explicit nil
### GetTopP

`func (o *CreateCompletionRequest) GetTopP() float32`

GetTopP returns the TopP field if non-nil, zero value otherwise.

### GetTopPOk

`func (o *CreateCompletionRequest) GetTopPOk() (*float32, bool)`

GetTopPOk returns a tuple with the TopP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopP

`func (o *CreateCompletionRequest) SetTopP(v float32)`

SetTopP sets TopP field to given value.

### HasTopP

`func (o *CreateCompletionRequest) HasTopP() bool`

HasTopP returns a boolean if a field has been set.

### SetTopPNil

`func (o *CreateCompletionRequest) SetTopPNil(b bool)`

 SetTopPNil sets the value for TopP to be an explicit nil

### UnsetTopP
`func (o *CreateCompletionRequest) UnsetTopP()`

UnsetTopP ensures that no value is present for TopP, not even an explicit nil
### GetN

`func (o *CreateCompletionRequest) GetN() int32`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *CreateCompletionRequest) GetNOk() (*int32, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *CreateCompletionRequest) SetN(v int32)`

SetN sets N field to given value.

### HasN

`func (o *CreateCompletionRequest) HasN() bool`

HasN returns a boolean if a field has been set.

### SetNNil

`func (o *CreateCompletionRequest) SetNNil(b bool)`

 SetNNil sets the value for N to be an explicit nil

### UnsetN
`func (o *CreateCompletionRequest) UnsetN()`

UnsetN ensures that no value is present for N, not even an explicit nil
### GetStream

`func (o *CreateCompletionRequest) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *CreateCompletionRequest) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *CreateCompletionRequest) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *CreateCompletionRequest) HasStream() bool`

HasStream returns a boolean if a field has been set.

### SetStreamNil

`func (o *CreateCompletionRequest) SetStreamNil(b bool)`

 SetStreamNil sets the value for Stream to be an explicit nil

### UnsetStream
`func (o *CreateCompletionRequest) UnsetStream()`

UnsetStream ensures that no value is present for Stream, not even an explicit nil
### GetLogprobs

`func (o *CreateCompletionRequest) GetLogprobs() int32`

GetLogprobs returns the Logprobs field if non-nil, zero value otherwise.

### GetLogprobsOk

`func (o *CreateCompletionRequest) GetLogprobsOk() (*int32, bool)`

GetLogprobsOk returns a tuple with the Logprobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogprobs

`func (o *CreateCompletionRequest) SetLogprobs(v int32)`

SetLogprobs sets Logprobs field to given value.

### HasLogprobs

`func (o *CreateCompletionRequest) HasLogprobs() bool`

HasLogprobs returns a boolean if a field has been set.

### SetLogprobsNil

`func (o *CreateCompletionRequest) SetLogprobsNil(b bool)`

 SetLogprobsNil sets the value for Logprobs to be an explicit nil

### UnsetLogprobs
`func (o *CreateCompletionRequest) UnsetLogprobs()`

UnsetLogprobs ensures that no value is present for Logprobs, not even an explicit nil
### GetEcho

`func (o *CreateCompletionRequest) GetEcho() bool`

GetEcho returns the Echo field if non-nil, zero value otherwise.

### GetEchoOk

`func (o *CreateCompletionRequest) GetEchoOk() (*bool, bool)`

GetEchoOk returns a tuple with the Echo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcho

`func (o *CreateCompletionRequest) SetEcho(v bool)`

SetEcho sets Echo field to given value.

### HasEcho

`func (o *CreateCompletionRequest) HasEcho() bool`

HasEcho returns a boolean if a field has been set.

### SetEchoNil

`func (o *CreateCompletionRequest) SetEchoNil(b bool)`

 SetEchoNil sets the value for Echo to be an explicit nil

### UnsetEcho
`func (o *CreateCompletionRequest) UnsetEcho()`

UnsetEcho ensures that no value is present for Echo, not even an explicit nil
### GetStop

`func (o *CreateCompletionRequest) GetStop() CreateCompletionRequestStop`

GetStop returns the Stop field if non-nil, zero value otherwise.

### GetStopOk

`func (o *CreateCompletionRequest) GetStopOk() (*CreateCompletionRequestStop, bool)`

GetStopOk returns a tuple with the Stop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStop

`func (o *CreateCompletionRequest) SetStop(v CreateCompletionRequestStop)`

SetStop sets Stop field to given value.

### HasStop

`func (o *CreateCompletionRequest) HasStop() bool`

HasStop returns a boolean if a field has been set.

### SetStopNil

`func (o *CreateCompletionRequest) SetStopNil(b bool)`

 SetStopNil sets the value for Stop to be an explicit nil

### UnsetStop
`func (o *CreateCompletionRequest) UnsetStop()`

UnsetStop ensures that no value is present for Stop, not even an explicit nil
### GetPresencePenalty

`func (o *CreateCompletionRequest) GetPresencePenalty() float32`

GetPresencePenalty returns the PresencePenalty field if non-nil, zero value otherwise.

### GetPresencePenaltyOk

`func (o *CreateCompletionRequest) GetPresencePenaltyOk() (*float32, bool)`

GetPresencePenaltyOk returns a tuple with the PresencePenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresencePenalty

`func (o *CreateCompletionRequest) SetPresencePenalty(v float32)`

SetPresencePenalty sets PresencePenalty field to given value.

### HasPresencePenalty

`func (o *CreateCompletionRequest) HasPresencePenalty() bool`

HasPresencePenalty returns a boolean if a field has been set.

### SetPresencePenaltyNil

`func (o *CreateCompletionRequest) SetPresencePenaltyNil(b bool)`

 SetPresencePenaltyNil sets the value for PresencePenalty to be an explicit nil

### UnsetPresencePenalty
`func (o *CreateCompletionRequest) UnsetPresencePenalty()`

UnsetPresencePenalty ensures that no value is present for PresencePenalty, not even an explicit nil
### GetFrequencyPenalty

`func (o *CreateCompletionRequest) GetFrequencyPenalty() float32`

GetFrequencyPenalty returns the FrequencyPenalty field if non-nil, zero value otherwise.

### GetFrequencyPenaltyOk

`func (o *CreateCompletionRequest) GetFrequencyPenaltyOk() (*float32, bool)`

GetFrequencyPenaltyOk returns a tuple with the FrequencyPenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyPenalty

`func (o *CreateCompletionRequest) SetFrequencyPenalty(v float32)`

SetFrequencyPenalty sets FrequencyPenalty field to given value.

### HasFrequencyPenalty

`func (o *CreateCompletionRequest) HasFrequencyPenalty() bool`

HasFrequencyPenalty returns a boolean if a field has been set.

### SetFrequencyPenaltyNil

`func (o *CreateCompletionRequest) SetFrequencyPenaltyNil(b bool)`

 SetFrequencyPenaltyNil sets the value for FrequencyPenalty to be an explicit nil

### UnsetFrequencyPenalty
`func (o *CreateCompletionRequest) UnsetFrequencyPenalty()`

UnsetFrequencyPenalty ensures that no value is present for FrequencyPenalty, not even an explicit nil
### GetBestOf

`func (o *CreateCompletionRequest) GetBestOf() int32`

GetBestOf returns the BestOf field if non-nil, zero value otherwise.

### GetBestOfOk

`func (o *CreateCompletionRequest) GetBestOfOk() (*int32, bool)`

GetBestOfOk returns a tuple with the BestOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestOf

`func (o *CreateCompletionRequest) SetBestOf(v int32)`

SetBestOf sets BestOf field to given value.

### HasBestOf

`func (o *CreateCompletionRequest) HasBestOf() bool`

HasBestOf returns a boolean if a field has been set.

### SetBestOfNil

`func (o *CreateCompletionRequest) SetBestOfNil(b bool)`

 SetBestOfNil sets the value for BestOf to be an explicit nil

### UnsetBestOf
`func (o *CreateCompletionRequest) UnsetBestOf()`

UnsetBestOf ensures that no value is present for BestOf, not even an explicit nil
### GetLogitBias

`func (o *CreateCompletionRequest) GetLogitBias() map[string]interface{}`

GetLogitBias returns the LogitBias field if non-nil, zero value otherwise.

### GetLogitBiasOk

`func (o *CreateCompletionRequest) GetLogitBiasOk() (*map[string]interface{}, bool)`

GetLogitBiasOk returns a tuple with the LogitBias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogitBias

`func (o *CreateCompletionRequest) SetLogitBias(v map[string]interface{})`

SetLogitBias sets LogitBias field to given value.

### HasLogitBias

`func (o *CreateCompletionRequest) HasLogitBias() bool`

HasLogitBias returns a boolean if a field has been set.

### SetLogitBiasNil

`func (o *CreateCompletionRequest) SetLogitBiasNil(b bool)`

 SetLogitBiasNil sets the value for LogitBias to be an explicit nil

### UnsetLogitBias
`func (o *CreateCompletionRequest) UnsetLogitBias()`

UnsetLogitBias ensures that no value is present for LogitBias, not even an explicit nil
### GetUser

`func (o *CreateCompletionRequest) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CreateCompletionRequest) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CreateCompletionRequest) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *CreateCompletionRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


