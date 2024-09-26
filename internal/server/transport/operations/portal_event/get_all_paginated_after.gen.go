package portal_event_operations

import (
	portal_event "github.com/bakeable/bkry/data/entities/portal_event"
)


func afterGetAllPaginated(entities []portal_event.PortalEvent, pageSize int, pageNumber int, orderBy string, ascending bool) []portal_event.PortalEvent {
	return entities
}