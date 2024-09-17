package media_operations

import (
	media "github.com/bakeable/bkry/data/entities/media"
)


func afterGetAllPaginated(entities []media.Media, pageSize int, pageNumber int, orderBy string, ascending bool) []media.Media {
	return entities
}