package category_localisation
import "github.com/bakeable/bkry/internal/server/models/general/auditinfo"

type CategoryLocalisation struct {
	// The unique identifier for the CategoryLocalisation
	ID string `json:"id"`
	// The audit information for the modification of the CategoryLocalisation
	Modified auditinfo.AuditInfo `json:"modified"`
	// The audit information for the creation of the CategoryLocalisation
	Created auditinfo.AuditInfo `json:"created"`
	// The ID of the parent entity for which this is the child document
	CategoryID string `json:"category_id"`
	// The label of the category
	Label string `json:"label"`
	// The default language code for the category
	LanguageID string `json:"language_id"`
}






