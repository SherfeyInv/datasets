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

// V1AnnotatedAssemblies struct for V1AnnotatedAssemblies
type V1AnnotatedAssemblies struct {
	Accession *string `json:"accession,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewV1AnnotatedAssemblies instantiates a new V1AnnotatedAssemblies object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1AnnotatedAssemblies() *V1AnnotatedAssemblies {
	this := V1AnnotatedAssemblies{}
	return &this
}

// NewV1AnnotatedAssembliesWithDefaults instantiates a new V1AnnotatedAssemblies object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1AnnotatedAssembliesWithDefaults() *V1AnnotatedAssemblies {
	this := V1AnnotatedAssemblies{}
	return &this
}

// GetAccession returns the Accession field value if set, zero value otherwise.
func (o *V1AnnotatedAssemblies) GetAccession() string {
	if o == nil || o.Accession == nil {
		var ret string
		return ret
	}
	return *o.Accession
}

// GetAccessionOk returns a tuple with the Accession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AnnotatedAssemblies) GetAccessionOk() (*string, bool) {
	if o == nil || o.Accession == nil {
		return nil, false
	}
	return o.Accession, true
}

// HasAccession returns a boolean if a field has been set.
func (o *V1AnnotatedAssemblies) HasAccession() bool {
	if o != nil && o.Accession != nil {
		return true
	}

	return false
}

// SetAccession gets a reference to the given string and assigns it to the Accession field.
func (o *V1AnnotatedAssemblies) SetAccession(v string) {
	o.Accession = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1AnnotatedAssemblies) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AnnotatedAssemblies) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1AnnotatedAssemblies) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1AnnotatedAssemblies) SetName(v string) {
	o.Name = &v
}

func (o V1AnnotatedAssemblies) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Accession != nil  {
		toSerialize["accession"] = o.Accession
	}
	if o.Name != nil  {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableV1AnnotatedAssemblies struct {
	value *V1AnnotatedAssemblies
	isSet bool
}

func (v NullableV1AnnotatedAssemblies) Get() *V1AnnotatedAssemblies {
	return v.value
}

func (v *NullableV1AnnotatedAssemblies) Set(val *V1AnnotatedAssemblies) {
	v.value = val
	v.isSet = true
}

func (v NullableV1AnnotatedAssemblies) IsSet() bool {
	return v.isSet
}

func (v *NullableV1AnnotatedAssemblies) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1AnnotatedAssemblies(val *V1AnnotatedAssemblies) *NullableV1AnnotatedAssemblies {
	return &NullableV1AnnotatedAssemblies{value: val, isSet: true}
}

func (v NullableV1AnnotatedAssemblies) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1AnnotatedAssemblies) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


