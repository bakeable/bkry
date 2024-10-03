package tag_localisation
import "github.com/bakeable/bkry/internal/server/models/general/auditinfo"

type TagLocalisation struct {
	// The unique identifier for the TagLocalisation
	ID string `json:"id"`
	// The audit information for the modification of the TagLocalisation
	Modified auditinfo.AuditInfo `json:"modified"`
	// The audit information for the creation of the TagLocalisation
	Created auditinfo.AuditInfo `json:"created"`
	// The label of the tag
	Label string `json:"label"`
	// The default language code for the tag
	LanguageID string `json:"language_id"`
	// The ID of the parent entity for which this is the child document
	TagID string `json:"tag_id"`
}






