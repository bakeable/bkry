package auditinfo

import "fmt"

// Get the formatted date
func (x AuditInfo) ToString() string {
	return fmt.Sprintf("Audited at timestamp: %d, by user: %s", x.Timestamp, x.UserId)
}