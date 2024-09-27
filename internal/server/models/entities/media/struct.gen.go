package media
import "github.com/bakeable/bkry/internal/server/models/general/auditinfo"

type Media struct {
	ID string `json:"id"`
	Modified auditinfo.AuditInfo `json:"modified"`
	Created auditinfo.AuditInfo `json:"created"`
	// The type of content: image, video, audio or file
	ContentType ContentType `json:"contentType"`
	// The description of the media
	Description string `json:"description"`
	// The file extension
	Extension string `json:"extension"`
	// The filename of the media
	Filename string `json:"filename"`
	// The official file mimetype
	MimeType string `json:"mimeType"`
	// The size of the file in Kb
	Size string `json:"size"`
	// The Firebase Storage location
	StoragePath string `json:"storagePath"`
	// The public URL to get media directly from the internet
	Url string `json:"url"`
}




type ContentType string




