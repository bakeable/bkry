package tag
import "github.com/bakeable/bkry/internal/server/models/general/auditinfo"

type Tag struct {
	// The unique identifier for the Tag
	ID string `json:"id"`
	// The audit information for the modification of the Tag
	Modified auditinfo.AuditInfo `json:"modified"`
	// The audit information for the creation of the Tag
	Created auditinfo.AuditInfo `json:"created"`
	// The label of the tag
	Label string `json:"label"`
	// The default language code for the tag
	LanguageID string `json:"language_id"`
	// A list of language id's in which the question bundle is available
	LanguageIds []string `json:"language_ids"`
}






