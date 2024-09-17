package quotation

import (
	"strings"
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var collectionPath = "Quotation"
func GetCollectionPath() string {
	output := collectionPath
	return output
}

func (entity Quotation) GetCollectionPath() string {
	return GetCollectionPath()
}

var documentPath = "Quotation/[quotationID]"
func GetDocumentPath(quotationID string) string {
	output := documentPath
	output = strings.Replace(output, "[quotationID]", quotationID, -1)
	return output
}

func (entity Quotation) GetDocumentPath() string {
    return GetDocumentPath( entity.ID)
}

func (entity Quotation) Encode() map[string]interface{} {
	return Encode(entity)
}

func (entity Quotation) GetID() string {
	return entity.ID
}
