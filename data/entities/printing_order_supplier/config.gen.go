package printing_order_supplier

import (
	"strings"
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var collectionPath = "PrintingOrderSupplier"
func GetCollectionPath() string {
	output := collectionPath
	return output
}

func (entity PrintingOrderSupplier) GetCollectionPath() string {
	return GetCollectionPath()
}

var documentPath = "PrintingOrderSupplier/[printingOrderSupplierID]"
func GetDocumentPath(printingOrderSupplierID string) string {
	output := documentPath
	output = strings.Replace(output, "[printingOrderSupplierID]", printingOrderSupplierID, -1)
	return output
}

func (entity PrintingOrderSupplier) GetDocumentPath() string {
    return GetDocumentPath( entity.ID)
}

func (entity PrintingOrderSupplier) Encode() map[string]interface{} {
	return Encode(entity)
}

func (entity PrintingOrderSupplier) GetID() string {
	return entity.ID
}
