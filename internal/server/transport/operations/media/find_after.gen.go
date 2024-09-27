package media_operations

import (
	media "github.com/bakeable/bkry/internal/server/models/entities/media"
)

func afterFind(entity media.Media) media.Media {
	return entity
}