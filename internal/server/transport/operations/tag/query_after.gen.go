package tag_operations

import (
	tag "github.com/bakeable/bkry/internal/server/models/entities/tag"
)


func afterQuery(entities []tag.Tag) []tag.Tag {
	return entities
}