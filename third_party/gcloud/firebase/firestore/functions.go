package firestore

import (
	"time"
)

func audit(userId *string) map[string]interface{} {
	auditInfo := map[string]interface{}{
		"timestamp": time.Now().Unix(),
		"user_id": userId,
	}

	if userId == nil {
		auditInfo["user_id"] = "system"
	} else {
		auditInfo["user_id"] = *userId
	}

	return auditInfo
}