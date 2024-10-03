package question_bundle
import "github.com/bakeable/bkry/internal/server/models/general/auditinfo"

type QuestionBundle struct {
	// The unique identifier for the QuestionBundle
	ID string `json:"id"`
	// The audit information for the modification of the QuestionBundle
	Modified auditinfo.AuditInfo `json:"modified"`
	// The audit information for the creation of the QuestionBundle
	Created auditinfo.AuditInfo `json:"created"`
	// How the QuestionBundle is comprised: Manual, ByCategory, or ByTag
	BundleType BundleType `json:"bundle_type"`
	// Optional further description of the contents of the bundle
	Description string `json:"description"`
	// The default language code for the question bundle
	LanguageID string `json:"language_id"`
	// A list of language id's in which the question bundle is available
	LanguageIds []string `json:"language_ids"`
	// The timestamp of the last export of the question bundle
	LastExportDateTime int `json:"last_export_timestamp"`
	// The list of QuestionContextId strings that are either added manually or queried when the QuestionBundle is saved
	QuestionContextIds []string `json:"question_context_ids"`
	// The current status of the question bundle
	Status Status `json:"status"`
	// The subtitle of the question bundle
	Subtitle string `json:"subtitle"`
	// The title of the question bundle
	Title string `json:"title"`
}






