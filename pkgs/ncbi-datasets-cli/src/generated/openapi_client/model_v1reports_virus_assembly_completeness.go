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

// V1reportsVirusAssemblyCompleteness the model 'V1reportsVirusAssemblyCompleteness'
type V1reportsVirusAssemblyCompleteness string

// List of v1reportsVirusAssemblyCompleteness
const (
	V1REPORTSVIRUSASSEMBLYCOMPLETENESS_UNKNOWN V1reportsVirusAssemblyCompleteness = "UNKNOWN"
	V1REPORTSVIRUSASSEMBLYCOMPLETENESS_COMPLETE V1reportsVirusAssemblyCompleteness = "COMPLETE"
	V1REPORTSVIRUSASSEMBLYCOMPLETENESS_PARTIAL V1reportsVirusAssemblyCompleteness = "PARTIAL"
)

var allowedV1reportsVirusAssemblyCompletenessEnumValues = []V1reportsVirusAssemblyCompleteness{
	"UNKNOWN",
	"COMPLETE",
	"PARTIAL",
}

func (v *V1reportsVirusAssemblyCompleteness) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := V1reportsVirusAssemblyCompleteness(value)
	for _, existing := range allowedV1reportsVirusAssemblyCompletenessEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid V1reportsVirusAssemblyCompleteness", value)
}

// NewV1reportsVirusAssemblyCompletenessFromValue returns a pointer to a valid V1reportsVirusAssemblyCompleteness
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewV1reportsVirusAssemblyCompletenessFromValue(v string) (*V1reportsVirusAssemblyCompleteness, error) {
	ev := V1reportsVirusAssemblyCompleteness(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for V1reportsVirusAssemblyCompleteness: valid values are %v", v, allowedV1reportsVirusAssemblyCompletenessEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v V1reportsVirusAssemblyCompleteness) IsValid() bool {
	for _, existing := range allowedV1reportsVirusAssemblyCompletenessEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to v1reportsVirusAssemblyCompleteness value
func (v V1reportsVirusAssemblyCompleteness) Ptr() *V1reportsVirusAssemblyCompleteness {
	return &v
}

type NullableV1reportsVirusAssemblyCompleteness struct {
	value *V1reportsVirusAssemblyCompleteness
	isSet bool
}

func (v NullableV1reportsVirusAssemblyCompleteness) Get() *V1reportsVirusAssemblyCompleteness {
	return v.value
}

func (v *NullableV1reportsVirusAssemblyCompleteness) Set(val *V1reportsVirusAssemblyCompleteness) {
	v.value = val
	v.isSet = true
}

func (v NullableV1reportsVirusAssemblyCompleteness) IsSet() bool {
	return v.isSet
}

func (v *NullableV1reportsVirusAssemblyCompleteness) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1reportsVirusAssemblyCompleteness(val *V1reportsVirusAssemblyCompleteness) *NullableV1reportsVirusAssemblyCompleteness {
	return &NullableV1reportsVirusAssemblyCompleteness{value: val, isSet: true}
}

func (v NullableV1reportsVirusAssemblyCompleteness) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1reportsVirusAssemblyCompleteness) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

