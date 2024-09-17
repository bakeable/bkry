package printing_order

import (
	"strings"
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var collectionPath = "PrintingOrder"
func GetCollectionPath() string {
	output := collectionPath
	return output
}

func (entity PrintingOrder) GetCollectionPath() string {
	return GetCollectionPath()
}

var documentPath = "PrintingOrder/[printingOrderID]"
func GetDocumentPath(printingOrderID string) string {
	output := documentPath
	output = strings.Replace(output, "[printingOrderID]", printingOrderID, -1)
	return output
}

func (entity PrintingOrder) GetDocumentPath() string {
    return GetDocumentPath( entity.ID)
}

func (entity PrintingOrder) Encode() map[string]interface{} {
	return Encode(entity)
}

func (entity PrintingOrder) GetID() string {
	return entity.ID
}
