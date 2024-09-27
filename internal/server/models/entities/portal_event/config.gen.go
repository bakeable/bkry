package portal_event

import (
	"strings"
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var collectionPath = "PortalEvent"
func GetCollectionPath() string {
	output := collectionPath
	return output
}

func (entity PortalEvent) GetCollectionPath() string {
	return GetCollectionPath()
}

var documentPath = "PortalEvent/[portalEventID]"
func GetDocumentPath(portalEventID string) string {
	output := documentPath
	output = strings.Replace(output, "[portalEventID]", portalEventID, -1)
	return output
}

func (entity PortalEvent) GetDocumentPath() string {
    return GetDocumentPath( entity.ID)
}

func (entity PortalEvent) Encode() map[string]interface{} {
	return Encode(entity)
}

func (entity PortalEvent) GetID() string {
	return entity.ID
}
