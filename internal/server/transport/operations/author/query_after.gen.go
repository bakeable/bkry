package author_operations

import (
	author "github.com/bakeable/bkry/internal/server/models/entities/author"
)


func afterQuery(entities []author.Author) []author.Author {
	return entities
}