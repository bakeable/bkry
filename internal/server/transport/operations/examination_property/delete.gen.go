package examination_property_operations

import (
	repo "github.com/bakeable/bkry/internal/server/database/repository"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func Delete(repository repo.IRepository, examinationPropertyID string) {
	// Perform before delete actions for ExaminationProperty entity
	beforeDelete(examinationPropertyID)

	// Delete ExaminationProperty
	err := repository.DeleteExaminationProperty(examinationPropertyID)
	if err != nil {
		panic(err)
	}

	// Perform after delete actions for ExaminationProperty entity
	afterDelete(examinationPropertyID)
}