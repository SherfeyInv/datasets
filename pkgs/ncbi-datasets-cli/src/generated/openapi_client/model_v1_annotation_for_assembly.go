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

// V1AnnotationForAssembly struct for V1AnnotationForAssembly
type V1AnnotationForAssembly struct {
	Name *string `json:"name,omitempty"`
	Source *string `json:"source,omitempty"`
	ReleaseNumber *string `json:"release_number,omitempty"`
	ReleaseDate *string `json:"release_date,omitempty"`
	ReportUrl *string `json:"report_url,omitempty"`
	File *[]V1AnnotationForAssemblyFile `json:"file,omitempty"`
	Stats *V1FeatureCounts `json:"stats,omitempty"`
	Busco *V1BuscoStat `json:"busco,omitempty"`
}

// NewV1AnnotationForAssembly instantiates a new V1AnnotationForAssembly object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1AnnotationForAssembly() *V1AnnotationForAssembly {
	this := V1AnnotationForAssembly{}
	return &this
}

// NewV1AnnotationForAssemblyWithDefaults instantiates a new V1AnnotationForAssembly object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1AnnotationForAssemblyWithDefaults() *V1AnnotationForAssembly {
	this := V1AnnotationForAssembly{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1AnnotationForAssembly) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AnnotationForAssembly) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1AnnotationForAssembly) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1AnnotationForAssembly) SetName(v string) {
	o.Name = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *V1AnnotationForAssembly) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AnnotationForAssembly) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *V1AnnotationForAssembly) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *V1AnnotationForAssembly) SetSource(v string) {
	o.Source = &v
}

// GetReleaseNumber returns the ReleaseNumber field value if set, zero value otherwise.
func (o *V1AnnotationForAssembly) GetReleaseNumber() string {
	if o == nil || o.ReleaseNumber == nil {
		var ret string
		return ret
	}
	return *o.ReleaseNumber
}

// GetReleaseNumberOk returns a tuple with the ReleaseNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AnnotationForAssembly) GetReleaseNumberOk() (*string, bool) {
	if o == nil || o.ReleaseNumber == nil {
		return nil, false
	}
	return o.ReleaseNumber, true
}

// HasReleaseNumber returns a boolean if a field has been set.
func (o *V1AnnotationForAssembly) HasReleaseNumber() bool {
	if o != nil && o.ReleaseNumber != nil {
		return true
	}

	return false
}

// SetReleaseNumber gets a reference to the given string and assigns it to the ReleaseNumber field.
func (o *V1AnnotationForAssembly) SetReleaseNumber(v string) {
	o.ReleaseNumber = &v
}

// GetReleaseDate returns the ReleaseDate field value if set, zero value otherwise.
func (o *V1AnnotationForAssembly) GetReleaseDate() string {
	if o == nil || o.ReleaseDate == nil {
		var ret string
		return ret
	}
	return *o.ReleaseDate
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AnnotationForAssembly) GetReleaseDateOk() (*string, bool) {
	if o == nil || o.ReleaseDate == nil {
		return nil, false
	}
	return o.ReleaseDate, true
}

// HasReleaseDate returns a boolean if a field has been set.
func (o *V1AnnotationForAssembly) HasReleaseDate() bool {
	if o != nil && o.ReleaseDate != nil {
		return true
	}

	return false
}

// SetReleaseDate gets a reference to the given string and assigns it to the ReleaseDate field.
func (o *V1AnnotationForAssembly) SetReleaseDate(v string) {
	o.ReleaseDate = &v
}

// GetReportUrl returns the ReportUrl field value if set, zero value otherwise.
func (o *V1AnnotationForAssembly) GetReportUrl() string {
	if o == nil || o.ReportUrl == nil {
		var ret string
		return ret
	}
	return *o.ReportUrl
}

// GetReportUrlOk returns a tuple with the ReportUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AnnotationForAssembly) GetReportUrlOk() (*string, bool) {
	if o == nil || o.ReportUrl == nil {
		return nil, false
	}
	return o.ReportUrl, true
}

// HasReportUrl returns a boolean if a field has been set.
func (o *V1AnnotationForAssembly) HasReportUrl() bool {
	if o != nil && o.ReportUrl != nil {
		return true
	}

	return false
}

// SetReportUrl gets a reference to the given string and assigns it to the ReportUrl field.
func (o *V1AnnotationForAssembly) SetReportUrl(v string) {
	o.ReportUrl = &v
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *V1AnnotationForAssembly) GetFile() []V1AnnotationForAssemblyFile {
	if o == nil || o.File == nil {
		var ret []V1AnnotationForAssemblyFile
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AnnotationForAssembly) GetFileOk() (*[]V1AnnotationForAssemblyFile, bool) {
	if o == nil || o.File == nil {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *V1AnnotationForAssembly) HasFile() bool {
	if o != nil && o.File != nil {
		return true
	}

	return false
}

// SetFile gets a reference to the given []V1AnnotationForAssemblyFile and assigns it to the File field.
func (o *V1AnnotationForAssembly) SetFile(v []V1AnnotationForAssemblyFile) {
	o.File = &v
}

// GetStats returns the Stats field value if set, zero value otherwise.
func (o *V1AnnotationForAssembly) GetStats() V1FeatureCounts {
	if o == nil || o.Stats == nil {
		var ret V1FeatureCounts
		return ret
	}
	return *o.Stats
}

// GetStatsOk returns a tuple with the Stats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AnnotationForAssembly) GetStatsOk() (*V1FeatureCounts, bool) {
	if o == nil || o.Stats == nil {
		return nil, false
	}
	return o.Stats, true
}

// HasStats returns a boolean if a field has been set.
func (o *V1AnnotationForAssembly) HasStats() bool {
	if o != nil && o.Stats != nil {
		return true
	}

	return false
}

// SetStats gets a reference to the given V1FeatureCounts and assigns it to the Stats field.
func (o *V1AnnotationForAssembly) SetStats(v V1FeatureCounts) {
	o.Stats = &v
}

// GetBusco returns the Busco field value if set, zero value otherwise.
func (o *V1AnnotationForAssembly) GetBusco() V1BuscoStat {
	if o == nil || o.Busco == nil {
		var ret V1BuscoStat
		return ret
	}
	return *o.Busco
}

// GetBuscoOk returns a tuple with the Busco field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AnnotationForAssembly) GetBuscoOk() (*V1BuscoStat, bool) {
	if o == nil || o.Busco == nil {
		return nil, false
	}
	return o.Busco, true
}

// HasBusco returns a boolean if a field has been set.
func (o *V1AnnotationForAssembly) HasBusco() bool {
	if o != nil && o.Busco != nil {
		return true
	}

	return false
}

// SetBusco gets a reference to the given V1BuscoStat and assigns it to the Busco field.
func (o *V1AnnotationForAssembly) SetBusco(v V1BuscoStat) {
	o.Busco = &v
}

func (o V1AnnotationForAssembly) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil  {
		toSerialize["name"] = o.Name
	}
	if o.Source != nil  {
		toSerialize["source"] = o.Source
	}
	if o.ReleaseNumber != nil  {
		toSerialize["release_number"] = o.ReleaseNumber
	}
	if o.ReleaseDate != nil  {
		toSerialize["release_date"] = o.ReleaseDate
	}
	if o.ReportUrl != nil  {
		toSerialize["report_url"] = o.ReportUrl
	}
	if o.File != nil && len(o.GetFile()) > 0  {
		toSerialize["file"] = o.File
	}
	if o.Stats != nil  {
		toSerialize["stats"] = o.Stats
	}
	if o.Busco != nil  {
		toSerialize["busco"] = o.Busco
	}
	return json.Marshal(toSerialize)
}

type NullableV1AnnotationForAssembly struct {
	value *V1AnnotationForAssembly
	isSet bool
}

func (v NullableV1AnnotationForAssembly) Get() *V1AnnotationForAssembly {
	return v.value
}

func (v *NullableV1AnnotationForAssembly) Set(val *V1AnnotationForAssembly) {
	v.value = val
	v.isSet = true
}

func (v NullableV1AnnotationForAssembly) IsSet() bool {
	return v.isSet
}

func (v *NullableV1AnnotationForAssembly) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1AnnotationForAssembly(val *V1AnnotationForAssembly) *NullableV1AnnotationForAssembly {
	return &NullableV1AnnotationForAssembly{value: val, isSet: true}
}

func (v NullableV1AnnotationForAssembly) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1AnnotationForAssembly) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


