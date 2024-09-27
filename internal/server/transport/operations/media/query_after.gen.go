package media_operations

import (
	media "github.com/bakeable/bkry/internal/server/models/entities/media"
)


func afterQuery(entities []media.Media) []media.Media {
	return entities
}