package portal_event_operations

import (
	portal_event "github.com/bakeable/bkry/internal/server/models/entities/portal_event"
)

func afterGet(portalEventID string, entity portal_event.PortalEvent) portal_event.PortalEvent {
	return entity
}