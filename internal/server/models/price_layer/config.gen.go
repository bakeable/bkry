package price_layer

import (
	"strings"
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var collectionPath = "PriceLayer"
func GetCollectionPath() string {
	output := collectionPath
	return output
}

func (entity PriceLayer) GetCollectionPath() string {
	return GetCollectionPath()
}

var documentPath = "PriceLayer/[priceLayerID]"
func GetDocumentPath(priceLayerID string) string {
	output := documentPath
	output = strings.Replace(output, "[priceLayerID]", priceLayerID, -1)
	return output
}

func (entity PriceLayer) GetDocumentPath() string {
    return GetDocumentPath( entity.ID)
}

func (entity PriceLayer) Encode() map[string]interface{} {
	return Encode(entity)
}

func (entity PriceLayer) GetID() string {
	return entity.ID
}
