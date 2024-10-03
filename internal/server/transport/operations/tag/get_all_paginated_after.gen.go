package tag_operations

import (
	tag "github.com/bakeable/bkry/internal/server/models/entities/tag"
)


func afterGetAllPaginated(entities []tag.Tag, pageSize int, pageNumber int, orderBy string, ascending bool) []tag.Tag {
	return entities
}