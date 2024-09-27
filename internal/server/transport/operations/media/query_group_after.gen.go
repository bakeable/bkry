package media_operations

import (
	media "github.com/bakeable/bkry/internal/server/models/entities/media"
)


func afterQueryGroup(entities []media.Media) []media.Media {
	return entities
}