package auditinfo

// Default AuditInfo struct
type AuditInfo struct {
	Timestamp int64 `json:"timestamp"`
	UserId string `json:"user_id"`
}