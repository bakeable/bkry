package cloud_function_operations

import (
	cloud_function "github.com/bakeable/bkry/internal/server/models/entities/cloud_function"
)

func beforeSave(entity cloud_function.CloudFunction, editorID *string) cloud_function.CloudFunction {
	// Return CloudFunction
	return entity
}