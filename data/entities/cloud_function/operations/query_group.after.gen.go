package cloud_function_operations

import (
	cloud_function "github.com/bakeable/bkry/data/entities/cloud_function"
)


func afterQueryGroup(entities []cloud_function.CloudFunction) []cloud_function.CloudFunction {
	return entities
}