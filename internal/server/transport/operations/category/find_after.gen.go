package category_operations

import (
	category "github.com/bakeable/bkry/internal/server/models/entities/category"
)

func afterFind(entity category.Category) category.Category {
	return entity
}