package category_operations

import (
	category "github.com/bakeable/bkry/internal/server/models/entities/category"
)


func afterGetAllPaginated(entities []category.Category, pageSize int, pageNumber int, orderBy string, ascending bool) []category.Category {
	return entities
}