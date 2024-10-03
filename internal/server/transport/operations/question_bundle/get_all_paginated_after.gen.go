package question_bundle_operations

import (
	question_bundle "github.com/bakeable/bkry/internal/server/models/entities/question_bundle"
)


func afterGetAllPaginated(entities []question_bundle.QuestionBundle, pageSize int, pageNumber int, orderBy string, ascending bool) []question_bundle.QuestionBundle {
	return entities
}