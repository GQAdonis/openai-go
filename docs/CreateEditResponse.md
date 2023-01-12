# CreateEditResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Object** | **string** |  | 
**Created** | **int32** |  | 
**Model** | **string** |  | 
**Choices** | [**[]CreateCompletionResponseChoicesInner**](CreateCompletionResponseChoicesInner.md) |  | 
**Usage** | [**CreateCompletionResponseUsage**](CreateCompletionResponseUsage.md) |  | 

## Methods

### NewCreateEditResponse

`func NewCreateEditResponse(id string, object string, created int32, model string, choices []CreateCompletionResponseChoicesInner, usage CreateCompletionResponseUsage, ) *CreateEditResponse`

NewCreateEditResponse instantiates a new CreateEditResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEditResponseWithDefaults

`func NewCreateEditResponseWithDefaults() *CreateEditResponse`

NewCreateEditResponseWithDefaults instantiates a new CreateEditResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateEditResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateEditResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateEditResponse) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *CreateEditResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CreateEditResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CreateEditResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetCreated

`func (o *CreateEditResponse) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CreateEditResponse) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CreateEditResponse) SetCreated(v int32)`

SetCreated sets Created field to given value.


### GetModel

`func (o *CreateEditResponse) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CreateEditResponse) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CreateEditResponse) SetModel(v string)`

SetModel sets Model field to given value.


### GetChoices

`func (o *CreateEditResponse) GetChoices() []CreateCompletionResponseChoicesInner`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *CreateEditResponse) GetChoicesOk() (*[]CreateCompletionResponseChoicesInner, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *CreateEditResponse) SetChoices(v []CreateCompletionResponseChoicesInner)`

SetChoices sets Choices field to given value.


### GetUsage

`func (o *CreateEditResponse) GetUsage() CreateCompletionResponseUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *CreateEditResponse) GetUsageOk() (*CreateCompletionResponseUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *CreateEditResponse) SetUsage(v CreateCompletionResponseUsage)`

SetUsage sets Usage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


