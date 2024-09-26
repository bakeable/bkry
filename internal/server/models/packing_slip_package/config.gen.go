package packing_slip_package

import (
	"strings"
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var collectionPath = "PackingSlipPackage"
func GetCollectionPath() string {
	output := collectionPath
	return output
}

func (entity PackingSlipPackage) GetCollectionPath() string {
	return GetCollectionPath()
}

var documentPath = "PackingSlipPackage/[packingSlipPackageID]"
func GetDocumentPath(packingSlipPackageID string) string {
	output := documentPath
	output = strings.Replace(output, "[packingSlipPackageID]", packingSlipPackageID, -1)
	return output
}

func (entity PackingSlipPackage) GetDocumentPath() string {
    return GetDocumentPath( entity.ID)
}

func (entity PackingSlipPackage) Encode() map[string]interface{} {
	return Encode(entity)
}

func (entity PackingSlipPackage) GetID() string {
	return entity.ID
}
