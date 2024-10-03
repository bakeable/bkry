package category_localisation_operations

import (
	category_localisation "github.com/bakeable/bkry/internal/server/models/entities/category_localisation"
)


func afterGetAllPaginated(categoryID string, entities []category_localisation.CategoryLocalisation, pageSize int, pageNumber int, orderBy string, ascending bool) []category_localisation.CategoryLocalisation {
	return entities
}