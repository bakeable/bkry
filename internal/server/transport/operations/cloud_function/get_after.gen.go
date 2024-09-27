package cloud_function_operations

import (
	cloud_function "github.com/bakeable/bkry/internal/server/models/entities/cloud_function"
)

func afterGet(cloudFunctionID string, entity cloud_function.CloudFunction) cloud_function.CloudFunction {
	return entity
}