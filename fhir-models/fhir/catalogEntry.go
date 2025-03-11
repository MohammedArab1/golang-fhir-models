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

// CatalogEntry is documented here http://hl7.org/fhir/StructureDefinition/CatalogEntry
type CatalogEntry struct {
	Id                       *string                    `bson:"id,omitempty" json:"id,omitempty"`
	Meta                     *Meta                      `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules            *string                    `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language                 *string                    `bson:"language,omitempty" json:"language,omitempty"`
	Text                     *Narrative                 `bson:"text,omitempty" json:"text,omitempty"`
	Extension                []Extension                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension        []Extension                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier               []Identifier               `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Type                     *CodeableConcept           `bson:"type,omitempty" json:"type,omitempty"`
	Orderable                bool                       `bson:"orderable,omitempty" json:"orderable,omitempty"`
	ReferencedItem           Reference                  `bson:"referencedItem,omitempty" json:"referencedItem,omitempty"`
	AdditionalIdentifier     []Identifier               `bson:"additionalIdentifier,omitempty" json:"additionalIdentifier,omitempty"`
	Classification           []CodeableConcept          `bson:"classification,omitempty" json:"classification,omitempty"`
	Status                   *PublicationStatus         `bson:"status,omitempty" json:"status,omitempty"`
	ValidityPeriod           *Period                    `bson:"validityPeriod,omitempty" json:"validityPeriod,omitempty"`
	ValidTo                  *string                    `bson:"validTo,omitempty" json:"validTo,omitempty"`
	LastUpdated              *string                    `bson:"lastUpdated,omitempty" json:"lastUpdated,omitempty"`
	AdditionalCharacteristic []CodeableConcept          `bson:"additionalCharacteristic,omitempty" json:"additionalCharacteristic,omitempty"`
	AdditionalClassification []CodeableConcept          `bson:"additionalClassification,omitempty" json:"additionalClassification,omitempty"`
	RelatedEntry             []CatalogEntryRelatedEntry `bson:"relatedEntry,omitempty" json:"relatedEntry,omitempty"`
}
type CatalogEntryRelatedEntry struct {
	Id                *string                  `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension              `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension              `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Relationtype      CatalogEntryRelationType `bson:"relationtype,omitempty" json:"relationtype,omitempty"`
	Item              Reference                `bson:"item,omitempty" json:"item,omitempty"`
}
type OtherCatalogEntry CatalogEntry

// MarshalJSON marshals the given CatalogEntry as JSON into a byte slice
func (r CatalogEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCatalogEntry
		ResourceType string `json:"resourceType"`
	}{
		OtherCatalogEntry: OtherCatalogEntry(r),
		ResourceType:      "CatalogEntry",
	})
}

// UnmarshalCatalogEntry unmarshals a CatalogEntry.
func UnmarshalCatalogEntry(b []byte) (CatalogEntry, error) {
	var catalogEntry CatalogEntry
	if err := json.Unmarshal(b, &catalogEntry); err != nil {
		return catalogEntry, err
	}
	return catalogEntry, nil
}
