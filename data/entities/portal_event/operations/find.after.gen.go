package portal_event_operations

import (
	portal_event "github.com/bakeable/bkry/data/entities/portal_event"
)

func afterFind(entity portal_event.PortalEvent) portal_event.PortalEvent {
	return entity
}