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

// Flag is documented here http://hl7.org/fhir/StructureDefinition/Flag
type Flag struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta             `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string           `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string           `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative        `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            FlagStatus        `bson:"status,omitempty" json:"status,omitempty"`
	Category          []CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	Code              CodeableConcept   `bson:"code,omitempty" json:"code,omitempty"`
	Subject           Reference         `bson:"subject,omitempty" json:"subject,omitempty"`
	Period            *Period           `bson:"period,omitempty" json:"period,omitempty"`
	Encounter         *Reference        `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Author            *Reference        `bson:"author,omitempty" json:"author,omitempty"`
}
type OtherFlag Flag

// MarshalJSON marshals the given Flag as JSON into a byte slice
func (r Flag) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherFlag
		ResourceType string `json:"resourceType"`
	}{
		OtherFlag:    OtherFlag(r),
		ResourceType: "Flag",
	})
}

// UnmarshalFlag unmarshals a Flag.
func UnmarshalFlag(b []byte) (Flag, error) {
	var flag Flag
	if err := json.Unmarshal(b, &flag); err != nil {
		return flag, err
	}
	return flag, nil
}
