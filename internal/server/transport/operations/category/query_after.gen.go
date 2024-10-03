package category_operations

import (
	category "github.com/bakeable/bkry/internal/server/models/entities/category"
)


func afterQuery(entities []category.Category) []category.Category {
	return entities
}