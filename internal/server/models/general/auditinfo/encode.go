package auditinfo

// Convert Key struct to database map
func (x AuditInfo) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"timestamp": x.Timestamp,
		"user_id": x.UserId,
	}
}