package airline_pricing

import (
	"strings"
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var collectionPath = "AirlinePricing"
func GetCollectionPath() string {
	output := collectionPath
	return output
}

func (entity AirlinePricing) GetCollectionPath() string {
	return GetCollectionPath()
}

var documentPath = "AirlinePricing/[airlinePricingID]"
func GetDocumentPath(airlinePricingID string) string {
	output := documentPath
	output = strings.Replace(output, "[airlinePricingID]", airlinePricingID, -1)
	return output
}

func (entity AirlinePricing) GetDocumentPath() string {
    return GetDocumentPath( entity.ID)
}

func (entity AirlinePricing) Encode() map[string]interface{} {
	return Encode(entity)
}

func (entity AirlinePricing) GetID() string {
	return entity.ID
}
