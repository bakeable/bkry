package tag_localisation

import (
	"strings"
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var collectionPath = "tags/[tagID]/tag_localisations"
func GetCollectionPath(tagID string) string {
	output := collectionPath
	output = strings.Replace(output, "[tagID]", tagID, -1)
	return output
}

func (entity TagLocalisation) GetCollectionPath() string {
	return GetCollectionPath(entity.TagID)
}

var documentPath = "tags/[tagID]/tag_localisations/[tagLocalisationID]"
func GetDocumentPath(tagID string, tagLocalisationID string) string {
	output := documentPath
	output = strings.Replace(output, "[tagID]", tagID, -1)
	output = strings.Replace(output, "[tagLocalisationID]", tagLocalisationID, -1)
	return output
}

func (entity TagLocalisation) GetDocumentPath() string {
    return GetDocumentPath(entity.TagID,  entity.ID)
}

func (entity TagLocalisation) ToMap() map[string]interface{} {
	return ToMap(entity)
}

func (entity TagLocalisation) GetID() string {
	return entity.ID
}
