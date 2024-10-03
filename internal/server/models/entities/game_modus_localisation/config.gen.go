package game_modus_localisation

import (
	"strings"
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var collectionPath = "game_modus/[gameModusID]/game_modus_localisations"
func GetCollectionPath(gameModusID string) string {
	output := collectionPath
	output = strings.Replace(output, "[gameModusID]", gameModusID, -1)
	return output
}

func (entity GameModusLocalisation) GetCollectionPath() string {
	return GetCollectionPath(entity.GameModusID)
}

var documentPath = "game_modus/[gameModusID]/game_modus_localisations/[gameModusLocalisationID]"
func GetDocumentPath(gameModusID string, gameModusLocalisationID string) string {
	output := documentPath
	output = strings.Replace(output, "[gameModusID]", gameModusID, -1)
	output = strings.Replace(output, "[gameModusLocalisationID]", gameModusLocalisationID, -1)
	return output
}

func (entity GameModusLocalisation) GetDocumentPath() string {
    return GetDocumentPath(entity.GameModusID,  entity.ID)
}

func (entity GameModusLocalisation) ToMap() map[string]interface{} {
	return ToMap(entity)
}

func (entity GameModusLocalisation) GetID() string {
	return entity.ID
}
