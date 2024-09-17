package packing_slip

import (
	"strings"
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var collectionPath = "PackingSlip"
func GetCollectionPath() string {
	output := collectionPath
	return output
}

func (entity PackingSlip) GetCollectionPath() string {
	return GetCollectionPath()
}

var documentPath = "PackingSlip/[packingSlipID]"
func GetDocumentPath(packingSlipID string) string {
	output := documentPath
	output = strings.Replace(output, "[packingSlipID]", packingSlipID, -1)
	return output
}

func (entity PackingSlip) GetDocumentPath() string {
    return GetDocumentPath( entity.ID)
}

func (entity PackingSlip) Encode() map[string]interface{} {
	return Encode(entity)
}

func (entity PackingSlip) GetID() string {
	return entity.ID
}
