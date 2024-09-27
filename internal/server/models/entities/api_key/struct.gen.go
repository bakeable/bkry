package api_key
import "github.com/bakeable/bkry/internal/server/models/general/auditinfo"

type ApiKey struct {
	ID string `json:"id"`
	Modified auditinfo.AuditInfo `json:"modified"`
	Created auditinfo.AuditInfo `json:"created"`
	// The kind of the object
	Kind string `json:"_kind"`
	// The public key
	publicKey string `json:"publicKey"`
	// The secret key
	secretKey string `json:"secretKey"`
	// The user ID
	userId string `json:"userId"`
}








