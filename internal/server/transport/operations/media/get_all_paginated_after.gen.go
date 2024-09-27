package media_operations

import (
	media "github.com/bakeable/bkry/internal/server/models/entities/media"
)


func afterGetAllPaginated(entities []media.Media, pageSize int, pageNumber int, orderBy string, ascending bool) []media.Media {
	return entities
}