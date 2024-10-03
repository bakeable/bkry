package question_bundle_localisation
import "github.com/bakeable/bkry/internal/server/models/general/auditinfo"

type QuestionBundleLocalisation struct {
	// The unique identifier for the QuestionBundleLocalisation
	ID string `json:"id"`
	// The audit information for the modification of the QuestionBundleLocalisation
	Modified auditinfo.AuditInfo `json:"modified"`
	// The audit information for the creation of the QuestionBundleLocalisation
	Created auditinfo.AuditInfo `json:"created"`
	// Optional further description of the contents of the bundle
	Description string `json:"description"`
	// The default language code for the question bundle
	LanguageID string `json:"language_id"`
	// The ID of the parent entity for which this is the child document
	QuestionBundleID string `json:"question_bundle_id"`
	// The current status of the question bundle
	Status Status `json:"status"`
	// The subtitle of the question bundle
	Subtitle string `json:"subtitle"`
	// The title of the question bundle
	Title string `json:"title"`
}






