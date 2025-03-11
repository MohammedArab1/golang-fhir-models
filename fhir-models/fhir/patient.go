// Copyright 2019 - 2022 The Samply Community
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fhir

import "encoding/json"

// THIS FILE IS GENERATED BY https://github.com/samply/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// Patient is documented here http://hl7.org/fhir/StructureDefinition/Patient
type Patient struct {
	Id                   *string                `bson:"id,omitempty" json:"id,omitempty"`
	Meta                 *Meta                  `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules        *string                `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language             *string                `bson:"language,omitempty" json:"language,omitempty"`
	Text                 *Narrative             `bson:"text,omitempty" json:"text,omitempty"`
	Extension            []Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier           []Identifier           `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Active               *bool                  `bson:"active,omitempty" json:"active,omitempty"`
	Name                 []HumanName            `bson:"name,omitempty" json:"name,omitempty"`
	Telecom              []ContactPoint         `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Gender               *AdministrativeGender  `bson:"gender,omitempty" json:"gender,omitempty"`
	BirthDate            *string                `bson:"birthDate,omitempty" json:"birthDate,omitempty"`
	DeceasedBoolean      *bool                  `bson:"deceasedBoolean,omitempty" json:"deceasedBoolean,omitempty"`
	DeceasedDateTime     *string                `bson:"deceasedDateTime,omitempty" json:"deceasedDateTime,omitempty"`
	Address              []Address              `bson:"address,omitempty" json:"address,omitempty"`
	MaritalStatus        *CodeableConcept       `bson:"maritalStatus,omitempty" json:"maritalStatus,omitempty"`
	MultipleBirthBoolean *bool                  `bson:"multipleBirthBoolean,omitempty" json:"multipleBirthBoolean,omitempty"`
	MultipleBirthInteger *int                   `bson:"multipleBirthInteger,omitempty" json:"multipleBirthInteger,omitempty"`
	Photo                []Attachment           `bson:"photo,omitempty" json:"photo,omitempty"`
	Contact              []PatientContact       `bson:"contact,omitempty" json:"contact,omitempty"`
	Communication        []PatientCommunication `bson:"communication,omitempty" json:"communication,omitempty"`
	GeneralPractitioner  []Reference            `bson:"generalPractitioner,omitempty" json:"generalPractitioner,omitempty"`
	ManagingOrganization *Reference             `bson:"managingOrganization,omitempty" json:"managingOrganization,omitempty"`
	Link                 []PatientLink          `bson:"link,omitempty" json:"link,omitempty"`
}
type PatientContact struct {
	Id                *string               `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Relationship      []CodeableConcept     `bson:"relationship,omitempty" json:"relationship,omitempty"`
	Name              *HumanName            `bson:"name,omitempty" json:"name,omitempty"`
	Telecom           []ContactPoint        `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Address           *Address              `bson:"address,omitempty" json:"address,omitempty"`
	Gender            *AdministrativeGender `bson:"gender,omitempty" json:"gender,omitempty"`
	Organization      *Reference            `bson:"organization,omitempty" json:"organization,omitempty"`
	Period            *Period               `bson:"period,omitempty" json:"period,omitempty"`
}
type PatientCommunication struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Language          CodeableConcept `bson:"language,omitempty" json:"language,omitempty"`
	Preferred         *bool           `bson:"preferred,omitempty" json:"preferred,omitempty"`
}
type PatientLink struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Other             Reference   `bson:"other,omitempty" json:"other,omitempty"`
	Type              LinkType    `bson:"type,omitempty" json:"type,omitempty"`
}
type OtherPatient Patient

// MarshalJSON marshals the given Patient as JSON into a byte slice
func (r Patient) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPatient
		ResourceType string `json:"resourceType"`
	}{
		OtherPatient: OtherPatient(r),
		ResourceType: "Patient",
	})
}

// UnmarshalPatient unmarshals a Patient.
func UnmarshalPatient(b []byte) (Patient, error) {
	var patient Patient
	if err := json.Unmarshal(b, &patient); err != nil {
		return patient, err
	}
	return patient, nil
}
