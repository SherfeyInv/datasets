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

// V1VirusTableField the model 'V1VirusTableField'
type V1VirusTableField string

// List of v1VirusTableField
const (
	V1VIRUSTABLEFIELD_UNSPECIFIED V1VirusTableField = "unspecified"
	V1VIRUSTABLEFIELD_NUCLEOTIDE_ACCESSION V1VirusTableField = "nucleotide_accession"
	V1VIRUSTABLEFIELD_SPECIES_TAX_ID V1VirusTableField = "species_tax_id"
	V1VIRUSTABLEFIELD_SPECIES_NAME V1VirusTableField = "species_name"
	V1VIRUSTABLEFIELD_GENUS V1VirusTableField = "genus"
	V1VIRUSTABLEFIELD_FAMILY V1VirusTableField = "family"
	V1VIRUSTABLEFIELD_NUCLEOTIDE_LENGTH V1VirusTableField = "nucleotide_length"
	V1VIRUSTABLEFIELD_ISOLATE_NAME V1VirusTableField = "isolate_name"
	V1VIRUSTABLEFIELD_SEQUENCE_TYPE V1VirusTableField = "sequence_type"
	V1VIRUSTABLEFIELD_NUC_COMPLETENESS V1VirusTableField = "nuc_completeness"
	V1VIRUSTABLEFIELD_GEO_LOCATION V1VirusTableField = "geo_location"
	V1VIRUSTABLEFIELD_US_STATE V1VirusTableField = "us_state"
	V1VIRUSTABLEFIELD_HOST_NAME V1VirusTableField = "host_name"
	V1VIRUSTABLEFIELD_HOST_TAX_ID V1VirusTableField = "host_tax_id"
	V1VIRUSTABLEFIELD_COLLECTION_DATE V1VirusTableField = "collection_date"
	V1VIRUSTABLEFIELD_BIOPROJECT V1VirusTableField = "bioproject"
	V1VIRUSTABLEFIELD_BIOSAMPLE V1VirusTableField = "biosample"
	V1VIRUSTABLEFIELD_POLYPROTEIN_NAME V1VirusTableField = "polyprotein_name"
	V1VIRUSTABLEFIELD_PROTEIN_NAME V1VirusTableField = "protein_name"
	V1VIRUSTABLEFIELD_PROTEIN_ACCESSION V1VirusTableField = "protein_accession"
	V1VIRUSTABLEFIELD_PROTEIN_SYNONYM V1VirusTableField = "protein_synonym"
	V1VIRUSTABLEFIELD_CDS_SPAN V1VirusTableField = "cds_span"
)

var allowedV1VirusTableFieldEnumValues = []V1VirusTableField{
	"unspecified",
	"nucleotide_accession",
	"species_tax_id",
	"species_name",
	"genus",
	"family",
	"nucleotide_length",
	"isolate_name",
	"sequence_type",
	"nuc_completeness",
	"geo_location",
	"us_state",
	"host_name",
	"host_tax_id",
	"collection_date",
	"bioproject",
	"biosample",
	"polyprotein_name",
	"protein_name",
	"protein_accession",
	"protein_synonym",
	"cds_span",
}

func (v *V1VirusTableField) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := V1VirusTableField(value)
	for _, existing := range allowedV1VirusTableFieldEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid V1VirusTableField", value)
}

// NewV1VirusTableFieldFromValue returns a pointer to a valid V1VirusTableField
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewV1VirusTableFieldFromValue(v string) (*V1VirusTableField, error) {
	ev := V1VirusTableField(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for V1VirusTableField: valid values are %v", v, allowedV1VirusTableFieldEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v V1VirusTableField) IsValid() bool {
	for _, existing := range allowedV1VirusTableFieldEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to v1VirusTableField value
func (v V1VirusTableField) Ptr() *V1VirusTableField {
	return &v
}

type NullableV1VirusTableField struct {
	value *V1VirusTableField
	isSet bool
}

func (v NullableV1VirusTableField) Get() *V1VirusTableField {
	return v.value
}

func (v *NullableV1VirusTableField) Set(val *V1VirusTableField) {
	v.value = val
	v.isSet = true
}

func (v NullableV1VirusTableField) IsSet() bool {
	return v.isSet
}

func (v *NullableV1VirusTableField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1VirusTableField(val *V1VirusTableField) *NullableV1VirusTableField {
	return &NullableV1VirusTableField{value: val, isSet: true}
}

func (v NullableV1VirusTableField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1VirusTableField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

