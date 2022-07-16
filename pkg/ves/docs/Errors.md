# Errors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** | DVLA reference code | [optional] 
**Title** | **string** | Error title | 
**Detail** | Pointer to **string** | A meaningful description of the error which has occurred | [optional] 

## Methods

### NewErrors

`func NewErrors(title string, ) *Errors`

NewErrors instantiates a new Errors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsWithDefaults

`func NewErrorsWithDefaults() *Errors`

NewErrorsWithDefaults instantiates a new Errors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Errors) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Errors) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Errors) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Errors) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCode

`func (o *Errors) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Errors) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Errors) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Errors) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetTitle

`func (o *Errors) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Errors) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Errors) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDetail

`func (o *Errors) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *Errors) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *Errors) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *Errors) HasDetail() bool`

HasDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


