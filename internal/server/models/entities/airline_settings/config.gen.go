package airline_settings

import (
	"strings"
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var collectionPath = "AirlineSettings"
func GetCollectionPath() string {
	output := collectionPath
	return output
}

func (entity AirlineSettings) GetCollectionPath() string {
	return GetCollectionPath()
}

var documentPath = "AirlineSettings/[airlineSettingsID]"
func GetDocumentPath(airlineSettingsID string) string {
	output := documentPath
	output = strings.Replace(output, "[airlineSettingsID]", airlineSettingsID, -1)
	return output
}

func (entity AirlineSettings) GetDocumentPath() string {
    return GetDocumentPath( entity.ID)
}

func (entity AirlineSettings) Encode() map[string]interface{} {
	return Encode(entity)
}

func (entity AirlineSettings) GetID() string {
	return entity.ID
}
