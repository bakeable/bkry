package delivery_entry

import (
	"strings"
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var collectionPath = "DeliveryEntry"
func GetCollectionPath() string {
	output := collectionPath
	return output
}

func (entity DeliveryEntry) GetCollectionPath() string {
	return GetCollectionPath()
}

var documentPath = "DeliveryEntry/[deliveryEntryID]"
func GetDocumentPath(deliveryEntryID string) string {
	output := documentPath
	output = strings.Replace(output, "[deliveryEntryID]", deliveryEntryID, -1)
	return output
}

func (entity DeliveryEntry) GetDocumentPath() string {
    return GetDocumentPath( entity.ID)
}

func (entity DeliveryEntry) Encode() map[string]interface{} {
	return Encode(entity)
}

func (entity DeliveryEntry) GetID() string {
	return entity.ID
}
