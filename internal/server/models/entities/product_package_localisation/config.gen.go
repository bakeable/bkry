package product_package_localisation

import (
	"strings"
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var collectionPath = "product_packages/[productPackageID]/product_package_localisations"
func GetCollectionPath(productPackageID string) string {
	output := collectionPath
	output = strings.Replace(output, "[productPackageID]", productPackageID, -1)
	return output
}

func (entity ProductPackageLocalisation) GetCollectionPath() string {
	return GetCollectionPath(entity.ProductPackageID)
}

var documentPath = "product_packages/[productPackageID]/product_package_localisations/[productPackageLocalisationID]"
func GetDocumentPath(productPackageID string, productPackageLocalisationID string) string {
	output := documentPath
	output = strings.Replace(output, "[productPackageID]", productPackageID, -1)
	output = strings.Replace(output, "[productPackageLocalisationID]", productPackageLocalisationID, -1)
	return output
}

func (entity ProductPackageLocalisation) GetDocumentPath() string {
    return GetDocumentPath(entity.ProductPackageID,  entity.ID)
}

func (entity ProductPackageLocalisation) ToMap() map[string]interface{} {
	return ToMap(entity)
}

func (entity ProductPackageLocalisation) GetID() string {
	return entity.ID
}
