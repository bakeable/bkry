package auditinfo

import utils "github.com/bakeable/bkry/tools"

// Decode a database map to an AuditInfo struct
func Decode(data interface{}) AuditInfo {
	if data == nil {
		return AuditInfo{}
	}
	if val, ok := data.(map[string]interface{}); ok {
		return AuditInfo{
			Timestamp: int64(utils.DecodeInt(val["timestamp"], 0)),
			UserId:    utils.DecodeString(val["user_id"], ""),
		}
	}
	return AuditInfo{}
}