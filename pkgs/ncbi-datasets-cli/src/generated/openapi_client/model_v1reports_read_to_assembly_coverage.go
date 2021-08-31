/*
 * NCBI Datasets API
 *
 * ### NCBI Datasets is a resource that lets you easily gather data from NCBI. The Datasets API is still in alpha, and we're updating it often to add new functionality, iron out bugs and enhance usability. For some larger downloads, you may want to download a [dehydrated bag](https://www.ncbi.nlm.nih.gov/datasets/docs/rehydrate/), and retrieve the individual data files at a later time. 
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datasets

import (
	"encoding/json"
)

// V1reportsReadToAssemblyCoverage struct for V1reportsReadToAssemblyCoverage
type V1reportsReadToAssemblyCoverage struct {
	Contig *int32 `json:"contig,omitempty"`
	Assembly *int32 `json:"assembly,omitempty"`
	Ratio *float32 `json:"ratio,omitempty"`
}

// NewV1reportsReadToAssemblyCoverage instantiates a new V1reportsReadToAssemblyCoverage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1reportsReadToAssemblyCoverage() *V1reportsReadToAssemblyCoverage {
	this := V1reportsReadToAssemblyCoverage{}
	return &this
}

// NewV1reportsReadToAssemblyCoverageWithDefaults instantiates a new V1reportsReadToAssemblyCoverage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1reportsReadToAssemblyCoverageWithDefaults() *V1reportsReadToAssemblyCoverage {
	this := V1reportsReadToAssemblyCoverage{}
	return &this
}

// GetContig returns the Contig field value if set, zero value otherwise.
func (o *V1reportsReadToAssemblyCoverage) GetContig() int32 {
	if o == nil || o.Contig == nil {
		var ret int32
		return ret
	}
	return *o.Contig
}

// GetContigOk returns a tuple with the Contig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsReadToAssemblyCoverage) GetContigOk() (*int32, bool) {
	if o == nil || o.Contig == nil {
		return nil, false
	}
	return o.Contig, true
}

// HasContig returns a boolean if a field has been set.
func (o *V1reportsReadToAssemblyCoverage) HasContig() bool {
	if o != nil && o.Contig != nil {
		return true
	}

	return false
}

// SetContig gets a reference to the given int32 and assigns it to the Contig field.
func (o *V1reportsReadToAssemblyCoverage) SetContig(v int32) {
	o.Contig = &v
}

// GetAssembly returns the Assembly field value if set, zero value otherwise.
func (o *V1reportsReadToAssemblyCoverage) GetAssembly() int32 {
	if o == nil || o.Assembly == nil {
		var ret int32
		return ret
	}
	return *o.Assembly
}

// GetAssemblyOk returns a tuple with the Assembly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsReadToAssemblyCoverage) GetAssemblyOk() (*int32, bool) {
	if o == nil || o.Assembly == nil {
		return nil, false
	}
	return o.Assembly, true
}

// HasAssembly returns a boolean if a field has been set.
func (o *V1reportsReadToAssemblyCoverage) HasAssembly() bool {
	if o != nil && o.Assembly != nil {
		return true
	}

	return false
}

// SetAssembly gets a reference to the given int32 and assigns it to the Assembly field.
func (o *V1reportsReadToAssemblyCoverage) SetAssembly(v int32) {
	o.Assembly = &v
}

// GetRatio returns the Ratio field value if set, zero value otherwise.
func (o *V1reportsReadToAssemblyCoverage) GetRatio() float32 {
	if o == nil || o.Ratio == nil {
		var ret float32
		return ret
	}
	return *o.Ratio
}

// GetRatioOk returns a tuple with the Ratio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsReadToAssemblyCoverage) GetRatioOk() (*float32, bool) {
	if o == nil || o.Ratio == nil {
		return nil, false
	}
	return o.Ratio, true
}

// HasRatio returns a boolean if a field has been set.
func (o *V1reportsReadToAssemblyCoverage) HasRatio() bool {
	if o != nil && o.Ratio != nil {
		return true
	}

	return false
}

// SetRatio gets a reference to the given float32 and assigns it to the Ratio field.
func (o *V1reportsReadToAssemblyCoverage) SetRatio(v float32) {
	o.Ratio = &v
}

func (o V1reportsReadToAssemblyCoverage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Contig != nil  {
		toSerialize["contig"] = o.Contig
	}
	if o.Assembly != nil  {
		toSerialize["assembly"] = o.Assembly
	}
	if o.Ratio != nil  {
		toSerialize["ratio"] = o.Ratio
	}
	return json.Marshal(toSerialize)
}

type NullableV1reportsReadToAssemblyCoverage struct {
	value *V1reportsReadToAssemblyCoverage
	isSet bool
}

func (v NullableV1reportsReadToAssemblyCoverage) Get() *V1reportsReadToAssemblyCoverage {
	return v.value
}

func (v *NullableV1reportsReadToAssemblyCoverage) Set(val *V1reportsReadToAssemblyCoverage) {
	v.value = val
	v.isSet = true
}

func (v NullableV1reportsReadToAssemblyCoverage) IsSet() bool {
	return v.isSet
}

func (v *NullableV1reportsReadToAssemblyCoverage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1reportsReadToAssemblyCoverage(val *V1reportsReadToAssemblyCoverage) *NullableV1reportsReadToAssemblyCoverage {
	return &NullableV1reportsReadToAssemblyCoverage{value: val, isSet: true}
}

func (v NullableV1reportsReadToAssemblyCoverage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1reportsReadToAssemblyCoverage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


