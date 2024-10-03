package author_operations

import (
	author "github.com/bakeable/bkry/internal/server/models/entities/author"
)


func afterGetAllPaginated(entities []author.Author, pageSize int, pageNumber int, orderBy string, ascending bool) []author.Author {
	return entities
}