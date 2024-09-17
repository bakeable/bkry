package examination_property_operations

import (
	examination_property "github.com/bakeable/bkry/data/entities/examination_property"
	repo "github.com/bakeable/bkry/data/repository/entities"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func GetAll(repository repo.IRepository, ) []examination_property.ExaminationProperty {
	// Get ExaminationProperties
	entities, err := repository.GetAllExaminationProperties()
	if err != nil {
		panic(err)
	}

	// Process ExaminationProperty entities
	entities = afterGetAll(entities)

	// Return ExaminationProperties
	return entities
}