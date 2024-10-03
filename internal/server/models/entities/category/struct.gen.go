package category
import "github.com/bakeable/bkry/internal/server/models/general/auditinfo"

type Category struct {
	// The unique identifier for the Category
	ID string `json:"id"`
	// The audit information for the modification of the Category
	Modified auditinfo.AuditInfo `json:"modified"`
	// The audit information for the creation of the Category
	Created auditinfo.AuditInfo `json:"created"`
	// The label of the category
	Label string `json:"label"`
	// The default language code for the category
	LanguageID string `json:"language_id"`
	// A list of language id's in which the question bundle is available
	LanguageIds []string `json:"language_ids"`
	// The amount of questions in this category
	QuestionCount int `json:"question_count"`
}






