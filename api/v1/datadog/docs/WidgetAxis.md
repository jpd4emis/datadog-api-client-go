# WidgetAxis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeZero** | Pointer to **bool** |  | [optional] 
**Label** | Pointer to **string** | The label of the axis to display on the graph. | [optional] 
**Max** | Pointer to **string** | Specifies the maximum value to show on the y-axis. It takes a number, or auto for default behavior. | [optional] [default to "auto"]
**Min** | Pointer to **string** | Specifies minimum value to show on the y-axis. It takes a number, or auto for default behavior. | [optional] [default to "auto"]
**Scale** | Pointer to **string** | Specifies the scale type. Possible values are linear, log, sqrt, pow## (eg. pow2, pow0.5..) | [optional] [default to "linear"]

## Methods

### NewWidgetAxis

`func NewWidgetAxis() *WidgetAxis`

NewWidgetAxis instantiates a new WidgetAxis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWidgetAxisWithDefaults

`func NewWidgetAxisWithDefaults() *WidgetAxis`

NewWidgetAxisWithDefaults instantiates a new WidgetAxis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeZero

`func (o *WidgetAxis) GetIncludeZero() bool`

GetIncludeZero returns the IncludeZero field if non-nil, zero value otherwise.

### GetIncludeZeroOk

`func (o *WidgetAxis) GetIncludeZeroOk() (bool, bool)`

GetIncludeZeroOk returns a tuple with the IncludeZero field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIncludeZero

`func (o *WidgetAxis) HasIncludeZero() bool`

HasIncludeZero returns a boolean if a field has been set.

### SetIncludeZero

`func (o *WidgetAxis) SetIncludeZero(v bool)`

SetIncludeZero gets a reference to the given bool and assigns it to the IncludeZero field.

### GetLabel

`func (o *WidgetAxis) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WidgetAxis) GetLabelOk() (string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLabel

`func (o *WidgetAxis) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabel

`func (o *WidgetAxis) SetLabel(v string)`

SetLabel gets a reference to the given string and assigns it to the Label field.

### GetMax

`func (o *WidgetAxis) GetMax() string`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *WidgetAxis) GetMaxOk() (string, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMax

`func (o *WidgetAxis) HasMax() bool`

HasMax returns a boolean if a field has been set.

### SetMax

`func (o *WidgetAxis) SetMax(v string)`

SetMax gets a reference to the given string and assigns it to the Max field.

### GetMin

`func (o *WidgetAxis) GetMin() string`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *WidgetAxis) GetMinOk() (string, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMin

`func (o *WidgetAxis) HasMin() bool`

HasMin returns a boolean if a field has been set.

### SetMin

`func (o *WidgetAxis) SetMin(v string)`

SetMin gets a reference to the given string and assigns it to the Min field.

### GetScale

`func (o *WidgetAxis) GetScale() string`

GetScale returns the Scale field if non-nil, zero value otherwise.

### GetScaleOk

`func (o *WidgetAxis) GetScaleOk() (string, bool)`

GetScaleOk returns a tuple with the Scale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScale

`func (o *WidgetAxis) HasScale() bool`

HasScale returns a boolean if a field has been set.

### SetScale

`func (o *WidgetAxis) SetScale(v string)`

SetScale gets a reference to the given string and assigns it to the Scale field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

