package tag_localisation_operations

import (
	tag_localisation "github.com/bakeable/bkry/internal/server/models/entities/tag_localisation"
)

func afterSave(tagID string, entity tag_localisation.TagLocalisation, editorID *string) {}