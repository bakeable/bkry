package tag_localisation_operations

import (
	tag_localisation "github.com/bakeable/bkry/internal/server/models/entities/tag_localisation"
)


func afterQueryGroup(entities []tag_localisation.TagLocalisation) []tag_localisation.TagLocalisation {
	return entities
}