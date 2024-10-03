package tag_localisation_operations

import (
	tag_localisation "github.com/bakeable/bkry/internal/server/models/entities/tag_localisation"
)


func afterGetAllPaginated(tagID string, entities []tag_localisation.TagLocalisation, pageSize int, pageNumber int, orderBy string, ascending bool) []tag_localisation.TagLocalisation {
	return entities
}