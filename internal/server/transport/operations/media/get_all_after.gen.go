package media_operations

import (
	media "github.com/bakeable/bkry/internal/server/models/entities/media"
)


func afterGetAll(entities []media.Media) []media.Media {
	return entities
}