package cloud_function_operations

import (
	cloud_function "github.com/bakeable/bkry/internal/server/models/entities/cloud_function"
)


func afterGetAllPaginated(entities []cloud_function.CloudFunction, pageSize int, pageNumber int, orderBy string, ascending bool) []cloud_function.CloudFunction {
	return entities
}