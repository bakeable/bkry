package category_localisation

import (
	"strings"
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var collectionPath = "categories/[categoryID]/category_localisations"
func GetCollectionPath(categoryID string) string {
	output := collectionPath
	output = strings.Replace(output, "[categoryID]", categoryID, -1)
	return output
}

func (entity CategoryLocalisation) GetCollectionPath() string {
	return GetCollectionPath(entity.CategoryID)
}

var documentPath = "categories/[categoryID]/category_localisations/[categoryLocalisationID]"
func GetDocumentPath(categoryID string, categoryLocalisationID string) string {
	output := documentPath
	output = strings.Replace(output, "[categoryID]", categoryID, -1)
	output = strings.Replace(output, "[categoryLocalisationID]", categoryLocalisationID, -1)
	return output
}

func (entity CategoryLocalisation) GetDocumentPath() string {
    return GetDocumentPath(entity.CategoryID,  entity.ID)
}

func (entity CategoryLocalisation) ToMap() map[string]interface{} {
	return ToMap(entity)
}

func (entity CategoryLocalisation) GetID() string {
	return entity.ID
}
