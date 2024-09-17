package quotation_operations

import (
	quotation "github.com/bakeable/bkry/data/entities/quotation"
	repo "github.com/bakeable/bkry/data/repository/entities"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func GetAll(repository repo.IRepository, ) []quotation.Quotation {
	// Get Quotations
	entities, err := repository.GetAllQuotations()
	if err != nil {
		panic(err)
	}

	// Process Quotation entities
	entities = afterGetAll(entities)

	// Return Quotations
	return entities
}