package product_package

import (
	"strings"
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var collectionPath = "product_packages"
func GetCollectionPath() string {
	output := collectionPath
	return output
}

func (entity ProductPackage) GetCollectionPath() string {
	return GetCollectionPath()
}

var documentPath = "product_packages/[productPackageID]"
func GetDocumentPath(productPackageID string) string {
	output := documentPath
	output = strings.Replace(output, "[productPackageID]", productPackageID, -1)
	return output
}

func (entity ProductPackage) GetDocumentPath() string {
    return GetDocumentPath( entity.ID)
}

func (entity ProductPackage) ToMap() map[string]interface{} {
	return ToMap(entity)
}

func (entity ProductPackage) GetID() string {
	return entity.ID
}
