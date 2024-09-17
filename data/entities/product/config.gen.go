package product

import (
	"strings"
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var collectionPath = "Product"
func GetCollectionPath() string {
	output := collectionPath
	return output
}

func (entity Product) GetCollectionPath() string {
	return GetCollectionPath()
}

var documentPath = "Product/[productID]"
func GetDocumentPath(productID string) string {
	output := documentPath
	output = strings.Replace(output, "[productID]", productID, -1)
	return output
}

func (entity Product) GetDocumentPath() string {
    return GetDocumentPath( entity.ID)
}

func (entity Product) Encode() map[string]interface{} {
	return Encode(entity)
}

func (entity Product) GetID() string {
	return entity.ID
}
