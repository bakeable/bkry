package product_package_localisation
import (
	"github.com/bakeable/bkry/internal/server/models/general/docref"
   "github.com/bakeable/bkry/internal/server/models/general/auditinfo"
)

type ProductPackageLocalisation struct {
	// The unique identifier for the ProductPackageLocalisation
	ID string `json:"id"`
	// The audit information for the modification of the ProductPackageLocalisation
	Modified auditinfo.AuditInfo `json:"modified"`
	// The audit information for the creation of the ProductPackageLocalisation
	Created auditinfo.AuditInfo `json:"created"`
	// Optional further description of the contents of the package
	Description string `json:"description"`
	// The default language code for the package
	LanguageID string `json:"language_id"`
	// Reference to the media associated with the package
	Media docref.DocRef `json:"media"`
	// The ID of the parent entity for which this is the child document
	ProductPackageID string `json:"product_package_id"`
	// The subtitle of the package
	Subtitle string `json:"subtitle"`
	// The title of the package
	Title string `json:"title"`
}






