package examination_property_operations

import (
	examination_property "github.com/bakeable/bkry/data/entities/examination_property"
	repo "github.com/bakeable/bkry/data/repository/entities"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"	
	
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func GetAllPaginated(repository repo.IRepository, pagination datastore.Pagination) ([]examination_property.ExaminationProperty, datastore.Pagination) {
	// Get ExaminationProperties
	entities, pagination, err := repository.GetAllExaminationPropertiesPaginated(pagination)
	if err != nil {
		panic(err)
	}

	// Process ExaminationProperty entities
	entities = afterGetAllPaginated(entities, pagination.PageSize, pagination.PageNumber, pagination.OrderBy, pagination.Ascending)

	// Return ExaminationProperties
	return entities, pagination
}