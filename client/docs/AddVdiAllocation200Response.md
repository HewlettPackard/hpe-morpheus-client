# AddVdiAllocation200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Desktop** | Pointer to [**ListVdi200ResponseAllOfDesktopsInner**](ListVdi200ResponseAllOfDesktopsInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddVdiAllocation200Response

`func NewAddVdiAllocation200Response() *AddVdiAllocation200Response`

NewAddVdiAllocation200Response instantiates a new AddVdiAllocation200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddVdiAllocation200ResponseWithDefaults

`func NewAddVdiAllocation200ResponseWithDefaults() *AddVdiAllocation200Response`

NewAddVdiAllocation200ResponseWithDefaults instantiates a new AddVdiAllocation200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDesktop

`func (o *AddVdiAllocation200Response) GetDesktop() ListVdi200ResponseAllOfDesktopsInner`

GetDesktop returns the Desktop field if non-nil, zero value otherwise.

### GetDesktopOk

`func (o *AddVdiAllocation200Response) GetDesktopOk() (*ListVdi200ResponseAllOfDesktopsInner, bool)`

GetDesktopOk returns a tuple with the Desktop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktop

`func (o *AddVdiAllocation200Response) SetDesktop(v ListVdi200ResponseAllOfDesktopsInner)`

SetDesktop sets Desktop field to given value.

### HasDesktop

`func (o *AddVdiAllocation200Response) HasDesktop() bool`

HasDesktop returns a boolean if a field has been set.

### GetSuccess

`func (o *AddVdiAllocation200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddVdiAllocation200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddVdiAllocation200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddVdiAllocation200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


