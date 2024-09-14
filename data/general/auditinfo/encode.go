package auditinfo

// Convert Key struct to database map
func (x AuditInfo) Encode() map[string]interface{} {
	return map[string]interface{}{
		"timestamp": x.Timestamp,
		"user_id": x.UserId,
	}
}