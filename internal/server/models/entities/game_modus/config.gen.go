package game_modus

import (
	"strings"
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var collectionPath = "game_modus"
func GetCollectionPath() string {
	output := collectionPath
	return output
}

func (entity GameModus) GetCollectionPath() string {
	return GetCollectionPath()
}

var documentPath = "game_modus/[gameModusID]"
func GetDocumentPath(gameModusID string) string {
	output := documentPath
	output = strings.Replace(output, "[gameModusID]", gameModusID, -1)
	return output
}

func (entity GameModus) GetDocumentPath() string {
    return GetDocumentPath( entity.ID)
}

func (entity GameModus) ToMap() map[string]interface{} {
	return ToMap(entity)
}

func (entity GameModus) GetID() string {
	return entity.ID
}
