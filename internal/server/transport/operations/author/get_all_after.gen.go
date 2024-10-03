package author_operations

import (
	author "github.com/bakeable/bkry/internal/server/models/entities/author"
)


func afterGetAll(entities []author.Author) []author.Author {
	return entities
}