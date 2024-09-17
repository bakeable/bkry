package printing_order_item

import (
	"strings"
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var collectionPath = "PrintingOrderItem"
func GetCollectionPath() string {
	output := collectionPath
	return output
}

func (entity PrintingOrderItem) GetCollectionPath() string {
	return GetCollectionPath()
}

var documentPath = "PrintingOrderItem/[printingOrderItemID]"
func GetDocumentPath(printingOrderItemID string) string {
	output := documentPath
	output = strings.Replace(output, "[printingOrderItemID]", printingOrderItemID, -1)
	return output
}

func (entity PrintingOrderItem) GetDocumentPath() string {
    return GetDocumentPath( entity.ID)
}

func (entity PrintingOrderItem) Encode() map[string]interface{} {
	return Encode(entity)
}

func (entity PrintingOrderItem) GetID() string {
	return entity.ID
}
