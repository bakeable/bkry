package price_configuration

import (
	"strings"
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var collectionPath = "PriceConfiguration"
func GetCollectionPath() string {
	output := collectionPath
	return output
}

func (entity PriceConfiguration) GetCollectionPath() string {
	return GetCollectionPath()
}

var documentPath = "PriceConfiguration/[priceConfigurationID]"
func GetDocumentPath(priceConfigurationID string) string {
	output := documentPath
	output = strings.Replace(output, "[priceConfigurationID]", priceConfigurationID, -1)
	return output
}

func (entity PriceConfiguration) GetDocumentPath() string {
    return GetDocumentPath( entity.ID)
}

func (entity PriceConfiguration) Encode() map[string]interface{} {
	return Encode(entity)
}

func (entity PriceConfiguration) GetID() string {
	return entity.ID
}
