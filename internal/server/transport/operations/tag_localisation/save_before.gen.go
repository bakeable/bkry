package tag_localisation_operations

import (
	tag_localisation "github.com/bakeable/bkry/internal/server/models/entities/tag_localisation"
)

func beforeSave(tagID string, entity tag_localisation.TagLocalisation, editorID *string) tag_localisation.TagLocalisation {
	// Return TagLocalisation
	return entity
}