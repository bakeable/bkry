package airline_order_group

import (
	"strings"
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var collectionPath = "AirlineOrderGroup"
func GetCollectionPath() string {
	output := collectionPath
	return output
}

func (entity AirlineOrderGroup) GetCollectionPath() string {
	return GetCollectionPath()
}

var documentPath = "AirlineOrderGroup/[airlineOrderGroupID]"
func GetDocumentPath(airlineOrderGroupID string) string {
	output := documentPath
	output = strings.Replace(output, "[airlineOrderGroupID]", airlineOrderGroupID, -1)
	return output
}

func (entity AirlineOrderGroup) GetDocumentPath() string {
    return GetDocumentPath( entity.ID)
}

func (entity AirlineOrderGroup) Encode() map[string]interface{} {
	return Encode(entity)
}

func (entity AirlineOrderGroup) GetID() string {
	return entity.ID
}
