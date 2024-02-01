/*
JQuants

JQuants

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package jquants

import (
	"encoding/json"
)

// checks if the ListedInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListedInfo{}

// ListedInfo listed info
type ListedInfo struct {
	Date *string `json:"Date,omitempty"`
	Code *string `json:"Code,omitempty"`
	CompanyName *string `json:"CompanyName,omitempty"`
	Sector17Code *string `json:"Sector17Code,omitempty"`
	Sector17CodeName *string `json:"Sector17CodeName,omitempty"`
	Sector33Code *string `json:"Sector33Code,omitempty"`
	Sector33CodeName *string `json:"Sector33CodeName,omitempty"`
	ScaleCategory *string `json:"ScaleCategory,omitempty"`
	MarketCode *string `json:"MarketCode,omitempty"`
	MarketCodeName *string `json:"MarketCodeName,omitempty"`
	MarginCode *string `json:"MarginCode,omitempty"`
	MarginCodeName *string `json:"MarginCodeName,omitempty"`
}

// NewListedInfo instantiates a new ListedInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListedInfo() *ListedInfo {
	this := ListedInfo{}
	return &this
}

// NewListedInfoWithDefaults instantiates a new ListedInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListedInfoWithDefaults() *ListedInfo {
	this := ListedInfo{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *ListedInfo) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListedInfo) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *ListedInfo) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *ListedInfo) SetDate(v string) {
	o.Date = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ListedInfo) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListedInfo) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ListedInfo) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ListedInfo) SetCode(v string) {
	o.Code = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *ListedInfo) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListedInfo) GetCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyName) {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *ListedInfo) HasCompanyName() bool {
	if o != nil && !IsNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *ListedInfo) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetSector17Code returns the Sector17Code field value if set, zero value otherwise.
func (o *ListedInfo) GetSector17Code() string {
	if o == nil || IsNil(o.Sector17Code) {
		var ret string
		return ret
	}
	return *o.Sector17Code
}

// GetSector17CodeOk returns a tuple with the Sector17Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListedInfo) GetSector17CodeOk() (*string, bool) {
	if o == nil || IsNil(o.Sector17Code) {
		return nil, false
	}
	return o.Sector17Code, true
}

// HasSector17Code returns a boolean if a field has been set.
func (o *ListedInfo) HasSector17Code() bool {
	if o != nil && !IsNil(o.Sector17Code) {
		return true
	}

	return false
}

// SetSector17Code gets a reference to the given string and assigns it to the Sector17Code field.
func (o *ListedInfo) SetSector17Code(v string) {
	o.Sector17Code = &v
}

// GetSector17CodeName returns the Sector17CodeName field value if set, zero value otherwise.
func (o *ListedInfo) GetSector17CodeName() string {
	if o == nil || IsNil(o.Sector17CodeName) {
		var ret string
		return ret
	}
	return *o.Sector17CodeName
}

// GetSector17CodeNameOk returns a tuple with the Sector17CodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListedInfo) GetSector17CodeNameOk() (*string, bool) {
	if o == nil || IsNil(o.Sector17CodeName) {
		return nil, false
	}
	return o.Sector17CodeName, true
}

// HasSector17CodeName returns a boolean if a field has been set.
func (o *ListedInfo) HasSector17CodeName() bool {
	if o != nil && !IsNil(o.Sector17CodeName) {
		return true
	}

	return false
}

// SetSector17CodeName gets a reference to the given string and assigns it to the Sector17CodeName field.
func (o *ListedInfo) SetSector17CodeName(v string) {
	o.Sector17CodeName = &v
}

// GetSector33Code returns the Sector33Code field value if set, zero value otherwise.
func (o *ListedInfo) GetSector33Code() string {
	if o == nil || IsNil(o.Sector33Code) {
		var ret string
		return ret
	}
	return *o.Sector33Code
}

// GetSector33CodeOk returns a tuple with the Sector33Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListedInfo) GetSector33CodeOk() (*string, bool) {
	if o == nil || IsNil(o.Sector33Code) {
		return nil, false
	}
	return o.Sector33Code, true
}

// HasSector33Code returns a boolean if a field has been set.
func (o *ListedInfo) HasSector33Code() bool {
	if o != nil && !IsNil(o.Sector33Code) {
		return true
	}

	return false
}

// SetSector33Code gets a reference to the given string and assigns it to the Sector33Code field.
func (o *ListedInfo) SetSector33Code(v string) {
	o.Sector33Code = &v
}

// GetSector33CodeName returns the Sector33CodeName field value if set, zero value otherwise.
func (o *ListedInfo) GetSector33CodeName() string {
	if o == nil || IsNil(o.Sector33CodeName) {
		var ret string
		return ret
	}
	return *o.Sector33CodeName
}

// GetSector33CodeNameOk returns a tuple with the Sector33CodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListedInfo) GetSector33CodeNameOk() (*string, bool) {
	if o == nil || IsNil(o.Sector33CodeName) {
		return nil, false
	}
	return o.Sector33CodeName, true
}

// HasSector33CodeName returns a boolean if a field has been set.
func (o *ListedInfo) HasSector33CodeName() bool {
	if o != nil && !IsNil(o.Sector33CodeName) {
		return true
	}

	return false
}

// SetSector33CodeName gets a reference to the given string and assigns it to the Sector33CodeName field.
func (o *ListedInfo) SetSector33CodeName(v string) {
	o.Sector33CodeName = &v
}

// GetScaleCategory returns the ScaleCategory field value if set, zero value otherwise.
func (o *ListedInfo) GetScaleCategory() string {
	if o == nil || IsNil(o.ScaleCategory) {
		var ret string
		return ret
	}
	return *o.ScaleCategory
}

// GetScaleCategoryOk returns a tuple with the ScaleCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListedInfo) GetScaleCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.ScaleCategory) {
		return nil, false
	}
	return o.ScaleCategory, true
}

// HasScaleCategory returns a boolean if a field has been set.
func (o *ListedInfo) HasScaleCategory() bool {
	if o != nil && !IsNil(o.ScaleCategory) {
		return true
	}

	return false
}

// SetScaleCategory gets a reference to the given string and assigns it to the ScaleCategory field.
func (o *ListedInfo) SetScaleCategory(v string) {
	o.ScaleCategory = &v
}

// GetMarketCode returns the MarketCode field value if set, zero value otherwise.
func (o *ListedInfo) GetMarketCode() string {
	if o == nil || IsNil(o.MarketCode) {
		var ret string
		return ret
	}
	return *o.MarketCode
}

// GetMarketCodeOk returns a tuple with the MarketCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListedInfo) GetMarketCodeOk() (*string, bool) {
	if o == nil || IsNil(o.MarketCode) {
		return nil, false
	}
	return o.MarketCode, true
}

// HasMarketCode returns a boolean if a field has been set.
func (o *ListedInfo) HasMarketCode() bool {
	if o != nil && !IsNil(o.MarketCode) {
		return true
	}

	return false
}

// SetMarketCode gets a reference to the given string and assigns it to the MarketCode field.
func (o *ListedInfo) SetMarketCode(v string) {
	o.MarketCode = &v
}

// GetMarketCodeName returns the MarketCodeName field value if set, zero value otherwise.
func (o *ListedInfo) GetMarketCodeName() string {
	if o == nil || IsNil(o.MarketCodeName) {
		var ret string
		return ret
	}
	return *o.MarketCodeName
}

// GetMarketCodeNameOk returns a tuple with the MarketCodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListedInfo) GetMarketCodeNameOk() (*string, bool) {
	if o == nil || IsNil(o.MarketCodeName) {
		return nil, false
	}
	return o.MarketCodeName, true
}

// HasMarketCodeName returns a boolean if a field has been set.
func (o *ListedInfo) HasMarketCodeName() bool {
	if o != nil && !IsNil(o.MarketCodeName) {
		return true
	}

	return false
}

// SetMarketCodeName gets a reference to the given string and assigns it to the MarketCodeName field.
func (o *ListedInfo) SetMarketCodeName(v string) {
	o.MarketCodeName = &v
}

// GetMarginCode returns the MarginCode field value if set, zero value otherwise.
func (o *ListedInfo) GetMarginCode() string {
	if o == nil || IsNil(o.MarginCode) {
		var ret string
		return ret
	}
	return *o.MarginCode
}

// GetMarginCodeOk returns a tuple with the MarginCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListedInfo) GetMarginCodeOk() (*string, bool) {
	if o == nil || IsNil(o.MarginCode) {
		return nil, false
	}
	return o.MarginCode, true
}

// HasMarginCode returns a boolean if a field has been set.
func (o *ListedInfo) HasMarginCode() bool {
	if o != nil && !IsNil(o.MarginCode) {
		return true
	}

	return false
}

// SetMarginCode gets a reference to the given string and assigns it to the MarginCode field.
func (o *ListedInfo) SetMarginCode(v string) {
	o.MarginCode = &v
}

// GetMarginCodeName returns the MarginCodeName field value if set, zero value otherwise.
func (o *ListedInfo) GetMarginCodeName() string {
	if o == nil || IsNil(o.MarginCodeName) {
		var ret string
		return ret
	}
	return *o.MarginCodeName
}

// GetMarginCodeNameOk returns a tuple with the MarginCodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListedInfo) GetMarginCodeNameOk() (*string, bool) {
	if o == nil || IsNil(o.MarginCodeName) {
		return nil, false
	}
	return o.MarginCodeName, true
}

// HasMarginCodeName returns a boolean if a field has been set.
func (o *ListedInfo) HasMarginCodeName() bool {
	if o != nil && !IsNil(o.MarginCodeName) {
		return true
	}

	return false
}

// SetMarginCodeName gets a reference to the given string and assigns it to the MarginCodeName field.
func (o *ListedInfo) SetMarginCodeName(v string) {
	o.MarginCodeName = &v
}

func (o ListedInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListedInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Date) {
		toSerialize["Date"] = o.Date
	}
	if !IsNil(o.Code) {
		toSerialize["Code"] = o.Code
	}
	if !IsNil(o.CompanyName) {
		toSerialize["CompanyName"] = o.CompanyName
	}
	if !IsNil(o.Sector17Code) {
		toSerialize["Sector17Code"] = o.Sector17Code
	}
	if !IsNil(o.Sector17CodeName) {
		toSerialize["Sector17CodeName"] = o.Sector17CodeName
	}
	if !IsNil(o.Sector33Code) {
		toSerialize["Sector33Code"] = o.Sector33Code
	}
	if !IsNil(o.Sector33CodeName) {
		toSerialize["Sector33CodeName"] = o.Sector33CodeName
	}
	if !IsNil(o.ScaleCategory) {
		toSerialize["ScaleCategory"] = o.ScaleCategory
	}
	if !IsNil(o.MarketCode) {
		toSerialize["MarketCode"] = o.MarketCode
	}
	if !IsNil(o.MarketCodeName) {
		toSerialize["MarketCodeName"] = o.MarketCodeName
	}
	if !IsNil(o.MarginCode) {
		toSerialize["MarginCode"] = o.MarginCode
	}
	if !IsNil(o.MarginCodeName) {
		toSerialize["MarginCodeName"] = o.MarginCodeName
	}
	return toSerialize, nil
}

type NullableListedInfo struct {
	value *ListedInfo
	isSet bool
}

func (v NullableListedInfo) Get() *ListedInfo {
	return v.value
}

func (v *NullableListedInfo) Set(val *ListedInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableListedInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableListedInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListedInfo(val *ListedInfo) *NullableListedInfo {
	return &NullableListedInfo{value: val, isSet: true}
}

func (v NullableListedInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListedInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

