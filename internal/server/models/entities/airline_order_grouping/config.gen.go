package airline_order_grouping

import (
	"strings"
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var collectionPath = "AirlineOrderGrouping"
func GetCollectionPath() string {
	output := collectionPath
	return output
}

func (entity AirlineOrderGrouping) GetCollectionPath() string {
	return GetCollectionPath()
}

var documentPath = "AirlineOrderGrouping/[airlineOrderGroupingID]"
func GetDocumentPath(airlineOrderGroupingID string) string {
	output := documentPath
	output = strings.Replace(output, "[airlineOrderGroupingID]", airlineOrderGroupingID, -1)
	return output
}

func (entity AirlineOrderGrouping) GetDocumentPath() string {
    return GetDocumentPath( entity.ID)
}

func (entity AirlineOrderGrouping) Encode() map[string]interface{} {
	return Encode(entity)
}

func (entity AirlineOrderGrouping) GetID() string {
	return entity.ID
}
