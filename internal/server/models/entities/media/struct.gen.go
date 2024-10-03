package media
import (
	"github.com/bakeable/bkry/internal/server/models/general/docref"
   "github.com/bakeable/bkry/internal/server/models/general/auditinfo"
)

type Media struct {
	// The unique identifier for the Media
	ID string `json:"id"`
	// The audit information for the modification of the Media
	Modified auditinfo.AuditInfo `json:"modified"`
	// The audit information for the creation of the Media
	Created auditinfo.AuditInfo `json:"created"`
	// DocumentReference to the Author
	Author docref.DocRef `json:"author"`
	// The type of content: Image, Video, or Audio
	ContentType ContentType `json:"content_type"`
	// The description of the media file
	Description Description `json:"description"`
	// The file extension
	Extension string `json:"extension"`
	// The array of language codes to which this media applies
	LanguageIDs []string `json:"language_ids"`
	// The official file mimetype
	MimeType string `json:"mime_type"`
	// Indicates if the media file is multilingual
	Multilingual bool `json:"multilingual"`
	// The size of the file in Kb
	Size string `json:"size"`
	// The Firebase Storage location
	StoragePath string `json:"storage_path"`
	// The public URL to get media directly from the internet
	Url string `json:"url"`
}




type Description struct {
	// The default locale of the description
	DefaultLocale string `json:"default_locale"`
	// The default description of the Media file
	DefaultValue string `json:"default_value"`
	// A mapping containing the LanguageCode => Description values
	Locales map[string]string `json:"locales"`
}



