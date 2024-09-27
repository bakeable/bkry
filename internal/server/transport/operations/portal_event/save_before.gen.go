package portal_event_operations

import (
	portal_event "github.com/bakeable/bkry/internal/server/models/entities/portal_event"
)

func beforeSave(entity portal_event.PortalEvent, editorID *string) portal_event.PortalEvent {
	// Return PortalEvent
	return entity
}