package tag_localisation_operations

import (
	tag_localisation "github.com/bakeable/bkry/internal/server/models/entities/tag_localisation"
)

func afterGet(tagID string, tagLocalisationID string, entity tag_localisation.TagLocalisation) tag_localisation.TagLocalisation {
	return entity
}