package portal_event_operations

import (
	portal_event "github.com/bakeable/bkry/internal/server/models/entities/portal_event"
)

func afterSave(entity portal_event.PortalEvent, editorID *string) {}