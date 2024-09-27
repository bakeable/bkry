package airline_order

import (
	"strings"
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var collectionPath = "AirlineOrder"
func GetCollectionPath() string {
	output := collectionPath
	return output
}

func (entity AirlineOrder) GetCollectionPath() string {
	return GetCollectionPath()
}

var documentPath = "AirlineOrder/[airlineOrderID]"
func GetDocumentPath(airlineOrderID string) string {
	output := documentPath
	output = strings.Replace(output, "[airlineOrderID]", airlineOrderID, -1)
	return output
}

func (entity AirlineOrder) GetDocumentPath() string {
    return GetDocumentPath( entity.ID)
}

func (entity AirlineOrder) Encode() map[string]interface{} {
	return Encode(entity)
}

func (entity AirlineOrder) GetID() string {
	return entity.ID
}
