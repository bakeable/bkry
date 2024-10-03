package product_package
import (
	"github.com/bakeable/bkry/internal/server/models/general/docref"
   "github.com/bakeable/bkry/internal/server/models/general/auditinfo"
)

type ProductPackage struct {
	// The unique identifier for the ProductPackage
	ID string `json:"id"`
	// The audit information for the modification of the ProductPackage
	Modified auditinfo.AuditInfo `json:"modified"`
	// The audit information for the creation of the ProductPackage
	Created auditinfo.AuditInfo `json:"created"`
	// Optional further description of the contents of the package
	Description string `json:"description"`
	// The items included in the package
	Items []Item `json:"items"`
	// The default language code for the package
	LanguageID string `json:"language_id"`
	// A list of language id's in which the question bundle is available
	LanguageIds []string `json:"language_ids"`
	// Reference to the media associated with the package
	Media docref.DocRef `json:"media"`
	// The subtitle of the package
	Subtitle string `json:"subtitle"`
	// The title of the package
	Title string `json:"title"`
}




type Item struct {
	// The entity type that is sold as an item of the package
	EntityType EntityType `json:"entity_type"`
	// The Id of the entity being sold
	ID string `json:"id"`
}



