package tag_operations

import (
	tag "github.com/bakeable/bkry/internal/server/models/entities/tag"
)


func afterGetAll(entities []tag.Tag) []tag.Tag {
	return entities
}