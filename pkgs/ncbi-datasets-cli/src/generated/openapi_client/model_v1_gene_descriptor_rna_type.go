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
	"fmt"
)

// V1GeneDescriptorRnaType the model 'V1GeneDescriptorRnaType'
type V1GeneDescriptorRnaType string

// List of v1GeneDescriptorRnaType
const (
	V1GENEDESCRIPTORRNATYPE_RNA_UNKNOWN V1GeneDescriptorRnaType = "rna_UNKNOWN"
	V1GENEDESCRIPTORRNATYPE_PREMSG V1GeneDescriptorRnaType = "premsg"
	V1GENEDESCRIPTORRNATYPE_TM_RNA V1GeneDescriptorRnaType = "tmRna"
)

var allowedV1GeneDescriptorRnaTypeEnumValues = []V1GeneDescriptorRnaType{
	"rna_UNKNOWN",
	"premsg",
	"tmRna",
}

func (v *V1GeneDescriptorRnaType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := V1GeneDescriptorRnaType(value)
	for _, existing := range allowedV1GeneDescriptorRnaTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid V1GeneDescriptorRnaType", value)
}

// NewV1GeneDescriptorRnaTypeFromValue returns a pointer to a valid V1GeneDescriptorRnaType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewV1GeneDescriptorRnaTypeFromValue(v string) (*V1GeneDescriptorRnaType, error) {
	ev := V1GeneDescriptorRnaType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for V1GeneDescriptorRnaType: valid values are %v", v, allowedV1GeneDescriptorRnaTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v V1GeneDescriptorRnaType) IsValid() bool {
	for _, existing := range allowedV1GeneDescriptorRnaTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to v1GeneDescriptorRnaType value
func (v V1GeneDescriptorRnaType) Ptr() *V1GeneDescriptorRnaType {
	return &v
}

type NullableV1GeneDescriptorRnaType struct {
	value *V1GeneDescriptorRnaType
	isSet bool
}

func (v NullableV1GeneDescriptorRnaType) Get() *V1GeneDescriptorRnaType {
	return v.value
}

func (v *NullableV1GeneDescriptorRnaType) Set(val *V1GeneDescriptorRnaType) {
	v.value = val
	v.isSet = true
}

func (v NullableV1GeneDescriptorRnaType) IsSet() bool {
	return v.isSet
}

func (v *NullableV1GeneDescriptorRnaType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1GeneDescriptorRnaType(val *V1GeneDescriptorRnaType) *NullableV1GeneDescriptorRnaType {
	return &NullableV1GeneDescriptorRnaType{value: val, isSet: true}
}

func (v NullableV1GeneDescriptorRnaType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1GeneDescriptorRnaType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

